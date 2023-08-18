package discogs

import (
	"context"
	"fmt"

	pb "github.com/brotherlogic/discogs/proto"
)

type GetFolderResponse struct {
	Folders []Folder
}

type Folder struct {
	Id   int
	Name string
}

func (p *prodClient) GetUserFolders(ctx context.Context) ([]*pb.Folder, error) {
	gfr := &GetFolderResponse{}
	err := p.makeDiscogsRequest("GET", fmt.Sprintf("/users/%v/collection/folders", p.user.GetUsername()), "", gfr)
	if err != nil {
		return nil, err
	}

	var folders []*pb.Folder
	for _, folder := range gfr.Folders {
		folders = append(folders, &pb.Folder{
			Id:   int32(folder.Id),
			Name: folder.Name,
		})
	}

	return folders, nil

}
