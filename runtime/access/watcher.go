package access

import (
	"time"

	"github.com/cloudwan/goten-sdk/runtime/resource"
)

type WatcherConfig struct {
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

func NewWatcherConfig() *WatcherConfig {
	return &WatcherConfig{
		WatcherEventBufferSize: 1000,
		RetryTimeout:           time.Second * 15,
		RecoveryDeadline:       time.Minute * 5,
	}
}

// WatcherEventChange is a higher-level object built on top of API resource.ResourceChange object
// decorated with extra information like previous resource (in case of modify) or contents of deleted resource
// (in case of deletion).
type WatcherEventChange interface {
	IsAdd() bool
	IsModify() bool
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
}

type WatcherEventbase struct {
	lostSync bool
	resync   bool
}

func NewWatcherEvent(resync bool) WatcherEventbase {
	return WatcherEventbase{lostSync: false, resync: resync}
}

func NewWatcherEventLostSync() WatcherEventbase {
	return WatcherEventbase{lostSync: true, resync: false}
}

func (eb WatcherEventbase) LostSync() bool {
	return eb.lostSync
}

func (eb WatcherEventbase) Resync() bool {
	return eb.resync
}
