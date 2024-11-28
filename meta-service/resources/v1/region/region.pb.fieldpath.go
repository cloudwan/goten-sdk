// Code generated by protoc-gen-goten-object
// File: goten/meta-service/proto/v1/region.proto
// DO NOT EDIT!!!

package region

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoregistry"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	"github.com/cloudwan/goten-sdk/runtime/strcase"
)

// proto imports
import (
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = new(json.Marshaler)
	_ = new(fmt.Stringer)
	_ = reflect.DeepEqual
	_ = strings.Builder{}
	_ = time.Second

	_ = strcase.ToLowerCamel
	_ = codes.NotFound
	_ = status.Status{}
	_ = protojson.UnmarshalOptions{}
	_ = new(proto.Message)
	_ = protoregistry.GlobalTypes

	_ = new(gotenobject.FieldPath)
)

// make sure we're using proto imports
var (
	_ = &meta.Meta{}
)

// FieldPath provides implementation to handle
// https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto
type Region_FieldPath interface {
	gotenobject.FieldPath
	Selector() Region_FieldPathSelector
	Get(source *Region) []interface{}
	GetSingle(source *Region) (interface{}, bool)
	ClearValue(item *Region)

	// Those methods build corresponding Region_FieldPathValue
	// (or array of values) and holds passed value. Panics if injected type is incorrect.
	WithIValue(value interface{}) Region_FieldPathValue
	WithIArrayOfValues(values interface{}) Region_FieldPathArrayOfValues
	WithIArrayItemValue(value interface{}) Region_FieldPathArrayItemValue
}

type Region_FieldPathSelector int32

const (
	Region_FieldPathSelectorName     Region_FieldPathSelector = 0
	Region_FieldPathSelectorMetadata Region_FieldPathSelector = 1
	Region_FieldPathSelectorTitle    Region_FieldPathSelector = 2
)

func (s Region_FieldPathSelector) String() string {
	switch s {
	case Region_FieldPathSelectorName:
		return "name"
	case Region_FieldPathSelectorMetadata:
		return "metadata"
	case Region_FieldPathSelectorTitle:
		return "title"
	default:
		panic(fmt.Sprintf("Invalid selector for Region: %d", s))
	}
}

func BuildRegion_FieldPath(fp gotenobject.RawFieldPath) (Region_FieldPath, error) {
	if len(fp) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty field path for object Region")
	}
	if len(fp) == 1 {
		switch fp[0] {
		case "name":
			return &Region_FieldTerminalPath{selector: Region_FieldPathSelectorName}, nil
		case "metadata":
			return &Region_FieldTerminalPath{selector: Region_FieldPathSelectorMetadata}, nil
		case "title":
			return &Region_FieldTerminalPath{selector: Region_FieldPathSelectorTitle}, nil
		}
	} else {
		switch fp[0] {
		case "metadata":
			if subpath, err := meta.BuildMeta_FieldPath(fp[1:]); err != nil {
				return nil, err
			} else {
				return &Region_FieldSubPath{selector: Region_FieldPathSelectorMetadata, subPath: subpath}, nil
			}
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "unknown field path '%s' for object Region", fp)
}

func ParseRegion_FieldPath(rawField string) (Region_FieldPath, error) {
	fp, err := gotenobject.ParseRawFieldPath(rawField)
	if err != nil {
		return nil, err
	}
	return BuildRegion_FieldPath(fp)
}

func MustParseRegion_FieldPath(rawField string) Region_FieldPath {
	fp, err := ParseRegion_FieldPath(rawField)
	if err != nil {
		panic(err)
	}
	return fp
}

type Region_FieldTerminalPath struct {
	selector Region_FieldPathSelector
}

var _ Region_FieldPath = (*Region_FieldTerminalPath)(nil)

func (fp *Region_FieldTerminalPath) Selector() Region_FieldPathSelector {
	return fp.selector
}

// String returns path representation in proto convention
func (fp *Region_FieldTerminalPath) String() string {
	return fp.selector.String()
}

