// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: goten/runtime/api/view.proto

package view

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

// View defines list of fields of the resource that will
// be obtained in servers response.
type View int32

const (
	View_UNSPECIFIED View = 0
	View_NAME        View = 1
	View_BASIC       View = 2
	View_DETAIL      View = 3
	View_FULL        View = 4
)

// Enum value maps for View.
var (
	View_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "NAME",
		2: "BASIC",
		3: "DETAIL",
		4: "FULL",
	}
	View_value = map[string]int32{
		"UNSPECIFIED": 0,
		"NAME":        1,
		"BASIC":       2,
		"DETAIL":      3,
		"FULL":        4,
	}
)

func (x View) Enum() *View {
	p := new(View)
	*p = x
	return p
}

func (x View) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (View) Descriptor() protoreflect.EnumDescriptor {
	return file_goten_runtime_api_view_proto_enumTypes[0].Descriptor()
}

func (View) Type() protoreflect.EnumType {
	return &file_goten_runtime_api_view_proto_enumTypes[0]
}

func (x View) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use View.Descriptor instead.
func (View) EnumDescriptor() ([]byte, []int) {
	return file_goten_runtime_api_view_proto_rawDescGZIP(), []int{0}
}

var File_goten_runtime_api_view_proto protoreflect.FileDescriptor

var file_goten_runtime_api_view_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x2a, 0x42, 0x0a, 0x04, 0x56, 0x69,
	0x65, 0x77, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x01, 0x12, 0x09, 0x0a,
	0x05, 0x42, 0x41, 0x53, 0x49, 0x43, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x54, 0x41,
	0x49, 0x4c, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x55, 0x4c, 0x4c, 0x10, 0x04, 0x42, 0x4a,
	0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e,
	0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x50, 0x01, 0x5a, 0x2a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x77, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_goten_runtime_api_view_proto_rawDescOnce sync.Once
	file_goten_runtime_api_view_proto_rawDescData = file_goten_runtime_api_view_proto_rawDesc
)

func file_goten_runtime_api_view_proto_rawDescGZIP() []byte {
	file_goten_runtime_api_view_proto_rawDescOnce.Do(func() {
		file_goten_runtime_api_view_proto_rawDescData = protoimpl.X.CompressGZIP(file_goten_runtime_api_view_proto_rawDescData)
	})
	return file_goten_runtime_api_view_proto_rawDescData
}

var file_goten_runtime_api_view_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_goten_runtime_api_view_proto_goTypes = []interface{}{
	(View)(0), // 0: goten.view.View
}
var file_goten_runtime_api_view_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_goten_runtime_api_view_proto_init() }
func file_goten_runtime_api_view_proto_init() {
	if File_goten_runtime_api_view_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_goten_runtime_api_view_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_goten_runtime_api_view_proto_goTypes,
		DependencyIndexes: file_goten_runtime_api_view_proto_depIdxs,
		EnumInfos:         file_goten_runtime_api_view_proto_enumTypes,
	}.Build()
	File_goten_runtime_api_view_proto = out.File
	file_goten_runtime_api_view_proto_rawDesc = nil
	file_goten_runtime_api_view_proto_goTypes = nil
	file_goten_runtime_api_view_proto_depIdxs = nil
}
