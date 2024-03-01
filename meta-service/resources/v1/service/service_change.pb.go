// Code generated by protoc-gen-goten-go
// File: goten/meta-service/proto/v1/service_change.proto
// DO NOT EDIT!!!

package service

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
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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
	_ = &fieldmaskpb.FieldMask{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ServiceChange is used by Watch notifications Responses to describe change of
// single Service One of Added, Modified, Removed
type ServiceChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Service change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*ServiceChange_Added_
	//	*ServiceChange_Modified_
	//	*ServiceChange_Current_
	//	*ServiceChange_Removed_
	ChangeType isServiceChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *ServiceChange) Reset() {
	*m = ServiceChange{}
	if protoimpl.UnsafeEnabled {
		mi := &goten_meta_service_proto_v1_service_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ServiceChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ServiceChange) ProtoMessage() {}

func (m *ServiceChange) ProtoReflect() preflect.Message {
	mi := &goten_meta_service_proto_v1_service_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ServiceChange) GotenMessage() {}

// Deprecated, Use ServiceChange.ProtoReflect.Descriptor instead.
func (*ServiceChange) Descriptor() ([]byte, []int) {
	return goten_meta_service_proto_v1_service_change_proto_rawDescGZIP(), []int{0}
}

func (m *ServiceChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ServiceChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ServiceChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ServiceChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isServiceChange_ChangeType interface {
	isServiceChange_ChangeType()
}

type ServiceChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *ServiceChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof" firestore:"added"`
}
type ServiceChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *ServiceChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof" firestore:"modified"`
}
type ServiceChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *ServiceChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof" firestore:"current"`
}
type ServiceChange_Removed_ struct {
	// Removed is returned when Service is deleted or leaves Query view
	Removed *ServiceChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof" firestore:"removed"`
}

func (*ServiceChange_Added_) isServiceChange_ChangeType()    {}
func (*ServiceChange_Modified_) isServiceChange_ChangeType() {}
func (*ServiceChange_Current_) isServiceChange_ChangeType()  {}
func (*ServiceChange_Removed_) isServiceChange_ChangeType()  {}
func (m *ServiceChange) GetChangeType() isServiceChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *ServiceChange) GetAdded() *ServiceChange_Added {
	if x, ok := m.GetChangeType().(*ServiceChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *ServiceChange) GetModified() *ServiceChange_Modified {
	if x, ok := m.GetChangeType().(*ServiceChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *ServiceChange) GetCurrent() *ServiceChange_Current {
	if x, ok := m.GetChangeType().(*ServiceChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *ServiceChange) GetRemoved() *ServiceChange_Removed {
	if x, ok := m.GetChangeType().(*ServiceChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *ServiceChange) SetChangeType(ofv isServiceChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isServiceChange_ChangeType", "ServiceChange"))
	}
	m.ChangeType = ofv
}
func (m *ServiceChange) SetAdded(fv *ServiceChange_Added) {
	m.SetChangeType(&ServiceChange_Added_{Added: fv})
}
func (m *ServiceChange) SetModified(fv *ServiceChange_Modified) {
	m.SetChangeType(&ServiceChange_Modified_{Modified: fv})
}
func (m *ServiceChange) SetCurrent(fv *ServiceChange_Current) {
	m.SetChangeType(&ServiceChange_Current_{Current: fv})
}
func (m *ServiceChange) SetRemoved(fv *ServiceChange_Removed) {
	m.SetChangeType(&ServiceChange_Removed_{Removed: fv})
}

// Service has been added to query view
type ServiceChange_Added struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Service       *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty" firestore:"service"`
	// Integer describing index of added Service in resulting query view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *ServiceChange_Added) Reset() {
	*m = ServiceChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &goten_meta_service_proto_v1_service_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ServiceChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ServiceChange_Added) ProtoMessage() {}

func (m *ServiceChange_Added) ProtoReflect() preflect.Message {
	mi := &goten_meta_service_proto_v1_service_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ServiceChange_Added) GotenMessage() {}

// Deprecated, Use ServiceChange_Added.ProtoReflect.Descriptor instead.
func (*ServiceChange_Added) Descriptor() ([]byte, []int) {
	return goten_meta_service_proto_v1_service_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *ServiceChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ServiceChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ServiceChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ServiceChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ServiceChange_Added) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *ServiceChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *ServiceChange_Added) SetService(fv *Service) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Service", "ServiceChange_Added"))
	}
	m.Service = fv
}

