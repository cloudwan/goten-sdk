// Code generated by protoc-gen-goten-go
// File: goten/runtime/resource/sharding/ring_config.proto
// DO NOT EDIT!!!

package sharding

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

// RingConfig is an information about single shard range for one specific purpose on a set of resources.
// Example:
//  {
//    size: 4096
//    field_paths: [
//     { path: "meta.shards.byProjectIdMod4096", label: "byProjectId" }
//    ]
//  }
// In the example above, we indicate that some set of resources (in this case, those that have project ID specified)
// are shardable and their respective shard IDs are stored in the field path with label "byProjectId". Also, size of
// the ring describing this shard range is equal to 4096.
// Labels are useful when there is more than one field path for given ring. This can happen, if we have two resources
// which must be within same ring, but have different paths. Or another case, when one resource instance belongs to two
// different rings of the same type -> for example, IAM/Group resources are bucketed into rings by their identifiers
// and, if one group belongs to another, their retrospective GroupMember resources are visible in two rings. In other
// words, resources that form a graph and resources "on the borders" should be visible in both groups.
// Notes:
// * Path pointed by FieldPath must be of type int64. Also map<..., int64> is also supported (meta.shards) can be a map.
// RingConfig in the controller config must be in sync with at least one RingConfig in server side config. Server side
// config can use multiple rings at once.
// * While controller uses just one RingConfig at the time, server can use many simultaneously. It is required that
// labelling across array of RingConfigs is consistent.
// * Common components for sharding offer ShardDecorator. It is a generic component that decorates each relevant
// resource with shard ID. It handles common logic. While developer is still responsible for providing algorithm to it,
// calculation of the actual shard ID (combination of hash + ring size), iterating over all ring configs to provide
// number for all the rings is done there. This component however requires list of algorithms where each algorithm
// must tell, via Label, which field path intends to write.
type RingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Ring size. It affects possible shard numbers assigned to all relevant resources, which are in range: [0:size).
	Size uint64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty" firestore:"size"`
	// Field path in the resource where shard number is stored, for example "meta.shards.byProjectIdMod4096". Expected
	// underlying type is int64.
	FieldPaths []*RingConfig_FieldPath `protobuf:"bytes,2,rep,name=field_paths,json=fieldPaths,proto3" json:"field_paths,omitempty" firestore:"fieldPaths"`
}

func (m *RingConfig) Reset() {
	*m = RingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &goten_runtime_resource_sharding_ring_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RingConfig) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RingConfig) ProtoMessage() {}

func (m *RingConfig) ProtoReflect() preflect.Message {
	mi := &goten_runtime_resource_sharding_ring_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RingConfig) GotenMessage() {}

// Deprecated, Use RingConfig.ProtoReflect.Descriptor instead.
func (*RingConfig) Descriptor() ([]byte, []int) {
	return goten_runtime_resource_sharding_ring_config_proto_rawDescGZIP(), []int{0}
}

func (m *RingConfig) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RingConfig) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RingConfig) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RingConfig) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RingConfig) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return uint64(0)
}

func (m *RingConfig) GetFieldPaths() []*RingConfig_FieldPath {
	if m != nil {
		return m.FieldPaths
	}
	return nil
}

func (m *RingConfig) SetSize(fv uint64) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Size", "RingConfig"))
	}
	m.Size = fv
}

func (m *RingConfig) SetFieldPaths(fv []*RingConfig_FieldPath) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldPaths", "RingConfig"))
	}
	m.FieldPaths = fv
}

type RingConfig_FieldPath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Path in the resource where shard number is stored.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty" firestore:"path"`
	// Label identifying field path.
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty" firestore:"label"`
}

func (m *RingConfig_FieldPath) Reset() {
	*m = RingConfig_FieldPath{}
	if protoimpl.UnsafeEnabled {
		mi := &goten_runtime_resource_sharding_ring_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RingConfig_FieldPath) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RingConfig_FieldPath) ProtoMessage() {}

func (m *RingConfig_FieldPath) ProtoReflect() preflect.Message {
	mi := &goten_runtime_resource_sharding_ring_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RingConfig_FieldPath) GotenMessage() {}

// Deprecated, Use RingConfig_FieldPath.ProtoReflect.Descriptor instead.
func (*RingConfig_FieldPath) Descriptor() ([]byte, []int) {
	return goten_runtime_resource_sharding_ring_config_proto_rawDescGZIP(), []int{0, 0}
}

func (m *RingConfig_FieldPath) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RingConfig_FieldPath) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RingConfig_FieldPath) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RingConfig_FieldPath) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RingConfig_FieldPath) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *RingConfig_FieldPath) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *RingConfig_FieldPath) SetPath(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Path", "RingConfig_FieldPath"))
	}
	m.Path = fv
}

func (m *RingConfig_FieldPath) SetLabel(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Label", "RingConfig_FieldPath"))
	}
	m.Label = fv
}

var goten_runtime_resource_sharding_ring_config_proto preflect.FileDescriptor

var goten_runtime_resource_sharding_ring_config_proto_rawDesc = []byte{
	0x0a, 0x31, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x2f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a, 0x0a, 0x52, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x61, 0x74,
	0x68, 0x52, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x61, 0x74, 0x68, 0x73, 0x1a, 0x35, 0x0a,
	0x09, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	goten_runtime_resource_sharding_ring_config_proto_rawDescOnce sync.Once
	goten_runtime_resource_sharding_ring_config_proto_rawDescData = goten_runtime_resource_sharding_ring_config_proto_rawDesc
)

func goten_runtime_resource_sharding_ring_config_proto_rawDescGZIP() []byte {
	goten_runtime_resource_sharding_ring_config_proto_rawDescOnce.Do(func() {
		goten_runtime_resource_sharding_ring_config_proto_rawDescData = protoimpl.X.CompressGZIP(goten_runtime_resource_sharding_ring_config_proto_rawDescData)
	})
	return goten_runtime_resource_sharding_ring_config_proto_rawDescData
}

var goten_runtime_resource_sharding_ring_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var goten_runtime_resource_sharding_ring_config_proto_goTypes = []interface{}{
	(*RingConfig)(nil),           // 0: goten.sharding.RingConfig
	(*RingConfig_FieldPath)(nil), // 1: goten.sharding.RingConfig.FieldPath
}
var goten_runtime_resource_sharding_ring_config_proto_depIdxs = []int32{
	1, // 0: goten.sharding.RingConfig.field_paths:type_name -> goten.sharding.RingConfig.FieldPath
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { goten_runtime_resource_sharding_ring_config_proto_init() }
func goten_runtime_resource_sharding_ring_config_proto_init() {
	if goten_runtime_resource_sharding_ring_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		goten_runtime_resource_sharding_ring_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RingConfig); i {
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
		goten_runtime_resource_sharding_ring_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RingConfig_FieldPath); i {
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
			RawDescriptor: goten_runtime_resource_sharding_ring_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           goten_runtime_resource_sharding_ring_config_proto_goTypes,
		DependencyIndexes: goten_runtime_resource_sharding_ring_config_proto_depIdxs,
		MessageInfos:      goten_runtime_resource_sharding_ring_config_proto_msgTypes,
	}.Build()
	goten_runtime_resource_sharding_ring_config_proto = out.File
	goten_runtime_resource_sharding_ring_config_proto_rawDesc = nil
	goten_runtime_resource_sharding_ring_config_proto_goTypes = nil
	goten_runtime_resource_sharding_ring_config_proto_depIdxs = nil
}