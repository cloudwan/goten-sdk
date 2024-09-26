package object

import (
	"encoding/base64"
	"encoding/json"
	"reflect"
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func ParseFieldPathValue(fp FieldPath, valueStr string) (FieldPathValue, error) {
	value, err := parseValueForType(fp.GetDefault(), valueStr)
	if err != nil {
		return nil, err
	}
	return fp.WithRawIValue(value), nil
}

func ParseFieldPathArrayItemValue(fp FieldPath, valueStr string) (FieldPathArrayItemValue, error) {
	itemType := reflect.TypeOf(fp.GetDefault()).Elem()
	value, err := parseValueForType(reflect.New(itemType).Elem().Interface(), valueStr)
	if err != nil {
		return nil, err
	}
	return fp.WithRawIArrayItemValue(value), nil
}

func ParseFieldPathArrayOfValues(fp FieldPath, valuesStr string) (FieldPathArrayOfValues, error) {
	// Not most efficient: we convert into []interface{}, then each element into string and back
	// to proper value. On the other hand, reuses code and parsing small arrays of strings, numbers,
	// enums should not take much time anyway
	// But optionally code can be changed to take into account proper underlying type

	tmpSlice := make([]interface{}, 0)
	err := json.Unmarshal([]byte(valuesStr), &tmpSlice)
	if err != nil {
		return nil, err
	}

	sliceType := reflect.SliceOf(reflect.TypeOf(fp.GetDefault()))
	dstSlice := reflect.MakeSlice(sliceType, 0, len(tmpSlice))
	for _, sliceElement := range tmpSlice {
		strValueOfElem, err := json.Marshal(sliceElement)
		if err != nil {
			return nil, err
		}
		elemValue, err := parseValueForType(fp.GetDefault(), string(strValueOfElem))
		if err != nil {
			return nil, err
		}
		dstSlice = reflect.Append(dstSlice, reflect.ValueOf(elemValue))
	}
	return fp.WithRawIArrayOfValues(dstSlice.Interface()), nil
}

func parseValueForType(defaultValue interface{}, valueStr string) (interface{}, error) {
	// special case for byte[] array
	if _, isBytesArray := defaultValue.([]byte); isBytesArray {
		if len(valueStr) > 0 && (valueStr[0] == '"' || valueStr[0] == '\'') {
			valueStr = valueStr[1:]
		}
		if len(valueStr) > 0 && (valueStr[len(valueStr)-1] == '"' || valueStr[len(valueStr)-1] == '\'') {
			valueStr = valueStr[:len(valueStr)-1]
		}
		bytesArray, err := base64.StdEncoding.DecodeString(valueStr)
		if err != nil {
			return nil, err
		}
		return bytesArray, nil
	}
	valueType := reflect.TypeOf(defaultValue)

	// Elements in containers (maps, slices) are often different compared to their go types.
	// resource.Name is mapped into string (but go type is struct), timestamps and durations
	// are string in json, but proto.Message in Go, Field masks also are strings in JSON
	// representation. For those containers, we need to use map[$KEY]interface{} or []interface{}
	// intermediate objects. Then convert each item back into string and then to proper value.
	if valueType.Kind() == reflect.Map {
		mapType := reflect.MapOf(valueType.Key(), reflect.TypeOf(map[string]interface{}{}).Elem())
		tmpMapPtr := reflect.New(mapType)
		err := json.Unmarshal([]byte(valueStr), tmpMapPtr.Interface())
		if err != nil {
			return nil, err
		}

		tmpMap := tmpMapPtr.Elem()
		mapKeys := tmpMap.MapKeys()
		elementDefaultType := reflect.New(valueType.Elem()).Elem().Interface()
		dstMap := reflect.MakeMapWithSize(valueType, len(mapKeys))
		for _, mapKey := range mapKeys {
			value := tmpMap.MapIndex(mapKey).Interface()
			valueStrOfElem, err := json.Marshal(value)
			if err != nil {
				return nil, err
			}
			elemValue, err := newSimpleValue(elementDefaultType, string(valueStrOfElem))
			if err != nil {
				return nil, err
			}
			dstMap.SetMapIndex(mapKey, reflect.ValueOf(elemValue))
		}
		return dstMap.Interface(), nil
	} else if valueType.Kind() == reflect.Slice {
		tmpSlice := make([]interface{}, 0)
		err := json.Unmarshal([]byte(valueStr), &tmpSlice)
		if err != nil {
			return nil, err
		}

		dstSlice := reflect.MakeSlice(valueType, 0, len(tmpSlice))
		elementDefaultType := reflect.New(valueType.Elem()).Elem().Interface()
		for _, sliceElement := range tmpSlice {
			strValueOfElem, err := json.Marshal(sliceElement)
			if err != nil {
				return nil, err
			}
			elemValue, err := newSimpleValue(elementDefaultType, string(strValueOfElem))
			if err != nil {
				return nil, err
			}
			dstSlice = reflect.Append(dstSlice, reflect.ValueOf(elemValue))
		}
		return dstSlice.Interface(), nil
	} else {
		return newSimpleValue(defaultValue, valueStr)
	}
}

func newSimpleValue(defaultType interface{}, valueStr string) (interface{}, error) {
	switch typedDefaultType := defaultType.(type) {
	case preflect.Enum:
		enumType := typedDefaultType.Type()
		enum, err := newEnumFromValue(valueStr, enumType)
		if err != nil {
			return nil, err
		}
		return enum, nil
	case FieldMask:
		protoMaskValuePtr := new(fieldmaskpb.FieldMask)
		err := protojson.Unmarshal([]byte(valueStr), protoMaskValuePtr)
		if err != nil {
			return nil, err
		}
		underlyingValue := reflect.New(reflect.TypeOf(defaultType).Elem()).Interface().(FieldMask)
		if err := underlyingValue.FromProtoFieldMask(protoMaskValuePtr); err != nil {
			return nil, err
		}
		return underlyingValue, nil
	case proto.Message:
		underlyingMsg := typedDefaultType.ProtoReflect().New().Interface()
		err := protojson.Unmarshal([]byte(valueStr), underlyingMsg)
		if err != nil {
			return nil, err
		}
		return underlyingMsg, nil
	case preflect.ProtoStringer:
		protoStringValuePtr := new(string)
		err := json.Unmarshal([]byte(valueStr), protoStringValuePtr)
		if err != nil {
			return nil, err
		}
		underlyingValue := reflect.New(reflect.TypeOf(defaultType).Elem()).Interface().(preflect.ProtoStringer)
		if err := underlyingValue.ParseProtoString(*protoStringValuePtr); err != nil {
			return nil, err
		}
		return underlyingValue, nil
	default:
		// booleans, integers, strings... could be pointer to them (proto2, optionals)
		defValType := reflect.TypeOf(defaultType)
		isPtr := defValType.Kind() == reflect.Ptr
		if isPtr {
			defValType = defValType.Elem()
		}
		underlyingValue := reflect.New(defValType)
		err := json.Unmarshal([]byte(valueStr), underlyingValue.Interface())
		if err != nil {
			return nil, err
		}
		if isPtr {
			return underlyingValue.Interface(), nil
		} else {
			return underlyingValue.Elem().Interface(), nil
		}
	}
}

func newEnumFromValue(valueStr string, enumType preflect.EnumType) (preflect.Enum, error) {
	values := enumType.Descriptor().Values()
	var valueDescriptor preflect.EnumValueDescriptor
	if strings.HasPrefix(valueStr, "\"") {
		valueStr = strings.TrimSuffix(strings.TrimPrefix(valueStr, "\""), "\"")
		valueDescriptor = values.ByName(preflect.Name(valueStr))
	} else {
		num, err := strconv.Atoi(valueStr)
		if err != nil {
			return nil, err
		}
		valueDescriptor = values.ByNumber(preflect.EnumNumber(num))
	}
	if valueDescriptor != nil {
		return enumType.New(valueDescriptor.Number()), nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "Enum Value %s was not found in %s",
		valueStr, enumType.Descriptor().Name())
}
