package discogs

import (
	"context"
	"encoding/json"
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
	Title            string
	Notes            []Note
	BasicInformation BasicInformation `json:"basic_information"`
}

type Note struct {
	FieldId int `json:"field_id"`
	Value   string
}

type BasicInformation struct {
	Formats []Format
	Labels  []Label
	Title   string
	Artists []Artist
}

type Label struct {
	Name  string
	Catno string
	Id    int
}

type Format struct {
	Descriptions []string
	Name         string
	Qty          string
}

type CreateFolderResponse struct {
	Id   int
	Name string
}

type DeleteFolderResponse struct{}

func (d *prodClient) DeleteFolder(ctx context.Context, folderId int32) error {
	return d.makeDiscogsRequest("DELETE", fmt.Sprintf("/users/%v/collection/folders/%v",
		d.user.GetUsername(), folderId), "", "/users/uname/collection/folders/fnum", &DeleteFolderResponse{})

}

func (d *prodClient) CreateFolder(ctx context.Context, folderName string) (*pb.Folder, error) {
	data := &Folder{
		Name: folderName,
	}
	v, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	cfr := &CreateFolderResponse{}
	err = d.makeDiscogsRequest("POST", fmt.Sprintf("/users/%v/collection/folders",
		d.user.GetUsername()), string(v), "/users/uname/collection/folders", cfr)
	if err != nil {
		return nil, err
	}
	return &pb.Folder{
		Name: cfr.Name,
		Id:   int32(cfr.Id),
	}, nil
}

func (d *prodClient) GetCollection(ctx context.Context, page int32) ([]*pb.Release, *pb.Pagination, error) {
	cr := &CollectionResponse{}
	err := d.makeDiscogsRequest("GET", fmt.Sprintf("/users/%v/collection/folders/0/releases?page=%v", d.user.GetUsername(), page), "", "/users/uname/collection/folders/0/releases", cr)
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
			Title:      release.BasicInformation.Title,
			Notes:      make(map[int32]string),
		}

		for _, note := range release.Notes {
			r.Notes[int32(note.FieldId)] = note.Value
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

		var labels []*pb.Label
		for _, label := range release.BasicInformation.Labels {
			labels = append(labels, &pb.Label{
				Name:  label.Name,
				Catno: label.Catno,
				Id:    int32(label.Id),
			})
		}
		r.Labels = labels

		var artists []*pb.Artist
		for _, artist := range release.BasicInformation.Artists {
			artists = append(artists, &pb.Artist{
				Name: artist.Name,
				Id:   int64(artist.Id),
			})
		}
		r.Artists = artists

		rs = append(rs, r)
	}

	return rs, &pb.Pagination{Page: int32(cr.Pagination.Page), Pages: int32(cr.Pagination.Pages)}, nil
}
