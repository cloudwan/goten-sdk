package client

import "fmt"

var (
	globalRegistry *Registry
)

// Registry contains all known service descriptors and allows accessing
// them by domain name and version.
// It is responsibility of developer to import all resource packages in order to
// ensure their registration.
type Registry struct {
	methodDescriptors map[string]MethodDescriptor
	apiDescriptors    map[string]ApiDescriptor
	svcDescriptors    map[string]ServiceDescriptor
}

// FindMethodDescriptor gives method descriptor based on full method name, for example
// "/ntt.iam.v1alpha.ProjectService/ListProjects" - can be obtained from grpc.UnaryServerInfo
// or grpc.StreamServerInfo
func (r *Registry) FindMethodDescriptor(fullMethodName string) MethodDescriptor {
	return r.methodDescriptors[fullMethodName]
}

// FindApiDescriptor gives API descriptor based on full api name, for example
// "/ntt.iam.v1alpha.ProjectService" - basically this is same as full method name from
// grpc.UnaryServerInfo or grpc.StreamServerInfo without method name and trailing '/'
func (r *Registry) FindApiDescriptor(fullApiName string) ApiDescriptor {
	return r.apiDescriptors[fullApiName]
}

func (r *Registry) FindAllVersionsOfServiceDescriptors(serviceDomain string) []ServiceDescriptor {
	svcDescriptors := make([]ServiceDescriptor, 0)
	for _, svcDescriptor := range r.svcDescriptors {
		if svcDescriptor.GetServiceDomain() == serviceDomain {
			svcDescriptors = append(svcDescriptors, svcDescriptor)
		}
	}
	return svcDescriptors
}

func (r *Registry) FindNewestVersionServiceDescriptor(serviceDomain string) ServiceDescriptor {
	for _, svcDescriptor := range r.svcDescriptors {
		if svcDescriptor.GetServiceDomain() == serviceDomain && svcDescriptor.GetNextVersion() == "" {
			return svcDescriptor
		}
	}
	return nil
}

func (r *Registry) FindVersionedServiceDescriptor(serviceDomain, version string) ServiceDescriptor {
	return r.svcDescriptors[fmt.Sprintf("%s/%s", serviceDomain, version)]
}

func (r *Registry) RegisterMethodDescriptor(descriptor MethodDescriptor) {
	r.methodDescriptors[descriptor.GetFullMethodName()] = descriptor
}

func (r *Registry) RegisterApiDescriptor(descriptor ApiDescriptor) {
	r.apiDescriptors[descriptor.GetFullAPIName()] = descriptor
}

func (r *Registry) RegisterSvcDescriptor(descriptor ServiceDescriptor) {
	r.svcDescriptors[fmt.Sprintf("%s/%s", descriptor.GetServiceDomain(), descriptor.GetVersion())] = descriptor
}

// GetRegistry returns global registry instance
func GetRegistry() *Registry {
	ensureRegistryExists()
	return globalRegistry
}

func ensureRegistryExists() {
	if globalRegistry == nil {
		globalRegistry = &Registry{
			methodDescriptors: make(map[string]MethodDescriptor),
			apiDescriptors:    make(map[string]ApiDescriptor),
			svcDescriptors:    make(map[string]ServiceDescriptor),
		}
	}
}

func init() {
	ensureRegistryExists()
}
