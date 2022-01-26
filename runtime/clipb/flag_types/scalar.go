package flag_types

import (
	"fmt"
	"reflect"

	"github.com/spf13/pflag"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cloudwan/goten-sdk/runtime/object"
)

type ScalarVar struct {
	Fp  object.FieldPath
	Fd  preflect.FieldDescriptor
	Msg proto.Message
	Tp  reflect.Type
}

var _ pflag.Value = (*ScalarVar)(nil)

func (v *ScalarVar) String() string {
	rawValue, present := v.Fp.GetSingleRaw(v.Msg)
	if !present {
		return fmt.Sprint(reflect.New(v.Tp))
	}
	if stringer, ok := rawValue.(fmt.Stringer); ok {
		return stringer.String()
	}
	return fmt.Sprint(rawValue)
}

func (v *ScalarVar) Set(raw string) error {
	val, err := makeProtoValue(raw, v.Fd, v.Tp)
	if err != nil {
		return err
	}

	switch v.Fd.Kind() {
	case preflect.EnumKind:
		enumType := reflect.New(v.Tp).Interface().(preflect.Enum).Type()
		v.Fp.WithRawIValue(enumType.New(val.Enum())).SetToRaw(v.Msg)
	case preflect.MessageKind, preflect.GroupKind:
		v.Fp.WithRawIValue(val.Message().Interface()).SetToRaw(v.Msg)
	default:
		v.Fp.WithRawIValue(val.Interface()).SetToRaw(v.Msg)
	}
	return nil
}

func (v *ScalarVar) Type() string {
	switch v.Fd.Kind() {
	case preflect.GroupKind, preflect.MessageKind:
		return string(v.Fd.Message().Name())
	case preflect.EnumKind:
		return string(v.Fd.Enum().Name())
	default:
		return v.Fd.Kind().String()
	}
}
