package discogs

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	pb "github.com/brotherlogic/discogs/proto"
	"github.com/dghubble/oauth1"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Option for doing remote downloads
type Downloader interface {
	Download(ctx context.Context, url string) (string, error)
}

const (
	THROTTLE_REQUESTS = 60
	THROTTLE_WINDOW   = time.Minute
)

type prodClient struct {
	secret   string
	key      string
	callback string

	getter        clientGetter
	user          *pb.User
	personalToken string

	requestTimes []time.Time

	downloader Downloader
}

func (d *prodClient) Throttle() {
	// Clean the request Times
	var nrt []time.Time
	for _, rt := range d.requestTimes {
		if time.Since(rt) < time.Minute {
			nrt = append(nrt, rt)
		}
	}
	d.requestTimes = nrt

	if len(d.requestTimes) > THROTTLE_REQUESTS {
		time.Sleep(time.Minute - time.Since(d.requestTimes[0]))
	}
}

func (d *prodClient) GetUserId() int32 {
	return d.user.GetDiscogsUserId()
}

type clientGetter interface {
	get() myClient
	getDefault() myClient
	config() oauth1.Config
}

type oauthGetter struct {
	key    string
	secret string
	conf   oauth1.Config
}

func (o *oauthGetter) getDefault() myClient {
	return http.DefaultClient
}

func (o *oauthGetter) get() myClient {
	oauthToken := oauth1.NewToken(o.key, o.secret)
	return o.conf.Client(oauth1.NoContext, oauthToken)
}

func (o *oauthGetter) config() oauth1.Config {
	return o.conf
}

type myClient interface {
	Get(url string) (resp *http.Response, err error)
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
	Do(req *http.Request) (*http.Response, error)
}

func DiscogsWithAuth(key, secret, callback string) Discogs {
	return &prodClient{
		key:      key,
		secret:   secret,
		callback: callback,
		getter: &oauthGetter{conf: oauth1.Config{
			ConsumerKey:    key,
			ConsumerSecret: secret,
			CallbackURL:    callback,
			Endpoint:       authenticateEndpoint,
		}},
	}
}

func (p *prodClient) SetDownloader(dl Downloader) {
	p.downloader = dl
}

func (p *prodClient) ForUser(user *pb.User) Discogs {
	return &prodClient{
		key:      p.key,
		secret:   p.secret,
		callback: p.callback,
		getter: &oauthGetter{key: user.GetUserToken(), secret: user.GetUserSecret(),
			conf: p.getter.config(),
		},
		personalToken: user.GetPersonalToken(),
		user:          user,
		downloader:    p.downloader,
	}
}

var (
	requestCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "discogs_requests",
	}, []string{"request_type", "endpoint", "response", "response_code"})
)

func (d *prodClient) makeDiscogsRequest(rtype, path string, data string, ep string, obj interface{}) error {
	if rtype == "SGET" {
		if d.downloader != nil {
			log.Printf("Running SGET with specified downloader")
			data, err := d.downloader.Download(context.Background(), path)
			if err != nil {
				return err
			}

			nobj := obj.(*strpass)
			nobj.Value = data
			return nil
		}

		log.Printf("Running SGET with local downloader")
		httpClient := d.getter.getDefault()
		resp, err := httpClient.Get(path)
		if err != nil {
			return err
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if resp.StatusCode != 200 && resp.StatusCode != 201 {
			return fmt.Errorf("%v: %v", resp.StatusCode, string(body))
		}

		log.Printf("READ %v", len(string(body)))
		nobj := obj.(*strpass)
		nobj.Value = string(body)
		return nil
	}

	if !strings.HasPrefix(path, "/") {
		return status.Errorf(codes.FailedPrecondition, "Path needs to start with / :'%v'", path)
	}

	fullPath := fmt.Sprintf("https://api.discogs.com%v", path)

	httpClient := d.getter.get()
	var resp *http.Response
	var err error

	// Setup for personal token if we have it listed
	if d.personalToken != "" {
		httpClient = d.getter.getDefault()
		if strings.Contains(fullPath, "?") {
			fullPath = fmt.Sprintf("%v&token=%v", fullPath, d.personalToken)
		} else {
			fullPath = fmt.Sprintf("%v?token=%v", fullPath, d.personalToken)
		}
	}

	log.Printf("DISCOGS_REQUEST %v:%v", rtype, fullPath)

	switch rtype {
	case "POST":
		resp, err = httpClient.Post(fullPath, "application/json", bytes.NewBuffer([]byte(data)))
	case "GET":
		resp, err = httpClient.Get(fullPath)
	case "PUT":
		req, _ := http.NewRequest("PUT", fullPath, bytes.NewBuffer([]byte(data)))
		req.Header.Set("Content-Type", "application/json")
		resp, err = httpClient.Do(req)
	case "DELETE":
		req, _ := http.NewRequest("DELETE", fullPath, bytes.NewBuffer([]byte("")))
		req.Header.Set("Content-Type", "application/json")
		resp, err = httpClient.Do(req)
	default:
		return fmt.Errorf("Unable to handle %v requests", rtype)
	}
	if err != nil {
		requestCounter.With(prometheus.Labels{"request_type": "POST", "endpoint": ep, "response": fmt.Sprintf("%v", status.Code(err)), "response_code": "-1"})
		return err
	}
	requestCounter.With(prometheus.Labels{"request_type": "POST", "endpoint": ep, "response": fmt.Sprintf("%v", status.Code(err)), "response_code": fmt.Sprintf("%v", resp.StatusCode)})

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Not Found
	if resp.StatusCode == 404 {
		return status.Errorf(codes.NotFound, string(body))
	}

	// Throttling
	if resp.StatusCode == 429 {
		return status.Errorf(codes.ResourceExhausted, "Discogs is throttling us")
	}

	// Permission denied
	if resp.StatusCode == 403 {
		return status.Errorf(codes.PermissionDenied, string(body))
	}

	// 422 Unprocessable
	if resp.StatusCode == 422 {
		return status.Errorf(codes.FailedPrecondition, string(body))
	}

	if resp.StatusCode != 200 && resp.StatusCode != 204 && resp.StatusCode != 201 {
		return status.Errorf(codes.Unknown, "Unknown response code: %v with body %v", resp.StatusCode, string(body))
	}

	if len(body) > 0 {
		err = json.Unmarshal(body, obj)
		if err != nil {
			return fmt.Errorf("unmarshal error (processing %v): %v from %v", err, string(body), data)
		}
	}
	return nil
}
