package discogs

import (
	"context"
	"testing"
)

func TestGetFolders_Success(t *testing.T) {
	d := GetTestDiscogs()

	folders, err := d.GetUserFolders(context.Background())
	if err != nil {
		t.Fatalf("Bad folder return: %v", err)
	}

	if len(folders) == 0 {
		t.Fatalf("No folders returned")
	}

	found := false
	for _, folder := range folders {
		if folder.GetName() == "12s" {
			found = true
		}
	}

	if !found {
		t.Errorf("Could not find folder")
	}
}
