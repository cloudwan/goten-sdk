package resource

import (
	"github.com/cloudwan/goten-sdk/runtime/goten"
	"github.com/cloudwan/goten-sdk/runtime/object"
)

type Resource interface {
	goten.GotenMessage
	object.GotenObjectExt
	GetRawName() Name
	GetResourceDescriptor() Descriptor
}

type ResourceChange interface {
	goten.GotenMessage
	IsAdd() bool
	IsModify() bool
	IsDelete() bool
	IsCurrent() bool
	GetResource() Resource
	GetRawName() Name
	SetAddedRaw(Resource, int)
	SetModifiedRaw(Name, Resource, int, int)
	SetDeletedRaw(Name, int)
	SetCurrentRaw(Resource)
	GetCurrentViewIndex() int32
	GetPreviousViewIndex() int32
}

type DisplayableResource interface {
	Resource
	GetDisplayName() string
}
