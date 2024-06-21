// Code generated by protoc-gen-goten-object
// File: goten/meta-service/proto/v1/resource.proto
// DO NOT EDIT!!!

package resource

import (
	"encoding/json"
	"strings"

	firestorepb "google.golang.org/genproto/googleapis/firestore/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = new(json.Marshaler)
	_ = strings.Builder{}

	_ = firestorepb.Value{}
	_ = codes.NotFound
	_ = status.Status{}
	_ = new(proto.Message)
	_ = new(preflect.Message)
	_ = googlefieldmaskpb.FieldMask{}

	_ = new(gotenobject.FieldMask)
)

// make sure we're using proto imports
var (
	_ = &service.Service{}
	_ = &meta.Meta{}
)

type Resource_FieldMask struct {
	Paths []Resource_FieldPath
}

func FullResource_FieldMask() *Resource_FieldMask {
	res := &Resource_FieldMask{}
	res.Paths = append(res.Paths, &Resource_FieldTerminalPath{selector: Resource_FieldPathSelectorName})
	res.Paths = append(res.Paths, &Resource_FieldTerminalPath{selector: Resource_FieldPathSelectorMetadata})
	res.Paths = append(res.Paths, &Resource_FieldTerminalPath{selector: Resource_FieldPathSelectorPluralName})
	res.Paths = append(res.Paths, &Resource_FieldTerminalPath{selector: Resource_FieldPathSelectorFqn})
	res.Paths = append(res.Paths, &Resource_FieldTerminalPath{selector: Resource_FieldPathSelectorVersions})
	res.Paths = append(res.Paths, &Resource_FieldTerminalPath{selector: Resource_FieldPathSelectorVersionedInfos})
	return res
}

func (fieldMask *Resource_FieldMask) String() string {
	if fieldMask == nil {
		return "<nil>"
	}
	pathsStr := make([]string, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		pathsStr = append(pathsStr, path.String())
	}
	return strings.Join(pathsStr, ", ")
}

// firestore encoding/decoding integration
func (fieldMask *Resource_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
	if fieldMask == nil {
		return &firestorepb.Value{ValueType: &firestorepb.Value_NullValue{}}, nil
	}
	arrayValues := make([]*firestorepb.Value, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.GetPaths() {
		arrayValues = append(arrayValues, &firestorepb.Value{ValueType: &firestorepb.Value_StringValue{StringValue: path.String()}})
	}
	return &firestorepb.Value{
		ValueType: &firestorepb.Value_ArrayValue{ArrayValue: &firestorepb.ArrayValue{Values: arrayValues}},
	}, nil
}

func (fieldMask *Resource_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseResource_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *Resource_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 6)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*Resource_FieldTerminalPath); ok {
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

func (fieldMask *Resource_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseResource_FieldPath(raw)
	})
}

func (fieldMask *Resource_FieldMask) ProtoMessage() {}

