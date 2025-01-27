// Code generated by protoc-gen-goten-go
// File: goten/meta-service/proto/v1/common.proto
// DO NOT EDIT!!!

package common

import (
	"fmt"
	"reflect"
	"sync"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import ()

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = fmt.Errorf
	_ = reflect.Method{}
	_ = sync.Once{}

	_ = protojson.MarshalOptions{}
	_ = proto.MarshalOptions{}
	_ = preflect.Value{}
	_ = protoimpl.DescBuilder{}
)

// make sure we're using proto imports
var ()

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// LabelledDomain provides connectivity information in a context of Service and
// specific label. Service/Deployment may provide multiple addresses, for
// example offering different quality class, like optimized for high/low
// throughput.
type LabelledDomain struct {
	state            protoimpl.MessageState
	sizeCache        protoimpl.SizeCache
	unknownFields    protoimpl.UnknownFields
	Label            string   `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty" firestore:"label"`
	Domain           string   `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty" firestore:"domain"`
	AvailableMixins  []string `protobuf:"bytes,3,rep,name=available_mixins,json=availableMixins,proto3" json:"available_mixins,omitempty" firestore:"availableMixins"`
	WebGrpcAvailable bool     `protobuf:"varint,4,opt,name=web_grpc_available,json=webGrpcAvailable,proto3" json:"web_grpc_available,omitempty" firestore:"webGrpcAvailable"`
	RestApiAvailable bool     `protobuf:"varint,5,opt,name=rest_api_available,json=restApiAvailable,proto3" json:"rest_api_available,omitempty" firestore:"restApiAvailable"`
	IsPrivate        bool     `protobuf:"varint,6,opt,name=is_private,json=isPrivate,proto3" json:"is_private,omitempty" firestore:"isPrivate"`
}

func (m *LabelledDomain) Reset() {
	*m = LabelledDomain{}
	if protoimpl.UnsafeEnabled {
		mi := &goten_meta_service_proto_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *LabelledDomain) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*LabelledDomain) ProtoMessage() {}

func (m *LabelledDomain) ProtoReflect() preflect.Message {
	mi := &goten_meta_service_proto_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*LabelledDomain) GotenMessage() {}

// Deprecated, Use LabelledDomain.ProtoReflect.Descriptor instead.
func (*LabelledDomain) Descriptor() ([]byte, []int) {
	return goten_meta_service_proto_v1_common_proto_rawDescGZIP(), []int{0}
}

func (m *LabelledDomain) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *LabelledDomain) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *LabelledDomain) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *LabelledDomain) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *LabelledDomain) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *LabelledDomain) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *LabelledDomain) GetAvailableMixins() []string {
	if m != nil {
		return m.AvailableMixins
	}
	return nil
}

func (m *LabelledDomain) GetWebGrpcAvailable() bool {
	if m != nil {
		return m.WebGrpcAvailable
	}
	return false
}

func (m *LabelledDomain) GetRestApiAvailable() bool {
	if m != nil {
		return m.RestApiAvailable
	}
	return false
}

func (m *LabelledDomain) GetIsPrivate() bool {
	if m != nil {
		return m.IsPrivate
	}
	return false
}

func (m *LabelledDomain) SetLabel(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Label", "LabelledDomain"))
	}
	m.Label = fv
}

func (m *LabelledDomain) SetDomain(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Domain", "LabelledDomain"))
	}
	m.Domain = fv
}

func (m *LabelledDomain) SetAvailableMixins(fv []string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AvailableMixins", "LabelledDomain"))
	}
	m.AvailableMixins = fv
}

func (m *LabelledDomain) SetWebGrpcAvailable(fv bool) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "WebGrpcAvailable", "LabelledDomain"))
	}
	m.WebGrpcAvailable = fv
}

func (m *LabelledDomain) SetRestApiAvailable(fv bool) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "RestApiAvailable", "LabelledDomain"))
	}
	m.RestApiAvailable = fv
}

func (m *LabelledDomain) SetIsPrivate(fv bool) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "IsPrivate", "LabelledDomain"))
	}
	m.IsPrivate = fv
}

var goten_meta_service_proto_v1_common_proto preflect.FileDescriptor

var goten_meta_service_proto_v1_common_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x02, 0x0a, 0x0e, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x22, 0x0a,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xca, 0xc6,
	0x27, 0x08, 0x2a, 0x06, 0x22, 0x02, 0x08, 0x40, 0x68, 0x01, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x12, 0x29, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x11, 0xca, 0xc6, 0x27, 0x0d, 0x2a, 0x0b, 0x22, 0x03, 0x08, 0xff, 0x01, 0x52, 0x02,
	0x60, 0x01, 0x68, 0x01, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x29, 0x0a, 0x10,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6d, 0x69, 0x78, 0x69, 0x6e, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x4d, 0x69, 0x78, 0x69, 0x6e, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x77, 0x65, 0x62, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x10, 0x77, 0x65, 0x62, 0x47, 0x72, 0x70, 0x63, 0x41, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70,
	0x69, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x10, 0x72, 0x65, 0x73, 0x74, 0x41, 0x70, 0x69, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x42, 0x6c, 0xe8, 0xde, 0x21, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x0b,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77,
	0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	goten_meta_service_proto_v1_common_proto_rawDescOnce sync.Once
	goten_meta_service_proto_v1_common_proto_rawDescData = goten_meta_service_proto_v1_common_proto_rawDesc
)

func goten_meta_service_proto_v1_common_proto_rawDescGZIP() []byte {
	goten_meta_service_proto_v1_common_proto_rawDescOnce.Do(func() {
		goten_meta_service_proto_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(goten_meta_service_proto_v1_common_proto_rawDescData)
	})
	return goten_meta_service_proto_v1_common_proto_rawDescData
}

var goten_meta_service_proto_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var goten_meta_service_proto_v1_common_proto_goTypes = []interface{}{
	(*LabelledDomain)(nil), // 0: goten.meta.v1.LabelledDomain
}
var goten_meta_service_proto_v1_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { goten_meta_service_proto_v1_common_proto_init() }
func goten_meta_service_proto_v1_common_proto_init() {
	if goten_meta_service_proto_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		goten_meta_service_proto_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelledDomain); i {
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
			RawDescriptor: goten_meta_service_proto_v1_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           goten_meta_service_proto_v1_common_proto_goTypes,
		DependencyIndexes: goten_meta_service_proto_v1_common_proto_depIdxs,
		MessageInfos:      goten_meta_service_proto_v1_common_proto_msgTypes,
	}.Build()
	goten_meta_service_proto_v1_common_proto = out.File
	goten_meta_service_proto_v1_common_proto_rawDesc = nil
	goten_meta_service_proto_v1_common_proto_goTypes = nil
	goten_meta_service_proto_v1_common_proto_depIdxs = nil
}