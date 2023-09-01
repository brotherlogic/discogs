package discogs

import (
	"context"
	"encoding/json"
	"fmt"

	pb "github.com/brotherlogic/discogs/proto"
)

type GetFolderResponse struct {
	Folders []Folder
}

type SetFolderResponse struct{}

type Folder struct {
	Id   int
	Name string
}

type NewFolder struct {
	FolderId int `json:"folder_id"`
}

func (p *prodClient) SetFolder(ctx context.Context, instanceId, releaseId, folderId, newFolderId int64) error {
	data := &NewFolder{
		FolderId: int(newFolderId),
	}
	v, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = p.makeDiscogsRequest("POST", fmt.Sprintf("/users/%v/collection/folders/%v/releases/%v/instances/%v",
		p.user.GetUsername(), folderId, releaseId, instanceId), string(v), &SetFolderResponse{})
	if err != nil {
		return err
	}
	return nil
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
