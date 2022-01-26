package flag_types

import (
	"encoding/json"
	"reflect"

	"github.com/spf13/pflag"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cloudwan/goten-sdk/runtime/object"
)

type MapVar struct {
	Fd  preflect.FieldDescriptor
	Msg proto.Message
	Fp  object.FieldPath
	Tp  reflect.Type
}

var _ pflag.Value = (*MapVar)(nil)

func (v *MapVar) String() string {
	rawMap, present := v.Fp.GetSingleRaw(v.Msg)
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
	v.Fp.WithRawIValue(reflect.ValueOf(mapType).Elem().Interface()).SetToRaw(v.Msg)
	return nil
}

func (v *MapVar) Type() string {
	return v.Tp.String()
}
