package discogs

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	pb "github.com/brotherlogic/discogs/proto"
)

type CreateSaleResponse struct {
	ListingId int64 `json:"listing_id"`
}

type SaleJson struct {
	ReleaseId int64   `json:"release_id"`
	Condition string  `json:"condition"`
	Price     float32 `json:"price"`
}

type GetSaleResponse struct {
	Status  string
	Release Release
}

type Release struct {
	Id int64
}

func convertStatus(status string) pb.SaleStatus {
	switch status {
	case "For Sale":
		return pb.SaleStatus_FOR_SALE
	}

	log.Fatalf("Unknown Sale State: %v", status)
	return pb.SaleStatus_UNKNOWN
}

func (p *prodClient) GetSale(ctx context.Context, saleId int64) (*pb.SaleItem, error) {
	gsURL := fmt.Sprintf("/marketplace/listings/%v", saleId)

	gsr := &GetSaleResponse{}
	err := p.makeDiscogsRequest("GET", gsURL, "", gsr)
	if err != nil {
		return nil, err
	}

	return &pb.SaleItem{
		Status:    convertStatus(gsr.Status),
		ReleaseId: (gsr.Release.Id),
	}, nil
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
