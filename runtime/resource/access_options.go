package resource

type saveOptions struct {
	previousResource Resource
}

type deleteOptions struct {
	currentResource Resource
}

type batchGetOptions struct {
	mustResolveAll bool
}

type SaveOptions interface {
	GetPreviousResource() Resource
}

type DeleteOptions interface {
	GetDeletedResource() Resource
}

type BatchGetOptions interface {
	MustResolveAll() bool
}

func (so *saveOptions) GetPreviousResource() Resource {
	return so.previousResource
}

func (do *deleteOptions) GetDeletedResource() Resource {
	return do.currentResource
}

func (bgo *batchGetOptions) MustResolveAll() bool {
	return bgo.mustResolveAll
}

type SaveOption func(*saveOptions)

type DeleteOption func(*deleteOptions)

type BatchGetOption func(*batchGetOptions)

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

func WithPreviousResource(previous Resource) SaveOption {
	return func(o *saveOptions) {
		o.previousResource = previous
	}
}

func WithCurrentResource(current Resource) DeleteOption {
	return func(o *deleteOptions) {
		o.currentResource = current
	}
}

func WithMustResolveAll() BatchGetOption {
	return func(o *batchGetOptions) {
		o.mustResolveAll = true
	}
}
