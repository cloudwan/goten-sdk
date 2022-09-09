package client

import (
	"github.com/cloudwan/goten-sdk/runtime/resource"
)

type ServiceImportInfo struct {
	Domain  string
	Version string
}

// ServiceDescriptor allows writing code operating on Goten Service without knowing exact type.
// It can be used to access all APIs/Methods, their properties, all resource types.
// ServiceDescriptor corresponds to service definition in api-skeleton file (one per version).
type ServiceDescriptor interface {
	// GetServiceDomain returns domain of Goten service, as defined by field "name" in api-skeleton yaml file,
	// for example library.edgelq.com
	GetServiceDomain() string

	// GetVersion returns version of Goten service, as defined by field "proto.package.currentVersion"
	// field in api-skeleton yaml file, for example "v1"
	GetVersion() string

	// GetNextVersion returns next version of Goten service, as defined by field "proto.package.nextVersion"
	// field in api-skeleton yaml file, for example "v1"
	GetNextVersion() string

	// AllResourceDescriptors returns descriptors for all owned resources
	AllResourceDescriptors() []resource.Descriptor

	// AllApiDescriptors returns descriptors for all owned APIs
	AllApiDescriptors() []ApiDescriptor

	// AllImportedServiceInfos returns information about all imported services.
	// Registry can be used to obtain proper descriptor, as long as it was imported
	// (We do not enforce all the imports in generated code)
	AllImportedServiceInfos() []ServiceImportInfo
}
