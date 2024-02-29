package utils

import (
	"fmt"
	"reflect"

	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"
)

func GetFieldTypeForOneOf(protoMsg proto.Message, fieldDesc preflect.FieldDescriptor) (reflect.StructField, bool) {
	if fieldDesc.ContainingOneof() == nil {
		panic(fmt.Sprintf("getFieldTypeForOneOf should be called only for oneof fields, but got %s/%s",
			fieldDesc.Parent().Name(), fieldDesc.Name()))
	}
	// if field is part of oneof, then unfortunately we cannot use reflection on
	// message - field type is interface, but we need struct that implements that
	// interface...
	// TODO: This implementation relies on presence of field OneofWrappers
	// within preflect.MessageType interface (!!!) - it holds GO types of our oneofs.
	// We can iterate them and find type we actually need. But as you may notice, we rely
	// on implementation within protobuf library, which is a bit far from perfect...
	// It would be nice if  protoMsg.ProtoReflect().NewField(fieldDesc) would return us
	// something better (which may be problem for ProtoStringer types especially).
	oneOfTypeForThisField := strcase.ToCamel(string(fieldDesc.Name()))
	oneOfTypeForThisFieldAlternative := oneOfTypeForThisField + "_"
	parentDesc := fieldDesc.Parent()
	for parentDesc != nil {
		if _, isMsg := parentDesc.(preflect.MessageDescriptor); isMsg {
			oneOfTypeForThisField = fmt.Sprintf("%s_%s", parentDesc.Name(), oneOfTypeForThisField)
			oneOfTypeForThisFieldAlternative = oneOfTypeForThisField + "_"
		} else {
			break
		}
		parentDesc = parentDesc.Parent()
	}

	hiddenMsgInfo := reflect.ValueOf(protoMsg.ProtoReflect().Type().(interface{}))
	for _, oneOfWrapper := range hiddenMsgInfo.Elem().FieldByName("OneofWrappers").Interface().([]interface{}) {
		tf := reflect.TypeOf(oneOfWrapper).Elem()
		if tf.Name() == oneOfTypeForThisField || tf.Name() == oneOfTypeForThisFieldAlternative {
			fieldType, ok := tf.FieldByName(strcase.ToCamel(string(fieldDesc.Name())))
			if !ok {
				fieldType, _ = tf.FieldByName(strcase.ToCamel(string(fieldDesc.Name())) + "_")
			}
			return fieldType, true
		}
	}
	return reflect.StructField{}, false
}
