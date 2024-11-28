package utils

import (
	"fmt"
	"reflect"
	"strings"

	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cloudwan/goten-sdk/runtime/strcase"
)

// ToProtoCase Converts name to proto convention (replacing capital
// letters with _ and lowercase letter).
// No special treatment for digits.
func ToProtoCase(name string) string {
	var b []byte
	for i := 0; i < len(name); i++ {
		c := name[i]
		if 'A' <= c && c <= 'Z' {
			b = append(b, '_')
			c += 'a' - 'A' // convert to lowercase
		}
		b = append(b, c)
	}
	return string(b)
}

// TODO: Consider removing GetValueFromProtoPath

func GetValueFromProtoPath(msg proto.Message, rawPath string) (interface{}, bool) {
	descriptor := msg.ProtoReflect().Descriptor()
	pathItems := strings.Split(rawPath, ".")
	lastMsgItem := msg

	for i, item := range pathItems {
		fd := descriptor.Fields().ByName(preflect.Name(ToProtoCase(item)))
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
				if !rfieldValue.IsValid() || rfieldValue.IsZero() {
					// for oneof, it can be field name with "_" as suffix.
					rfieldValue = underlyingValue.Elem().FieldByName(strcase.ToCamel(string(fd.Name())) + "_")
					if !rfieldValue.IsValid() || rfieldValue.IsZero() {
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

func GetValuesFromProtoPath(msg proto.Message, rawPath string) ([]interface{}, bool, error) {
	origDescriptor := msg.ProtoReflect().Descriptor()
	pathItems := strings.Split(rawPath, ".")

	var getValues func(msg proto.Message, remainingPathItems []string) ([]interface{}, bool, error)
	getValues = func(msg proto.Message, remainingPathItems []string) ([]interface{}, bool, error) {
		if len(remainingPathItems) == 0 {
			return nil, false, fmt.Errorf("empty field path: %s", rawPath)
		}
		descriptor := msg.ProtoReflect().Descriptor()
		pathItem := remainingPathItems[0]
		fd := descriptor.Fields().ByName(preflect.Name(ToProtoCase(pathItem)))
		if fd == nil {
			return nil, false, fmt.Errorf("field path %s not found in message %s (item %s)", rawPath, origDescriptor.Name(), pathItem)
		}
		fieldValue := msg.ProtoReflect().Get(fd)
		if !fieldValue.IsValid() {
			return nil, false, nil
		}
		remainingPathItems = remainingPathItems[1:]

		// Handle the last item
		if len(remainingPathItems) == 0 || (fd.IsMap() && len(remainingPathItems) == 1) {
			if fd.IsMap() {
				refMsgType := reflect.ValueOf(msg.ProtoReflect().Interface())
				if refMsgType.IsZero() {
					return nil, false, nil
				}
				mapObject := refMsgType.Elem().FieldByName(strcase.ToCamel(string(fd.Name())))
				if len(remainingPathItems) == 0 {
					return []interface{}{mapObject.Interface()}, true, nil
				}
				if mapObject.Type().Key().Kind() != reflect.String {
					return nil, false, fmt.Errorf("only string-type map keys are supported")
				}
				mapKey := reflect.ValueOf(remainingPathItems[0])
				mapItem := mapObject.MapIndex(mapKey)
				if !mapItem.IsValid() || mapItem.IsZero() {
					return nil, false, nil
				}
				return []interface{}{mapItem.Interface()}, true, nil
			}
			if fd.IsList() {
				refMsgType := reflect.ValueOf(msg.ProtoReflect().Interface())
				if refMsgType.IsZero() {
					return nil, false, nil
				}
				listValue := refMsgType.Elem().FieldByName(strcase.ToCamel(string(fd.Name())))
				rValue := make([]interface{}, 0, listValue.Len())
				for i := 0; i < listValue.Len(); i++ {
					rValue = append(rValue, listValue.Index(i).Interface())
				}
				return rValue, true, nil
			}
			if fd.ContainingOneof() != nil {
				ooFieldValue := reflect.ValueOf(msg.ProtoReflect().Interface()).Elem().FieldByName(strcase.ToCamel(string(fd.ContainingOneof().Name())))
				if !ooFieldValue.IsValid() {
					return nil, false, nil
				}
				underlyingValue := reflect.ValueOf(ooFieldValue.Interface())
				rfieldValue := underlyingValue.Elem().FieldByName(strcase.ToCamel(string(fd.Name())))
				if !rfieldValue.IsValid() || rfieldValue.IsZero() {
					// for oneof, it can be field name with "_" as suffix.
					rfieldValue = underlyingValue.Elem().FieldByName(strcase.ToCamel(string(fd.Name())) + "_")
					if !rfieldValue.IsValid() || rfieldValue.IsZero() {
						return nil, false, nil
					}
				}
				return []interface{}{rfieldValue.Interface()}, true, nil
			}
			rMsg := reflect.ValueOf(msg.ProtoReflect().Interface())
			if !rMsg.IsValid() || rMsg.IsZero() || !rMsg.Elem().IsValid() || rMsg.Elem().IsZero() {
				return nil, false, nil
			}
			rfieldValue := rMsg.Elem().FieldByName(strcase.ToCamel(string(fd.Name())))
			if rfieldValue.IsZero() || !rfieldValue.IsValid() {
				return nil, false, nil
			}
			return []interface{}{rfieldValue.Interface()}, true, nil
		}

		// Still more path items to go...
		if fd.Kind() != preflect.MessageKind && fd.Kind() != preflect.GroupKind {
			return nil, false, fmt.Errorf("path %s containins non-message before the end (remaining part is %s)",
				rawPath, strings.Join(remainingPathItems, "."))
		}
		if fd.IsList() {
			listValue := fieldValue.List()

			rValue := make([]interface{}, 0, listValue.Len())
			for i := 0; i < listValue.Len(); i++ {
				listValue.Get(i).Message().Interface()
				subItems, ok, err := getValues(listValue.Get(i).Message().Interface(), remainingPathItems)
				if err != nil {
					return nil, false, err
				}
				if ok {
					rValue = append(rValue, subItems...)
				}
			}
			return rValue, true, nil
		} else if fd.IsMap() {
			if fd.MapKey().Kind() != preflect.StringKind {
				return nil, false, fmt.Errorf("only string-type map keys are supported")
			}
			mapKey := preflect.MapKey(preflect.ValueOfString(remainingPathItems[0]))
			remainingPathItems = remainingPathItems[1:]
			mapValue := fieldValue.Map().Get(mapKey)
			if !mapValue.IsValid() {
				return nil, false, nil
			}
			return getValues(mapValue.Message().Interface(), remainingPathItems)
		}
		return getValues(fieldValue.Message().Interface(), remainingPathItems)
	}
	return getValues(msg, pathItems)
}

func SetFieldPathValueToProtoMsg(target proto.Message, rawPath string, v any) {
	descriptor := target.ProtoReflect().Descriptor()
	pathItems := strings.Split(rawPath, ".")
	lastMsgItem := target
	for i, item := range pathItems {
		fd := descriptor.Fields().ByName(preflect.Name(ToProtoCase(item)))
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
