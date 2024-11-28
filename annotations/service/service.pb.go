// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.25.2
// source: goten/annotations/service.proto

package service

import (
	goten "github.com/cloudwan/goten-sdk/annotations/goten"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Method contains meta-information about RPC method like what type of resource it works on,
// whether it operates on single instance, set or whole collection. It informs what are the
// field paths to resource name, body or collection. It also influences generated method descriptor
// and associated request/response handlers.
// This object is generated by goten-bootstrap tool based on action (bootstrap) parameters.
type Method struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource name
	Resource string `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// Action verb
	Verb string `protobuf:"bytes,2,opt,name=verb,proto3" json:"verb,omitempty"`
	// whether method operates on resource collection. This generally includes
	// methods like List or Search (all those with "parent" field). It also
	// contains Create methods, as act of creation concerns whole collection
	// (like checking for uniqueness). It does not include batch versions
	// of single-resource actions (like batch get).
	IsCollection bool `protobuf:"varint,3,opt,name=is_collection,json=isCollection,proto3" json:"is_collection,omitempty"`
	// whether method operates on multiple instances of resource - List, Search,
	// BatchGet...
	IsPlural bool `protobuf:"varint,4,opt,name=is_plural,json=isPlural,proto3" json:"is_plural,omitempty"`
	// List of interesting field paths in request object
	RequestPaths *Method_ObjectFieldPaths `protobuf:"bytes,5,opt,name=request_paths,json=requestPaths,proto3" json:"request_paths,omitempty"`
	// List of interesting field paths in response object
	ResponsePaths *Method_ObjectFieldPaths `protobuf:"bytes,6,opt,name=response_paths,json=responsePaths,proto3" json:"response_paths,omitempty"`
}

func (x *Method) Reset() {
	*x = Method{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goten_annotations_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Method) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Method) ProtoMessage() {}

func (x *Method) ProtoReflect() protoreflect.Message {
	mi := &file_goten_annotations_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Method.ProtoReflect.Descriptor instead.
func (*Method) Descriptor() ([]byte, []int) {
	return file_goten_annotations_service_proto_rawDescGZIP(), []int{0}
}

func (x *Method) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *Method) GetVerb() string {
	if x != nil {
		return x.Verb
	}
	return ""
}

func (x *Method) GetIsCollection() bool {
	if x != nil {
		return x.IsCollection
	}
	return false
}

func (x *Method) GetIsPlural() bool {
	if x != nil {
		return x.IsPlural
	}
	return false
}

func (x *Method) GetRequestPaths() *Method_ObjectFieldPaths {
	if x != nil {
		return x.RequestPaths
	}
	return nil
}

func (x *Method) GetResponsePaths() *Method_ObjectFieldPaths {
	if x != nil {
		return x.ResponsePaths
	}
	return nil
}

// Goten service package is a bundle of resources and RPC services that binds
// everything together. It is an extension of proto package -> Proto package can
// have only one instance of ServicePackage defined. All resources and RPC
// services found within proto package are assigned automatically to this
// ServicePackage.
type ServicePackage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Simple service name, for example IAM, Devices, etc.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Service name, for example library.goten.com
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// Current version of the service
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// Name of a proto package which contains newer version of current
	// ServicePackage
	NextVersionPkg string `protobuf:"bytes,4,opt,name=next_version_pkg,json=nextVersionPkg,proto3" json:"next_version_pkg,omitempty"`
	// All imported service packages. For example, library.goten.com may require
	// book.registry.goten.com.
	ImportedServices []*ServicePackage_SvcImport `protobuf:"bytes,5,rep,name=imported_services,json=importedServices,proto3" json:"imported_services,omitempty"`
	// If multi-region support is disabled
	DisableMultiRegion bool `protobuf:"varint,6,opt,name=disable_multi_region,json=disableMultiRegion,proto3" json:"disable_multi_region,omitempty"`
}

func (x *ServicePackage) Reset() {
	*x = ServicePackage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goten_annotations_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServicePackage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServicePackage) ProtoMessage() {}

func (x *ServicePackage) ProtoReflect() protoreflect.Message {
	mi := &file_goten_annotations_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServicePackage.ProtoReflect.Descriptor instead.
func (*ServicePackage) Descriptor() ([]byte, []int) {
	return file_goten_annotations_service_proto_rawDescGZIP(), []int{1}
}

func (x *ServicePackage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServicePackage) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *ServicePackage) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ServicePackage) GetNextVersionPkg() string {
	if x != nil {
		return x.NextVersionPkg
	}
	return ""
}

func (x *ServicePackage) GetImportedServices() []*ServicePackage_SvcImport {
	if x != nil {
		return x.ImportedServices
	}
	return nil
}

func (x *ServicePackage) GetDisableMultiRegion() bool {
	if x != nil {
		return x.DisableMultiRegion
	}
	return false
}

type Method_ObjectFieldPaths struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// All field paths pointing to resource name (or names for plural resource actions).
	ResourceName []string `protobuf:"bytes,1,rep,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// All field paths pointing to resource body (or bodies for plural resource actions).
	// In certain circumstances, it may contain resource change bodies.
	ResourceBody []string `protobuf:"bytes,2,rep,name=resource_body,json=resourceBody,proto3" json:"resource_body,omitempty"`
	// All field paths pointing to resource collection. This field is valid only
	// for collection actions only.
	ResourceParent []string `protobuf:"bytes,3,rep,name=resource_parent,json=resourceParent,proto3" json:"resource_parent,omitempty"`
}

