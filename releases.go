package discogs

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

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
	url := fmt.Sprintf("/releases/%v", releaseId)

	resp := &IndividualRelease{}
	err := p.makeDiscogsRequest("GET", url, "", "/releases/rid/", resp)
	if err != nil {
		return nil, err
	}

	r := &pb.Release{
		InstanceId: int64(resp.InstanceId),
		Id:         int64(resp.Id),
		FolderId:   int32(resp.FolderId),
		Rating:     int32(resp.Rating),
		Title:      resp.Title,
	}

	var formats []*pb.Format
	for _, form := range resp.Formats {
		val, _ := strconv.ParseInt(form.Qty, 10, 32)
		formats = append(formats, &pb.Format{
			Name:         form.Name,
			Descriptions: form.Descriptions,
			Quantity:     int32(val),
		})
	}
	r.Formats = formats

	var labels []*pb.Label
	for _, label := range resp.Labels {
		labels = append(labels, &pb.Label{
			Name:  label.Name,
			Catno: label.Catno,
			Id:    int32(label.Id),
		})
	}
	r.Labels = labels

	var artists []*pb.Artist
	for _, artist := range resp.Artists {
		artists = append(artists, &pb.Artist{
			Name: artist.Name,
			Id:   int64(artist.Id),
		})
	}
	r.Artists = artists

	switch strings.Count(resp.Released, "-") {
	case 2:
		if strings.HasSuffix(resp.Released, "-00") {
			rd, err := time.Parse("2006-01-00", resp.Released)
			if err != nil {
				return nil, fmt.Errorf("unable to parse %v -> %v", resp.Released, err)
			}
			r.ReleaseDate = rd.Unix()
		} else {
			rd, err := time.Parse("2006-01-02", resp.Released)
			if err != nil {
				return nil, fmt.Errorf("unable to parse %v -> %v", resp.Released, err)
			}
			r.ReleaseDate = rd.Unix()
		}
	case 1:
		rd, err := time.Parse("2006-01", resp.Released)
		if err != nil {
			return nil, fmt.Errorf("unable to parse %v -> %v", resp.Released, err)
		}
		r.ReleaseDate = rd.Unix()
	case 0:
		rd, err := time.Parse("2006", resp.Released)
		if err != nil {
			return nil, fmt.Errorf("unable to parse %v -> %v", resp.Released, err)
		}
		r.ReleaseDate = rd.Unix()
	}

	return r, nil
}
