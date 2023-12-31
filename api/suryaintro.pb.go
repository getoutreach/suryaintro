// Copyright 2023 Outreach Corporation. All Rights Reserved.
// Please modify this to match the interface specified in suryaintro.go

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: suryaintro.proto

package api

import (
	context "context"
	reflect "reflect"
	sync "sync"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Define your grpc service structures here
// PingRequest is the request for ping
type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_suryaintro_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_suryaintro_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_suryaintro_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// PingResponse is the response for echo.
type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_suryaintro_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_suryaintro_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_suryaintro_proto_rawDescGZIP(), []int{1}
}

func (x *PingResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// PongRequest is the request for ping
type PongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PongRequest) Reset() {
	*x = PongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_suryaintro_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PongRequest) ProtoMessage() {}

func (x *PongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_suryaintro_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PongRequest.ProtoReflect.Descriptor instead.
func (*PongRequest) Descriptor() ([]byte, []int) {
	return file_suryaintro_proto_rawDescGZIP(), []int{2}
}

func (x *PongRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// PongResponse is the response for echo.
type PongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PongResponse) Reset() {
	*x = PongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_suryaintro_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PongResponse) ProtoMessage() {}

func (x *PongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_suryaintro_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PongResponse.ProtoReflect.Descriptor instead.
func (*PongResponse) Descriptor() ([]byte, []int) {
	return file_suryaintro_proto_rawDescGZIP(), []int{3}
}

func (x *PongResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// JokeRequest is the request for ping
type JokeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *JokeRequest) Reset() {
	*x = JokeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_suryaintro_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JokeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JokeRequest) ProtoMessage() {}

func (x *JokeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_suryaintro_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JokeRequest.ProtoReflect.Descriptor instead.
func (*JokeRequest) Descriptor() ([]byte, []int) {
	return file_suryaintro_proto_rawDescGZIP(), []int{4}
}

// JokeResponse is the response for echo.
type JokeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Joke string `protobuf:"bytes,1,opt,name=joke,proto3" json:"joke,omitempty"`
}

func (x *JokeResponse) Reset() {
	*x = JokeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_suryaintro_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JokeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JokeResponse) ProtoMessage() {}

func (x *JokeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_suryaintro_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JokeResponse.ProtoReflect.Descriptor instead.
func (*JokeResponse) Descriptor() ([]byte, []int) {
	return file_suryaintro_proto_rawDescGZIP(), []int{5}
}

func (x *JokeResponse) GetJoke() string {
	if x != nil {
		return x.Joke
	}
	return ""
}

var File_suryaintro_proto protoreflect.FileDescriptor

var file_suryaintro_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x75, 0x72, 0x79, 0x61, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x73, 0x75, 0x72, 0x79, 0x61, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x22, 0x27, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x0c, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x27, 0x0a, 0x0b, 0x50, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x28,
	0x0a, 0x0c, 0x50, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x4a, 0x6f, 0x6b, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x22, 0x0a, 0x0c, 0x4a, 0x6f, 0x6b, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6a, 0x6f, 0x6b, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6a, 0x6f, 0x6b, 0x65, 0x32, 0xdb, 0x01, 0x0a, 0x0a,
	0x53, 0x75, 0x72, 0x79, 0x61, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x12, 0x43, 0x0a, 0x04, 0x50, 0x69,
	0x6e, 0x67, 0x12, 0x1b, 0x2e, 0x73, 0x75, 0x72, 0x79, 0x61, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x73, 0x75, 0x72, 0x79, 0x61, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x43, 0x0a, 0x04, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x1b, 0x2e, 0x73, 0x75, 0x72, 0x79, 0x61, 0x69,
	0x6e, 0x74, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x75, 0x72, 0x79, 0x61, 0x69, 0x6e, 0x74, 0x72,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x04, 0x4a, 0x6f, 0x6b, 0x65, 0x12, 0x1b, 0x2e, 0x73,
	0x75, 0x72, 0x79, 0x61, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4a, 0x6f,
	0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x75, 0x72, 0x79,
	0x61, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4a, 0x6f, 0x6b, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3a, 0x5a, 0x25, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x74, 0x6f, 0x75, 0x74, 0x72, 0x65,
	0x61, 0x63, 0x68, 0x2f, 0x73, 0x75, 0x72, 0x79, 0x61, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0xea, 0x02, 0x10, 0x53, 0x75, 0x72, 0x79, 0x61, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_suryaintro_proto_rawDescOnce sync.Once
	file_suryaintro_proto_rawDescData = file_suryaintro_proto_rawDesc
)

func file_suryaintro_proto_rawDescGZIP() []byte {
	file_suryaintro_proto_rawDescOnce.Do(func() {
		file_suryaintro_proto_rawDescData = protoimpl.X.CompressGZIP(file_suryaintro_proto_rawDescData)
	})
	return file_suryaintro_proto_rawDescData
}

var file_suryaintro_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_suryaintro_proto_goTypes = []interface{}{
	(*PingRequest)(nil),  // 0: suryaintro.api.PingRequest
	(*PingResponse)(nil), // 1: suryaintro.api.PingResponse
	(*PongRequest)(nil),  // 2: suryaintro.api.PongRequest
	(*PongResponse)(nil), // 3: suryaintro.api.PongResponse
	(*JokeRequest)(nil),  // 4: suryaintro.api.JokeRequest
	(*JokeResponse)(nil), // 5: suryaintro.api.JokeResponse
}
var file_suryaintro_proto_depIdxs = []int32{
	0, // 0: suryaintro.api.Suryaintro.Ping:input_type -> suryaintro.api.PingRequest
	2, // 1: suryaintro.api.Suryaintro.Pong:input_type -> suryaintro.api.PongRequest
	4, // 2: suryaintro.api.Suryaintro.Joke:input_type -> suryaintro.api.JokeRequest
	1, // 3: suryaintro.api.Suryaintro.Ping:output_type -> suryaintro.api.PingResponse
	3, // 4: suryaintro.api.Suryaintro.Pong:output_type -> suryaintro.api.PongResponse
	5, // 5: suryaintro.api.Suryaintro.Joke:output_type -> suryaintro.api.JokeResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_suryaintro_proto_init() }
func file_suryaintro_proto_init() {
	if File_suryaintro_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_suryaintro_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_suryaintro_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_suryaintro_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PongRequest); i {
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
		file_suryaintro_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PongResponse); i {
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
		file_suryaintro_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JokeRequest); i {
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
		file_suryaintro_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JokeResponse); i {
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
			RawDescriptor: file_suryaintro_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_suryaintro_proto_goTypes,
		DependencyIndexes: file_suryaintro_proto_depIdxs,
		MessageInfos:      file_suryaintro_proto_msgTypes,
	}.Build()
	File_suryaintro_proto = out.File
	file_suryaintro_proto_rawDesc = nil
	file_suryaintro_proto_goTypes = nil
	file_suryaintro_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SuryaintroClient is the client API for Suryaintro service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SuryaintroClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	Pong(ctx context.Context, in *PongRequest, opts ...grpc.CallOption) (*PongResponse, error)
	Joke(ctx context.Context, in *JokeRequest, opts ...grpc.CallOption) (*JokeResponse, error)
}

type suryaintroClient struct {
	cc grpc.ClientConnInterface
}

func NewSuryaintroClient(cc grpc.ClientConnInterface) SuryaintroClient {
	return &suryaintroClient{cc}
}

func (c *suryaintroClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/suryaintro.api.Suryaintro/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *suryaintroClient) Pong(ctx context.Context, in *PongRequest, opts ...grpc.CallOption) (*PongResponse, error) {
	out := new(PongResponse)
	err := c.cc.Invoke(ctx, "/suryaintro.api.Suryaintro/Pong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *suryaintroClient) Joke(ctx context.Context, in *JokeRequest, opts ...grpc.CallOption) (*JokeResponse, error) {
	out := new(JokeResponse)
	err := c.cc.Invoke(ctx, "/suryaintro.api.Suryaintro/Joke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SuryaintroServer is the server API for Suryaintro service.
type SuryaintroServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	Pong(context.Context, *PongRequest) (*PongResponse, error)
	Joke(context.Context, *JokeRequest) (*JokeResponse, error)
}

// UnimplementedSuryaintroServer can be embedded to have forward compatible implementations.
type UnimplementedSuryaintroServer struct {
}

func (*UnimplementedSuryaintroServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedSuryaintroServer) Pong(context.Context, *PongRequest) (*PongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pong not implemented")
}
func (*UnimplementedSuryaintroServer) Joke(context.Context, *JokeRequest) (*JokeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Joke not implemented")
}

func RegisterSuryaintroServer(s *grpc.Server, srv SuryaintroServer) {
	s.RegisterService(&_Suryaintro_serviceDesc, srv)
}

func _Suryaintro_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuryaintroServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/suryaintro.api.Suryaintro/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuryaintroServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Suryaintro_Pong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuryaintroServer).Pong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/suryaintro.api.Suryaintro/Pong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuryaintroServer).Pong(ctx, req.(*PongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Suryaintro_Joke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JokeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuryaintroServer).Joke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/suryaintro.api.Suryaintro/Joke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuryaintroServer).Joke(ctx, req.(*JokeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Suryaintro_serviceDesc = grpc.ServiceDesc{
	ServiceName: "suryaintro.api.Suryaintro",
	HandlerType: (*SuryaintroServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Suryaintro_Ping_Handler,
		},
		{
			MethodName: "Pong",
			Handler:    _Suryaintro_Pong_Handler,
		},
		{
			MethodName: "Joke",
			Handler:    _Suryaintro_Joke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "suryaintro.proto",
}
