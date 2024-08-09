package discogs

import (
	"context"
	"log"
	"testing"
	"time"

	pb "github.com/brotherlogic/discogs/proto"
)

func TestSetRating(t *testing.T) {
	d := GetTestDiscogs()

	err := d.SetRating(context.Background(), 3139057, 5)
	if err != nil {
		t.Errorf("Failed to set rating: %v", err)
	}
}

func TestGetMasterReleases(t *testing.T) {
	d := GetTestDiscogs()

	rs, err := d.GetMasterReleases(context.Background(), 1693557, 1, pb.MasterSort_BY_YEAR)
	if err != nil {
		t.Fatalf("Unable to get master releases: %v", err)
	}

	if len(rs) != 3 {
		t.Errorf("Bad number of releases: %v", rs)
	}

	found := false
	for _, release := range rs {
		if release.GetId() == 15777447 && release.GetYear() == 2020 {
			found = true
		}
	}

	if !found {
		t.Errorf("Did not find key release: %v", rs)
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

	if r.GetMasterId() != 38998 {
		t.Errorf("Failed to get master: %v", r.GetMasterId())
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

func TestGetRelease_EdgeCase2(t *testing.T) {
	d := GetTestDiscogs()

	r, err := d.GetRelease(context.Background(), 1059056)
	if err != nil {
		t.Fatalf("Failed to get release: %v", err)
	}

	storedDate := "2002"
	dVal, err := time.Parse("2006", storedDate)
	if err != nil {
		t.Fatalf("Date parsing problem: %v", err)
	}

	if r.GetReleaseDate() != dVal.Unix() {
		t.Errorf("Bad release returned: %v -> should have been %v", time.Unix(r.GetReleaseDate(), 0), dVal)
	}
}
