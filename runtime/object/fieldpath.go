package object

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type RawFieldPath []string

func (fp RawFieldPath) String() string {
	return strings.Join(fp, ".")
}

func ParseRawFieldPath(rawPath string) (RawFieldPath, error) {
	if rawPath == "" {
		return nil, status.Errorf(codes.InvalidArgument, "field path cannot be empty string")
	}

	fps := strings.Split(rawPath, ".")
	res := make(RawFieldPath, 0, len(fps))
	for _, fp := range fps {
		res = append(res, fp)
	}

	return res, nil
}

func (fp RawFieldPath) LastItem() string {
	if len(fp) == 0 {
		return ""
	}
	return fp[len(fp)-1]
}

type FieldPath interface {
	fmt.Stringer
	JSONString() string
	GetRaw(proto.Message) []interface{}
	GetSingleRaw(proto.Message) (interface{}, bool)
	GetDefault() interface{}
	ClearValueRaw(item proto.Message)
	WithRawIValue(value interface{}) FieldPathValue
	WithRawIArrayOfValues(values interface{}) FieldPathArrayOfValues
	WithRawIArrayItemValue(value interface{}) FieldPathArrayItemValue
	IsLeaf() bool
	SplitIntoTerminalIPaths() []FieldPath
}

type FieldPathValue interface {
	FieldPath
	GetRawValue() interface{}
	SetToRaw(target proto.Message)
	CompareWithRaw(msg proto.Message) (cmp int, comparable bool)
}

type FieldPathArrayItemValue interface {
	FieldPath
	GetRawItemValue() interface{}
}

type FieldPathArrayOfValues interface {
	FieldPath
	GetRawValues() []interface{}
}
