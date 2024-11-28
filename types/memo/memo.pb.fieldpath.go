// Code generated by protoc-gen-goten-object
// File: goten/types/memo.proto
// DO NOT EDIT!!!

package memo

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
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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
	_ = &timestamppb.Timestamp{}
)

// FieldPath provides implementation to handle
// https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto
type Memo_FieldPath interface {
	gotenobject.FieldPath
	Selector() Memo_FieldPathSelector
	Get(source *Memo) []interface{}
	GetSingle(source *Memo) (interface{}, bool)
	ClearValue(item *Memo)

	// Those methods build corresponding Memo_FieldPathValue
	// (or array of values) and holds passed value. Panics if injected type is incorrect.
	WithIValue(value interface{}) Memo_FieldPathValue
	WithIArrayOfValues(values interface{}) Memo_FieldPathArrayOfValues
	WithIArrayItemValue(value interface{}) Memo_FieldPathArrayItemValue
}

type Memo_FieldPathSelector int32

const (
	Memo_FieldPathSelectorCreateTime Memo_FieldPathSelector = 0
	Memo_FieldPathSelectorUpdateTime Memo_FieldPathSelector = 1
	Memo_FieldPathSelectorMessage    Memo_FieldPathSelector = 2
	Memo_FieldPathSelectorCreatedBy  Memo_FieldPathSelector = 3
)

func (s Memo_FieldPathSelector) String() string {
	switch s {
	case Memo_FieldPathSelectorCreateTime:
		return "create_time"
	case Memo_FieldPathSelectorUpdateTime:
		return "update_time"
	case Memo_FieldPathSelectorMessage:
		return "message"
	case Memo_FieldPathSelectorCreatedBy:
		return "created_by"
	default:
		panic(fmt.Sprintf("Invalid selector for Memo: %d", s))
	}
}

func BuildMemo_FieldPath(fp gotenobject.RawFieldPath) (Memo_FieldPath, error) {
	if len(fp) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty field path for object Memo")
	}
	if len(fp) == 1 {
		switch fp[0] {
		case "create_time", "createTime", "create-time":
			return &Memo_FieldTerminalPath{selector: Memo_FieldPathSelectorCreateTime}, nil
		case "update_time", "updateTime", "update-time":
			return &Memo_FieldTerminalPath{selector: Memo_FieldPathSelectorUpdateTime}, nil
		case "message":
			return &Memo_FieldTerminalPath{selector: Memo_FieldPathSelectorMessage}, nil
		case "created_by", "createdBy", "created-by":
			return &Memo_FieldTerminalPath{selector: Memo_FieldPathSelectorCreatedBy}, nil
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "unknown field path '%s' for object Memo", fp)
}

func ParseMemo_FieldPath(rawField string) (Memo_FieldPath, error) {
	fp, err := gotenobject.ParseRawFieldPath(rawField)
	if err != nil {
		return nil, err
	}
	return BuildMemo_FieldPath(fp)
}

func MustParseMemo_FieldPath(rawField string) Memo_FieldPath {
	fp, err := ParseMemo_FieldPath(rawField)
	if err != nil {
		panic(err)
	}
	return fp
}

type Memo_FieldTerminalPath struct {
	selector Memo_FieldPathSelector
}

var _ Memo_FieldPath = (*Memo_FieldTerminalPath)(nil)

func (fp *Memo_FieldTerminalPath) Selector() Memo_FieldPathSelector {
	return fp.selector
}

// String returns path representation in proto convention
func (fp *Memo_FieldTerminalPath) String() string {
	return fp.selector.String()
}

// JSONString returns path representation is JSON convention
func (fp *Memo_FieldTerminalPath) JSONString() string {
	return strcase.ToLowerCamel(fp.String())
}

// Get returns all values pointed by specific field from source Memo
func (fp *Memo_FieldTerminalPath) Get(source *Memo) (values []interface{}) {
	if source != nil {
		switch fp.selector {
		case Memo_FieldPathSelectorCreateTime:
			if source.CreateTime != nil {
				values = append(values, source.CreateTime)
			}
		case Memo_FieldPathSelectorUpdateTime:
			if source.UpdateTime != nil {
				values = append(values, source.UpdateTime)
			}
		case Memo_FieldPathSelectorMessage:
			values = append(values, source.Message)
		case Memo_FieldPathSelectorCreatedBy:
			values = append(values, source.CreatedBy)
		default:
			panic(fmt.Sprintf("Invalid selector for Memo: %d", fp.selector))
		}
	}
	return
}

func (fp *Memo_FieldTerminalPath) GetRaw(source proto.Message) []interface{} {
	return fp.Get(source.(*Memo))
}

