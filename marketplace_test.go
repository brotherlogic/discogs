package discogs

import (
	"context"
	"testing"
)

func TestCreateSale_Success(t *testing.T) {
	td := GetTestDiscogs()

	saleid, err := td.CreateSale(context.Background(), SaleParams{
		ReleaseId: 27962688,
	})

	if err != nil {
		t.Fatalf("Error creating sale: %v", err)
	}

	if saleid != 2695549717 {
		t.Errorf("Bad saleid return: %v", saleid)
	}
}
