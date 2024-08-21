package discogs

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
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

func GetTestDiscogsWithPersonalToken() Discogs {
	return &prodClient{
		getter:        &testGetter{},
		user:          &pb.User{Username: "brotherlogic"},
		personalToken: "personal_token",
	}
}

type testGetter struct{}

func (tg *testGetter) getDefault() myClient {
	return &tClient{}
}

func (tg *testGetter) get() myClient {
	return &tClient{}
}

func (tg *testGetter) config() oauth1.Config {
	return oauth1.Config{}
}

type tClient struct{}

func (t *tClient) Do(req *http.Request) (*http.Response, error) {
	log.Printf("HERE: %+v", req)
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	if req.Method == "DELETE" {
		return t.Get(fmt.Sprintf("%v", req.URL))
	}
	log.Printf("HASH: %v", string(body))
	return t.Get(fmt.Sprintf("%v_%v", req.URL, hash(string(body))))
}

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
	response.StatusCode = 200

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
	log.Printf("POSTING: %v", buf.String())
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

	if pag.GetPages() != 48 {
		t.Errorf("Bad pagination return: %v", pag)
	}

	found := false
	for _, release := range coll {
		if release.GetId() == 5964926 {
			found = true
			if release.GetInstanceId() != 780876394 {
				t.Errorf("Bad instance id: %v", release)
			}
			if release.GetFolderId() != 267116 {
				t.Errorf("Bad folder id: %v", release)
			}
			if release.GetRating() != 5 {
				t.Errorf("Bad rating: %v", release)
			}

			if release.GetTitle() != "Loveseed" {
				t.Errorf("Title has not been returned")
			}

			if release.GetNotes()[4] != "1.6" {
				t.Errorf("Notes have not been returned: %v", release.GetNotes())
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

			foundLabel := false
			for _, label := range release.GetLabels() {
				if label.GetName() == "Amestay" && label.Catno == "1001" && label.Id == 724668 {
					foundLabel = true
				}
			}
			if !foundLabel {
				t.Errorf("The label was not found: %v", release)
			}

			foundArtist := false
			for _, artist := range release.GetArtists() {
				if artist.GetName() == "The Loveseed" && artist.GetId() == 3954346 {
					foundArtist = true
				}
			}

			if !foundArtist {
				t.Errorf("Artist was not found: %v", release)
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

func TestSetFolder(t *testing.T) {
	td := GetTestDiscogs()

	err := td.SetFolder(context.Background(), 1427071368, 27915987, 3578980, 242017)
	if err != nil {
		t.Errorf("Error setting folder: %v", err)
	}
}

func TestCreateFolder(t *testing.T) {
	td := GetTestDiscogs()
	folder, err := td.CreateFolder(context.Background(), "TestFolder")
	if err != nil {
		t.Fatalf("Error creating folder: %v", err)
	}

	if folder.GetName() != "TestFolder" || folder.Id != 6259627 {
		t.Errorf("Bad folder create: %v", folder)
	}
}

func TestDeleteFolder(t *testing.T) {
	td := GetTestDiscogs()
	err := td.DeleteFolder(context.Background(), 6259627)
	if err != nil {
		t.Fatalf("Error deleting folder: %v", err)
	}
}

func TestGetCollectionRelease(t *testing.T) {
	td := GetTestDiscogs()
	r, _, err := td.GetCollectionRelease(context.Background(), 550078, 1)
	if err != nil {
		t.Fatalf("Unable to get release: %v", err)
	}

	if len(r) != 1 {
		t.Fatalf("Bad pull: %v", len(r))
	}

	if r[0].GetTitle() != "Tra La La / Stand Up And Be A Man" {
		t.Errorf("Bad title pull: %v", r[0])
	}

	if r[0].GetFolderId() != 1613206 {
		t.Errorf("Bad folder id: %v", r[0])
	}

}
