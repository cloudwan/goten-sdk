// Code generated by protoc-gen-goten-object
// File: goten/meta-service/proto/v1/common.proto
// DO NOT EDIT!!!

package common

import (
	"encoding/json"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import ()

// ensure the imports are used
var (
	_ = new(json.Marshaler)
	_ = strings.Builder{}

	_ = codes.NotFound
	_ = status.Status{}
	_ = new(proto.Message)
	_ = new(preflect.Message)
	_ = googlefieldmaskpb.FieldMask{}

	_ = new(gotenobject.FieldMask)
)

// make sure we're using proto imports
var ()

type LabelledDomain_FieldMask struct {
	Paths []LabelledDomain_FieldPath
}

func FullLabelledDomain_FieldMask() *LabelledDomain_FieldMask {
	res := &LabelledDomain_FieldMask{}
	res.Paths = append(res.Paths, &LabelledDomain_FieldTerminalPath{selector: LabelledDomain_FieldPathSelectorLabel})
	res.Paths = append(res.Paths, &LabelledDomain_FieldTerminalPath{selector: LabelledDomain_FieldPathSelectorDomain})
	res.Paths = append(res.Paths, &LabelledDomain_FieldTerminalPath{selector: LabelledDomain_FieldPathSelectorAvailableMixins})
	res.Paths = append(res.Paths, &LabelledDomain_FieldTerminalPath{selector: LabelledDomain_FieldPathSelectorWebGrpcAvailable})
	res.Paths = append(res.Paths, &LabelledDomain_FieldTerminalPath{selector: LabelledDomain_FieldPathSelectorRestApiAvailable})
	res.Paths = append(res.Paths, &LabelledDomain_FieldTerminalPath{selector: LabelledDomain_FieldPathSelectorIsPrivate})
	return res
}

func (fieldMask *LabelledDomain_FieldMask) String() string {
	if fieldMask == nil {
		return "<nil>"
	}
	pathsStr := make([]string, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		pathsStr = append(pathsStr, path.String())
	}
	return strings.Join(pathsStr, ", ")
}

func (fieldMask *LabelledDomain_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 6)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*LabelledDomain_FieldTerminalPath); ok {
			presentSelectors[int(asFinal.selector)] = true
		}
	}
	for _, flag := range presentSelectors {
		if !flag {
			return false
		}
	}
	return true
}

func (fieldMask *LabelledDomain_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseLabelledDomain_FieldPath(raw)
	})
}

func (fieldMask *LabelledDomain_FieldMask) ProtoMessage() {}

func (fieldMask *LabelledDomain_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *LabelledDomain_FieldMask) Subtract(other *LabelledDomain_FieldMask) *LabelledDomain_FieldMask {
	result := &LabelledDomain_FieldMask{}
	removedSelectors := make([]bool, 6)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *LabelledDomain_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			result.Paths = append(result.Paths, path)
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *LabelledDomain_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*LabelledDomain_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *LabelledDomain_FieldMask) FilterInputFields() *LabelledDomain_FieldMask {
	result := &LabelledDomain_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *LabelledDomain_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *LabelledDomain_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]LabelledDomain_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseLabelledDomain_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask LabelledDomain_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *LabelledDomain_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *LabelledDomain_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask LabelledDomain_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *LabelledDomain_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *LabelledDomain_FieldMask) AppendPath(path LabelledDomain_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *LabelledDomain_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(LabelledDomain_FieldPath))
}

func (fieldMask *LabelledDomain_FieldMask) GetPaths() []LabelledDomain_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *LabelledDomain_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *LabelledDomain_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseLabelledDomain_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *LabelledDomain_FieldMask) Set(target, source *LabelledDomain) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *LabelledDomain_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*LabelledDomain), source.(*LabelledDomain))
}

func (fieldMask *LabelledDomain_FieldMask) Project(source *LabelledDomain) *LabelledDomain {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &LabelledDomain{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *LabelledDomain_FieldTerminalPath:
			switch tp.selector {
			case LabelledDomain_FieldPathSelectorLabel:
				result.Label = source.Label
			case LabelledDomain_FieldPathSelectorDomain:
				result.Domain = source.Domain
			case LabelledDomain_FieldPathSelectorAvailableMixins:
				result.AvailableMixins = source.AvailableMixins
			case LabelledDomain_FieldPathSelectorWebGrpcAvailable:
				result.WebGrpcAvailable = source.WebGrpcAvailable
			case LabelledDomain_FieldPathSelectorRestApiAvailable:
				result.RestApiAvailable = source.RestApiAvailable
			case LabelledDomain_FieldPathSelectorIsPrivate:
				result.IsPrivate = source.IsPrivate
			}
		}
	}
	return result
}

func (fieldMask *LabelledDomain_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*LabelledDomain))
}

func (fieldMask *LabelledDomain_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
