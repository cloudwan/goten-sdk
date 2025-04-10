// Code generated by protoc-gen-goten-object
// File: goten/meta-service/proto/v1/common.proto
// DO NOT EDIT!!!

package common

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
import ()

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
var ()

// FieldPath provides implementation to handle
// https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto
type LabelledDomain_FieldPath interface {
	gotenobject.FieldPath
	Selector() LabelledDomain_FieldPathSelector
	Get(source *LabelledDomain) []interface{}
	GetSingle(source *LabelledDomain) (interface{}, bool)
	ClearValue(item *LabelledDomain)

	// Those methods build corresponding LabelledDomain_FieldPathValue
	// (or array of values) and holds passed value. Panics if injected type is incorrect.
	WithIValue(value interface{}) LabelledDomain_FieldPathValue
	WithIArrayOfValues(values interface{}) LabelledDomain_FieldPathArrayOfValues
	WithIArrayItemValue(value interface{}) LabelledDomain_FieldPathArrayItemValue
}

type LabelledDomain_FieldPathSelector int32

const (
	LabelledDomain_FieldPathSelectorLabel            LabelledDomain_FieldPathSelector = 0
	LabelledDomain_FieldPathSelectorDomain           LabelledDomain_FieldPathSelector = 1
	LabelledDomain_FieldPathSelectorAvailableMixins  LabelledDomain_FieldPathSelector = 2
	LabelledDomain_FieldPathSelectorWebGrpcAvailable LabelledDomain_FieldPathSelector = 3
	LabelledDomain_FieldPathSelectorRestApiAvailable LabelledDomain_FieldPathSelector = 4
	LabelledDomain_FieldPathSelectorIsPrivate        LabelledDomain_FieldPathSelector = 5
)

func (s LabelledDomain_FieldPathSelector) String() string {
	switch s {
	case LabelledDomain_FieldPathSelectorLabel:
		return "label"
	case LabelledDomain_FieldPathSelectorDomain:
		return "domain"
	case LabelledDomain_FieldPathSelectorAvailableMixins:
		return "available_mixins"
	case LabelledDomain_FieldPathSelectorWebGrpcAvailable:
		return "web_grpc_available"
	case LabelledDomain_FieldPathSelectorRestApiAvailable:
		return "rest_api_available"
	case LabelledDomain_FieldPathSelectorIsPrivate:
		return "is_private"
	default:
		panic(fmt.Sprintf("Invalid selector for LabelledDomain: %d", s))
	}
}

func BuildLabelledDomain_FieldPath(fp gotenobject.RawFieldPath) (LabelledDomain_FieldPath, error) {
	if len(fp) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty field path for object LabelledDomain")
	}
	if len(fp) == 1 {
		switch fp[0] {
		case "label":
			return &LabelledDomain_FieldTerminalPath{selector: LabelledDomain_FieldPathSelectorLabel}, nil
		case "domain":
			return &LabelledDomain_FieldTerminalPath{selector: LabelledDomain_FieldPathSelectorDomain}, nil
		case "available_mixins", "availableMixins", "available-mixins":
			return &LabelledDomain_FieldTerminalPath{selector: LabelledDomain_FieldPathSelectorAvailableMixins}, nil
		case "web_grpc_available", "webGrpcAvailable", "web-grpc-available":
			return &LabelledDomain_FieldTerminalPath{selector: LabelledDomain_FieldPathSelectorWebGrpcAvailable}, nil
		case "rest_api_available", "restApiAvailable", "rest-api-available":
			return &LabelledDomain_FieldTerminalPath{selector: LabelledDomain_FieldPathSelectorRestApiAvailable}, nil
		case "is_private", "isPrivate", "is-private":
			return &LabelledDomain_FieldTerminalPath{selector: LabelledDomain_FieldPathSelectorIsPrivate}, nil
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "unknown field path '%s' for object LabelledDomain", fp)
}

func ParseLabelledDomain_FieldPath(rawField string) (LabelledDomain_FieldPath, error) {
	fp, err := gotenobject.ParseRawFieldPath(rawField)
	if err != nil {
		return nil, err
	}
	return BuildLabelledDomain_FieldPath(fp)
}

