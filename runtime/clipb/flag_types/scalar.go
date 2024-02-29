package flag_types

import (
	"fmt"
	"reflect"

	"github.com/spf13/pflag"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"
)

type ScalarVar struct {
	FdPath []preflect.FieldDescriptor
	DVals  []interface{}
	Fd     preflect.FieldDescriptor
	Msg    proto.Message
	Tp     reflect.Type
}

var _ pflag.Value = (*ScalarVar)(nil)

func (v *ScalarVar) String() string {
	rawValue, present := getCurrentRawValue(v.Msg, v.FdPath)
	if !present {
		return fmt.Sprint(reflect.New(v.Tp))
	}
	return fmt.Sprintf("%v", rawValue)
}

func (v *ScalarVar) Set(raw string) error {
	val, err := makeProtoValue(raw, v.Fd, v.Tp)
	if err != nil {
		return err
	}
	pathDefaults, pathFds := stripPathToLast(v.DVals, v.FdPath)
	currentMsg := getFieldHolderAndEnsurePath(v.Msg, pathDefaults, pathFds)
	currentMsg.Set(v.Fd, val)
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
