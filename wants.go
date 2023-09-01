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

type AddWantResponse struct {
	Id               int
	BasicInformation BasicInformation `json:"basic_information"`
}

func (p *prodClient) AddWant(ctx context.Context, releaseId int64) (*pb.Want, error) {
	cr := &AddWantResponse{}
	err := p.makeDiscogsRequest("PUT", fmt.Sprintf("/users/%v/wants/%v", p.user.GetUsername(), releaseId), "", cr)
	if err != nil {
		return nil, err
	}

	return &pb.Want{
		Id:    int64(cr.Id),
		Title: cr.BasicInformation.Title,
	}, nil
}

func (p *prodClient) DeleteWant(ctx context.Context, wantId int64) error {
	cr := &AddWantResponse{}
	return p.makeDiscogsRequest("DELETE", fmt.Sprintf("/users/%v/wants/%v", p.user.GetUsername(), wantId), "", cr)
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