// JSONString returns path representation is JSON convention
func (fp *Region_FieldTerminalPath) JSONString() string {
	return strcase.ToLowerCamel(fp.String())
}

// Get returns all values pointed by specific field from source Region
func (fp *Region_FieldTerminalPath) Get(source *Region) (values []interface{}) {
	if source != nil {
		switch fp.selector {
		case Region_FieldPathSelectorName:
			if source.Name != nil {
				values = append(values, source.Name)
			}
		case Region_FieldPathSelectorMetadata:
			if source.Metadata != nil {
				values = append(values, source.Metadata)
			}
		case Region_FieldPathSelectorTitle:
			values = append(values, source.Title)
		default:
			panic(fmt.Sprintf("Invalid selector for Region: %d", fp.selector))
		}
	}
	return
}

func (fp *Region_FieldTerminalPath) GetRaw(source proto.Message) []interface{} {
	return fp.Get(source.(*Region))
}

// GetSingle returns value pointed by specific field of from source Region
func (fp *Region_FieldTerminalPath) GetSingle(source *Region) (interface{}, bool) {
	switch fp.selector {
	case Region_FieldPathSelectorName:
		res := source.GetName()
		return res, res != nil
	case Region_FieldPathSelectorMetadata:
		res := source.GetMetadata()
		return res, res != nil
	case Region_FieldPathSelectorTitle:
		return source.GetTitle(), source != nil
	default:
		panic(fmt.Sprintf("Invalid selector for Region: %d", fp.selector))
	}
}

func (fp *Region_FieldTerminalPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fp.GetSingle(source.(*Region))
}

// GetDefault returns a default value of the field type
func (fp *Region_FieldTerminalPath) GetDefault() interface{} {
	switch fp.selector {
	case Region_FieldPathSelectorName:
		return (*Name)(nil)
	case Region_FieldPathSelectorMetadata:
		return (*meta.Meta)(nil)
	case Region_FieldPathSelectorTitle:
		return ""
	default:
		panic(fmt.Sprintf("Invalid selector for Region: %d", fp.selector))
	}
}

func (fp *Region_FieldTerminalPath) ClearValue(item *Region) {
	if item != nil {
		switch fp.selector {
		case Region_FieldPathSelectorName:
			item.Name = nil
		case Region_FieldPathSelectorMetadata:
			item.Metadata = nil
		case Region_FieldPathSelectorTitle:
			item.Title = ""
		default:
			panic(fmt.Sprintf("Invalid selector for Region: %d", fp.selector))
		}
	}
}

func (fp *Region_FieldTerminalPath) ClearValueRaw(item proto.Message) {
	fp.ClearValue(item.(*Region))
}

// IsLeaf - whether field path is holds simple value
func (fp *Region_FieldTerminalPath) IsLeaf() bool {
	return fp.selector == Region_FieldPathSelectorName ||
		fp.selector == Region_FieldPathSelectorTitle
}

func (fp *Region_FieldTerminalPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	return []gotenobject.FieldPath{fp}
}

func (fp *Region_FieldTerminalPath) WithIValue(value interface{}) Region_FieldPathValue {
	switch fp.selector {
	case Region_FieldPathSelectorName:
		return &Region_FieldTerminalPathValue{Region_FieldTerminalPath: *fp, value: value.(*Name)}
	case Region_FieldPathSelectorMetadata:
		return &Region_FieldTerminalPathValue{Region_FieldTerminalPath: *fp, value: value.(*meta.Meta)}
	case Region_FieldPathSelectorTitle:
		return &Region_FieldTerminalPathValue{Region_FieldTerminalPath: *fp, value: value.(string)}
	default:
		panic(fmt.Sprintf("Invalid selector for Region: %d", fp.selector))
	}
}

func (fp *Region_FieldTerminalPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fp.WithIValue(value)
}