func (x *Method_ObjectFieldPaths) Reset() {
	*x = Method_ObjectFieldPaths{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goten_annotations_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Method_ObjectFieldPaths) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Method_ObjectFieldPaths) ProtoMessage() {}

func (x *Method_ObjectFieldPaths) ProtoReflect() protoreflect.Message {
	mi := &file_goten_annotations_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Method_ObjectFieldPaths.ProtoReflect.Descriptor instead.
func (*Method_ObjectFieldPaths) Descriptor() ([]byte, []int) {
	return file_goten_annotations_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Method_ObjectFieldPaths) GetResourceName() []string {
	if x != nil {
		return x.ResourceName
	}
	return nil
}

func (x *Method_ObjectFieldPaths) GetResourceBody() []string {
	if x != nil {
		return x.ResourceBody
	}
	return nil
}

func (x *Method_ObjectFieldPaths) GetResourceParent() []string {
	if x != nil {
		return x.ResourceParent
	}
	return nil
}

// SvcImport describes imported service
type ServicePackage_SvcImport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Service domain, for example library.goten.com
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// Proto package where service is defined
	ProtoPkg string `protobuf:"bytes,2,opt,name=proto_pkg,json=protoPkg,proto3" json:"proto_pkg,omitempty"`
	// Version of the service
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ServicePackage_SvcImport) Reset() {
	*x = ServicePackage_SvcImport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goten_annotations_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServicePackage_SvcImport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServicePackage_SvcImport) ProtoMessage() {}

func (x *ServicePackage_SvcImport) ProtoReflect() protoreflect.Message {
	mi := &file_goten_annotations_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServicePackage_SvcImport.ProtoReflect.Descriptor instead.
func (*ServicePackage_SvcImport) Descriptor() ([]byte, []int) {
	return file_goten_annotations_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ServicePackage_SvcImport) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *ServicePackage_SvcImport) GetProtoPkg() string {
	if x != nil {
		return x.ProtoPkg
	}
	return ""
}

func (x *ServicePackage_SvcImport) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var file_goten_annotations_service_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*Method)(nil),
		Field:         69040,
		Name:          "goten.annotations.method",
		Tag:           "bytes,69040,opt,name=method",
		Filename:      "goten/annotations/service.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*goten.GoPackage)(nil),
		Field:         690160,
		Name:          "goten.annotations.server_go_package",
		Tag:           "bytes,690160,opt,name=server_go_package",
		Filename:      "goten/annotations/service.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*goten.GoPackage)(nil),
		Field:         690161,
		Name:          "goten.annotations.client_go_package",
		Tag:           "bytes,690161,opt,name=client_go_package",
		Filename:      "goten/annotations/service.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*goten.GoPackage)(nil),
		Field:         690370,
		Name:          "goten.annotations.versioning_go_package",
		Tag:           "bytes,690370,opt,name=versioning_go_package",
		Filename:      "goten/annotations/service.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*goten.GoPackage)(nil),
		Field:         690180,
		Name:          "goten.annotations.access_go_package",
		Tag:           "bytes,690180,opt,name=access_go_package",
		Filename:      "goten/annotations/service.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*goten.GoPackage)(nil),
		Field:         690120,
		Name:          "goten.annotations.cli_go_package",
		Tag:           "bytes,690120,opt,name=cli_go_package",
		Filename:      "goten/annotations/service.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*goten.GoPackage)(nil),
		Field:         690170,
		Name:          "goten.annotations.store_go_package",
		Tag:           "bytes,690170,opt,name=store_go_package",
		Filename:      "goten/annotations/service.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*ServicePackage)(nil),
		Field:         690162,
		Name:          "goten.annotations.service_pkg",
		Tag:           "bytes,690162,opt,name=service_pkg",
		Filename:      "goten/annotations/service.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional goten.annotations.Method method = 69040;
	E_Method = &file_goten_annotations_service_proto_extTypes[0]
)

