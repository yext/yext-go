package yext

import (
	"log"
	"reflect"
)

func instanceOf(val interface{}) interface{} {
	var (
		isPtr = reflect.ValueOf(val).Kind() == reflect.Ptr
		tmp   interface{}
	)
	if isPtr {
		tmp = reflect.ValueOf(val).Elem().Interface()
	} else {
		tmp = val
	}
	return reflect.New(reflect.TypeOf(tmp)).Interface()
}

func diff(a interface{}, b interface{}, nilIsEmptyA bool, nilIsEmptyB bool) (interface{}, bool) {
	var (
		aV, bV = reflect.ValueOf(a), reflect.ValueOf(b)
		isDiff = false
		delta  = instanceOf(a)
	)

	if aV.Kind() == reflect.Ptr {
		aV = indirect(aV)
	}
	if bV.Kind() == reflect.Ptr {
		if bV.IsNil() {
			return delta, isDiff
		}
		bV = indirect(bV)
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
		log.Println(nameA)

		if nameA == "nilIsEmpty" {
			continue
		}

		if valB.Kind() == reflect.Ptr && valB.IsNil() {
			continue
		}

		if !valB.CanSet() {
			continue
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
		} else if valA.Kind() == reflect.Ptr && indirectKind(valA) == reflect.Struct {
			// Handle case where new is &Address{} and base is &Address{"Line1"}
			if isZeroValue(valB, nilIsEmptyB) && !isZeroValue(valA, nilIsEmptyA) {
				isDiff = true
				reflect.ValueOf(delta).Elem().FieldByName(nameA).Set(valB)
			} else {
				d, diff := diff(valA.Interface(), valB.Interface(), nilIsEmptyA, nilIsEmptyB)
				if diff {
					isDiff = true
					reflect.ValueOf(delta).Elem().FieldByName(nameA).Set(reflect.ValueOf(d))
				}
			}
			continue
		}

		if isZeroValue(valA, nilIsEmptyA) && isZeroValue(valB, nilIsEmptyB) {
			continue
		}

		aI, bI := valA.Interface(), valB.Interface()

		comparableA, aOk := aI.(Comparable)
		comparableB, bOk := bI.(Comparable)

		if aOk && bOk {
			if !comparableA.Equal(comparableB) {
				reflect.ValueOf(delta).Elem().FieldByName(nameA).Set(valB)
				isDiff = true
			}
		} else if !reflect.DeepEqual(aI, bI) {
			reflect.ValueOf(delta).Elem().FieldByName(nameA).Set(valB)
			isDiff = true
		}
	}
	return delta, isDiff
}

func indirect(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}

func indirectKind(v reflect.Value) reflect.Kind {
	return indirect(v).Kind()
}

// Diff(a, b): a is base, b is new
func Diff(a Entity, b Entity) (Entity, bool) {
	if a.GetEntityType() != b.GetEntityType() {
		return nil, true
	}

	delta, isDiff := diff(a, b, getNilIsEmpty(a), getNilIsEmpty(b))
	if !isDiff {
		return nil, isDiff
	}
	return delta.(Entity), isDiff
}