func (fp *Region_FieldTerminalPath) WithIArrayOfValues(values interface{}) Region_FieldPathArrayOfValues {
	fpaov := &Region_FieldTerminalPathArrayOfValues{Region_FieldTerminalPath: *fp}
	switch fp.selector {
	case Region_FieldPathSelectorName:
		return &Region_FieldTerminalPathArrayOfValues{Region_FieldTerminalPath: *fp, values: values.([]*Name)}
	case Region_FieldPathSelectorMetadata:
		return &Region_FieldTerminalPathArrayOfValues{Region_FieldTerminalPath: *fp, values: values.([]*meta.Meta)}
	case Region_FieldPathSelectorTitle:
		return &Region_FieldTerminalPathArrayOfValues{Region_FieldTerminalPath: *fp, values: values.([]string)}
	default:
		panic(fmt.Sprintf("Invalid selector for Region: %d", fp.selector))
	}
	return fpaov
}

func (fp *Region_FieldTerminalPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fp.WithIArrayOfValues(values)
}

func (fp *Region_FieldTerminalPath) WithIArrayItemValue(value interface{}) Region_FieldPathArrayItemValue {
	switch fp.selector {
	default:
		panic(fmt.Sprintf("Invalid selector for Region: %d", fp.selector))
	}
}

func (fp *Region_FieldTerminalPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fp.WithIArrayItemValue(value)
}

type Region_FieldSubPath struct {
	selector Region_FieldPathSelector
	subPath  gotenobject.FieldPath
}

var _ Region_FieldPath = (*Region_FieldSubPath)(nil)

func (fps *Region_FieldSubPath) Selector() Region_FieldPathSelector {
	return fps.selector
}
func (fps *Region_FieldSubPath) AsMetadataSubPath() (meta.Meta_FieldPath, bool) {
	res, ok := fps.subPath.(meta.Meta_FieldPath)
	return res, ok
}

// String returns path representation in proto convention
func (fps *Region_FieldSubPath) String() string {
	return fps.selector.String() + "." + fps.subPath.String()
}

// JSONString returns path representation is JSON convention
func (fps *Region_FieldSubPath) JSONString() string {
	return strcase.ToLowerCamel(fps.selector.String()) + "." + fps.subPath.JSONString()
}

// Get returns all values pointed by selected field from source Region
func (fps *Region_FieldSubPath) Get(source *Region) (values []interface{}) {
	switch fps.selector {
	case Region_FieldPathSelectorMetadata:
		values = append(values, fps.subPath.GetRaw(source.GetMetadata())...)
	default:
		panic(fmt.Sprintf("Invalid selector for Region: %d", fps.selector))
	}
	return
}

func (fps *Region_FieldSubPath) GetRaw(source proto.Message) []interface{} {
	return fps.Get(source.(*Region))
}