func (fieldMask *Resource_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *Resource_FieldMask) Subtract(other *Resource_FieldMask) *Resource_FieldMask {
	result := &Resource_FieldMask{}
	removedSelectors := make([]bool, 6)
	otherSubMasks := map[Resource_FieldPathSelector]gotenobject.FieldMask{
		Resource_FieldPathSelectorMetadata:       &meta.Meta_FieldMask{},
		Resource_FieldPathSelectorVersionedInfos: &Resource_VersionedInfo_FieldMask{},
	}
	mySubMasks := map[Resource_FieldPathSelector]gotenobject.FieldMask{
		Resource_FieldPathSelectorMetadata:       &meta.Meta_FieldMask{},
		Resource_FieldPathSelectorVersionedInfos: &Resource_VersionedInfo_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *Resource_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *Resource_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*Resource_FieldTerminalPath); ok {
					switch tp.selector {
					case Resource_FieldPathSelectorMetadata:
						mySubMasks[Resource_FieldPathSelectorMetadata] = meta.FullMeta_FieldMask()
					case Resource_FieldPathSelectorVersionedInfos:
						mySubMasks[Resource_FieldPathSelectorVersionedInfos] = FullResource_VersionedInfo_FieldMask()
					}
				} else if tp, ok := path.(*Resource_FieldSubPath); ok {
					mySubMasks[tp.selector].AppendRawPath(tp.subPath)
				}
			} else {
				result.Paths = append(result.Paths, path)
			}
		}
	}
	for selector, mySubMask := range mySubMasks {
		if mySubMask.PathsCount() > 0 {
			for _, allowedPath := range mySubMask.SubtractRaw(otherSubMasks[selector]).GetRawPaths() {
				result.Paths = append(result.Paths, &Resource_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *Resource_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*Resource_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *Resource_FieldMask) FilterInputFields() *Resource_FieldMask {
	result := &Resource_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case Resource_FieldPathSelectorFqn:
		case Resource_FieldPathSelectorMetadata:
			if _, ok := path.(*Resource_FieldTerminalPath); ok {
				for _, subpath := range meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &Resource_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*Resource_FieldSubPath); ok {
				selectedMask := &meta.Meta_FieldMask{
					Paths: []meta.Meta_FieldPath{sub.subPath.(meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &Resource_FieldSubPath{selector: Resource_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *Resource_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *Resource_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]Resource_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseResource_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask Resource_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *Resource_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Resource_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask Resource_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *Resource_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Resource_FieldMask) AppendPath(path Resource_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *Resource_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(Resource_FieldPath))
}

func (fieldMask *Resource_FieldMask) GetPaths() []Resource_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *Resource_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *Resource_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseResource_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *Resource_FieldMask) Set(target, source *Resource) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *Resource_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*Resource), source.(*Resource))
}

func (fieldMask *Resource_FieldMask) Project(source *Resource) *Resource {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &Resource{}
	metadataMask := &meta.Meta_FieldMask{}
	wholeMetadataAccepted := false
	versionedInfosMask := &Resource_VersionedInfo_FieldMask{}
	wholeVersionedInfosAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *Resource_FieldTerminalPath:
			switch tp.selector {
			case Resource_FieldPathSelectorName:
				result.Name = source.Name
			case Resource_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			case Resource_FieldPathSelectorPluralName:
				result.PluralName = source.PluralName
			case Resource_FieldPathSelectorFqn:
				result.Fqn = source.Fqn
			case Resource_FieldPathSelectorVersions:
				result.Versions = source.Versions
			case Resource_FieldPathSelectorVersionedInfos:
				result.VersionedInfos = source.VersionedInfos
				wholeVersionedInfosAccepted = true
			}
		case *Resource_FieldSubPath:
			switch tp.selector {
			case Resource_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(meta.Meta_FieldPath))
			case Resource_FieldPathSelectorVersionedInfos:
				versionedInfosMask.AppendPath(tp.subPath.(ResourceVersionedInfo_FieldPath))
			}
		}
	}
	if wholeMetadataAccepted == false && len(metadataMask.Paths) > 0 {
		result.Metadata = metadataMask.Project(source.GetMetadata())
	}
	if wholeVersionedInfosAccepted == false && len(versionedInfosMask.Paths) > 0 {
		for _, sourceItem := range source.GetVersionedInfos() {
			result.VersionedInfos = append(result.VersionedInfos, versionedInfosMask.Project(sourceItem))
		}
	}
	return result
}

func (fieldMask *Resource_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*Resource))
}

