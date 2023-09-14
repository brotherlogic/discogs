package discogs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	pb "github.com/brotherlogic/discogs/proto"
	"github.com/dghubble/oauth1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type prodClient struct {
	secret   string
	key      string
	callback string

	getter clientGetter
	user   *pb.User
}

func (d *prodClient) GetUserId() int32 {
	return d.user.GetDiscogsUserId()
}

type clientGetter interface {
	get() myClient
	config() oauth1.Config
}

type oauthGetter struct {
	key    string
	secret string
	conf   oauth1.Config
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

func (p *prodClient) ForUser(user *pb.User) Discogs {
	return &prodClient{
		key:      p.key,
		secret:   p.secret,
		callback: p.callback,
		getter: &oauthGetter{key: user.GetUserToken(), secret: user.GetUserSecret(),
			conf: p.getter.config(),
		},
		user: user,
	}
}

func (d *prodClient) makeDiscogsRequest(rtype, path string, data string, obj interface{}) error {
	if !strings.HasPrefix(path, "/") {
		return status.Errorf(codes.FailedPrecondition, "Path needs to start with / :'%v'", path)
	}
	fullPath := fmt.Sprintf("https://api.discogs.com%v", path)
	httpClient := d.getter.get()
	var resp *http.Response
	var err error

	switch rtype {
	case "POST":
		resp, err = httpClient.Post(fullPath, "application/json", bytes.NewBuffer([]byte(data)))
	case "GET":
		resp, err = httpClient.Get(fullPath)
	default:
		return fmt.Errorf("Unable to handle %v requests", rtype)
	}
	if err != nil {
		return err
	}

	// Throttling
	if resp.StatusCode == 429 {
		return status.Errorf(codes.ResourceExhausted, "Discogs is throttling us")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if len(body) > 0 {
		err = json.Unmarshal(body, obj)
		if err != nil {
			return fmt.Errorf("Unarshal error (processing %v): %v from %v", err, string(body), data)
		}
	}
	return nil
}