func (m *ServiceChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "ServiceChange_Added"))
	}
	m.ViewIndex = fv
}

// Service changed some of it's fields - contains either full document or
// masked change
type ServiceChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified Service
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// New version of Service or masked difference, depending on mask_changes
	// instrumentation of issued [WatchServiceRequest] or [WatchServicesRequest]
	Service *Service `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty" firestore:"service"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *Service_FieldMask `protobuf:"bytes,3,opt,customtype=Service_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty" firestore:"fieldMask"`
	// Previous view index specifies previous position of modified Service.
	// When modification doesn't affect sorted order, value will remain
	// identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty" firestore:"previousViewIndex"`
	// Integer specifying Service new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *ServiceChange_Modified) Reset() {
	*m = ServiceChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &goten_meta_service_proto_v1_service_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ServiceChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ServiceChange_Modified) ProtoMessage() {}

func (m *ServiceChange_Modified) ProtoReflect() preflect.Message {
	mi := &goten_meta_service_proto_v1_service_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ServiceChange_Modified) GotenMessage() {}

// Deprecated, Use ServiceChange_Modified.ProtoReflect.Descriptor instead.
func (*ServiceChange_Modified) Descriptor() ([]byte, []int) {
	return goten_meta_service_proto_v1_service_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *ServiceChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ServiceChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ServiceChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ServiceChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ServiceChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ServiceChange_Modified) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *ServiceChange_Modified) GetFieldMask() *Service_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *ServiceChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *ServiceChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *ServiceChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "ServiceChange_Modified"))
	}
	m.Name = fv
}

func (m *ServiceChange_Modified) SetService(fv *Service) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Service", "ServiceChange_Modified"))
	}
	m.Service = fv
}

func (m *ServiceChange_Modified) SetFieldMask(fv *Service_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "ServiceChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *ServiceChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "ServiceChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *ServiceChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "ServiceChange_Modified"))
	}
	m.ViewIndex = fv
}

// Service has been added or modified in a query view. Version used for
// stateless watching
type ServiceChange_Current struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Service       *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty" firestore:"service"`
}

func (m *ServiceChange_Current) Reset() {
	*m = ServiceChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &goten_meta_service_proto_v1_service_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ServiceChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ServiceChange_Current) ProtoMessage() {}

func (m *ServiceChange_Current) ProtoReflect() preflect.Message {
	mi := &goten_meta_service_proto_v1_service_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ServiceChange_Current) GotenMessage() {}

// Deprecated, Use ServiceChange_Current.ProtoReflect.Descriptor instead.
func (*ServiceChange_Current) Descriptor() ([]byte, []int) {
	return goten_meta_service_proto_v1_service_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *ServiceChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ServiceChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ServiceChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ServiceChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ServiceChange_Current) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *ServiceChange_Current) SetService(fv *Service) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Service", "ServiceChange_Current"))
	}
	m.Service = fv
}

// Removed is returned when Service is deleted or leaves Query view
type ServiceChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Integer specifying removed Service index. Not populated in stateless
	// watch type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *ServiceChange_Removed) Reset() {
	*m = ServiceChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &goten_meta_service_proto_v1_service_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ServiceChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ServiceChange_Removed) ProtoMessage() {}

func (m *ServiceChange_Removed) ProtoReflect() preflect.Message {
	mi := &goten_meta_service_proto_v1_service_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ServiceChange_Removed) GotenMessage() {}

// Deprecated, Use ServiceChange_Removed.ProtoReflect.Descriptor instead.
func (*ServiceChange_Removed) Descriptor() ([]byte, []int) {
	return goten_meta_service_proto_v1_service_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *ServiceChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ServiceChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ServiceChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ServiceChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ServiceChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ServiceChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *ServiceChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "ServiceChange_Removed"))
	}
	m.Name = fv
}

func (m *ServiceChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "ServiceChange_Removed"))
	}
	m.ViewIndex = fv
}

var goten_meta_service_proto_v1_service_change_proto preflect.FileDescriptor

