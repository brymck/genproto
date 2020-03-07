// Code generated by protoc-gen-go. DO NOT EDIT.
// source: brymck/calendar/v1/calendar_api.proto

package calendarv1

import (
	context "context"
	fmt "fmt"
	v1 "github.com/brymck/genproto/dates/v1"
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

type GetDatesRequest struct {
	StartDate            *v1.Date `protobuf:"bytes,1,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate              *v1.Date `protobuf:"bytes,2,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDatesRequest) Reset()         { *m = GetDatesRequest{} }
func (m *GetDatesRequest) String() string { return proto.CompactTextString(m) }
func (*GetDatesRequest) ProtoMessage()    {}
func (*GetDatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_336f27c4c603af63, []int{0}
}

func (m *GetDatesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDatesRequest.Unmarshal(m, b)
}
func (m *GetDatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDatesRequest.Marshal(b, m, deterministic)
}
func (m *GetDatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDatesRequest.Merge(m, src)
}
func (m *GetDatesRequest) XXX_Size() int {
	return xxx_messageInfo_GetDatesRequest.Size(m)
}
func (m *GetDatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDatesRequest proto.InternalMessageInfo

func (m *GetDatesRequest) GetStartDate() *v1.Date {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *GetDatesRequest) GetEndDate() *v1.Date {
	if m != nil {
		return m.EndDate
	}
	return nil
}

type GetDatesResponse struct {
	Dates                []*v1.Date `protobuf:"bytes,1,rep,name=dates,proto3" json:"dates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetDatesResponse) Reset()         { *m = GetDatesResponse{} }
func (m *GetDatesResponse) String() string { return proto.CompactTextString(m) }
func (*GetDatesResponse) ProtoMessage()    {}
func (*GetDatesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_336f27c4c603af63, []int{1}
}

func (m *GetDatesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDatesResponse.Unmarshal(m, b)
}
func (m *GetDatesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDatesResponse.Marshal(b, m, deterministic)
}
func (m *GetDatesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDatesResponse.Merge(m, src)
}
func (m *GetDatesResponse) XXX_Size() int {
	return xxx_messageInfo_GetDatesResponse.Size(m)
}
func (m *GetDatesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDatesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDatesResponse proto.InternalMessageInfo

func (m *GetDatesResponse) GetDates() []*v1.Date {
	if m != nil {
		return m.Dates
	}
	return nil
}

type GetLatestBusinessDayRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLatestBusinessDayRequest) Reset()         { *m = GetLatestBusinessDayRequest{} }
func (m *GetLatestBusinessDayRequest) String() string { return proto.CompactTextString(m) }
func (*GetLatestBusinessDayRequest) ProtoMessage()    {}
func (*GetLatestBusinessDayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_336f27c4c603af63, []int{2}
}

func (m *GetLatestBusinessDayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLatestBusinessDayRequest.Unmarshal(m, b)
}
func (m *GetLatestBusinessDayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLatestBusinessDayRequest.Marshal(b, m, deterministic)
}
func (m *GetLatestBusinessDayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLatestBusinessDayRequest.Merge(m, src)
}
func (m *GetLatestBusinessDayRequest) XXX_Size() int {
	return xxx_messageInfo_GetLatestBusinessDayRequest.Size(m)
}
func (m *GetLatestBusinessDayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLatestBusinessDayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLatestBusinessDayRequest proto.InternalMessageInfo

type GetLatestBusinessDayResponse struct {
	Date                 *v1.Date `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLatestBusinessDayResponse) Reset()         { *m = GetLatestBusinessDayResponse{} }
func (m *GetLatestBusinessDayResponse) String() string { return proto.CompactTextString(m) }
func (*GetLatestBusinessDayResponse) ProtoMessage()    {}
func (*GetLatestBusinessDayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_336f27c4c603af63, []int{3}
}

func (m *GetLatestBusinessDayResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLatestBusinessDayResponse.Unmarshal(m, b)
}
func (m *GetLatestBusinessDayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLatestBusinessDayResponse.Marshal(b, m, deterministic)
}
func (m *GetLatestBusinessDayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLatestBusinessDayResponse.Merge(m, src)
}
func (m *GetLatestBusinessDayResponse) XXX_Size() int {
	return xxx_messageInfo_GetLatestBusinessDayResponse.Size(m)
}
func (m *GetLatestBusinessDayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLatestBusinessDayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLatestBusinessDayResponse proto.InternalMessageInfo

func (m *GetLatestBusinessDayResponse) GetDate() *v1.Date {
	if m != nil {
		return m.Date
	}
	return nil
}

func init() {
	proto.RegisterType((*GetDatesRequest)(nil), "brymck.calendar.v1.GetDatesRequest")
	proto.RegisterType((*GetDatesResponse)(nil), "brymck.calendar.v1.GetDatesResponse")
	proto.RegisterType((*GetLatestBusinessDayRequest)(nil), "brymck.calendar.v1.GetLatestBusinessDayRequest")
	proto.RegisterType((*GetLatestBusinessDayResponse)(nil), "brymck.calendar.v1.GetLatestBusinessDayResponse")
}

func init() {
	proto.RegisterFile("brymck/calendar/v1/calendar_api.proto", fileDescriptor_336f27c4c603af63)
}

var fileDescriptor_336f27c4c603af63 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0x4a, 0xf3, 0x30,
	0x14, 0xc0, 0xe9, 0xf6, 0x7d, 0x3a, 0xcf, 0x2e, 0x1c, 0x41, 0x61, 0x54, 0x07, 0xa3, 0x2a, 0x28,
	0x42, 0xb6, 0x4c, 0xef, 0xc5, 0x6e, 0x30, 0x06, 0x5e, 0x8c, 0x81, 0x43, 0x64, 0x30, 0xb2, 0x35,
	0x68, 0xd1, 0xb5, 0xb5, 0x49, 0x0b, 0x7d, 0x1d, 0x2f, 0x7d, 0x14, 0x5f, 0xc2, 0x57, 0x91, 0x34,
	0xc9, 0x04, 0x6d, 0xa7, 0x77, 0x87, 0x9c, 0xdf, 0xef, 0xfc, 0x23, 0x70, 0xb2, 0x88, 0xb3, 0xd5,
	0xf2, 0xa9, 0xb3, 0xa4, 0xcf, 0x2c, 0xf0, 0x68, 0xdc, 0x49, 0xc9, 0x3a, 0x9e, 0xd3, 0xc8, 0xc7,
	0x51, 0x1c, 0x8a, 0x10, 0x21, 0x85, 0x61, 0x93, 0xc2, 0x29, 0xb1, 0x6d, 0xad, 0x7a, 0x54, 0x30,
	0x2e, 0x3d, 0x19, 0x28, 0xde, 0xc9, 0x60, 0x77, 0xc8, 0xc4, 0x40, 0x66, 0x26, 0xec, 0x25, 0x61,
	0x5c, 0xa0, 0x4b, 0x00, 0x2e, 0x68, 0x2c, 0xe6, 0x12, 0x6b, 0x5a, 0x6d, 0xeb, 0xb4, 0xde, 0xdb,
	0xc7, 0xba, 0x6e, 0x5e, 0x03, 0xa7, 0x04, 0x4b, 0x65, 0xb2, 0x93, 0x83, 0x32, 0x44, 0x5d, 0xa8,
	0xb1, 0xc0, 0x53, 0x4e, 0x65, 0x93, 0xb3, 0xcd, 0x02, 0x4f, 0x06, 0xce, 0x15, 0x34, 0xbe, 0x5a,
	0xf3, 0x28, 0x0c, 0x38, 0x43, 0xe7, 0xf0, 0x3f, 0xa7, 0x9b, 0x56, 0xbb, 0x5a, 0x5e, 0x42, 0x31,
	0x4e, 0x0b, 0x0e, 0x86, 0x4c, 0xdc, 0xc8, 0x58, 0xb8, 0x09, 0xf7, 0x03, 0xc6, 0xf9, 0x80, 0x66,
	0x7a, 0x0f, 0x67, 0x04, 0x87, 0xc5, 0x69, 0xdd, 0xeb, 0x0c, 0xfe, 0xfd, 0xbe, 0x61, 0x8e, 0xf4,
	0x3e, 0x2c, 0xa8, 0xf7, 0xf5, 0x45, 0xaf, 0xc7, 0x23, 0x74, 0x0b, 0x35, 0x33, 0x3a, 0x3a, 0xc2,
	0x3f, 0x4f, 0x8e, 0xbf, 0xdd, 0xd4, 0x3e, 0xde, 0x0c, 0xe9, 0x89, 0x32, 0xd8, 0x2b, 0x9a, 0x18,
	0x75, 0x4a, 0xec, 0xb2, 0xd5, 0xed, 0xee, 0xdf, 0x05, 0xd5, 0xda, 0x0d, 0xa1, 0xb5, 0x0c, 0x57,
	0xf8, 0xc1, 0x17, 0x8f, 0xc9, 0xa2, 0xc0, 0x76, 0x1b, 0xeb, 0xfd, 0x23, 0x7f, 0x2c, 0xbf, 0xce,
	0xd8, 0xba, 0x07, 0x03, 0xa4, 0xe4, 0xb5, 0x52, 0x75, 0xfb, 0x77, 0x6f, 0x15, 0xe4, 0x2a, 0xd5,
	0xd0, 0x78, 0x4a, 0xde, 0xcd, 0xe3, 0xcc, 0x3c, 0xce, 0xa6, 0x64, 0xb1, 0x95, 0xff, 0xbf, 0x8b,
	0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x09, 0x53, 0xd0, 0x18, 0xd8, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalendarAPIClient is the client API for CalendarAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalendarAPIClient interface {
	// GetDates retrieves dates.
	GetDates(ctx context.Context, in *GetDatesRequest, opts ...grpc.CallOption) (*GetDatesResponse, error)
	// GetLatestBusinessDay retrieves the latest business day.
	GetLatestBusinessDay(ctx context.Context, in *GetLatestBusinessDayRequest, opts ...grpc.CallOption) (*GetLatestBusinessDayResponse, error)
}

type calendarAPIClient struct {
	cc *grpc.ClientConn
}

func NewCalendarAPIClient(cc *grpc.ClientConn) CalendarAPIClient {
	return &calendarAPIClient{cc}
}

func (c *calendarAPIClient) GetDates(ctx context.Context, in *GetDatesRequest, opts ...grpc.CallOption) (*GetDatesResponse, error) {
	out := new(GetDatesResponse)
	err := c.cc.Invoke(ctx, "/brymck.calendar.v1.CalendarAPI/GetDates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarAPIClient) GetLatestBusinessDay(ctx context.Context, in *GetLatestBusinessDayRequest, opts ...grpc.CallOption) (*GetLatestBusinessDayResponse, error) {
	out := new(GetLatestBusinessDayResponse)
	err := c.cc.Invoke(ctx, "/brymck.calendar.v1.CalendarAPI/GetLatestBusinessDay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarAPIServer is the server API for CalendarAPI service.
type CalendarAPIServer interface {
	// GetDates retrieves dates.
	GetDates(context.Context, *GetDatesRequest) (*GetDatesResponse, error)
	// GetLatestBusinessDay retrieves the latest business day.
	GetLatestBusinessDay(context.Context, *GetLatestBusinessDayRequest) (*GetLatestBusinessDayResponse, error)
}

func RegisterCalendarAPIServer(s *grpc.Server, srv CalendarAPIServer) {
	s.RegisterService(&_CalendarAPI_serviceDesc, srv)
}

func _CalendarAPI_GetDates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarAPIServer).GetDates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brymck.calendar.v1.CalendarAPI/GetDates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarAPIServer).GetDates(ctx, req.(*GetDatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarAPI_GetLatestBusinessDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestBusinessDayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarAPIServer).GetLatestBusinessDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brymck.calendar.v1.CalendarAPI/GetLatestBusinessDay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarAPIServer).GetLatestBusinessDay(ctx, req.(*GetLatestBusinessDayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalendarAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "brymck.calendar.v1.CalendarAPI",
	HandlerType: (*CalendarAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDates",
			Handler:    _CalendarAPI_GetDates_Handler,
		},
		{
			MethodName: "GetLatestBusinessDay",
			Handler:    _CalendarAPI_GetLatestBusinessDay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "brymck/calendar/v1/calendar_api.proto",
}
