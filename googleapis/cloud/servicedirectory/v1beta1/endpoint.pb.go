// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/servicedirectory/v1beta1/endpoint.proto

package servicedirectory

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

// An individual endpoint that provides a
// [service][google.cloud.servicedirectory.v1beta1.Service]. The service must
// already exist to create an endpoint.
type Endpoint struct {
	// Immutable. The resource name for the endpoint in the format
	// 'projects/*/locations/*/namespaces/*/services/*/endpoints/*'.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. An IPv4 or IPv6 address. Service Directory will reject bad
	// addresses like:
	//   "8.8.8"
	//   "8.8.8.8:53"
	//   "test:bad:address"
	//   "[::1]"
	//   "[::1]:8080"
	// Limited to 45 characters.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// Optional. Service Directory will reject values outside of [0, 65535].
	Port int32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	// Optional. Metadata for the endpoint. This data can be consumed by service
	// clients.  The entire metadata dictionary may contain up to 512 characters,
	// spread accoss all key-value pairs. Metadata that goes beyond any these
	// limits will be rejected.
	Metadata             map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_b682dd88b410ee12, []int{0}
}

func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (m *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(m, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Endpoint) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Endpoint) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Endpoint) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "google.cloud.servicedirectory.v1beta1.Endpoint")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.servicedirectory.v1beta1.Endpoint.MetadataEntry")
}

func init() {
	proto.RegisterFile("google/cloud/servicedirectory/v1beta1/endpoint.proto", fileDescriptor_b682dd88b410ee12)
}

var fileDescriptor_b682dd88b410ee12 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x8a, 0xd4, 0x40,
	0x10, 0xa5, 0x93, 0x1d, 0x5d, 0x5b, 0x16, 0x24, 0x08, 0x8e, 0x83, 0xe2, 0x20, 0x2c, 0xc4, 0x4b,
	0x37, 0xa3, 0x1e, 0x24, 0x8b, 0x87, 0x1d, 0x5d, 0x3c, 0x09, 0xcb, 0x88, 0x7b, 0x90, 0x41, 0xe9,
	0xe9, 0xd4, 0xc6, 0x68, 0xa6, 0x2b, 0x74, 0xf7, 0x0c, 0x0c, 0x21, 0x1f, 0xb1, 0x1e, 0xfc, 0x08,
	0x3f, 0xc5, 0xaf, 0xf0, 0xbc, 0x5f, 0xe0, 0x51, 0x92, 0xee, 0xc4, 0x75, 0x3c, 0xec, 0xdc, 0x5e,
	0xea, 0xbd, 0x57, 0x79, 0x55, 0x5d, 0xf4, 0x79, 0x86, 0x98, 0x15, 0xc0, 0x65, 0x81, 0xab, 0x94,
	0x1b, 0xd0, 0xeb, 0x5c, 0x42, 0x9a, 0x6b, 0x90, 0x16, 0xf5, 0x86, 0xaf, 0x27, 0x0b, 0xb0, 0x62,
	0xc2, 0x41, 0xa5, 0x25, 0xe6, 0xca, 0xb2, 0x52, 0xa3, 0xc5, 0xe8, 0xd0, 0xb9, 0x58, 0xeb, 0x62,
	0xdb, 0x2e, 0xe6, 0x5d, 0xa3, 0x47, 0xbe, 0xb9, 0x28, 0x73, 0x7e, 0x9e, 0x43, 0x91, 0x7e, 0x5a,
	0xc0, 0x67, 0xb1, 0xce, 0x51, 0xbb, 0x3e, 0xa3, 0xfb, 0x57, 0x04, 0x1a, 0x0c, 0xae, 0xb4, 0x04,
	0x4f, 0x3d, 0xb8, 0x42, 0x09, 0xa5, 0xd0, 0x0a, 0x9b, 0xa3, 0x32, 0x8e, 0x7d, 0x7c, 0x11, 0xd2,
	0xfd, 0x13, 0x9f, 0x29, 0xba, 0x47, 0xf7, 0x94, 0x58, 0xc2, 0x90, 0x8c, 0x49, 0x7c, 0x6b, 0x1a,
	0xfe, 0x3a, 0x1e, 0xcc, 0xda, 0x42, 0xf4, 0x90, 0xde, 0x14, 0x69, 0xaa, 0xc1, 0x98, 0x61, 0xd0,
	0x71, 0x64, 0xd6, 0xd5, 0x1a, 0x5f, 0x89, 0xda, 0x0e, 0xc3, 0x31, 0x89, 0x07, 0x8e, 0x6b, 0x0b,
	0xd1, 0x47, 0xba, 0xbf, 0x04, 0x2b, 0x52, 0x61, 0xc5, 0x70, 0x6f, 0x1c, 0xc6, 0xb7, 0x9f, 0xbe,
	0x64, 0x3b, 0x4d, 0xcc, 0xba, 0x4c, 0xec, 0xad, 0xf7, 0x9f, 0x28, 0xab, 0x37, 0xae, 0x77, 0xdf,
	0x73, 0x74, 0x44, 0x0f, 0xfe, 0xe1, 0xa3, 0x3b, 0x34, 0xfc, 0x0a, 0x1b, 0x37, 0xc0, 0xac, 0x81,
	0xd1, 0x5d, 0x3a, 0x58, 0x8b, 0x62, 0x05, 0x2e, 0xf8, 0xcc, 0x7d, 0x24, 0xc1, 0x0b, 0x92, 0x7c,
	0x27, 0x97, 0xc7, 0xdf, 0x08, 0x8d, 0xff, 0xcb, 0xe0, 0x12, 0x8a, 0x32, 0x37, 0x4c, 0xe2, 0x92,
	0xf7, 0xeb, 0x39, 0x2f, 0x35, 0x7e, 0x01, 0x69, 0x0d, 0xaf, 0x3c, 0xaa, 0x79, 0x81, 0xd2, 0xed,
	0x93, 0x57, 0x1d, 0xac, 0x79, 0xb3, 0x33, 0x53, 0x0a, 0x09, 0x86, 0x57, 0x3d, 0xae, 0xbb, 0xa3,
	0x30, 0xbc, 0xf2, 0xa8, 0xee, 0xcf, 0xc1, 0xf0, 0xaa, 0x83, 0xf5, 0xf4, 0x22, 0xa0, 0x4f, 0x24,
	0x2e, 0x77, 0xdb, 0xd4, 0xf4, 0xa0, 0xcb, 0x77, 0xda, 0x3c, 0xe8, 0x29, 0xf9, 0xf0, 0xde, 0xfb,
	0x32, 0x2c, 0x84, 0xca, 0x18, 0xea, 0x8c, 0x67, 0xa0, 0xda, 0xe7, 0xe6, 0x7f, 0x47, 0xbb, 0xe6,
	0x50, 0x8f, 0xb6, 0x89, 0xdf, 0x84, 0xfc, 0x08, 0x0e, 0xdf, 0xb8, 0xd6, 0xaf, 0xda, 0x48, 0xef,
	0x9c, 0xe4, 0x75, 0x1f, 0xe9, 0x6c, 0x32, 0x6d, 0xbc, 0x3f, 0x3b, 0xdd, 0xbc, 0xd5, 0xcd, 0xb7,
	0x75, 0xf3, 0x33, 0xf7, 0x8f, 0xcb, 0x20, 0x76, 0xba, 0x24, 0x69, 0x85, 0x49, 0xb2, 0xad, 0x4c,
	0x12, 0x2f, 0x5d, 0xdc, 0x68, 0xf3, 0x3f, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x32, 0x5e,
	0xe6, 0x67, 0x03, 0x00, 0x00,
}