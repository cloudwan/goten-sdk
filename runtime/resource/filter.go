package resource

import (
	"fmt"

	preflect "google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cloudwan/goten-sdk/runtime/object"
	filterParser "github.com/cloudwan/goten-sdk/runtime/resource/filter"
)

// Interface for goten filter
type Filter interface {
	fmt.Stringer
	preflect.ProtoStringer
	EvaluateRaw(Resource) bool
	GetRawCondition() FilterCondition
}

type FilterCondition interface {
	fmt.Stringer
	EvaluateRaw(Resource) bool
	SatisfiesRaw(FilterCondition) bool
	SpecifiesRawFieldPath(object.FieldPath) bool
}

type ConditionComposite interface {
	FilterCondition
	GetOperator() filterParser.CompositeOperator
	GetSubConditions() []FilterCondition
	ConditionComposite()
}

type ConditionNot interface {
	FilterCondition
	GetSubCondition() FilterCondition
	ConditionNot()
}

type FieldPathCondition interface {
	FilterCondition
	GetRawFieldPath() object.FieldPath
}

type ConditionCompare interface {
	FieldPathCondition
	GetOperator() filterParser.CompareOperator
	GetRawFieldPathValue() object.FieldPathValue
	ConditionCompare()
}

type UnknownConditionContainsType struct {
	Type ConditionContainsType
}

func NewUnknownConditionContainsType(ct ConditionContainsType) error {
	return &UnknownConditionContainsType{ct}
}

func (ucct *UnknownConditionContainsType) Error() string {
	return fmt.Sprintf("unknown Condition Contains Type: %s: %d", ucct.Type, ucct.Type)
}

type ConditionContainsType int

const (
	ConditionContainsTypeUnspecified = iota
	ConditionContainsTypeValue
	ConditionContainsTypeAny
	ConditionContainsTypeAll
)

func ConditionContainsTypeFromParser(contains *filterParser.Contains) ConditionContainsType {
	if contains.Value != nil {
		return ConditionContainsTypeValue
	} else if contains.Any != nil {
		return ConditionContainsTypeAny
	} else if contains.All != nil {
		return ConditionContainsTypeAll
	} else {
		panic("empty filter parser contains")
	}
}

func (cct ConditionContainsType) IsValue() bool {
	return cct == ConditionContainsTypeValue
}

func (cct ConditionContainsType) IsAny() bool {
	return cct == ConditionContainsTypeAny
}

func (cct ConditionContainsType) IsAll() bool {
	return cct == ConditionContainsTypeAll
}

func (cct ConditionContainsType) String() string {
	switch cct {
	case ConditionContainsTypeValue:
		return "VALUE"
	case ConditionContainsTypeAny:
		return "ANY"
	case ConditionContainsTypeAll:
		return "ALL"
	default:
		panic(fmt.Errorf("unknown condition contains type: %d", cct))
	}
}

type ConditionContains interface {
	FieldPathCondition
	ConditionContainsType() ConditionContainsType
	GetRawFieldPathItemValue() object.FieldPathArrayItemValue    // Value
	GetRawFieldPathItemValues() []object.FieldPathArrayItemValue // Any | All
}

type ConditionIn interface {
	FieldPathCondition
	GetRawFieldPathArrayOfValues() object.FieldPathArrayOfValues
	ConditionIn()
}

type ConditionNotIn interface {
	FieldPathCondition
	GetRawFieldPathArrayOfValues() object.FieldPathArrayOfValues
	ConditionNotIn()
}

type ConditionIsNull interface {
	FieldPathCondition
	NotNull() bool
	GetRawFieldPath() object.FieldPath
	ConditionIsNull()
}

type ConditionIsNaN interface {
	FieldPathCondition
	GetRawFieldPath() object.FieldPath
	ConditionIsNaN()
}
