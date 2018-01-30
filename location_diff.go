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

		if nameA == "Hours" {
			if HoursAreEquivalent(getUnderlyingValue(valA.Interface()).(string), getUnderlyingValue(valB.Interface()).(string)) {
				continue
			}
		}

		if isZeroValue(valA, y.nilIsEmpty) && isZeroValue(valB, b.nilIsEmpty) {
			continue
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
					value = getUnderlyingValue(value)
					if aValue, ok := y.CustomFields[field]; ok {
						aValue = getUnderlyingValue(aValue)
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
					} else if !(isZeroValue(reflect.ValueOf(value), b.nilIsEmpty) && y.nilIsEmpty) {
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

// getUnderlyingValue unwraps a pointer/interface to get the underlying value.
// If the value is already unwrapped, it returns it as is.
func getUnderlyingValue(v interface{}) interface{} {
	if v == nil {
		return nil
	}

	rv := reflect.ValueOf(v)

	switch rv.Kind() {
	case reflect.Ptr, reflect.Interface:
		rv = rv.Elem()
	}

	return rv.Interface()
}

func isZeroValue(v reflect.Value, interpretNilAsZeroValue bool) bool {
	if !v.IsValid() {
		return true
	}

	switch v.Kind() {
	case reflect.Slice, reflect.Map:
		return v.Len() == 0
	case reflect.Bool:
		return v.Bool() == false
	case reflect.String:
		return v.String() == ""
	case reflect.Float64:
		return v.Float() == 0.0
	case reflect.Ptr, reflect.Interface:
		if v.IsNil() && !interpretNilAsZeroValue {
			return false
		}
		return isZeroValue(v.Elem(), interpretNilAsZeroValue)
	case reflect.Struct:
		for i, n := 0, v.NumField(); i < n; i++ {
			if !isZeroValue(v.Field(i), interpretNilAsZeroValue) {
				return false
			}
		}
		return true
	default:
		return v.IsNil() && interpretNilAsZeroValue
	}
}

var closedHoursEquivalents = map[string]struct{}{
	"": struct{}{},
	"1:closed,2:closed,3:closed,4:closed,5:closed,6:closed,7:closed": struct{}{},
}

func HoursAreEquivalent(a, b string) bool {
	_, aIsClosed := closedHoursEquivalents[a]
	_, bIsClosed := closedHoursEquivalents[b]

	return a == b || aIsClosed && bIsClosed
}
