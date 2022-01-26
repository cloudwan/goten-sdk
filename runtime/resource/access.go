package resource

import (
	"context"
)

// Access is an interface allowing basic interaction with set of supported resources
type Access interface {
	// Get returns single resource. Raises NotFound grpc error if
	// resource is not found.
	Get(ctx context.Context, q GetQuery) (Resource, error)

	// Query returns snapshot of resources within given query spec.
	Query(ctx context.Context, q ListQuery) (QueryResultSnapshot, error)

	// Search is like Query, but also provides results based on text-search
	Search(ctx context.Context, q SearchQuery) (SearchQueryResultSnapshot, error)

	// Watch is blocking call till is finished or error occurred.
	// Initially, provides ResourceChange of Add type containing full
	// resource state (minus provided field mask). Then it contains single
	// changes only.
	Watch(ctx context.Context, q GetQuery, cb func(ch ResourceChange) error) error

	// WatchQuery is blocking call till is finished or error occurred.
	// Initially provides snapshot of query like would be returned for Query.
	// Then it returns changes as they happen.
	WatchQuery(ctx context.Context, q WatchQuery, cb func(ch QueryResultChange) error) error

	// Save creates/updates single resource
	Save(ctx context.Context, res Resource, opts ...SaveOption) error

	// Delete removes single resource
	Delete(ctx context.Context, ref Reference, opts ...DeleteOption) error

	// BatchGet resolves passed references all at once.
	// References pointing to non-existent resources won't be resolved.
	// In order to ensure all references are resolved, use WithMustResolveAll
	// option
	BatchGet(ctx context.Context, toGet []Reference, opts ...BatchGetOption) error

	// BatchSave saves multiple resources at once
	// TODO: Implement
	//BatchSave(ctx context.Context, toSave []resource.Resource) error

	// BatchDelete removes multiple resources at once
	// TODO: Implement
	//BatchDelete(ctx context.Context, toDelete []resource.Reference) error

	// GetResourceDescriptors returns all supported resource descriptors under this Access
	GetResourceDescriptors() []Descriptor
}

// NewCompositeAccess constructs Access over multiple smaller ones
func NewCompositeAccess(accesses ...Access) Access {
	composite := &compositeAccess{
		items: map[Descriptor]Access{},
	}
	for _, access := range accesses {
		for _, descriptor := range access.GetResourceDescriptors() {
			composite.items[descriptor] = access
		}
	}
	for descriptor := range composite.items {
		composite.linear = append(composite.linear, descriptor)
	}
	return composite
}

type compositeAccess struct {
	items  map[Descriptor]Access
	linear []Descriptor
}

func (ca *compositeAccess) Get(ctx context.Context, q GetQuery) (Resource, error) {
	return ca.items[q.GetResourceDescriptor()].Get(ctx, q)
}

func (ca *compositeAccess) Query(ctx context.Context, q ListQuery) (QueryResultSnapshot, error) {
	return ca.items[q.GetResourceDescriptor()].Query(ctx, q)
}

func (ca *compositeAccess) Search(ctx context.Context, q SearchQuery) (SearchQueryResultSnapshot, error) {
	return ca.items[q.GetResourceDescriptor()].Search(ctx, q)
}

func (ca *compositeAccess) Watch(ctx context.Context, q GetQuery, cb func(ch ResourceChange) error) error {
	return ca.items[q.GetResourceDescriptor()].Watch(ctx, q, cb)
}

func (ca *compositeAccess) WatchQuery(ctx context.Context, q WatchQuery, cb func(ch QueryResultChange) error) error {
	return ca.items[q.GetResourceDescriptor()].WatchQuery(ctx, q, cb)
}

func (ca *compositeAccess) Save(ctx context.Context, res Resource, opts ...SaveOption) error {
	return ca.items[res.GetResourceDescriptor()].Save(ctx, res, opts...)
}

func (ca *compositeAccess) Delete(ctx context.Context, ref Reference, opts ...DeleteOption) error {
	return ca.items[ref.GetResourceDescriptor()].Delete(ctx, ref, opts...)
}

func (ca *compositeAccess) BatchGet(ctx context.Context, toGet []Reference, opts ...BatchGetOption) error {
	byDescriptor := map[Descriptor][]Reference{}
	for _, ref := range toGet {
		byDescriptor[ref.GetResourceDescriptor()] = append(byDescriptor[ref.GetResourceDescriptor()], ref)
	}
	for descriptor, refs := range byDescriptor {
		if err := ca.items[descriptor].BatchGet(ctx, refs, opts...); err != nil {
			return err
		}
	}
	return nil
}

func (ca *compositeAccess) GetResourceDescriptors() []Descriptor {
	return ca.linear
}
