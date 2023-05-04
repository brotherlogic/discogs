package discogs

import (
	"context"
	"fmt"

	pb "github.com/brotherlogic/discogs/proto"
)

type FieldsResponse struct {
	Fields []Field
}

type Field struct {
	Id   int
	Name string
}

func (d *prodClient) GetFields(ctx context.Context) ([]*pb.Field, error) {
	cr := &FieldsResponse{}
	err := d.makeDiscogsRequest("GET", fmt.Sprintf("/users/%v/collection/fields", d.user.GetUsername()), "", cr)
	if err != nil {
		return nil, err
	}

	var fields []*pb.Field
	for _, field := range cr.Fields {
		fields = append(fields, &pb.Field{
			Id:   int32(field.Id),
			Name: field.Name,
		})
	}

	return fields, nil
}
