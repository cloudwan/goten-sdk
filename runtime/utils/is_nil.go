package utils

import "reflect"

func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}
	switch reflect.ValueOf(v).Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func:
		return reflect.ValueOf(v).IsNil()
	default:
		return false
	}
}
