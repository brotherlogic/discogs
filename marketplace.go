package discogs

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"

	pb "github.com/brotherlogic/discogs/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateSaleResponse struct {
	ListingId int64 `json:"listing_id"`
}

type Price struct {
	Value    float32
	Currency string
}

type GetSaleResponse struct {
	Status    string
	Id        int64
	Release   Release
	Price     Price
	Condition string
}

type Release struct {
	Id int64
}

type InventoryResponse struct {
	Pagination Pagination
	Listings   []GetSaleResponse
}

func convertSaleStatus(status pb.SaleStatus) string {
	switch status {
	case pb.SaleStatus_FOR_SALE:
		return "For Sale"
	case pb.SaleStatus_DRAFT:
		return "Draft"
	case pb.SaleStatus_EXPIRED:
		return "Expired"
	}

	log.Fatalf("Unknown Sale State: %v", status)
	return "unknown"
}

func convertStatus(status string) pb.SaleStatus {
	switch status {
	case "For Sale":
		return pb.SaleStatus_FOR_SALE
	case "Violation":
		return pb.SaleStatus_VIOLATION
	case "Sold":
		return pb.SaleStatus_SOLD
	case "Draft":
		return pb.SaleStatus_DRAFT
	case "Expired":
		return pb.SaleStatus_EXPIRED
	}

	log.Fatalf("Unknown Sale State: %v", status)
	return pb.SaleStatus_UNKNOWN
}

func convertPrice(price Price) *pb.Price {
	return &pb.Price{
		Value:    int32(price.Value * 100),
		Currency: price.Currency,
	}
}

type strpass struct {
	Value string
}

func (p *prodClient) GetReleaseStats(ctx context.Context, releaseId int64) (*pb.ReleaseStats, error) {
	url := fmt.Sprintf("https://www.discogs.com/release/%v", releaseId)
	str := &strpass{}
	err := p.makeDiscogsRequest("SGET", url, "", "release", str)
	if err != nil {
		return nil, err
	}

	rs := &pb.ReleaseStats{}

	results := regexp.MustCompile("Median<!.*?span.*?span>(.*?)<").FindAllStringSubmatch(str.Value, 1)
	if len(results) > 0 && len(results[0]) > 0 {
		strvl := results[0][1]

		// Release has no median price
		if strvl == "--" {
			return &pb.ReleaseStats{MedianPrice: 0}, nil
		}

		num, err := strconv.ParseFloat(strvl[1:], 16)
		if err != nil {
			return nil, err
		}
		rs.MedianPrice = int32(num * 100)
	}

	results = regexp.MustCompile("Low<!.*?span.*?span>(.*?)<").FindAllStringSubmatch(str.Value, 1)
	if len(results) > 0 && len(results[0]) > 0 {
		strvl := results[0][1]

		// Release has no median price
		if strvl == "--" {
			return &pb.ReleaseStats{MedianPrice: 0}, nil
		}

		num, err := strconv.ParseFloat(strvl[1:], 16)
		if err != nil {
			return nil, err
		}
		rs.LowPrice = int32(num * 100)
	}

	results = regexp.MustCompile("High<!.*?span.*?span>(.*?)<").FindAllStringSubmatch(str.Value, 1)
	if len(results) > 0 && len(results[0]) > 0 {
		strvl := results[0][1]

		// Release has no median price
		if strvl == "--" {
			return &pb.ReleaseStats{MedianPrice: 0}, nil
		}

		num, err := strconv.ParseFloat(strvl[1:], 16)
		if err != nil {
			return nil, err
		}
		rs.HighPrice = int32(num * 100)
	}

	return rs, nil
}

func (p *prodClient) ListSales(ctx context.Context, page int32) ([]*pb.SaleItem, *pb.Pagination, error) {
	cr := &InventoryResponse{}
	err := p.makeDiscogsRequest(
		"GET",
		fmt.Sprintf("/users/%v/inventory?page=%v", p.user.GetUsername(), page),
		"",
		"/users/uname/inventory",
		cr,
	)
	if err != nil {
		return nil, nil, err
	}

	var listings []*pb.SaleItem
	for _, listing := range cr.Listings {
		listings = append(listings, &pb.SaleItem{
			ReleaseId: listing.Release.Id,
			Status:    convertStatus(listing.Status),
			SaleId:    listing.Id,
			Price:     convertPrice(listing.Price),
			Condition: listing.Condition,
		})
	}

	return listings, &pb.Pagination{Page: int32(cr.Pagination.Page), Pages: int32(cr.Pagination.Pages)}, nil
}

type OrderItem struct {
	Release Release
	Price   Price
	Id      int64
}

type Order struct {
	Id     string
	Status string
	Items  []OrderItem
}

func (p *prodClient) GetOrder(ctx context.Context, orderId string) (*pb.Order, error) {
	gsURL := fmt.Sprintf("/marketplace/orders/%v", orderId)

	gsr := &Order{}
	err := p.makeDiscogsRequest(
		"GET",
		gsURL,
		"",
		"/marketplace/orders/oid",
		gsr,
	)

	if err != nil {
		return nil, err
	}

	return &pb.Order{
		Id:     gsr.Id,
		Status: gsr.Status,
	}, nil
}