// GetSingle returns value pointed by specific field of from source Memo
func (fp *Memo_FieldTerminalPath) GetSingle(source *Memo) (interface{}, bool) {
	switch fp.selector {
	case Memo_FieldPathSelectorCreateTime:
		res := source.GetCreateTime()
		return res, res != nil
	case Memo_FieldPathSelectorUpdateTime:
		res := source.GetUpdateTime()
		return res, res != nil
	case Memo_FieldPathSelectorMessage:
		return source.GetMessage(), source != nil
	case Memo_FieldPathSelectorCreatedBy:
		return source.GetCreatedBy(), source != nil
	default:
		panic(fmt.Sprintf("Invalid selector for Memo: %d", fp.selector))
	}
}

func (fp *Memo_FieldTerminalPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fp.GetSingle(source.(*Memo))
}

// GetDefault returns a default value of the field type
func (fp *Memo_FieldTerminalPath) GetDefault() interface{} {
	switch fp.selector {
	case Memo_FieldPathSelectorCreateTime:
		return (*timestamppb.Timestamp)(nil)
	case Memo_FieldPathSelectorUpdateTime:
		return (*timestamppb.Timestamp)(nil)
	case Memo_FieldPathSelectorMessage:
		return ""
	case Memo_FieldPathSelectorCreatedBy:
		return ""
	default:
		panic(fmt.Sprintf("Invalid selector for Memo: %d", fp.selector))
	}
}

func (fp *Memo_FieldTerminalPath) ClearValue(item *Memo) {
	if item != nil {
		switch fp.selector {
		case Memo_FieldPathSelectorCreateTime:
			item.CreateTime = nil
		case Memo_FieldPathSelectorUpdateTime:
			item.UpdateTime = nil
		case Memo_FieldPathSelectorMessage:
			item.Message = ""
		case Memo_FieldPathSelectorCreatedBy:
			item.CreatedBy = ""
		default:
			panic(fmt.Sprintf("Invalid selector for Memo: %d", fp.selector))
		}
	}
}

func (fp *Memo_FieldTerminalPath) ClearValueRaw(item proto.Message) {
	fp.ClearValue(item.(*Memo))
}

// IsLeaf - whether field path is holds simple value
func (fp *Memo_FieldTerminalPath) IsLeaf() bool {
	return fp.selector == Memo_FieldPathSelectorCreateTime ||
		fp.selector == Memo_FieldPathSelectorUpdateTime ||
		fp.selector == Memo_FieldPathSelectorMessage ||
		fp.selector == Memo_FieldPathSelectorCreatedBy
}

func (fp *Memo_FieldTerminalPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	return []gotenobject.FieldPath{fp}
}

func (fp *Memo_FieldTerminalPath) WithIValue(value interface{}) Memo_FieldPathValue {
	switch fp.selector {
	case Memo_FieldPathSelectorCreateTime:
		return &Memo_FieldTerminalPathValue{Memo_FieldTerminalPath: *fp, value: value.(*timestamppb.Timestamp)}
	case Memo_FieldPathSelectorUpdateTime:
		return &Memo_FieldTerminalPathValue{Memo_FieldTerminalPath: *fp, value: value.(*timestamppb.Timestamp)}
	case Memo_FieldPathSelectorMessage:
		return &Memo_FieldTerminalPathValue{Memo_FieldTerminalPath: *fp, value: value.(string)}
	case Memo_FieldPathSelectorCreatedBy:
		return &Memo_FieldTerminalPathValue{Memo_FieldTerminalPath: *fp, value: value.(string)}
	default:
		panic(fmt.Sprintf("Invalid selector for Memo: %d", fp.selector))
	}
}

func (fp *Memo_FieldTerminalPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fp.WithIValue(value)
}

func (fp *Memo_FieldTerminalPath) WithIArrayOfValues(values interface{}) Memo_FieldPathArrayOfValues {
	fpaov := &Memo_FieldTerminalPathArrayOfValues{Memo_FieldTerminalPath: *fp}
	switch fp.selector {
	case Memo_FieldPathSelectorCreateTime:
		return &Memo_FieldTerminalPathArrayOfValues{Memo_FieldTerminalPath: *fp, values: values.([]*timestamppb.Timestamp)}
	case Memo_FieldPathSelectorUpdateTime:
		return &Memo_FieldTerminalPathArrayOfValues{Memo_FieldTerminalPath: *fp, values: values.([]*timestamppb.Timestamp)}
	case Memo_FieldPathSelectorMessage:
		return &Memo_FieldTerminalPathArrayOfValues{Memo_FieldTerminalPath: *fp, values: values.([]string)}
	case Memo_FieldPathSelectorCreatedBy:
		return &Memo_FieldTerminalPathArrayOfValues{Memo_FieldTerminalPath: *fp, values: values.([]string)}
	default:
		panic(fmt.Sprintf("Invalid selector for Memo: %d", fp.selector))
	}
	return fpaov
}

