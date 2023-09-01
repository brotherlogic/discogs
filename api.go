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
	AllowOffers     bool
	Status          string
	ExternalId      string
	Location        string
	Weight          int32
	FormatQuantity  int32
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
	CreateFolder(ctx context.Context, folderName string) (*pb.Folder, error)
	DeleteFolder(ctx context.Context, folderId int32) error

	CreateSale(ctx context.Context, params SaleParams) (int64, error)
	GetSale(ctx context.Context, saleId int64) (*pb.SaleItem, error)
	ListSales(ctx context.Context, page int32) ([]*pb.SaleItem, *pb.Pagination, error)
	GetOrder(ctx context.Context, orderId string) (*pb.Order, error)

	SetFolder(ctx context.Context, instanceId, releaseId, folderId, newFolderId int64) error

	GetWants(ctx context.Context, page int32) ([]*pb.Want, *pb.Pagination, error)
	AddWant(ctx context.Context, releaseId int64) (*pb.Want, error)
}
