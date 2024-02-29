package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"
)

func GetValueFromProtoPath(msg proto.Message, rawPath string) (interface{}, bool) {
	descriptor := msg.ProtoReflect().Descriptor()
	pathItems := strings.Split(rawPath, ".")
	lastMsgItem := msg

	for i, item := range pathItems {
		fd := descriptor.Fields().ByName(preflect.Name(strcase.ToSnake(item)))
		if fd == nil {
			panic(fmt.Sprintf("Field path %s not found in message %s", rawPath, descriptor.Name()))
		}
		fieldValue := lastMsgItem.ProtoReflect().Get(fd)
		if !fieldValue.IsValid() {
			return nil, false
		}
		if i == len(pathItems)-1 {
			if fd.IsList() || fd.IsMap() {
				refMsgType := reflect.ValueOf(lastMsgItem.ProtoReflect().Interface())
				if refMsgType.IsZero() {
					return nil, false
				}
				return refMsgType.Elem().FieldByName(strcase.ToCamel(string(fd.Name()))).Interface(), true
			}
			if fd.ContainingOneof() != nil {
				ooFieldValue := reflect.ValueOf(lastMsgItem.ProtoReflect().Interface()).Elem().FieldByName(strcase.ToCamel(string(fd.ContainingOneof().Name())))
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
			} else {
				rMsg := reflect.ValueOf(lastMsgItem.ProtoReflect().Interface())
				if !rMsg.IsValid() || rMsg.IsZero() || !rMsg.Elem().IsValid() || rMsg.Elem().IsZero() {
					return nil, false
				}
				rfieldValue := rMsg.Elem().FieldByName(strcase.ToCamel(string(fd.Name())))
				if rfieldValue.IsZero() || !rfieldValue.IsValid() {
					return nil, false
				}
				return rfieldValue.Interface(), true
			}
		} else {
			descriptor = fieldValue.Message().Descriptor()
			lastMsgItem = fieldValue.Message().Interface()
		}
	}
	return nil, false
}

func SetFieldPathValueToProtoMsg(target proto.Message, rawPath string, v any) {
	descriptor := target.ProtoReflect().Descriptor()
	pathItems := strings.Split(rawPath, ".")
	lastMsgItem := target
	for i, item := range pathItems {
		fd := descriptor.Fields().ByName(preflect.Name(strcase.ToSnake(item)))
		if fd == nil {
			panic(fmt.Sprintf("Field path %s not found in message %s", rawPath, descriptor.Name()))
		}
		if i == len(pathItems)-1 {
			SetFieldValueToProtoMsg(lastMsgItem, fd, v)
		} else {
			fieldValue := lastMsgItem.ProtoReflect().Get(fd)
			if !fieldValue.IsValid() || !fieldValue.Message().IsValid() {
				var fieldValueType reflect.Type
				if fd.ContainingOneof() == nil {
					refStructField, _ := reflect.TypeOf(lastMsgItem).Elem().FieldByName(strcase.ToCamel(item))
					fieldValueType = refStructField.Type
				} else {
					refStructField, _ := GetFieldTypeForOneOf(lastMsgItem, fd)
					fieldValueType = refStructField.Type
				}
				nextNested := reflect.New(fieldValueType.Elem()).Interface().(proto.Message)
				lastMsgItem.ProtoReflect().Set(fd, preflect.ValueOfMessage(nextNested.ProtoReflect()))
				lastMsgItem = nextNested
			} else {
				lastMsgItem = fieldValue.Message().Interface()
			}
			descriptor = lastMsgItem.ProtoReflect().Descriptor()
		}
	}
}

func SetFieldValueToProtoMsg(target proto.Message, fieldDesc preflect.FieldDescriptor, v any) {
	if fieldDesc.IsList() {
		rListValue := reflect.ValueOf(v)
		listContent := target.ProtoReflect().NewField(fieldDesc).List()
		for i := 0; i < rListValue.Len(); i++ {
			value := rListValue.Index(i).Interface()
			var preflectItemValue preflect.Value
			if asProtoMsg, ok := value.(proto.Message); ok {
				preflectItemValue = preflect.ValueOfMessage(asProtoMsg.ProtoReflect())
			} else if asEnum, ok := value.(preflect.Enum); ok {
				preflectItemValue = preflect.ValueOfEnum(asEnum.Number())
			} else {
				preflectItemValue = preflect.ValueOf(value)
			}
			listContent.Append(preflectItemValue)
		}
		target.ProtoReflect().Set(fieldDesc, preflect.ValueOfList(listContent))

	} else if fieldDesc.IsMap() {
		rMapValue := reflect.ValueOf(v)
		objReflect := target.ProtoReflect()

		mapContent := objReflect.NewField(fieldDesc).Map()
		for _, key := range rMapValue.MapKeys() {
			value := rMapValue.MapIndex(key).Interface()
			var preflectItemValue preflect.Value
			if asProtoMsg, ok := value.(proto.Message); ok {
				preflectItemValue = preflect.ValueOfMessage(asProtoMsg.ProtoReflect())
			} else if asEnum, ok := value.(preflect.Enum); ok {
				preflectItemValue = preflect.ValueOfEnum(asEnum.Number())
			} else {
				preflectItemValue = preflect.ValueOf(value)
			}
			mapContent.Set(preflect.ValueOf(key.Interface()).MapKey(), preflectItemValue)
		}
		objReflect.Set(fieldDesc, preflect.ValueOfMap(mapContent))

	} else {
		if asProtoMsg, ok := v.(proto.Message); ok {
			target.ProtoReflect().Set(fieldDesc, preflect.ValueOfMessage(asProtoMsg.ProtoReflect()))
		} else if asEnum, ok := v.(preflect.Enum); ok {
			target.ProtoReflect().Set(fieldDesc, preflect.ValueOfEnum(asEnum.Number()))
		} else {
			target.ProtoReflect().Set(fieldDesc, preflect.ValueOf(v))
		}
	}
}
