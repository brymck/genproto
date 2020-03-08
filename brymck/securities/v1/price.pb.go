// Code generated by protoc-gen-go. DO NOT EDIT.
// source: brymck/securities/v1/price.proto

package securitiesv1

import (
	fmt "fmt"
	v1 "github.com/brymck/genproto/brymck/dates/v1"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

// Price is a date and price.
type Price struct {
	Date                 *v1.Date `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Price                float64  `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Price) Reset()         { *m = Price{} }
func (m *Price) String() string { return proto.CompactTextString(m) }
func (*Price) ProtoMessage()    {}
func (*Price) Descriptor() ([]byte, []int) {
	return fileDescriptor_050da46d09d040e0, []int{0}
}

func (m *Price) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Price.Unmarshal(m, b)
}
func (m *Price) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Price.Marshal(b, m, deterministic)
}
func (m *Price) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Price.Merge(m, src)
}
func (m *Price) XXX_Size() int {
	return xxx_messageInfo_Price.Size(m)
}
func (m *Price) XXX_DiscardUnknown() {
	xxx_messageInfo_Price.DiscardUnknown(m)
}

var xxx_messageInfo_Price proto.InternalMessageInfo

func (m *Price) GetDate() *v1.Date {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *Price) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func init() {
	proto.RegisterType((*Price)(nil), "brymck.securities.v1.Price")
}

func init() { proto.RegisterFile("brymck/securities/v1/price.proto", fileDescriptor_050da46d09d040e0) }

var fileDescriptor_050da46d09d040e0 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0x2a, 0xaa, 0xcc,
	0x4d, 0xce, 0xd6, 0x2f, 0x4e, 0x4d, 0x2e, 0x2d, 0xca, 0x2c, 0xc9, 0x4c, 0x2d, 0xd6, 0x2f, 0x33,
	0xd4, 0x2f, 0x28, 0xca, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x81, 0xa8,
	0xd0, 0x43, 0xa8, 0xd0, 0x2b, 0x33, 0x94, 0x92, 0x82, 0xea, 0x4b, 0x49, 0x2c, 0x81, 0x68, 0x01,
	0x31, 0x20, 0x3a, 0x94, 0x3c, 0xb8, 0x58, 0x03, 0x40, 0x06, 0x08, 0x69, 0x72, 0xb1, 0x80, 0x84,
	0x25, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x44, 0xf5, 0xa0, 0x26, 0x81, 0xf5, 0xe8, 0x95, 0x19,
	0xea, 0xb9, 0x24, 0x96, 0xa4, 0x06, 0x81, 0x95, 0x08, 0x89, 0x70, 0xb1, 0x82, 0x2d, 0x95, 0x60,
	0x52, 0x60, 0xd4, 0x60, 0x0c, 0x82, 0x70, 0x9c, 0x0a, 0xb9, 0xe4, 0x93, 0xf3, 0x73, 0xf5, 0xd2,
	0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0xb0, 0x39, 0xc4, 0x89, 0x0b, 0x6c, 0x55, 0x00, 0xc8, 0xe2,
	0x00, 0xc6, 0x28, 0x1e, 0x84, 0x64, 0x99, 0xe1, 0x22, 0x26, 0x66, 0xa7, 0xe0, 0x88, 0x55, 0x4c,
	0x22, 0x4e, 0x10, 0x8d, 0xc1, 0x08, 0x8d, 0x61, 0x86, 0xa7, 0x60, 0xc2, 0x31, 0x08, 0xe1, 0x98,
	0x30, 0xc3, 0x24, 0x36, 0xb0, 0x1f, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x85, 0x3f, 0x2c,
	0x2c, 0x19, 0x01, 0x00, 0x00,
}
