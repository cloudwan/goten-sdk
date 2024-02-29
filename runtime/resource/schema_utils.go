package resource

import (
	"reflect"

	"google.golang.org/protobuf/proto"

	"github.com/cloudwan/goten-sdk/runtime/object"
)

func UnsetReferenceFromResource(res Resource, refFp object.FieldPath, refToUnset Reference) bool {
	var unsetFunc func(obj proto.Message, terminalPaths []object.FieldPath) bool
	unsetFunc = func(obj proto.Message, terminalPaths []object.FieldPath) bool {
		pathItemValue, ok := terminalPaths[0].GetSingleRaw(obj)
		if !ok {
			return false
		}
		rPathItemValue := reflect.ValueOf(pathItemValue)

		if len(terminalPaths) == 1 {
			// Array of references... or reference... TODO: Map support?
			if rPathItemValue.Kind() == reflect.Slice {
				anythingChanged := false
				newRefList := reflect.New(rPathItemValue.Type()).Elem()
				for i := 0; i < rPathItemValue.Len(); i++ {
					refItem := rPathItemValue.Index(i)
					if !refToUnset.GotenEqual(refItem.Interface()) {
						newRefList = reflect.Append(newRefList, refItem)
					} else {
						anythingChanged = true
					}
				}
				terminalPaths[0].WithRawIValue(newRefList.Interface()).SetToRaw(obj)
				return anythingChanged
			} else {
				if refToUnset.GotenEqual(pathItemValue) {
					terminalPaths[0].ClearValueRaw(obj)
					return true
				}
				return false
			}
		} else {
			if rPathItemValue.Kind() == reflect.Slice {
				anythingWasCleared := false
				for i := 0; i < rPathItemValue.Len(); i++ {
					arrayItem := rPathItemValue.Index(i).Interface().(proto.Message)
					if unsetFunc(arrayItem, terminalPaths[1:]) {
						anythingWasCleared = true
					}
				}
				return anythingWasCleared
			} else {
				return unsetFunc(pathItemValue.(proto.Message), terminalPaths[1:])
			}
		}
	}
	return unsetFunc(res, refFp.SplitIntoTerminalIPaths())
}
