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

	if saleid != 2695553917 {
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
}
