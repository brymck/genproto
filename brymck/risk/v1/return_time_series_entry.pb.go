// Code generated by protoc-gen-go. DO NOT EDIT.
// source: brymck/risk/v1/return_time_series_entry.proto

package riskv1

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

// A ReturnTimeSeriesEntry contains a date and relevant return metrics.
type ReturnTimeSeriesEntry struct {
	Date                 *v1.Date `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Return               float64  `protobuf:"fixed64,2,opt,name=return,proto3" json:"return,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReturnTimeSeriesEntry) Reset()         { *m = ReturnTimeSeriesEntry{} }
func (m *ReturnTimeSeriesEntry) String() string { return proto.CompactTextString(m) }
func (*ReturnTimeSeriesEntry) ProtoMessage()    {}
func (*ReturnTimeSeriesEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_d191b92cb1a59e36, []int{0}
}

func (m *ReturnTimeSeriesEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReturnTimeSeriesEntry.Unmarshal(m, b)
}
func (m *ReturnTimeSeriesEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReturnTimeSeriesEntry.Marshal(b, m, deterministic)
}
func (m *ReturnTimeSeriesEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReturnTimeSeriesEntry.Merge(m, src)
}
func (m *ReturnTimeSeriesEntry) XXX_Size() int {
	return xxx_messageInfo_ReturnTimeSeriesEntry.Size(m)
}
func (m *ReturnTimeSeriesEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ReturnTimeSeriesEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ReturnTimeSeriesEntry proto.InternalMessageInfo

func (m *ReturnTimeSeriesEntry) GetDate() *v1.Date {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *ReturnTimeSeriesEntry) GetReturn() float64 {
	if m != nil {
		return m.Return
	}
	return 0
}

func init() {
	proto.RegisterType((*ReturnTimeSeriesEntry)(nil), "brymck.risk.v1.ReturnTimeSeriesEntry")
}

func init() {
	proto.RegisterFile("brymck/risk/v1/return_time_series_entry.proto", fileDescriptor_d191b92cb1a59e36)
}

var fileDescriptor_d191b92cb1a59e36 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4d, 0x2a, 0xaa, 0xcc,
	0x4d, 0xce, 0xd6, 0x2f, 0xca, 0x2c, 0xce, 0xd6, 0x2f, 0x33, 0xd4, 0x2f, 0x4a, 0x2d, 0x29, 0x2d,
	0xca, 0x8b, 0x2f, 0xc9, 0xcc, 0x4d, 0x8d, 0x2f, 0x4e, 0x2d, 0xca, 0x4c, 0x2d, 0x8e, 0x4f, 0xcd,
	0x2b, 0x29, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x83, 0x28, 0xd7, 0x03, 0x29,
	0xd7, 0x2b, 0x33, 0x94, 0x92, 0x82, 0x6a, 0x4f, 0x49, 0x2c, 0x49, 0x2d, 0x06, 0xe9, 0x07, 0x31,
	0x20, 0x6a, 0x95, 0xa2, 0xb8, 0x44, 0x83, 0xc0, 0xa6, 0x85, 0x64, 0xe6, 0xa6, 0x06, 0x83, 0xcd,
	0x72, 0x05, 0x19, 0x25, 0xa4, 0xc9, 0xc5, 0x02, 0x52, 0x26, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d,
	0x24, 0xaa, 0x07, 0x35, 0x13, 0x6c, 0x86, 0x5e, 0x99, 0xa1, 0x9e, 0x4b, 0x62, 0x49, 0x6a, 0x10,
	0x58, 0x89, 0x90, 0x18, 0x17, 0x1b, 0xc4, 0x45, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x8c, 0x41, 0x50,
	0x9e, 0x53, 0x26, 0x97, 0x64, 0x72, 0x7e, 0xae, 0x5e, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x1e,
	0xaa, 0xa3, 0x9c, 0xa4, 0xb0, 0x5a, 0x1b, 0x00, 0x72, 0x54, 0x00, 0x63, 0x14, 0x1b, 0x48, 0x59,
	0x99, 0xe1, 0x22, 0x26, 0x66, 0xa7, 0xa0, 0x88, 0x55, 0x4c, 0x7c, 0x4e, 0x10, 0xcd, 0x41, 0x20,
	0xcd, 0x61, 0x86, 0xa7, 0x60, 0x02, 0x31, 0x20, 0x81, 0x98, 0x30, 0xc3, 0x24, 0x36, 0xb0, 0x6f,
	0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xa8, 0xff, 0x44, 0x2a, 0x01, 0x00, 0x00,
}
