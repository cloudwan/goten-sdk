// Code generated by protoc-gen-goten-object
// File: goten/meta-service/proto/v1/service.proto
// DO NOT EDIT!!!

package service

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
	meta "github.com/cloudwan/goten-sdk/types/meta"
	multi_region_policy "github.com/cloudwan/goten-sdk/types/multi_region_policy"
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
	_ = &meta.Meta{}
	_ = &multi_region_policy.MultiRegionPolicy{}
)

type Service_FieldMask struct {
	Paths []Service_FieldPath
}

func FullService_FieldMask() *Service_FieldMask {
	res := &Service_FieldMask{}
	res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorName})
	res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorMetadata})
	res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorMultiRegionPolicy})
	res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorDisplayName})
	res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorAllVersions})
	res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorGlobalDomain})
	res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorImportedServices})
	res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorUsedServices})
	res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorImportedVersions})
	res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorEnvRegistryGeneration})
	res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorAutomaticVersionSwitch})
	res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorServicesCtrl})
	return res
}

func (fieldMask *Service_FieldMask) String() string {
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
func (fieldMask *Service_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *Service_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseService_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *Service_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 12)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*Service_FieldTerminalPath); ok {
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

func (fieldMask *Service_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseService_FieldPath(raw)
	})
}

func (fieldMask *Service_FieldMask) ProtoMessage() {}

func (fieldMask *Service_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *Service_FieldMask) Subtract(other *Service_FieldMask) *Service_FieldMask {
	result := &Service_FieldMask{}
	removedSelectors := make([]bool, 12)
	otherSubMasks := map[Service_FieldPathSelector]gotenobject.FieldMask{
		Service_FieldPathSelectorMetadata:          &meta.Meta_FieldMask{},
		Service_FieldPathSelectorMultiRegionPolicy: &multi_region_policy.MultiRegionPolicy_FieldMask{},
		Service_FieldPathSelectorImportedVersions:  &Service_ImportedVersions_FieldMask{},
		Service_FieldPathSelectorServicesCtrl:      &Service_AllowedServicesCtrlFlag_FieldMask{},
	}
	mySubMasks := map[Service_FieldPathSelector]gotenobject.FieldMask{
		Service_FieldPathSelectorMetadata:          &meta.Meta_FieldMask{},
		Service_FieldPathSelectorMultiRegionPolicy: &multi_region_policy.MultiRegionPolicy_FieldMask{},
		Service_FieldPathSelectorImportedVersions:  &Service_ImportedVersions_FieldMask{},
		Service_FieldPathSelectorServicesCtrl:      &Service_AllowedServicesCtrlFlag_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *Service_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *Service_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*Service_FieldTerminalPath); ok {
					switch tp.selector {
					case Service_FieldPathSelectorMetadata:
						mySubMasks[Service_FieldPathSelectorMetadata] = meta.FullMeta_FieldMask()
					case Service_FieldPathSelectorMultiRegionPolicy:
						mySubMasks[Service_FieldPathSelectorMultiRegionPolicy] = multi_region_policy.FullMultiRegionPolicy_FieldMask()
					case Service_FieldPathSelectorImportedVersions:
						mySubMasks[Service_FieldPathSelectorImportedVersions] = FullService_ImportedVersions_FieldMask()
					case Service_FieldPathSelectorServicesCtrl:
						mySubMasks[Service_FieldPathSelectorServicesCtrl] = FullService_AllowedServicesCtrlFlag_FieldMask()
					}
				} else if tp, ok := path.(*Service_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &Service_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *Service_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*Service_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *Service_FieldMask) FilterInputFields() *Service_FieldMask {
	result := &Service_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case Service_FieldPathSelectorServicesCtrl:
		case Service_FieldPathSelectorMetadata:
			if _, ok := path.(*Service_FieldTerminalPath); ok {
				for _, subpath := range meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &Service_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*Service_FieldSubPath); ok {
				selectedMask := &meta.Meta_FieldMask{
					Paths: []meta.Meta_FieldPath{sub.subPath.(meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &Service_FieldSubPath{selector: Service_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *Service_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *Service_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]Service_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseService_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask Service_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *Service_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Service_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask Service_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *Service_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Service_FieldMask) AppendPath(path Service_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *Service_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(Service_FieldPath))
}

func (fieldMask *Service_FieldMask) GetPaths() []Service_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *Service_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *Service_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseService_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *Service_FieldMask) Set(target, source *Service) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *Service_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*Service), source.(*Service))
}

