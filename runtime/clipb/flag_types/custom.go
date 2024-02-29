package flag_types

import (
	"fmt"
	"reflect"

	"github.com/spf13/pflag"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"
)

type CustomTypeCliValue interface {
	fmt.Stringer
	SetFromCliFlag(raw string) error
}

type CustomTypeCliValueAdapter struct {
	FdPath []preflect.FieldDescriptor
	DVals  []interface{}
	Fd     preflect.FieldDescriptor
	Msg    proto.Message
	Tp     reflect.Type
}

var _ pflag.Value = (*CustomTypeCliValueAdapter)(nil)

func (v *CustomTypeCliValueAdapter) Set(raw string) error {
	customValue, present := getCurrentRawValue(v.Msg, v.FdPath)
	if !present {
		customValue = reflect.New(v.Tp).Interface()
	}
	if err := customValue.(CustomTypeCliValue).SetFromCliFlag(raw); err != nil {
		return err
	}
	pathDefaults, pathFds := stripPathToLast(v.DVals, v.FdPath)
	currentMsg := getFieldHolderAndEnsurePath(v.Msg, pathDefaults, pathFds)
	if asProtoMsg, ok := customValue.(proto.Message); !ok {
		currentMsg.Set(v.Fd, preflect.ValueOf(customValue))
	} else {
		currentMsg.Set(v.Fd, preflect.ValueOfMessage(asProtoMsg.ProtoReflect()))
	}
	return nil
}

func (v *CustomTypeCliValueAdapter) String() string {
	rawValue, present := getCurrentRawValue(v.Msg, v.FdPath)
	if !present {
		return reflect.New(reflect.PtrTo(v.Tp)).Elem().Interface().(CustomTypeCliValue).String()
	}
	return rawValue.(CustomTypeCliValue).String()
}

func (v *CustomTypeCliValueAdapter) Type() string {
	return v.Tp.String()
}
