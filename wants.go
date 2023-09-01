package discogs

import (
	"context"
	"fmt"

	pb "github.com/brotherlogic/discogs/proto"
)

type Artist struct {
	Name string
}

type Want struct {
	Title   string
	Id      int
	Artists []Artist
}
type GetWantsResponse struct {
	Pagination Pagination
	Wants      []Want
}

func (p *prodClient) GetWants(ctx context.Context, page int32) ([]*pb.Want, *pb.Pagination, error) {
	cr := &GetWantsResponse{}
	err := p.makeDiscogsRequest("GET", fmt.Sprintf("/users/%v/wants?page=%v", p.user.GetUsername(), page), "", cr)
	if err != nil {
		return nil, nil, err
	}

	var wants []*pb.Want
	for _, want := range cr.Wants {
		var artists []*pb.Artist
		for _, artist := range want.Artists {
			artists = append(artists, &pb.Artist{
				Name: artist.Name,
			})
		}
		wants = append(wants, &pb.Want{
			Title:   want.Title,
			Id:      int64(want.Id),
			Artists: artists,
		})
	}

	return wants, &pb.Pagination{Page: int32(cr.Pagination.Page), Pages: int32(cr.Pagination.Pages)}, nil
}
