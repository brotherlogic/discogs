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

	if saleid != 2696998546 {
		t.Errorf("Bad saleid return: %v", saleid)
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
			if sale.GetStatus() != pb.SaleStatus_SOLD || sale.GetSaleId() != 769427368 || sale.GetPrice().GetValue() != 1363 {
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

	err := td.UpdateSale(context.Background(), 2708115424, 5655)

	if err != nil {
		t.Fatalf("Bad list sales: %v", err)
	}
}
