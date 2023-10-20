package discogs

import (
	"context"
	"encoding/json"
	"fmt"

	pb "github.com/brotherlogic/discogs/proto"
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

func (p *prodClient) GetRelease(ctx context.Context, releaseId int64) (*pb.Release, error) {
	return &pb.Release{}, nil
}
