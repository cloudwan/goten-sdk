package client

import (
	"github.com/cloudwan/goten-sdk/runtime/resource"
)

// ServiceDescriptor allows writing code operating on Goten Service without knowing exact type.
// It can be used to access all APIs/Methods, their properties, all resource types.
// ServiceDescriptor corresponds to service definition in api-skeleton file (one per version).
//
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

	AllResourceDescriptors() []resource.Descriptor
	AllApiDescriptors() []ApiDescriptor
}
