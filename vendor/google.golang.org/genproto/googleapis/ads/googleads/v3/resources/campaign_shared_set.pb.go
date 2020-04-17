// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/campaign_shared_set.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v3/enums"
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

// CampaignSharedSets are used for managing the shared sets associated with a
// campaign.
type CampaignSharedSet struct {
	// Immutable. The resource name of the campaign shared set.
	// Campaign shared set resource names have the form:
	//
	// `customers/{customer_id}/campaignSharedSets/{campaign_id}~{shared_set_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Immutable. The campaign to which the campaign shared set belongs.
	Campaign *wrappers.StringValue `protobuf:"bytes,3,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// Immutable. The shared set associated with the campaign. This may be a negative keyword
	// shared set of another customer. This customer should be a manager of the
	// other customer, otherwise the campaign shared set will exist but have no
	// serving effect. Only negative keyword shared sets can be associated with
	// Shopping campaigns. Only negative placement shared sets can be associated
	// with Display mobile app campaigns.
	SharedSet *wrappers.StringValue `protobuf:"bytes,4,opt,name=shared_set,json=sharedSet,proto3" json:"shared_set,omitempty"`
	// Output only. The status of this campaign shared set. Read only.
	Status               enums.CampaignSharedSetStatusEnum_CampaignSharedSetStatus `protobuf:"varint,2,opt,name=status,proto3,enum=google.ads.googleads.v3.enums.CampaignSharedSetStatusEnum_CampaignSharedSetStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                  `json:"-"`
	XXX_unrecognized     []byte                                                    `json:"-"`
	XXX_sizecache        int32                                                     `json:"-"`
}

func (m *CampaignSharedSet) Reset()         { *m = CampaignSharedSet{} }
func (m *CampaignSharedSet) String() string { return proto.CompactTextString(m) }
func (*CampaignSharedSet) ProtoMessage()    {}
func (*CampaignSharedSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d4c077729b864b, []int{0}
}

func (m *CampaignSharedSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignSharedSet.Unmarshal(m, b)
}
func (m *CampaignSharedSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignSharedSet.Marshal(b, m, deterministic)
}
func (m *CampaignSharedSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignSharedSet.Merge(m, src)
}
func (m *CampaignSharedSet) XXX_Size() int {
	return xxx_messageInfo_CampaignSharedSet.Size(m)
}
func (m *CampaignSharedSet) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignSharedSet.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignSharedSet proto.InternalMessageInfo

func (m *CampaignSharedSet) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignSharedSet) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *CampaignSharedSet) GetSharedSet() *wrappers.StringValue {
	if m != nil {
		return m.SharedSet
	}
	return nil
}

func (m *CampaignSharedSet) GetStatus() enums.CampaignSharedSetStatusEnum_CampaignSharedSetStatus {
	if m != nil {
		return m.Status
	}
	return enums.CampaignSharedSetStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*CampaignSharedSet)(nil), "google.ads.googleads.v3.resources.CampaignSharedSet")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/campaign_shared_set.proto", fileDescriptor_41d4c077729b864b)
}

