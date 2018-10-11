package yext

import (
	"reflect"
)

func copyEntity(val interface{}) interface{} {
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
	return reflect.New(reflect.TypeOf(newVal)).Interface()
}

func diff(a interface{}, b interface{}, delta interface{}) bool {
	var (
		aV, bV = reflect.ValueOf(a), reflect.ValueOf(b)
		isDiff = false
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

		if nameA == "BaseEntity" { // TODO: Need to handle this case
			continue
		}

		if valA.Kind() == reflect.Struct {
			diff := diff(valA.Addr().Interface(), valB.Addr().Interface(), reflect.ValueOf(delta).Elem().FieldByName(nameA).Addr().Interface())
			if diff {
				isDiff = true
			}
			continue
		} else if valA.Kind() == reflect.Ptr {
			if !valB.IsNil() || !valB.CanSet() { // should we check can set?
				valAIndirect := valA.Elem()
				valBIndirect := valB.Elem()
				if valAIndirect.Kind() == reflect.Struct {
					reflect.ValueOf(delta).Elem().FieldByName(nameA).Set(valB)
					diff := diff(valAIndirect.Addr().Interface(), valBIndirect.Addr().Interface(), reflect.ValueOf(delta).Elem().FieldByName(nameA).Interface())
					if diff {
						isDiff = true
					}
					continue
				}
			}
		}

		if valB.IsNil() || !valB.CanSet() {
			continue
		}

		// if isZeroValue(valA, y.nilIsEmpty) && isZeroValue(valB, b.nilIsEmpty) {
		// 	continue
		// }
		//

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
	return isDiff
}

func Diff(a Entity, b Entity) (Entity, bool) {
	if a.GetEntityType() != b.GetEntityType() {
		return nil, true
	}

	baseEntity := BaseEntity{
		Meta: &EntityMeta{
			Id: String(a.GetEntityId()),
		},
	}

	// TODO: figure out how to handle other base entity attributes
	delta := copyEntity(a)
	reflect.ValueOf(delta).Elem().FieldByName("BaseEntity").Set(reflect.ValueOf(baseEntity))

	isDiff := diff(a, b, delta)
	if !isDiff {
		return nil, isDiff
	}
	return delta.(Entity), isDiff
}
