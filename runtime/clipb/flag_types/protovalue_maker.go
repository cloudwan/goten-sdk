package flag_types

import (
	"encoding/base64"
	"reflect"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"
)

func makeProtoValue(raw string, fd preflect.FieldDescriptor, tp reflect.Type) (preflect.Value, error) {
	if tp.Implements(reflect.TypeOf((*CustomTypeCliValue)(nil)).Elem()) {
		custom := reflect.New(tp.Elem()).Interface().(CustomTypeCliValue)
		if err := custom.SetFromCliFlag(raw); err != nil {
			return preflect.Value{}, err
		}
		if asProtoStringer, ok := custom.(preflect.ProtoStringer); ok {
			return preflect.ValueOfProtoString(asProtoStringer), nil
		}
		return preflect.ValueOfMessage(custom.(proto.Message).ProtoReflect()), nil
	}

	switch fd.Kind() {
	case preflect.Int32Kind, preflect.Sint32Kind, preflect.Sfixed32Kind:
		tmp, err := strconv.ParseInt(raw, 10, 32)
		if err != nil {
			return preflect.Value{}, err
		}
		return preflect.ValueOfInt32(int32(tmp)), nil
	case preflect.Int64Kind, preflect.Sint64Kind, preflect.Sfixed64Kind:
		tmp, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			return preflect.Value{}, err
		}
		return preflect.ValueOfInt64(tmp), nil
	case preflect.Uint32Kind, preflect.Fixed32Kind:
		tmp, err := strconv.ParseUint(raw, 10, 32)
		if err != nil {
			return preflect.Value{}, err
		}
		return preflect.ValueOfUint32(uint32(tmp)), nil
	case preflect.Uint64Kind, preflect.Fixed64Kind:
		tmp, err := strconv.ParseUint(raw, 10, 64)
		if err != nil {
			return preflect.Value{}, err
		}
		return preflect.ValueOfUint64(tmp), nil
	case preflect.FloatKind:
		tmp, err := strconv.ParseFloat(raw, 32)
		if err != nil {
			return preflect.Value{}, err
		}
		return preflect.ValueOfFloat32(float32(tmp)), nil
	case preflect.DoubleKind:
		tmp, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			return preflect.Value{}, err
		}
		return preflect.ValueOfFloat64(tmp), nil
	case preflect.BoolKind:
		tmp, err := strconv.ParseBool(raw)
		if err != nil {
			return preflect.Value{}, err
		}
		return preflect.ValueOfBool(tmp), nil
	case preflect.StringKind:
		return preflect.ValueOfString(raw), nil
	case preflect.BytesKind:
		bytes, err := base64.StdEncoding.DecodeString(raw)
		if err != nil {
			return preflect.Value{}, err
		}
		return preflect.ValueOfBytes(bytes), nil
	case preflect.EnumKind:
		enValDes := fd.Enum().Values().ByName(preflect.Name(raw))
		if enValDes == nil {
			return preflect.Value{}, status.Errorf(
				codes.InvalidArgument, "Enum value %s not found in %s", raw, fd.Enum().Name())
		}
		return preflect.ValueOfEnum(enValDes.Number()), nil
	case preflect.MessageKind, preflect.GroupKind:
		subMsg := reflect.New(tp.Elem()).Interface().(proto.Message)
		if err := protojson.Unmarshal([]byte(raw), subMsg); err != nil {
			return preflect.Value{}, err
		}
		return preflect.ValueOfMessage(subMsg.ProtoReflect()), nil
	}

	return preflect.Value{}, status.Errorf(
		codes.InvalidArgument, "Unrecognized kind %s for raw value %s", fd.Name(), raw)
}
