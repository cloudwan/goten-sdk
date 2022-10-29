package access

import (
	"context"

	"github.com/cloudwan/goten-sdk/runtime/resource"
)

type QueryWatcher interface {
	QueryWatcher()
	Run(ctx context.Context) error
}

type QueryWatcherEvent interface {
	GetWatcherIdentifier() int
	GetChanges() resource.ResourceChangeList
	IsReset() bool
	IsLostSync() bool
	IsSync() bool
	GetSnapshotSize() int64
	HasSnapshotSize() bool
}
