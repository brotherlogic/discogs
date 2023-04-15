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

	GetCollection(ctx context.Context, user *pb.User, page int32) ([]*pb.Release, *pb.Pagination, error)
}

type prodClient struct {
	secret   string
	key      string
	callback string

	getter clientGetter
}

type clientGetter interface {
	get() myClient
}

type oauthGetter struct {
	key    string
	secret string
	config oauth1.Config
}

func (o *oauthGetter) get() myClient {
	oauthToken := oauth1.NewToken(o.key, o.secret)
	return o.config.Client(oauth1.NoContext, oauthToken)
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

func DiscogsWithToken(token, secret string, consumerKey, consumerToken, callback string) Discogs {
	return &prodClient{
		getter: &oauthGetter{key: token, secret: secret,
			config: oauth1.Config{
				ConsumerKey:    consumerKey,
				ConsumerSecret: consumerToken,
				CallbackURL:    callback,
				Endpoint:       authenticateEndpoint},
		},
	}
}