func (fieldMask *Resource_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type Resource_VersionedInfo_FieldMask struct {
	Paths []ResourceVersionedInfo_FieldPath
}

func FullResource_VersionedInfo_FieldMask() *Resource_VersionedInfo_FieldMask {
	res := &Resource_VersionedInfo_FieldMask{}
	res.Paths = append(res.Paths, &ResourceVersionedInfo_FieldTerminalPath{selector: ResourceVersionedInfo_FieldPathSelectorVersion})
	res.Paths = append(res.Paths, &ResourceVersionedInfo_FieldTerminalPath{selector: ResourceVersionedInfo_FieldPathSelectorIsRegional})
	return res
}

func (fieldMask *Resource_VersionedInfo_FieldMask) String() string {
	if fieldMask == nil {
		return "<nil>"
	}
	pathsStr := make([]string, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		pathsStr = append(pathsStr, path.String())
	}
	return strings.Join(pathsStr, ", ")
}

// firestore encoding/decoding integration
func (fieldMask *Resource_VersionedInfo_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
	if fieldMask == nil {
		return &firestorepb.Value{ValueType: &firestorepb.Value_NullValue{}}, nil
	}
	arrayValues := make([]*firestorepb.Value, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.GetPaths() {
		arrayValues = append(arrayValues, &firestorepb.Value{ValueType: &firestorepb.Value_StringValue{StringValue: path.String()}})
	}
	return &firestorepb.Value{
		ValueType: &firestorepb.Value_ArrayValue{ArrayValue: &firestorepb.ArrayValue{Values: arrayValues}},
	}, nil
}

func (fieldMask *Resource_VersionedInfo_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseResourceVersionedInfo_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *Resource_VersionedInfo_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 2)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*ResourceVersionedInfo_FieldTerminalPath); ok {
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

func (fieldMask *Resource_VersionedInfo_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseResourceVersionedInfo_FieldPath(raw)
	})
}

func (fieldMask *Resource_VersionedInfo_FieldMask) ProtoMessage() {}

func (fieldMask *Resource_VersionedInfo_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *Resource_VersionedInfo_FieldMask) Subtract(other *Resource_VersionedInfo_FieldMask) *Resource_VersionedInfo_FieldMask {
	result := &Resource_VersionedInfo_FieldMask{}
	removedSelectors := make([]bool, 2)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *ResourceVersionedInfo_FieldTerminalPath:
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

func (fieldMask *Resource_VersionedInfo_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*Resource_VersionedInfo_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *Resource_VersionedInfo_FieldMask) FilterInputFields() *Resource_VersionedInfo_FieldMask {
	result := &Resource_VersionedInfo_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *Resource_VersionedInfo_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *Resource_VersionedInfo_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]ResourceVersionedInfo_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseResourceVersionedInfo_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask Resource_VersionedInfo_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *Resource_VersionedInfo_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Resource_VersionedInfo_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask Resource_VersionedInfo_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *Resource_VersionedInfo_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Resource_VersionedInfo_FieldMask) AppendPath(path ResourceVersionedInfo_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *Resource_VersionedInfo_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(ResourceVersionedInfo_FieldPath))
}

func (fieldMask *Resource_VersionedInfo_FieldMask) GetPaths() []ResourceVersionedInfo_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *Resource_VersionedInfo_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *Resource_VersionedInfo_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseResourceVersionedInfo_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *Resource_VersionedInfo_FieldMask) Set(target, source *Resource_VersionedInfo) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *Resource_VersionedInfo_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*Resource_VersionedInfo), source.(*Resource_VersionedInfo))
}

func (fieldMask *Resource_VersionedInfo_FieldMask) Project(source *Resource_VersionedInfo) *Resource_VersionedInfo {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &Resource_VersionedInfo{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *ResourceVersionedInfo_FieldTerminalPath:
			switch tp.selector {
			case ResourceVersionedInfo_FieldPathSelectorVersion:
				result.Version = source.Version
			case ResourceVersionedInfo_FieldPathSelectorIsRegional:
				result.IsRegional = source.IsRegional
			}
		}
	}
	return result
}

func (fieldMask *Resource_VersionedInfo_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*Resource_VersionedInfo))
}

func (fieldMask *Resource_VersionedInfo_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
