// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.12.4
// source: common.proto

package protobufs

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HotstuffSignedMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data      []byte      `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Signature []byte      `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	TsQc      *HotStuffQC `protobuf:"bytes,3,opt,name=TsQc,proto3" json:"TsQc,omitempty"`
	DigestPar []byte      `protobuf:"bytes,4,opt,name=digest_par,json=digestPar,proto3" json:"digest_par,omitempty"`
}

func (x *HotstuffSignedMsg) Reset() {
	*x = HotstuffSignedMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HotstuffSignedMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HotstuffSignedMsg) ProtoMessage() {}

func (x *HotstuffSignedMsg) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HotstuffSignedMsg.ProtoReflect.Descriptor instead.
func (*HotstuffSignedMsg) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *HotstuffSignedMsg) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *HotstuffSignedMsg) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *HotstuffSignedMsg) GetTsQc() *HotStuffQC {
	if x != nil {
		return x.TsQc
	}
	return nil
}

func (x *HotstuffSignedMsg) GetDigestPar() []byte {
	if x != nil {
		return x.DigestPar
	}
	return nil
}

type SignedMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data      []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SignedMsg) Reset() {
	*x = SignedMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedMsg) ProtoMessage() {}

func (x *SignedMsg) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedMsg.ProtoReflect.Descriptor instead.
func (*SignedMsg) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *SignedMsg) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SignedMsg) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type Timeout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sn   int32 `protobuf:"varint,1,opt,name=sn,proto3" json:"sn,omitempty"`
	View int32 `protobuf:"varint,2,opt,name=view,proto3" json:"view,omitempty"`
}

func (x *Timeout) Reset() {
	*x = Timeout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timeout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timeout) ProtoMessage() {}

func (x *Timeout) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timeout.ProtoReflect.Descriptor instead.
func (*Timeout) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *Timeout) GetSn() int32 {
	if x != nil {
		return x.Sn
	}
	return 0
}

func (x *Timeout) GetView() int32 {
	if x != nil {
		return x.View
	}
	return 0
}

type Heartbeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Heartbeat) Reset() {
	*x = Heartbeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Heartbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Heartbeat) ProtoMessage() {}

func (x *Heartbeat) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Heartbeat.ProtoReflect.Descriptor instead.
func (*Heartbeat) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x1a, 0x15, 0x68, 0x6f, 0x74, 0x73, 0x74,
	0x75, 0x66, 0x66, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8f, 0x01, 0x0a, 0x11, 0x48, 0x6f, 0x74, 0x73, 0x74, 0x75, 0x66, 0x66, 0x53, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x54, 0x73, 0x51, 0x63,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x73, 0x2e, 0x48, 0x6f, 0x74, 0x53, 0x74, 0x75, 0x66, 0x66, 0x51, 0x43, 0x52, 0x04, 0x54,
	0x73, 0x51, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x61,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x50,
	0x61, 0x72, 0x22, 0x3d, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x73, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x22, 0x2d, 0x0a, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x73, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x73, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x76, 0x69, 0x65, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x76, 0x69, 0x65, 0x77,
	0x22, 0x0b, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x42, 0x0e, 0x5a,
	0x0c, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_common_proto_goTypes = []interface{}{
	(*HotstuffSignedMsg)(nil), // 0: protobufs.HotstuffSignedMsg
	(*SignedMsg)(nil),         // 1: protobufs.SignedMsg
	(*Timeout)(nil),           // 2: protobufs.Timeout
	(*Heartbeat)(nil),         // 3: protobufs.Heartbeat
	(*HotStuffQC)(nil),        // 4: protobufs.HotStuffQC
}
var file_common_proto_depIdxs = []int32{
	4, // 0: protobufs.HotstuffSignedMsg.TsQc:type_name -> protobufs.HotStuffQC
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	file_hotstufforderer_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HotstuffSignedMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timeout); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Heartbeat); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
