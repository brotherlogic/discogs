package proto

import (
	"testing"
)

func TestOrderProto(t *testing.T) {
	var nilOrder *Order
	if nilOrder.GetId() != "" || nilOrder.GetStatus() != "" || nilOrder.GetCreated() != 0 || nilOrder.GetItems() != nil || nilOrder.GetTotal() != nil || nilOrder.GetLastActivity() != 0 {
		t.Errorf("nil Order getter returned unexpected non-zero value")
	}

	var nilOrderItem *OrderItem
	if nilOrderItem.GetId() != 0 || nilOrderItem.GetReleaseId() != 0 || nilOrderItem.GetPrice() != nil || nilOrderItem.GetCondition() != "" || nilOrderItem.GetSleeveCondition() != "" {
		t.Errorf("nil OrderItem getter returned unexpected non-zero value")
	}

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

	_ = order.String()
	_ = order.ProtoReflect()
	_ = item.String()
	_ = item.ProtoReflect()

	orderCopy := &Order{Id: "test"}
	orderCopy.Reset()
	if orderCopy.GetId() != "" {
		t.Errorf("Reset failed for Order")
	}

	itemCopy := &OrderItem{Id: 100}
	itemCopy.Reset()
	if itemCopy.GetId() != 0 {
		t.Errorf("Reset failed for OrderItem")
	}
}