var fileDescriptor_41d4c077729b864b = []byte{
	// 487 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4f, 0x6b, 0xdb, 0x3e,
	0x18, 0xc6, 0x4e, 0x7f, 0xe5, 0x57, 0xed, 0x0f, 0xcc, 0x87, 0x91, 0x95, 0xb2, 0x25, 0xdd, 0x0a,
	0x59, 0x19, 0x12, 0xc4, 0x37, 0x95, 0x0d, 0xe4, 0x31, 0x0a, 0x3b, 0x8c, 0xe2, 0x40, 0x06, 0x23,
	0x60, 0x14, 0x5b, 0x55, 0xcd, 0x62, 0xc9, 0x48, 0x72, 0x76, 0x18, 0x3d, 0xed, 0x63, 0xec, 0xb6,
	0xe3, 0x3e, 0xca, 0x3e, 0x45, 0xcf, 0xfd, 0x08, 0x3b, 0x8d, 0xd8, 0x92, 0x52, 0xc8, 0xb2, 0x6e,
	0xb7, 0x27, 0xbc, 0xcf, 0x9f, 0x37, 0x8f, 0x5f, 0x81, 0x13, 0x2e, 0x25, 0x5f, 0x30, 0x44, 0x0b,
	0x8d, 0x3a, 0xb8, 0x42, 0xcb, 0x18, 0x29, 0xa6, 0x65, 0xa3, 0x72, 0xa6, 0x51, 0x4e, 0xab, 0x9a,
	0x96, 0x5c, 0x64, 0xfa, 0x82, 0x2a, 0x56, 0x64, 0x9a, 0x19, 0x58, 0x2b, 0x69, 0x64, 0x34, 0xec,
	0x14, 0x90, 0x16, 0x1a, 0x7a, 0x31, 0x5c, 0xc6, 0xd0, 0x8b, 0xf7, 0x5f, 0x6d, 0xf3, 0x67, 0xa2,
	0xa9, 0x7e, 0xeb, 0x9d, 0x69, 0x43, 0x4d, 0xa3, 0xbb, 0x88, 0xfd, 0x27, 0x4e, 0x5f, 0x97, 0xe8,
	0xbc, 0x64, 0x8b, 0x22, 0x9b, 0xb3, 0x0b, 0xba, 0x2c, 0xa5, 0xb2, 0x84, 0x47, 0x37, 0x08, 0x2e,
	0xd6, 0x8e, 0x1e, 0xdb, 0x51, 0xfb, 0x6b, 0xde, 0x9c, 0xa3, 0x4f, 0x8a, 0xd6, 0x35, 0x53, 0xce,
	0xfb, 0xe0, 0x86, 0x94, 0x0a, 0x21, 0x0d, 0x35, 0xa5, 0x14, 0x76, 0x7a, 0xf8, 0x75, 0x07, 0x3c,
	0x78, 0x6d, 0xd7, 0x9b, 0xb4, 0xdb, 0x4d, 0x98, 0x89, 0xde, 0x83, 0x7b, 0x2e, 0x25, 0x13, 0xb4,
	0x62, 0xfd, 0x60, 0x10, 0x8c, 0xf6, 0x92, 0xf1, 0x15, 0xf9, 0xef, 0x27, 0x79, 0x01, 0x8e, 0xd7,
	0x35, 0x58, 0x54, 0x97, 0x1a, 0xe6, 0xb2, 0x42, 0x1b, 0x56, 0xe9, 0x5d, 0x67, 0xf4, 0x8e, 0x56,
	0x2c, 0xca, 0xc1, 0xff, 0xae, 0x8c, 0x7e, 0x6f, 0x10, 0x8c, 0xee, 0x8c, 0x0f, 0xac, 0x05, 0x74,
	0xfb, 0xc3, 0x89, 0x51, 0xa5, 0xe0, 0x53, 0xba, 0x68, 0x58, 0xf2, 0xbc, 0x4d, 0x7c, 0x0a, 0x86,
	0xb7, 0x26, 0xa6, 0xde, 0x38, 0xe2, 0x00, 0xac, 0x8b, 0xee, 0xef, 0xfc, 0x45, 0xcc, 0x71, 0x1b,
	0xf3, 0x0c, 0x1c, 0x6e, 0x8d, 0x59, 0xff, 0xa1, 0x3d, 0xed, 0x6b, 0x92, 0x60, 0xb7, 0xfb, 0x8c,
	0xfd, 0x70, 0x10, 0x8c, 0xee, 0x8f, 0x53, 0xb8, 0xed, 0x54, 0xda, 0x3b, 0x80, 0x1b, 0xed, 0x4c,
	0x5a, 0xf5, 0x1b, 0xd1, 0x54, 0xdb, 0x66, 0x49, 0xef, 0x8a, 0xf4, 0x52, 0x1b, 0x83, 0xc5, 0x35,
	0xf9, 0xf8, 0x2f, 0xed, 0x47, 0x2f, 0xf3, 0x46, 0x1b, 0x59, 0x31, 0xa5, 0xd1, 0x67, 0x07, 0x2f,
	0xfd, 0x3d, 0x7a, 0xde, 0x6a, 0xba, 0x79, 0xa3, 0x97, 0xc9, 0x97, 0x10, 0x1c, 0xe5, 0xb2, 0x82,
	0xb7, 0xbe, 0x80, 0xe4, 0xe1, 0x46, 0xf6, 0xd9, 0xaa, 0xe8, 0xb3, 0xe0, 0xc3, 0x5b, 0x2b, 0xe6,
	0x72, 0x41, 0x05, 0x87, 0x52, 0x71, 0xc4, 0x99, 0x68, 0x3f, 0x03, 0x5a, 0xef, 0xff, 0x87, 0xa7,
	0x79, 0xe2, 0xd1, 0xb7, 0xb0, 0x77, 0x4a, 0xc8, 0xf7, 0x70, 0x78, 0xda, 0x59, 0x92, 0x42, 0xc3,
	0x0e, 0xae, 0xd0, 0x34, 0x86, 0xa9, 0x63, 0xfe, 0x70, 0x9c, 0x19, 0x29, 0xf4, 0xcc, 0x73, 0x66,
	0xd3, 0x78, 0xe6, 0x39, 0xd7, 0xe1, 0x51, 0x37, 0xc0, 0x98, 0x14, 0x1a, 0x63, 0xcf, 0xc2, 0x78,
	0x1a, 0x63, 0xec, 0x79, 0xf3, 0xdd, 0x76, 0xd9, 0xf8, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x60,
	0x4c, 0xbe, 0x4a, 0x46, 0x04, 0x00, 0x00,
}
