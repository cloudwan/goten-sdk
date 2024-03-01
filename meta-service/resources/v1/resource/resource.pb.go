// Code generated by protoc-gen-goten-go
// File: goten/meta-service/proto/v1/resource.proto
// DO NOT EDIT!!!

package resource

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
import (
	service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

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
var (
	_ = &service.Service{}
	_ = &meta.Meta{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Resource Resource
type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of Resource
	// When creating a new instance, this field is optional and if not provided,
	// it will be generated automatically. Last ID segment must conform to the
	// following regex: [a-zA-Z]{1,128}
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Metadata is an object with information like create, update and delete time
	// (for async deleted resources), has user labels/annotations, sharding
	// information, multi-region syncing information and may have non-schema
	// owners (useful for taking ownership of resources belonging to lower level
	// services by higher ones).
	Metadata *meta.Meta `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
	// Plural name of resource in CamelForm, for example "Devices"
	PluralName string `protobuf:"bytes,3,opt,name=plural_name,json=pluralName,proto3" json:"plural_name,omitempty" firestore:"pluralName"`
	// Fully qualified name, for example devices.edgelq.com/Device.
	// It can also be derived from name too.
	Fqn string `protobuf:"bytes,4,opt,name=fqn,proto3" json:"fqn,omitempty" firestore:"fqn"`
	// List of service versions where this resource is known.
	// This list will be sorted if created by meta service provided by Goten.
	// The newest version is first, the oldest last.
	Versions []string `protobuf:"bytes,5,rep,name=versions,proto3" json:"versions,omitempty" firestore:"versions"`
	// Versioned information holds values that may be valid for specific versions
	// only. Sorted exactly like versions list.
	VersionedInfos []*Resource_VersionedInfo `protobuf:"bytes,6,rep,name=versioned_infos,json=versionedInfos,proto3" json:"versioned_infos,omitempty" firestore:"versionedInfos"`
	// Internal allowed services generation field from Service resource.
	// Used for metadata synchronization.
	AllowedServicesGeneration int64 `protobuf:"varint,7,opt,name=allowed_services_generation,json=allowedServicesGeneration,proto3" json:"allowed_services_generation,omitempty" firestore:"allowedServicesGeneration"`
}

func (m *Resource) Reset() {
	*m = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &goten_meta_service_proto_v1_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Resource) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Resource) ProtoMessage() {}

func (m *Resource) ProtoReflect() preflect.Message {
	mi := &goten_meta_service_proto_v1_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Resource) GotenMessage() {}

// Deprecated, Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return goten_meta_service_proto_v1_resource_proto_rawDescGZIP(), []int{0}
}

func (m *Resource) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Resource) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Resource) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Resource) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Resource) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Resource) GetMetadata() *meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Resource) GetPluralName() string {
	if m != nil {
		return m.PluralName
	}
	return ""
}

func (m *Resource) GetFqn() string {
	if m != nil {
		return m.Fqn
	}
	return ""
}

func (m *Resource) GetVersions() []string {
	if m != nil {
		return m.Versions
	}
	return nil
}

func (m *Resource) GetVersionedInfos() []*Resource_VersionedInfo {
	if m != nil {
		return m.VersionedInfos
	}
	return nil
}

func (m *Resource) GetAllowedServicesGeneration() int64 {
	if m != nil {
		return m.AllowedServicesGeneration
	}
	return int64(0)
}

func (m *Resource) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "Resource"))
	}
	m.Name = fv
}

func (m *Resource) SetMetadata(fv *meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "Resource"))
	}
	m.Metadata = fv
}

func (m *Resource) SetPluralName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PluralName", "Resource"))
	}
	m.PluralName = fv
}

func (m *Resource) SetFqn(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Fqn", "Resource"))
	}
	m.Fqn = fv
}

func (m *Resource) SetVersions(fv []string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Versions", "Resource"))
	}
	m.Versions = fv
}

func (m *Resource) SetVersionedInfos(fv []*Resource_VersionedInfo) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "VersionedInfos", "Resource"))
	}
	m.VersionedInfos = fv
}

func (m *Resource) SetAllowedServicesGeneration(fv int64) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AllowedServicesGeneration", "Resource"))
	}
	m.AllowedServicesGeneration = fv
}

// VersionedInfo contains specification part that is versioned.
type Resource_VersionedInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Resource version this information applies to.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty" firestore:"version"`
	// Whether resource is regional.
	IsRegional bool `protobuf:"varint,2,opt,name=is_regional,json=isRegional,proto3" json:"is_regional,omitempty" firestore:"isRegional"`
}

func (m *Resource_VersionedInfo) Reset() {
	*m = Resource_VersionedInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &goten_meta_service_proto_v1_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Resource_VersionedInfo) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Resource_VersionedInfo) ProtoMessage() {}

func (m *Resource_VersionedInfo) ProtoReflect() preflect.Message {
	mi := &goten_meta_service_proto_v1_resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Resource_VersionedInfo) GotenMessage() {}

// Deprecated, Use Resource_VersionedInfo.ProtoReflect.Descriptor instead.
func (*Resource_VersionedInfo) Descriptor() ([]byte, []int) {
	return goten_meta_service_proto_v1_resource_proto_rawDescGZIP(), []int{0, 0}
}

func (m *Resource_VersionedInfo) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Resource_VersionedInfo) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Resource_VersionedInfo) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Resource_VersionedInfo) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Resource_VersionedInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Resource_VersionedInfo) GetIsRegional() bool {
	if m != nil {
		return m.IsRegional
	}
	return false
}

