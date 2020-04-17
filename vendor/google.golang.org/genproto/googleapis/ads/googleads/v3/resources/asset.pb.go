// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/asset.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v3/common"
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

// Asset is a part of an ad which can be shared across multiple ads.
// It can be an image (ImageAsset), a video (YoutubeVideoAsset), etc.
type Asset struct {
	// Immutable. The resource name of the asset.
	// Asset resource names have the form:
	//
	// `customers/{customer_id}/assets/{asset_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the asset.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Optional name of the asset.
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Type of the asset.
	Type enums.AssetTypeEnum_AssetType `protobuf:"varint,4,opt,name=type,proto3,enum=google.ads.googleads.v3.enums.AssetTypeEnum_AssetType" json:"type,omitempty"`
	// The specific type of the asset.
	//
	// Types that are valid to be assigned to AssetData:
	//	*Asset_YoutubeVideoAsset
	//	*Asset_MediaBundleAsset
	//	*Asset_ImageAsset
	//	*Asset_TextAsset
	AssetData            isAsset_AssetData `protobuf_oneof:"asset_data"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_8cc8f6a72ea26cf4, []int{0}
}

func (m *Asset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset.Unmarshal(m, b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
}
func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}
func (m *Asset) XXX_Size() int {
	return xxx_messageInfo_Asset.Size(m)
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Asset) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Asset) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Asset) GetType() enums.AssetTypeEnum_AssetType {
	if m != nil {
		return m.Type
	}
	return enums.AssetTypeEnum_UNSPECIFIED
}

type isAsset_AssetData interface {
	isAsset_AssetData()
}

type Asset_YoutubeVideoAsset struct {
	YoutubeVideoAsset *common.YoutubeVideoAsset `protobuf:"bytes,5,opt,name=youtube_video_asset,json=youtubeVideoAsset,proto3,oneof"`
}

type Asset_MediaBundleAsset struct {
	MediaBundleAsset *common.MediaBundleAsset `protobuf:"bytes,6,opt,name=media_bundle_asset,json=mediaBundleAsset,proto3,oneof"`
}

type Asset_ImageAsset struct {
	ImageAsset *common.ImageAsset `protobuf:"bytes,7,opt,name=image_asset,json=imageAsset,proto3,oneof"`
}

type Asset_TextAsset struct {
	TextAsset *common.TextAsset `protobuf:"bytes,8,opt,name=text_asset,json=textAsset,proto3,oneof"`
}

func (*Asset_YoutubeVideoAsset) isAsset_AssetData() {}

func (*Asset_MediaBundleAsset) isAsset_AssetData() {}

func (*Asset_ImageAsset) isAsset_AssetData() {}

func (*Asset_TextAsset) isAsset_AssetData() {}

func (m *Asset) GetAssetData() isAsset_AssetData {
	if m != nil {
		return m.AssetData
	}
	return nil
}

func (m *Asset) GetYoutubeVideoAsset() *common.YoutubeVideoAsset {
	if x, ok := m.GetAssetData().(*Asset_YoutubeVideoAsset); ok {
		return x.YoutubeVideoAsset
	}
	return nil
}

func (m *Asset) GetMediaBundleAsset() *common.MediaBundleAsset {
	if x, ok := m.GetAssetData().(*Asset_MediaBundleAsset); ok {
		return x.MediaBundleAsset
	}
	return nil
}

func (m *Asset) GetImageAsset() *common.ImageAsset {
	if x, ok := m.GetAssetData().(*Asset_ImageAsset); ok {
		return x.ImageAsset
	}
	return nil
}

func (m *Asset) GetTextAsset() *common.TextAsset {
	if x, ok := m.GetAssetData().(*Asset_TextAsset); ok {
		return x.TextAsset
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Asset) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Asset_YoutubeVideoAsset)(nil),
		(*Asset_MediaBundleAsset)(nil),
		(*Asset_ImageAsset)(nil),
		(*Asset_TextAsset)(nil),
	}
}

func init() {
	proto.RegisterType((*Asset)(nil), "google.ads.googleads.v3.resources.Asset")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/asset.proto", fileDescriptor_8cc8f6a72ea26cf4)
}

var fileDescriptor_8cc8f6a72ea26cf4 = []byte{
	// 598 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0x97, 0x74, 0x1d, 0xcc, 0x1b, 0x08, 0xc2, 0xa5, 0x8c, 0x69, 0x74, 0xa0, 0xa1, 0x81,
	0x84, 0x53, 0x16, 0xb4, 0x43, 0x38, 0x25, 0xd2, 0xb4, 0x0d, 0x04, 0x9b, 0xc2, 0x54, 0x09, 0x54,
	0x29, 0x72, 0x6b, 0x2f, 0xb3, 0x54, 0xdb, 0x51, 0xec, 0x94, 0x55, 0xd3, 0x5e, 0x86, 0x23, 0x8f,
	0xc2, 0x0b, 0x70, 0xdd, 0x79, 0x2f, 0x80, 0xc4, 0x09, 0xc5, 0x4e, 0xd2, 0x6e, 0xa8, 0xf4, 0xd4,
	0x2f, 0xf9, 0xfe, 0xff, 0xdf, 0xf7, 0x97, 0xf3, 0xb9, 0xe0, 0x75, 0x22, 0x44, 0x32, 0x24, 0x2e,
	0xc2, 0xd2, 0x35, 0x65, 0x51, 0x8d, 0x3c, 0x37, 0x23, 0x52, 0xe4, 0xd9, 0x80, 0x48, 0x17, 0x49,
	0x49, 0x14, 0x4c, 0x33, 0xa1, 0x84, 0xb3, 0x69, 0x34, 0x10, 0x61, 0x09, 0x6b, 0x39, 0x1c, 0x79,
	0xb0, 0x96, 0xaf, 0x75, 0x66, 0x11, 0x07, 0x82, 0x31, 0xc1, 0x0d, 0x2e, 0x56, 0xe3, 0x94, 0x48,
	0x03, 0x5d, 0x83, 0xb3, 0x1c, 0x84, 0xe7, 0x4c, 0x4e, 0x19, 0x4a, 0xfd, 0xd3, 0x4a, 0x9f, 0x52,
	0xf7, 0x94, 0x92, 0x21, 0x8e, 0xfb, 0xe4, 0x0c, 0x8d, 0xa8, 0xc8, 0x4a, 0xc1, 0xe3, 0x29, 0x41,
	0x15, 0xac, 0x6c, 0x6d, 0x94, 0x2d, 0xfd, 0xd4, 0xcf, 0x4f, 0xdd, 0x6f, 0x19, 0x4a, 0x53, 0x92,
	0x55, 0x59, 0xd6, 0xa7, 0xac, 0x88, 0x73, 0xa1, 0x90, 0xa2, 0x82, 0x97, 0xdd, 0x67, 0xbf, 0x9a,
	0xa0, 0x19, 0x14, 0x71, 0x9c, 0x0f, 0xe0, 0x5e, 0x45, 0x8e, 0x39, 0x62, 0xa4, 0x65, 0xb5, 0xad,
	0xed, 0xe5, 0xf0, 0xc5, 0x55, 0xd0, 0xfc, 0x13, 0xb4, 0xc1, 0xc6, 0xe4, 0x70, 0xca, 0x2a, 0xa5,
	0x12, 0x0e, 0x04, 0x73, 0xb5, 0x3d, 0x5a, 0xad, 0xcc, 0x9f, 0x10, 0x23, 0x4e, 0x07, 0xd8, 0x14,
	0xb7, 0xec, 0xb6, 0xb5, 0xbd, 0xb2, 0xf3, 0xa4, 0x34, 0xc0, 0x2a, 0x21, 0x3c, 0xe4, 0x6a, 0xf7,
	0x6d, 0x17, 0x0d, 0x73, 0x12, 0x36, 0xae, 0x82, 0x46, 0x64, 0x53, 0xec, 0x74, 0xc0, 0xa2, 0x9e,
	0xda, 0xd0, 0x9e, 0xf5, 0x7f, 0x3c, 0x9f, 0x55, 0x46, 0x79, 0xa2, 0x4d, 0x91, 0x56, 0x3a, 0x47,
	0x60, 0xb1, 0x38, 0xc2, 0xd6, 0x62, 0xdb, 0xda, 0xbe, 0xbf, 0xb3, 0x0b, 0x67, 0x7d, 0x48, 0x7d,
	0xe6, 0x50, 0xa7, 0x3c, 0x19, 0xa7, 0x64, 0x8f, 0xe7, 0x6c, 0xf2, 0x64, 0x02, 0x68, 0x90, 0x73,
	0x06, 0x1e, 0x8d, 0x45, 0xae, 0xf2, 0x3e, 0x89, 0x47, 0x14, 0x13, 0x11, 0xeb, 0xef, 0xd4, 0x6a,
	0xea, 0x44, 0x6f, 0x66, 0xf2, 0xcd, 0x16, 0xc0, 0x2f, 0xc6, 0xda, 0x2d, 0x9c, 0x1a, 0x5f, 0xa0,
	0x9b, 0x07, 0x0b, 0xd1, 0xc3, 0xf1, 0xed, 0x8e, 0x43, 0x80, 0xc3, 0x08, 0xa6, 0x28, 0xee, 0xe7,
	0x1c, 0x0f, 0x49, 0x39, 0x68, 0x49, 0x0f, 0xea, 0xcc, 0x1b, 0xf4, 0xb1, 0x70, 0x86, 0xda, 0x78,
	0x63, 0xce, 0x03, 0x76, 0xab, 0xe1, 0x44, 0x60, 0x85, 0x32, 0x94, 0x54, 0xfc, 0x3b, 0x9a, 0xff,
	0x6a, 0x1e, 0xff, 0xb0, 0xb0, 0xd4, 0xe4, 0xc6, 0xc1, 0x42, 0x04, 0x68, 0xfd, 0xca, 0x39, 0x02,
	0x40, 0x91, 0x73, 0x55, 0x22, 0xef, 0x6a, 0xe4, 0xcb, 0x79, 0xc8, 0x13, 0x72, 0xae, 0x6e, 0x10,
	0x97, 0x55, 0xf5, 0xc6, 0x3f, 0xb8, 0x0e, 0xf6, 0xe6, 0x6d, 0x97, 0xf3, 0x7c, 0x90, 0x4b, 0x25,
	0x18, 0xc9, 0xa4, 0x7b, 0x51, 0x95, 0x97, 0xe6, 0x1e, 0x49, 0xf7, 0x42, 0xff, 0x5e, 0x86, 0xab,
	0x00, 0x98, 0x9b, 0x85, 0x91, 0x42, 0xe1, 0x6f, 0x0b, 0x6c, 0x0d, 0x04, 0x83, 0x73, 0xef, 0x77,
	0x08, 0xf4, 0x8c, 0xe3, 0x62, 0xd3, 0x8e, 0xad, 0xaf, 0xef, 0x4b, 0x43, 0x22, 0x86, 0x88, 0x27,
	0x50, 0x64, 0x89, 0x9b, 0x10, 0xae, 0xf7, 0xd0, 0x9d, 0x64, 0xfb, 0xcf, 0xdf, 0xcb, 0xbb, 0xba,
	0xfa, 0x6e, 0x37, 0xf6, 0x83, 0xe0, 0x87, 0xbd, 0xb9, 0x6f, 0x90, 0x01, 0x96, 0xd0, 0x94, 0x45,
	0xd5, 0xf5, 0x60, 0x54, 0x29, 0x7f, 0x56, 0x9a, 0x5e, 0x80, 0x65, 0xaf, 0xd6, 0xf4, 0xba, 0x5e,
	0xaf, 0xd6, 0x5c, 0xdb, 0x5b, 0xa6, 0xe1, 0xfb, 0x01, 0x96, 0xbe, 0x5f, 0xab, 0x7c, 0xbf, 0xeb,
	0xf9, 0x7e, 0xad, 0xeb, 0x2f, 0xe9, 0xb0, 0xde, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xca, 0x80,
	0x51, 0xf4, 0x0a, 0x05, 0x00, 0x00,
}