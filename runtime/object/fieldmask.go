package object

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// FieldMask describes an interface Interface of goten field mask.
type FieldMask interface {
	fmt.Stringer
	proto.Message

	// ToProtoFieldMask converts to google.protobuf.FieldMask
	ToProtoFieldMask() *fieldmaskpb.FieldMask

	// FromProtoFieldMask erases current contents and re-fills using
	// paths provided by google.protobuf.FieldMask
	FromProtoFieldMask(*fieldmaskpb.FieldMask) error

	// AppendRawPath appends new field path
	AppendRawPath(path FieldPath)

	// SubtractRaw creates new FieldMask by subtracting current
	// instance from the one given as argument.
	// Subtracting operation is using following rules:
	// * F1 {A, C} - F2 {A, B} = F3 {C} (A removes A)
	// * F1 {A} - F2 {A.B} = F3 {A.X1, ... A.Xn} (where X
	//   is any possible subpath of A that is not B).
	//   Exception to this rule are field paths containing
	//   map keys, where possible Xn combinations are
	//   open-ended. Sub-paths ending on map leaves are
	//   ignored in F2 (left side argument).
	// * F1 {A.B} - F2 {A} = F3 {} (if we remove A, we remove
	//   all sub-paths of A too).
	SubtractRaw(other FieldMask) FieldMask

	// GetRawPaths returns all current field paths
	GetRawPaths() []FieldPath

	// IsFull indicates if mask is full. Mask is considered as
	// full, if proto.Equal(a, mask.CloneRaw(a)) equals
	// true.
	IsFull() bool

	// SetRaw copies shallowly values pointed by internally hold
	// fields paths from source to the target.
	SetRaw(target, source GotenObjectExt)

	// ProjectRaw makes projection (shallow copy) of the given
	// resource for internal field paths only.
	// If the mask is nil, then same object is returned.
	ProjectRaw(source GotenObjectExt) GotenObjectExt

	// PathsCount returns number of field paths included
	PathsCount() int
}
