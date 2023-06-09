package discogs

import (
	"context"

	pb "github.com/brotherlogic/discogs/proto"
)

type TestDiscogsClient struct {
	collectionRecords []*pb.Release
}

func (t *TestDiscogsClient) GetUserId() int32 {
	return int32(10)
}

func (t *TestDiscogsClient) ForUser(user *pb.User) Discogs {
	return t
}

func (t *TestDiscogsClient) GetFields(_ context.Context) ([]*pb.Field, error) {
	return []*pb.Field{}, nil
}

func (t *TestDiscogsClient) SetField(ctx context.Context, r *pb.Release, fnum int, value string) error {
	return nil
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

func (t *TestDiscogsClient) GetCollection(ctx context.Context, page int32) ([]*pb.Release, *pb.Pagination, error) {
	return t.collectionRecords, &pb.Pagination{}, nil
}
