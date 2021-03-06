// Code generated by protoc-gen-go. DO NOT EDIT.
// source: brymck/securities/v1/securities_api.proto

package securitiesv1

import (
	context "context"
	fmt "fmt"
	v1 "github.com/brymck/genproto/brymck/dates/v1"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

// A request for security information.
type GetSecurityRequest struct {
	// A unique identifier for a security.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The main symbol for the security per market convention.
	Symbol               string   `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSecurityRequest) Reset()         { *m = GetSecurityRequest{} }
func (m *GetSecurityRequest) String() string { return proto.CompactTextString(m) }
func (*GetSecurityRequest) ProtoMessage()    {}
func (*GetSecurityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_59e644c81e27fbff, []int{0}
}

func (m *GetSecurityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSecurityRequest.Unmarshal(m, b)
}
func (m *GetSecurityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSecurityRequest.Marshal(b, m, deterministic)
}
func (m *GetSecurityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSecurityRequest.Merge(m, src)
}
func (m *GetSecurityRequest) XXX_Size() int {
	return xxx_messageInfo_GetSecurityRequest.Size(m)
}
func (m *GetSecurityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSecurityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSecurityRequest proto.InternalMessageInfo

func (m *GetSecurityRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetSecurityRequest) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

type GetSecurityResponse struct {
	Security             *Security `protobuf:"bytes,1,opt,name=security,proto3" json:"security,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetSecurityResponse) Reset()         { *m = GetSecurityResponse{} }
func (m *GetSecurityResponse) String() string { return proto.CompactTextString(m) }
func (*GetSecurityResponse) ProtoMessage()    {}
func (*GetSecurityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_59e644c81e27fbff, []int{1}
}

func (m *GetSecurityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSecurityResponse.Unmarshal(m, b)
}
func (m *GetSecurityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSecurityResponse.Marshal(b, m, deterministic)
}
func (m *GetSecurityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSecurityResponse.Merge(m, src)
}
func (m *GetSecurityResponse) XXX_Size() int {
	return xxx_messageInfo_GetSecurityResponse.Size(m)
}
func (m *GetSecurityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSecurityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSecurityResponse proto.InternalMessageInfo

func (m *GetSecurityResponse) GetSecurity() *Security {
	if m != nil {
		return m.Security
	}
	return nil
}

type InsertSecurityRequest struct {
	Security             *Security `protobuf:"bytes,1,opt,name=security,proto3" json:"security,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *InsertSecurityRequest) Reset()         { *m = InsertSecurityRequest{} }
func (m *InsertSecurityRequest) String() string { return proto.CompactTextString(m) }
func (*InsertSecurityRequest) ProtoMessage()    {}
func (*InsertSecurityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_59e644c81e27fbff, []int{2}
}

func (m *InsertSecurityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertSecurityRequest.Unmarshal(m, b)
}
func (m *InsertSecurityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertSecurityRequest.Marshal(b, m, deterministic)
}
func (m *InsertSecurityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertSecurityRequest.Merge(m, src)
}
func (m *InsertSecurityRequest) XXX_Size() int {
	return xxx_messageInfo_InsertSecurityRequest.Size(m)
}
func (m *InsertSecurityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertSecurityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InsertSecurityRequest proto.InternalMessageInfo

func (m *InsertSecurityRequest) GetSecurity() *Security {
	if m != nil {
		return m.Security
	}
	return nil
}

type InsertSecurityResponse struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InsertSecurityResponse) Reset()         { *m = InsertSecurityResponse{} }
func (m *InsertSecurityResponse) String() string { return proto.CompactTextString(m) }
func (*InsertSecurityResponse) ProtoMessage()    {}
func (*InsertSecurityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_59e644c81e27fbff, []int{3}
}

func (m *InsertSecurityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertSecurityResponse.Unmarshal(m, b)
}
func (m *InsertSecurityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertSecurityResponse.Marshal(b, m, deterministic)
}
func (m *InsertSecurityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertSecurityResponse.Merge(m, src)
}
func (m *InsertSecurityResponse) XXX_Size() int {
	return xxx_messageInfo_InsertSecurityResponse.Size(m)
}
func (m *InsertSecurityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertSecurityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InsertSecurityResponse proto.InternalMessageInfo

func (m *InsertSecurityResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetPricesRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StartDate            *v1.Date `protobuf:"bytes,2,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate              *v1.Date `protobuf:"bytes,3,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPricesRequest) Reset()         { *m = GetPricesRequest{} }
func (m *GetPricesRequest) String() string { return proto.CompactTextString(m) }
func (*GetPricesRequest) ProtoMessage()    {}
func (*GetPricesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_59e644c81e27fbff, []int{4}
}

func (m *GetPricesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPricesRequest.Unmarshal(m, b)
}
func (m *GetPricesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPricesRequest.Marshal(b, m, deterministic)
}
func (m *GetPricesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPricesRequest.Merge(m, src)
}
func (m *GetPricesRequest) XXX_Size() int {
	return xxx_messageInfo_GetPricesRequest.Size(m)
}
func (m *GetPricesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPricesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPricesRequest proto.InternalMessageInfo

func (m *GetPricesRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetPricesRequest) GetStartDate() *v1.Date {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *GetPricesRequest) GetEndDate() *v1.Date {
	if m != nil {
		return m.EndDate
	}
	return nil
}

type GetPricesResponse struct {
	Prices               []*Price `protobuf:"bytes,1,rep,name=prices,proto3" json:"prices,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPricesResponse) Reset()         { *m = GetPricesResponse{} }
func (m *GetPricesResponse) String() string { return proto.CompactTextString(m) }
func (*GetPricesResponse) ProtoMessage()    {}
func (*GetPricesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_59e644c81e27fbff, []int{5}
}

func (m *GetPricesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPricesResponse.Unmarshal(m, b)
}
func (m *GetPricesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPricesResponse.Marshal(b, m, deterministic)
}
func (m *GetPricesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPricesResponse.Merge(m, src)
}
func (m *GetPricesResponse) XXX_Size() int {
	return xxx_messageInfo_GetPricesResponse.Size(m)
}
func (m *GetPricesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPricesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPricesResponse proto.InternalMessageInfo

func (m *GetPricesResponse) GetPrices() []*Price {
	if m != nil {
		return m.Prices
	}
	return nil
}

type UpdatePricesRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Symbol               string   `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Full                 bool     `protobuf:"varint,3,opt,name=full,proto3" json:"full,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePricesRequest) Reset()         { *m = UpdatePricesRequest{} }
func (m *UpdatePricesRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePricesRequest) ProtoMessage()    {}
func (*UpdatePricesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_59e644c81e27fbff, []int{6}
}

func (m *UpdatePricesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePricesRequest.Unmarshal(m, b)
}
func (m *UpdatePricesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePricesRequest.Marshal(b, m, deterministic)
}
func (m *UpdatePricesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePricesRequest.Merge(m, src)
}
func (m *UpdatePricesRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePricesRequest.Size(m)
}
func (m *UpdatePricesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePricesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePricesRequest proto.InternalMessageInfo

func (m *UpdatePricesRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdatePricesRequest) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *UpdatePricesRequest) GetFull() bool {
	if m != nil {
		return m.Full
	}
	return false
}

type UpdatePricesResponse struct {
	Count                uint64   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePricesResponse) Reset()         { *m = UpdatePricesResponse{} }
func (m *UpdatePricesResponse) String() string { return proto.CompactTextString(m) }
func (*UpdatePricesResponse) ProtoMessage()    {}
func (*UpdatePricesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_59e644c81e27fbff, []int{7}
}

func (m *UpdatePricesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePricesResponse.Unmarshal(m, b)
}
func (m *UpdatePricesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePricesResponse.Marshal(b, m, deterministic)
}
func (m *UpdatePricesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePricesResponse.Merge(m, src)
}
func (m *UpdatePricesResponse) XXX_Size() int {
	return xxx_messageInfo_UpdatePricesResponse.Size(m)
}
func (m *UpdatePricesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePricesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePricesResponse proto.InternalMessageInfo

func (m *UpdatePricesResponse) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*GetSecurityRequest)(nil), "brymck.securities.v1.GetSecurityRequest")
	proto.RegisterType((*GetSecurityResponse)(nil), "brymck.securities.v1.GetSecurityResponse")
	proto.RegisterType((*InsertSecurityRequest)(nil), "brymck.securities.v1.InsertSecurityRequest")
	proto.RegisterType((*InsertSecurityResponse)(nil), "brymck.securities.v1.InsertSecurityResponse")
	proto.RegisterType((*GetPricesRequest)(nil), "brymck.securities.v1.GetPricesRequest")
	proto.RegisterType((*GetPricesResponse)(nil), "brymck.securities.v1.GetPricesResponse")
	proto.RegisterType((*UpdatePricesRequest)(nil), "brymck.securities.v1.UpdatePricesRequest")
	proto.RegisterType((*UpdatePricesResponse)(nil), "brymck.securities.v1.UpdatePricesResponse")
}

func init() {
	proto.RegisterFile("brymck/securities/v1/securities_api.proto", fileDescriptor_59e644c81e27fbff)
}

var fileDescriptor_59e644c81e27fbff = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x25, 0xe9, 0x5a, 0xdb, 0xdb, 0x75, 0xd1, 0xb1, 0xbb, 0x94, 0x08, 0x5a, 0x22, 0x68, 0x56,
	0x97, 0xac, 0xe9, 0xfa, 0x24, 0xbe, 0x58, 0x84, 0x75, 0xdf, 0x6a, 0x8a, 0x8b, 0x48, 0x61, 0x69,
	0x92, 0x51, 0x87, 0x6d, 0x93, 0x98, 0x99, 0x14, 0xfa, 0x17, 0xfc, 0x19, 0x82, 0x2f, 0xfe, 0x14,
	0x7f, 0x95, 0xcc, 0x47, 0x9b, 0x64, 0x3b, 0xa9, 0x85, 0x7d, 0xbb, 0x33, 0x39, 0xe7, 0xdc, 0x73,
	0xef, 0x19, 0x02, 0xc7, 0x41, 0xb6, 0x9c, 0x87, 0xd7, 0xa7, 0x14, 0x87, 0x79, 0x46, 0x18, 0xc1,
	0xf4, 0x74, 0xe1, 0x95, 0x4e, 0x57, 0xd3, 0x94, 0xb8, 0x69, 0x96, 0xb0, 0x04, 0x75, 0x25, 0xd4,
	0x2d, 0x3e, 0xba, 0x0b, 0xcf, 0xb2, 0x94, 0x40, 0x34, 0x65, 0x92, 0xcb, 0x0b, 0xc9, 0xb0, 0xfa,
	0x5a, 0xf1, 0x34, 0x23, 0xe1, 0x0a, 0xf1, 0x74, 0x5b, 0xfb, 0xa5, 0x04, 0xd9, 0x6f, 0x01, 0x9d,
	0x63, 0x36, 0x56, 0x97, 0x3e, 0xfe, 0x91, 0x63, 0xca, 0xd0, 0x01, 0x98, 0x24, 0xea, 0x19, 0x7d,
	0xc3, 0xd9, 0xf3, 0x4d, 0x12, 0xa1, 0x23, 0x68, 0xd2, 0xe5, 0x3c, 0x48, 0x66, 0x3d, 0xb3, 0x6f,
	0x38, 0x6d, 0x5f, 0x9d, 0xec, 0x8f, 0xf0, 0xb0, 0xc2, 0xa6, 0x69, 0x12, 0x53, 0x8c, 0xde, 0x40,
	0x6b, 0xd5, 0x46, 0x88, 0x74, 0x06, 0x8f, 0x5d, 0xdd, 0x80, 0xee, 0x9a, 0xb9, 0xc6, 0xdb, 0x63,
	0x38, 0xbc, 0x88, 0x29, 0xce, 0x36, 0x3c, 0xdd, 0x46, 0xd4, 0x81, 0xa3, 0x9b, 0xa2, 0xca, 0xea,
	0x8d, 0x49, 0xed, 0x9f, 0x06, 0xdc, 0x3f, 0xc7, 0x6c, 0xc4, 0xf7, 0x48, 0xeb, 0xd6, 0xf1, 0x1a,
	0x80, 0xb2, 0x69, 0xc6, 0xae, 0x78, 0x1e, 0x62, 0x25, 0x9d, 0xc1, 0xe1, 0xca, 0x8c, 0x08, 0x8b,
	0xfb, 0x78, 0x3f, 0x65, 0xd8, 0x6f, 0x0b, 0x20, 0x2f, 0xd1, 0x2b, 0x68, 0xe1, 0x38, 0x92, 0x9c,
	0xc6, 0x36, 0xce, 0x5d, 0x1c, 0x47, 0xbc, 0xb0, 0x3f, 0xc0, 0x83, 0x92, 0x17, 0xe5, 0xf8, 0x0c,
	0x9a, 0x22, 0x65, 0xda, 0x33, 0xfa, 0x0d, 0xa7, 0x33, 0x78, 0xa4, 0xdf, 0x82, 0x60, 0xf9, 0x0a,
	0xca, 0x83, 0xfa, 0x94, 0xf2, 0x2e, 0xdb, 0x07, 0xab, 0xc9, 0x19, 0x21, 0xd8, 0xfb, 0x9a, 0xcf,
	0x66, 0xc2, 0x76, 0xcb, 0x17, 0xb5, 0x7d, 0x02, 0xdd, 0xaa, 0xa4, 0xf2, 0xd7, 0x85, 0x3b, 0x61,
	0x92, 0xc7, 0x4c, 0xc9, 0xca, 0xc3, 0xe0, 0x77, 0x03, 0xee, 0x8d, 0xd7, 0x06, 0xdf, 0x8d, 0x2e,
	0x50, 0x00, 0x9d, 0xd2, 0xdb, 0x41, 0x8e, 0x7e, 0x8c, 0xcd, 0xc7, 0x69, 0x1d, 0xef, 0x80, 0x54,
	0x5e, 0xae, 0xe1, 0xa0, 0x9a, 0x3b, 0x7a, 0xa9, 0x27, 0x6b, 0x9f, 0x9c, 0x75, 0xb2, 0x1b, 0x58,
	0x35, 0x9b, 0x40, 0x7b, 0x9d, 0x16, 0x7a, 0x56, 0x6b, 0xb2, 0x92, 0x80, 0xf5, 0xfc, 0xbf, 0x38,
	0xa5, 0x8e, 0x61, 0xbf, 0xbc, 0x6e, 0x54, 0xb3, 0x05, 0x4d, 0xca, 0xd6, 0x8b, 0x5d, 0xa0, 0xb2,
	0xcd, 0x70, 0x09, 0x4f, 0xc2, 0x64, 0xee, 0x7e, 0x23, 0xec, 0x7b, 0x1e, 0x68, 0x79, 0x43, 0x54,
	0xca, 0x31, 0x25, 0x23, 0xfe, 0x1b, 0x19, 0x19, 0x5f, 0xf6, 0x0b, 0xd0, 0xc2, 0xfb, 0x65, 0x36,
	0x86, 0xe3, 0xcf, 0x7f, 0xcc, 0xee, 0x50, 0x0a, 0x14, 0x0c, 0xf7, 0xd2, 0xfb, 0xbb, 0xba, 0x9e,
	0x14, 0xd7, 0x93, 0x4b, 0x2f, 0x68, 0x8a, 0x3f, 0xd2, 0xd9, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xaa, 0x39, 0x11, 0x62, 0x37, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SecuritiesAPIClient is the client API for SecuritiesAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecuritiesAPIClient interface {
	// Retrieve a full security object.
	GetSecurity(ctx context.Context, in *GetSecurityRequest, opts ...grpc.CallOption) (*GetSecurityResponse, error)
	// Insert a new security.
	InsertSecurity(ctx context.Context, in *InsertSecurityRequest, opts ...grpc.CallOption) (*InsertSecurityResponse, error)
	// Retrieve security prices.
	GetPrices(ctx context.Context, in *GetPricesRequest, opts ...grpc.CallOption) (*GetPricesResponse, error)
	// Update security prices.
	UpdatePrices(ctx context.Context, in *UpdatePricesRequest, opts ...grpc.CallOption) (*UpdatePricesResponse, error)
}

type securitiesAPIClient struct {
	cc *grpc.ClientConn
}

func NewSecuritiesAPIClient(cc *grpc.ClientConn) SecuritiesAPIClient {
	return &securitiesAPIClient{cc}
}

func (c *securitiesAPIClient) GetSecurity(ctx context.Context, in *GetSecurityRequest, opts ...grpc.CallOption) (*GetSecurityResponse, error) {
	out := new(GetSecurityResponse)
	err := c.cc.Invoke(ctx, "/brymck.securities.v1.SecuritiesAPI/GetSecurity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securitiesAPIClient) InsertSecurity(ctx context.Context, in *InsertSecurityRequest, opts ...grpc.CallOption) (*InsertSecurityResponse, error) {
	out := new(InsertSecurityResponse)
	err := c.cc.Invoke(ctx, "/brymck.securities.v1.SecuritiesAPI/InsertSecurity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securitiesAPIClient) GetPrices(ctx context.Context, in *GetPricesRequest, opts ...grpc.CallOption) (*GetPricesResponse, error) {
	out := new(GetPricesResponse)
	err := c.cc.Invoke(ctx, "/brymck.securities.v1.SecuritiesAPI/GetPrices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securitiesAPIClient) UpdatePrices(ctx context.Context, in *UpdatePricesRequest, opts ...grpc.CallOption) (*UpdatePricesResponse, error) {
	out := new(UpdatePricesResponse)
	err := c.cc.Invoke(ctx, "/brymck.securities.v1.SecuritiesAPI/UpdatePrices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecuritiesAPIServer is the server API for SecuritiesAPI service.
type SecuritiesAPIServer interface {
	// Retrieve a full security object.
	GetSecurity(context.Context, *GetSecurityRequest) (*GetSecurityResponse, error)
	// Insert a new security.
	InsertSecurity(context.Context, *InsertSecurityRequest) (*InsertSecurityResponse, error)
	// Retrieve security prices.
	GetPrices(context.Context, *GetPricesRequest) (*GetPricesResponse, error)
	// Update security prices.
	UpdatePrices(context.Context, *UpdatePricesRequest) (*UpdatePricesResponse, error)
}

func RegisterSecuritiesAPIServer(s *grpc.Server, srv SecuritiesAPIServer) {
	s.RegisterService(&_SecuritiesAPI_serviceDesc, srv)
}

func _SecuritiesAPI_GetSecurity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecurityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecuritiesAPIServer).GetSecurity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brymck.securities.v1.SecuritiesAPI/GetSecurity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecuritiesAPIServer).GetSecurity(ctx, req.(*GetSecurityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecuritiesAPI_InsertSecurity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertSecurityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecuritiesAPIServer).InsertSecurity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brymck.securities.v1.SecuritiesAPI/InsertSecurity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecuritiesAPIServer).InsertSecurity(ctx, req.(*InsertSecurityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecuritiesAPI_GetPrices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPricesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecuritiesAPIServer).GetPrices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brymck.securities.v1.SecuritiesAPI/GetPrices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecuritiesAPIServer).GetPrices(ctx, req.(*GetPricesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecuritiesAPI_UpdatePrices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePricesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecuritiesAPIServer).UpdatePrices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brymck.securities.v1.SecuritiesAPI/UpdatePrices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecuritiesAPIServer).UpdatePrices(ctx, req.(*UpdatePricesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecuritiesAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "brymck.securities.v1.SecuritiesAPI",
	HandlerType: (*SecuritiesAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSecurity",
			Handler:    _SecuritiesAPI_GetSecurity_Handler,
		},
		{
			MethodName: "InsertSecurity",
			Handler:    _SecuritiesAPI_InsertSecurity_Handler,
		},
		{
			MethodName: "GetPrices",
			Handler:    _SecuritiesAPI_GetPrices_Handler,
		},
		{
			MethodName: "UpdatePrices",
			Handler:    _SecuritiesAPI_UpdatePrices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "brymck/securities/v1/securities_api.proto",
}
