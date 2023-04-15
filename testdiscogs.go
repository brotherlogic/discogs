package discogs

import (
	"context"

	pb "github.com/brotherlogic/discogs/proto"
)

type TestDiscogsClient struct {
	collectionRecords []*pb.Release
}

func (t *TestDiscogsClient) AddCollectionRelease(r *pb.Release) {
	if t.collectionRecords == nil {
		t.collectionRecords = make([]*pb.Release, 0)
	}
	t.collectionRecords = append(t.collectionRecords, r)
}

func (t *TestDiscogsClient) GetLoginURL() (string, string, string, error) {
	return "", "", "", nil
}
func (t *TestDiscogsClient) HandleDiscogsResponse(ctx context.Context, secret, token, verifier string) (string, string, error) {
	return "", "", nil
}

func (t *TestDiscogsClient) GetDiscogsUser(ctx context.Context) (*pb.User, error) {
	return nil, nil
}

func (t *TestDiscogsClient) GetCollection(ctx context.Context, user *pb.User, page int32) ([]*pb.Release, *pb.Pagination, error) {
	return t.collectionRecords, &pb.Pagination{}, nil
}
