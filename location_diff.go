package yext

import (
	"reflect"
)

type Comparable interface {
	Equal(Comparable) bool
}

// Diff calculates the differences between a base Location (a) and a proposed set of changes
// represented by a second Location (b).  The diffing logic will ignore fields in the proposed
// Location that aren't set (nil).  This characteristic makes the function ideal for
// calculting the minimal PUT required to bring an object "up-to-date" via the API.
//
// Example:
//   var (
//     // Typically, this would come from an authoritative source, like the API
//     base     = Location{Name: "Yext, Inc.", Address1: "123 Mulberry St"}
//     proposed = Location{Name: "Yext, Inc.", Address1: "456 Washington St"}
//   )
//
//   delta, isDiff := base.Diff(proposed)
//   // isDiff -> true
//   // delta -> &Location{Address1: "456 Washington St"}
//
//   delta, isDiff = base.Diff(base)
//   // isDiff -> false
//   // delta -> nil
func (y Location) Diff(b *Location) (d *Location, diff bool) {
	// TODO (bjm): remove this once the ETL is upgraded to use hydration
	if (!y.hydrated || !b.hydrated) && y.String() == b.String() {
		return nil, false
	}

	d = &Location{Id: String(y.GetId())}

	var (
		aV, bV = reflect.ValueOf(y), reflect.ValueOf(b).Elem()
		aT     = reflect.TypeOf(y)
		numA   = aV.NumField()
	)

	for i := 0; i < numA; i++ {
		var (
			nameA = aT.Field(i).Name
			valA  = aV.Field(i)
			valB  = bV.Field(i)
		)

		if !valB.CanSet() {
			continue
		}

		if valB.IsNil() || nameA == "Id" {
			continue
		}

		if y.nilIsEmpty && valA.IsNil() {
			valA = GetZeroOf(valA)
		}
		if b.nilIsEmpty && valB.IsNil() {
			valB = GetZeroOf(valB)
		}

		aI, bI := valA.Interface(), valB.Interface()

		comparableA, aOk := aI.(Comparable)
		comparableB, bOk := bI.(Comparable)

		if aOk && bOk {
			if !comparableA.Equal(comparableB) {
				diff = true
				reflect.ValueOf(d).Elem().FieldByName(nameA).Set(valB)
			}
		} else if !reflect.DeepEqual(aI, bI) {
			if nameA == "CustomFields" {
				d.CustomFields = make(map[string]interface{})
				for field, value := range b.CustomFields {
					if aValue, ok := y.CustomFields[field]; ok {
						var valueDiff bool
						if v, ok := aValue.(Comparable); ok {
							valueDiff = !v.Equal(value.(Comparable))
						} else if !reflect.DeepEqual(aValue, value) {
							valueDiff = true
						}
						if valueDiff {
							diff = true
							d.CustomFields[field] = value
						}
					} else {
						d.CustomFields[field] = value
						diff = true
					}
				}
			} else {
				diff = true
				reflect.ValueOf(d).Elem().FieldByName(nameA).Set(valB)
			}
		}
	}
	if diff == false {
		// ensure d remains nil if nothing has changed
		d = nil
	}

	return d, diff
}

func GetZeroOf(val reflect.Value) reflect.Value {
	switch val.Type() {
	case reflect.PtrTo(reflect.TypeOf("")):
		return reflect.New(reflect.TypeOf(""))
	case reflect.PtrTo(reflect.TypeOf(false)):
		return reflect.New(reflect.TypeOf(false))
	case reflect.TypeOf(&[]string{}):
		return reflect.ValueOf(&[]string{})
	case reflect.TypeOf(&GoogleAttributes{}):
		return reflect.ValueOf(&GoogleAttributes{})
	case reflect.TypeOf(&LocationPhoto{}):
		return reflect.ValueOf(&LocationPhoto{})
	case reflect.TypeOf(&[]LocationPhoto{}):
		return reflect.ValueOf(&[]LocationPhoto{})
	case reflect.TypeOf(&LocationClosed{}):
		return reflect.ValueOf(&LocationClosed{})
	}
	return reflect.Zero(val.Type())
}