// GetSingle returns value of selected field from source Region
func (fps *Region_FieldSubPath) GetSingle(source *Region) (interface{}, bool) {
	switch fps.selector {
	case Region_FieldPathSelectorMetadata:
		if source.GetMetadata() == nil {
			return nil, false
		}
		return fps.subPath.GetSingleRaw(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for Region: %d", fps.selector))
	}
}

func (fps *Region_FieldSubPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fps.GetSingle(source.(*Region))
}

// GetDefault returns a default value of the field type
func (fps *Region_FieldSubPath) GetDefault() interface{} {
	return fps.subPath.GetDefault()
}

func (fps *Region_FieldSubPath) ClearValue(item *Region) {
	if item != nil {
		switch fps.selector {
		case Region_FieldPathSelectorMetadata:
			fps.subPath.ClearValueRaw(item.Metadata)
		default:
			panic(fmt.Sprintf("Invalid selector for Region: %d", fps.selector))
		}
	}
}

func (fps *Region_FieldSubPath) ClearValueRaw(item proto.Message) {
	fps.ClearValue(item.(*Region))
}

// IsLeaf - whether field path is holds simple value
func (fps *Region_FieldSubPath) IsLeaf() bool {
	return fps.subPath.IsLeaf()
}

func (fps *Region_FieldSubPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	iPaths := []gotenobject.FieldPath{&Region_FieldTerminalPath{selector: fps.selector}}
	iPaths = append(iPaths, fps.subPath.SplitIntoTerminalIPaths()...)
	return iPaths
}

func (fps *Region_FieldSubPath) WithIValue(value interface{}) Region_FieldPathValue {
	return &Region_FieldSubPathValue{fps, fps.subPath.WithRawIValue(value)}
}

func (fps *Region_FieldSubPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fps.WithIValue(value)
}

func (fps *Region_FieldSubPath) WithIArrayOfValues(values interface{}) Region_FieldPathArrayOfValues {
	return &Region_FieldSubPathArrayOfValues{fps, fps.subPath.WithRawIArrayOfValues(values)}
}

func (fps *Region_FieldSubPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fps.WithIArrayOfValues(values)
}

func (fps *Region_FieldSubPath) WithIArrayItemValue(value interface{}) Region_FieldPathArrayItemValue {
	return &Region_FieldSubPathArrayItemValue{fps, fps.subPath.WithRawIArrayItemValue(value)}
}

func (fps *Region_FieldSubPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fps.WithIArrayItemValue(value)
}

// Region_FieldPathValue allows storing values for Region fields according to their type
type Region_FieldPathValue interface {
	Region_FieldPath
	gotenobject.FieldPathValue
	SetTo(target **Region)
	CompareWith(*Region) (cmp int, comparable bool)
}

func ParseRegion_FieldPathValue(pathStr, valueStr string) (Region_FieldPathValue, error) {
	fp, err := ParseRegion_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpv, err := gotenobject.ParseFieldPathValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Region field path value from %s: %v", valueStr, err)
	}
	return fpv.(Region_FieldPathValue), nil
}

func MustParseRegion_FieldPathValue(pathStr, valueStr string) Region_FieldPathValue {
	fpv, err := ParseRegion_FieldPathValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpv
}

type Region_FieldTerminalPathValue struct {
	Region_FieldTerminalPath
	value interface{}
}

var _ Region_FieldPathValue = (*Region_FieldTerminalPathValue)(nil)

// GetRawValue returns raw value stored under selected path for 'Region' as interface{}
func (fpv *Region_FieldTerminalPathValue) GetRawValue() interface{} {
	return fpv.value
}
func (fpv *Region_FieldTerminalPathValue) AsNameValue() (*Name, bool) {
	res, ok := fpv.value.(*Name)
	return res, ok
}
func (fpv *Region_FieldTerminalPathValue) AsMetadataValue() (*meta.Meta, bool) {
	res, ok := fpv.value.(*meta.Meta)
	return res, ok
}
func (fpv *Region_FieldTerminalPathValue) AsTitleValue() (string, bool) {
	res, ok := fpv.value.(string)
	return res, ok
}

// SetTo stores value for selected field for object Region
func (fpv *Region_FieldTerminalPathValue) SetTo(target **Region) {
	if *target == nil {
		*target = new(Region)
	}
	switch fpv.selector {
	case Region_FieldPathSelectorName:
		(*target).Name = fpv.value.(*Name)
	case Region_FieldPathSelectorMetadata:
		(*target).Metadata = fpv.value.(*meta.Meta)
	case Region_FieldPathSelectorTitle:
		(*target).Title = fpv.value.(string)
	default:
		panic(fmt.Sprintf("Invalid selector for Region: %d", fpv.selector))
	}
}

func (fpv *Region_FieldTerminalPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*Region)
	fpv.SetTo(&typedObject)
}

// CompareWith compares value in the 'Region_FieldTerminalPathValue' with the value under path in 'Region'.
func (fpv *Region_FieldTerminalPathValue) CompareWith(source *Region) (int, bool) {
	switch fpv.selector {
	case Region_FieldPathSelectorName:
		leftValue := fpv.value.(*Name)
		rightValue := source.GetName()
		if leftValue == nil {
			if rightValue != nil {
				return -1, true
			}
			return 0, true
		}
		if rightValue == nil {
			return 1, true
		}
		if leftValue.String() == rightValue.String() {
			return 0, true
		} else if leftValue.String() < rightValue.String() {
			return -1, true
		} else {
			return 1, true
		}
	case Region_FieldPathSelectorMetadata:
		return 0, false
	case Region_FieldPathSelectorTitle:
		leftValue := fpv.value.(string)
		rightValue := source.GetTitle()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	default:
		panic(fmt.Sprintf("Invalid selector for Region: %d", fpv.selector))
	}
}

