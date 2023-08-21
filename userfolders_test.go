package discogs

import (
	"context"
	"testing"
)

func TestGetFolders_Success(t *testing.T) {
	d := GetTestDiscogs()

	folders, err := d.GetUserFolders(context.Background())
	if err != nil {
		t.Fatalf("bad folder return: %v", err)
	}

	if len(folders) == 0 {
		t.Fatalf("no folders returned: %v", folders)
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
