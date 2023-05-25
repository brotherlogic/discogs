package main

import (
	"context"
	"log"
	"os"

	"github.com/brotherlogic/discogs"

	pb "github.com/brotherlogic/discogs/proto"
)

func main() {
	d := discogs.DiscogsWithAuth(os.Args[1], os.Args[2], "https://gramophile.brotherlogic-backend.com/callback")
	ud := d.ForUser(&pb.User{UserToken: os.Args[3], UserSecret: os.Args[4], Username: "brotherlogic"})
	fields, err := ud.GetFields(context.Background())
	if err != nil {
		log.Fatalf("Err: %v", err)
	}
	log.Printf("Fields: %v", fields)
}
