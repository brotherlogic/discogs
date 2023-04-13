package discogs

import (
	"context"
	"fmt"

	pb "github.com/brotherlogic/discogs/proto"
)

type Pagination struct {
	pages int
	page  int
}

type CollectionResponse struct {
	Pagination Pagination
	Releases   []CollectionRelease
}

type CollectionRelease struct {
	Id         int
	InstanceId int `json:"instance_id"`
	FolderId   int `json:"folder_id"`
	Rating     int `json:"rating"`
}

func (d *Discogs) GetCollection(ctx context.Context, user *pb.User, page int32) ([]*pb.Release, *pb.Pagination, error) {
	cr := &CollectionResponse{}
	err := d.makeDiscogsRequest("GET", fmt.Sprintf("/users/%v/collection/folders/0/releases?page=%v", user.GetUsername(), page), "", cr)
	if err != nil {
		return nil, nil, err
	}

	var rs []*pb.Release
	for _, release := range cr.Releases {
		rs = append(rs, &pb.Release{
			InstanceId: int64(release.InstanceId),
			Id:         int64(release.Id),
			FolderId:   int32(release.FolderId),
			Rating:     int32(release.Rating),
		})
	}

	return rs, &pb.Pagination{Page: int32(cr.Pagination.page), Pages: int32(cr.Pagination.pages)}, nil
}
