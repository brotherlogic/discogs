package discogs

import (
	"context"

	pb "github.com/brotherlogic/discogs/proto"
)

type DiscogsUser struct {
	AvatarURL    string `json:"avatar_url"`
	ID           int32  `json:"id"`
	Username     string `json:"username"`
	CurrencyAbbr string `json:"curr_abbr"`
}

func (d *prodClient) GetDiscogsUser(ctx context.Context) (*pb.User, error) {
	user := &DiscogsUser{}
	err := d.makeDiscogsRequest(
		"GET",
		"/oauth/identity",
		"",
		"/oauth/identity",
		user,
	)
	return &pb.User{
		Username:      user.Username,
		DiscogsUserId: user.ID,
	}, err
}