func MustParseLabelledDomain_FieldPath(rawField string) LabelledDomain_FieldPath {
	fp, err := ParseLabelledDomain_FieldPath(rawField)
	if err != nil {
		panic(err)
	}
	return fp
}

type LabelledDomain_FieldTerminalPath struct {
	selector LabelledDomain_FieldPathSelector
}

var _ LabelledDomain_FieldPath = (*LabelledDomain_FieldTerminalPath)(nil)

func (fp *LabelledDomain_FieldTerminalPath) Selector() LabelledDomain_FieldPathSelector {
	return fp.selector
}

// String returns path representation in proto convention
func (fp *LabelledDomain_FieldTerminalPath) String() string {
	return fp.selector.String()
}

// JSONString returns path representation is JSON convention
func (fp *LabelledDomain_FieldTerminalPath) JSONString() string {
	return strcase.ToLowerCamel(fp.String())
}

// Get returns all values pointed by specific field from source LabelledDomain
func (fp *LabelledDomain_FieldTerminalPath) Get(source *LabelledDomain) (values []interface{}) {
	if source != nil {
		switch fp.selector {
		case LabelledDomain_FieldPathSelectorLabel:
			values = append(values, source.Label)
		case LabelledDomain_FieldPathSelectorDomain:
			values = append(values, source.Domain)
		case LabelledDomain_FieldPathSelectorAvailableMixins:
			for _, value := range source.GetAvailableMixins() {
				values = append(values, value)
			}
		case LabelledDomain_FieldPathSelectorWebGrpcAvailable:
			values = append(values, source.WebGrpcAvailable)
		case LabelledDomain_FieldPathSelectorRestApiAvailable:
			values = append(values, source.RestApiAvailable)
		case LabelledDomain_FieldPathSelectorIsPrivate:
			values = append(values, source.IsPrivate)
		default:
			panic(fmt.Sprintf("Invalid selector for LabelledDomain: %d", fp.selector))
		}
	}
	return
}

func (fp *LabelledDomain_FieldTerminalPath) GetRaw(source proto.Message) []interface{} {
	return fp.Get(source.(*LabelledDomain))
}

// GetSingle returns value pointed by specific field of from source LabelledDomain
func (fp *LabelledDomain_FieldTerminalPath) GetSingle(source *LabelledDomain) (interface{}, bool) {
	switch fp.selector {
	case LabelledDomain_FieldPathSelectorLabel:
		return source.GetLabel(), source != nil
	case LabelledDomain_FieldPathSelectorDomain:
		return source.GetDomain(), source != nil
	case LabelledDomain_FieldPathSelectorAvailableMixins:
		res := source.GetAvailableMixins()
		return res, res != nil
	case LabelledDomain_FieldPathSelectorWebGrpcAvailable:
		return source.GetWebGrpcAvailable(), source != nil
	case LabelledDomain_FieldPathSelectorRestApiAvailable:
		return source.GetRestApiAvailable(), source != nil
	case LabelledDomain_FieldPathSelectorIsPrivate:
		return source.GetIsPrivate(), source != nil
	default:
		panic(fmt.Sprintf("Invalid selector for LabelledDomain: %d", fp.selector))
	}
}

func (fp *LabelledDomain_FieldTerminalPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fp.GetSingle(source.(*LabelledDomain))
}

// GetDefault returns a default value of the field type
func (fp *LabelledDomain_FieldTerminalPath) GetDefault() interface{} {
	switch fp.selector {
	case LabelledDomain_FieldPathSelectorLabel:
		return ""
	case LabelledDomain_FieldPathSelectorDomain:
		return ""
	case LabelledDomain_FieldPathSelectorAvailableMixins:
		return ([]string)(nil)
	case LabelledDomain_FieldPathSelectorWebGrpcAvailable:
		return false
	case LabelledDomain_FieldPathSelectorRestApiAvailable:
		return false
	case LabelledDomain_FieldPathSelectorIsPrivate:
		return false
	default:
		panic(fmt.Sprintf("Invalid selector for LabelledDomain: %d", fp.selector))
	}
}

