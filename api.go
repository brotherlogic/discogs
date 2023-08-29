package discogs

import (
	"context"

	pb "github.com/brotherlogic/discogs/proto"
)

type SaleParams struct {
	ReleaseId       int32
	Condition       string
	SleeveCondition string
	Price           int32
	Comments        string
}

type Discogs interface {
	GetLoginURL() (string, string, string, error)
	HandleDiscogsResponse(ctx context.Context, secret, token, verifier string) (string, string, error)

	GetDiscogsUser(ctx context.Context) (*pb.User, error)

	GetCollection(ctx context.Context, page int32) ([]*pb.Release, *pb.Pagination, error)

	GetUserId() int32

	ForUser(user *pb.User) Discogs

	GetFields(ctx context.Context) ([]*pb.Field, error)
	SetField(ctx context.Context, r *pb.Release, fnum int, value string) error

	GetUserFolders(ctx context.Context) ([]*pb.Folder, error)

	CreateSale(ctx context.Context, params SaleParams) (int64, error)
}
