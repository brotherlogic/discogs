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

	if len(wants) != 32 || pagination.Pages != 1 {
		t.Errorf("Bad sale return %v -> %v", wants, pagination)
	}
}

func TestAddWant_Success(t *testing.T) {
	td := GetTestDiscogs()

	want, err := td.AddWant(context.Background(), 12778444)

	if err != nil {
		t.Fatalf("Bad list sales: %v", err)
	}

	if want.GetTitle() != "In A Mood" {
		t.Errorf("Bad want return: %v", want)
	}
}