func (fp *LabelledDomain_FieldTerminalPath) ClearValue(item *LabelledDomain) {
	if item != nil {
		switch fp.selector {
		case LabelledDomain_FieldPathSelectorLabel:
			item.Label = ""
		case LabelledDomain_FieldPathSelectorDomain:
			item.Domain = ""
		case LabelledDomain_FieldPathSelectorAvailableMixins:
			item.AvailableMixins = nil
		case LabelledDomain_FieldPathSelectorWebGrpcAvailable:
			item.WebGrpcAvailable = false
		case LabelledDomain_FieldPathSelectorRestApiAvailable:
			item.RestApiAvailable = false
		case LabelledDomain_FieldPathSelectorIsPrivate:
			item.IsPrivate = false
		default:
			panic(fmt.Sprintf("Invalid selector for LabelledDomain: %d", fp.selector))
		}
	}
}

func (fp *LabelledDomain_FieldTerminalPath) ClearValueRaw(item proto.Message) {
	fp.ClearValue(item.(*LabelledDomain))
}

// IsLeaf - whether field path is holds simple value
func (fp *LabelledDomain_FieldTerminalPath) IsLeaf() bool {
	return fp.selector == LabelledDomain_FieldPathSelectorLabel ||
		fp.selector == LabelledDomain_FieldPathSelectorDomain ||
		fp.selector == LabelledDomain_FieldPathSelectorAvailableMixins ||
		fp.selector == LabelledDomain_FieldPathSelectorWebGrpcAvailable ||
		fp.selector == LabelledDomain_FieldPathSelectorRestApiAvailable ||
		fp.selector == LabelledDomain_FieldPathSelectorIsPrivate
}

func (fp *LabelledDomain_FieldTerminalPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	return []gotenobject.FieldPath{fp}
}

func (fp *LabelledDomain_FieldTerminalPath) WithIValue(value interface{}) LabelledDomain_FieldPathValue {
	switch fp.selector {
	case LabelledDomain_FieldPathSelectorLabel:
		return &LabelledDomain_FieldTerminalPathValue{LabelledDomain_FieldTerminalPath: *fp, value: value.(string)}
	case LabelledDomain_FieldPathSelectorDomain:
		return &LabelledDomain_FieldTerminalPathValue{LabelledDomain_FieldTerminalPath: *fp, value: value.(string)}
	case LabelledDomain_FieldPathSelectorAvailableMixins:
		return &LabelledDomain_FieldTerminalPathValue{LabelledDomain_FieldTerminalPath: *fp, value: value.([]string)}
	case LabelledDomain_FieldPathSelectorWebGrpcAvailable:
		return &LabelledDomain_FieldTerminalPathValue{LabelledDomain_FieldTerminalPath: *fp, value: value.(bool)}
	case LabelledDomain_FieldPathSelectorRestApiAvailable:
		return &LabelledDomain_FieldTerminalPathValue{LabelledDomain_FieldTerminalPath: *fp, value: value.(bool)}
	case LabelledDomain_FieldPathSelectorIsPrivate:
		return &LabelledDomain_FieldTerminalPathValue{LabelledDomain_FieldTerminalPath: *fp, value: value.(bool)}
	default:
		panic(fmt.Sprintf("Invalid selector for LabelledDomain: %d", fp.selector))
	}
}

func (fp *LabelledDomain_FieldTerminalPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fp.WithIValue(value)
}

func (fp *LabelledDomain_FieldTerminalPath) WithIArrayOfValues(values interface{}) LabelledDomain_FieldPathArrayOfValues {
	fpaov := &LabelledDomain_FieldTerminalPathArrayOfValues{LabelledDomain_FieldTerminalPath: *fp}
	switch fp.selector {
	case LabelledDomain_FieldPathSelectorLabel:
		return &LabelledDomain_FieldTerminalPathArrayOfValues{LabelledDomain_FieldTerminalPath: *fp, values: values.([]string)}
	case LabelledDomain_FieldPathSelectorDomain:
		return &LabelledDomain_FieldTerminalPathArrayOfValues{LabelledDomain_FieldTerminalPath: *fp, values: values.([]string)}
	case LabelledDomain_FieldPathSelectorAvailableMixins:
		return &LabelledDomain_FieldTerminalPathArrayOfValues{LabelledDomain_FieldTerminalPath: *fp, values: values.([][]string)}
	case LabelledDomain_FieldPathSelectorWebGrpcAvailable:
		return &LabelledDomain_FieldTerminalPathArrayOfValues{LabelledDomain_FieldTerminalPath: *fp, values: values.([]bool)}
	case LabelledDomain_FieldPathSelectorRestApiAvailable:
		return &LabelledDomain_FieldTerminalPathArrayOfValues{LabelledDomain_FieldTerminalPath: *fp, values: values.([]bool)}
	case LabelledDomain_FieldPathSelectorIsPrivate:
		return &LabelledDomain_FieldTerminalPathArrayOfValues{LabelledDomain_FieldTerminalPath: *fp, values: values.([]bool)}
	default:
		panic(fmt.Sprintf("Invalid selector for LabelledDomain: %d", fp.selector))
	}
	return fpaov
}

