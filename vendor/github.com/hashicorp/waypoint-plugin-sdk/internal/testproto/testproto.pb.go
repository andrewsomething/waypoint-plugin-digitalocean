// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.2
// source: testproto.proto

package testproto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Data is just some data, used for tests so meant to be meaningless.
type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value  string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Number int32  `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testproto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_testproto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_testproto_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Data) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

// Other types that we can use. There isn't any purpose for this other
// than to provide message types that can be used for tests.
type A struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *A) Reset() {
	*x = A{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testproto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *A) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*A) ProtoMessage() {}

func (x *A) ProtoReflect() protoreflect.Message {
	mi := &file_testproto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use A.ProtoReflect.Descriptor instead.
func (*A) Descriptor() ([]byte, []int) {
	return file_testproto_proto_rawDescGZIP(), []int{1}
}

func (x *A) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type B struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *B) Reset() {
	*x = B{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testproto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *B) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*B) ProtoMessage() {}

func (x *B) ProtoReflect() protoreflect.Message {
	mi := &file_testproto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use B.ProtoReflect.Descriptor instead.
func (*B) Descriptor() ([]byte, []int) {
	return file_testproto_proto_rawDescGZIP(), []int{2}
}

func (x *B) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_testproto_proto protoreflect.FileDescriptor

var file_testproto_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x19, 0x0a, 0x01, 0x41, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x19, 0x0a,
	0x01, 0x42, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x74, 0x65,
	0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testproto_proto_rawDescOnce sync.Once
	file_testproto_proto_rawDescData = file_testproto_proto_rawDesc
)

func file_testproto_proto_rawDescGZIP() []byte {
	file_testproto_proto_rawDescOnce.Do(func() {
		file_testproto_proto_rawDescData = protoimpl.X.CompressGZIP(file_testproto_proto_rawDescData)
	})
	return file_testproto_proto_rawDescData
}

var file_testproto_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_testproto_proto_goTypes = []interface{}{
	(*Data)(nil), // 0: testproto.Data
	(*A)(nil),    // 1: testproto.A
	(*B)(nil),    // 2: testproto.B
}
var file_testproto_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_testproto_proto_init() }
func file_testproto_proto_init() {
	if File_testproto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testproto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_testproto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*A); i {
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
		file_testproto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*B); i {
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
			RawDescriptor: file_testproto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_testproto_proto_goTypes,
		DependencyIndexes: file_testproto_proto_depIdxs,
		MessageInfos:      file_testproto_proto_msgTypes,
	}.Build()
	File_testproto_proto = out.File
	file_testproto_proto_rawDesc = nil
	file_testproto_proto_goTypes = nil
	file_testproto_proto_depIdxs = nil
}
