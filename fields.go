package discogs

import (
	"context"
	"fmt"
	"log"

	"encoding/json"

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
	log.Printf("USER: %v with %+v", d.user, d)
	cr := &FieldsResponse{}
	err := d.makeDiscogsRequest("GET", fmt.Sprintf("/users/%v/collection/fields", d.user.GetUsername()), "", cr)
	if err != nil {
		return nil, err
	}
	log.Printf("Resp %v,%v", err, cr)

	var fields []*pb.Field
	for _, field := range cr.Fields {
		fields = append(fields, &pb.Field{
			Id:   int32(field.Id),
			Name: field.Name,
		})
	}

	return fields, nil
}

type fieldUpdate struct {
	Value string `json:"value"`
}

func (d *prodClient) SetField(ctx context.Context, r *pb.Release, fnum int, value string) error {
	cr := &FieldsResponse{}
	vjson := &fieldUpdate{Value: value}
	vstr, _ := json.Marshal(vjson)
	return d.makeDiscogsRequest("POST", fmt.Sprintf("/users/%v/collection/folders/0/releases/%v/instances/%v/fields/%v", d.user.GetUsername(), r.GetId(), r.GetInstanceId(), fnum), string(vstr), cr)
}
