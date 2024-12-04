package resource

import (
	"github.com/cloudwan/goten-sdk/runtime/object"
)

type saveOptions struct {
	updateOnly  bool
	createOnly  bool
	updateMask  object.FieldMask
	compareMask object.FieldMask
	compareRes  Resource
}

type deleteOptions struct {
}

type getOptions struct {
	skipCache bool
}

type batchGetOptions struct {
	mustResolveAll bool
	masks          map[Descriptor]object.FieldMask
	skipCache      bool
}

type queryOptions struct {
	skipCache bool
}

type SaveOptions interface {
	OnlyCreate() bool
	OnlyUpdate() bool
	// DEPRECATED, returns always nil
	GetPreviousResource() Resource
	GetUpdateMask() object.FieldMask
	GetCAS() (object.FieldMask, Resource)
}

type DeleteOptions interface {
	// DEPRECATED, returns always nil
	GetDeletedResource() Resource
}

type GetOptions interface {
	GetSkipCache() bool
}

type BatchGetOptions interface {
	MustResolveAll() bool
	GetFieldMask(Descriptor) object.FieldMask
	GetFieldMasks() map[Descriptor]object.FieldMask
	GetSkipCache() bool
}

type QueryOptions interface {
	GetSkipCache() bool
}

func (so *saveOptions) GetPreviousResource() Resource {
	return nil
}

func (so *saveOptions) OnlyCreate() bool {
	return so.createOnly
}

func (so *saveOptions) OnlyUpdate() bool {
	return so.updateOnly
}

func (so *saveOptions) GetUpdateMask() object.FieldMask {
	return so.updateMask
}

func (so *saveOptions) GetCAS() (object.FieldMask, Resource) {
	return so.compareMask, so.compareRes
}

// DEPRECATED, always returns nil
func (do *deleteOptions) GetDeletedResource() Resource {
	return nil
}

func (opt *getOptions) GetSkipCache() bool {
	return opt.skipCache
}

func (opt *queryOptions) GetSkipCache() bool {
	return opt.skipCache
}

func (bgo *batchGetOptions) MustResolveAll() bool {
	return bgo.mustResolveAll
}

func (bgo *batchGetOptions) GetFieldMask(descriptor Descriptor) object.FieldMask {
	return bgo.masks[descriptor]
}

func (bgo *batchGetOptions) GetFieldMasks() map[Descriptor]object.FieldMask {
	return bgo.masks
}

func (bgo *batchGetOptions) GetSkipCache() bool {
	return bgo.skipCache
}

type SaveOption func(*saveOptions)

type DeleteOption func(*deleteOptions)

type BatchGetOption func(*batchGetOptions)

type GetOption func(*getOptions)

type QueryOption func(*queryOptions)

func MakeSaveOptions(opts []SaveOption) SaveOptions {
	sOpts := saveOptions{}
	for _, opt := range opts {
		opt(&sOpts)
	}
	return &sOpts
}

func MakeDeleteOptions(opts []DeleteOption) DeleteOptions {
	dOpts := deleteOptions{}
	for _, opt := range opts {
		opt(&dOpts)
	}
	return &dOpts
}

func MakeBatchGetOptions(opts []BatchGetOption) BatchGetOptions {
	dOpts := batchGetOptions{}
	for _, opt := range opts {
		opt(&dOpts)
	}
	return &dOpts
}

func MakeGetOptions(opts []GetOption) GetOptions {
	dOpts := getOptions{}
	for _, opt := range opts {
		opt(&dOpts)
	}
	return &dOpts
}

func MakeQueryOptions(opts []QueryOption) QueryOptions {
	dOpts := queryOptions{}
	for _, opt := range opts {
		opt(&dOpts)
	}
	return &dOpts
}

// DEPRECATED, has no effect
func WithPreviousResource(_ Resource) SaveOption {
	return func(o *saveOptions) {
	}
}

// DEPRECATED, has no effect
func WithCurrentResource(_ Resource) DeleteOption {
	return func(o *deleteOptions) {
	}
}

func WithMustResolveAll() BatchGetOption {
	return func(o *batchGetOptions) {
		o.mustResolveAll = true
	}
}

func WithBatchGetFieldMask(desc Descriptor, mask object.FieldMask) BatchGetOption {
	return func(o *batchGetOptions) {
		if o.masks == nil {
			o.masks = map[Descriptor]object.FieldMask{}
		}
		o.masks[desc] = mask
	}
}

func WithBatchGetSkipCache() BatchGetOption {
	return func(o *batchGetOptions) {
		o.skipCache = true
	}
}

func WithGetSkipCache() GetOption {
	return func(o *getOptions) {
		o.skipCache = true
	}
}

func WithQuerySkipCache() QueryOption {
	return func(o *queryOptions) {
		o.skipCache = true
	}
}

func WithCreateModeOnly() SaveOption {
	return func(o *saveOptions) {
		if o.updateOnly {
			panic("WithCreateModeOnly called after WithUpdateModeOnly")
		}
		o.createOnly = true
	}
}

func WithUpdateModeOnly() SaveOption {
	return func(o *saveOptions) {
		if o.createOnly {
			panic("WithUpdateModeOnly called after WithCreateModeOnly")
		}
		o.updateOnly = true
	}
}

func WithUpdateMask(mask object.FieldMask) SaveOption {
	return func(o *saveOptions) {
		if o.createOnly {
			panic("WithUpdateMask called after WithCreateModeOnly")
		}
		o.updateMask = mask
	}
}

func WithCompareAndSwap(state Resource, mask object.FieldMask) SaveOption {
	return func(o *saveOptions) {
		if o.createOnly {
			panic("WithCompareAndSwap called after WithCreateModeOnly")
		}
		o.updateOnly = true
		o.compareMask = mask
		o.compareRes = state
	}
}
