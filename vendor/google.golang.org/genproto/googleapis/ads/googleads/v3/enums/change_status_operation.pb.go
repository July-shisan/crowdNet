// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/change_status_operation.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Status of the changed resource
type ChangeStatusOperationEnum_ChangeStatusOperation int32

const (
	// No value has been specified.
	ChangeStatusOperationEnum_UNSPECIFIED ChangeStatusOperationEnum_ChangeStatusOperation = 0
	// Used for return value only. Represents an unclassified resource unknown
	// in this version.
	ChangeStatusOperationEnum_UNKNOWN ChangeStatusOperationEnum_ChangeStatusOperation = 1
	// The resource was created.
	ChangeStatusOperationEnum_ADDED ChangeStatusOperationEnum_ChangeStatusOperation = 2
	// The resource was modified.
	ChangeStatusOperationEnum_CHANGED ChangeStatusOperationEnum_ChangeStatusOperation = 3
	// The resource was removed.
	ChangeStatusOperationEnum_REMOVED ChangeStatusOperationEnum_ChangeStatusOperation = 4
)

var ChangeStatusOperationEnum_ChangeStatusOperation_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ADDED",
	3: "CHANGED",
	4: "REMOVED",
}

var ChangeStatusOperationEnum_ChangeStatusOperation_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ADDED":       2,
	"CHANGED":     3,
	"REMOVED":     4,
}

func (x ChangeStatusOperationEnum_ChangeStatusOperation) String() string {
	return proto.EnumName(ChangeStatusOperationEnum_ChangeStatusOperation_name, int32(x))
}

func (ChangeStatusOperationEnum_ChangeStatusOperation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_52f8f4521d0e81d3, []int{0, 0}
}

// Container for enum describing operations for the ChangeStatus resource.
type ChangeStatusOperationEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeStatusOperationEnum) Reset()         { *m = ChangeStatusOperationEnum{} }
func (m *ChangeStatusOperationEnum) String() string { return proto.CompactTextString(m) }
func (*ChangeStatusOperationEnum) ProtoMessage()    {}
func (*ChangeStatusOperationEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_52f8f4521d0e81d3, []int{0}
}

func (m *ChangeStatusOperationEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeStatusOperationEnum.Unmarshal(m, b)
}
func (m *ChangeStatusOperationEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeStatusOperationEnum.Marshal(b, m, deterministic)
}
func (m *ChangeStatusOperationEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeStatusOperationEnum.Merge(m, src)
}
func (m *ChangeStatusOperationEnum) XXX_Size() int {
	return xxx_messageInfo_ChangeStatusOperationEnum.Size(m)
}
func (m *ChangeStatusOperationEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeStatusOperationEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeStatusOperationEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.ChangeStatusOperationEnum_ChangeStatusOperation", ChangeStatusOperationEnum_ChangeStatusOperation_name, ChangeStatusOperationEnum_ChangeStatusOperation_value)
	proto.RegisterType((*ChangeStatusOperationEnum)(nil), "google.ads.googleads.v3.enums.ChangeStatusOperationEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/change_status_operation.proto", fileDescriptor_52f8f4521d0e81d3)
}

var fileDescriptor_52f8f4521d0e81d3 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x4e, 0x02, 0x31,
	0x14, 0xc6, 0x65, 0xf0, 0x4f, 0x2c, 0x0b, 0x27, 0x93, 0xb8, 0x90, 0xc8, 0x02, 0x0e, 0xd0, 0x2e,
	0xba, 0x2b, 0xab, 0x42, 0x2b, 0x12, 0xe3, 0x40, 0x24, 0x8c, 0x09, 0x99, 0x84, 0x54, 0x66, 0x52,
	0x49, 0xa0, 0x9d, 0xd0, 0x01, 0xef, 0xe3, 0xd2, 0xa3, 0x78, 0x14, 0x37, 0x5e, 0xc1, 0xb4, 0x75,
	0x66, 0x85, 0x6e, 0x9a, 0xaf, 0xef, 0x7b, 0xbf, 0xaf, 0xaf, 0x0f, 0xf4, 0xa5, 0xd6, 0x72, 0x93,
	0x23, 0x91, 0x19, 0xe4, 0xa5, 0x55, 0x07, 0x8c, 0x72, 0xb5, 0xdf, 0x1a, 0xb4, 0x7a, 0x15, 0x4a,
	0xe6, 0x4b, 0x53, 0x8a, 0x72, 0x6f, 0x96, 0xba, 0xc8, 0x77, 0xa2, 0x5c, 0x6b, 0x05, 0x8b, 0x9d,
	0x2e, 0x75, 0xd4, 0xf1, 0x04, 0x14, 0x99, 0x81, 0x35, 0x0c, 0x0f, 0x18, 0x3a, 0xb8, 0x7d, 0x5b,
	0x65, 0x17, 0x6b, 0x24, 0x94, 0xd2, 0xa5, 0x63, 0x8d, 0x87, 0x7b, 0x6f, 0xe0, 0x66, 0xe8, 0xd2,
	0x67, 0x2e, 0x7c, 0x52, 0x65, 0x73, 0xb5, 0xdf, 0xf6, 0x16, 0xe0, 0xfa, 0xa8, 0x19, 0x5d, 0x81,
	0xd6, 0x3c, 0x9e, 0x4d, 0xf9, 0x70, 0x7c, 0x37, 0xe6, 0x2c, 0x3c, 0x89, 0x5a, 0xe0, 0x62, 0x1e,
	0x3f, 0xc4, 0x93, 0xe7, 0x38, 0x6c, 0x44, 0x97, 0xe0, 0x8c, 0x32, 0xc6, 0x59, 0x18, 0xd8, 0xfa,
	0xf0, 0x9e, 0xc6, 0x23, 0xce, 0xc2, 0xa6, 0xbd, 0x3c, 0xf1, 0xc7, 0x49, 0xc2, 0x59, 0x78, 0x3a,
	0xf8, 0x6e, 0x80, 0xee, 0x4a, 0x6f, 0xe1, 0xbf, 0xc3, 0x0f, 0xda, 0x47, 0xdf, 0x9f, 0xda, 0xd1,
	0xa7, 0x8d, 0xc5, 0xe0, 0x17, 0x96, 0x7a, 0x23, 0x94, 0x84, 0x7a, 0x27, 0x91, 0xcc, 0x95, 0xfb,
	0x58, 0xb5, 0xc6, 0x62, 0x6d, 0xfe, 0xd8, 0x6a, 0xdf, 0x9d, 0xef, 0x41, 0x73, 0x44, 0xe9, 0x47,
	0xd0, 0x19, 0xf9, 0x28, 0x9a, 0x19, 0xe8, 0xa5, 0x55, 0x09, 0x86, 0x76, 0x11, 0xe6, 0xb3, 0xf2,
	0x53, 0x9a, 0x99, 0xb4, 0xf6, 0xd3, 0x04, 0xa7, 0xce, 0xff, 0x0a, 0xba, 0xbe, 0x48, 0x08, 0xcd,
	0x0c, 0x21, 0x75, 0x07, 0x21, 0x09, 0x26, 0xc4, 0xf5, 0xbc, 0x9c, 0xbb, 0xc1, 0xf0, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x09, 0x2a, 0x56, 0xb5, 0xed, 0x01, 0x00, 0x00,
}