func (fieldMask *Service_FieldMask) Project(source *Service) *Service {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &Service{}
	metadataMask := &meta.Meta_FieldMask{}
	wholeMetadataAccepted := false
	multiRegionPolicyMask := &multi_region_policy.MultiRegionPolicy_FieldMask{}
	wholeMultiRegionPolicyAccepted := false
	importedVersionsMask := &Service_ImportedVersions_FieldMask{}
	wholeImportedVersionsAccepted := false
	servicesCtrlMask := &Service_AllowedServicesCtrlFlag_FieldMask{}
	wholeServicesCtrlAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *Service_FieldTerminalPath:
			switch tp.selector {
			case Service_FieldPathSelectorName:
				result.Name = source.Name
			case Service_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			case Service_FieldPathSelectorMultiRegionPolicy:
				result.MultiRegionPolicy = source.MultiRegionPolicy
				wholeMultiRegionPolicyAccepted = true
			case Service_FieldPathSelectorDisplayName:
				result.DisplayName = source.DisplayName
			case Service_FieldPathSelectorAllVersions:
				result.AllVersions = source.AllVersions
			case Service_FieldPathSelectorGlobalDomain:
				result.GlobalDomain = source.GlobalDomain
			case Service_FieldPathSelectorImportedServices:
				result.ImportedServices = source.ImportedServices
			case Service_FieldPathSelectorUsedServices:
				result.UsedServices = source.UsedServices
			case Service_FieldPathSelectorImportedVersions:
				result.ImportedVersions = source.ImportedVersions
				wholeImportedVersionsAccepted = true
			case Service_FieldPathSelectorEnvRegistryGeneration:
				result.EnvRegistryGeneration = source.EnvRegistryGeneration
			case Service_FieldPathSelectorAutomaticVersionSwitch:
				result.AutomaticVersionSwitch = source.AutomaticVersionSwitch
			case Service_FieldPathSelectorServicesCtrl:
				result.ServicesCtrl = source.ServicesCtrl
				wholeServicesCtrlAccepted = true
			}
		case *Service_FieldSubPath:
			switch tp.selector {
			case Service_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(meta.Meta_FieldPath))
			case Service_FieldPathSelectorMultiRegionPolicy:
				multiRegionPolicyMask.AppendPath(tp.subPath.(multi_region_policy.MultiRegionPolicy_FieldPath))
			case Service_FieldPathSelectorImportedVersions:
				importedVersionsMask.AppendPath(tp.subPath.(ServiceImportedVersions_FieldPath))
			case Service_FieldPathSelectorServicesCtrl:
				servicesCtrlMask.AppendPath(tp.subPath.(ServiceAllowedServicesCtrlFlag_FieldPath))
			}
		}
	}
	if wholeMetadataAccepted == false && len(metadataMask.Paths) > 0 {
		result.Metadata = metadataMask.Project(source.GetMetadata())
	}
	if wholeMultiRegionPolicyAccepted == false && len(multiRegionPolicyMask.Paths) > 0 {
		result.MultiRegionPolicy = multiRegionPolicyMask.Project(source.GetMultiRegionPolicy())
	}
	if wholeImportedVersionsAccepted == false && len(importedVersionsMask.Paths) > 0 {
		for _, sourceItem := range source.GetImportedVersions() {
			result.ImportedVersions = append(result.ImportedVersions, importedVersionsMask.Project(sourceItem))
		}
	}
	if wholeServicesCtrlAccepted == false && len(servicesCtrlMask.Paths) > 0 {
		result.ServicesCtrl = servicesCtrlMask.Project(source.GetServicesCtrl())
	}
	return result
}

