package flag_types

import (
	"fmt"
	"reflect"

	"github.com/spf13/pflag"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cloudwan/goten-sdk/runtime/object"
)

type CustomTypeCliValue interface {
	fmt.Stringer
	SetFromCliFlag(raw string) error
}

type CustomTypeCliValueAdapter struct {
	Fd  preflect.FieldDescriptor
	Msg proto.Message
	Fp  object.FieldPath
	Tp  reflect.Type
}

var _ pflag.Value = (*CustomTypeCliValueAdapter)(nil)

func (v *CustomTypeCliValueAdapter) Set(raw string) error {
	customValue, present := v.Fp.GetSingleRaw(v.Msg)
	if !present {
		customValue = reflect.New(v.Tp).Interface()
	}
	if err := customValue.(CustomTypeCliValue).SetFromCliFlag(raw); err != nil {
		return err
	}
	v.Fp.WithRawIValue(customValue).SetToRaw(v.Msg)
	return nil
}

func (v *CustomTypeCliValueAdapter) String() string {
	rawValue, present := v.Fp.GetSingleRaw(v.Msg)
	if !present {
		return reflect.New(reflect.PtrTo(v.Tp)).Elem().Interface().(CustomTypeCliValue).String()
	}
	return rawValue.(CustomTypeCliValue).String()
}

func (v *CustomTypeCliValueAdapter) Type() string {
	return v.Tp.String()
}
