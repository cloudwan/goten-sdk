package clipb

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/pflag"
	preflect "google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cloudwan/goten-sdk/runtime/clipb/flag_types"
	"github.com/cloudwan/goten-sdk/runtime/goten"
	"github.com/cloudwan/goten-sdk/runtime/object"
)

func MakeVarForField(msg goten.GotenMessage, fieldPath object.FieldPath) pflag.Value {
	descriptor := msg.ProtoReflect().Descriptor()
	var fd preflect.FieldDescriptor
	for _, item := range strings.Split(fieldPath.String(), ".") {
		if fd == nil {
			fd = descriptor.Fields().ByName(preflect.Name(item))
		} else {
			fd = fd.Message().Fields().ByName(preflect.Name(item))
		}
	}
	if fd == nil {
		panic(fmt.Sprintf("Field path %s not found in message %s", fieldPath, msg.ProtoReflect().Descriptor().Name()))
	}

	reflectType := reflect.TypeOf(fieldPath.GetDefault())

	if fd.IsMap() {
		return &flag_types.MapVar{Fd: fd, Msg: msg, Fp: fieldPath, Tp: reflectType}
	}
	if fd.IsList() {
		return &flag_types.SliceVar{Fd: fd, Msg: msg, Fp: fieldPath, Tp: reflectType}
	}
	if reflectType.Implements(reflect.TypeOf((*flag_types.CustomTypeCliValue)(nil)).Elem()) {
		return &flag_types.CustomTypeCliValueAdapter{Fd: fd, Msg: msg, Fp: fieldPath, Tp: reflectType.Elem()}
	}
	return &flag_types.ScalarVar{Fd: fd, Msg: msg, Fp: fieldPath, Tp: reflectType}
}
