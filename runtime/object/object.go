package object

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

// GotenObjectExt is implemented by a Go representation of
// a protobuf message for which protoc-goten-gen-object
// plugin was used. GotenObjectExt is basically an extension
// of a regular go-protobuf object containing many additional
// goten-specific features like type-safe embedded field paths,
// field masks, diff support, customized merge support etc.
// Those extensions are mandatory for request or resource objects,
// but can be turned off for most other types.
type GotenObjectExt interface {
	fmt.Stringer
	proto.Message
	GotenObjectExt()

	// CloneRaw clones current object.
	CloneRaw() GotenObjectExt

	// MergeRaw merges provided source into current object.
	MergeRaw(GotenObjectExt)

	// MakeRawFullFieldMask returns full mask. Full field mask contains
	// paths of all fields without sub paths items,
	// as it is considered that field path already contains
	// sub paths.
	MakeRawFullFieldMask() FieldMask

	// MakeRawDiffFieldMask returns diff mask compared with given resource.
	// Panics, if other resource is not of same type.
	MakeRawDiffFieldMask(other GotenObjectExt) FieldMask
}
