// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: buzztracker.proto

package protos

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type GetBeerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBeerRequest) Reset() {
	*x = GetBeerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buzztracker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBeerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBeerRequest) ProtoMessage() {}

func (x *GetBeerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buzztracker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBeerRequest.ProtoReflect.Descriptor instead.
func (*GetBeerRequest) Descriptor() ([]byte, []int) {
	return file_buzztracker_proto_rawDescGZIP(), []int{0}
}

func (x *GetBeerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetBeerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Beer *Beer `protobuf:"bytes,1,opt,name=beer,proto3" json:"beer,omitempty"`
}

func (x *GetBeerResponse) Reset() {
	*x = GetBeerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buzztracker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBeerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBeerResponse) ProtoMessage() {}

func (x *GetBeerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buzztracker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBeerResponse.ProtoReflect.Descriptor instead.
func (*GetBeerResponse) Descriptor() ([]byte, []int) {
	return file_buzztracker_proto_rawDescGZIP(), []int{1}
}

func (x *GetBeerResponse) GetBeer() *Beer {
	if x != nil {
		return x.Beer
	}
	return nil
}

type Beer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Brewery     string               `protobuf:"bytes,3,opt,name=brewery,proto3" json:"brewery,omitempty"`
	BeerType    string               `protobuf:"bytes,4,opt,name=beerType,proto3" json:"beerType,omitempty"`
	LastUpdated *timestamp.Timestamp `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

func (x *Beer) Reset() {
	*x = Beer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buzztracker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Beer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Beer) ProtoMessage() {}

func (x *Beer) ProtoReflect() protoreflect.Message {
	mi := &file_buzztracker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Beer.ProtoReflect.Descriptor instead.
func (*Beer) Descriptor() ([]byte, []int) {
	return file_buzztracker_proto_rawDescGZIP(), []int{2}
}

func (x *Beer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Beer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Beer) GetBrewery() string {
	if x != nil {
		return x.Brewery
	}
	return ""
}

func (x *Beer) GetBeerType() string {
	if x != nil {
		return x.BeerType
	}
	return ""
}

func (x *Beer) GetLastUpdated() *timestamp.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

var File_buzztracker_proto protoreflect.FileDescriptor

var file_buzztracker_proto_rawDesc = []byte{
	0x0a, 0x11, 0x62, 0x75, 0x7a, 0x7a, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x62, 0x75, 0x7a, 0x7a, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x42, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x42, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x42, 0x65, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2f, 0x0a, 0x04, 0x62, 0x65, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x62, 0x75, 0x7a, 0x7a, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x42, 0x65, 0x65, 0x72, 0x52, 0x04, 0x62, 0x65, 0x65,
	0x72, 0x22, 0x9f, 0x01, 0x0a, 0x04, 0x42, 0x65, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x72, 0x65, 0x77, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x62, 0x72, 0x65, 0x77, 0x65, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x65, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x65, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x32, 0x6a, 0x0a, 0x0e, 0x42, 0x75, 0x7a, 0x7a, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x65, 0x72, 0x41, 0x50, 0x49, 0x12, 0x58, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x65, 0x65, 0x72,
	0x12, 0x25, 0x2e, 0x62, 0x75, 0x7a, 0x7a, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x65, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x62, 0x75, 0x7a, 0x7a, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x65, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x51, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61,
	0x6d, 0x6d, 0x69, 0x74, 0x62, 0x69, 0x6c, 0x6c, 0x79, 0x30, 0x6e, 0x65, 0x2f, 0x62, 0x75, 0x7a,
	0x7a, 0x2d, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x7a, 0x7a, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x61, 0x70, 0x69, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buzztracker_proto_rawDescOnce sync.Once
	file_buzztracker_proto_rawDescData = file_buzztracker_proto_rawDesc
)

func file_buzztracker_proto_rawDescGZIP() []byte {
	file_buzztracker_proto_rawDescOnce.Do(func() {
		file_buzztracker_proto_rawDescData = protoimpl.X.CompressGZIP(file_buzztracker_proto_rawDescData)
	})
	return file_buzztracker_proto_rawDescData
}

var file_buzztracker_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_buzztracker_proto_goTypes = []interface{}{
	(*GetBeerRequest)(nil),      // 0: buzztracker.api.alpha.GetBeerRequest
	(*GetBeerResponse)(nil),     // 1: buzztracker.api.alpha.GetBeerResponse
	(*Beer)(nil),                // 2: buzztracker.api.alpha.Beer
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_buzztracker_proto_depIdxs = []int32{
	2, // 0: buzztracker.api.alpha.GetBeerResponse.beer:type_name -> buzztracker.api.alpha.Beer
	3, // 1: buzztracker.api.alpha.Beer.last_updated:type_name -> google.protobuf.Timestamp
	0, // 2: buzztracker.api.alpha.BuzzTrackerAPI.GetBeer:input_type -> buzztracker.api.alpha.GetBeerRequest
	1, // 3: buzztracker.api.alpha.BuzzTrackerAPI.GetBeer:output_type -> buzztracker.api.alpha.GetBeerResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_buzztracker_proto_init() }
func file_buzztracker_proto_init() {
	if File_buzztracker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buzztracker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBeerRequest); i {
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
		file_buzztracker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBeerResponse); i {
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
		file_buzztracker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Beer); i {
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
			RawDescriptor: file_buzztracker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_buzztracker_proto_goTypes,
		DependencyIndexes: file_buzztracker_proto_depIdxs,
		MessageInfos:      file_buzztracker_proto_msgTypes,
	}.Build()
	File_buzztracker_proto = out.File
	file_buzztracker_proto_rawDesc = nil
	file_buzztracker_proto_goTypes = nil
	file_buzztracker_proto_depIdxs = nil
}
