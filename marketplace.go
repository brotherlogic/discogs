package discogs

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

type CreateSaleResponse struct {
	ListingId int64 `json:"listing_id"`
}

type SaleJson struct {
	ReleaseId int64   `json:"release_id"`
	Condition string  `json:"condition"`
	Price     float32 `json:"price"`
}

func (p *prodClient) CreateSale(ctx context.Context, params SaleParams) (int64, error) {
	csURL := fmt.Sprintf("/marketplace/listings")

	data := &SaleJson{
		ReleaseId: int64(params.ReleaseId),
		Condition: params.Condition,
		Price:     float32(params.Price) / 100,
	}
	v, err := json.Marshal(data)
	if err != nil {
		return -1, err
	}

	csr := &CreateSaleResponse{}
	log.Printf("%v", string(v))
	err = p.makeDiscogsRequest("POST", csURL, string(v), csr)
	if err != nil {
		return -1, err
	}

	return csr.ListingId, nil
}
