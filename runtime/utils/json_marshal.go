package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func JsonMarshal(v interface{}) ([]byte, error) {
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Slice {
		items := make([]string, 0, vv.Len())
		for i := 0; i < vv.Len(); i++ {
			itemBytes, err := JsonMarshal(vv.Index(i).Interface())
			if err != nil {
				return nil, err
			}
			items = append(items, string(itemBytes))
		}
		data := fmt.Sprintf("[%s]", strings.Join(items, ","))
		return []byte(data), nil
	} else if vv.Kind() == reflect.Map {
		items := make([]string, 0, vv.Len())
		for _, mapKey := range vv.MapKeys() {
			mapValue := vv.MapIndex(mapKey)
			jsonKey, err := JsonMarshal(mapKey.Interface())
			if err != nil {
				return nil, err
			}
			jsonVal, err := JsonMarshal(mapValue.Interface())
			if err != nil {
				return nil, err
			}
			items = append(items, fmt.Sprintf("%s:%s", string(jsonKey), string(jsonVal)))
		}
		data := fmt.Sprintf("{%s}", strings.Join(items, ","))
		return []byte(data), nil
	} else {
		if m, ok := v.(proto.Message); ok {
			return protojson.Marshal(m)
		} else if s, ok := v.(fmt.Stringer); ok {
			return json.Marshal(s.String())
		} else {
			return json.Marshal(v)
		}
	}
}
