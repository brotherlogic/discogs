package discogs

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	pb "github.com/brotherlogic/discogs/proto"
	"github.com/dghubble/oauth1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetTestDiscogs() Discogs {
	return &prodClient{
		getter: &testGetter{},
		user:   &pb.User{Username: "brotherlogic"},
	}
}

type testGetter struct{}

func (tg *testGetter) get() myClient {
	return &tClient{}
}

func (tg *testGetter) config() oauth1.Config {
	return oauth1.Config{}
}

type tClient struct{}

func (t *tClient) Get(url string) (*http.Response, error) {
	response := &http.Response{}
	testFile := strings.Replace(strings.Replace(url[23:], "?", "_", -1), "&", "_", -1)

	stat, err := os.Stat("testdata" + testFile)
	if err != nil {
		return nil, err
	}

	adder := ""
	if stat.IsDir() {
		adder = "/FILE"
	}

	blah, err := os.Open("testdata" + testFile + adder)

	if err != nil {
		return nil, err
	}
	response.Body = blah

	return response, nil
}

func hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func (t *tClient) Post(url, contentType string, body io.Reader) (*http.Response, error) {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, body)
	if err != nil {
		return nil, err
	}
	return t.Get(url + "_" + hash(buf.String()))
}

func TestGetCollection(t *testing.T) {
	td := GetTestDiscogs()

	coll, pag, err := td.GetCollection(context.Background(), 1)

	if err != nil {
		t.Fatalf("Unable to retrieve collection: %v -> %v,%v", err, coll, pag)
	}

	if len(coll) != 100 {
		t.Errorf("Bad collection size, expected 50, got %v", len(coll))
	}

	if pag.GetPages() != 46 {
		t.Errorf("Bad pagination return: %v", pag)
	}

	found := false
	for _, release := range coll {
		if release.GetId() == 570258 {
			found = true
			if release.GetInstanceId() != 365214833 {
				t.Errorf("Bad instance id: %v", release)
			}
			if release.GetFolderId() != 1727264 {
				t.Errorf("Bad folder id: %v", release)
			}
			if release.GetRating() != 0 {
				t.Errorf("Bad rating: %v", release)
			}

			foundVinyl := false
			for _, format := range release.GetFormats() {
				if format.GetName() == "Vinyl" && format.Quantity == 1 {
					foundVinyl = true
				}
			}
			if !foundVinyl {
				t.Errorf("Unable to find vinyl: %v", release)
			}

			foundAnkst := false
			for _, label := range release.GetLabels() {
				if label.GetName() == "Ankst" && label.Catno == "ANKST 059" && label.Id == 33378 {
					foundAnkst = true
				}
			}
			if !foundAnkst {
				t.Errorf("The label was not found: %v", release)
			}

		}
	}

	if !found {
		t.Errorf("Release 570258 was not found")
	}
}

func TestGetCollectionPageOutOfBounds(t *testing.T) {
	td := GetTestDiscogs()

	coll, pag, err := td.GetCollection(context.Background(), 100)

	if status.Code(err) != codes.OutOfRange {
		t.Fatalf("Did not return out of of range: %v -> %v,%v", err, coll, pag)
	}
}
