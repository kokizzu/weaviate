//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package grpc

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassName  string            `protobuf:"bytes,1,opt,name=className,proto3" json:"className,omitempty"`
	Limit      uint32            `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	NearVector *NearVectorParams `protobuf:"bytes,3,opt,name=nearVector,proto3" json:"nearVector,omitempty"`
	NearObject *NearObjectParams `protobuf:"bytes,4,opt,name=nearObject,proto3" json:"nearObject,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_weaviate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_weaviate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_grpc_weaviate_proto_rawDescGZIP(), []int{0}
}

func (x *SearchRequest) GetClassName() string {
	if x != nil {
		return x.ClassName
	}
	return ""
}

func (x *SearchRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *SearchRequest) GetNearVector() *NearVectorParams {
	if x != nil {
		return x.NearVector
	}
	return nil
}

func (x *SearchRequest) GetNearObject() *NearObjectParams {
	if x != nil {
		return x.NearObject
	}
	return nil
}

type NearVectorParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vector []float32 `protobuf:"fixed32,1,rep,packed,name=vector,proto3" json:"vector,omitempty"`
}

func (x *NearVectorParams) Reset() {
	*x = NearVectorParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_weaviate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NearVectorParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NearVectorParams) ProtoMessage() {}

func (x *NearVectorParams) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_weaviate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NearVectorParams.ProtoReflect.Descriptor instead.
func (*NearVectorParams) Descriptor() ([]byte, []int) {
	return file_grpc_weaviate_proto_rawDescGZIP(), []int{1}
}

func (x *NearVectorParams) GetVector() []float32 {
	if x != nil {
		return x.Vector
	}
	return nil
}

type NearObjectParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NearObjectParams) Reset() {
	*x = NearObjectParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_weaviate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NearObjectParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NearObjectParams) ProtoMessage() {}

func (x *NearObjectParams) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_weaviate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NearObjectParams.ProtoReflect.Descriptor instead.
func (*NearObjectParams) Descriptor() ([]byte, []int) {
	return file_grpc_weaviate_proto_rawDescGZIP(), []int{2}
}

func (x *NearObjectParams) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SearchReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*SearchResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Took    float32         `protobuf:"fixed32,2,opt,name=took,proto3" json:"took,omitempty"`
}

func (x *SearchReply) Reset() {
	*x = SearchReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_weaviate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchReply) ProtoMessage() {}

func (x *SearchReply) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_weaviate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchReply.ProtoReflect.Descriptor instead.
func (*SearchReply) Descriptor() ([]byte, []int) {
	return file_grpc_weaviate_proto_rawDescGZIP(), []int{3}
}

func (x *SearchReply) GetResults() []*SearchResult {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *SearchReply) GetTook() float32 {
	if x != nil {
		return x.Took
	}
	return 0
}

type SearchResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SearchResult) Reset() {
	*x = SearchResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_weaviate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResult) ProtoMessage() {}

func (x *SearchResult) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_weaviate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResult.ProtoReflect.Descriptor instead.
func (*SearchResult) Descriptor() ([]byte, []int) {
	return file_grpc_weaviate_proto_rawDescGZIP(), []int{4}
}

func (x *SearchResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_grpc_weaviate_proto protoreflect.FileDescriptor

var file_grpc_weaviate_proto_rawDesc = []byte{
	0x0a, 0x13, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x67,
	0x72, 0x70, 0x63, 0x22, 0xc3, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x6e, 0x65, 0x61,
	0x72, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4e, 0x65, 0x61,
	0x72, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x0a, 0x6e,
	0x65, 0x61, 0x72, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x3e, 0x0a, 0x0a, 0x6e, 0x65, 0x61,
	0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4e, 0x65, 0x61,
	0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x0a, 0x6e,
	0x65, 0x61, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2a, 0x0a, 0x10, 0x4e, 0x65, 0x61,
	0x72, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x02, 0x52, 0x06, 0x76,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x22, 0x0a, 0x10, 0x4e, 0x65, 0x61, 0x72, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x57, 0x0a, 0x0b, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x34, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x77, 0x65, 0x61, 0x76,
	0x69, 0x61, 0x74, 0x65, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x6f, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x74, 0x6f,
	0x6f, 0x6b, 0x22, 0x1e, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x32, 0x4e, 0x0a, 0x08, 0x57, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x12, 0x42,
	0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1b, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69,
	0x61, 0x74, 0x65, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61,
	0x74, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_weaviate_proto_rawDescOnce sync.Once
	file_grpc_weaviate_proto_rawDescData = file_grpc_weaviate_proto_rawDesc
)

func file_grpc_weaviate_proto_rawDescGZIP() []byte {
	file_grpc_weaviate_proto_rawDescOnce.Do(func() {
		file_grpc_weaviate_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_weaviate_proto_rawDescData)
	})
	return file_grpc_weaviate_proto_rawDescData
}

var file_grpc_weaviate_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_grpc_weaviate_proto_goTypes = []interface{}{
	(*SearchRequest)(nil),    // 0: weaviategrpc.SearchRequest
	(*NearVectorParams)(nil), // 1: weaviategrpc.NearVectorParams
	(*NearObjectParams)(nil), // 2: weaviategrpc.NearObjectParams
	(*SearchReply)(nil),      // 3: weaviategrpc.SearchReply
	(*SearchResult)(nil),     // 4: weaviategrpc.SearchResult
}
var file_grpc_weaviate_proto_depIdxs = []int32{
	1, // 0: weaviategrpc.SearchRequest.nearVector:type_name -> weaviategrpc.NearVectorParams
	2, // 1: weaviategrpc.SearchRequest.nearObject:type_name -> weaviategrpc.NearObjectParams
	4, // 2: weaviategrpc.SearchReply.results:type_name -> weaviategrpc.SearchResult
	0, // 3: weaviategrpc.Weaviate.Search:input_type -> weaviategrpc.SearchRequest
	3, // 4: weaviategrpc.Weaviate.Search:output_type -> weaviategrpc.SearchReply
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_grpc_weaviate_proto_init() }
func file_grpc_weaviate_proto_init() {
	if File_grpc_weaviate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_weaviate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_grpc_weaviate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NearVectorParams); i {
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
		file_grpc_weaviate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NearObjectParams); i {
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
		file_grpc_weaviate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchReply); i {
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
		file_grpc_weaviate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResult); i {
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
			RawDescriptor: file_grpc_weaviate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_weaviate_proto_goTypes,
		DependencyIndexes: file_grpc_weaviate_proto_depIdxs,
		MessageInfos:      file_grpc_weaviate_proto_msgTypes,
	}.Build()
	File_grpc_weaviate_proto = out.File
	file_grpc_weaviate_proto_rawDesc = nil
	file_grpc_weaviate_proto_goTypes = nil
	file_grpc_weaviate_proto_depIdxs = nil
}
