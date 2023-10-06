package discogs

import (
	"context"
	"log"

	pb "github.com/brotherlogic/discogs/proto"
)

type TestDiscogsClient struct {
	collectionRecords []*pb.Release
	UserId            int32
	Fields            []*pb.Field
	Folders           []*pb.Folder
	Sales             []*pb.SaleItem
	Rating            map[int64]int32
}

func GetTestClient() *TestDiscogsClient {
	return &TestDiscogsClient{Rating: make(map[int64]int32)}
}

func (t *TestDiscogsClient) Throttle() {

}

func (t *TestDiscogsClient) SetRating(ctx context.Context, releaseId int64, newScore int32) error {
	t.Rating[releaseId] = newScore
	return nil
}

func (t *TestDiscogsClient) GetUserId() int32 {
	return t.UserId
}

func (t *TestDiscogsClient) CreateSale(ctx context.Context, params SaleParams) (int64, error) {
	return 1234, nil
}

func (t *TestDiscogsClient) CreateFolder(ctx context.Context, folderName string) (*pb.Folder, error) {
	folder := &pb.Folder{Name: folderName, Id: 123}
	t.Folders = append(t.Folders, folder)
	log.Printf("Added folder: %v", t.Folders)
	return folder, nil
}

func (t *TestDiscogsClient) DeleteFolder(ctx context.Context, folderId int32) error {
	return nil
}

func (t *TestDiscogsClient) GetOrder(ctx context.Context, orderId string) (*pb.Order, error) {
	return &pb.Order{}, nil
}

func (t *TestDiscogsClient) AddWant(ctx context.Context, releaseId int64) (*pb.Want, error) {
	return &pb.Want{}, nil
}

func (t *TestDiscogsClient) DeleteWant(ctx context.Context, wantId int64) error {
	return nil
}

func (t *TestDiscogsClient) GetSale(ctx context.Context, saleId int64) (*pb.SaleItem, error) {
	for _, sale := range t.Sales {
		if sale.GetSaleId() == saleId {
			return sale, nil
		}
	}
	return &pb.SaleItem{}, nil
}

func (t *TestDiscogsClient) UpdateSale(ctx context.Context, saleId int64, releaseId int64, condition string, newPrice int32) error {
	for _, sale := range t.Sales {
		if sale.GetSaleId() == saleId {
			sale.Price.Value = newPrice
		}
	}
	return nil
}

func (t *TestDiscogsClient) GetWants(ctx context.Context, page int32) ([]*pb.Want, *pb.Pagination, error) {
	return []*pb.Want{}, &pb.Pagination{}, nil
}

func (t *TestDiscogsClient) ListSales(ctx context.Context, page int32) ([]*pb.SaleItem, *pb.Pagination, error) {
	return t.Sales, &pb.Pagination{}, nil
}

func (t *TestDiscogsClient) SetFolder(ctx context.Context, instanceId, releaseId, folderId, newFolderId int64) error {
	return nil
}

func (t *TestDiscogsClient) GetReleaseStats(ctx context.Context, releaseId int32) (*pb.ReleaseStats, error) {
	return &pb.ReleaseStats{}, nil
}

func (t *TestDiscogsClient) ForUser(user *pb.User) Discogs {
	return t
}

func (t *TestDiscogsClient) GetUserFolders(_ context.Context) ([]*pb.Folder, error) {
	log.Printf("Getting folders: %v", t.Folders)
	return t.Folders, nil
}

func (t *TestDiscogsClient) GetFields(_ context.Context) ([]*pb.Field, error) {
	return t.Fields, nil
}

func (t *TestDiscogsClient) SetField(ctx context.Context, r *pb.Release, fnum int, value string) error {
	t.Fields = append(t.Fields, &pb.Field{Name: value, Id: int32(fnum)})
	return nil
}

func (t *TestDiscogsClient) AddCollectionRelease(r *pb.Release) {
	if t.collectionRecords == nil {
		t.collectionRecords = make([]*pb.Release, 0)
	}
	t.collectionRecords = append(t.collectionRecords, r)
}

func (t *TestDiscogsClient) GetLoginURL() (string, string, string, error) {
	return "", "madeuptoken", "madeupysecret", nil
}
func (t *TestDiscogsClient) HandleDiscogsResponse(ctx context.Context, secret, token, verifier string) (string, string, error) {
	return "", "", nil
}

func (t *TestDiscogsClient) GetDiscogsUser(ctx context.Context) (*pb.User, error) {
	return &pb.User{Username: "madeupuser"}, nil
}

func (t *TestDiscogsClient) GetCollection(ctx context.Context, page int32) ([]*pb.Release, *pb.Pagination, error) {
	return t.collectionRecords, &pb.Pagination{}, nil
}
