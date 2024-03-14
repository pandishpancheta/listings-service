// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: pkg/tokenization/pb/tokenization.proto

package tokenization

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

type TokenizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenId string `protobuf:"bytes,1,opt,name=tokenId,proto3" json:"tokenId,omitempty"`
	Chunk   []byte `protobuf:"bytes,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *TokenizationRequest) Reset() {
	*x = TokenizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_tokenization_pb_tokenization_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenizationRequest) ProtoMessage() {}

func (x *TokenizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_tokenization_pb_tokenization_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenizationRequest.ProtoReflect.Descriptor instead.
func (*TokenizationRequest) Descriptor() ([]byte, []int) {
	return file_pkg_tokenization_pb_tokenization_proto_rawDescGZIP(), []int{0}
}

func (x *TokenizationRequest) GetTokenId() string {
	if x != nil {
		return x.TokenId
	}
	return ""
}

func (x *TokenizationRequest) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type TokenizationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenURI string `protobuf:"bytes,1,opt,name=tokenURI,proto3" json:"tokenURI,omitempty"`
}

func (x *TokenizationResponse) Reset() {
	*x = TokenizationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_tokenization_pb_tokenization_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenizationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenizationResponse) ProtoMessage() {}

func (x *TokenizationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_tokenization_pb_tokenization_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenizationResponse.ProtoReflect.Descriptor instead.
func (*TokenizationResponse) Descriptor() ([]byte, []int) {
	return file_pkg_tokenization_pb_tokenization_proto_rawDescGZIP(), []int{1}
}

func (x *TokenizationResponse) GetTokenURI() string {
	if x != nil {
		return x.TokenURI
	}
	return ""
}

var File_pkg_tokenization_pb_tokenization_proto protoreflect.FileDescriptor

var file_pkg_tokenization_pb_tokenization_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x45, 0x0a, 0x13, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x32, 0x0a,
	0x14, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x55, 0x52,
	0x49, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x55, 0x52,
	0x49, 0x32, 0x6a, 0x0a, 0x13, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x08, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x69, 0x7a, 0x65, 0x12, 0x21, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x45, 0x5a,
	0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x6e, 0x64,
	0x69, 0x73, 0x68, 0x70, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x74, 0x61, 0x2f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_tokenization_pb_tokenization_proto_rawDescOnce sync.Once
	file_pkg_tokenization_pb_tokenization_proto_rawDescData = file_pkg_tokenization_pb_tokenization_proto_rawDesc
)

func file_pkg_tokenization_pb_tokenization_proto_rawDescGZIP() []byte {
	file_pkg_tokenization_pb_tokenization_proto_rawDescOnce.Do(func() {
		file_pkg_tokenization_pb_tokenization_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_tokenization_pb_tokenization_proto_rawDescData)
	})
	return file_pkg_tokenization_pb_tokenization_proto_rawDescData
}

var file_pkg_tokenization_pb_tokenization_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_tokenization_pb_tokenization_proto_goTypes = []interface{}{
	(*TokenizationRequest)(nil),  // 0: tokenization.TokenizationRequest
	(*TokenizationResponse)(nil), // 1: tokenization.TokenizationResponse
}
var file_pkg_tokenization_pb_tokenization_proto_depIdxs = []int32{
	0, // 0: tokenization.TokenizationService.Tokenize:input_type -> tokenization.TokenizationRequest
	1, // 1: tokenization.TokenizationService.Tokenize:output_type -> tokenization.TokenizationResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_tokenization_pb_tokenization_proto_init() }
func file_pkg_tokenization_pb_tokenization_proto_init() {
	if File_pkg_tokenization_pb_tokenization_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_tokenization_pb_tokenization_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenizationRequest); i {
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
		file_pkg_tokenization_pb_tokenization_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenizationResponse); i {
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
			RawDescriptor: file_pkg_tokenization_pb_tokenization_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_tokenization_pb_tokenization_proto_goTypes,
		DependencyIndexes: file_pkg_tokenization_pb_tokenization_proto_depIdxs,
		MessageInfos:      file_pkg_tokenization_pb_tokenization_proto_msgTypes,
	}.Build()
	File_pkg_tokenization_pb_tokenization_proto = out.File
	file_pkg_tokenization_pb_tokenization_proto_rawDesc = nil
	file_pkg_tokenization_pb_tokenization_proto_goTypes = nil
	file_pkg_tokenization_pb_tokenization_proto_depIdxs = nil
}