// Extension fields to descriptorpb.FileOptions.
var (
	// Placement for server side Go packages
	//
	// optional goten.annotations.GoPackage server_go_package = 690160;
	E_ServerGoPackage = &file_goten_annotations_service_proto_extTypes[1]
	// Placement for client side Go packages
	//
	// optional goten.annotations.GoPackage client_go_package = 690161;
	E_ClientGoPackage = &file_goten_annotations_service_proto_extTypes[2]
	// Placement for versioning Go packages (if any)
	//
	// optional goten.annotations.GoPackage versioning_go_package = 690370;
	E_VersioningGoPackage = &file_goten_annotations_service_proto_extTypes[3]
	// Placement for API Access related Go packages
	//
	// optional goten.annotations.GoPackage access_go_package = 690180;
	E_AccessGoPackage = &file_goten_annotations_service_proto_extTypes[4]
	// Package for CLI related Go packages
	//
	// optional goten.annotations.GoPackage cli_go_package = 690120;
	E_CliGoPackage = &file_goten_annotations_service_proto_extTypes[5]
	// Package for Store Access related Go packages
	//
	// optional goten.annotations.GoPackage store_go_package = 690170;
	E_StoreGoPackage = &file_goten_annotations_service_proto_extTypes[6]
	// Specification of a ServicePackage. This option can be defined only once in
	// one file within whole proto package. It must be defined if proto package
	// describes goten service.
	//
	// optional goten.annotations.ServicePackage service_pkg = 690162;
	E_ServicePkg = &file_goten_annotations_service_proto_extTypes[7]
)

var File_goten_annotations_service_proto protoreflect.FileDescriptor

var file_goten_annotations_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x03, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x76, 0x65, 0x72, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x76, 0x65, 0x72, 0x62,
	0x12, 0x23, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x70, 0x6c, 0x75, 0x72,
	0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x50, 0x6c, 0x75, 0x72,
	0x61, 0x6c, 0x12, 0x4f, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x50, 0x61, 0x74, 0x68, 0x73, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61,
	0x74, 0x68, 0x73, 0x12, 0x51, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x50, 0x61, 0x74, 0x68, 0x73, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x73, 0x1a, 0x85, 0x01, 0x0a, 0x10, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x61, 0x74, 0x68, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x22, 0xe8,
	0x02, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x65, 0x78, 0x74, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6b, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x6e, 0x65, 0x78, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6b,
	0x67, 0x12, 0x58, 0x0a, 0x11, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e,
	0x53, 0x76, 0x63, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x10, 0x69, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x1a, 0x5a, 0x0a,
	0x09, 0x53, 0x76, 0x63, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x70, 0x6b, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6b, 0x67, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x53, 0x0a, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xb0, 0x9b, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x3a, 0x68,
	0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xf0, 0x8f, 0x2a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x6f,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47,
	0x6f, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x3a, 0x68, 0x0a, 0x11, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf1, 0x8f, 0x2a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x6f, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x52, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x6f, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x3a, 0x70, 0x0a, 0x15, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67,
	0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc2, 0x91, 0x2a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x6f, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52,
	0x13, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x47, 0x6f, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x3a, 0x68, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x67,
	0x6f, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x84, 0x90, 0x2a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x6f, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x0f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x47, 0x6f, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x3a, 0x62,
	0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc8,
	0x8f, 0x2a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x6f, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x47, 0x6f, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x3a, 0x66, 0x0a, 0x10, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x67, 0x6f, 0x5f, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xfa, 0x8f, 0x2a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x47, 0x6f, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x0e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x47, 0x6f, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x3a, 0x62, 0x0a, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x6b, 0x67, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf2, 0x8f, 0x2a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6b, 0x67, 0x42, 0x67,
	0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x42, 0x11, 0x47, 0x6f, 0x74, 0x65, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x01, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goten_annotations_service_proto_rawDescOnce sync.Once
	file_goten_annotations_service_proto_rawDescData = file_goten_annotations_service_proto_rawDesc
)

