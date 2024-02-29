package flag_types

import (
	"reflect"

	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"
)

func stripPathToLast(defaults []interface{}, fds []preflect.FieldDescriptor) ([]interface{}, []preflect.FieldDescriptor) {
	if len(fds) < 2 {
		return nil, nil
	}
	return defaults[:len(defaults)-1], fds[:len(fds)-1]
}

func getCurrentRawValue(msg proto.Message, fdPath []preflect.FieldDescriptor) (interface{}, bool) {
	currentMsg := msg.ProtoReflect()
	for i, fd := range fdPath {
		fieldValue := currentMsg.Get(fd)
		if !fieldValue.IsValid() {
			return nil, false
		}
		if i == len(fdPath)-1 {
			if fd.IsList() || fd.IsMap() {
				refMsgType := reflect.ValueOf(currentMsg.Interface())
				if refMsgType.IsZero() {
					return nil, false
				}
				return reflect.ValueOf(currentMsg.Interface()).Elem().
					FieldByName(strcase.ToCamel(string(fd.Name()))).Interface(), true
			}
			if fd.ContainingOneof() != nil {
				ooFieldValue := reflect.ValueOf(currentMsg.Interface()).Elem().FieldByName(strcase.ToCamel(string(fd.ContainingOneof().Name())))
				if !ooFieldValue.IsValid() {
					return nil, false
				}
				underlyingValue := reflect.ValueOf(ooFieldValue.Interface())
				rfieldValue := underlyingValue.Elem().FieldByName(strcase.ToCamel(string(fd.Name())))
				if rfieldValue.IsZero() || !rfieldValue.IsValid() {
					// for oneof, it can be field name with "_" as suffix.
					rfieldValue = underlyingValue.Elem().FieldByName(strcase.ToCamel(string(fd.Name())) + "_")
					if rfieldValue.IsZero() || !rfieldValue.IsValid() {
						return nil, false
					}
				}
				return rfieldValue.Interface(), true
			}
			rMsg := reflect.ValueOf(currentMsg.Interface())
			if !rMsg.IsValid() || rMsg.IsZero() || !rMsg.Elem().IsValid() || rMsg.Elem().IsZero() {
				return nil, false
			}
			rfieldValue := rMsg.Elem().FieldByName(strcase.ToCamel(string(fd.Name())))
			if rfieldValue.IsZero() || !rfieldValue.IsValid() {
				return nil, false
			}
			return rfieldValue.Interface(), true
		}
		currentMsg = fieldValue.Message()
	}
	return nil, false
}

func getFieldHolderAndEnsurePath(msg proto.Message, defaults []interface{}, fds []preflect.FieldDescriptor) preflect.Message {
	currentMsg := msg.ProtoReflect()
	for i := 0; i < len(fds); i++ {
		fieldValue := currentMsg.Get(fds[i])
		if !fieldValue.IsValid() || !fieldValue.Message().IsValid() {
			msgValue := reflect.New(reflect.TypeOf(defaults[i]).Elem()).Interface().(proto.Message)
			protoMsgValue := preflect.ValueOfMessage(msgValue.ProtoReflect())
			currentMsg.Set(fds[i], protoMsgValue)
			currentMsg = protoMsgValue.Message()
		} else {
			currentMsg = fieldValue.Message()
		}
	}
	return currentMsg
}
