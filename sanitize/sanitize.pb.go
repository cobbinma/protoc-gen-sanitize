// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: sanitize.proto

package sanitize

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type Sanitization int32

const (
	Sanitization_NONE Sanitization = 0
	Sanitization_TEXT Sanitization = 1
	Sanitization_HTML Sanitization = 2
)

// Enum value maps for Sanitization.
var (
	Sanitization_name = map[int32]string{
		0: "NONE",
		1: "TEXT",
		2: "HTML",
	}
	Sanitization_value = map[string]int32{
		"NONE": 0,
		"TEXT": 1,
		"HTML": 2,
	}
)

func (x Sanitization) Enum() *Sanitization {
	p := new(Sanitization)
	*p = x
	return p
}

func (x Sanitization) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Sanitization) Descriptor() protoreflect.EnumDescriptor {
	return file_sanitize_proto_enumTypes[0].Descriptor()
}

func (Sanitization) Type() protoreflect.EnumType {
	return &file_sanitize_proto_enumTypes[0]
}

func (x Sanitization) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Sanitization.Descriptor instead.
func (Sanitization) EnumDescriptor() ([]byte, []int) {
	return file_sanitize_proto_rawDescGZIP(), []int{0}
}

var file_sanitize_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.FileOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         2181,
		Name:          "sanitize.disable_file",
		Tag:           "varint,2181,opt,name=disable_file",
		Filename:      "sanitize.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         2181,
		Name:          "sanitize.disable_message",
		Tag:           "varint,2181,opt,name=disable_message",
		Filename:      "sanitize.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*Sanitization)(nil),
		Field:         2181,
		Name:          "sanitize.kind",
		Tag:           "varint,2181,opt,name=kind,enum=sanitize.Sanitization",
		Filename:      "sanitize.proto",
	},
}

// Extension fields to descriptor.FileOptions.
var (
	// optional bool disable_file = 2181;
	E_DisableFile = &file_sanitize_proto_extTypes[0]
)

// Extension fields to descriptor.MessageOptions.
var (
	// optional bool disable_message = 2181;
	E_DisableMessage = &file_sanitize_proto_extTypes[1]
)

// Extension fields to descriptor.FieldOptions.
var (
	// optional sanitize.Sanitization kind = 2181;
	E_Kind = &file_sanitize_proto_extTypes[2]
)

var File_sanitize_proto protoreflect.FileDescriptor

var file_sanitize_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x61, 0x6e, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x73, 0x61, 0x6e, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x2c, 0x0a, 0x0c,
	0x53, 0x61, 0x6e, 0x69, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x45, 0x58, 0x54, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x48, 0x54, 0x4d, 0x4c, 0x10, 0x02, 0x3a, 0x43, 0x0a, 0x0c, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x85, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x3a,
	0x4c, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x85, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x3a, 0x4d, 0x0a,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x85, 0x11, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x73, 0x61,
	0x6e, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x2e, 0x53, 0x61, 0x6e, 0x69, 0x74, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x88, 0x01, 0x01, 0x42, 0x33, 0x5a, 0x31,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x74, 0x72, 0x69,
	0x6e, 0x73, 0x65, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x73, 0x61, 0x6e, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x2f, 0x73, 0x61, 0x6e, 0x69, 0x74, 0x69, 0x7a,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sanitize_proto_rawDescOnce sync.Once
	file_sanitize_proto_rawDescData = file_sanitize_proto_rawDesc
)

func file_sanitize_proto_rawDescGZIP() []byte {
	file_sanitize_proto_rawDescOnce.Do(func() {
		file_sanitize_proto_rawDescData = protoimpl.X.CompressGZIP(file_sanitize_proto_rawDescData)
	})
	return file_sanitize_proto_rawDescData
}

var file_sanitize_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sanitize_proto_goTypes = []interface{}{
	(Sanitization)(0),                 // 0: sanitize.Sanitization
	(*descriptor.FileOptions)(nil),    // 1: google.protobuf.FileOptions
	(*descriptor.MessageOptions)(nil), // 2: google.protobuf.MessageOptions
	(*descriptor.FieldOptions)(nil),   // 3: google.protobuf.FieldOptions
}
var file_sanitize_proto_depIdxs = []int32{
	1, // 0: sanitize.disable_file:extendee -> google.protobuf.FileOptions
	2, // 1: sanitize.disable_message:extendee -> google.protobuf.MessageOptions
	3, // 2: sanitize.kind:extendee -> google.protobuf.FieldOptions
	0, // 3: sanitize.kind:type_name -> sanitize.Sanitization
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	3, // [3:4] is the sub-list for extension type_name
	0, // [0:3] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sanitize_proto_init() }
func file_sanitize_proto_init() {
	if File_sanitize_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sanitize_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 3,
			NumServices:   0,
		},
		GoTypes:           file_sanitize_proto_goTypes,
		DependencyIndexes: file_sanitize_proto_depIdxs,
		EnumInfos:         file_sanitize_proto_enumTypes,
		ExtensionInfos:    file_sanitize_proto_extTypes,
	}.Build()
	File_sanitize_proto = out.File
	file_sanitize_proto_rawDesc = nil
	file_sanitize_proto_goTypes = nil
	file_sanitize_proto_depIdxs = nil
}
