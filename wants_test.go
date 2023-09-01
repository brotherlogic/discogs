package discogs

import (
	"context"
	"testing"
)

func TestGetWants_Success(t *testing.T) {
	td := GetTestDiscogs()

	wants, pagination, err := td.GetWants(context.Background(), 1)

	if err != nil {
		t.Fatalf("Bad list sales: %v", err)
	}

	if len(wants) != 32 {
		t.Errorf("Bad sale return %v -> %v", wants, pagination)
	}
}
