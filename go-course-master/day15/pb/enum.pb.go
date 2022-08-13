// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: pb/enum.proto

package pb

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

type Corpus int32

const (
	Corpus_UNIVERSAL Corpus = 0
	Corpus_WEB       Corpus = 1
	Corpus_IMAGES    Corpus = 2
	Corpus_LOCAL     Corpus = 3
	Corpus_NEWS      Corpus = 4
	Corpus_PRODUCTS  Corpus = 5
	Corpus_VIDEO     Corpus = 6
)

// Enum value maps for Corpus.
var (
	Corpus_name = map[int32]string{
		0: "UNIVERSAL",
		1: "WEB",
		2: "IMAGES",
		3: "LOCAL",
		4: "NEWS",
		5: "PRODUCTS",
		6: "VIDEO",
	}
	Corpus_value = map[string]int32{
		"UNIVERSAL": 0,
		"WEB":       1,
		"IMAGES":    2,
		"LOCAL":     3,
		"NEWS":      4,
		"PRODUCTS":  5,
		"VIDEO":     6,
	}
)

func (x Corpus) Enum() *Corpus {
	p := new(Corpus)
	*p = x
	return p
}

func (x Corpus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Corpus) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_enum_proto_enumTypes[0].Descriptor()
}

func (Corpus) Type() protoreflect.EnumType {
	return &file_pb_enum_proto_enumTypes[0]
}

func (x Corpus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Corpus.Descriptor instead.
func (Corpus) EnumDescriptor() ([]byte, []int) {
	return file_pb_enum_proto_rawDescGZIP(), []int{0}
}

var File_pb_enum_proto protoreflect.FileDescriptor

var file_pb_enum_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2a, 0x5a, 0x0a, 0x06, 0x43, 0x6f, 0x72, 0x70, 0x75, 0x73,
	0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x49, 0x56, 0x45, 0x52, 0x53, 0x41, 0x4c, 0x10, 0x00, 0x12,
	0x07, 0x0a, 0x03, 0x57, 0x45, 0x42, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x4d, 0x41, 0x47,
	0x45, 0x53, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x10, 0x03, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x45, 0x57, 0x53, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x52, 0x4f,
	0x44, 0x55, 0x43, 0x54, 0x53, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x49, 0x44, 0x45, 0x4f,
	0x10, 0x06, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x2f, 0x64, 0x61, 0x79, 0x32, 0x31, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_enum_proto_rawDescOnce sync.Once
	file_pb_enum_proto_rawDescData = file_pb_enum_proto_rawDesc
)

func file_pb_enum_proto_rawDescGZIP() []byte {
	file_pb_enum_proto_rawDescOnce.Do(func() {
		file_pb_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_enum_proto_rawDescData)
	})
	return file_pb_enum_proto_rawDescData
}

var file_pb_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_enum_proto_goTypes = []interface{}{
	(Corpus)(0), // 0: hello.Corpus
}
var file_pb_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_enum_proto_init() }
func file_pb_enum_proto_init() {
	if File_pb_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_enum_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_enum_proto_goTypes,
		DependencyIndexes: file_pb_enum_proto_depIdxs,
		EnumInfos:         file_pb_enum_proto_enumTypes,
	}.Build()
	File_pb_enum_proto = out.File
	file_pb_enum_proto_rawDesc = nil
	file_pb_enum_proto_goTypes = nil
	file_pb_enum_proto_depIdxs = nil
}
