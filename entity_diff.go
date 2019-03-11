package yext

import (
	"fmt"
	"reflect"
)

func InstanceOf(val interface{}) interface{} {
	var (
		isPtr = reflect.ValueOf(val).Kind() == reflect.Ptr
		tmp   interface{}
	)

	if isPtr {
		var (
			ptr         = reflect.New(reflect.TypeOf(val).Elem()).Interface()
			numPointers = 0
		)
		for reflect.ValueOf(val).Kind() == reflect.Ptr {
			val = reflect.ValueOf(val).Elem().Interface()
			numPointers++
		}

		tmp = reflect.New(reflect.TypeOf(val)).Interface()
		if numPointers == 1 {
			return tmp
		}
		// This will only work for ** pointers, no *** pointers
		reflect.ValueOf(ptr).Elem().Set(reflect.ValueOf(tmp))
		return ptr
	}
	return reflect.New(reflect.TypeOf(val)).Interface()
}

func diff(a interface{}, b interface{}, nilIsEmptyA bool, nilIsEmptyB bool) (interface{}, bool) {
	var (
		aV, bV = reflect.ValueOf(a), reflect.ValueOf(b)
		isDiff = false
		delta  = InstanceOf(a)
	)

	if aV.Kind() == reflect.Ptr {
		aV = Indirect(aV)
	}
	if bV.Kind() == reflect.Ptr {
		if bV.IsNil() {
			return delta, isDiff
		}
		bV = Indirect(bV)
	}

	var (
		aT   = aV.Type()
		numA = aV.NumField()
	)

	for i := 0; i < numA; i++ {
		var (
			nameA = aT.Field(i).Name
			valA  = aV.Field(i)
			valB  = bV.Field(i)
		)

		if nameA == "nilIsEmpty" {
			continue
		}

		if IsNil(valB) {
			continue
		}

		if !valB.CanSet() {
			continue
		}

		aI, bI := valA.Interface(), valB.Interface()
		// Comparable does not handle the nil is empty case:
		// So if valA is nil, don't call comparable (valB checked for nil above)
		if !IsNil(valA) {
			comparableA, aOk := aI.(Comparable)
			comparableB, bOk := bI.(Comparable)
			if aOk && bOk {
				if !comparableA.Equal(comparableB) {
					isDiff = true
					Indirect(reflect.ValueOf(delta)).FieldByName(nameA).Set(valB)
				}
				continue
			}
		}

		// First, use recursion to handle a field that is a struct or a pointer to a struct
		// If Kind() == struct, this is likely an embedded struct
		if valA.Kind() == reflect.Struct {
			d, diff := diff(valA.Addr().Interface(), valB.Addr().Interface(), nilIsEmptyA, nilIsEmptyB)
			if diff {
				isDiff = true
				reflect.ValueOf(delta).Elem().FieldByName(nameA).Set(reflect.ValueOf(d).Elem())
			}
			continue
			// If it's a pointer to a struct we need to handle it in a special way:
		} else if valA.Kind() == reflect.Ptr && Indirect(valA).Kind() == reflect.Struct {
			// Handle case where new is &Address{} and base is &Address{"Line1"}
			if IsZeroValue(valB, nilIsEmptyB) && !IsZeroValue(valA, nilIsEmptyA) {
				isDiff = true
				Indirect(reflect.ValueOf(delta)).FieldByName(nameA).Set(valB)
			} else {
				d, diff := diff(valA.Interface(), valB.Interface(), nilIsEmptyA, nilIsEmptyB)
				if diff {
					isDiff = true
					Indirect(reflect.ValueOf(delta)).FieldByName(nameA).Set(reflect.ValueOf(d))
				}
			}
			continue
		}

		if IsZeroValue(valA, nilIsEmptyA) && IsZeroValue(valB, nilIsEmptyB) {
			continue
		}

		if !reflect.DeepEqual(aI, bI) {
			isDiff = true
			Indirect(reflect.ValueOf(delta)).FieldByName(nameA).Set(valB)
		}
	}
	return delta, isDiff
}

func Indirect(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}

func IsNil(v reflect.Value) bool {
	if v.Kind() == reflect.Ptr {
		return v.IsNil()
	}
	return false
}

// Diff(a, b): a is base, b is new
func Diff(a Entity, b Entity) (Entity, bool, error) {
	if a.GetEntityType() != b.GetEntityType() {
		return nil, true, fmt.Errorf("Entity Types do not match: '%s', '%s'", a.GetEntityType(), b.GetEntityType())
	}

	delta, isDiff := diff(a, b, getNilIsEmpty(a), getNilIsEmpty(b))
	if !isDiff {
		return nil, isDiff, nil
	}
	return delta.(Entity), isDiff, nil
}
