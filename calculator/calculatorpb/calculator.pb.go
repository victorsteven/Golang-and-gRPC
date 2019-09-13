// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calculatorpb/calculator.proto

package calculatorpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SumRequest struct {
	FirstNumber          int32    `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber         int32    `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{0}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetFirstNumber() int32 {
	if m != nil {
		return m.FirstNumber
	}
	return 0
}

func (m *SumRequest) GetSecondNumber() int32 {
	if m != nil {
		return m.SecondNumber
	}
	return 0
}

type SumResponse struct {
	SumResult            int32    `protobuf:"varint,1,opt,name=sum_result,json=sumResult,proto3" json:"sum_result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{1}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetSumResult() int32 {
	if m != nil {
		return m.SumResult
	}
	return 0
}

type PrimeRequest struct {
	TheNumber            int32    `protobuf:"varint,1,opt,name=the_number,json=theNumber,proto3" json:"the_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeRequest) Reset()         { *m = PrimeRequest{} }
func (m *PrimeRequest) String() string { return proto.CompactTextString(m) }
func (*PrimeRequest) ProtoMessage()    {}
func (*PrimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{2}
}

func (m *PrimeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeRequest.Unmarshal(m, b)
}
func (m *PrimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeRequest.Marshal(b, m, deterministic)
}
func (m *PrimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeRequest.Merge(m, src)
}
func (m *PrimeRequest) XXX_Size() int {
	return xxx_messageInfo_PrimeRequest.Size(m)
}
func (m *PrimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeRequest proto.InternalMessageInfo

func (m *PrimeRequest) GetTheNumber() int32 {
	if m != nil {
		return m.TheNumber
	}
	return 0
}

type PrimeResponse struct {
	PrimeFactor          int32    `protobuf:"varint,1,opt,name=prime_factor,json=primeFactor,proto3" json:"prime_factor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeResponse) Reset()         { *m = PrimeResponse{} }
func (m *PrimeResponse) String() string { return proto.CompactTextString(m) }
func (*PrimeResponse) ProtoMessage()    {}
func (*PrimeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{3}
}

func (m *PrimeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeResponse.Unmarshal(m, b)
}
func (m *PrimeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeResponse.Marshal(b, m, deterministic)
}
func (m *PrimeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeResponse.Merge(m, src)
}
func (m *PrimeResponse) XXX_Size() int {
	return xxx_messageInfo_PrimeResponse.Size(m)
}
func (m *PrimeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeResponse proto.InternalMessageInfo

func (m *PrimeResponse) GetPrimeFactor() int32 {
	if m != nil {
		return m.PrimeFactor
	}
	return 0
}

type NumberRequest struct {
	Number               int64    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NumberRequest) Reset()         { *m = NumberRequest{} }
func (m *NumberRequest) String() string { return proto.CompactTextString(m) }
func (*NumberRequest) ProtoMessage()    {}
func (*NumberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{4}
}

func (m *NumberRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NumberRequest.Unmarshal(m, b)
}
func (m *NumberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NumberRequest.Marshal(b, m, deterministic)
}
func (m *NumberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumberRequest.Merge(m, src)
}
func (m *NumberRequest) XXX_Size() int {
	return xxx_messageInfo_NumberRequest.Size(m)
}
func (m *NumberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NumberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NumberRequest proto.InternalMessageInfo

func (m *NumberRequest) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type NumberResponse struct {
	Average              float64  `protobuf:"fixed64,1,opt,name=average,proto3" json:"average,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NumberResponse) Reset()         { *m = NumberResponse{} }
func (m *NumberResponse) String() string { return proto.CompactTextString(m) }
func (*NumberResponse) ProtoMessage()    {}
func (*NumberResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{5}
}

func (m *NumberResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NumberResponse.Unmarshal(m, b)
}
func (m *NumberResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NumberResponse.Marshal(b, m, deterministic)
}
func (m *NumberResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumberResponse.Merge(m, src)
}
func (m *NumberResponse) XXX_Size() int {
	return xxx_messageInfo_NumberResponse.Size(m)
}
func (m *NumberResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NumberResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NumberResponse proto.InternalMessageInfo

func (m *NumberResponse) GetAverage() float64 {
	if m != nil {
		return m.Average
	}
	return 0
}

type FindMaximumRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaximumRequest) Reset()         { *m = FindMaximumRequest{} }
func (m *FindMaximumRequest) String() string { return proto.CompactTextString(m) }
func (*FindMaximumRequest) ProtoMessage()    {}
func (*FindMaximumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{6}
}

func (m *FindMaximumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaximumRequest.Unmarshal(m, b)
}
func (m *FindMaximumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaximumRequest.Marshal(b, m, deterministic)
}
func (m *FindMaximumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaximumRequest.Merge(m, src)
}
func (m *FindMaximumRequest) XXX_Size() int {
	return xxx_messageInfo_FindMaximumRequest.Size(m)
}
func (m *FindMaximumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaximumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaximumRequest proto.InternalMessageInfo

func (m *FindMaximumRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type FindMaximumResponse struct {
	Maximum              int32    `protobuf:"varint,1,opt,name=maximum,proto3" json:"maximum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaximumResponse) Reset()         { *m = FindMaximumResponse{} }
func (m *FindMaximumResponse) String() string { return proto.CompactTextString(m) }
func (*FindMaximumResponse) ProtoMessage()    {}
func (*FindMaximumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{7}
}

func (m *FindMaximumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaximumResponse.Unmarshal(m, b)
}
func (m *FindMaximumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaximumResponse.Marshal(b, m, deterministic)
}
func (m *FindMaximumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaximumResponse.Merge(m, src)
}
func (m *FindMaximumResponse) XXX_Size() int {
	return xxx_messageInfo_FindMaximumResponse.Size(m)
}
func (m *FindMaximumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaximumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaximumResponse proto.InternalMessageInfo

func (m *FindMaximumResponse) GetMaximum() int32 {
	if m != nil {
		return m.Maximum
	}
	return 0
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "calculator.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "calculator.SumResponse")
	proto.RegisterType((*PrimeRequest)(nil), "calculator.PrimeRequest")
	proto.RegisterType((*PrimeResponse)(nil), "calculator.PrimeResponse")
	proto.RegisterType((*NumberRequest)(nil), "calculator.NumberRequest")
	proto.RegisterType((*NumberResponse)(nil), "calculator.NumberResponse")
	proto.RegisterType((*FindMaximumRequest)(nil), "calculator.FindMaximumRequest")
	proto.RegisterType((*FindMaximumResponse)(nil), "calculator.FindMaximumResponse")
}

func init() {
	proto.RegisterFile("calculator/calculatorpb/calculator.proto", fileDescriptor_7f42938f8c8365cf)
}

var fileDescriptor_7f42938f8c8365cf = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x4f, 0x4f, 0xf2, 0x40,
	0x10, 0xc6, 0xdf, 0x42, 0xe0, 0x0d, 0xd3, 0x42, 0xe2, 0x9a, 0x20, 0x36, 0x41, 0xa5, 0x1e, 0x24,
	0x06, 0x81, 0xe0, 0xc5, 0xab, 0x92, 0x70, 0x31, 0x1a, 0x53, 0x38, 0x79, 0x21, 0xa5, 0x2c, 0xd2,
	0xa4, 0xdb, 0xd6, 0xfd, 0x43, 0xfc, 0x68, 0x7e, 0x3c, 0xe3, 0x76, 0x97, 0x6e, 0x03, 0xde, 0x3a,
	0xcf, 0x3c, 0xfb, 0xdb, 0xc9, 0x33, 0x5d, 0xe8, 0x87, 0x41, 0x1c, 0x8a, 0x38, 0xe0, 0x29, 0x1d,
	0x15, 0x9f, 0xd9, 0xca, 0x28, 0x86, 0x19, 0x4d, 0x79, 0x8a, 0xa0, 0x50, 0xbc, 0x05, 0xc0, 0x5c,
	0x10, 0x1f, 0x7f, 0x0a, 0xcc, 0x38, 0xea, 0x81, 0xb3, 0x89, 0x28, 0xe3, 0xcb, 0x44, 0x90, 0x15,
	0xa6, 0x1d, 0xeb, 0xca, 0xea, 0xd7, 0x7c, 0x5b, 0x6a, 0xaf, 0x52, 0x42, 0xd7, 0xd0, 0x64, 0x38,
	0x4c, 0x93, 0xb5, 0xf6, 0x54, 0xa4, 0xc7, 0xc9, 0xc5, 0xdc, 0xe4, 0x0d, 0xc0, 0x96, 0x54, 0x96,
	0xa5, 0x09, 0xc3, 0xa8, 0x0b, 0xc0, 0x04, 0x59, 0x52, 0xcc, 0x44, 0xcc, 0x15, 0xb4, 0xc1, 0xa4,
	0x41, 0xc4, 0xdc, 0xbb, 0x03, 0xe7, 0x8d, 0x46, 0x04, 0xeb, 0x29, 0xba, 0x00, 0x7c, 0x8b, 0xcb,
	0x33, 0x34, 0xf8, 0x16, 0x2b, 0xf8, 0x04, 0x9a, 0xca, 0xae, 0xf0, 0x3d, 0x70, 0xb2, 0x5f, 0x61,
	0xb9, 0x09, 0x42, 0x9e, 0xee, 0xa7, 0x96, 0xda, 0x4c, 0x4a, 0xde, 0x0d, 0x34, 0xf3, 0xd3, 0xfa,
	0x8e, 0x36, 0xd4, 0x0d, 0x7e, 0xd5, 0x57, 0x95, 0x77, 0x0b, 0x2d, 0x6d, 0x54, 0xf4, 0x0e, 0xfc,
	0x0f, 0x76, 0x98, 0x06, 0x1f, 0x58, 0x5a, 0x2d, 0x5f, 0x97, 0xde, 0x00, 0xd0, 0x2c, 0x4a, 0xd6,
	0x2f, 0xc1, 0x57, 0x44, 0x8a, 0x0c, 0xcb, 0xe4, 0xda, 0x9e, 0x3c, 0x82, 0xd3, 0x92, 0xbb, 0xc0,
	0x93, 0x5c, 0x52, 0x7e, 0x5d, 0x4e, 0xbe, 0x2b, 0x70, 0x32, 0xdd, 0x6f, 0x6a, 0x8e, 0xe9, 0x2e,
	0x0a, 0x31, 0x7a, 0x80, 0xea, 0x5c, 0x10, 0xd4, 0x1e, 0x1a, 0x6b, 0x2d, 0x36, 0xe8, 0x9e, 0x1d,
	0xe8, 0xf9, 0x3d, 0xde, 0x3f, 0x34, 0x03, 0x5b, 0xe6, 0xa6, 0x16, 0xd9, 0x31, 0x9d, 0x66, 0xfe,
	0xee, 0xf9, 0x91, 0x8e, 0xa6, 0x8c, 0x2d, 0xf4, 0x0c, 0xad, 0x69, 0x4a, 0x32, 0xc1, 0xf1, 0x63,
	0x1e, 0x04, 0x2a, 0x1d, 0x28, 0xe5, 0xec, 0xba, 0xc7, 0x5a, 0x1a, 0xd6, 0xb7, 0xd0, 0x02, 0x6c,
	0x23, 0x15, 0x74, 0x61, 0xda, 0x0f, 0xc3, 0x75, 0x2f, 0xff, 0xec, 0x17, 0xcc, 0xb1, 0xf5, 0xd4,
	0x7a, 0x77, 0xcc, 0x27, 0xb0, 0xaa, 0xcb, 0x1f, 0xff, 0xfe, 0x27, 0x00, 0x00, 0xff, 0xff, 0xee,
	0x52, 0xac, 0xca, 0x24, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	PrimeNumber(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberClient, error)
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error)
	FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_FindMaximumClient, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) PrimeNumber(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculator.CalculatorService/PrimeNumber", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeNumberClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeNumberClient interface {
	Recv() (*PrimeResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeNumberClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeNumberClient) Recv() (*PrimeResponse, error) {
	m := new(PrimeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[1], "/calculator.CalculatorService/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceComputeAverageClient{stream}
	return x, nil
}

type CalculatorService_ComputeAverageClient interface {
	Send(*NumberRequest) error
	CloseAndRecv() (*NumberResponse, error)
	grpc.ClientStream
}

type calculatorServiceComputeAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceComputeAverageClient) Send(m *NumberRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceComputeAverageClient) CloseAndRecv() (*NumberResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(NumberResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_FindMaximumClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[2], "/calculator.CalculatorService/FindMaximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceFindMaximumClient{stream}
	return x, nil
}

type CalculatorService_FindMaximumClient interface {
	Send(*FindMaximumRequest) error
	Recv() (*FindMaximumResponse, error)
	grpc.ClientStream
}

type calculatorServiceFindMaximumClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceFindMaximumClient) Send(m *FindMaximumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceFindMaximumClient) Recv() (*FindMaximumResponse, error) {
	m := new(FindMaximumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	PrimeNumber(*PrimeRequest, CalculatorService_PrimeNumberServer) error
	ComputeAverage(CalculatorService_ComputeAverageServer) error
	FindMaximum(CalculatorService_FindMaximumServer) error
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Sum(ctx context.Context, req *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedCalculatorServiceServer) PrimeNumber(req *PrimeRequest, srv CalculatorService_PrimeNumberServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeNumber not implemented")
}
func (*UnimplementedCalculatorServiceServer) ComputeAverage(srv CalculatorService_ComputeAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAverage not implemented")
}
func (*UnimplementedCalculatorServiceServer) FindMaximum(srv CalculatorService_FindMaximumServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMaximum not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_PrimeNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).PrimeNumber(m, &calculatorServicePrimeNumberServer{stream})
}

type CalculatorService_PrimeNumberServer interface {
	Send(*PrimeResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeNumberServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeNumberServer) Send(m *PrimeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).ComputeAverage(&calculatorServiceComputeAverageServer{stream})
}

type CalculatorService_ComputeAverageServer interface {
	SendAndClose(*NumberResponse) error
	Recv() (*NumberRequest, error)
	grpc.ServerStream
}

type calculatorServiceComputeAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceComputeAverageServer) SendAndClose(m *NumberResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceComputeAverageServer) Recv() (*NumberRequest, error) {
	m := new(NumberRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_FindMaximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).FindMaximum(&calculatorServiceFindMaximumServer{stream})
}

type CalculatorService_FindMaximumServer interface {
	Send(*FindMaximumResponse) error
	Recv() (*FindMaximumRequest, error)
	grpc.ServerStream
}

type calculatorServiceFindMaximumServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceFindMaximumServer) Send(m *FindMaximumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceFindMaximumServer) Recv() (*FindMaximumRequest, error) {
	m := new(FindMaximumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumber",
			Handler:       _CalculatorService_PrimeNumber_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAverage",
			Handler:       _CalculatorService_ComputeAverage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMaximum",
			Handler:       _CalculatorService_FindMaximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calculatorpb/calculator.proto",
}
