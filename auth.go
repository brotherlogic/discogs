package discogs

import (
	"context"

	"github.com/dghubble/oauth1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var authenticateEndpoint = oauth1.Endpoint{
	RequestTokenURL: "https://api.discogs.com/oauth/request_token",
	AuthorizeURL:    "https://www.discogs.com/oauth/authorize",
	AccessTokenURL:  "https://api.discogs.com/oauth/access_token",
}

// GenerateLoginURL generates a URL to login by
func (d *prodClient) GetLoginURL() (string, string, string, error) {
	if d.callback == "" {
		return "", "", "", status.Errorf(codes.FailedPrecondition, "unable to get login url without api params (missing callback)")
	}
	config :=
		&oauth1.Config{
			ConsumerKey:    d.key,
			ConsumerSecret: d.secret,
			CallbackURL:    d.callback,
			Endpoint:       authenticateEndpoint,
		}

	requestToken, secret, err := config.RequestToken()
	if err != nil {
		return "", "", "", err
	}
	authorizationURL, err := config.AuthorizationURL(requestToken)
	if err != nil {
		return "", "", "", err
	}
	return authorizationURL.String(), requestToken, secret, nil
}

// HandleDiscogsResponse handles the response from discogs on the last oauth step
func (d *prodClient) HandleDiscogsResponse(ctx context.Context, secret, token, verifier string) (string, string, error) {
	config :=
		&oauth1.Config{
			ConsumerKey:    d.key,
			ConsumerSecret: d.secret,
			CallbackURL:    d.callback,
			Endpoint:       authenticateEndpoint,
		}

	return config.AccessToken(token, secret, verifier)
}
