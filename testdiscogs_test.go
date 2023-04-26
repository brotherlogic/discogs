package discogs

import (
	"log"
	"testing"
)

func TestDiscogsTestClient(t *testing.T) {
	var d Discogs = &TestDiscogsClient{}
	log.Printf("TEST %v", d)
}
