package yext

import (
	"fmt"
	"reflect"
)

type Registry map[string]interface{}

func (r Registry) Register(key string, val interface{}) {
	//From: https://github.com/reggo/reggo/blob/master/common/common.go#L169
	isPtr := reflect.ValueOf(val).Kind() == reflect.Ptr
	var newVal interface{}
	var tmp interface{}
	if isPtr {
		tmp = reflect.ValueOf(val).Elem().Interface()
	} else {
		tmp = val
	}
	newVal = reflect.New(reflect.TypeOf(tmp)).Elem().Interface()
	r[key] = newVal
}

func (r Registry) Lookup(key string) (interface{}, error) {
	val, ok := r[key]
	if !ok {
		return nil, fmt.Errorf("Unable to find key %s in registry %v", key, r)
	}
	return reflect.New(reflect.TypeOf(val)).Interface(), nil
}
