package flag_types

import (
	"encoding/json"
	"reflect"

	"github.com/iancoleman/strcase"
	"github.com/spf13/pflag"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"
)

type SliceVar struct {
	FdPath []preflect.FieldDescriptor
	DVals  []interface{}
	Fd     preflect.FieldDescriptor
	Msg    proto.Message
	Tp     reflect.Type
}

var _ pflag.Value = (*SliceVar)(nil)

func (v *SliceVar) String() string {
	currentSlice, _ := getCurrentRawValue(v.Msg, v.FdPath)
	data, err := json.Marshal(currentSlice)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func (v *SliceVar) Set(raw string) error {
	current, isCurrentSet := getCurrentRawValue(v.Msg, v.FdPath)
	if !isCurrentSet {
		current = reflect.New(reflect.TypeOf(v.DVals[len(v.DVals)-1]).Elem()).Elem().Interface()
	}
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

	pathDefaults, pathFds := stripPathToLast(v.DVals, v.FdPath)
	currentMsg := getFieldHolderAndEnsurePath(v.Msg, pathDefaults, pathFds)

	// A bit tricky, but since slices cannot be in oneofs, we can just simply
	// use reflection and spare effort to convert into preflect.List
	reflect.ValueOf(currentMsg.Interface()).Elem().
		FieldByName(strcase.ToCamel(string(v.Fd.Name()))).
		Set(newSlice)
	return nil
}

func (v *SliceVar) Type() string {
	return "slice"
}
