package yext

import (
	"encoding/json"
	"log"
	"reflect"
	"strings"
)

func unmarshal(i interface{}, m map[string]interface{}) interface{} {
	val := Indirect(reflect.ValueOf(i))
	if val.Kind() == reflect.Struct {
		var jsonTagToKey = map[string]string{}
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
					t := v.Type()
					if t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Ptr {
						for t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Ptr {
							t = t.Elem()
						}
						typedNil := reflect.New(t)

						defer func() {
							if r := recover(); r != nil {
								log.Fatalf("Error while unmarshaling field '%s': %s", jsonTagToKey[tag], r)
							}
						}()

						Indirect(reflect.ValueOf(i)).FieldByName(jsonTagToKey[tag]).Set(reflect.ValueOf(typedNil.Interface()))
					}
				} else if vMap, ok := val.(map[string]interface{}); ok {
					v := Indirect(reflect.ValueOf(i)).FieldByName(jsonTagToKey[tag])
					r := unmarshal(v.Interface(), vMap)
					Indirect(reflect.ValueOf(i)).FieldByName(jsonTagToKey[tag]).Set(reflect.ValueOf(r))
				}
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
	i = unmarshal(i, m)
	return nil
}
