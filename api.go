package discogs

import (
	"context"

	pb "github.com/brotherlogic/discogs/proto"
)

type Discogs interface {
	GetCallCount() int32

	GetLoginURL() (string, string, string, error)
	HandleDiscogsResponse(ctx context.Context, secret, token, verifier string) (string, string, error)

	GetDiscogsUser(ctx context.Context) (*pb.User, error)

	GetCollection(ctx context.Context, page int32) ([]*pb.Release, *pb.Pagination, error)
	GetCollectionRelease(ctx context.Context, iid int64, page int32) ([]*pb.Release, *pb.Pagination, error)

	GetUserId() int32

	SetDownloader(downloader Downloader)

	ForUser(user *pb.User) Discogs

	GetFields(ctx context.Context) ([]*pb.Field, error)
	SetField(ctx context.Context, r *pb.Release, fnum int, value string) error

	GetUserFolders(ctx context.Context) ([]*pb.Folder, error)
	CreateFolder(ctx context.Context, folderName string) (*pb.Folder, error)
	DeleteFolder(ctx context.Context, folderId int32) error

	CreateSale(ctx context.Context, params *pb.SaleParams) (int64, error)
	GetSale(ctx context.Context, saleId int64) (*pb.SaleItem, error)
	UpdateSale(ctx context.Context, saleId int64, releaseId int64, condition string, newPrice int32) error
	UpdateSaleState(ctx context.Context, saleId, releaseId int64, condition string, saleSate pb.SaleStatus) error
	ListSales(ctx context.Context, page int32) ([]*pb.SaleItem, *pb.Pagination, error)
	GetOrder(ctx context.Context, orderId string) (*pb.Order, error)

	GetReleaseStats(ctx context.Context, releaseId int64) (*pb.ReleaseStats, error)

	SetFolder(ctx context.Context, instanceId, releaseId int64, folderId, newFolderId int32) error
	AddRelease(ctx context.Context, id, folder int64) (int64, error)

	GetRelease(ctx context.Context, releaseId int64) (*pb.Release, error)
	SetRating(ctx context.Context, releaseId int64, rating int32) error

	GetWants(ctx context.Context, page int32) ([]*pb.Want, *pb.Pagination, error)
	AddWant(ctx context.Context, releaseId int64) (*pb.Want, error)
	DeleteWant(ctx context.Context, wantId int64) error

	GetMasterReleases(ctx context.Context, masterId int64, page int32, sort pb.MasterSort) ([]*pb.MasterRelease, error)

	Throttle()
}
