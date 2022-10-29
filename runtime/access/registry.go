package access

import (
	"time"

	timestamppb "github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"

	"github.com/cloudwan/goten-sdk/runtime/api/view"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	"github.com/cloudwan/goten-sdk/runtime/object"
	"github.com/cloudwan/goten-sdk/runtime/resource"
)

var (
	globalRegistry *Registry
)

// Registry contains all known builders of resource-oriented API access objects.
// It is responsibility of developer to import all resource access packages in order to
// ensure FindApiAccessBuilder does not return nil object.
type Registry struct {
	entries map[resource.Descriptor]*ApiAccessBuilder
}

func (r *Registry) FindApiAccessBuilder(descriptor resource.Descriptor) *ApiAccessBuilder {
	return r.entries[descriptor]
}

func (r *Registry) RegisterWatcherConstructor(descriptor resource.Descriptor, constructor WatcherConstructor) {
	r.getOrMakeEntry(descriptor).watcherConstructor = constructor
}

func (r *Registry) RegisterQueryWatcherConstructor(descriptor resource.Descriptor, constructor QueryWatcherConstructor) {
	r.getOrMakeEntry(descriptor).queryWatcherConstructor = constructor
}

func (r *Registry) RegisterWatcherFilterConstructor(descriptor resource.Descriptor, constructor WatcherFilterConstructor) {
	r.getOrMakeEntry(descriptor).watcherFilterConstructor = constructor
}

func (r *Registry) RegisterApiAccessConstructor(descriptor resource.Descriptor, constructor ApiAccessConstructor) {
	r.getOrMakeEntry(descriptor).apiAccessConstructor = constructor
}

func (r *Registry) getOrMakeEntry(descriptor resource.Descriptor) *ApiAccessBuilder {
	entry := r.entries[descriptor]
	if entry == nil {
		entry = &ApiAccessBuilder{descriptor: descriptor}
		r.entries[descriptor] = entry
	}
	return entry
}

// GetRegistry returns global registry instance
func GetRegistry() *Registry {
	ensureRegistryExists()
	return globalRegistry
}

func ensureRegistryExists() {
	if globalRegistry == nil {
		globalRegistry = &Registry{
			entries: make(map[resource.Descriptor]*ApiAccessBuilder),
		}
	}
}

func init() {
	ensureRegistryExists()
}

type WatcherConfigParams struct {
	CfgBase    *WatcherConfigBase
	FieldMask  object.FieldMask
	OrderBy    resource.OrderBy
	WatchType  watch_type.WatchType
	View       view.View
	ChunkSize  int
}

type QueryWatcherConfigParams struct {
	Parent           resource.Reference
	Filter           resource.Filter
	Cursor           resource.Cursor
	FieldMask        object.FieldMask
	OrderBy          resource.OrderBy
	WatchType        watch_type.WatchType
	View             view.View
	ChunkSize        int
	PageSize         int
	StartingTime     *timestamppb.Timestamp
	RecoveryDeadline time.Duration
	RetryTimeout     time.Duration
}

type QueryWatcherConstructor func(id int, grpcConn grpc.ClientConnInterface, params *QueryWatcherConfigParams, ch chan QueryWatcherEvent) QueryWatcher

type WatcherConstructor func(grpcConn grpc.ClientConnInterface, params *WatcherConfigParams, filters ...WatcherFilterParams) Watcher

type WatcherFilterConstructor func(filter resource.Filter, parentRef resource.Reference) WatcherFilterParams

type ApiAccessConstructor func(grpcConn grpc.ClientConnInterface) resource.Access

// ApiAccessBuilder is a helper object allowing constructing various API-based and resource-oriented access objects
// without explicit resource type.
type ApiAccessBuilder struct {
	descriptor               resource.Descriptor
	watcherConstructor       WatcherConstructor
	queryWatcherConstructor  QueryWatcherConstructor
	watcherFilterConstructor WatcherFilterConstructor
	apiAccessConstructor     ApiAccessConstructor
}

func (re *ApiAccessBuilder) GetDescriptor() resource.Descriptor {
	return re.descriptor
}

func (re *ApiAccessBuilder) MakeWatcherFilterParams(filter resource.Filter, parentRef resource.Reference) WatcherFilterParams {
	return re.watcherFilterConstructor(filter, parentRef)
}

func (re *ApiAccessBuilder) MakeWatcher(grpcConn grpc.ClientConnInterface, params *WatcherConfigParams, filters ...WatcherFilterParams) Watcher {
	return re.watcherConstructor(grpcConn, params, filters...)
}

func (re *ApiAccessBuilder) MakeQueryWatcher(id int, grpcConn grpc.ClientConnInterface, params *QueryWatcherConfigParams, ch chan QueryWatcherEvent) QueryWatcher {
	return re.queryWatcherConstructor(id, grpcConn, params, ch)
}

func (re *ApiAccessBuilder) MakeApiAccess(grpcConn grpc.ClientConnInterface) resource.Access {
	return re.apiAccessConstructor(grpcConn)
}
