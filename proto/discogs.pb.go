// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: discogs.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SaleStatus int32

const (
	SaleStatus_UNKNOWN   SaleStatus = 0
	SaleStatus_FOR_SALE  SaleStatus = 1
	SaleStatus_SOLD      SaleStatus = 2
	SaleStatus_VIOLATION SaleStatus = 3
)

// Enum value maps for SaleStatus.
var (
	SaleStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "FOR_SALE",
		2: "SOLD",
		3: "VIOLATION",
	}
	SaleStatus_value = map[string]int32{
		"UNKNOWN":   0,
		"FOR_SALE":  1,
		"SOLD":      2,
		"VIOLATION": 3,
	}
)

func (x SaleStatus) Enum() *SaleStatus {
	p := new(SaleStatus)
	*p = x
	return p
}

func (x SaleStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SaleStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_discogs_proto_enumTypes[0].Descriptor()
}

func (SaleStatus) Type() protoreflect.EnumType {
	return &file_discogs_proto_enumTypes[0]
}

func (x SaleStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SaleStatus.Descriptor instead.
func (SaleStatus) EnumDescriptor() ([]byte, []int) {
	return file_discogs_proto_rawDescGZIP(), []int{0}
}

type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discogs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_discogs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_discogs_proto_rawDescGZIP(), []int{0}
}

func (x *Field) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Field) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Folder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Folder) Reset() {
	*x = Folder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discogs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Folder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Folder) ProtoMessage() {}

func (x *Folder) ProtoReflect() protoreflect.Message {
	mi := &file_discogs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Folder.ProtoReflect.Descriptor instead.
func (*Folder) Descriptor() ([]byte, []int) {
	return file_discogs_proto_rawDescGZIP(), []int{1}
}

func (x *Folder) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Folder) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Pages int32 `protobuf:"varint,2,opt,name=pages,proto3" json:"pages,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discogs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_discogs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_discogs_proto_rawDescGZIP(), []int{2}
}

func (x *Pagination) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Pagination) GetPages() int32 {
	if x != nil {
		return x.Pages
	}
	return 0
}

type Format struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Descriptions []string `protobuf:"bytes,1,rep,name=descriptions,proto3" json:"descriptions,omitempty"`
	Name         string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Quantity     int32    `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Format) Reset() {
	*x = Format{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discogs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Format) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Format) ProtoMessage() {}