var goten_meta_service_proto_v1_service_change_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76,
	0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x29, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x06, 0x0a, 0x0d, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x05,
	0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x65, 0x64, 0x48,
	0x00, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x43, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x40, 0x0a,
	0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12,
	0x40, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x64, 0x1a, 0x58, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0xfc, 0x01, 0x0a, 0x08,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xb2, 0xda, 0x21, 0x0b, 0x0a, 0x09, 0x0a, 0x07,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4a, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42,
	0x0f, 0xb2, 0xda, 0x21, 0x0b, 0x32, 0x09, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x2e, 0x0a, 0x13, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f,
	0x75, 0x73, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x76,
	0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x3b, 0x0a, 0x07, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x4d, 0x0a, 0x07, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0f, 0xb2, 0xda, 0x21, 0x0b, 0x0a, 0x09, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65,
	0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x0d, 0x9a, 0xd9, 0x21, 0x09, 0x0a, 0x07, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x42, 0x75, 0xe8, 0xde, 0x21, 0x00, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31,
	0x42, 0x12, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	goten_meta_service_proto_v1_service_change_proto_rawDescOnce sync.Once
	goten_meta_service_proto_v1_service_change_proto_rawDescData = goten_meta_service_proto_v1_service_change_proto_rawDesc
)

func goten_meta_service_proto_v1_service_change_proto_rawDescGZIP() []byte {
	goten_meta_service_proto_v1_service_change_proto_rawDescOnce.Do(func() {
		goten_meta_service_proto_v1_service_change_proto_rawDescData = protoimpl.X.CompressGZIP(goten_meta_service_proto_v1_service_change_proto_rawDescData)
	})
	return goten_meta_service_proto_v1_service_change_proto_rawDescData
}

var goten_meta_service_proto_v1_service_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var goten_meta_service_proto_v1_service_change_proto_goTypes = []interface{}{
	(*ServiceChange)(nil),          // 0: goten.meta.v1.ServiceChange
	(*ServiceChange_Added)(nil),    // 1: goten.meta.v1.ServiceChange.Added
	(*ServiceChange_Modified)(nil), // 2: goten.meta.v1.ServiceChange.Modified
	(*ServiceChange_Current)(nil),  // 3: goten.meta.v1.ServiceChange.Current
	(*ServiceChange_Removed)(nil),  // 4: goten.meta.v1.ServiceChange.Removed
	(*Service)(nil),                // 5: goten.meta.v1.Service
	(*Service_FieldMask)(nil),      // 6: goten.meta.v1.Service_FieldMask
}
var goten_meta_service_proto_v1_service_change_proto_depIdxs = []int32{
	1, // 0: goten.meta.v1.ServiceChange.added:type_name -> goten.meta.v1.ServiceChange.Added
	2, // 1: goten.meta.v1.ServiceChange.modified:type_name -> goten.meta.v1.ServiceChange.Modified
	3, // 2: goten.meta.v1.ServiceChange.current:type_name -> goten.meta.v1.ServiceChange.Current
	4, // 3: goten.meta.v1.ServiceChange.removed:type_name -> goten.meta.v1.ServiceChange.Removed
	5, // 4: goten.meta.v1.ServiceChange.Added.service:type_name -> goten.meta.v1.Service
	5, // 5: goten.meta.v1.ServiceChange.Modified.service:type_name -> goten.meta.v1.Service
	6, // 6: goten.meta.v1.ServiceChange.Modified.field_mask:type_name -> goten.meta.v1.Service_FieldMask
	5, // 7: goten.meta.v1.ServiceChange.Current.service:type_name -> goten.meta.v1.Service
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { goten_meta_service_proto_v1_service_change_proto_init() }
func goten_meta_service_proto_v1_service_change_proto_init() {
	if goten_meta_service_proto_v1_service_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		goten_meta_service_proto_v1_service_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceChange); i {
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
		goten_meta_service_proto_v1_service_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceChange_Added); i {
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
		goten_meta_service_proto_v1_service_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceChange_Modified); i {
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
		goten_meta_service_proto_v1_service_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceChange_Current); i {
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
		goten_meta_service_proto_v1_service_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceChange_Removed); i {
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

	goten_meta_service_proto_v1_service_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ServiceChange_Added_)(nil),
		(*ServiceChange_Modified_)(nil),
		(*ServiceChange_Current_)(nil),
		(*ServiceChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: goten_meta_service_proto_v1_service_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           goten_meta_service_proto_v1_service_change_proto_goTypes,
		DependencyIndexes: goten_meta_service_proto_v1_service_change_proto_depIdxs,
		MessageInfos:      goten_meta_service_proto_v1_service_change_proto_msgTypes,
	}.Build()
	goten_meta_service_proto_v1_service_change_proto = out.File
	goten_meta_service_proto_v1_service_change_proto_rawDesc = nil
	goten_meta_service_proto_v1_service_change_proto_goTypes = nil
	goten_meta_service_proto_v1_service_change_proto_depIdxs = nil
}