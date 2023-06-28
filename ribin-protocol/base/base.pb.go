// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.0
// source: base/base.proto

package base

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

// Request
type MsgType int32

const (
	MsgType_E_MSGTYPE_INVALID MsgType = 0
	MsgType_E_MSGTYPE_CHAT    MsgType = 1
)

// Enum value maps for MsgType.
var (
	MsgType_name = map[int32]string{
		0: "E_MSGTYPE_INVALID",
		1: "E_MSGTYPE_CHAT",
	}
	MsgType_value = map[string]int32{
		"E_MSGTYPE_INVALID": 0,
		"E_MSGTYPE_CHAT":    1,
	}
)

func (x MsgType) Enum() *MsgType {
	p := new(MsgType)
	*p = x
	return p
}

func (x MsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_base_base_proto_enumTypes[0].Descriptor()
}

func (MsgType) Type() protoreflect.EnumType {
	return &file_base_base_proto_enumTypes[0]
}

func (x MsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgType.Descriptor instead.
func (MsgType) EnumDescriptor() ([]byte, []int) {
	return file_base_base_proto_rawDescGZIP(), []int{0}
}

type Client2ServerReqCmd int32

const (
	Client2ServerReqCmd_E_CMD_INVALID    Client2ServerReqCmd = 0
	Client2ServerReqCmd_E_CMD_HEART_BEAT Client2ServerReqCmd = 10
)

// Enum value maps for Client2ServerReqCmd.
var (
	Client2ServerReqCmd_name = map[int32]string{
		0:  "E_CMD_INVALID",
		10: "E_CMD_HEART_BEAT",
	}
	Client2ServerReqCmd_value = map[string]int32{
		"E_CMD_INVALID":    0,
		"E_CMD_HEART_BEAT": 10,
	}
)

func (x Client2ServerReqCmd) Enum() *Client2ServerReqCmd {
	p := new(Client2ServerReqCmd)
	*p = x
	return p
}

func (x Client2ServerReqCmd) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Client2ServerReqCmd) Descriptor() protoreflect.EnumDescriptor {
	return file_base_base_proto_enumTypes[1].Descriptor()
}

func (Client2ServerReqCmd) Type() protoreflect.EnumType {
	return &file_base_base_proto_enumTypes[1]
}

func (x Client2ServerReqCmd) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Client2ServerReqCmd.Descriptor instead.
func (Client2ServerReqCmd) EnumDescriptor() ([]byte, []int) {
	return file_base_base_proto_rawDescGZIP(), []int{1}
}

type Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seq  string `protobuf:"bytes,1,opt,name=seq,proto3" json:"seq,omitempty"`
	Body []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Test) Reset() {
	*x = Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_base_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_base_base_proto_rawDescGZIP(), []int{0}
}

func (x *Test) GetSeq() string {
	if x != nil {
		return x.Seq
	}
	return ""
}

func (x *Test) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type ReqBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *ReqBody) Reset() {
	*x = ReqBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqBody) ProtoMessage() {}

func (x *ReqBody) ProtoReflect() protoreflect.Message {
	mi := &file_base_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqBody.ProtoReflect.Descriptor instead.
func (*ReqBody) Descriptor() ([]byte, []int) {
	return file_base_base_proto_rawDescGZIP(), []int{1}
}

func (x *ReqBody) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type Client2ServerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd Client2ServerReqCmd `protobuf:"varint,1,opt,name=cmd,proto3,enum=Client2ServerReqCmd" json:"cmd,omitempty"`
	Seq string              `protobuf:"bytes,2,opt,name=seq,proto3" json:"seq,omitempty"`
}

func (x *Client2ServerReq) Reset() {
	*x = Client2ServerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client2ServerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client2ServerReq) ProtoMessage() {}

func (x *Client2ServerReq) ProtoReflect() protoreflect.Message {
	mi := &file_base_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client2ServerReq.ProtoReflect.Descriptor instead.
func (*Client2ServerReq) Descriptor() ([]byte, []int) {
	return file_base_base_proto_rawDescGZIP(), []int{2}
}

func (x *Client2ServerReq) GetCmd() Client2ServerReqCmd {
	if x != nil {
		return x.Cmd
	}
	return Client2ServerReqCmd_E_CMD_INVALID
}

func (x *Client2ServerReq) GetSeq() string {
	if x != nil {
		return x.Seq
	}
	return ""
}

type Server2ClientRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seq  string              `protobuf:"bytes,1,opt,name=seq,proto3" json:"seq,omitempty"`
	Cmd  Client2ServerReqCmd `protobuf:"varint,2,opt,name=cmd,proto3,enum=Client2ServerReqCmd" json:"cmd,omitempty"`
	Code int32               `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string              `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Server2ClientRsp) Reset() {
	*x = Server2ClientRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_base_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server2ClientRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server2ClientRsp) ProtoMessage() {}

