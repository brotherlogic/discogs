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

	artistFound := false
	for _, artist := range r.GetArtists() {
		if artist.GetName() == "The Fall" && artist.GetId() == 2228 {
			artistFound = true
		}
	}
	if !artistFound {
		t.Errorf("Did not find artist: %v", r)
	}

	labelFound := false
	for _, label := range r.GetLabels() {
		if label.GetName() == "Step-Forward Records" && label.GetId() == 14962 {
			labelFound = true
		}
	}
	if !labelFound {
		t.Errorf("Did not find label: %v", r)
	}
}

func TestGetRelease_JustYear(t *testing.T) {
	d := GetTestDiscogs()

	r, err := d.GetRelease(context.Background(), 1929402)

	if err != nil {
		t.Fatalf("Failed to get release: %v", err)
	}

	storedDate := "1994"
	dVal, err := time.Parse("2006", storedDate)
	if err != nil {
		t.Fatalf("Date parsing problem: %v", err)
	}

	if r.GetReleaseDate() != dVal.Unix() {
		t.Errorf("Bad release returned: %v -> should have been %v", time.Unix(r.GetReleaseDate(), 0), dVal)
	}
}

func TestGetRelease_JustYearAndMonth(t *testing.T) {
	d := GetTestDiscogs()

	r, err := d.GetRelease(context.Background(), 372019)

	if err != nil {
		t.Fatalf("Failed to get release: %v", err)
	}

	storedDate := "May 1980"
	dVal, err := time.Parse("Jan 2006", storedDate)
	if err != nil {
		t.Fatalf("Date parsing problem: %v", err)
	}

	if r.GetReleaseDate() != dVal.Unix() {
		t.Errorf("Bad release returned: %v -> should have been %v", time.Unix(r.GetReleaseDate(), 0), dVal)
	}
}

func TestGetRelease_EdgeCase(t *testing.T) {
	d := GetTestDiscogs()

	r, err := d.GetRelease(context.Background(), 939775)

	if err != nil {
		t.Fatalf("Failed to get release: %v", err)
	}

	if r.GetReleaseDate() != 0 {
		t.Errorf("Bad release returned: %v -> should have been %v", time.Unix(r.GetReleaseDate(), 0), 0)
	}
}