func (fp *Memo_FieldTerminalPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fp.WithIArrayOfValues(values)
}

func (fp *Memo_FieldTerminalPath) WithIArrayItemValue(value interface{}) Memo_FieldPathArrayItemValue {
	switch fp.selector {
	default:
		panic(fmt.Sprintf("Invalid selector for Memo: %d", fp.selector))
	}
}

func (fp *Memo_FieldTerminalPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fp.WithIArrayItemValue(value)
}

// Memo_FieldPathValue allows storing values for Memo fields according to their type
type Memo_FieldPathValue interface {
	Memo_FieldPath
	gotenobject.FieldPathValue
	SetTo(target **Memo)
	CompareWith(*Memo) (cmp int, comparable bool)
}

func ParseMemo_FieldPathValue(pathStr, valueStr string) (Memo_FieldPathValue, error) {
	fp, err := ParseMemo_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpv, err := gotenobject.ParseFieldPathValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Memo field path value from %s: %v", valueStr, err)
	}
	return fpv.(Memo_FieldPathValue), nil
}

func MustParseMemo_FieldPathValue(pathStr, valueStr string) Memo_FieldPathValue {
	fpv, err := ParseMemo_FieldPathValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpv
}

type Memo_FieldTerminalPathValue struct {
	Memo_FieldTerminalPath
	value interface{}
}

var _ Memo_FieldPathValue = (*Memo_FieldTerminalPathValue)(nil)

// GetRawValue returns raw value stored under selected path for 'Memo' as interface{}
func (fpv *Memo_FieldTerminalPathValue) GetRawValue() interface{} {
	return fpv.value
}
func (fpv *Memo_FieldTerminalPathValue) AsCreateTimeValue() (*timestamppb.Timestamp, bool) {
	res, ok := fpv.value.(*timestamppb.Timestamp)
	return res, ok
}
func (fpv *Memo_FieldTerminalPathValue) AsUpdateTimeValue() (*timestamppb.Timestamp, bool) {
	res, ok := fpv.value.(*timestamppb.Timestamp)
	return res, ok
}
func (fpv *Memo_FieldTerminalPathValue) AsMessageValue() (string, bool) {
	res, ok := fpv.value.(string)
	return res, ok
}
func (fpv *Memo_FieldTerminalPathValue) AsCreatedByValue() (string, bool) {
	res, ok := fpv.value.(string)
	return res, ok
}

// SetTo stores value for selected field for object Memo
func (fpv *Memo_FieldTerminalPathValue) SetTo(target **Memo) {
	if *target == nil {
		*target = new(Memo)
	}
	switch fpv.selector {
	case Memo_FieldPathSelectorCreateTime:
		(*target).CreateTime = fpv.value.(*timestamppb.Timestamp)
	case Memo_FieldPathSelectorUpdateTime:
		(*target).UpdateTime = fpv.value.(*timestamppb.Timestamp)
	case Memo_FieldPathSelectorMessage:
		(*target).Message = fpv.value.(string)
	case Memo_FieldPathSelectorCreatedBy:
		(*target).CreatedBy = fpv.value.(string)
	default:
		panic(fmt.Sprintf("Invalid selector for Memo: %d", fpv.selector))
	}
}

func (fpv *Memo_FieldTerminalPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*Memo)
	fpv.SetTo(&typedObject)
}

// CompareWith compares value in the 'Memo_FieldTerminalPathValue' with the value under path in 'Memo'.
func (fpv *Memo_FieldTerminalPathValue) CompareWith(source *Memo) (int, bool) {
	switch fpv.selector {
	case Memo_FieldPathSelectorCreateTime:
		leftValue := fpv.value.(*timestamppb.Timestamp)
		rightValue := source.GetCreateTime()
		if leftValue == nil {
			if rightValue != nil {
				return -1, true
			}
			return 0, true
		}
		if rightValue == nil {
			return 1, true
		}
		if leftValue.AsTime().Equal(rightValue.AsTime()) {
			return 0, true
		} else if leftValue.AsTime().Before(rightValue.AsTime()) {
			return -1, true
		} else {
			return 1, true
		}
	case Memo_FieldPathSelectorUpdateTime:
		leftValue := fpv.value.(*timestamppb.Timestamp)
		rightValue := source.GetUpdateTime()
		if leftValue == nil {
			if rightValue != nil {
				return -1, true
			}
			return 0, true
		}
		if rightValue == nil {
			return 1, true
		}
		if leftValue.AsTime().Equal(rightValue.AsTime()) {
			return 0, true
		} else if leftValue.AsTime().Before(rightValue.AsTime()) {
			return -1, true
		} else {
			return 1, true
		}
	case Memo_FieldPathSelectorMessage:
		leftValue := fpv.value.(string)
		rightValue := source.GetMessage()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case Memo_FieldPathSelectorCreatedBy:
		leftValue := fpv.value.(string)
		rightValue := source.GetCreatedBy()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	default:
		panic(fmt.Sprintf("Invalid selector for Memo: %d", fpv.selector))
	}
}

