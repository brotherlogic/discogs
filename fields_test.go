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
		t.Fatalf("Unable to retrieve collection: %v,%v", fields, err)
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
