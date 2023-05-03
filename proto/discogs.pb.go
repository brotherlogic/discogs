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
		mi := &file_discogs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_discogs_proto_rawDescGZIP(), []int{0}
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
		mi := &file_discogs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Format) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Format) ProtoMessage() {}

func (x *Format) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Format.ProtoReflect.Descriptor instead.
func (*Format) Descriptor() ([]byte, []int) {
	return file_discogs_proto_rawDescGZIP(), []int{1}
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
		mi := &file_discogs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_discogs_proto_rawDescGZIP(), []int{2}
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
}

func (x *Release) Reset() {
	*x = Release{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discogs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Release) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Release) ProtoMessage() {}

func (x *Release) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Release.ProtoReflect.Descriptor instead.
func (*Release) Descriptor() ([]byte, []int) {
	return file_discogs_proto_rawDescGZIP(), []int{3}
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

var File_discogs_proto protoreflect.FileDescriptor

var file_discogs_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x67, 0x73, 0x22, 0x36, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73,
	0x22, 0x5c, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0c, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x8a,
	0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x67, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x67, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x9a, 0x01, 0x0a, 0x07,
	0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x29, 0x0a,
	0x07, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x67, 0x73, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52,
	0x07, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x73, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x6c, 0x6f,
	0x67, 0x69, 0x63, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x67, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_discogs_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_discogs_proto_goTypes = []interface{}{
	(*Pagination)(nil), // 0: discogs.Pagination
	(*Format)(nil),     // 1: discogs.Format
	(*User)(nil),       // 2: discogs.User
	(*Release)(nil),    // 3: discogs.Release
}
var file_discogs_proto_depIdxs = []int32{
	1, // 0: discogs.Release.formats:type_name -> discogs.Format
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_discogs_proto_init() }
func file_discogs_proto_init() {
	if File_discogs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_discogs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_discogs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_discogs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_discogs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_discogs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_discogs_proto_goTypes,
		DependencyIndexes: file_discogs_proto_depIdxs,
		MessageInfos:      file_discogs_proto_msgTypes,
	}.Build()
	File_discogs_proto = out.File
	file_discogs_proto_rawDesc = nil
	file_discogs_proto_goTypes = nil
	file_discogs_proto_depIdxs = nil
}
