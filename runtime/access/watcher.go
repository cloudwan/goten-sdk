package access

import (
	"context"
	"time"

	"github.com/cloudwan/goten-sdk/runtime/resource"
)

// Watcher is a high-level component using watch collection API method. Unlike
// raw watch, Watcher is capable of managing multiple queries at once (achieving
// OR filter, not possible via raw API). It also handles in transparent way
// brief reconnections, hiding all recovery implementation details. It only
// communicates LostSync/Resync events to users in case of long-term connection
// failures.
// Although each Watcher implements this interface (allowing users to develop
// applications in abstraction of type if they need), each struct also has strong-typed
// methods specific to their underlying type
type Watcher interface {
	// IEvents returns an abstract channel of WatcherEvent instances.
	// It must not be called if its strong-typed version Events() was ever called.
	IEvents() <-chan WatcherEvent

	// InSync informs if watched collection is in sync.
	InSync() bool

	// GetIFilters returns all currently active filters.
	GetIFilters() []WatcherFilterParams

	// ResetIFilters resets active filters. It will trigger LostSync event.
	// It can be called anytime during watcher runtime.
	// Watcher will no longer send events related to old filters, but there
	// may be still old events already in the events queue. It is necessary to check
	// integer number (filter versions) returned by this function and compare
	// with filter version returned by any WatcherEvent.
	ResetIFilters(ctx context.Context, filters ...WatcherFilterParams) (int32, error)

	// Run starts watcher.
	Run(ctx context.Context) error
}

// WatcherEventChange is a higher-level object built on top of API resource.ResourceChange object
// decorated with extra information like previous resource (in case of modify) or contents of deleted resource
// (in case of deletion).
type WatcherEventChange interface {
	// IsAdd tells if resource is newly added to the observed collection
	IsAdd() bool

	// IsModify tells if resource has been modified within observed collection
	IsModify() bool

	// IsDelete tells if resource has been removed from collection.
	IsDelete() bool

	// GetRawName returns name regardless of type
	GetRawName() resource.Name

	// GetRawAdded returns non-nil resource if IsAdd returns true
	GetRawAdded() resource.Resource

	// GetRawPrevious returns non-nil resource if IsModify returns true
	GetRawPrevious() resource.Resource

	// GetRawDeleted returns non-nil resource if IsDelete returns true
	GetRawDeleted() resource.Resource

	// GetRawCurrent returns non-nil resource if IsAdd or IsModify returns true
	GetRawCurrent() resource.Resource
}

// WatcherEvent informs user of Watcher about events & changes happening on observed collection.
type WatcherEvent interface {
	// LostSync communicates issue like network connectivity problem
	LostSync() bool

	// Resync returns true if changes contains all data in current snapshot (Added changes)
	Resync() bool

	// GetRawAt returns change at given index, or nil if outside bounds
	GetRawAt(index int) WatcherEventChange

	// Length returns size of the change (always 0 if LostSync is true)
	Length() int

	// AppendRawChange adds additional change. Panic for incorrect type
	AppendRawChange(change WatcherEventChange)

	// FilterVersion returns version of filter that generated this event.
	FilterVersion() int32
}

// WatcherFilterParams is a wrap over resource filter and parent.
// Parent may not apply (by always nil) if resource does not have any parent.
type WatcherFilterParams interface {
	GetIParentRef() resource.Reference
	GetIFilter() resource.Filter
}

// WatcherConfigBase is a base struct for all actual implementations of WatcherConfig
type WatcherConfigBase struct {
	// Event buffer size
	WatcherEventBufferSize int

	// Time that watcher will wait before attempting to establish connection with server in the event of failure like
	// connectivity lost. This does not generate any event downstream. IF you need to receive notification about failure,
	// see RecoveryDeadline
	RetryTimeout time.Duration

	// In case of single failure of connectivity with server watcher will make attempts to recover
	// without communicating LostSync event - unless recovery exceeds specified RecoveryDeadline.
	// If watcher manages to recover eventually after deadline, next event will come with Resync flag set.
	// If you dont want to have any lost sync events and you want watcher handle it, set value to 0.
	RecoveryDeadline time.Duration
}

func NewWatcherConfigBase() *WatcherConfigBase {
	return &WatcherConfigBase{
		WatcherEventBufferSize: 1000,
		RetryTimeout:           time.Second * 15,
		RecoveryDeadline:       time.Minute * 5,
	}
}

// WatcherEventBase is a base struct for all actual implementations of WatcherEvent.
type WatcherEventBase struct {
	lostSync  bool
	resync    bool
	filterVer int32
}

func NewWatcherEventBase(resync bool, filterVer int32) WatcherEventBase {
	return WatcherEventBase{lostSync: false, resync: resync, filterVer: filterVer}
}

func NewWatcherEventBaseLostSync(filterVer int32) WatcherEventBase {
	return WatcherEventBase{lostSync: true, resync: false, filterVer: filterVer}
}

func (eb WatcherEventBase) LostSync() bool {
	return eb.lostSync
}

func (eb WatcherEventBase) Resync() bool {
	return eb.resync
}

func (eb WatcherEventBase) FilterVersion() int32 {
	return eb.filterVer
}