func (fp *LabelledDomain_FieldTerminalPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fp.WithIArrayOfValues(values)
}

func (fp *LabelledDomain_FieldTerminalPath) WithIArrayItemValue(value interface{}) LabelledDomain_FieldPathArrayItemValue {
	switch fp.selector {
	case LabelledDomain_FieldPathSelectorAvailableMixins:
		return &LabelledDomain_FieldTerminalPathArrayItemValue{LabelledDomain_FieldTerminalPath: *fp, value: value.(string)}
	default:
		panic(fmt.Sprintf("Invalid selector for LabelledDomain: %d", fp.selector))
	}
}

func (fp *LabelledDomain_FieldTerminalPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fp.WithIArrayItemValue(value)
}

// LabelledDomain_FieldPathValue allows storing values for LabelledDomain fields according to their type
type LabelledDomain_FieldPathValue interface {
	LabelledDomain_FieldPath
	gotenobject.FieldPathValue
	SetTo(target **LabelledDomain)
	CompareWith(*LabelledDomain) (cmp int, comparable bool)
}

func ParseLabelledDomain_FieldPathValue(pathStr, valueStr string) (LabelledDomain_FieldPathValue, error) {
	fp, err := ParseLabelledDomain_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpv, err := gotenobject.ParseFieldPathValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing LabelledDomain field path value from %s: %v", valueStr, err)
	}
	return fpv.(LabelledDomain_FieldPathValue), nil
}

func MustParseLabelledDomain_FieldPathValue(pathStr, valueStr string) LabelledDomain_FieldPathValue {
	fpv, err := ParseLabelledDomain_FieldPathValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpv
}

type LabelledDomain_FieldTerminalPathValue struct {
	LabelledDomain_FieldTerminalPath
	value interface{}
}

var _ LabelledDomain_FieldPathValue = (*LabelledDomain_FieldTerminalPathValue)(nil)

// GetRawValue returns raw value stored under selected path for 'LabelledDomain' as interface{}
func (fpv *LabelledDomain_FieldTerminalPathValue) GetRawValue() interface{} {
	return fpv.value
}
func (fpv *LabelledDomain_FieldTerminalPathValue) AsLabelValue() (string, bool) {
	res, ok := fpv.value.(string)
	return res, ok
}
func (fpv *LabelledDomain_FieldTerminalPathValue) AsDomainValue() (string, bool) {
	res, ok := fpv.value.(string)
	return res, ok
}
func (fpv *LabelledDomain_FieldTerminalPathValue) AsAvailableMixinsValue() ([]string, bool) {
	res, ok := fpv.value.([]string)
	return res, ok
}
func (fpv *LabelledDomain_FieldTerminalPathValue) AsWebGrpcAvailableValue() (bool, bool) {
	res, ok := fpv.value.(bool)
	return res, ok
}
func (fpv *LabelledDomain_FieldTerminalPathValue) AsRestApiAvailableValue() (bool, bool) {
	res, ok := fpv.value.(bool)
	return res, ok
}
func (fpv *LabelledDomain_FieldTerminalPathValue) AsIsPrivateValue() (bool, bool) {
	res, ok := fpv.value.(bool)
	return res, ok
}

