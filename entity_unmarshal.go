package yext

import (
	"encoding/json"
	"reflect"
	"strings"
)

func UnmarshalEntityJSON(i interface{}, data []byte) error {
	var jsonTagToKey = map[string]string{}
	val := reflect.ValueOf(i).Elem()
	for i := 0; i < val.Type().NumField(); i++ {
		field := val.Type().Field(i)
		tag := strings.Replace(field.Tag.Get("json"), ",omitempty", "", -1)
		jsonTagToKey[tag] = field.Name
	}

	var m map[string]interface{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}

	for tag, val := range m {
		if _, ok := jsonTagToKey[tag]; ok && val == nil {
			v := reflect.ValueOf(i).Elem().FieldByName(jsonTagToKey[tag])

			// Check if double pointer
			if v.Type().Kind() == reflect.Ptr {
				t := v.Type()
				for t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Ptr {
					t = t.Elem()
				}
				typedNil := reflect.New(t)
				reflect.ValueOf(i).Elem().FieldByName(jsonTagToKey[tag]).Set(reflect.ValueOf(typedNil.Interface()))
			}
		}
	}
	return nil
}
