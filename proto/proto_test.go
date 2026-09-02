package proto

import (
	"testing"
)

func TestOrderProto(t *testing.T) {
	order := &Order{
		Id:      "12345-6789",
		Status:  "Payment Received",
		Created: 1700000000,
		Items: []*OrderItem{
			{
				Id:              101,
				ReleaseId:       202,
				Price:           &Price{Currency: "USD", Value: 2500},
				Condition:       "Near Mint (NM or M-)",
				SleeveCondition: "Generic",
			},
		},
		Total:        &Price{Currency: "USD", Value: 2500},
		LastActivity: 1700000500,
	}

	if order.GetId() != "12345-6789" {
		t.Errorf("Expected id 12345-6789, got %v", order.GetId())
	}
	if order.GetStatus() != "Payment Received" {
		t.Errorf("Expected status Payment Received, got %v", order.GetStatus())
	}
	if order.GetCreated() != 1700000000 {
		t.Errorf("Expected created 1700000000, got %v", order.GetCreated())
	}
	if len(order.GetItems()) != 1 {
		t.Fatalf("Expected 1 item, got %d", len(order.GetItems()))
	}
	item := order.GetItems()[0]
	if item.GetId() != 101 || item.GetReleaseId() != 202 {
		t.Errorf("Unexpected item fields: %+v", item)
	}
	if item.GetPrice().GetValue() != 2500 || item.GetPrice().GetCurrency() != "USD" {
		t.Errorf("Unexpected item price: %+v", item.GetPrice())
	}
	if item.GetCondition() != "Near Mint (NM or M-)" || item.GetSleeveCondition() != "Generic" {
		t.Errorf("Unexpected item condition: %+v", item)
	}
	if order.GetTotal().GetValue() != 2500 || order.GetTotal().GetCurrency() != "USD" {
		t.Errorf("Unexpected order total: %+v", order.GetTotal())
	}
	if order.GetLastActivity() != 1700000500 {
		t.Errorf("Expected last activity 1700000500, got %v", order.GetLastActivity())
	}
}
