package discogs

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/discogs/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreateSale_Success(t *testing.T) {
	td := GetTestDiscogs()

	saleid, err := td.CreateSale(context.Background(), &pb.SaleParams{
		ReleaseId: 27962688,
		Condition: "Mint (M)",
		Price:     10023,
		Status:    "For Sale",
	})

	if err != nil {
		t.Fatalf("Error creating sale: %v", err)
	}

	if saleid != 2851616155 {
		t.Errorf("Bad saleid return: %v", saleid)
	}
}

func TestGetSaleStats(t *testing.T) {
	td := GetTestDiscogs()

	stats, err := td.GetSaleStats(context.Background(), 189766)
	if err != nil {
		t.Fatalf("Bad get sale stats: %v", err)
	}

	if stats.GetVgPrice() < 16.98 || stats.GetVgPrice() > 16.99 {
		t.Errorf("Bad vg price: %v -> 16.99", stats.GetVgPrice())
	}

	if stats.GetMPrice() < 35.86 || stats.GetMPrice() > 35.87 {
		t.Errorf("Bad m price: %v -> 16.99", stats.GetMPrice())
	}

}

func TestExpireSale(t *testing.T) {
	td := GetTestDiscogs()

	err := td.UpdateSaleState(context.Background(), 2828937565, 1349214, "Very Good Plus (VG+)", pb.SaleStatus_EXPIRED)
	if err != nil {
		t.Errorf("Unable to expire sale: %v", err)
	}
}

func TestGetSale_Success(t *testing.T) {
	td := GetTestDiscogs()

	sale, err := td.GetSale(context.Background(), 2695553917)
	if err != nil {
		t.Fatalf("Error getting sale: %v", err)
	}

	if sale.GetStatus() != pb.SaleStatus_FOR_SALE {
		t.Errorf("Bad sale state: %v", sale)
	}

	if sale.GetReleaseId() != 27962688 {
		t.Errorf("Bad sale sate (wrong release id): %v", sale)
	}

	if sale.GetCondition() != "Mint (M)" {
		t.Errorf("Bad sale condition: %v", sale.GetCondition())
	}
}

func TestGetReleaseStats(t *testing.T) {
	td := GetTestDiscogs()
	stats, err := td.GetReleaseStats(context.Background(), 625928)
	if err != nil {
		t.Fatalf("bad get: %v", err)
	}

	if stats.GetMedianPrice() != 1647 {
		t.Errorf("Wrong median price should have been 1647, was %v", stats.GetMedianPrice())
	}

	if stats.GetLowPrice() != 922 {
		t.Errorf("Wrong low price; should have been 922, was %v", stats.GetLowPrice())
	}

	if stats.GetHighPrice() != 2407 {
		t.Errorf("Wrong low price; should have been 2407, was %v", stats.GetHighPrice())
	}
}

func TestGetReleaseStats_CornerCase(t *testing.T) {
	td := GetTestDiscogs()
	_, err := td.GetReleaseStats(context.Background(), 28154152)
	if err == nil || status.Code(err) != codes.NotFound {
		t.Fatalf("bad get: %v", err)
	}
}

func TestGetReleaseStats_CornerCase2(t *testing.T) {
	td := GetTestDiscogs()
	stats, err := td.GetReleaseStats(context.Background(), 1606771)
	if err != nil {
		t.Fatalf("bad get: %v", err)
	}

	if stats.GetMedianPrice() != 1347 {
		t.Errorf("Wrong median price should have been 0, was %v", stats.GetMedianPrice())
	}
}

func TestGetReleaseStats_Validation4243427(t *testing.T) {
	td := GetTestDiscogs()
	stats, err := td.GetReleaseStats(context.Background(), 4243427)
	if err != nil {
		t.Fatalf("Failed to get release stats for 4243427: %v", err)
	}

	if stats.GetLowPrice() != 581 {
		t.Errorf("Wrong low price: expected 581, got %v", stats.GetLowPrice())
	}
	if stats.GetMedianPrice() != 1855 {
		t.Errorf("Wrong median price: expected 1855, got %v", stats.GetMedianPrice())
	}
	if stats.GetHighPrice() != 3082 {
		t.Errorf("Wrong high price: expected 3082, got %v", stats.GetHighPrice())
	}
}

func TestGetReleaseStats_Validation556184(t *testing.T) {
	td := GetTestDiscogs()
	stats, err := td.GetReleaseStats(context.Background(), 556184)
	if err != nil {
		t.Fatalf("Failed to get release stats for 556184: %v", err)
	}

	if stats.GetLowPrice() != 698 {
		t.Errorf("Wrong low price: expected 698, got %v", stats.GetLowPrice())
	}
	if stats.GetMedianPrice() != 1369 {
		t.Errorf("Wrong median price: expected 1369, got %v", stats.GetMedianPrice())
	}
	if stats.GetHighPrice() != 9000 {
		t.Errorf("Wrong high price: expected 9000, got %v", stats.GetHighPrice())
	}
}