// SetTo stores value for selected field for object LabelledDomain
func (fpv *LabelledDomain_FieldTerminalPathValue) SetTo(target **LabelledDomain) {
	if *target == nil {
		*target = new(LabelledDomain)
	}
	switch fpv.selector {
	case LabelledDomain_FieldPathSelectorLabel:
		(*target).Label = fpv.value.(string)
	case LabelledDomain_FieldPathSelectorDomain:
		(*target).Domain = fpv.value.(string)
	case LabelledDomain_FieldPathSelectorAvailableMixins:
		(*target).AvailableMixins = fpv.value.([]string)
	case LabelledDomain_FieldPathSelectorWebGrpcAvailable:
		(*target).WebGrpcAvailable = fpv.value.(bool)
	case LabelledDomain_FieldPathSelectorRestApiAvailable:
		(*target).RestApiAvailable = fpv.value.(bool)
	case LabelledDomain_FieldPathSelectorIsPrivate:
		(*target).IsPrivate = fpv.value.(bool)
	default:
		panic(fmt.Sprintf("Invalid selector for LabelledDomain: %d", fpv.selector))
	}
}

func (fpv *LabelledDomain_FieldTerminalPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*LabelledDomain)
	fpv.SetTo(&typedObject)
}

// CompareWith compares value in the 'LabelledDomain_FieldTerminalPathValue' with the value under path in 'LabelledDomain'.
func (fpv *LabelledDomain_FieldTerminalPathValue) CompareWith(source *LabelledDomain) (int, bool) {
	switch fpv.selector {
	case LabelledDomain_FieldPathSelectorLabel:
		leftValue := fpv.value.(string)
		rightValue := source.GetLabel()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case LabelledDomain_FieldPathSelectorDomain:
		leftValue := fpv.value.(string)
		rightValue := source.GetDomain()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case LabelledDomain_FieldPathSelectorAvailableMixins:
		return 0, false
	case LabelledDomain_FieldPathSelectorWebGrpcAvailable:
		leftValue := fpv.value.(bool)
		rightValue := source.GetWebGrpcAvailable()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if !(leftValue) && (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case LabelledDomain_FieldPathSelectorRestApiAvailable:
		leftValue := fpv.value.(bool)
		rightValue := source.GetRestApiAvailable()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if !(leftValue) && (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case LabelledDomain_FieldPathSelectorIsPrivate:
		leftValue := fpv.value.(bool)
		rightValue := source.GetIsPrivate()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if !(leftValue) && (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	default:
		panic(fmt.Sprintf("Invalid selector for LabelledDomain: %d", fpv.selector))
	}
}

func (fpv *LabelledDomain_FieldTerminalPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpv.CompareWith(source.(*LabelledDomain))
}

// LabelledDomain_FieldPathArrayItemValue allows storing single item in Path-specific values for LabelledDomain according to their type
// Present only for array (repeated) types.
type LabelledDomain_FieldPathArrayItemValue interface {
	gotenobject.FieldPathArrayItemValue
	LabelledDomain_FieldPath
	ContainsValue(*LabelledDomain) bool
}

// ParseLabelledDomain_FieldPathArrayItemValue parses string and JSON-encoded value to its Value
func ParseLabelledDomain_FieldPathArrayItemValue(pathStr, valueStr string) (LabelledDomain_FieldPathArrayItemValue, error) {
	fp, err := ParseLabelledDomain_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaiv, err := gotenobject.ParseFieldPathArrayItemValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing LabelledDomain field path array item value from %s: %v", valueStr, err)
	}
	return fpaiv.(LabelledDomain_FieldPathArrayItemValue), nil
}

func MustParseLabelledDomain_FieldPathArrayItemValue(pathStr, valueStr string) LabelledDomain_FieldPathArrayItemValue {
	fpaiv, err := ParseLabelledDomain_FieldPathArrayItemValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpaiv
}

type LabelledDomain_FieldTerminalPathArrayItemValue struct {
	LabelledDomain_FieldTerminalPath
	value interface{}
}

var _ LabelledDomain_FieldPathArrayItemValue = (*LabelledDomain_FieldTerminalPathArrayItemValue)(nil)

// GetRawValue returns stored element value for array in object LabelledDomain as interface{}
func (fpaiv *LabelledDomain_FieldTerminalPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaiv.value
}
func (fpaiv *LabelledDomain_FieldTerminalPathArrayItemValue) AsAvailableMixinsItemValue() (string, bool) {
	res, ok := fpaiv.value.(string)
	return res, ok
}

func (fpaiv *LabelledDomain_FieldTerminalPathArrayItemValue) GetSingle(source *LabelledDomain) (interface{}, bool) {
	return nil, false
}

func (fpaiv *LabelledDomain_FieldTerminalPathArrayItemValue) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpaiv.GetSingle(source.(*LabelledDomain))
}

// Contains returns a boolean indicating if value that is being held is present in given 'LabelledDomain'
func (fpaiv *LabelledDomain_FieldTerminalPathArrayItemValue) ContainsValue(source *LabelledDomain) bool {
	slice := fpaiv.LabelledDomain_FieldTerminalPath.Get(source)
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

// LabelledDomain_FieldPathArrayOfValues allows storing slice of values for LabelledDomain fields according to their type
type LabelledDomain_FieldPathArrayOfValues interface {
	gotenobject.FieldPathArrayOfValues
	LabelledDomain_FieldPath
}

func ParseLabelledDomain_FieldPathArrayOfValues(pathStr, valuesStr string) (LabelledDomain_FieldPathArrayOfValues, error) {
	fp, err := ParseLabelledDomain_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaov, err := gotenobject.ParseFieldPathArrayOfValues(fp, valuesStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing LabelledDomain field path array of values from %s: %v", valuesStr, err)
	}
	return fpaov.(LabelledDomain_FieldPathArrayOfValues), nil
}

func MustParseLabelledDomain_FieldPathArrayOfValues(pathStr, valuesStr string) LabelledDomain_FieldPathArrayOfValues {
	fpaov, err := ParseLabelledDomain_FieldPathArrayOfValues(pathStr, valuesStr)
	if err != nil {
		panic(err)
	}
	return fpaov
}

type LabelledDomain_FieldTerminalPathArrayOfValues struct {
	LabelledDomain_FieldTerminalPath
	values interface{}
}

var _ LabelledDomain_FieldPathArrayOfValues = (*LabelledDomain_FieldTerminalPathArrayOfValues)(nil)

func (fpaov *LabelledDomain_FieldTerminalPathArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpaov.selector {
	case LabelledDomain_FieldPathSelectorLabel:
		for _, v := range fpaov.values.([]string) {
			values = append(values, v)
		}
	case LabelledDomain_FieldPathSelectorDomain:
		for _, v := range fpaov.values.([]string) {
			values = append(values, v)
		}
	case LabelledDomain_FieldPathSelectorAvailableMixins:
		for _, v := range fpaov.values.([][]string) {
			values = append(values, v)
		}
	case LabelledDomain_FieldPathSelectorWebGrpcAvailable:
		for _, v := range fpaov.values.([]bool) {
			values = append(values, v)
		}
	case LabelledDomain_FieldPathSelectorRestApiAvailable:
		for _, v := range fpaov.values.([]bool) {
			values = append(values, v)
		}
	case LabelledDomain_FieldPathSelectorIsPrivate:
		for _, v := range fpaov.values.([]bool) {
			values = append(values, v)
		}
	}
	return
}
func (fpaov *LabelledDomain_FieldTerminalPathArrayOfValues) AsLabelArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}
func (fpaov *LabelledDomain_FieldTerminalPathArrayOfValues) AsDomainArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}
func (fpaov *LabelledDomain_FieldTerminalPathArrayOfValues) AsAvailableMixinsArrayOfValues() ([][]string, bool) {
	res, ok := fpaov.values.([][]string)
	return res, ok
}
func (fpaov *LabelledDomain_FieldTerminalPathArrayOfValues) AsWebGrpcAvailableArrayOfValues() ([]bool, bool) {
	res, ok := fpaov.values.([]bool)
	return res, ok
}
func (fpaov *LabelledDomain_FieldTerminalPathArrayOfValues) AsRestApiAvailableArrayOfValues() ([]bool, bool) {
	res, ok := fpaov.values.([]bool)
	return res, ok
}
func (fpaov *LabelledDomain_FieldTerminalPathArrayOfValues) AsIsPrivateArrayOfValues() ([]bool, bool) {
	res, ok := fpaov.values.([]bool)
	return res, ok
}
