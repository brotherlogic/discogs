package discogs

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestSetRating(t *testing.T) {
	d := GetTestDiscogs()

	err := d.SetRating(context.Background(), 3139057, 5)
	if err != nil {
		t.Errorf("Failed to set rating: %v", err)
	}
}

func TestGetRelease(t *testing.T) {
	d := GetTestDiscogs()

	r, err := d.GetRelease(context.Background(), 372000)

	if err != nil {
		t.Fatalf("Failed to get release: %v", err)
	}

	tt := time.Now()
	log.Printf("HERE: %v", tt.Format("_2 Jan 2006"))
	storedDate := "16 Mar 1979"
	dVal, err := time.Parse("_2 Jan 2006", storedDate)
	if err != nil {
		t.Fatalf("Date parsing problem: %v", err)
	}

	if r.GetReleaseDate() != dVal.Unix() {
		t.Errorf("Bad release returned: %v -> should have been %v", time.Unix(r.GetReleaseDate(), 0), dVal)
	}
}
