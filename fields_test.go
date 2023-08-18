package discogs

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/discogs/proto"
)

func TestGetFields(t *testing.T) {
	td := GetTestDiscogs()

	fields, err := td.GetFields(context.Background())

	if err != nil {
		t.Fatalf("unable to retrieve collection: %v,%v", fields, err)
	}

	var f *pb.Field
	for _, field := range fields {
		if field.GetName() == "Cleaned" {
			f = field
		}
	}

	if f == nil || f.Id != 5 {
		t.Errorf("Bad field pull: %v -> %v", f, fields)
	}
}

func TestSetFields(t *testing.T) {
	td := GetTestDiscogs()

	err := td.SetField(context.Background(), &pb.Release{Id: 1163112, InstanceId: 19867414}, 5, "Testing")

	if err != nil {
		t.Fatalf("Unable to set field: %v", err)
	}
}
