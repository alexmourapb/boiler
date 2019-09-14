// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: user.proto

package entity

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type User struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Created              int64    `protobuf:"varint,3,opt,name=Created,proto3" json:"Created,omitempty"`
	Updated              int64    `protobuf:"varint,4,opt,name=Updated,proto3" json:"Updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *User) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "entity.User")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 118 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0x54, 0x8a,
	0xe3, 0x62, 0x09, 0x2d, 0x4e, 0x2d, 0x12, 0xe2, 0xe3, 0x62, 0xf2, 0x74, 0x91, 0x60, 0x54, 0x60,
	0xd4, 0x60, 0x0e, 0x62, 0xf2, 0x74, 0x11, 0x12, 0xe2, 0x62, 0xf1, 0x4b, 0xcc, 0x4d, 0x95, 0x60,
	0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x24, 0xb8, 0xd8, 0x9d, 0x8b, 0x52, 0x13, 0x4b,
	0x52, 0x53, 0x24, 0x98, 0xc1, 0x0a, 0x61, 0x5c, 0x90, 0x4c, 0x68, 0x41, 0x0a, 0x58, 0x86, 0x05,
	0x22, 0x03, 0xe5, 0x26, 0xb1, 0x81, 0xad, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x34,
	0xe2, 0xbf, 0x7c, 0x00, 0x00, 0x00,
}