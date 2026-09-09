package discogs

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	pb "github.com/brotherlogic/discogs/proto"
)

func TestDiscogsTestClient(t *testing.T) {
	var d Discogs = &TestDiscogsClient{}
	log.Printf("TEST %v", d)
}

func TestTestDiscogsClient_ListOrders(t *testing.T) {
	client := GetTestClient()
	ctx := context.Background()

	// 1. Initially empty
	orders, pagination, err := client.ListOrders(ctx, time.Time{}, 1)
	if err != nil {
		t.Fatalf("ListOrders failed: %v", err)
	}
	if len(orders) != 0 {
		t.Errorf("Expected 0 orders, got %d", len(orders))
	}
	if pagination.GetPage() != 1 || pagination.GetPages() != 0 {
		t.Errorf("Expected pagination {Page: 1, Pages: 0}, got %+v", pagination)
	}

	// 2. Add mock orders with different created timestamps
	baseTime := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	for i := 1; i <= 150; i++ {
		client.AddOrder(&pb.Order{
			Id:      fmt.Sprintf("order-%d", i),
			Status:  "Payment Received",
			Created: baseTime.Add(time.Duration(i) * time.Hour).Unix(),
		})
	}

	// 3. Test pagination without filter (page 1: 100 items, page 2: 50 items, page 3: 0 items)
	p1Orders, p1Pag, err := client.ListOrders(ctx, time.Time{}, 1)
	if err != nil {
		t.Fatalf("ListOrders page 1 failed: %v", err)
	}
	if len(p1Orders) != 100 {
		t.Errorf("Expected 100 orders on page 1, got %d", len(p1Orders))
	}
	if p1Pag.GetPage() != 1 || p1Pag.GetPages() != 2 {
		t.Errorf("Expected page 1 pagination {Page: 1, Pages: 2}, got %+v", p1Pag)
	}
	if p1Orders[0].GetId() != "order-1" || p1Orders[99].GetId() != "order-100" {
		t.Errorf("Unexpected order ids on page 1: first=%s, last=%s", p1Orders[0].GetId(), p1Orders[99].GetId())
	}

	p2Orders, p2Pag, err := client.ListOrders(ctx, time.Time{}, 2)
	if err != nil {
		t.Fatalf("ListOrders page 2 failed: %v", err)
	}
	if len(p2Orders) != 50 {
		t.Errorf("Expected 50 orders on page 2, got %d", len(p2Orders))
	}
	if p2Pag.GetPage() != 2 || p2Pag.GetPages() != 2 {
		t.Errorf("Expected page 2 pagination {Page: 2, Pages: 2}, got %+v", p2Pag)
	}
	if p2Orders[0].GetId() != "order-101" || p2Orders[49].GetId() != "order-150" {
		t.Errorf("Unexpected order ids on page 2: first=%s, last=%s", p2Orders[0].GetId(), p2Orders[49].GetId())
	}

	p3Orders, p3Pag, err := client.ListOrders(ctx, time.Time{}, 3)
	if err != nil {
		t.Fatalf("ListOrders page 3 failed: %v", err)
	}
	if len(p3Orders) != 0 {
		t.Errorf("Expected 0 orders on page 3, got %d", len(p3Orders))
	}
	if p3Pag.GetPage() != 3 || p3Pag.GetPages() != 2 {
		t.Errorf("Expected page 3 pagination {Page: 3, Pages: 2}, got %+v", p3Pag)
	}

	// 4. Test filtering with createdAfter
	// Filter for orders created after baseTime + 100 hours (orders 101 to 150 = 50 orders)
	filterTime := baseTime.Add(101 * time.Hour)
	filteredOrders, filteredPag, err := client.ListOrders(ctx, filterTime, 1)
	if err != nil {
		t.Fatalf("ListOrders with filter failed: %v", err)
	}
	if len(filteredOrders) != 50 {
		t.Errorf("Expected 50 filtered orders, got %d", len(filteredOrders))
	}
	if filteredPag.GetPage() != 1 || filteredPag.GetPages() != 1 {
		t.Errorf("Expected filtered pagination {Page: 1, Pages: 1}, got %+v", filteredPag)
	}
	if filteredOrders[0].GetId() != "order-101" {
		t.Errorf("Expected first filtered order to be order-101, got %s", filteredOrders[0].GetId())
	}

	// 5. Test page <= 0 defaults to page 1
	p0Orders, p0Pag, err := client.ListOrders(ctx, time.Time{}, 0)
	if err != nil {
		t.Fatalf("ListOrders page 0 failed: %v", err)
	}
	if len(p0Orders) != 100 {
		t.Errorf("Expected 100 orders on page 0, got %d", len(p0Orders))
	}
	if p0Pag.GetPage() != 1 || p0Pag.GetPages() != 2 {
		t.Errorf("Expected page 0 pagination {Page: 1, Pages: 2}, got %+v", p0Pag)
	}
}
