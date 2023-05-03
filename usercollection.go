package discogs

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/discogs/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Pagination struct {
	Pages int `json:"pages"`
	Page  int `json:"page"`
}

type CollectionResponse struct {
	Pagination Pagination
	Releases   []CollectionRelease
	Message    string
}

type CollectionRelease struct {
	Id               int
	InstanceId       int `json:"instance_id"`
	FolderId         int `json:"folder_id"`
	Rating           int
	BasicInformation BasicInformation `json:"basic_information"`
}

type BasicInformation struct {
	Formats []Format
}

type Format struct {
	Descriptions []string
	Name         string
	Qty          string
}

func (d *prodClient) GetCollection(ctx context.Context, page int32) ([]*pb.Release, *pb.Pagination, error) {
	cr := &CollectionResponse{}
	err := d.makeDiscogsRequest("GET", fmt.Sprintf("/users/%v/collection/folders/0/releases?page=%v", d.user.GetUsername(), page), "", cr)
	if err != nil {
		return nil, nil, err
	}

	if len(cr.Message) > 0 {
		if strings.Contains(cr.Message, "is outside of valid range") {
			return nil, nil, status.Errorf(codes.OutOfRange, cr.Message)
		}
	}

	var rs []*pb.Release
	for _, release := range cr.Releases {
		r := &pb.Release{
			InstanceId: int64(release.InstanceId),
			Id:         int64(release.Id),
			FolderId:   int32(release.FolderId),
			Rating:     int32(release.Rating),
		}

		var formats []*pb.Format
		for _, form := range release.BasicInformation.Formats {
			val, _ := strconv.ParseInt(form.Qty, 10, 32)
			formats = append(formats, &pb.Format{
				Name:         form.Name,
				Descriptions: form.Descriptions,
				Quantity:     int32(val),
			})
		}
		r.Formats = formats

		rs = append(rs, r)
	}

	return rs, &pb.Pagination{Page: int32(cr.Pagination.Page), Pages: int32(cr.Pagination.Pages)}, nil
}
