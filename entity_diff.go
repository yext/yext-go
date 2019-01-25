package yext

import (
	"fmt"
	"log"
	"reflect"

	"github.com/davecgh/go-spew/spew"
)

func instanceOf(val interface{}) interface{} {
	var (
		isPtr = reflect.ValueOf(val).Kind() == reflect.Ptr
		tmp   interface{}
	)
	if isPtr {
		var numPointers = 0
		for reflect.ValueOf(val).Kind() == reflect.Ptr {
			log.Println(reflect.ValueOf(val).Kind())
			val = reflect.ValueOf(val).Elem().Interface()
			numPointers++
		}
		log.Println("numPointers", numPointers)
		tmp = val
		i := 0
		for i < numPointers {
			log.Println(i)
			// if i == numPointers-1 {
			// 	tmp = reflect.New(reflect.TypeOf(tmp)).Interface()
			// } else {
			tmp = reflect.New(reflect.TypeOf(tmp)).Interface()
			//}
			i++
		}
		log.Println("type of ", reflect.ValueOf(tmp).Kind())
		// log.Println("here")
		// var s = DoubleString("blah")
		// var v = reflect.ValueOf(s).Elem().Elem().Interface()
		// log.Println(v, reflect.ValueOf(s).Kind(), reflect.ValueOf(v).Kind())
		// log.Println("typeof ", reflect.TypeOf(v))
		// log.Println(reflect.New(reflect.TypeOf(v)).Kind())
		// var ptr = reflect.New(reflect.TypeOf(v)).Interface()
		// log.Println("typeof ", reflect.TypeOf(ptr))
		// log.Println(reflect.New(reflect.TypeOf(ptr)))
		return tmp
		//tmp = reflect.ValueOf(val).Elem().Interface()
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

		if nameA == "nilIsEmpty" {
			continue
		}

		if isNil(valB) {
			continue
		}

		if !valB.CanSet() {
			continue
		}

		aI, bI := valA.Interface(), valB.Interface()
		// Comparable does not handle the nil is empty case:
		// So if valA is nil, don't call comparable (valB checked for nil above)
		if !isNil(valA) {
			comparableA, aOk := aI.(Comparable)
			comparableB, bOk := bI.(Comparable)
			if aOk && bOk {
				if !comparableA.Equal(comparableB) {
					return b, true
				}
				return nil, false
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
		} else if valA.Kind() == reflect.Ptr && indirect(valA).Kind() == reflect.Struct {
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

		log.Println(nameA)
		log.Println(valB)
		if !reflect.DeepEqual(aI, bI) {
			log.Print("kind ", reflect.ValueOf(delta).Kind())
			var d = indirect(reflect.ValueOf(delta)) // reflect.ValueOf(delta).Elem()
			log.Println("delta")
			spew.Dump(delta)
			d.FieldByName(nameA).Set(valB)
			isDiff = true
		}
	}
	return delta, isDiff
}

func indirect(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		log.Println(v, "was ptr")
		v = v.Elem()
	}
	return v
}

func isNil(v reflect.Value) bool {
	if v.Kind() == reflect.Ptr {
		return v.IsNil()
	}
	return false
}

// Diff(a, b): a is base, b is new
func Diff(a Entity, b Entity) (Entity, bool, error) {
	// TODO: should the below return an error? If not should return an empty b object with entity type set?
	if a.GetEntityType() != b.GetEntityType() {
		return nil, true, fmt.Errorf("Entity Types do not match: '%s', '%s'", a.GetEntityType(), b.GetEntityType())
	}

	delta, isDiff := diff(a, b, getNilIsEmpty(a), getNilIsEmpty(b))
	if !isDiff {
		return nil, isDiff, nil
	}
	return delta.(Entity), isDiff, nil
}
