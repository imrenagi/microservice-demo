// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/proto/payment/payment.proto

package payment

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PaymentCreated struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	OrderID              string   `protobuf:"bytes,2,opt,name=orderID,proto3" json:"orderID,omitempty"`
	Value                float32  `protobuf:"fixed32,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentCreated) Reset()         { *m = PaymentCreated{} }
func (m *PaymentCreated) String() string { return proto.CompactTextString(m) }
func (*PaymentCreated) ProtoMessage()    {}
func (*PaymentCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_a75e174f730c8e0b, []int{0}
}
func (m *PaymentCreated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentCreated.Unmarshal(m, b)
}
func (m *PaymentCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentCreated.Marshal(b, m, deterministic)
}
func (dst *PaymentCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentCreated.Merge(dst, src)
}
func (m *PaymentCreated) XXX_Size() int {
	return xxx_messageInfo_PaymentCreated.Size(m)
}
func (m *PaymentCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentCreated.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentCreated proto.InternalMessageInfo

func (m *PaymentCreated) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *PaymentCreated) GetOrderID() string {
	if m != nil {
		return m.OrderID
	}
	return ""
}

func (m *PaymentCreated) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*PaymentCreated)(nil), "payment.PaymentCreated")
}

func init() {
	proto.RegisterFile("pkg/proto/payment/payment.proto", fileDescriptor_payment_a75e174f730c8e0b)
}

var fileDescriptor_payment_a75e174f730c8e0b = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0xc8, 0x4e, 0xd7,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x48, 0xac, 0xcc, 0x4d, 0xcd, 0x2b, 0x81, 0xd1, 0x7a,
	0x60, 0x51, 0x21, 0x76, 0x28, 0x57, 0x29, 0x80, 0x8b, 0x2f, 0x00, 0xc2, 0x74, 0x2e, 0x4a, 0x4d,
	0x2c, 0x49, 0x4d, 0x11, 0xe2, 0xe3, 0x62, 0xf2, 0x74, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c,
	0x62, 0xf2, 0x74, 0x11, 0x92, 0xe0, 0x62, 0xcf, 0x2f, 0x4a, 0x49, 0x2d, 0xf2, 0x74, 0x91, 0x60,
	0x02, 0x0b, 0xc2, 0xb8, 0x42, 0x22, 0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x12, 0xcc, 0x0a,
	0x8c, 0x1a, 0x4c, 0x41, 0x10, 0x4e, 0x12, 0x1b, 0xd8, 0x06, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x23, 0x96, 0x8d, 0xd4, 0x84, 0x00, 0x00, 0x00,
}