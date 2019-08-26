package yext

import (
	"encoding/json"
	"reflect"
	"strings"
)

func customUnmarshal(i interface{}, m map[string]interface{}) interface{} {
	var (
		val         = Indirect(reflect.ValueOf(i))
		unknownKeys = map[string]interface{}{}
	)
	if val.Kind() == reflect.Struct {
		var jsonTagToKey = map[string]string{}
		for i := 0; i < val.Type().NumField(); i++ {
			field := val.Type().Field(i)
			tag := strings.Replace(field.Tag.Get("json"), ",omitempty", "", -1)
			jsonTagToKey[tag] = field.Name
		}
		//log.Println("tags", jsonTagToKey)

		for tag, val := range m {
			if strings.HasPrefix(tag, "c_") || strings.HasPrefix(tag, "cf_") {
				//log.Println("looking for tag", tag)
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
						r := customUnmarshal(v.Interface(), vMap)
						Indirect(reflect.ValueOf(i)).FieldByName(jsonTagToKey[tag]).Set(reflect.ValueOf(r))
					}
				} else {
					//log.Println("did not find tag", tag)
					unknownKeys[tag] = val
				}
			}
			if Indirect(reflect.ValueOf(i)).FieldByName("UnknownFields").CanSet() {
				Indirect(reflect.ValueOf(i)).FieldByName("UnknownFields").Set(reflect.ValueOf(&unknownKeys))
			}
		}
	}
	return i
}

func UnmarshalCustomEntityJSON(i interface{}, data []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	i = customUnmarshal(i, m)
	return nil
}

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

func HydrateUnknownFields(entity Entity, customFields []*CustomField) (Entity, error) {
	if entity == nil {
		return nil, nil
	}

	v := Indirect(reflect.ValueOf(entity)).FieldByName("UnknownFields")
	if v.IsValid() {
		var (
			hydratedFields = map[string]interface{}{}
			ptr            = v.Interface().(*map[string]interface{})
		)
		if ptr != nil {
			unknownFields := *ptr
			for fieldKey, fieldValue := range unknownFields {
				hydratedFields[fieldKey] = hydrateField(fieldKey, fieldValue, customFields)
			}
			Indirect(reflect.ValueOf(entity)).FieldByName("UnknownFields").Set(reflect.ValueOf(&hydratedFields))
		}
	}
	return entity, nil
}

func hydrateField(fieldKey string, fieldVal interface{}, customFields []*CustomField) interface{} {
	for _, customField := range customFields {
		if customField.GetId() == fieldKey {
			switch customField.Type {
			case CUSTOMFIELDTYPE_YESNO:
				return NullableBool(fieldVal.(bool))
			}
		}
	}
	return fieldVal
}
