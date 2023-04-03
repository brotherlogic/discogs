package discogs

import (
	"io"
	"net/http"

	"github.com/dghubble/oauth1"
)

type Discogs struct {
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

func DiscogsWithAuth(key, secret, callback string) *Discogs {
	return &Discogs{
		key:      key,
		secret:   secret,
		callback: callback,
	}
}

func DiscogsWithToken(token, secret string) *Discogs {
	return &Discogs{
		getter: &oauthGetter{key: token, secret: secret},
	}
}
