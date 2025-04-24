package discogs

import (
	"context"
	"log"

	pb "github.com/brotherlogic/discogs/proto"
	"golang.org/x/exp/maps"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TestDiscogsClient struct {
	collectionRecords    []*pb.Release
	nonCollectionRecords []*pb.Release
	UserId               int32
	Fields               []*pb.Field
	Folders              []*pb.Folder
	Sales                []*pb.SaleItem
	Rating               map[int64]int32
	Wants                map[int64]*pb.Want
	Masters              map[int64]*pb.Release
	callCount            int32
}

func GetTestClient() *TestDiscogsClient {
	return &TestDiscogsClient{
		Rating: make(map[int64]int32),
		Wants:  make(map[int64]*pb.Want)}
}

func (t *TestDiscogsClient) Throttle() {

}

func (t *TestDiscogsClient) GetCallCount() int32 {
	return t.callCount
}

func (t *TestDiscogsClient) UpdateSaleState(ctx context.Context, saleId int64, releaseId int64, condition string, saleState pb.SaleStatus) error {
	t.callCount++
	for _, sale := range t.Sales {
		if sale.GetSaleId() == saleId {
			sale.Status = saleState
		}
	}
	return nil
}

func (t *TestDiscogsClient) GetMasterReleases(ctx context.Context, masterId int64, page int32, sort pb.MasterSort) ([]*pb.MasterRelease, error) {
	t.callCount++
	var mr []*pb.MasterRelease
	for _, r := range t.collectionRecords {
		if r.GetMasterId() == masterId {
			mr = append(mr, &pb.MasterRelease{Id: r.GetId()})
		}
	}
	for _, r := range t.nonCollectionRecords {
		if r.GetMasterId() == masterId {
			mr = append(mr, &pb.MasterRelease{Id: r.GetId()})
		}
	}
	return mr, nil
}

func (t *TestDiscogsClient) SetRating(ctx context.Context, releaseId int64, newScore int32) error {
	t.callCount++
	if t.Rating == nil {
		t.Rating = make(map[int64]int32)
	}
	t.Rating[releaseId] = newScore
	return nil
}

func (t *TestDiscogsClient) GetUserId() int32 {
	return t.UserId
}

func (t *TestDiscogsClient) CreateSale(ctx context.Context, params SaleParams) (int64, error) {
	t.callCount++
	return 1234, nil
}

func (t *TestDiscogsClient) CreateFolder(ctx context.Context, folderName string) (*pb.Folder, error) {
	t.callCount++
	folder := &pb.Folder{Name: folderName, Id: 123}
	t.Folders = append(t.Folders, folder)
	log.Printf("Added folder: %v", t.Folders)
	return folder, nil
}

func (t *TestDiscogsClient) DeleteFolder(ctx context.Context, folderId int32) error {
	t.callCount++
	return nil
}

func (t *TestDiscogsClient) GetOrder(ctx context.Context, orderId string) (*pb.Order, error) {
	t.callCount++
	return &pb.Order{}, nil
}

func (t *TestDiscogsClient) AddWant(ctx context.Context, releaseId int64) (*pb.Want, error) {
	t.callCount++
	if t.Wants == nil {
		t.Wants = make(map[int64]*pb.Want)
	}
	t.Wants[releaseId] = &pb.Want{Id: releaseId}
	return &pb.Want{}, nil
}

func (t *TestDiscogsClient) DeleteWant(ctx context.Context, wantId int64) error {
	t.callCount++
	if t.Wants == nil {
		t.Wants = make(map[int64]*pb.Want)
	}
	delete(t.Wants, wantId)
	return nil
}

func (t *TestDiscogsClient) GetSale(ctx context.Context, saleId int64) (*pb.SaleItem, error) {
	t.callCount++
	for _, sale := range t.Sales {
		if sale.GetSaleId() == saleId {
			return sale, nil
		}
	}
	return &pb.SaleItem{}, nil
}

func (t *TestDiscogsClient) SetDownloader(_ Downloader) {
	// Do nothing
}

func (t *TestDiscogsClient) UpdateSale(ctx context.Context, saleId int64, releaseId int64, condition string, newPrice int32) error {
	t.callCount++
	for _, sale := range t.Sales {
		if sale.GetSaleId() == saleId {
			sale.Price.Value = newPrice
		}
	}
	return nil
}

func (t *TestDiscogsClient) GetWants(ctx context.Context, page int32) ([]*pb.Want, *pb.Pagination, error) {
	t.callCount++
	if t.Wants == nil {
		t.Wants = make(map[int64]*pb.Want)
	}
	return maps.Values(t.Wants), &pb.Pagination{}, nil
}

func (t *TestDiscogsClient) ListSales(ctx context.Context, page int32) ([]*pb.SaleItem, *pb.Pagination, error) {
	t.callCount++
	return t.Sales, &pb.Pagination{}, nil
}

func (t *TestDiscogsClient) GetRelease(ctx context.Context, releaseId int64) (*pb.Release, error) {
	t.callCount++
	for _, r := range t.collectionRecords {
		if r.GetId() == releaseId {
			return r, nil
		}
	}

	for _, r := range t.nonCollectionRecords {
		if r.GetId() == releaseId {
			return r, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "Unable to locate %v  (given %v)", releaseId, t.nonCollectionRecords)
}

func (t *TestDiscogsClient) SetFolder(ctx context.Context, instanceId, releaseId int64, folderId, newFolderId int32) error {
	t.callCount++
	return nil
}

func (t *TestDiscogsClient) GetReleaseStats(ctx context.Context, releaseId int64) (*pb.ReleaseStats, error) {
	t.callCount++
	return &pb.ReleaseStats{}, nil
}

func (t *TestDiscogsClient) ForUser(user *pb.User) Discogs {
	t.UserId = user.GetDiscogsUserId()
	return t
}

func (t *TestDiscogsClient) GetUserFolders(_ context.Context) ([]*pb.Folder, error) {
	t.callCount++
	log.Printf("Getting folders: %v", t.Folders)
	return t.Folders, nil
}

func (t *TestDiscogsClient) GetFields(_ context.Context) ([]*pb.Field, error) {
	t.callCount++
	return t.Fields, nil
}

func (t *TestDiscogsClient) SetField(ctx context.Context, r *pb.Release, fnum int, value string) error {
	t.callCount++
	t.Fields = append(t.Fields, &pb.Field{Name: value, Id: int32(fnum)})
	return nil
}

func (t *TestDiscogsClient) AddCollectionRelease(r *pb.Release) {
	t.callCount++
	if t.collectionRecords == nil {
		t.collectionRecords = make([]*pb.Release, 0)
	}
	t.collectionRecords = append(t.collectionRecords, r)
}

func (t *TestDiscogsClient) AddRelease(ctx context.Context, id, folder int64) (int64, error) {
	t.callCount++
	if t.collectionRecords == nil {
		t.collectionRecords = make([]*pb.Release, 0)
	}
	t.collectionRecords = append(t.collectionRecords, &pb.Release{Id: id, FolderId: int32(folder)})
	return id, nil
}

func (t *TestDiscogsClient) GetCollectionRelease(ctx context.Context, id int64, page int32) ([]*pb.Release, *pb.Pagination, error) {
	t.callCount++
	var ret []*pb.Release
	for _, r := range t.collectionRecords {
		if r.Id == id {
			ret = append(ret, r)
		}
	}
	return ret, &pb.Pagination{}, nil
}

func (t *TestDiscogsClient) AddNonCollectionRelease(r *pb.Release) {
	if t.nonCollectionRecords == nil {
		t.nonCollectionRecords = make([]*pb.Release, 0)
	}
	t.nonCollectionRecords = append(t.nonCollectionRecords, r)
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
	t.callCount++
	return append(t.collectionRecords), &pb.Pagination{}, nil
}
