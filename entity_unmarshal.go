package yext

import (
	"encoding/json"
	"reflect"
	"strings"
)

func recur(i interface{}, m map[string]interface{}) interface{} {
	var jsonTagToKey = map[string]string{}
	val := Indirect(reflect.ValueOf(i))
	for i := 0; i < val.Type().NumField(); i++ {
		field := val.Type().Field(i)
		tag := strings.Replace(field.Tag.Get("json"), ",omitempty", "", -1)
		jsonTagToKey[tag] = field.Name
	}

	for tag, val := range m {
		if _, ok := jsonTagToKey[tag]; ok {
			if val == nil {
				v := Indirect(reflect.ValueOf(i)).FieldByName(jsonTagToKey[tag])

				// Check if double pointer
				if v.Type().Kind() == reflect.Ptr {
					t := v.Type()
					for t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Ptr {
						t = t.Elem()
					}
					typedNil := reflect.New(t)
					Indirect(reflect.ValueOf(i)).FieldByName(jsonTagToKey[tag]).Set(reflect.ValueOf(typedNil.Interface()))
				}
			} else if vMap, ok := val.(map[string]interface{}); ok {
				v := Indirect(reflect.ValueOf(i)).FieldByName(jsonTagToKey[tag])
				r := recur(v.Interface(), vMap)
				Indirect(reflect.ValueOf(i)).FieldByName(jsonTagToKey[tag]).Set(reflect.ValueOf(r))
			}
		}
	}
	return i
}

func UnmarshalEntityJSON(i interface{}, data []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	i = recur(i, m)
	return nil
}