func (x *Format) ProtoReflect() protoreflect.Message {
	mi := &file_discogs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Format.ProtoReflect.Descriptor instead.
func (*Format) Descriptor() ([]byte, []int) {
	return file_discogs_proto_rawDescGZIP(), []int{3}
}

func (x *Format) GetDescriptions() []string {
	if x != nil {
		return x.Descriptions
	}
	return nil
}

func (x *Format) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Format) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiscogsUserId int32  `protobuf:"varint,1,opt,name=discogs_user_id,json=discogsUserId,proto3" json:"discogs_user_id,omitempty"`
	Username      string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	UserToken     string `protobuf:"bytes,3,opt,name=user_token,json=userToken,proto3" json:"user_token,omitempty"`
	UserSecret    string `protobuf:"bytes,4,opt,name=user_secret,json=userSecret,proto3" json:"user_secret,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discogs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_discogs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_discogs_proto_rawDescGZIP(), []int{4}
}

func (x *User) GetDiscogsUserId() int32 {
	if x != nil {
		return x.DiscogsUserId
	}
	return 0
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetUserToken() string {
	if x != nil {
		return x.UserToken
	}
	return ""
}

func (x *User) GetUserSecret() string {
	if x != nil {
		return x.UserSecret
	}
	return ""
}

type Release struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	InstanceId int64     `protobuf:"varint,2,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	FolderId   int32     `protobuf:"varint,3,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	Rating     int32     `protobuf:"varint,4,opt,name=rating,proto3" json:"rating,omitempty"`
	Formats    []*Format `protobuf:"bytes,5,rep,name=formats,proto3" json:"formats,omitempty"`
	Labels     []*Label  `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *Release) Reset() {
	*x = Release{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discogs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Release) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Release) ProtoMessage() {}

func (x *Release) ProtoReflect() protoreflect.Message {
	mi := &file_discogs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Release.ProtoReflect.Descriptor instead.
func (*Release) Descriptor() ([]byte, []int) {
	return file_discogs_proto_rawDescGZIP(), []int{5}
}

func (x *Release) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Release) GetInstanceId() int64 {
	if x != nil {
		return x.InstanceId
	}
	return 0
}

func (x *Release) GetFolderId() int32 {
	if x != nil {
		return x.FolderId
	}
	return 0
}

func (x *Release) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Release) GetFormats() []*Format {
	if x != nil {
		return x.Formats
	}
	return nil
}

func (x *Release) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

type Label struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Catno string `protobuf:"bytes,3,opt,name=catno,proto3" json:"catno,omitempty"`
}

func (x *Label) Reset() {
	*x = Label{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discogs_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Label) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Label) ProtoMessage() {}

func (x *Label) ProtoReflect() protoreflect.Message {
	mi := &file_discogs_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Label.ProtoReflect.Descriptor instead.
func (*Label) Descriptor() ([]byte, []int) {
	return file_discogs_proto_rawDescGZIP(), []int{6}
}

func (x *Label) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Label) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Label) GetCatno() string {
	if x != nil {
		return x.Catno
	}
	return ""
}

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Currency string `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	Value    int32  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discogs_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_discogs_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_discogs_proto_rawDescGZIP(), []int{7}
}

func (x *Price) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Price) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type SaleItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SaleId    int64      `protobuf:"varint,1,opt,name=sale_id,json=saleId,proto3" json:"sale_id,omitempty"`
	Status    SaleStatus `protobuf:"varint,2,opt,name=status,proto3,enum=discogs.SaleStatus" json:"status,omitempty"`
	Price     *Price     `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	ReleaseId int64      `protobuf:"varint,4,opt,name=release_id,json=releaseId,proto3" json:"release_id,omitempty"`
}

func (x *SaleItem) Reset() {
	*x = SaleItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discogs_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleItem) ProtoMessage() {}

func (x *SaleItem) ProtoReflect() protoreflect.Message {
	mi := &file_discogs_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleItem.ProtoReflect.Descriptor instead.
func (*SaleItem) Descriptor() ([]byte, []int) {
	return file_discogs_proto_rawDescGZIP(), []int{8}
}

func (x *SaleItem) GetSaleId() int64 {
	if x != nil {
		return x.SaleId
	}
	return 0
}

func (x *SaleItem) GetStatus() SaleStatus {
	if x != nil {
		return x.Status
	}
	return SaleStatus_UNKNOWN
}

func (x *SaleItem) GetPrice() *Price {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *SaleItem) GetReleaseId() int64 {
	if x != nil {
		return x.ReleaseId
	}
	return 0
}

type Artist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Artist) Reset() {
	*x = Artist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discogs_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Artist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Artist) ProtoMessage() {}

func (x *Artist) ProtoReflect() protoreflect.Message {
	mi := &file_discogs_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Artist.ProtoReflect.Descriptor instead.
func (*Artist) Descriptor() ([]byte, []int) {
	return file_discogs_proto_rawDescGZIP(), []int{9}
}

func (x *Artist) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Want struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string    `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Artists []*Artist `protobuf:"bytes,3,rep,name=artists,proto3" json:"artists,omitempty"`
}

func (x *Want) Reset() {
	*x = Want{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discogs_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Want) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Want) ProtoMessage() {}

