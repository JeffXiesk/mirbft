// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nodeidentity.proto

package protobufs

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

type NodeIdentity struct {
	NodeId               int32    `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	PublicAddr           string   `protobuf:"bytes,2,opt,name=public_addr,json=publicAddr,proto3" json:"public_addr,omitempty"`
	PrivateAddr          string   `protobuf:"bytes,3,opt,name=private_addr,json=privateAddr,proto3" json:"private_addr,omitempty"`
	Port                 int32    `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	PubKey               []byte   `protobuf:"bytes,5,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeIdentity) Reset()         { *m = NodeIdentity{} }
func (m *NodeIdentity) String() string { return proto.CompactTextString(m) }
func (*NodeIdentity) ProtoMessage()    {}
func (*NodeIdentity) Descriptor() ([]byte, []int) {
	return fileDescriptor_50f107c74fbb13b2, []int{0}
}

func (m *NodeIdentity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeIdentity.Unmarshal(m, b)
}
func (m *NodeIdentity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeIdentity.Marshal(b, m, deterministic)
}
func (m *NodeIdentity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeIdentity.Merge(m, src)
}
func (m *NodeIdentity) XXX_Size() int {
	return xxx_messageInfo_NodeIdentity.Size(m)
}
func (m *NodeIdentity) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeIdentity.DiscardUnknown(m)
}

var xxx_messageInfo_NodeIdentity proto.InternalMessageInfo

func (m *NodeIdentity) GetNodeId() int32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *NodeIdentity) GetPublicAddr() string {
	if m != nil {
		return m.PublicAddr
	}
	return ""
}

func (m *NodeIdentity) GetPrivateAddr() string {
	if m != nil {
		return m.PrivateAddr
	}
	return ""
}

func (m *NodeIdentity) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *NodeIdentity) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func init() {
	proto.RegisterType((*NodeIdentity)(nil), "protobufs.NodeIdentity")
}

func init() {
	proto.RegisterFile("nodeidentity.proto", fileDescriptor_50f107c74fbb13b2)
}

var fileDescriptor_50f107c74fbb13b2 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0xcb, 0x4f, 0x49,
	0xcd, 0x4c, 0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x04, 0x53, 0x49, 0xa5, 0x69, 0xc5, 0x4a, 0x33, 0x18, 0xb9, 0x78, 0xfc, 0xf2, 0x53, 0x52, 0x3d,
	0xa1, 0x2a, 0x84, 0xc4, 0xb9, 0xd8, 0x41, 0x3a, 0xe2, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35,
	0x58, 0x83, 0xd8, 0xf2, 0xc0, 0xd2, 0x42, 0xf2, 0x5c, 0xdc, 0x05, 0xa5, 0x49, 0x39, 0x99, 0xc9,
	0xf1, 0x89, 0x29, 0x29, 0x45, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x5c, 0x10, 0x21, 0xc7,
	0x94, 0x94, 0x22, 0x21, 0x45, 0x2e, 0x9e, 0x82, 0xa2, 0xcc, 0xb2, 0xc4, 0x92, 0x54, 0x88, 0x0a,
	0x66, 0xb0, 0x0a, 0x6e, 0xa8, 0x18, 0x58, 0x89, 0x10, 0x17, 0x4b, 0x41, 0x7e, 0x51, 0x89, 0x04,
	0x0b, 0xd8, 0x64, 0x30, 0x1b, 0x64, 0x61, 0x41, 0x69, 0x52, 0x7c, 0x76, 0x6a, 0xa5, 0x04, 0xab,
	0x02, 0xa3, 0x06, 0x4f, 0x10, 0x5b, 0x41, 0x69, 0x92, 0x77, 0x6a, 0xa5, 0x13, 0x5f, 0x14, 0x8f,
	0x9e, 0xbe, 0x35, 0xdc, 0xa9, 0x49, 0x6c, 0x60, 0xa6, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x43,
	0x35, 0xa2, 0xf8, 0xd2, 0x00, 0x00, 0x00,
}