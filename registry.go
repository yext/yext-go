package yext

import (
	"fmt"
	"reflect"
	"strings"
)

type Registry map[string]interface{}

func (r Registry) Register(key string, val interface{}) {
	//From: https://github.com/reggo/reggo/blob/master/common/common.go#L169
	var (
		isPtr  = reflect.ValueOf(val).Kind() == reflect.Ptr
		newVal interface{}
		tmp    interface{}
	)
	if isPtr {
		tmp = reflect.ValueOf(val).Elem().Interface()
	} else {
		tmp = val
	}
	newVal = reflect.New(reflect.TypeOf(tmp)).Elem().Interface()
	r[key] = newVal
}

func (r Registry) Initialize(key string) (interface{}, error) {
	val, ok := r[key]
	if !ok {
		return nil, fmt.Errorf("Unable to find key %s in registry. Known keys: %s", key, strings.Join(r.Keys(), ","))
	}
	return reflect.New(reflect.TypeOf(val)).Interface(), nil
}

func (r Registry) Keys() []string {
	var keys = []string{}
	for key, _ := range r {
		keys = append(keys, key)
	}
	return keys
}
