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
		aV = aV.Elem()
	}
	if bV.Kind() == reflect.Ptr {
		bV = bV.Elem()
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

		log.Println("valA", valA)
		log.Println("valB", valB)

		// If Kind() == struct, this is likely an embedded struct
		if valA.Kind() == reflect.Struct {
			log.Println("is struct")
			d, diff := diff(valA.Addr().Interface(), valB.Addr().Interface(), nilIsEmptyA, nilIsEmptyB)
			if diff {
				isDiff = true
				reflect.ValueOf(delta).Elem().FieldByName(nameA).Set(reflect.ValueOf(d).Elem())
			}
			continue
		} else if valA.Kind() == reflect.Ptr {
			// log.Println("is pointer")
			// if valA.IsNil() { // implies valB is non-nil
			// 	log.Println("I am nil")
			// 	//valBIndirect := valB.Elem()
			// 	log.Println("Nil is empty b", nilIsEmptyB)
			// 	log.Println("Nil is empty a", nilIsEmptyA)
			// 	log.Println("is zero a", isZeroValue(valA, nilIsEmptyA))
			// 	log.Println("is zero b", isZeroValue(valB.Elem(), nilIsEmptyB))
			// 	if isZeroValue(valA, nilIsEmptyA) && isZeroValue(valB, nilIsEmptyB) {
			// 		continue
			// 	}
			// 	isDiff = true
			// 	reflect.ValueOf(delta).Elem().FieldByName(nameA).Set(valB)
			// 	continue
			log.Println("isZeroA", isZeroValue(valA, nilIsEmptyA))
			log.Println("isZeroB", isZeroValue(valB, nilIsEmptyB))
			if !valB.IsNil() && valB.CanSet() {
				log.Println("val B is not nil")
				valAIndirect := valA.Elem()
				valBIndirect := valB.Elem()
				// log.Println("valAInd", valAIndirect)
				// log.Println("valBInd", valBIndirect)
				// log.Println("kind:", valAIndirect.Kind())
				if valAIndirect.Kind() == reflect.Struct {
					// If base is &Address{Line1:"abc"} and new is &Address{}, we want &Address for the diff
					// if isZeroValue(valBIndirect, getNilIsEmpty(valBIndirect)) {
					// 	log.Println("val b is zero")
					// 	if !(isZeroValue(valAIndirect, getNilIsEmpty(valAIndirect))) {
					// 		log.Println("val a is not zero")
					// 		isDiff = true
					// 		reflect.ValueOf(delta).Elem().FieldByName(nameA).Set(valB)
					// 	}
					// 	continue
					// }
					reflect.ValueOf(delta).Elem().FieldByName(nameA).Set(valB)
					d, diff := diff(valAIndirect.Addr().Interface(), valBIndirect.Addr().Interface(), nilIsEmptyA, nilIsEmptyB)
					if diff {
						isDiff = true
						reflect.ValueOf(delta).Elem().FieldByName(nameA).Set(reflect.ValueOf(d))
					}
					continue
				}
			}
		}

		// before we want to recur we want to make sure neither is zoer

		if valB.Kind() == reflect.Ptr && valB.IsNil() {
			continue
		}
		if !valB.CanSet() {
			continue
		}

		if isZeroValue(valA, getNilIsEmpty(a)) && isZeroValue(valB, getNilIsEmpty(b)) {
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
