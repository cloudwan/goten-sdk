package flag_types

import (
	"encoding/json"
	"reflect"

	"github.com/spf13/pflag"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cloudwan/goten-sdk/runtime/object"
)

type SliceVar struct {
	Fd  preflect.FieldDescriptor
	Msg proto.Message
	Fp  object.FieldPath
	Tp  reflect.Type
}

var _ pflag.Value = (*SliceVar)(nil)

func (v *SliceVar) String() string {
	currentSlice := v.Fp.GetRaw(v.Msg)
	data, err := json.Marshal(currentSlice)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func (v *SliceVar) Set(raw string) error {
	current, _ := v.Fp.GetSingleRaw(v.Msg)
	elemType := reflect.TypeOf(current).Elem()
	val, err := makeProtoValue(raw, v.Fd, elemType)
	if err != nil {
		return err
	}

	var newSliceElement reflect.Value
	switch v.Fd.Kind() {
	case preflect.EnumKind:
		enumType := reflect.New(elemType).Interface().(preflect.Enum).Type()
		newSliceElement = reflect.ValueOf(enumType.New(val.Enum()))
	case preflect.MessageKind, preflect.GroupKind:
		newSliceElement = reflect.ValueOf(val.Message().Interface())
	default:
		newSliceElement = reflect.ValueOf(val.Interface())
	}
	newSlice := reflect.Append(reflect.ValueOf(current), newSliceElement)
	v.Fp.WithRawIValue(newSlice.Interface()).SetToRaw(v.Msg)
	return nil
}

func (v *SliceVar) Type() string {
	return "slice"
}
