package discogs

import (
	"context"
	"testing"
)

func TestSetRating(t *testing.T) {
	d := GetTestDiscogs()

	err := d.SetRating(context.Background(), 3139057, 5)
	if err != nil {
		t.Errorf("Failed to set rating: %v", err)
	}
}
