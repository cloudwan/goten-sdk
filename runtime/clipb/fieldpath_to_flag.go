package clipb

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/spf13/pflag"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cloudwan/goten-sdk/runtime/clipb/flag_types"
	"github.com/cloudwan/goten-sdk/runtime/goten"
	"github.com/cloudwan/goten-sdk/runtime/utils"
)

func MakeVarForField(msg goten.GotenMessage, rawFieldPath string) pflag.Value {
	descriptor := msg.ProtoReflect().Descriptor()
	fds := make([]preflect.FieldDescriptor, 0)
	defaultValues := make([]interface{}, 0)
	currentMsg := msg.(proto.Message)
	var lastRType reflect.Type
	var lastFd preflect.FieldDescriptor
	for _, item := range strings.Split(rawFieldPath, ".") {
		if len(fds) == 0 {
			lastFd = descriptor.Fields().ByName(preflect.Name(item))
		} else {
			lastFd = lastFd.Message().Fields().ByName(preflect.Name(item))
		}
		if lastFd == nil {
			panic(fmt.Sprintf("Field path %s not found in message %s", rawFieldPath, msg.ProtoReflect().Descriptor().Name()))
		}
		fds = append(fds, lastFd)
		if lastFd.ContainingOneof() == nil {
			refStructField, _ := reflect.TypeOf(currentMsg).Elem().FieldByName(strcase.ToCamel(item))
			lastRType = refStructField.Type
		} else {
			refStructField, _ := utils.GetFieldTypeForOneOf(currentMsg, lastFd)
			lastRType = refStructField.Type
		}
		if lastRType.Kind() == reflect.Pointer {
			lastRType = lastRType.Elem()
		}
		// Msg -> *Msg, []Slice -> *[]Slice, Map[] -> *Map[]...
		// consider slices and maps... Though in  SliceVar and MapVar we just do extra Elem()
		defValue := reflect.New(lastRType).Interface()
		if asProtoMsg, ok := defValue.(proto.Message); ok {
			currentMsg = asProtoMsg
		}
		defaultValues = append(defaultValues, defValue)
	}

	if lastFd.IsMap() {
		return &flag_types.MapVar{Fd: lastFd, Msg: msg, FdPath: fds, DVals: defaultValues, Tp: lastRType}
	}
	if lastFd.IsList() {
		return &flag_types.SliceVar{Fd: lastFd, Msg: msg, FdPath: fds, DVals: defaultValues, Tp: lastRType}
	}
	if reflect.New(lastRType).Type().Implements(reflect.TypeOf((*flag_types.CustomTypeCliValue)(nil)).Elem()) {
		return &flag_types.CustomTypeCliValueAdapter{Fd: lastFd, Msg: msg, FdPath: fds, DVals: defaultValues, Tp: lastRType}
	}
	return &flag_types.ScalarVar{Fd: lastFd, Msg: msg, FdPath: fds, DVals: defaultValues, Tp: lastRType}
}