func (m *Resource_VersionedInfo) SetVersion(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Version", "Resource_VersionedInfo"))
	}
	m.Version = fv
}

func (m *Resource_VersionedInfo) SetIsRegional(fv bool) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "IsRegional", "Resource_VersionedInfo"))
	}
	m.IsRegional = fv
}

var goten_meta_service_proto_v1_resource_proto preflect.FileDescriptor

var goten_meta_service_proto_v1_resource_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x67, 0x6f, 0x74, 0x65, 0x6e,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x04, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x10, 0xb2, 0xda, 0x21, 0x0c, 0x0a, 0x0a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x43, 0x0a, 0x0b, 0x70, 0x6c, 0x75,
	0x72, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22,
	0xca, 0xc6, 0x27, 0x1e, 0x2a, 0x1c, 0x52, 0x18, 0x42, 0x16, 0x5e, 0x5b, 0x41, 0x2d, 0x5a, 0x5d,
	0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x5d, 0x7b, 0x31, 0x2c, 0x31, 0x33, 0x32, 0x7d, 0x24,
	0x68, 0x01, 0x52, 0x0a, 0x70, 0x6c, 0x75, 0x72, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x03, 0x66, 0x71, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xf0, 0xd9, 0x21,
	0x01, 0x52, 0x03, 0x66, 0x71, 0x6e, 0x12, 0x30, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x42, 0x14, 0xca, 0xc6, 0x27, 0x10, 0x42, 0x0e,
	0x0a, 0x02, 0x08, 0x01, 0x18, 0x01, 0x22, 0x06, 0x2a, 0x04, 0x3a, 0x02, 0x08, 0x40, 0x52, 0x08,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x5c, 0x0a, 0x0f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0c, 0xca, 0xc6, 0x27, 0x08, 0x42, 0x06,
	0x0a, 0x02, 0x08, 0x01, 0x18, 0x01, 0x52, 0x0e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x65,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x44, 0x0a, 0x1b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x42, 0x04, 0xf0, 0xd9, 0x21,
	0x01, 0x52, 0x19, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x58, 0x0a, 0x0d,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c,
	0xca, 0xc6, 0x27, 0x08, 0x2a, 0x06, 0x3a, 0x02, 0x08, 0x40, 0x68, 0x01, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x3a, 0x8d, 0x01, 0xea, 0x41, 0x42, 0x0a, 0x17, 0x6d, 0x65,
	0x74, 0x61, 0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x27, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x7b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x7d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x7d, 0x92, 0xd9,
	0x21, 0x32, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x09, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2a, 0x0f, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x5d, 0x7b, 0x31, 0x2c, 0x31, 0x32,
	0x38, 0x7d, 0x38, 0x05, 0xda, 0x94, 0x23, 0x08, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0xe2, 0xde, 0x21, 0x02, 0x08, 0x01, 0x42, 0x92, 0x02, 0xe8, 0xde, 0x21, 0x01, 0xd2, 0xff, 0xd0,
	0x02, 0x4a, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x12, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0xa2, 0x80, 0xd1, 0x02,
	0x4c, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x0a, 0x14, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x62,
	0x2e, 0x76, 0x31, 0x42, 0x0d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f,
	0x6d, 0x65, 0x74, 0x61, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	goten_meta_service_proto_v1_resource_proto_rawDescOnce sync.Once
	goten_meta_service_proto_v1_resource_proto_rawDescData = goten_meta_service_proto_v1_resource_proto_rawDesc
)

func goten_meta_service_proto_v1_resource_proto_rawDescGZIP() []byte {
	goten_meta_service_proto_v1_resource_proto_rawDescOnce.Do(func() {
		goten_meta_service_proto_v1_resource_proto_rawDescData = protoimpl.X.CompressGZIP(goten_meta_service_proto_v1_resource_proto_rawDescData)
	})
	return goten_meta_service_proto_v1_resource_proto_rawDescData
}

var goten_meta_service_proto_v1_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var goten_meta_service_proto_v1_resource_proto_goTypes = []interface{}{
	(*Resource)(nil),               // 0: goten.meta.v1.Resource
	(*Resource_VersionedInfo)(nil), // 1: goten.meta.v1.Resource.VersionedInfo
	(*meta.Meta)(nil),              // 2: goten.types.Meta
}
var goten_meta_service_proto_v1_resource_proto_depIdxs = []int32{
	2, // 0: goten.meta.v1.Resource.metadata:type_name -> goten.types.Meta
	1, // 1: goten.meta.v1.Resource.versioned_infos:type_name -> goten.meta.v1.Resource.VersionedInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { goten_meta_service_proto_v1_resource_proto_init() }
func goten_meta_service_proto_v1_resource_proto_init() {
	if goten_meta_service_proto_v1_resource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		goten_meta_service_proto_v1_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		goten_meta_service_proto_v1_resource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource_VersionedInfo); i {
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
			RawDescriptor: goten_meta_service_proto_v1_resource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           goten_meta_service_proto_v1_resource_proto_goTypes,
		DependencyIndexes: goten_meta_service_proto_v1_resource_proto_depIdxs,
		MessageInfos:      goten_meta_service_proto_v1_resource_proto_msgTypes,
	}.Build()
	goten_meta_service_proto_v1_resource_proto = out.File
	goten_meta_service_proto_v1_resource_proto_rawDesc = nil
	goten_meta_service_proto_v1_resource_proto_goTypes = nil
	goten_meta_service_proto_v1_resource_proto_depIdxs = nil
}