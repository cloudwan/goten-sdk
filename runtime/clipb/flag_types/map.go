package flag_types

import (
	"encoding/json"
	"reflect"

	"github.com/spf13/pflag"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cloudwan/goten-sdk/runtime/strcase"
)

type MapVar struct {
	FdPath []preflect.FieldDescriptor
	DVals  []interface{}
	Fd     preflect.FieldDescriptor
	Msg    proto.Message
	Tp     reflect.Type
}

var _ pflag.Value = (*MapVar)(nil)

func (v *MapVar) String() string {
	rawMap, present := getCurrentRawValue(v.Msg, v.FdPath)
	if !present {
		return "{}"
	}
	data, err := json.Marshal(rawMap)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func (v *MapVar) Set(raw string) error {
	mapType := reflect.New(v.Tp).Interface()
	if err := json.Unmarshal([]byte(raw), mapType); err != nil {
		return err
	}
	pathDefaults, pathFds := stripPathToLast(v.DVals, v.FdPath)
	currentMsg := getFieldHolderAndEnsurePath(v.Msg, pathDefaults, pathFds)

	// A bit tricky, but since maps cannot be in oneofs, we can just simply
	// use reflection and spare effort to convert into preflect.Map
	reflect.ValueOf(currentMsg.Interface()).Elem().
		FieldByName(strcase.ToCamel(string(v.Fd.Name()))).
		Set(reflect.ValueOf(mapType).Elem())
	return nil
}

func (v *MapVar) Type() string {
	return v.Tp.String()
}