func TestAllProtosCoverage(t *testing.T) {
	field := &Field{Id: 1, Name: "name"}
	_ = field.GetId()
	_ = field.GetName()
	_ = field.String()
	_ = field.ProtoReflect()
	var nilField *Field
	_ = nilField.GetId()
	_ = nilField.GetName()
	field.Reset()

	folder := &Folder{Id: 1, Name: "folder"}
	_ = folder.GetId()
	_ = folder.GetName()
	_ = folder.String()
	_ = folder.ProtoReflect()
	var nilFolder *Folder
	_ = nilFolder.GetId()
	_ = nilFolder.GetName()
	folder.Reset()

	mr := &MasterRelease{Id: 1, Year: 1999}
	_ = mr.GetId()
	_ = mr.GetYear()
	_ = mr.String()
	_ = mr.ProtoReflect()
	var nilMr *MasterRelease
	_ = nilMr.GetId()
	_ = nilMr.GetYear()
	mr.Reset()

	pag := &Pagination{Page: 1, Pages: 10}
	_ = pag.GetPage()
	_ = pag.GetPages()
	_ = pag.String()
	_ = pag.ProtoReflect()
	var nilPag *Pagination
	_ = nilPag.GetPage()
	_ = nilPag.GetPages()
	pag.Reset()

	format := &Format{Descriptions: []string{"LP"}, Name: "Vinyl", Quantity: 1}
	_ = format.GetDescriptions()
	_ = format.GetName()
	_ = format.GetQuantity()
	_ = format.String()
	_ = format.ProtoReflect()
	var nilFormat *Format
	_ = nilFormat.GetDescriptions()
	_ = nilFormat.GetName()
	_ = nilFormat.GetQuantity()
	format.Reset()

	user := &User{DiscogsUserId: 1, Username: "user", UserToken: "token", UserSecret: "secret", PersonalToken: "ptoken"}
	_ = user.GetDiscogsUserId()
	_ = user.GetUsername()
	_ = user.GetUserToken()
	_ = user.GetUserSecret()
	_ = user.GetPersonalToken()
	_ = user.String()
	_ = user.ProtoReflect()
	var nilUser *User
	_ = nilUser.GetDiscogsUserId()
	_ = nilUser.GetUsername()
	_ = nilUser.GetUserToken()
	_ = nilUser.GetUserSecret()
	_ = nilUser.GetPersonalToken()
	user.Reset()

	rel := &Release{
		Id: 1, InstanceId: 2, FolderId: 3, Rating: 4, Title: "title", MasterId: 5,
		Condition: "NM", SleeveCondition: "Generic", Formats: []*Format{format},
		Labels: []*Label{{Id: 1, Name: "label", Catno: "cat"}},
		Artists: []*Artist{{Name: "artist", Id: 1}},
		Notes: map[int32]string{1: "note"}, ReleaseDate: 123, DateAdded: 456,
	}
	_ = rel.GetId()
	_ = rel.GetInstanceId()
	_ = rel.GetFolderId()
	_ = rel.GetRating()
	_ = rel.GetTitle()
	_ = rel.GetMasterId()
	_ = rel.GetCondition()
	_ = rel.GetSleeveCondition()
	_ = rel.GetFormats()
	_ = rel.GetLabels()
	_ = rel.GetArtists()
	_ = rel.GetNotes()
	_ = rel.GetReleaseDate()
	_ = rel.GetDateAdded()
	_ = rel.String()
	_ = rel.ProtoReflect()
	var nilRel *Release
	_ = nilRel.GetId()
	_ = nilRel.GetInstanceId()
	_ = nilRel.GetFolderId()
	_ = nilRel.GetRating()
	_ = nilRel.GetTitle()
	_ = nilRel.GetMasterId()
	_ = nilRel.GetCondition()
	_ = nilRel.GetSleeveCondition()
	_ = nilRel.GetFormats()
	_ = nilRel.GetLabels()
	_ = nilRel.GetArtists()
	_ = nilRel.GetNotes()
	_ = nilRel.GetReleaseDate()
	_ = nilRel.GetDateAdded()
	rel.Reset()

	rs := &ReleaseStats{MedianPrice: 10, LowPrice: 5, HighPrice: 20}
	_ = rs.GetMedianPrice()
	_ = rs.GetLowPrice()
	_ = rs.GetHighPrice()
	_ = rs.String()
	_ = rs.ProtoReflect()
	var nilRs *ReleaseStats
	_ = nilRs.GetMedianPrice()
	_ = nilRs.GetLowPrice()
	_ = nilRs.GetHighPrice()
	rs.Reset()

	lbl := &Label{Id: 1, Name: "Label", Catno: "CAT01"}
	_ = lbl.GetId()
	_ = lbl.GetName()
	_ = lbl.GetCatno()
	_ = lbl.String()
	_ = lbl.ProtoReflect()
	var nilLbl *Label
	_ = nilLbl.GetId()
	_ = nilLbl.GetName()
	_ = nilLbl.GetCatno()
	lbl.Reset()

	pr := &Price{Currency: "USD", Value: 100}
	_ = pr.GetCurrency()
	_ = pr.GetValue()
	_ = pr.String()
	_ = pr.ProtoReflect()
	var nilPr *Price
	_ = nilPr.GetCurrency()
	_ = nilPr.GetValue()
	pr.Reset()

	si := &SaleItem{SaleId: 1, Status: SaleStatus_FOR_SALE, Price: pr, ReleaseId: 2, Condition: "VG+"}
	_ = si.GetSaleId()
	_ = si.GetStatus()
	_ = si.GetPrice()
	_ = si.GetReleaseId()
	_ = si.GetCondition()
	_ = si.String()
	_ = si.ProtoReflect()
	var nilSi *SaleItem
	_ = nilSi.GetSaleId()
	_ = nilSi.GetStatus()
	_ = nilSi.GetPrice()
	_ = nilSi.GetReleaseId()
	_ = nilSi.GetCondition()
	si.Reset()

	art := &Artist{Name: "Artist", Id: 1}
	_ = art.GetName()
	_ = art.GetId()
	_ = art.String()
	_ = art.ProtoReflect()
	var nilArt *Artist
	_ = nilArt.GetName()
	_ = nilArt.GetId()
	art.Reset()

	w := &Want{Id: 1, Title: "Want", Artists: []*Artist{art}}
	_ = w.GetId()
	_ = w.GetTitle()
	_ = w.GetArtists()
	_ = w.String()
	_ = w.ProtoReflect()
	var nilW *Want
	_ = nilW.GetId()
	_ = nilW.GetTitle()
	_ = nilW.GetArtists()
	w.Reset()

	sparams := &SaleParams{
		ReleaseId: 1, Condition: "VG", SleeveCondition: "VG", Price: 10.5,
		Comments: "Nice", AllowOffers: true, Status: "For Sale",
		ExternalId: "ext1", Location: "bin1", Weight: 100, FormatQuantity: 1,
	}
	_ = sparams.GetReleaseId()
	_ = sparams.GetCondition()
	_ = sparams.GetSleeveCondition()
	_ = sparams.GetPrice()
	_ = sparams.GetComments()
	_ = sparams.GetAllowOffers()
	_ = sparams.GetStatus()
	_ = sparams.GetExternalId()
	_ = sparams.GetLocation()
	_ = sparams.GetWeight()
	_ = sparams.GetFormatQuantity()
	_ = sparams.String()
	_ = sparams.ProtoReflect()
	var nilSparams *SaleParams
	_ = nilSparams.GetReleaseId()
	_ = nilSparams.GetCondition()
	_ = nilSparams.GetSleeveCondition()
	_ = nilSparams.GetPrice()
	_ = nilSparams.GetComments()
	_ = nilSparams.GetAllowOffers()
	_ = nilSparams.GetStatus()
	_ = nilSparams.GetExternalId()
	_ = nilSparams.GetLocation()
	_ = nilSparams.GetWeight()
	_ = nilSparams.GetFormatQuantity()
	sparams.Reset()

	sstats := &SaleStats{
		VgPrice: 1.0, GplusPrice: 2.0, NmPrice: 3.0, GPrice: 4.0,
		VgplusPrice: 5.0, MPrice: 6.0, FPrice: 7.0, PPrice: 8.0,
	}
	_ = sstats.GetVgPrice()
	_ = sstats.GetGplusPrice()
	_ = sstats.GetNmPrice()
	_ = sstats.GetGPrice()
	_ = sstats.GetVgplusPrice()
	_ = sstats.GetMPrice()
	_ = sstats.GetFPrice()
	_ = sstats.GetPPrice()
	_ = sstats.String()
	_ = sstats.ProtoReflect()
	var nilSstats *SaleStats
	_ = nilSstats.GetVgPrice()
	_ = nilSstats.GetGplusPrice()
	_ = nilSstats.GetNmPrice()
	_ = nilSstats.GetGPrice()
	_ = nilSstats.GetVgplusPrice()
	_ = nilSstats.GetMPrice()
	_ = nilSstats.GetFPrice()
	_ = nilSstats.GetPPrice()
	sstats.Reset()

	_ = MasterSort_BY_YEAR.Enum()
	_ = MasterSort_BY_YEAR.Number()
	_ = MasterSort_BY_YEAR.String()
	_ = MasterSort_BY_YEAR.Type()
	_ = MasterSort_BY_YEAR.Descriptor()

	_ = SaleStatus_FOR_SALE.Enum()
	_ = SaleStatus_FOR_SALE.Number()
	_ = SaleStatus_FOR_SALE.String()
	_ = SaleStatus_FOR_SALE.Type()
	_ = SaleStatus_FOR_SALE.Descriptor()
}
