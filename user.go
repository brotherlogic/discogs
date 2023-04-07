package discogs

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/brotherlogic/discogs/proto"
)

type DiscogsUser struct {
	AvatarURL    string `json:"avatar_url"`
	ID           int32  `json:"id"`
	Username     string `json:"username"`
	CurrencyAbbr string `json:"curr_abbr"`
}

func (d *Discogs) GetDiscogsUser(ctx context.Context) (*pb.User, err) {
	user := &DiscogsUser{}
	err := d.makeDiscogsRequest("GET", "oauth/identity", "", user)
	return &pb.User{
		Username:      user.Username,
		DiscogsUserId: user.ID,
	}, err
}
func (d *Discogs) makeDiscogsRequest(rtype, path string, data string, obj interface{}) error {
	fullPath := fmt.Sprintf("https://api.discogs.com/%v", path)
	httpClient := d.getter.get()

	if rtype == "POST" {
		resp, err := httpClient.Post(fullPath, "application/json", bytes.NewBuffer([]byte(data)))
		if err != nil {
			return err
		}

		// Throttling
		if resp.StatusCode == 429 {
			return status.Errorf(codes.ResourceExhausted, "Discogs is throttling us")
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if len(body) > 0 {
			err = json.Unmarshal(body, obj)
			if err != nil {
				return fmt.Errorf("Unarshal error (processing %v): %v from %v", err, string(body), data)
			}
		}
		return nil

	}
	resp, err := httpClient.Get(fullPath)
	if err != nil {
		return err
	}

	if resp.StatusCode == 404 {
		return status.Errorf(codes.NotFound, "Unable to locate sale")
	}

	// Throttling
	if resp.StatusCode == 429 {
		return status.Errorf(codes.ResourceExhausted, "Discogs is throttling us")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Printf("GOT BODY: %v", body)
	if len(body) > 0 {
		err = json.Unmarshal(body, obj)
		if err != nil {
			return fmt.Errorf("Unmarshal error (processing %v): %v", string(body), err)
		}
	}

	return nil
}
