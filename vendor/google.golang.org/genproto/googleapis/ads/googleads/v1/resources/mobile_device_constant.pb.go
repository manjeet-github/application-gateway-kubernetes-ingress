// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/mobile_device_constant.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A mobile device constant.
type MobileDeviceConstant struct {
	// The resource name of the mobile device constant.
	// Mobile device constant resource names have the form:
	//
	// `mobileDeviceConstants/{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the mobile device constant.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the mobile device.
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The manufacturer of the mobile device.
	ManufacturerName *wrappers.StringValue `protobuf:"bytes,4,opt,name=manufacturer_name,json=manufacturerName,proto3" json:"manufacturer_name,omitempty"`
	// The operating system of the mobile device.
	OperatingSystemName *wrappers.StringValue `protobuf:"bytes,5,opt,name=operating_system_name,json=operatingSystemName,proto3" json:"operating_system_name,omitempty"`
	// The type of mobile device.
	Type                 enums.MobileDeviceTypeEnum_MobileDeviceType `protobuf:"varint,6,opt,name=type,proto3,enum=google.ads.googleads.v1.enums.MobileDeviceTypeEnum_MobileDeviceType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *MobileDeviceConstant) Reset()         { *m = MobileDeviceConstant{} }
func (m *MobileDeviceConstant) String() string { return proto.CompactTextString(m) }
func (*MobileDeviceConstant) ProtoMessage()    {}
func (*MobileDeviceConstant) Descriptor() ([]byte, []int) {
	return fileDescriptor_mobile_device_constant_bb8f0eee204a71a8, []int{0}
}
func (m *MobileDeviceConstant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MobileDeviceConstant.Unmarshal(m, b)
}
func (m *MobileDeviceConstant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MobileDeviceConstant.Marshal(b, m, deterministic)
}
func (dst *MobileDeviceConstant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MobileDeviceConstant.Merge(dst, src)
}
func (m *MobileDeviceConstant) XXX_Size() int {
	return xxx_messageInfo_MobileDeviceConstant.Size(m)
}
func (m *MobileDeviceConstant) XXX_DiscardUnknown() {
	xxx_messageInfo_MobileDeviceConstant.DiscardUnknown(m)
}

var xxx_messageInfo_MobileDeviceConstant proto.InternalMessageInfo

func (m *MobileDeviceConstant) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *MobileDeviceConstant) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *MobileDeviceConstant) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *MobileDeviceConstant) GetManufacturerName() *wrappers.StringValue {
	if m != nil {
		return m.ManufacturerName
	}
	return nil
}

func (m *MobileDeviceConstant) GetOperatingSystemName() *wrappers.StringValue {
	if m != nil {
		return m.OperatingSystemName
	}
	return nil
}

func (m *MobileDeviceConstant) GetType() enums.MobileDeviceTypeEnum_MobileDeviceType {
	if m != nil {
		return m.Type
	}
	return enums.MobileDeviceTypeEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*MobileDeviceConstant)(nil), "google.ads.googleads.v1.resources.MobileDeviceConstant")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/mobile_device_constant.proto", fileDescriptor_mobile_device_constant_bb8f0eee204a71a8)
}

var fileDescriptor_mobile_device_constant_bb8f0eee204a71a8 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0x6a, 0x13, 0x41,
	0x14, 0xc6, 0xd9, 0x4d, 0x2c, 0x38, 0xfe, 0x41, 0x57, 0x85, 0x58, 0x8b, 0xa4, 0x4a, 0x21, 0x20,
	0xcc, 0xba, 0x55, 0x7a, 0xb1, 0x82, 0xb0, 0xb5, 0x52, 0x2a, 0x28, 0x21, 0x95, 0x20, 0x12, 0x08,
	0x93, 0x9d, 0xd3, 0x61, 0x20, 0xf3, 0x87, 0x99, 0xd9, 0x48, 0x5e, 0xc0, 0x7b, 0x5f, 0xc1, 0x4b,
	0x1f, 0xc5, 0x47, 0xf1, 0x29, 0x24, 0x33, 0xbb, 0x43, 0xb1, 0x56, 0x7b, 0x77, 0x32, 0xe7, 0xfb,
	0x7e, 0xf9, 0xce, 0x9e, 0x83, 0x5e, 0x33, 0xa5, 0xd8, 0x12, 0x72, 0x42, 0x6d, 0x1e, 0xca, 0x4d,
	0xb5, 0x2a, 0x72, 0x03, 0x56, 0x35, 0xa6, 0x06, 0x9b, 0x0b, 0xb5, 0xe0, 0x4b, 0x98, 0x53, 0x58,
	0xf1, 0x1a, 0xe6, 0xb5, 0x92, 0xd6, 0x11, 0xe9, 0xb0, 0x36, 0xca, 0xa9, 0x6c, 0x37, 0x98, 0x30,
	0xa1, 0x16, 0x47, 0x3f, 0x5e, 0x15, 0x38, 0xfa, 0xb7, 0x0f, 0x2e, 0xfb, 0x0b, 0x90, 0x8d, 0xf8,
	0x13, 0xef, 0xd6, 0x1a, 0x02, 0x7a, 0xfb, 0x71, 0xeb, 0xf3, 0xbf, 0x16, 0xcd, 0x59, 0xfe, 0xc5,
	0x10, 0xad, 0xc1, 0xd8, 0xb6, 0xbf, 0xd3, 0x71, 0x35, 0xcf, 0x89, 0x94, 0xca, 0x11, 0xc7, 0x95,
	0x6c, 0xbb, 0x4f, 0xbe, 0xf5, 0xd0, 0xfd, 0xf7, 0x1e, 0x7d, 0xe4, 0xc9, 0x6f, 0xda, 0xdc, 0xd9,
	0x53, 0x74, 0xab, 0xcb, 0x36, 0x97, 0x44, 0xc0, 0x20, 0x19, 0x26, 0xa3, 0xeb, 0x93, 0x9b, 0xdd,
	0xe3, 0x07, 0x22, 0x20, 0x7b, 0x86, 0x52, 0x4e, 0x07, 0xe9, 0x30, 0x19, 0xdd, 0xd8, 0x7f, 0xd4,
	0x0e, 0x86, 0xbb, 0x20, 0xf8, 0x44, 0xba, 0x83, 0x97, 0x53, 0xb2, 0x6c, 0x60, 0x92, 0x72, 0x9a,
	0x3d, 0x47, 0x7d, 0x0f, 0xea, 0x79, 0xf9, 0xce, 0x05, 0xf9, 0xa9, 0x33, 0x5c, 0xb2, 0xa0, 0xf7,
	0xca, 0xec, 0x04, 0xdd, 0x15, 0x44, 0x36, 0x67, 0xa4, 0x76, 0x8d, 0x01, 0x13, 0x72, 0xf4, 0xaf,
	0x60, 0xbf, 0x73, 0xde, 0xe6, 0x93, 0x8e, 0xd1, 0x03, 0xa5, 0xc1, 0x10, 0xc7, 0x25, 0x9b, 0xdb,
	0xb5, 0x75, 0x20, 0x02, 0xee, 0xda, 0x15, 0x70, 0xf7, 0xa2, 0xf5, 0xd4, 0x3b, 0x3d, 0xf1, 0x13,
	0xea, 0x6f, 0xb6, 0x30, 0xd8, 0x1a, 0x26, 0xa3, 0xdb, 0xfb, 0x47, 0xf8, 0xb2, 0x0d, 0xfb, 0xf5,
	0xe1, 0xf3, 0xdf, 0xf8, 0xe3, 0x5a, 0xc3, 0x5b, 0xd9, 0x88, 0x0b, 0x8f, 0x13, 0x4f, 0x3c, 0xfc,
	0x9a, 0xa2, 0xbd, 0x5a, 0x09, 0xfc, 0xdf, 0x9b, 0x39, 0x7c, 0xf8, 0xb7, 0xd5, 0x8d, 0x37, 0x23,
	0x8c, 0x93, 0xcf, 0xef, 0x5a, 0x3f, 0x53, 0x4b, 0x22, 0x19, 0x56, 0x86, 0xe5, 0x0c, 0xa4, 0x1f,
	0xb0, 0x3b, 0x30, 0xcd, 0xed, 0x3f, 0x4e, 0xfa, 0x55, 0xac, 0xbe, 0xa7, 0xbd, 0xe3, 0xaa, 0xfa,
	0x91, 0xee, 0x1e, 0x07, 0x64, 0x45, 0x2d, 0x0e, 0xe5, 0xa6, 0x9a, 0x16, 0x78, 0xd2, 0x29, 0x7f,
	0x76, 0x9a, 0x59, 0x45, 0xed, 0x2c, 0x6a, 0x66, 0xd3, 0x62, 0x16, 0x35, 0xbf, 0xd2, 0xbd, 0xd0,
	0x28, 0xcb, 0x8a, 0xda, 0xb2, 0x8c, 0xaa, 0xb2, 0x9c, 0x16, 0x65, 0x19, 0x75, 0x8b, 0x2d, 0x1f,
	0xf6, 0xc5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x4e, 0xe7, 0xc5, 0x7e, 0x03, 0x00, 0x00,
}