func (p *prodClient) GetSale(ctx context.Context, saleId int64) (*pb.SaleItem, error) {
	gsURL := fmt.Sprintf("/marketplace/listings/%v", saleId)

	gsr := &GetSaleResponse{}
	err := p.makeDiscogsRequest(
		"GET",
		gsURL,
		"",
		"/marketplace/listings/sid",
		gsr,
	)
	if err != nil {
		return nil, err
	}

	return &pb.SaleItem{
		Status:    convertStatus(gsr.Status),
		ReleaseId: (gsr.Release.Id),
		Condition: gsr.Condition,
	}, nil
}

func (p *prodClient) CreateSale(ctx context.Context, params *pb.SaleParams) (int64, error) {
	csURL := fmt.Sprintf("/marketplace/listings")

	// Validate the sale parameters
	if params.ReleaseId == 0 {
		return 0, status.Errorf(codes.InvalidArgument, "No release ID provided")
	}
	if params.Condition == "" {
		return 0, status.Errorf(codes.InvalidArgument, "No condition provided")
	}
	if params.Price == 0 {
		return 0, status.Errorf(codes.InvalidArgument, "No price provided")
	}
	if params.Status == "" {
		return 0, status.Errorf(codes.InvalidArgument, "No status provided")
	}

	if params.Status != "For Sale" && params.Status != "Draft" {
		return 0, status.Errorf(codes.InvalidArgument, "Invalid status provided: %v", params.Status)
	}

	v, err := json.Marshal(params)
	if err != nil {
		return -1, err
	}

	csr := &CreateSaleResponse{}
	log.Printf("%v", string(v))
	err = p.makeDiscogsRequest(
		"POST",
		csURL,
		string(v),
		"/marketplace/listings",
		csr)
	if err != nil {
		return -1, err
	}

	return csr.ListingId, nil
}

func (p *prodClient) UpdateSale(ctx context.Context, saleId int64, releaseId int64, condition string, newPrice int32) error {
	csURL := fmt.Sprintf("/marketplace/listings/%v", saleId)

	data := &pb.SaleParams{
		Price:     float32(newPrice) / 100,
		ReleaseId: releaseId,
		Condition: condition,
		Status:    "For Sale", // Assumed that sale updates are for sale items
	}
	v, err := json.Marshal(data)
	if err != nil {
		return err
	}

	csr := &CreateSaleResponse{}
	log.Printf("%v", string(v))
	err = p.makeDiscogsRequest("POST", csURL, string(v), "/marketplace/listings/sid", csr)
	if err != nil {
		return err
	}

	return nil
}

func (p *prodClient) UpdateSaleState(ctx context.Context, saleId int64, releaseId int64, condition string, saleState pb.SaleStatus) error {
	csURL := fmt.Sprintf("/marketplace/listings/%v", saleId)

	data := &pb.SaleParams{
		Status:    convertSaleStatus(saleState),
		ReleaseId: releaseId,
		Condition: condition,
	}
	v, err := json.Marshal(data)
	if err != nil {
		return err
	}

	csr := &CreateSaleResponse{}
	log.Printf("%v", string(v))
	err = p.makeDiscogsRequest("POST", csURL, string(v), "/marketplace/listings/sid", csr)
	if err != nil {
		return err
	}

	return nil
}

type DiscogsPrice struct {
	Currency string  `json:"currency"`
	Value    float32 `json:"value"`
}

type SaleStatsResponse struct {
	VG  DiscogsPrice `json:"Very Good (VG)"`
	M   DiscogsPrice `json:"Mint (M)"`
	NM  DiscogsPrice `json: "Near Mint (NM or M-)"`
	VGP DiscogsPrice `json: "Very Good Plus (VG+)"`
	GP  DiscogsPrice `json: "Good Plus (G+)"`
	G   DiscogsPrice `json: "Good (G)"`
	F   DiscogsPrice `json: "Fair (F)"`
	P   DiscogsPrice `json: "Poor (P)"`
}

func (p *prodClient) GetSaleStats(ctx context.Context, releaseId int64) (*pb.SaleStats, error) {
	csURL := fmt.Sprintf("/marketplace/price_suggestions/%v", releaseId)
	ssr := &SaleStatsResponse{}
	err := p.makeDiscogsRequest("GET", csURL, "", "/marketplace/price_suggestions/rid", ssr)
	if err != nil {
		return nil, err
	}

	return &pb.SaleStats{
		VgPrice:     ssr.VG.Value,
		MPrice:      ssr.M.Value,
		NmPrice:     ssr.NM.Value,
		VgplusPrice: ssr.VG.Value,
		GplusPrice:  ssr.GP.Value,
		GPrice:      ssr.G.Value,
		FPrice:      ssr.F.Value,
		PPrice:      ssr.P.Value,
	}, nil
}
