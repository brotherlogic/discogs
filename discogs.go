package discogs

import (
	"context"
	"io"
	"net/http"

	pb "github.com/brotherlogic/discogs/proto"
	"github.com/dghubble/oauth1"
)

type Discogs interface {
	GetLoginURL() (string, string, string, error)
	HandleDiscogsResponse(ctx context.Context, secret, token, verifier string) (string, string, error)

	GetDiscogsUser(ctx context.Context) (*pb.User, error)

	GetCollection(ctx context.Context, page int32) ([]*pb.Release, *pb.Pagination, error)

	GetUserId() int32
}

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
	}
}

func (p *prodClient) ForUser(token, secret string) Discogs {
	return &prodClient{
		getter: &oauthGetter{key: token, secret: secret,
			conf: p.getter.config(),
		},
	}
}
