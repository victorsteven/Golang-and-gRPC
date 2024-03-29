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

type SquareRootRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquareRootRequest) Reset()         { *m = SquareRootRequest{} }
func (m *SquareRootRequest) String() string { return proto.CompactTextString(m) }
func (*SquareRootRequest) ProtoMessage()    {}
func (*SquareRootRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{8}
}

func (m *SquareRootRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquareRootRequest.Unmarshal(m, b)
}
func (m *SquareRootRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquareRootRequest.Marshal(b, m, deterministic)
}
func (m *SquareRootRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquareRootRequest.Merge(m, src)
}
func (m *SquareRootRequest) XXX_Size() int {
	return xxx_messageInfo_SquareRootRequest.Size(m)
}
func (m *SquareRootRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SquareRootRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SquareRootRequest proto.InternalMessageInfo

func (m *SquareRootRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type SquareRootResponse struct {
	NumberRoot           float64  `protobuf:"fixed64,1,opt,name=number_root,json=numberRoot,proto3" json:"number_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquareRootResponse) Reset()         { *m = SquareRootResponse{} }
func (m *SquareRootResponse) String() string { return proto.CompactTextString(m) }
func (*SquareRootResponse) ProtoMessage()    {}
func (*SquareRootResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{9}
}

func (m *SquareRootResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquareRootResponse.Unmarshal(m, b)
}
func (m *SquareRootResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquareRootResponse.Marshal(b, m, deterministic)
}
func (m *SquareRootResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquareRootResponse.Merge(m, src)
}
func (m *SquareRootResponse) XXX_Size() int {
	return xxx_messageInfo_SquareRootResponse.Size(m)
}
func (m *SquareRootResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SquareRootResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SquareRootResponse proto.InternalMessageInfo

func (m *SquareRootResponse) GetNumberRoot() float64 {
	if m != nil {
		return m.NumberRoot
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
	proto.RegisterType((*SquareRootRequest)(nil), "calculator.SquareRootRequest")
	proto.RegisterType((*SquareRootResponse)(nil), "calculator.SquareRootResponse")
}

func init() {
	proto.RegisterFile("calculator/calculatorpb/calculator.proto", fileDescriptor_7f42938f8c8365cf)
}

var fileDescriptor_7f42938f8c8365cf = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x86, 0x31, 0x51, 0x8b, 0x3a, 0x76, 0x22, 0x75, 0x91, 0x4a, 0xb0, 0x94, 0x96, 0x2e, 0x07,
	0x22, 0x28, 0x6d, 0x55, 0x84, 0xc4, 0x15, 0x2a, 0xe5, 0x82, 0x82, 0x90, 0x9d, 0x13, 0x17, 0xcb,
	0x71, 0x36, 0xc4, 0x92, 0xd7, 0xeb, 0xec, 0x47, 0xc4, 0x3f, 0xe0, 0x6f, 0xa3, 0xec, 0x47, 0xbc,
	0x26, 0x89, 0x7a, 0xdb, 0x79, 0xe7, 0xdd, 0x67, 0x47, 0x33, 0xb3, 0x30, 0x2e, 0xf2, 0xaa, 0x50,
	0x55, 0x2e, 0x19, 0xbf, 0x6b, 0x8f, 0xcd, 0xdc, 0x0b, 0x6e, 0x1b, 0xce, 0x24, 0x43, 0xd0, 0x2a,
	0x78, 0x06, 0x90, 0x2a, 0x9a, 0x90, 0xb5, 0x22, 0x42, 0xa2, 0x6b, 0x88, 0x96, 0x25, 0x17, 0x32,
	0xab, 0x15, 0x9d, 0x13, 0x3e, 0x0c, 0xde, 0x04, 0xe3, 0x93, 0x24, 0xd4, 0xda, 0x0f, 0x2d, 0xa1,
	0xb7, 0xd0, 0x17, 0xa4, 0x60, 0xf5, 0xc2, 0x79, 0x9e, 0x6b, 0x4f, 0x64, 0x44, 0x63, 0xc2, 0x37,
	0x10, 0x6a, 0xaa, 0x68, 0x58, 0x2d, 0x08, 0x1a, 0x01, 0x08, 0x45, 0x33, 0x4e, 0x84, 0xaa, 0xa4,
	0x85, 0x9e, 0x09, 0x6d, 0x50, 0x95, 0xc4, 0x1f, 0x21, 0xfa, 0xc9, 0x4b, 0x4a, 0x5c, 0x15, 0x23,
	0x00, 0xb9, 0x22, 0xdd, 0x1a, 0xce, 0xe4, 0x8a, 0x58, 0xf8, 0x03, 0xf4, 0xad, 0xdd, 0xe2, 0xaf,
	0x21, 0x6a, 0xb6, 0x42, 0xb6, 0xcc, 0x0b, 0xc9, 0x76, 0x55, 0x6b, 0x6d, 0xa2, 0x25, 0xfc, 0x0e,
	0xfa, 0xe6, 0xb6, 0x7b, 0xe3, 0x02, 0x4e, 0x3d, 0x7e, 0x2f, 0xb1, 0x11, 0x7e, 0x0f, 0x03, 0x67,
	0xb4, 0xf4, 0x21, 0xbc, 0xc8, 0x37, 0x84, 0xe7, 0xbf, 0x89, 0xb6, 0x06, 0x89, 0x0b, 0xf1, 0x0d,
	0xa0, 0x49, 0x59, 0x2f, 0xa6, 0xf9, 0x9f, 0x92, 0xb6, 0x3d, 0xec, 0x92, 0x4f, 0x76, 0xe4, 0x3b,
	0x78, 0xd9, 0x71, 0xb7, 0x78, 0x6a, 0x24, 0xeb, 0x77, 0x21, 0xfe, 0x00, 0xe7, 0xe9, 0x5a, 0xe5,
	0x9c, 0x24, 0x8c, 0xc9, 0xa7, 0xe8, 0x9f, 0x01, 0xf9, 0x66, 0x0b, 0xbf, 0x82, 0xd0, 0xe4, 0x33,
	0xce, 0x98, 0xb4, 0xf5, 0x83, 0x91, 0xb6, 0xc6, 0x87, 0xbf, 0x3d, 0x38, 0x7f, 0xdc, 0x6d, 0x43,
	0x4a, 0xf8, 0xa6, 0x2c, 0x08, 0xfa, 0x02, 0xbd, 0x54, 0x51, 0x74, 0x71, 0xeb, 0xad, 0x4e, 0xbb,
	0x25, 0xf1, 0xab, 0x3d, 0xdd, 0x3c, 0x87, 0x9f, 0xa1, 0x09, 0x84, 0x7a, 0x36, 0x76, 0x59, 0x86,
	0xbe, 0xd3, 0x9f, 0x71, 0xfc, 0xfa, 0x40, 0xc6, 0x51, 0xee, 0x03, 0xf4, 0x1d, 0x06, 0x8f, 0x8c,
	0x36, 0x4a, 0x92, 0xaf, 0xa6, 0xd9, 0xa8, 0x73, 0xa1, 0x33, 0xcb, 0x38, 0x3e, 0x94, 0x72, 0xb0,
	0x71, 0x80, 0x66, 0x10, 0x7a, 0x9d, 0x47, 0x97, 0xbe, 0x7d, 0x7f, 0x80, 0xf1, 0xd5, 0xd1, 0x7c,
	0xcb, 0xbc, 0x0f, 0xd0, 0x14, 0xa0, 0xed, 0x38, 0x1a, 0x75, 0x7a, 0xf2, 0xff, 0xd8, 0xe2, 0xcb,
	0x63, 0x69, 0x87, 0xfc, 0x36, 0xf8, 0x15, 0xf9, 0xbf, 0x76, 0x7e, 0xaa, 0xff, 0xea, 0xa7, 0x7f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x0d, 0x01, 0x65, 0xd7, 0x03, 0x00, 0x00,
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
	// Unary
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	PrimeNumber(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberClient, error)
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error)
	FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_FindMaximumClient, error)
	// Unary
	SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error)
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

func (c *calculatorServiceClient) SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error) {
	out := new(SquareRootResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/SquareRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	// Unary
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	PrimeNumber(*PrimeRequest, CalculatorService_PrimeNumberServer) error
	ComputeAverage(CalculatorService_ComputeAverageServer) error
	FindMaximum(CalculatorService_FindMaximumServer) error
	// Unary
	SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error)
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
func (*UnimplementedCalculatorServiceServer) SquareRoot(ctx context.Context, req *SquareRootRequest) (*SquareRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SquareRoot not implemented")
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

func _CalculatorService_SquareRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquareRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).SquareRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/SquareRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).SquareRoot(ctx, req.(*SquareRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
		{
			MethodName: "SquareRoot",
			Handler:    _CalculatorService_SquareRoot_Handler,
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
