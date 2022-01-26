package resource

// Interface for goten reference or parent reference types
type Reference interface {
	Name
	Resolved() bool
	ClearCached()
	GetRawResource() Resource
	ResolveRaw(Resource) error
}
