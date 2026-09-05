package discogs

import (
	"context"
	"testing"
	"time"

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

func TestListOrders_Parsing(t *testing.T) {
	td := GetTestDiscogs()

	orders, pagination, err := td.ListOrders(context.Background(), time.Time{}, 1)
	if err != nil {
		t.Fatalf("Failed to list orders: %v", err)
	}

	if pagination.GetPage() != 1 || pagination.GetPages() != 3 {
		t.Errorf("Bad pagination: got page=%v, pages=%v; expected 1 and 3", pagination.GetPage(), pagination.GetPages())
	}

	if len(orders) != 2 {
		t.Fatalf("Expected 2 orders, got %d", len(orders))
	}

	// First order
	o1 := orders[0]
	if o1.GetId() != "150295-1254" || o1.GetStatus() != "Shipped" {
		t.Errorf("Unexpected order 1 metadata: %+v", o1)
	}
	expectedCreated1, _ := time.Parse(time.RFC3339, "2023-08-19T05:25:38-07:00")
	if o1.GetCreated() != expectedCreated1.Unix() {
		t.Errorf("Expected created unix %v, got %v", expectedCreated1.Unix(), o1.GetCreated())
	}
	expectedLastActivity1, _ := time.Parse(time.RFC3339, "2023-08-25T20:05:05-07:00")
	if o1.GetLastActivity() != expectedLastActivity1.Unix() {
		t.Errorf("Expected last_activity unix %v, got %v", expectedLastActivity1.Unix(), o1.GetLastActivity())
	}
	if o1.GetTotal().GetValue() != 1555 || o1.GetTotal().GetCurrency() != "USD" {
		t.Errorf("Expected total 1555 USD, got %+v", o1.GetTotal())
	}
	if len(o1.GetItems()) != 1 {
		t.Fatalf("Expected 1 item in order 1, got %d", len(o1.GetItems()))
	}
	item1 := o1.GetItems()[0]
	if item1.GetId() != 2565159678 || item1.GetReleaseId() != 8419904 || item1.GetCondition() != "Very Good Plus (VG+)" || item1.GetSleeveCondition() != "Very Good Plus (VG+)" {
		t.Errorf("Unexpected item 1: %+v", item1)
	}
	if item1.GetPrice().GetValue() != 1555 || item1.GetPrice().GetCurrency() != "USD" {
		t.Errorf("Expected item price 1555 USD, got %+v", item1.GetPrice())
	}

	// Second order
	o2 := orders[1]
	if o2.GetId() != "150295-1255" || o2.GetStatus() != "Payment Received" {
		t.Errorf("Unexpected order 2 metadata: %+v", o2)
	}
	expectedCreated2, _ := time.Parse(time.RFC3339, "2023-08-20T10:00:00Z")
	if o2.GetCreated() != expectedCreated2.Unix() {
		t.Errorf("Expected created unix %v, got %v", expectedCreated2.Unix(), o2.GetCreated())
	}
	if o2.GetTotal().GetValue() != 3000 || o2.GetTotal().GetCurrency() != "EUR" {
		t.Errorf("Expected total 3000 EUR, got %+v", o2.GetTotal())
	}
	if len(o2.GetItems()) != 1 {
		t.Fatalf("Expected 1 item in order 2, got %d", len(o2.GetItems()))
	}
	item2 := o2.GetItems()[0]
	if item2.GetId() != 2565159679 || item2.GetReleaseId() != 123456 || item2.GetCondition() != "Near Mint (NM or M-)" || item2.GetSleeveCondition() != "Generic" {
		t.Errorf("Unexpected item 2: %+v", item2)
	}
	if item2.GetPrice().GetValue() != 3000 || item2.GetPrice().GetCurrency() != "EUR" {
		t.Errorf("Expected item price 3000 EUR, got %+v", item2.GetPrice())
	}
}

func TestListOrders_CreatedAfterFormatting(t *testing.T) {
	td := GetTestDiscogs()

	// With createdAfter set
	createdAfter := time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC)
	orders, pagination, err := td.ListOrders(context.Background(), createdAfter, 1)
	if err != nil {
		t.Fatalf("Failed to list orders with created_after: %v", err)
	}
	if len(orders) != 1 || orders[0].GetId() != "150295-1255" {
		t.Errorf("Expected 1 order with id 150295-1255, got %+v", orders)
	}
	if pagination.GetPage() != 1 || pagination.GetPages() != 1 {
		t.Errorf("Unexpected pagination: %+v", pagination)
	}

	// Without createdAfter (zero time)
	ordersZero, paginationZero, err := td.ListOrders(context.Background(), time.Time{}, 1)
	if err != nil {
		t.Fatalf("Failed to list orders with zero time: %v", err)
	}
	if len(ordersZero) != 2 {
		t.Errorf("Expected 2 orders with zero time, got %d", len(ordersZero))
	}
	if paginationZero.GetPage() != 1 || paginationZero.GetPages() != 3 {
		t.Errorf("Unexpected pagination with zero time: %+v", paginationZero)
	}
}

func TestListOrders_EmptyAndError(t *testing.T) {
	td := GetTestDiscogs()

	// Empty orders
	orders, pagination, err := td.ListOrders(context.Background(), time.Time{}, 99)
	if err != nil {
		t.Fatalf("Failed to list orders on empty page: %v", err)
	}
	if len(orders) != 0 {
		t.Errorf("Expected 0 orders on page 99, got %d", len(orders))
	}
	if pagination.GetPage() != 99 || pagination.GetPages() != 3 {
		t.Errorf("Unexpected pagination on page 99: %+v", pagination)
	}

	// Error case - invalid/unmocked page that causes error
	_, _, err = td.ListOrders(context.Background(), time.Time{}, 555)
	if err == nil {
		t.Fatalf("Expected error when requesting non-existent fixture, got nil")
	}
}