func (fpv *Memo_FieldTerminalPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpv.CompareWith(source.(*Memo))
}

// Memo_FieldPathArrayItemValue allows storing single item in Path-specific values for Memo according to their type
// Present only for array (repeated) types.
type Memo_FieldPathArrayItemValue interface {
	gotenobject.FieldPathArrayItemValue
	Memo_FieldPath
	ContainsValue(*Memo) bool
}

// ParseMemo_FieldPathArrayItemValue parses string and JSON-encoded value to its Value
func ParseMemo_FieldPathArrayItemValue(pathStr, valueStr string) (Memo_FieldPathArrayItemValue, error) {
	fp, err := ParseMemo_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaiv, err := gotenobject.ParseFieldPathArrayItemValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Memo field path array item value from %s: %v", valueStr, err)
	}
	return fpaiv.(Memo_FieldPathArrayItemValue), nil
}

func MustParseMemo_FieldPathArrayItemValue(pathStr, valueStr string) Memo_FieldPathArrayItemValue {
	fpaiv, err := ParseMemo_FieldPathArrayItemValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpaiv
}

type Memo_FieldTerminalPathArrayItemValue struct {
	Memo_FieldTerminalPath
	value interface{}
}

var _ Memo_FieldPathArrayItemValue = (*Memo_FieldTerminalPathArrayItemValue)(nil)

// GetRawValue returns stored element value for array in object Memo as interface{}
func (fpaiv *Memo_FieldTerminalPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaiv.value
}

func (fpaiv *Memo_FieldTerminalPathArrayItemValue) GetSingle(source *Memo) (interface{}, bool) {
	return nil, false
}

func (fpaiv *Memo_FieldTerminalPathArrayItemValue) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpaiv.GetSingle(source.(*Memo))
}

// Contains returns a boolean indicating if value that is being held is present in given 'Memo'
func (fpaiv *Memo_FieldTerminalPathArrayItemValue) ContainsValue(source *Memo) bool {
	slice := fpaiv.Memo_FieldTerminalPath.Get(source)
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

// Memo_FieldPathArrayOfValues allows storing slice of values for Memo fields according to their type
type Memo_FieldPathArrayOfValues interface {
	gotenobject.FieldPathArrayOfValues
	Memo_FieldPath
}

func ParseMemo_FieldPathArrayOfValues(pathStr, valuesStr string) (Memo_FieldPathArrayOfValues, error) {
	fp, err := ParseMemo_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaov, err := gotenobject.ParseFieldPathArrayOfValues(fp, valuesStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Memo field path array of values from %s: %v", valuesStr, err)
	}
	return fpaov.(Memo_FieldPathArrayOfValues), nil
}

func MustParseMemo_FieldPathArrayOfValues(pathStr, valuesStr string) Memo_FieldPathArrayOfValues {
	fpaov, err := ParseMemo_FieldPathArrayOfValues(pathStr, valuesStr)
	if err != nil {
		panic(err)
	}
	return fpaov
}

type Memo_FieldTerminalPathArrayOfValues struct {
	Memo_FieldTerminalPath
	values interface{}
}

var _ Memo_FieldPathArrayOfValues = (*Memo_FieldTerminalPathArrayOfValues)(nil)

func (fpaov *Memo_FieldTerminalPathArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpaov.selector {
	case Memo_FieldPathSelectorCreateTime:
		for _, v := range fpaov.values.([]*timestamppb.Timestamp) {
			values = append(values, v)
		}
	case Memo_FieldPathSelectorUpdateTime:
		for _, v := range fpaov.values.([]*timestamppb.Timestamp) {
			values = append(values, v)
		}
	case Memo_FieldPathSelectorMessage:
		for _, v := range fpaov.values.([]string) {
			values = append(values, v)
		}
	case Memo_FieldPathSelectorCreatedBy:
		for _, v := range fpaov.values.([]string) {
			values = append(values, v)
		}
	}
	return
}
func (fpaov *Memo_FieldTerminalPathArrayOfValues) AsCreateTimeArrayOfValues() ([]*timestamppb.Timestamp, bool) {
	res, ok := fpaov.values.([]*timestamppb.Timestamp)
	return res, ok
}
func (fpaov *Memo_FieldTerminalPathArrayOfValues) AsUpdateTimeArrayOfValues() ([]*timestamppb.Timestamp, bool) {
	res, ok := fpaov.values.([]*timestamppb.Timestamp)
	return res, ok
}
func (fpaov *Memo_FieldTerminalPathArrayOfValues) AsMessageArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}
func (fpaov *Memo_FieldTerminalPathArrayOfValues) AsCreatedByArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}