func (x *Server2ClientRsp) ProtoReflect() protoreflect.Message {
	mi := &file_base_base_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server2ClientRsp.ProtoReflect.Descriptor instead.
func (*Server2ClientRsp) Descriptor() ([]byte, []int) {
	return file_base_base_proto_rawDescGZIP(), []int{3}
}

func (x *Server2ClientRsp) GetSeq() string {
	if x != nil {
		return x.Seq
	}
	return ""
}

func (x *Server2ClientRsp) GetCmd() Client2ServerReqCmd {
	if x != nil {
		return x.Cmd
	}
	return Client2ServerReqCmd_E_CMD_INVALID
}

func (x *Server2ClientRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Server2ClientRsp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_base_base_proto protoreflect.FileDescriptor

var file_base_base_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x2c, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22,
	0x1b, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x4c, 0x0a, 0x10,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x32, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x12, 0x26, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x32, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x43, 0x6d, 0x64, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x71, 0x22, 0x72, 0x0a, 0x10, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x32, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x71,
	0x12, 0x26, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x32, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x43, 0x6d, 0x64, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x2a, 0x34,
	0x0a, 0x07, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x5f, 0x4d,
	0x53, 0x47, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00,
	0x12, 0x12, 0x0a, 0x0e, 0x45, 0x5f, 0x4d, 0x53, 0x47, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x48,
	0x41, 0x54, 0x10, 0x01, 0x2a, 0x3e, 0x0a, 0x13, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x32, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x43, 0x6d, 0x64, 0x12, 0x11, 0x0a, 0x0d, 0x45,
	0x5f, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14,
	0x0a, 0x10, 0x45, 0x5f, 0x43, 0x4d, 0x44, 0x5f, 0x48, 0x45, 0x41, 0x52, 0x54, 0x5f, 0x42, 0x45,
	0x41, 0x54, 0x10, 0x0a, 0x42, 0x15, 0x5a, 0x13, 0x72, 0x69, 0x62, 0x69, 0x6e, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_base_base_proto_rawDescOnce sync.Once
	file_base_base_proto_rawDescData = file_base_base_proto_rawDesc
)

func file_base_base_proto_rawDescGZIP() []byte {
	file_base_base_proto_rawDescOnce.Do(func() {
		file_base_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_base_proto_rawDescData)
	})
	return file_base_base_proto_rawDescData
}

var file_base_base_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_base_base_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_base_base_proto_goTypes = []interface{}{
	(MsgType)(0),             // 0: MsgType
	(Client2ServerReqCmd)(0), // 1: Client2ServerReqCmd
	(*Test)(nil),             // 2: Test
	(*ReqBody)(nil),          // 3: ReqBody
	(*Client2ServerReq)(nil), // 4: Client2ServerReq
	(*Server2ClientRsp)(nil), // 5: Server2ClientRsp
}
var file_base_base_proto_depIdxs = []int32{
	1, // 0: Client2ServerReq.cmd:type_name -> Client2ServerReqCmd
	1, // 1: Server2ClientRsp.cmd:type_name -> Client2ServerReqCmd
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_base_base_proto_init() }
func file_base_base_proto_init() {
	if File_base_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_base_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test); i {
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
		file_base_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqBody); i {
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
		file_base_base_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client2ServerReq); i {
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
		file_base_base_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server2ClientRsp); i {
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
			RawDescriptor: file_base_base_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_base_proto_goTypes,
		DependencyIndexes: file_base_base_proto_depIdxs,
		EnumInfos:         file_base_base_proto_enumTypes,
		MessageInfos:      file_base_base_proto_msgTypes,
	}.Build()
	File_base_base_proto = out.File
	file_base_base_proto_rawDesc = nil
	file_base_base_proto_goTypes = nil
	file_base_base_proto_depIdxs = nil
}
