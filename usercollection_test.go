package discogs

import (
	"context"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	pb "github.com/brotherlogic/discogs/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetTestDiscogs() *Discogs {
	return &Discogs{
		getter: &testGetter{},
	}
}

type testGetter struct{}

func (tg *testGetter) get() myClient {
	return &tClient{}
}

type tClient struct{}

func (t *tClient) Get(url string) (*http.Response, error) {
	response := &http.Response{}
	testFile := strings.Replace(strings.Replace(url[23:], "?", "_", -1), "&", "_", -1)

	blah, err := os.Open("testdata" + testFile)

	if err != nil {
		return nil, err
	}
	response.Body = blah

	return response, nil
}

func (t *tClient) Post(url, contentType string, body io.Reader) (*http.Response, error) {
	return t.Get(url)
}

func TestGetCollection(t *testing.T) {
	td := GetTestDiscogs()

	coll, pag, err := td.GetCollection(context.Background(), &pb.User{Username: "brotherlogic"}, 1)

	if err != nil {
		t.Fatalf("Unable to retrieve collection: %v -> %v,%v", err, coll, pag)
	}

	if len(coll) != 100 {
		t.Errorf("Bad collection size, expected 50, got %v", len(coll))
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
		}
	}

	if !found {
		t.Errorf("Release 570258 was not found")
	}
}

func TestGetCollectionPageOutOfBounds(t *testing.T) {
	td := GetTestDiscogs()

	coll, pag, err := td.GetCollection(context.Background(), &pb.User{Username: "brotherlogic"}, 100)

	if status.Code(err) != codes.OutOfRange {
		t.Fatalf("Did not return out of of range: %v -> %v,%v", err, coll, pag)
	}
}
