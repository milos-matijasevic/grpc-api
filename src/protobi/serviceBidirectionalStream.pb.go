// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: serviceBidirectionalStream.proto

package proto

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartDate  string `protobuf:"bytes,1,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate    string `protobuf:"bytes,2,opt,name=endDate,proto3" json:"endDate,omitempty"`
	EventName  string `protobuf:"bytes,3,opt,name=eventName,proto3" json:"eventName,omitempty"`
	Advertiser int32  `protobuf:"varint,4,opt,name=advertiser,proto3" json:"advertiser,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serviceBidirectionalStream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_serviceBidirectionalStream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_serviceBidirectionalStream_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *Request) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *Request) GetEventName() string {
	if x != nil {
		return x.EventName
	}
	return ""
}

func (x *Request) GetAdvertiser() int32 {
	if x != nil {
		return x.Advertiser
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Publisher int32  `protobuf:"varint,1,opt,name=publisher,proto3" json:"publisher,omitempty"`
	Revenue   int64  `protobuf:"varint,2,opt,name=revenue,proto3" json:"revenue,omitempty"`
	Device    string `protobuf:"bytes,3,opt,name=device,proto3" json:"device,omitempty"`
	EventNum  int32  `protobuf:"varint,4,opt,name=eventNum,proto3" json:"eventNum,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serviceBidirectionalStream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_serviceBidirectionalStream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_serviceBidirectionalStream_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetPublisher() int32 {
	if x != nil {
		return x.Publisher
	}
	return 0
}

func (x *Response) GetRevenue() int64 {
	if x != nil {
		return x.Revenue
	}
	return 0
}

func (x *Response) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *Response) GetEventNum() int32 {
	if x != nil {
		return x.EventNum
	}
	return 0
}

var File_serviceBidirectionalStream_proto protoreflect.FileDescriptor

var file_serviceBidirectionalStream_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x07, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x64,
	0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x72, 0x22, 0x76, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e,
	0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e,
	0x75, 0x6d, 0x32, 0x8d, 0x01, 0x0a, 0x08, 0x46, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3f, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x46, 0x6f, 0x72, 0x45, 0x61,
	0x63, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01,
	0x12, 0x40, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x46, 0x6f, 0x72, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_serviceBidirectionalStream_proto_rawDescOnce sync.Once
	file_serviceBidirectionalStream_proto_rawDescData = file_serviceBidirectionalStream_proto_rawDesc
)

func file_serviceBidirectionalStream_proto_rawDescGZIP() []byte {
	file_serviceBidirectionalStream_proto_rawDescOnce.Do(func() {
		file_serviceBidirectionalStream_proto_rawDescData = protoimpl.X.CompressGZIP(file_serviceBidirectionalStream_proto_rawDescData)
	})
	return file_serviceBidirectionalStream_proto_rawDescData
}

var file_serviceBidirectionalStream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_serviceBidirectionalStream_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: proto.Request
	(*Response)(nil), // 1: proto.Response
}
var file_serviceBidirectionalStream_proto_depIdxs = []int32{
	0, // 0: proto.FService.GetMoneyForEachPublisher:input_type -> proto.Request
	0, // 1: proto.FService.GetNumEventsForAdvertiser:input_type -> proto.Request
	1, // 2: proto.FService.GetMoneyForEachPublisher:output_type -> proto.Response
	1, // 3: proto.FService.GetNumEventsForAdvertiser:output_type -> proto.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_serviceBidirectionalStream_proto_init() }
func file_serviceBidirectionalStream_proto_init() {
	if File_serviceBidirectionalStream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_serviceBidirectionalStream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_serviceBidirectionalStream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_serviceBidirectionalStream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_serviceBidirectionalStream_proto_goTypes,
		DependencyIndexes: file_serviceBidirectionalStream_proto_depIdxs,
		MessageInfos:      file_serviceBidirectionalStream_proto_msgTypes,
	}.Build()
	File_serviceBidirectionalStream_proto = out.File
	file_serviceBidirectionalStream_proto_rawDesc = nil
	file_serviceBidirectionalStream_proto_goTypes = nil
	file_serviceBidirectionalStream_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FServiceClient is the client API for FService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FServiceClient interface {
	GetMoneyForEachPublisher(ctx context.Context, opts ...grpc.CallOption) (FService_GetMoneyForEachPublisherClient, error)
	GetNumEventsForAdvertiser(ctx context.Context, opts ...grpc.CallOption) (FService_GetNumEventsForAdvertiserClient, error)
}

type fServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFServiceClient(cc grpc.ClientConnInterface) FServiceClient {
	return &fServiceClient{cc}
}

func (c *fServiceClient) GetMoneyForEachPublisher(ctx context.Context, opts ...grpc.CallOption) (FService_GetMoneyForEachPublisherClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FService_serviceDesc.Streams[0], "/proto.FService/GetMoneyForEachPublisher", opts...)
	if err != nil {
		return nil, err
	}
	x := &fServiceGetMoneyForEachPublisherClient{stream}
	return x, nil
}

type FService_GetMoneyForEachPublisherClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type fServiceGetMoneyForEachPublisherClient struct {
	grpc.ClientStream
}

func (x *fServiceGetMoneyForEachPublisherClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fServiceGetMoneyForEachPublisherClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fServiceClient) GetNumEventsForAdvertiser(ctx context.Context, opts ...grpc.CallOption) (FService_GetNumEventsForAdvertiserClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FService_serviceDesc.Streams[1], "/proto.FService/GetNumEventsForAdvertiser", opts...)
	if err != nil {
		return nil, err
	}
	x := &fServiceGetNumEventsForAdvertiserClient{stream}
	return x, nil
}

type FService_GetNumEventsForAdvertiserClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type fServiceGetNumEventsForAdvertiserClient struct {
	grpc.ClientStream
}

func (x *fServiceGetNumEventsForAdvertiserClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fServiceGetNumEventsForAdvertiserClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FServiceServer is the server API for FService service.
type FServiceServer interface {
	GetMoneyForEachPublisher(FService_GetMoneyForEachPublisherServer) error
	GetNumEventsForAdvertiser(FService_GetNumEventsForAdvertiserServer) error
}

// UnimplementedFServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFServiceServer struct {
}

func (*UnimplementedFServiceServer) GetMoneyForEachPublisher(FService_GetMoneyForEachPublisherServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMoneyForEachPublisher not implemented")
}
func (*UnimplementedFServiceServer) GetNumEventsForAdvertiser(FService_GetNumEventsForAdvertiserServer) error {
	return status.Errorf(codes.Unimplemented, "method GetNumEventsForAdvertiser not implemented")
}

func RegisterFServiceServer(s *grpc.Server, srv FServiceServer) {
	s.RegisterService(&_FService_serviceDesc, srv)
}

func _FService_GetMoneyForEachPublisher_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FServiceServer).GetMoneyForEachPublisher(&fServiceGetMoneyForEachPublisherServer{stream})
}

type FService_GetMoneyForEachPublisherServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type fServiceGetMoneyForEachPublisherServer struct {
	grpc.ServerStream
}

func (x *fServiceGetMoneyForEachPublisherServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fServiceGetMoneyForEachPublisherServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FService_GetNumEventsForAdvertiser_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FServiceServer).GetNumEventsForAdvertiser(&fServiceGetNumEventsForAdvertiserServer{stream})
}

type FService_GetNumEventsForAdvertiserServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type fServiceGetNumEventsForAdvertiserServer struct {
	grpc.ServerStream
}

func (x *fServiceGetNumEventsForAdvertiserServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fServiceGetNumEventsForAdvertiserServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _FService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.FService",
	HandlerType: (*FServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetMoneyForEachPublisher",
			Handler:       _FService_GetMoneyForEachPublisher_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetNumEventsForAdvertiser",
			Handler:       _FService_GetNumEventsForAdvertiser_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "serviceBidirectionalStream.proto",
}