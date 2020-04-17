// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/detail_placement_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for [DetailPlacementViewService.GetDetailPlacementView][google.ads.googleads.v2.services.DetailPlacementViewService.GetDetailPlacementView].
type GetDetailPlacementViewRequest struct {
	// Required. The resource name of the Detail Placement view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDetailPlacementViewRequest) Reset()         { *m = GetDetailPlacementViewRequest{} }
func (m *GetDetailPlacementViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetDetailPlacementViewRequest) ProtoMessage()    {}
func (*GetDetailPlacementViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_24eb6ba92ccb047b, []int{0}
}

func (m *GetDetailPlacementViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDetailPlacementViewRequest.Unmarshal(m, b)
}
func (m *GetDetailPlacementViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDetailPlacementViewRequest.Marshal(b, m, deterministic)
}
func (m *GetDetailPlacementViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDetailPlacementViewRequest.Merge(m, src)
}
func (m *GetDetailPlacementViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetDetailPlacementViewRequest.Size(m)
}
func (m *GetDetailPlacementViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDetailPlacementViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDetailPlacementViewRequest proto.InternalMessageInfo

func (m *GetDetailPlacementViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetDetailPlacementViewRequest)(nil), "google.ads.googleads.v2.services.GetDetailPlacementViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/detail_placement_view_service.proto", fileDescriptor_24eb6ba92ccb047b)
}

var fileDescriptor_24eb6ba92ccb047b = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x31, 0x8b, 0xd4, 0x40,
	0x18, 0x25, 0x11, 0x04, 0x83, 0x36, 0x29, 0xf4, 0x8c, 0xca, 0x2d, 0xc7, 0x15, 0x72, 0xc8, 0x0c,
	0x44, 0x39, 0x64, 0xe4, 0x90, 0x59, 0x0e, 0xd6, 0xea, 0x58, 0x14, 0x02, 0x4a, 0x20, 0xcc, 0x25,
	0x9f, 0x71, 0x20, 0x99, 0x89, 0x99, 0xd9, 0x1c, 0x28, 0x36, 0x16, 0xfe, 0x01, 0x6b, 0x1b, 0x4b,
	0x7f, 0xca, 0xb5, 0x76, 0x56, 0x0a, 0x56, 0xfe, 0x04, 0x2b, 0xc9, 0xce, 0x4c, 0x6e, 0x17, 0x12,
	0xaf, 0x7b, 0xec, 0x7b, 0xdf, 0xfb, 0xbe, 0x79, 0x6f, 0x13, 0x1c, 0x97, 0x52, 0x96, 0x15, 0x60,
	0x56, 0x28, 0x6c, 0x60, 0x8f, 0xba, 0x18, 0x2b, 0x68, 0x3b, 0x9e, 0x83, 0xc2, 0x05, 0x68, 0xc6,
	0xab, 0xac, 0xa9, 0x58, 0x0e, 0x35, 0x08, 0x9d, 0x75, 0x1c, 0xce, 0x32, 0x4b, 0xa3, 0xa6, 0x95,
	0x5a, 0x86, 0x33, 0x33, 0x8a, 0x58, 0xa1, 0xd0, 0xe0, 0x82, 0xba, 0x18, 0x39, 0x97, 0xe8, 0x68,
	0x6a, 0x4f, 0x0b, 0x4a, 0xae, 0xda, 0xc9, 0x45, 0x66, 0x41, 0x74, 0xd7, 0x8d, 0x37, 0x1c, 0x33,
	0x21, 0xa4, 0x66, 0x9a, 0x4b, 0xa1, 0x2c, 0x7b, 0x6b, 0x83, 0xcd, 0x2b, 0x0e, 0x42, 0x5b, 0x62,
	0x77, 0x83, 0x78, 0xcd, 0xa1, 0x2a, 0xb2, 0x53, 0x78, 0xc3, 0x3a, 0x2e, 0x5b, 0x2b, 0xb8, 0xbd,
	0x21, 0x70, 0x97, 0x18, 0x6a, 0xef, 0x5d, 0x70, 0x6f, 0x01, 0xfa, 0x78, 0x7d, 0xd4, 0xd2, 0xdd,
	0x94, 0x70, 0x38, 0x7b, 0x0e, 0x6f, 0x57, 0xa0, 0x74, 0xf8, 0x32, 0xb8, 0xe1, 0x46, 0x32, 0xc1,
	0x6a, 0xd8, 0xf1, 0x66, 0xde, 0xfd, 0x6b, 0xf3, 0x47, 0x3f, 0xa9, 0xff, 0x97, 0xa2, 0xe0, 0xc1,
	0x45, 0x10, 0x16, 0x35, 0x5c, 0xa1, 0x5c, 0xd6, 0x78, 0xcc, 0xf3, 0xba, 0xb3, 0x3a, 0x61, 0x35,
	0xc4, 0x5f, 0xfc, 0x20, 0x1a, 0x51, 0xbd, 0x30, 0x69, 0x86, 0xbf, 0xbc, 0xe0, 0xe6, 0xf8, 0x6d,
	0xe1, 0x53, 0x74, 0x59, 0x15, 0xe8, 0xbf, 0xaf, 0x8a, 0x0e, 0x27, 0x0d, 0x86, 0xa6, 0xd0, 0xc8,
	0xf8, 0xde, 0xc9, 0x0f, 0xba, 0x1d, 0xc7, 0xc7, 0xef, 0xbf, 0x3f, 0xfb, 0x8f, 0xc3, 0xc3, 0xbe,
	0xe4, 0xf7, 0x5b, 0xcc, 0x51, 0xbe, 0x52, 0x5a, 0xd6, 0xd0, 0x2a, 0x7c, 0x60, 0x5b, 0xdf, 0xf2,
	0x52, 0xf8, 0xe0, 0x43, 0x74, 0xe7, 0x9c, 0xee, 0x4c, 0x25, 0x38, 0xff, 0xe4, 0x07, 0xfb, 0xb9,
	0xac, 0x2f, 0x7d, 0xeb, 0x7c, 0x77, 0x3a, 0xc5, 0x65, 0xdf, 0xf2, 0xd2, 0x7b, 0xf5, 0xcc, 0x9a,
	0x94, 0xb2, 0x62, 0xa2, 0x44, 0xb2, 0x2d, 0x71, 0x09, 0x62, 0xfd, 0x1f, 0xc0, 0x17, 0x6b, 0xa7,
	0x3f, 0x90, 0x27, 0x0e, 0x7c, 0xf5, 0xaf, 0x2c, 0x28, 0xfd, 0xe6, 0xcf, 0x16, 0xc6, 0x90, 0x16,
	0x0a, 0x19, 0xd8, 0xa3, 0x24, 0x46, 0x76, 0xb1, 0x3a, 0x77, 0x92, 0x94, 0x16, 0x2a, 0x1d, 0x24,
	0x69, 0x12, 0xa7, 0x4e, 0xf2, 0xc7, 0xdf, 0x37, 0xbf, 0x13, 0x42, 0x0b, 0x45, 0xc8, 0x20, 0x22,
	0x24, 0x89, 0x09, 0x71, 0xb2, 0xd3, 0xab, 0xeb, 0x3b, 0x1f, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff,
	0x7f, 0x40, 0xb5, 0x40, 0xc7, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DetailPlacementViewServiceClient is the client API for DetailPlacementViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DetailPlacementViewServiceClient interface {
	// Returns the requested Detail Placement view in full detail.
	GetDetailPlacementView(ctx context.Context, in *GetDetailPlacementViewRequest, opts ...grpc.CallOption) (*resources.DetailPlacementView, error)
}

type detailPlacementViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDetailPlacementViewServiceClient(cc grpc.ClientConnInterface) DetailPlacementViewServiceClient {
	return &detailPlacementViewServiceClient{cc}
}

func (c *detailPlacementViewServiceClient) GetDetailPlacementView(ctx context.Context, in *GetDetailPlacementViewRequest, opts ...grpc.CallOption) (*resources.DetailPlacementView, error) {
	out := new(resources.DetailPlacementView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.DetailPlacementViewService/GetDetailPlacementView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DetailPlacementViewServiceServer is the server API for DetailPlacementViewService service.
type DetailPlacementViewServiceServer interface {
	// Returns the requested Detail Placement view in full detail.
	GetDetailPlacementView(context.Context, *GetDetailPlacementViewRequest) (*resources.DetailPlacementView, error)
}

// UnimplementedDetailPlacementViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDetailPlacementViewServiceServer struct {
}

func (*UnimplementedDetailPlacementViewServiceServer) GetDetailPlacementView(ctx context.Context, req *GetDetailPlacementViewRequest) (*resources.DetailPlacementView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetailPlacementView not implemented")
}

func RegisterDetailPlacementViewServiceServer(s *grpc.Server, srv DetailPlacementViewServiceServer) {
	s.RegisterService(&_DetailPlacementViewService_serviceDesc, srv)
}

func _DetailPlacementViewService_GetDetailPlacementView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDetailPlacementViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DetailPlacementViewServiceServer).GetDetailPlacementView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.DetailPlacementViewService/GetDetailPlacementView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DetailPlacementViewServiceServer).GetDetailPlacementView(ctx, req.(*GetDetailPlacementViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DetailPlacementViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.DetailPlacementViewService",
	HandlerType: (*DetailPlacementViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDetailPlacementView",
			Handler:    _DetailPlacementViewService_GetDetailPlacementView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/detail_placement_view_service.proto",
}