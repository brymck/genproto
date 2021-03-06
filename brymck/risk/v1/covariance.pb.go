// Code generated by protoc-gen-go. DO NOT EDIT.
// source: brymck/risk/v1/covariance.proto

package riskv1

import (
	fmt "fmt"
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

// A CovariancePair is a tuple consisting of two security IDs and their covariance.
type CovariancePair struct {
	SecurityId1          uint64   `protobuf:"varint,1,opt,name=security_id1,json=securityId1,proto3" json:"security_id1,omitempty"`
	SecurityId2          uint64   `protobuf:"varint,2,opt,name=security_id2,json=securityId2,proto3" json:"security_id2,omitempty"`
	Covariance           float64  `protobuf:"fixed64,3,opt,name=covariance,proto3" json:"covariance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CovariancePair) Reset()         { *m = CovariancePair{} }
func (m *CovariancePair) String() string { return proto.CompactTextString(m) }
func (*CovariancePair) ProtoMessage()    {}
func (*CovariancePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_dad231244d5f3f58, []int{0}
}

func (m *CovariancePair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CovariancePair.Unmarshal(m, b)
}
func (m *CovariancePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CovariancePair.Marshal(b, m, deterministic)
}
func (m *CovariancePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CovariancePair.Merge(m, src)
}
func (m *CovariancePair) XXX_Size() int {
	return xxx_messageInfo_CovariancePair.Size(m)
}
func (m *CovariancePair) XXX_DiscardUnknown() {
	xxx_messageInfo_CovariancePair.DiscardUnknown(m)
}

var xxx_messageInfo_CovariancePair proto.InternalMessageInfo

func (m *CovariancePair) GetSecurityId1() uint64 {
	if m != nil {
		return m.SecurityId1
	}
	return 0
}

func (m *CovariancePair) GetSecurityId2() uint64 {
	if m != nil {
		return m.SecurityId2
	}
	return 0
}

func (m *CovariancePair) GetCovariance() float64 {
	if m != nil {
		return m.Covariance
	}
	return 0
}

func init() {
	proto.RegisterType((*CovariancePair)(nil), "brymck.risk.v1.CovariancePair")
}

func init() { proto.RegisterFile("brymck/risk/v1/covariance.proto", fileDescriptor_dad231244d5f3f58) }

var fileDescriptor_dad231244d5f3f58 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x2a, 0xaa, 0xcc,
	0x4d, 0xce, 0xd6, 0x2f, 0xca, 0x2c, 0xce, 0xd6, 0x2f, 0x33, 0xd4, 0x4f, 0xce, 0x2f, 0x4b, 0x2c,
	0xca, 0x4c, 0xcc, 0x4b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x83, 0x28, 0xd0,
	0x03, 0x29, 0xd0, 0x2b, 0x33, 0x54, 0x2a, 0xe3, 0xe2, 0x73, 0x86, 0xab, 0x09, 0x48, 0xcc, 0x2c,
	0x12, 0x52, 0xe4, 0xe2, 0x29, 0x4e, 0x4d, 0x2e, 0x2d, 0xca, 0x2c, 0xa9, 0x8c, 0xcf, 0x4c, 0x31,
	0x94, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x09, 0xe2, 0x86, 0x89, 0x79, 0xa6, 0x18, 0xa2, 0x29, 0x31,
	0x92, 0x60, 0x42, 0x57, 0x62, 0x24, 0x24, 0xc7, 0xc5, 0x85, 0xb0, 0x5b, 0x82, 0x59, 0x81, 0x51,
	0x83, 0x31, 0x08, 0x49, 0xc4, 0x29, 0x8e, 0x4b, 0x32, 0x39, 0x3f, 0x57, 0x2f, 0x3d, 0xb3, 0x24,
	0xa3, 0x34, 0x49, 0x0f, 0xd5, 0x51, 0x4e, 0xfc, 0x48, 0x4e, 0x02, 0xb9, 0x3a, 0x80, 0x31, 0x8a,
	0x0d, 0x24, 0x57, 0x66, 0xb8, 0x88, 0x89, 0xd9, 0x29, 0x28, 0x62, 0x15, 0x13, 0x9f, 0x13, 0x44,
	0x47, 0x10, 0x48, 0x47, 0x98, 0xe1, 0x29, 0x98, 0x40, 0x0c, 0x48, 0x20, 0x26, 0xcc, 0x30, 0x89,
	0x0d, 0xec, 0x5d, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x19, 0xbb, 0x06, 0x87, 0x11, 0x01,
	0x00, 0x00,
}