func (fpv *Region_FieldTerminalPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpv.CompareWith(source.(*Region))
}

type Region_FieldSubPathValue struct {
	Region_FieldPath
	subPathValue gotenobject.FieldPathValue
}

var _ Region_FieldPathValue = (*Region_FieldSubPathValue)(nil)

func (fpvs *Region_FieldSubPathValue) AsMetadataPathValue() (meta.Meta_FieldPathValue, bool) {
	res, ok := fpvs.subPathValue.(meta.Meta_FieldPathValue)
	return res, ok
}

func (fpvs *Region_FieldSubPathValue) SetTo(target **Region) {
	if *target == nil {
		*target = new(Region)
	}
	switch fpvs.Selector() {
	case Region_FieldPathSelectorMetadata:
		fpvs.subPathValue.(meta.Meta_FieldPathValue).SetTo(&(*target).Metadata)
	default:
		panic(fmt.Sprintf("Invalid selector for Region: %d", fpvs.Selector()))
	}
}

func (fpvs *Region_FieldSubPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*Region)
	fpvs.SetTo(&typedObject)
}

func (fpvs *Region_FieldSubPathValue) GetRawValue() interface{} {
	return fpvs.subPathValue.GetRawValue()
}

func (fpvs *Region_FieldSubPathValue) CompareWith(source *Region) (int, bool) {
	switch fpvs.Selector() {
	case Region_FieldPathSelectorMetadata:
		return fpvs.subPathValue.(meta.Meta_FieldPathValue).CompareWith(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for Region: %d", fpvs.Selector()))
	}
}

func (fpvs *Region_FieldSubPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpvs.CompareWith(source.(*Region))
}

// Region_FieldPathArrayItemValue allows storing single item in Path-specific values for Region according to their type
// Present only for array (repeated) types.
type Region_FieldPathArrayItemValue interface {
	gotenobject.FieldPathArrayItemValue
	Region_FieldPath
	ContainsValue(*Region) bool
}

// ParseRegion_FieldPathArrayItemValue parses string and JSON-encoded value to its Value
func ParseRegion_FieldPathArrayItemValue(pathStr, valueStr string) (Region_FieldPathArrayItemValue, error) {
	fp, err := ParseRegion_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaiv, err := gotenobject.ParseFieldPathArrayItemValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Region field path array item value from %s: %v", valueStr, err)
	}
	return fpaiv.(Region_FieldPathArrayItemValue), nil
}

func MustParseRegion_FieldPathArrayItemValue(pathStr, valueStr string) Region_FieldPathArrayItemValue {
	fpaiv, err := ParseRegion_FieldPathArrayItemValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpaiv
}

type Region_FieldTerminalPathArrayItemValue struct {
	Region_FieldTerminalPath
	value interface{}
}

var _ Region_FieldPathArrayItemValue = (*Region_FieldTerminalPathArrayItemValue)(nil)

// GetRawValue returns stored element value for array in object Region as interface{}
func (fpaiv *Region_FieldTerminalPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaiv.value
}

func (fpaiv *Region_FieldTerminalPathArrayItemValue) GetSingle(source *Region) (interface{}, bool) {
	return nil, false
}

func (fpaiv *Region_FieldTerminalPathArrayItemValue) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpaiv.GetSingle(source.(*Region))
}

// Contains returns a boolean indicating if value that is being held is present in given 'Region'
func (fpaiv *Region_FieldTerminalPathArrayItemValue) ContainsValue(source *Region) bool {
	slice := fpaiv.Region_FieldTerminalPath.Get(source)
	for _, v := range slice {
		if asProtoMsg, ok := fpaiv.value.(proto.Message); ok {
			if proto.Equal(asProtoMsg, v.(proto.Message)) {
				return true
			}
		} else if reflect.DeepEqual(v, fpaiv.value) {
			return true
		}
	}
	return false
}