func (x *Want) ProtoReflect() protoreflect.Message {
	mi := &file_discogs_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Want.ProtoReflect.Descriptor instead.
func (*Want) Descriptor() ([]byte, []int) {
	return file_discogs_proto_rawDescGZIP(), []int{10}
}

func (x *Want) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Want) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Want) GetArtists() []*Artist {
	if x != nil {
		return x.Artists
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discogs_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_discogs_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_discogs_proto_rawDescGZIP(), []int{11}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_discogs_proto protoreflect.FileDescriptor

var file_discogs_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x67, 0x73, 0x22, 0x2b, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x06, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x22, 0x5c, 0x0a, 0x06, 0x46,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x8a, 0x01, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x67, 0x73, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x67, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0xc2, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x29, 0x0a, 0x07, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x67, 0x73, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x07, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x67, 0x73, 0x2e, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0x41, 0x0a, 0x05, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x74, 0x6e,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x74, 0x6e, 0x6f, 0x22, 0x39,
	0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x08, 0x53, 0x61,
	0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x13, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x67, 0x73, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x67, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49,
	0x64, 0x22, 0x1c, 0x0a, 0x06, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x57, 0x0a, 0x04, 0x57, 0x61, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x29, 0x0a,
	0x07, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x67, 0x73, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x52,
	0x07, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x22, 0x2f, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x40, 0x0a, 0x0a, 0x53, 0x61, 0x6c,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x4f, 0x52, 0x5f, 0x53, 0x41, 0x4c, 0x45,
	0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x4f, 0x4c, 0x44, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09,
	0x56, 0x49, 0x4f, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x42, 0x27, 0x5a, 0x25, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72, 0x6f, 0x74, 0x68, 0x65,
	0x72, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x67, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_discogs_proto_rawDescOnce sync.Once
	file_discogs_proto_rawDescData = file_discogs_proto_rawDesc
)

func file_discogs_proto_rawDescGZIP() []byte {
	file_discogs_proto_rawDescOnce.Do(func() {
		file_discogs_proto_rawDescData = protoimpl.X.CompressGZIP(file_discogs_proto_rawDescData)
	})
	return file_discogs_proto_rawDescData
}

var file_discogs_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_discogs_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_discogs_proto_goTypes = []interface{}{
	(SaleStatus)(0),    // 0: discogs.SaleStatus
	(*Field)(nil),      // 1: discogs.Field
	(*Folder)(nil),     // 2: discogs.Folder
	(*Pagination)(nil), // 3: discogs.Pagination
	(*Format)(nil),     // 4: discogs.Format
	(*User)(nil),       // 5: discogs.User
	(*Release)(nil),    // 6: discogs.Release
	(*Label)(nil),      // 7: discogs.Label
	(*Price)(nil),      // 8: discogs.Price
	(*SaleItem)(nil),   // 9: discogs.SaleItem
	(*Artist)(nil),     // 10: discogs.Artist
	(*Want)(nil),       // 11: discogs.Want
	(*Order)(nil),      // 12: discogs.Order
}
var file_discogs_proto_depIdxs = []int32{
	4,  // 0: discogs.Release.formats:type_name -> discogs.Format
	7,  // 1: discogs.Release.labels:type_name -> discogs.Label
	0,  // 2: discogs.SaleItem.status:type_name -> discogs.SaleStatus
	8,  // 3: discogs.SaleItem.price:type_name -> discogs.Price
	10, // 4: discogs.Want.artists:type_name -> discogs.Artist
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_discogs_proto_init() }
func file_discogs_proto_init() {
	if File_discogs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_discogs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_discogs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Folder); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_discogs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_discogs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Format); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_discogs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_discogs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Release); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_discogs_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Label); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_discogs_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Price); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_discogs_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_discogs_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Artist); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_discogs_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Want); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_discogs_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_discogs_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_discogs_proto_goTypes,
		DependencyIndexes: file_discogs_proto_depIdxs,
		EnumInfos:         file_discogs_proto_enumTypes,
		MessageInfos:      file_discogs_proto_msgTypes,
	}.Build()
	File_discogs_proto = out.File
	file_discogs_proto_rawDesc = nil
	file_discogs_proto_goTypes = nil
	file_discogs_proto_depIdxs = nil
}
