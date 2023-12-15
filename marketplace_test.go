package discogs

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/discogs/proto"
)

func TestCreateSale_Success(t *testing.T) {
	td := GetTestDiscogs()

	saleid, err := td.CreateSale(context.Background(), SaleParams{
		ReleaseId: 27962688,
		Condition: "Mint (M)",
		Price:     10023,
	})

	if err != nil {
		t.Fatalf("Error creating sale: %v", err)
	}

	if saleid != 2851616155 {
		t.Errorf("Bad saleid return: %v", saleid)
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

	if stats.GetMedianPrice() != 1578 {
		t.Errorf("Wrong median price should have been 1578, was %v", stats.GetMedianPrice())
	}
}

func TestGetReleaseStats_CornerCase(t *testing.T) {
	td := GetTestDiscogs()
	stats, err := td.GetReleaseStats(context.Background(), 28154152)
	if err != nil {
		t.Fatalf("bad get: %v", err)
	}

	if stats.GetMedianPrice() != 0 {
		t.Errorf("Wrong median price should have been 0, was %v", stats.GetMedianPrice())
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

func TestListSales_Success(t *testing.T) {
	td := GetTestDiscogs()

	sales, pagination, err := td.ListSales(context.Background(), 1)

	if err != nil {
		t.Fatalf("Bad list sales: %v", err)
	}

	if len(sales) != 50 {
		t.Errorf("Bad sale return %v -> %v", sales, pagination)
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
