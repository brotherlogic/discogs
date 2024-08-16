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
	Rating int `json:"rating"`
}

type RatingResponse struct{}

type MasterReleasesResponse struct {
	Versions []MasterRelease
}

type MasterRelease struct {
	Id       int
	Released string
}

func (p *prodClient) GetMasterReleases(ctx context.Context, masterId int64, page int32, sort pb.MasterSort) ([]*pb.MasterRelease, error) {
	url := fmt.Sprintf("/masters/%v/versions?page=%v&per_page=100", masterId, page)
	switch sort {
	case pb.MasterSort_BY_YEAR:
		url += "&sort=released"
	}

	resp := &MasterReleasesResponse{}
	err := p.makeDiscogsRequest("GET", url, "", "/masters/mid/versions", resp)
	if err != nil {
		return nil, fmt.Errorf("unable to get mids: %w", err)
	}

	var vals []*pb.MasterRelease
	for _, version := range resp.Versions {
		conv, err := strconv.ParseInt(version.Released, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("unable to parse year: %w", err)
		}
		vals = append(vals, &pb.MasterRelease{
			Id:   int64(version.Id),
			Year: int32(conv),
		})
	}

	return vals, nil
}

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
		return nil, fmt.Errorf("unable to get release: %w", err)
	}

	r := &pb.Release{
		InstanceId: int64(resp.InstanceId),
		Id:         int64(resp.Id),
		FolderId:   int32(resp.FolderId),
		Rating:     int32(resp.Rating),
		Title:      resp.Title,
		MasterId:   int64(resp.MasterId),
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
		if strings.HasSuffix(resp.Released, "00-00") {
			rd, err := time.Parse("2006-00-00", resp.Released)
			if err != nil {
				return nil, fmt.Errorf("5 unable to parse %v -> %v", resp.Released, err)
			}
			r.ReleaseDate = rd.Unix()
		} else if strings.HasSuffix(resp.Released, "-00") {
			rd, err := time.Parse("2006-01-00", resp.Released)
			if err != nil {
				return nil, fmt.Errorf("2 unable to parse %v -> %v", resp.Released, err)
			}
			r.ReleaseDate = rd.Unix()
		} else {
			rd, err := time.Parse("2006-01-02", resp.Released)
			if err != nil {
				return nil, fmt.Errorf("2 unable to parse %v -> %v", resp.Released, err)
			}
			r.ReleaseDate = rd.Unix()
		}
	case 1:
		rd, err := time.Parse("2006-01", resp.Released)
		if err != nil {
			return nil, fmt.Errorf("1 unable to parse %v -> %v", resp.Released, err)
		}
		r.ReleaseDate = rd.Unix()
	case 0:
		if resp.Released == "" {
			r.ReleaseDate = 0
		} else {

			rd, err := time.Parse("2006", resp.Released)
			if err != nil {
				return nil, fmt.Errorf("0 unable to parse %v -> %v", resp.Released, err)
			}
			r.ReleaseDate = rd.Unix()
		}
	}

	return r, nil
}