func TestGetReleaseStats_SingleSale(t *testing.T) {
	td := GetTestDiscogs()
	stats, err := td.GetReleaseStats(context.Background(), 9999991)
	if err != nil {
		t.Fatalf("Failed to get single sale release stats: %v", err)
	}

	if stats.GetLowPrice() != 1500 {
		t.Errorf("Wrong low price: expected 1500, got %v", stats.GetLowPrice())
	}
	if stats.GetMedianPrice() != 1500 {
		t.Errorf("Wrong median price: expected 1500, got %v", stats.GetMedianPrice())
	}
	if stats.GetHighPrice() != 1500 {
		t.Errorf("Wrong high price: expected 1500, got %v", stats.GetHighPrice())
	}
}

func TestGetReleaseStats_Malformed(t *testing.T) {
	td := GetTestDiscogs()
	_, err := td.GetReleaseStats(context.Background(), 9999992)
	if err == nil {
		t.Fatalf("Expected error for malformed release HTML without statistics, got nil")
	}
}

func TestGetReleaseStats_NextData(t *testing.T) {
	td := GetTestDiscogs()
	stats, err := td.GetReleaseStats(context.Background(), 9999993)
	if err != nil {
		t.Fatalf("Failed to get release stats for Next.js payload: %v", err)
	}

	if stats.GetLowPrice() != 1050 {
		t.Errorf("Wrong low price: expected 1050, got %v", stats.GetLowPrice())
	}
	if stats.GetMedianPrice() != 2000 {
		t.Errorf("Wrong median price: expected 2000, got %v", stats.GetMedianPrice())
	}
	if stats.GetHighPrice() != 3525 {
		t.Errorf("Wrong high price: expected 3525, got %v", stats.GetHighPrice())
	}
}

func TestGetReleaseStats_LegacySpans(t *testing.T) {
	td := GetTestDiscogs()
	stats, err := td.GetReleaseStats(context.Background(), 9999994)
	if err != nil {
		t.Fatalf("Failed to get release stats for legacy spans: %v", err)
	}

	if stats.GetLowPrice() != 450 {
		t.Errorf("Wrong low price: expected 450, got %v", stats.GetLowPrice())
	}
	if stats.GetMedianPrice() != 1200 {
		t.Errorf("Wrong median price: expected 1200, got %v", stats.GetMedianPrice())
	}
	if stats.GetHighPrice() != 2550 {
		t.Errorf("Wrong high price: expected 2550, got %v", stats.GetHighPrice())
	}
}

func TestGetReleaseStats_LegacyUnsold(t *testing.T) {
	td := GetTestDiscogs()
	_, err := td.GetReleaseStats(context.Background(), 9999995)
	if err == nil || status.Code(err) != codes.NotFound {
		t.Fatalf("Expected NotFound error for legacy unsold release, got %v", err)
	}
}


func TestListSales_PerPage100(t *testing.T) {
	td := GetTestDiscogs()

	sales, pagination, err := td.ListSales(context.Background(), 1)

	if err != nil {
		t.Fatalf("Failed to list sales: %v", err)
	}

	if len(sales) != 50 {
		t.Errorf("Bad sale return %v -> %v", sales, pagination)
	}

	if pagination.GetPage() != 1 || pagination.GetPages() != 35 {
		t.Errorf("Bad pagination return: %v", pagination)
	}

	for _, sale := range sales {
		if sale.GetReleaseId() == 9624074 {
			if sale.GetStatus() != pb.SaleStatus_SOLD || sale.GetSaleId() != 769427368 || sale.GetPrice().GetValue() != 1363 || sale.GetCondition() != "Very Good Plus (VG+)" {
				t.Errorf("No sale id returned: %v", sale)
			}
		}
	}
}

func TestGetOrder_Success(t *testing.T) {
	td := GetTestDiscogs()

	order, err := td.GetOrder(context.Background(), "150295-1254")

	if err != nil {
		t.Fatalf("Bad list sales: %v", err)
	}

	if order.Status != "Shipped" {
		t.Errorf("Bad error returned: %v", order)
	}
}

func TestUpdateSale_Success(t *testing.T) {
	td := GetTestDiscogs()

	err := td.UpdateSale(context.Background(), 2708115424, 19975519, "Near Mint (NM or M-)", 5655)

	if err != nil {
		t.Fatalf("Bad list sales: %v", err)
	}
}
