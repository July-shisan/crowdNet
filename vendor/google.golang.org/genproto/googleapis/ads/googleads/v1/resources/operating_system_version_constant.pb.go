// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/operating_system_version_constant.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A mobile operating system version or a range of versions, depending on
// `operator_type`. List of available mobile platforms at
// https://developers.google.com/adwords/api/docs/appendix/codes-formats#mobile-platforms
type OperatingSystemVersionConstant struct {
	// Output only. The resource name of the operating system version constant.
	// Operating system version constant resource names have the form:
	//
	// `operatingSystemVersionConstants/{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the operating system version.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. Name of the operating system.
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The OS Major Version number.
	OsMajorVersion *wrappers.Int32Value `protobuf:"bytes,4,opt,name=os_major_version,json=osMajorVersion,proto3" json:"os_major_version,omitempty"`
	// Output only. The OS Minor Version number.
	OsMinorVersion *wrappers.Int32Value `protobuf:"bytes,5,opt,name=os_minor_version,json=osMinorVersion,proto3" json:"os_minor_version,omitempty"`
	// Output only. Determines whether this constant represents a single version or a range of
	// versions.
	OperatorType         enums.OperatingSystemVersionOperatorTypeEnum_OperatingSystemVersionOperatorType `protobuf:"varint,6,opt,name=operator_type,json=operatorType,proto3,enum=google.ads.googleads.v1.enums.OperatingSystemVersionOperatorTypeEnum_OperatingSystemVersionOperatorType" json:"operator_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                        `json:"-"`
	XXX_unrecognized     []byte                                                                          `json:"-"`
	XXX_sizecache        int32                                                                           `json:"-"`
}

func (m *OperatingSystemVersionConstant) Reset()         { *m = OperatingSystemVersionConstant{} }
func (m *OperatingSystemVersionConstant) String() string { return proto.CompactTextString(m) }
func (*OperatingSystemVersionConstant) ProtoMessage()    {}
func (*OperatingSystemVersionConstant) Descriptor() ([]byte, []int) {
	return fileDescriptor_255222b6c86bb69b, []int{0}
}

func (m *OperatingSystemVersionConstant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperatingSystemVersionConstant.Unmarshal(m, b)
}
func (m *OperatingSystemVersionConstant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperatingSystemVersionConstant.Marshal(b, m, deterministic)
}
func (m *OperatingSystemVersionConstant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperatingSystemVersionConstant.Merge(m, src)
}
func (m *OperatingSystemVersionConstant) XXX_Size() int {
	return xxx_messageInfo_OperatingSystemVersionConstant.Size(m)
}
func (m *OperatingSystemVersionConstant) XXX_DiscardUnknown() {
	xxx_messageInfo_OperatingSystemVersionConstant.DiscardUnknown(m)
}

var xxx_messageInfo_OperatingSystemVersionConstant proto.InternalMessageInfo

func (m *OperatingSystemVersionConstant) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *OperatingSystemVersionConstant) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *OperatingSystemVersionConstant) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *OperatingSystemVersionConstant) GetOsMajorVersion() *wrappers.Int32Value {
	if m != nil {
		return m.OsMajorVersion
	}
	return nil
}

func (m *OperatingSystemVersionConstant) GetOsMinorVersion() *wrappers.Int32Value {
	if m != nil {
		return m.OsMinorVersion
	}
	return nil
}

func (m *OperatingSystemVersionConstant) GetOperatorType() enums.OperatingSystemVersionOperatorTypeEnum_OperatingSystemVersionOperatorType {
	if m != nil {
		return m.OperatorType
	}
	return enums.OperatingSystemVersionOperatorTypeEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*OperatingSystemVersionConstant)(nil), "google.ads.googleads.v1.resources.OperatingSystemVersionConstant")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/operating_system_version_constant.proto", fileDescriptor_255222b6c86bb69b)
}

var fileDescriptor_255222b6c86bb69b = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0x26, 0xd9, 0xb6, 0x60, 0x6c, 0x8b, 0xe4, 0x6a, 0xad, 0xa5, 0x6e, 0x95, 0xc2, 0x5e, 0xcd,
	0xb8, 0x5b, 0x7f, 0x30, 0x5e, 0x48, 0xb6, 0x48, 0x69, 0x41, 0x5b, 0xb6, 0x92, 0x0b, 0x59, 0x08,
	0xb3, 0x9b, 0x69, 0x3a, 0xb2, 0x99, 0x13, 0x66, 0x26, 0x2b, 0x8b, 0x28, 0xf8, 0x06, 0xe2, 0x13,
	0x88, 0x97, 0x3e, 0x8a, 0x4f, 0xd1, 0xeb, 0x3e, 0x82, 0x57, 0x92, 0x64, 0x26, 0x8d, 0xe8, 0xa6,
	0xda, 0xbb, 0xb3, 0xfb, 0xfd, 0x9c, 0x6f, 0xe6, 0xcc, 0x89, 0x73, 0x10, 0x03, 0xc4, 0x53, 0x8a,
	0x49, 0x24, 0x71, 0x59, 0xe6, 0xd5, 0xac, 0x87, 0x05, 0x95, 0x90, 0x89, 0x09, 0x95, 0x18, 0x52,
	0x2a, 0x88, 0x62, 0x3c, 0x0e, 0xe5, 0x5c, 0x2a, 0x9a, 0x84, 0x33, 0x2a, 0x24, 0x03, 0x1e, 0x4e,
	0x80, 0x4b, 0x45, 0xb8, 0x42, 0xa9, 0x00, 0x05, 0xee, 0x76, 0xa9, 0x47, 0x24, 0x92, 0xa8, 0xb2,
	0x42, 0xb3, 0x1e, 0xaa, 0xac, 0x36, 0x0e, 0x17, 0x75, 0xa3, 0x3c, 0x4b, 0x1a, 0x3a, 0x95, 0x00,
	0x88, 0x50, 0xcd, 0x53, 0x5a, 0xb6, 0xdb, 0xb8, 0x6b, 0xbc, 0x52, 0x86, 0x4f, 0x19, 0x9d, 0x46,
	0xe1, 0x98, 0x9e, 0x91, 0x19, 0x03, 0xa1, 0x09, 0xb7, 0x6b, 0x04, 0x13, 0x41, 0x43, 0x5b, 0x1a,
	0x2a, 0x7e, 0x8d, 0xb3, 0x53, 0xfc, 0x4e, 0x90, 0x34, 0xa5, 0x42, 0x6a, 0x7c, 0xb3, 0x26, 0x25,
	0x9c, 0x83, 0x22, 0x8a, 0x01, 0xd7, 0xe8, 0xbd, 0xaf, 0xcb, 0xce, 0xd6, 0x91, 0x89, 0x7a, 0x52,
	0x24, 0x0d, 0xca, 0xa0, 0x7b, 0xfa, 0x46, 0xdc, 0xc8, 0x59, 0x33, 0x2d, 0x43, 0x4e, 0x12, 0xda,
	0xb6, 0x3a, 0x56, 0xf7, 0xc6, 0xe0, 0xf9, 0xb9, 0xdf, 0xfa, 0xe9, 0x3f, 0x75, 0x9e, 0x5c, 0xde,
	0x8f, 0xae, 0x52, 0x26, 0xd1, 0x04, 0x12, 0xdc, 0xec, 0x3b, 0x5c, 0x35, 0xae, 0xaf, 0x48, 0x42,
	0xdd, 0x07, 0x8e, 0xcd, 0xa2, 0xb6, 0xdd, 0xb1, 0xba, 0x37, 0xfb, 0x77, 0xb4, 0x13, 0x32, 0x67,
	0x42, 0x07, 0x5c, 0x3d, 0x7e, 0x18, 0x90, 0x69, 0x46, 0x07, 0xad, 0x73, 0xbf, 0x35, 0xb4, 0x59,
	0xe4, 0x3e, 0x72, 0x96, 0x8a, 0x38, 0xad, 0x42, 0xb3, 0xf9, 0x87, 0xe6, 0x44, 0x09, 0xc6, 0xe3,
	0x9a, 0xa8, 0xa0, 0xbb, 0x87, 0xce, 0x2d, 0x90, 0x61, 0x42, 0xde, 0x82, 0x30, 0x33, 0x69, 0x2f,
	0x2d, 0x6e, 0xbb, 0xdb, 0xaf, 0x39, 0xac, 0x83, 0x7c, 0x99, 0x0b, 0xf5, 0x51, 0x8c, 0x17, 0xe3,
	0x35, 0xaf, 0xe5, 0xff, 0xf0, 0xca, 0x85, 0xc6, 0xeb, 0x8b, 0xe5, 0xac, 0xfd, 0xf6, 0x36, 0xda,
	0x2b, 0x1d, 0xab, 0xbb, 0xde, 0x3f, 0x43, 0x8b, 0xde, 0x62, 0xf1, 0xd0, 0xd0, 0xdf, 0x6f, 0xf9,
	0x48, 0x3b, 0xbd, 0x9e, 0xa7, 0xf4, 0x05, 0xcf, 0x92, 0x7f, 0xa0, 0x95, 0xb1, 0x56, 0xa1, 0xf6,
	0x97, 0xf7, 0xc9, 0xba, 0xf0, 0x3f, 0x5e, 0x7b, 0xc4, 0xee, 0x1e, 0x34, 0xe2, 0x12, 0xbf, 0xbf,
	0x72, 0x21, 0x3f, 0x0c, 0x3e, 0xdb, 0xce, 0xce, 0x04, 0x12, 0x74, 0xe5, 0x4a, 0x0e, 0xee, 0x37,
	0xc7, 0x39, 0xce, 0x47, 0x71, 0x6c, 0xbd, 0xd1, 0x9b, 0x8b, 0x62, 0x98, 0x12, 0x1e, 0x23, 0x10,
	0x31, 0x8e, 0x29, 0x2f, 0x06, 0x85, 0x2f, 0xcf, 0xd7, 0xf0, 0x19, 0x79, 0x56, 0x55, 0xdf, 0xec,
	0xd6, 0xbe, 0xef, 0x7f, 0xb7, 0xb7, 0xf7, 0x4b, 0x4b, 0x3f, 0x92, 0xa8, 0x2c, 0xf3, 0x2a, 0xe8,
	0xa1, 0xa1, 0x61, 0xfe, 0x30, 0x9c, 0x91, 0x1f, 0xc9, 0x51, 0xc5, 0x19, 0x05, 0xbd, 0x51, 0xc5,
	0xb9, 0xb0, 0x77, 0x4a, 0xc0, 0xf3, 0xfc, 0x48, 0x7a, 0x5e, 0xc5, 0xf2, 0xbc, 0xa0, 0xe7, 0x79,
	0x15, 0x6f, 0xbc, 0x52, 0x84, 0xdd, 0xfd, 0x15, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xe9, 0x6b, 0xdc,
	0xf2, 0x04, 0x00, 0x00,
}
