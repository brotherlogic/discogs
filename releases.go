package discogs

import (
	"context"
	"encoding/json"
	"fmt"
)

type Rating struct {
	Rating int
}

type RatingResponse struct{}

func (p *prodClient) SetRating(ctx context.Context, releaseid int64, rating int32) error {
	url := fmt.Sprintf("/releases/%v/rating/%v", releaseid, p.user.GetUsername())
	data := &Rating{
		Rating: int(rating),
	}
	v, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp := &RatingResponse{}
	return p.makeDiscogsRequest("PUT", url, string(v), "/releases/rid/rating/uname", resp)
}