func file_goten_annotations_service_proto_rawDescGZIP() []byte {
	file_goten_annotations_service_proto_rawDescOnce.Do(func() {
		file_goten_annotations_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_goten_annotations_service_proto_rawDescData)
	})
	return file_goten_annotations_service_proto_rawDescData
}

var file_goten_annotations_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_goten_annotations_service_proto_goTypes = []interface{}{
	(*Method)(nil),                     // 0: goten.annotations.Method
	(*ServicePackage)(nil),             // 1: goten.annotations.ServicePackage
	(*Method_ObjectFieldPaths)(nil),    // 2: goten.annotations.Method.ObjectFieldPaths
	(*ServicePackage_SvcImport)(nil),   // 3: goten.annotations.ServicePackage.SvcImport
	(*descriptorpb.MethodOptions)(nil), // 4: google.protobuf.MethodOptions
	(*descriptorpb.FileOptions)(nil),   // 5: google.protobuf.FileOptions
	(*goten.GoPackage)(nil),            // 6: goten.annotations.GoPackage
}
var file_goten_annotations_service_proto_depIdxs = []int32{
	2,  // 0: goten.annotations.Method.request_paths:type_name -> goten.annotations.Method.ObjectFieldPaths
	2,  // 1: goten.annotations.Method.response_paths:type_name -> goten.annotations.Method.ObjectFieldPaths
	3,  // 2: goten.annotations.ServicePackage.imported_services:type_name -> goten.annotations.ServicePackage.SvcImport
	4,  // 3: goten.annotations.method:extendee -> google.protobuf.MethodOptions
	5,  // 4: goten.annotations.server_go_package:extendee -> google.protobuf.FileOptions
	5,  // 5: goten.annotations.client_go_package:extendee -> google.protobuf.FileOptions
	5,  // 6: goten.annotations.versioning_go_package:extendee -> google.protobuf.FileOptions
	5,  // 7: goten.annotations.access_go_package:extendee -> google.protobuf.FileOptions
	5,  // 8: goten.annotations.cli_go_package:extendee -> google.protobuf.FileOptions
	5,  // 9: goten.annotations.store_go_package:extendee -> google.protobuf.FileOptions
	5,  // 10: goten.annotations.service_pkg:extendee -> google.protobuf.FileOptions
	0,  // 11: goten.annotations.method:type_name -> goten.annotations.Method
	6,  // 12: goten.annotations.server_go_package:type_name -> goten.annotations.GoPackage
	6,  // 13: goten.annotations.client_go_package:type_name -> goten.annotations.GoPackage
	6,  // 14: goten.annotations.versioning_go_package:type_name -> goten.annotations.GoPackage
	6,  // 15: goten.annotations.access_go_package:type_name -> goten.annotations.GoPackage
	6,  // 16: goten.annotations.cli_go_package:type_name -> goten.annotations.GoPackage
	6,  // 17: goten.annotations.store_go_package:type_name -> goten.annotations.GoPackage
	1,  // 18: goten.annotations.service_pkg:type_name -> goten.annotations.ServicePackage
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	11, // [11:19] is the sub-list for extension type_name
	3,  // [3:11] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_goten_annotations_service_proto_init() }
func file_goten_annotations_service_proto_init() {
	if File_goten_annotations_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goten_annotations_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Method); i {
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
		file_goten_annotations_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServicePackage); i {
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
		file_goten_annotations_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Method_ObjectFieldPaths); i {
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
		file_goten_annotations_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServicePackage_SvcImport); i {
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
			RawDescriptor: file_goten_annotations_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 8,
			NumServices:   0,
		},
		GoTypes:           file_goten_annotations_service_proto_goTypes,
		DependencyIndexes: file_goten_annotations_service_proto_depIdxs,
		MessageInfos:      file_goten_annotations_service_proto_msgTypes,
		ExtensionInfos:    file_goten_annotations_service_proto_extTypes,
	}.Build()
	File_goten_annotations_service_proto = out.File
	file_goten_annotations_service_proto_rawDesc = nil
	file_goten_annotations_service_proto_goTypes = nil
	file_goten_annotations_service_proto_depIdxs = nil
}
