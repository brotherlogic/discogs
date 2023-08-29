package discogs

import (
	"context"
	"fmt"
)

type CreateSaleResponse struct {
	ListingId int64 `json:"listing_id"`
}

func (p *prodClient) CreateSale(ctx context.Context, params SaleParams) (int64, error) {
	csURL := fmt.Sprintf("/marketplace/listings")

	csr := &CreateSaleResponse{}
	err := p.makeDiscogsRequest("POST", csURL, "", csr)
	if err != nil {
		return -1, err
	}

	return csr.ListingId, nil
}
