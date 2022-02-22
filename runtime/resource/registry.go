package resource

var (
	globalRegistry *Registry
)

// Registry contains all known resource descriptors and allows accessing
// them by domain name and version.
// It is responsibility of developer to import all resource packages in order to
// ensure their registration.
type Registry struct {
	descriptors map[string]map[string]Descriptor
}

// FindResourceDescriptor gives resource descriptor based on fully qualified name and version
func (r *Registry) FindResourceDescriptor(fullyQualifiedName, version string) Descriptor {
	return r.descriptors[fullyQualifiedName][version]
}

// RegisterDescriptor is called by each imported resource package.
func (r *Registry) RegisterDescriptor(descriptor Descriptor) {
	typeName := descriptor.GetResourceTypeName()
	byVersion := r.descriptors[typeName.FullyQualifiedTypeName()]
	if byVersion == nil {
		byVersion = make(map[string]Descriptor)
		r.descriptors[typeName.FullyQualifiedTypeName()] = byVersion
	}
	byVersion[typeName.Version()] = descriptor
}

// GetRegistry returns global registry instance
func GetRegistry() *Registry {
	ensureRegistryExists()
	return globalRegistry
}

func ensureRegistryExists() {
	if globalRegistry == nil {
		globalRegistry = &Registry{
			descriptors: make(map[string]map[string]Descriptor),
		}
	}
}

func init() {
	ensureRegistryExists()
}
