package client

// ApiDescriptor allows writing code operating on Goten API without knowing exact type.
// It can be used to access all methods and their properties, use grpc.ClientConnInterface
// without knowing exact types.
// ApiDescriptor corresponds to single Goten-API service (represented by single GRPC proto definition)
// in specific proto file.
type ApiDescriptor interface {
	// GetFullAPIName returns fully qualified name that is same as used as PREFIX in
	// grpc.UnaryStreamInfo (or stream equivalent) objects for all , for example
	// /goten.library.v2.BookService
	GetFullAPIName() string

	// GetProtoPkgName returns proto package where corresponding RPC service is defined,
	// for example "goten.library.v2"
	GetProtoPkgName() string

	// GetApiName returns Goten API name (GRPC service as defined in proto file), for example BookService
	GetApiName() string

	// GetServiceDomain returns domain of Goten service, as defined by field "name" in api-skeleton yaml file,
	// for example library.goten.com
	GetServiceDomain() string

	// GetServiceVersion returns version of Goten service, as defined by field "proto.package.currentVersion"
	// field in api-skeleton yaml file, for example "v1"
	GetServiceVersion() string

	// AllMethodDescriptors returns all method descriptors
	AllMethodDescriptors() []MethodDescriptor
}