func (fieldMask *Service_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*Service))
}

func (fieldMask *Service_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type Service_ImportedVersions_FieldMask struct {
	Paths []ServiceImportedVersions_FieldPath
}

func FullService_ImportedVersions_FieldMask() *Service_ImportedVersions_FieldMask {
	res := &Service_ImportedVersions_FieldMask{}
	res.Paths = append(res.Paths, &ServiceImportedVersions_FieldTerminalPath{selector: ServiceImportedVersions_FieldPathSelectorTargetService})
	res.Paths = append(res.Paths, &ServiceImportedVersions_FieldTerminalPath{selector: ServiceImportedVersions_FieldPathSelectorTargetServiceVersion})
	res.Paths = append(res.Paths, &ServiceImportedVersions_FieldTerminalPath{selector: ServiceImportedVersions_FieldPathSelectorCurrentServiceVersion})
	return res
}

func (fieldMask *Service_ImportedVersions_FieldMask) String() string {
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
func (fieldMask *Service_ImportedVersions_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *Service_ImportedVersions_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseServiceImportedVersions_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *Service_ImportedVersions_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 3)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*ServiceImportedVersions_FieldTerminalPath); ok {
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

func (fieldMask *Service_ImportedVersions_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseServiceImportedVersions_FieldPath(raw)
	})
}

func (fieldMask *Service_ImportedVersions_FieldMask) ProtoMessage() {}

func (fieldMask *Service_ImportedVersions_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *Service_ImportedVersions_FieldMask) Subtract(other *Service_ImportedVersions_FieldMask) *Service_ImportedVersions_FieldMask {
	result := &Service_ImportedVersions_FieldMask{}
	removedSelectors := make([]bool, 3)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *ServiceImportedVersions_FieldTerminalPath:
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

func (fieldMask *Service_ImportedVersions_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*Service_ImportedVersions_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *Service_ImportedVersions_FieldMask) FilterInputFields() *Service_ImportedVersions_FieldMask {
	result := &Service_ImportedVersions_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *Service_ImportedVersions_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *Service_ImportedVersions_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]ServiceImportedVersions_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseServiceImportedVersions_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask Service_ImportedVersions_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *Service_ImportedVersions_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Service_ImportedVersions_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask Service_ImportedVersions_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *Service_ImportedVersions_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Service_ImportedVersions_FieldMask) AppendPath(path ServiceImportedVersions_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *Service_ImportedVersions_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(ServiceImportedVersions_FieldPath))
}

func (fieldMask *Service_ImportedVersions_FieldMask) GetPaths() []ServiceImportedVersions_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *Service_ImportedVersions_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *Service_ImportedVersions_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseServiceImportedVersions_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *Service_ImportedVersions_FieldMask) Set(target, source *Service_ImportedVersions) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *Service_ImportedVersions_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*Service_ImportedVersions), source.(*Service_ImportedVersions))
}

func (fieldMask *Service_ImportedVersions_FieldMask) Project(source *Service_ImportedVersions) *Service_ImportedVersions {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &Service_ImportedVersions{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *ServiceImportedVersions_FieldTerminalPath:
			switch tp.selector {
			case ServiceImportedVersions_FieldPathSelectorTargetService:
				result.TargetService = source.TargetService
			case ServiceImportedVersions_FieldPathSelectorTargetServiceVersion:
				result.TargetServiceVersion = source.TargetServiceVersion
			case ServiceImportedVersions_FieldPathSelectorCurrentServiceVersion:
				result.CurrentServiceVersion = source.CurrentServiceVersion
			}
		}
	}
	return result
}

func (fieldMask *Service_ImportedVersions_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*Service_ImportedVersions))
}