type Region_FieldSubPathArrayItemValue struct {
	Region_FieldPath
	subPathItemValue gotenobject.FieldPathArrayItemValue
}

// GetRawValue returns stored array item value
func (fpaivs *Region_FieldSubPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaivs.subPathItemValue.GetRawItemValue()
}
func (fpaivs *Region_FieldSubPathArrayItemValue) AsMetadataPathItemValue() (meta.Meta_FieldPathArrayItemValue, bool) {
	res, ok := fpaivs.subPathItemValue.(meta.Meta_FieldPathArrayItemValue)
	return res, ok
}

// Contains returns a boolean indicating if value that is being held is present in given 'Region'
func (fpaivs *Region_FieldSubPathArrayItemValue) ContainsValue(source *Region) bool {
	switch fpaivs.Selector() {
	case Region_FieldPathSelectorMetadata:
		return fpaivs.subPathItemValue.(meta.Meta_FieldPathArrayItemValue).ContainsValue(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for Region: %d", fpaivs.Selector()))
	}
}

// Region_FieldPathArrayOfValues allows storing slice of values for Region fields according to their type
type Region_FieldPathArrayOfValues interface {
	gotenobject.FieldPathArrayOfValues
	Region_FieldPath
}

func ParseRegion_FieldPathArrayOfValues(pathStr, valuesStr string) (Region_FieldPathArrayOfValues, error) {
	fp, err := ParseRegion_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaov, err := gotenobject.ParseFieldPathArrayOfValues(fp, valuesStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Region field path array of values from %s: %v", valuesStr, err)
	}
	return fpaov.(Region_FieldPathArrayOfValues), nil
}

func MustParseRegion_FieldPathArrayOfValues(pathStr, valuesStr string) Region_FieldPathArrayOfValues {
	fpaov, err := ParseRegion_FieldPathArrayOfValues(pathStr, valuesStr)
	if err != nil {
		panic(err)
	}
	return fpaov
}

type Region_FieldTerminalPathArrayOfValues struct {
	Region_FieldTerminalPath
	values interface{}
}

var _ Region_FieldPathArrayOfValues = (*Region_FieldTerminalPathArrayOfValues)(nil)

func (fpaov *Region_FieldTerminalPathArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpaov.selector {
	case Region_FieldPathSelectorName:
		for _, v := range fpaov.values.([]*Name) {
			values = append(values, v)
		}
	case Region_FieldPathSelectorMetadata:
		for _, v := range fpaov.values.([]*meta.Meta) {
			values = append(values, v)
		}
	case Region_FieldPathSelectorTitle:
		for _, v := range fpaov.values.([]string) {
			values = append(values, v)
		}
	}
	return
}
func (fpaov *Region_FieldTerminalPathArrayOfValues) AsNameArrayOfValues() ([]*Name, bool) {
	res, ok := fpaov.values.([]*Name)
	return res, ok
}
func (fpaov *Region_FieldTerminalPathArrayOfValues) AsMetadataArrayOfValues() ([]*meta.Meta, bool) {
	res, ok := fpaov.values.([]*meta.Meta)
	return res, ok
}
func (fpaov *Region_FieldTerminalPathArrayOfValues) AsTitleArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}

type Region_FieldSubPathArrayOfValues struct {
	Region_FieldPath
	subPathArrayOfValues gotenobject.FieldPathArrayOfValues
}

var _ Region_FieldPathArrayOfValues = (*Region_FieldSubPathArrayOfValues)(nil)

func (fpsaov *Region_FieldSubPathArrayOfValues) GetRawValues() []interface{} {
	return fpsaov.subPathArrayOfValues.GetRawValues()
}
func (fpsaov *Region_FieldSubPathArrayOfValues) AsMetadataPathArrayOfValues() (meta.Meta_FieldPathArrayOfValues, bool) {
	res, ok := fpsaov.subPathArrayOfValues.(meta.Meta_FieldPathArrayOfValues)
	return res, ok
}
