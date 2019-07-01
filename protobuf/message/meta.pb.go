// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message/meta.proto

package message

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

// 元消息
type Meta struct {
	Method               string   `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Pagination           string   `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
	ErrorCode            int32    `protobuf:"varint,4,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Meta) Reset()         { *m = Meta{} }
func (m *Meta) String() string { return proto.CompactTextString(m) }
func (*Meta) ProtoMessage()    {}
func (*Meta) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcbd3f844af06f4d, []int{0}
}

func (m *Meta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Meta.Unmarshal(m, b)
}
func (m *Meta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Meta.Marshal(b, m, deterministic)
}
func (m *Meta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Meta.Merge(m, src)
}
func (m *Meta) XXX_Size() int {
	return xxx_messageInfo_Meta.Size(m)
}
func (m *Meta) XXX_DiscardUnknown() {
	xxx_messageInfo_Meta.DiscardUnknown(m)
}

var xxx_messageInfo_Meta proto.InternalMessageInfo

func (m *Meta) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Meta) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Meta) GetPagination() string {
	if m != nil {
		return m.Pagination
	}
	return ""
}

func (m *Meta) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func init() {
	proto.RegisterType((*Meta)(nil), "grpc_demo_go.protobuf.message.Meta")
}

func init() { proto.RegisterFile("message/meta.proto", fileDescriptor_fcbd3f844af06f4d) }

var fileDescriptor_fcbd3f844af06f4d = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xb1, 0x8e, 0x83, 0x30,
	0x10, 0x44, 0x05, 0xc7, 0x71, 0x62, 0xaf, 0x73, 0x71, 0x72, 0xc3, 0x09, 0xa5, 0xa2, 0x01, 0x8a,
	0x74, 0x29, 0x49, 0x4d, 0x43, 0x99, 0x06, 0x19, 0xbc, 0x71, 0x28, 0xcc, 0x22, 0xe3, 0x28, 0xff,
	0x93, 0x2f, 0x8d, 0x30, 0x46, 0xa2, 0xdc, 0x99, 0xa7, 0x9d, 0x07, 0x4c, 0xe3, 0xb2, 0x08, 0x85,
	0x95, 0x46, 0x2b, 0xca, 0xd9, 0x90, 0x25, 0x96, 0x2a, 0x33, 0x0f, 0x9d, 0x44, 0x4d, 0x9d, 0xa2,
	0x2d, 0xeb, 0x9f, 0xf7, 0xd2, 0x93, 0xa7, 0x17, 0x44, 0x0d, 0x5a, 0xc1, 0xfe, 0x20, 0xd6, 0x68,
	0x1f, 0x24, 0x79, 0x90, 0x05, 0x79, 0xd2, 0xfa, 0x8b, 0x71, 0xf8, 0xf1, 0x28, 0x0f, 0x5d, 0xb1,
	0x9f, 0xec, 0x1f, 0x60, 0x16, 0x6a, 0x9c, 0x84, 0x1d, 0x69, 0xe2, 0x5f, 0xae, 0x3c, 0x24, 0x2c,
	0x05, 0x40, 0x63, 0xc8, 0x74, 0x03, 0x49, 0xe4, 0x51, 0x16, 0xe4, 0xdf, 0x6d, 0xe2, 0x92, 0x2b,
	0x49, 0xac, 0x2f, 0xf5, 0xef, 0x3a, 0xdc, 0x6c, 0xdf, 0x6e, 0x4e, 0xb3, 0x58, 0x35, 0x0b, 0x45,
	0xd5, 0xae, 0x59, 0xf9, 0xb1, 0x77, 0x78, 0x84, 0xfb, 0xd8, 0xf5, 0xe7, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xe1, 0xe9, 0x8e, 0xf3, 0xf0, 0x00, 0x00, 0x00,
}