func (fieldMask *Service_ImportedVersions_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type Service_AllowedServicesCtrlFlag_FieldMask struct {
	Paths []ServiceAllowedServicesCtrlFlag_FieldPath
}

func FullService_AllowedServicesCtrlFlag_FieldMask() *Service_AllowedServicesCtrlFlag_FieldMask {
	res := &Service_AllowedServicesCtrlFlag_FieldMask{}
	res.Paths = append(res.Paths, &ServiceAllowedServicesCtrlFlag_FieldTerminalPath{selector: ServiceAllowedServicesCtrlFlag_FieldPathSelectorIsDirty})
	res.Paths = append(res.Paths, &ServiceAllowedServicesCtrlFlag_FieldTerminalPath{selector: ServiceAllowedServicesCtrlFlag_FieldPathSelectorGeneration})
	return res
}

func (fieldMask *Service_AllowedServicesCtrlFlag_FieldMask) String() string {
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
func (fieldMask *Service_AllowedServicesCtrlFlag_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *Service_AllowedServicesCtrlFlag_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseServiceAllowedServicesCtrlFlag_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *Service_AllowedServicesCtrlFlag_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 2)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*ServiceAllowedServicesCtrlFlag_FieldTerminalPath); ok {
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

func (fieldMask *Service_AllowedServicesCtrlFlag_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseServiceAllowedServicesCtrlFlag_FieldPath(raw)
	})
}

func (fieldMask *Service_AllowedServicesCtrlFlag_FieldMask) ProtoMessage() {}

func (fieldMask *Service_AllowedServicesCtrlFlag_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *Service_AllowedServicesCtrlFlag_FieldMask) Subtract(other *Service_AllowedServicesCtrlFlag_FieldMask) *Service_AllowedServicesCtrlFlag_FieldMask {
	result := &Service_AllowedServicesCtrlFlag_FieldMask{}
	removedSelectors := make([]bool, 2)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *ServiceAllowedServicesCtrlFlag_FieldTerminalPath:
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

func (fieldMask *Service_AllowedServicesCtrlFlag_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*Service_AllowedServicesCtrlFlag_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *Service_AllowedServicesCtrlFlag_FieldMask) FilterInputFields() *Service_AllowedServicesCtrlFlag_FieldMask {
	result := &Service_AllowedServicesCtrlFlag_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *Service_AllowedServicesCtrlFlag_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *Service_AllowedServicesCtrlFlag_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]ServiceAllowedServicesCtrlFlag_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseServiceAllowedServicesCtrlFlag_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask Service_AllowedServicesCtrlFlag_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *Service_AllowedServicesCtrlFlag_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Service_AllowedServicesCtrlFlag_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask Service_AllowedServicesCtrlFlag_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *Service_AllowedServicesCtrlFlag_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Service_AllowedServicesCtrlFlag_FieldMask) AppendPath(path ServiceAllowedServicesCtrlFlag_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *Service_AllowedServicesCtrlFlag_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(ServiceAllowedServicesCtrlFlag_FieldPath))
}

func (fieldMask *Service_AllowedServicesCtrlFlag_FieldMask) GetPaths() []ServiceAllowedServicesCtrlFlag_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *Service_AllowedServicesCtrlFlag_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *Service_AllowedServicesCtrlFlag_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseServiceAllowedServicesCtrlFlag_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *Service_AllowedServicesCtrlFlag_FieldMask) Set(target, source *Service_AllowedServicesCtrlFlag) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *Service_AllowedServicesCtrlFlag_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*Service_AllowedServicesCtrlFlag), source.(*Service_AllowedServicesCtrlFlag))
}

func (fieldMask *Service_AllowedServicesCtrlFlag_FieldMask) Project(source *Service_AllowedServicesCtrlFlag) *Service_AllowedServicesCtrlFlag {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &Service_AllowedServicesCtrlFlag{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *ServiceAllowedServicesCtrlFlag_FieldTerminalPath:
			switch tp.selector {
			case ServiceAllowedServicesCtrlFlag_FieldPathSelectorIsDirty:
				result.IsDirty = source.IsDirty
			case ServiceAllowedServicesCtrlFlag_FieldPathSelectorGeneration:
				result.Generation = source.Generation
			}
		}
	}
	return result
}

func (fieldMask *Service_AllowedServicesCtrlFlag_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*Service_AllowedServicesCtrlFlag))
}

func (fieldMask *Service_AllowedServicesCtrlFlag_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
