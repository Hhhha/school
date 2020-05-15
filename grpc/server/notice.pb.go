// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.9.1
// source: notice.proto

package grpc

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type FilterInsertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notices []*Notice `protobuf:"bytes,1,rep,name=notices,proto3" json:"notices,omitempty"`
}

func (x *FilterInsertRequest) Reset() {
	*x = FilterInsertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterInsertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterInsertRequest) ProtoMessage() {}

func (x *FilterInsertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterInsertRequest.ProtoReflect.Descriptor instead.
func (*FilterInsertRequest) Descriptor() ([]byte, []int) {
	return file_notice_proto_rawDescGZIP(), []int{0}
}

func (x *FilterInsertRequest) GetNotices() []*Notice {
	if x != nil {
		return x.Notices
	}
	return nil
}

type FilterInsertReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *FilterInsertReply) Reset() {
	*x = FilterInsertReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterInsertReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterInsertReply) ProtoMessage() {}

func (x *FilterInsertReply) ProtoReflect() protoreflect.Message {
	mi := &file_notice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterInsertReply.ProtoReflect.Descriptor instead.
func (*FilterInsertReply) Descriptor() ([]byte, []int) {
	return file_notice_proto_rawDescGZIP(), []int{1}
}

func (x *FilterInsertReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type Notice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	CreateTime string `protobuf:"bytes,2,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	JumpUrl    string `protobuf:"bytes,3,opt,name=JumpUrl,proto3" json:"JumpUrl,omitempty"`
}

func (x *Notice) Reset() {
	*x = Notice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notice) ProtoMessage() {}

func (x *Notice) ProtoReflect() protoreflect.Message {
	mi := &file_notice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notice.ProtoReflect.Descriptor instead.
func (*Notice) Descriptor() ([]byte, []int) {
	return file_notice_proto_rawDescGZIP(), []int{2}
}

func (x *Notice) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Notice) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *Notice) GetJumpUrl() string {
	if x != nil {
		return x.JumpUrl
	}
	return ""
}

var File_notice_proto protoreflect.FileDescriptor

var file_notice_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x67, 0x72, 0x70, 0x63, 0x22, 0x3d, 0x0a, 0x13, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x49, 0x6e,
	0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x07, 0x6e,
	0x6f, 0x74, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52, 0x07, 0x6e, 0x6f, 0x74, 0x69,
	0x63, 0x65, 0x73, 0x22, 0x27, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73,
	0x65, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x58, 0x0a, 0x06,
	0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x4a, 0x75, 0x6d, 0x70, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4a,
	0x75, 0x6d, 0x70, 0x55, 0x72, 0x6c, 0x32, 0x55, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notice_proto_rawDescOnce sync.Once
	file_notice_proto_rawDescData = file_notice_proto_rawDesc
)

func file_notice_proto_rawDescGZIP() []byte {
	file_notice_proto_rawDescOnce.Do(func() {
		file_notice_proto_rawDescData = protoimpl.X.CompressGZIP(file_notice_proto_rawDescData)
	})
	return file_notice_proto_rawDescData
}

var file_notice_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_notice_proto_goTypes = []interface{}{
	(*FilterInsertRequest)(nil), // 0: grpc.FilterInsertRequest
	(*FilterInsertReply)(nil),   // 1: grpc.FilterInsertReply
	(*Notice)(nil),              // 2: grpc.Notice
}
var file_notice_proto_depIdxs = []int32{
	2, // 0: grpc.FilterInsertRequest.notices:type_name -> grpc.Notice
	0, // 1: grpc.NoticeService.FilterInsert:input_type -> grpc.FilterInsertRequest
	1, // 2: grpc.NoticeService.FilterInsert:output_type -> grpc.FilterInsertReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_notice_proto_init() }
func file_notice_proto_init() {
	if File_notice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterInsertRequest); i {
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
		file_notice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterInsertReply); i {
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
		file_notice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notice); i {
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
			RawDescriptor: file_notice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notice_proto_goTypes,
		DependencyIndexes: file_notice_proto_depIdxs,
		MessageInfos:      file_notice_proto_msgTypes,
	}.Build()
	File_notice_proto = out.File
	file_notice_proto_rawDesc = nil
	file_notice_proto_goTypes = nil
	file_notice_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NoticeServiceClient is the client API for NoticeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NoticeServiceClient interface {
	FilterInsert(ctx context.Context, in *FilterInsertRequest, opts ...grpc.CallOption) (*FilterInsertReply, error)
}

type noticeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNoticeServiceClient(cc grpc.ClientConnInterface) NoticeServiceClient {
	return &noticeServiceClient{cc}
}

func (c *noticeServiceClient) FilterInsert(ctx context.Context, in *FilterInsertRequest, opts ...grpc.CallOption) (*FilterInsertReply, error) {
	out := new(FilterInsertReply)
	err := c.cc.Invoke(ctx, "/grpc.NoticeService/FilterInsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NoticeServiceServer is the server API for NoticeService service.
type NoticeServiceServer interface {
	FilterInsert(context.Context, *FilterInsertRequest) (*FilterInsertReply, error)
}

// UnimplementedNoticeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNoticeServiceServer struct {
}

func (*UnimplementedNoticeServiceServer) FilterInsert(context.Context, *FilterInsertRequest) (*FilterInsertReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FilterInsert not implemented")
}

func RegisterNoticeServiceServer(s *grpc.Server, srv NoticeServiceServer) {
	s.RegisterService(&_NoticeService_serviceDesc, srv)
}

func _NoticeService_FilterInsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterInsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoticeServiceServer).FilterInsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.NoticeService/FilterInsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoticeServiceServer).FilterInsert(ctx, req.(*FilterInsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NoticeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.NoticeService",
	HandlerType: (*NoticeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FilterInsert",
			Handler:    _NoticeService_FilterInsert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notice.proto",
}