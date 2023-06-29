// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.0
// source: server-data/server_data.proto

package server_data

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int32  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Seq string `protobuf:"bytes,2,opt,name=seq,proto3" json:"seq,omitempty"`
	Ts  int32  `protobuf:"varint,3,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_data_server_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_server_data_server_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_server_data_server_data_proto_rawDescGZIP(), []int{0}
}

func (x *Ping) GetUid() int32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *Ping) GetSeq() string {
	if x != nil {
		return x.Seq
	}
	return ""
}

func (x *Ping) GetTs() int32 {
	if x != nil {
		return x.Ts
	}
	return 0
}

type Pong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int32  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Seq string `protobuf:"bytes,2,opt,name=seq,proto3" json:"seq,omitempty"`
	Ts  int32  `protobuf:"varint,3,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *Pong) Reset() {
	*x = Pong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_data_server_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pong) ProtoMessage() {}

func (x *Pong) ProtoReflect() protoreflect.Message {
	mi := &file_server_data_server_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pong.ProtoReflect.Descriptor instead.
func (*Pong) Descriptor() ([]byte, []int) {
	return file_server_data_server_data_proto_rawDescGZIP(), []int{1}
}

func (x *Pong) GetUid() int32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *Pong) GetSeq() string {
	if x != nil {
		return x.Seq
	}
	return ""
}

func (x *Pong) GetTs() int32 {
	if x != nil {
		return x.Ts
	}
	return 0
}

var File_server_data_server_data_proto protoreflect.FileDescriptor

var file_server_data_server_data_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x3a, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x74,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x74, 0x73, 0x22, 0x3a, 0x0a, 0x04, 0x50,
	0x6f, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x74, 0x73, 0x32, 0x29, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65,
	0x61, 0x74, 0x12, 0x05, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x1a, 0x05, 0x2e, 0x50, 0x6f, 0x6e, 0x67,
	0x22, 0x00, 0x42, 0x1c, 0x5a, 0x1a, 0x72, 0x69, 0x62, 0x69, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_data_server_data_proto_rawDescOnce sync.Once
	file_server_data_server_data_proto_rawDescData = file_server_data_server_data_proto_rawDesc
)

func file_server_data_server_data_proto_rawDescGZIP() []byte {
	file_server_data_server_data_proto_rawDescOnce.Do(func() {
		file_server_data_server_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_data_server_data_proto_rawDescData)
	})
	return file_server_data_server_data_proto_rawDescData
}

var file_server_data_server_data_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_server_data_server_data_proto_goTypes = []interface{}{
	(*Ping)(nil), // 0: Ping
	(*Pong)(nil), // 1: Pong
}
var file_server_data_server_data_proto_depIdxs = []int32{
	0, // 0: ServerData.HeartBeat:input_type -> Ping
	1, // 1: ServerData.HeartBeat:output_type -> Pong
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_server_data_server_data_proto_init() }
func file_server_data_server_data_proto_init() {
	if File_server_data_server_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_data_server_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_server_data_server_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pong); i {
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
			RawDescriptor: file_server_data_server_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_data_server_data_proto_goTypes,
		DependencyIndexes: file_server_data_server_data_proto_depIdxs,
		MessageInfos:      file_server_data_server_data_proto_msgTypes,
	}.Build()
	File_server_data_server_data_proto = out.File
	file_server_data_server_data_proto_rawDesc = nil
	file_server_data_server_data_proto_goTypes = nil
	file_server_data_server_data_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServerDataClient is the client API for ServerData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerDataClient interface {
	HeartBeat(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error)
}

type serverDataClient struct {
	cc grpc.ClientConnInterface
}

func NewServerDataClient(cc grpc.ClientConnInterface) ServerDataClient {
	return &serverDataClient{cc}
}

func (c *serverDataClient) HeartBeat(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/ServerData/HeartBeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerDataServer is the server API for ServerData service.
type ServerDataServer interface {
	HeartBeat(context.Context, *Ping) (*Pong, error)
}

// UnimplementedServerDataServer can be embedded to have forward compatible implementations.
type UnimplementedServerDataServer struct {
}

func (*UnimplementedServerDataServer) HeartBeat(context.Context, *Ping) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartBeat not implemented")
}

func RegisterServerDataServer(s *grpc.Server, srv ServerDataServer) {
	s.RegisterService(&_ServerData_serviceDesc, srv)
}

func _ServerData_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerDataServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ServerData/HeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerDataServer).HeartBeat(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServerData_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ServerData",
	HandlerType: (*ServerDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HeartBeat",
			Handler:    _ServerData_HeartBeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server-data/server_data.proto",
}