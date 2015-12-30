package yext

import (
	"reflect"
)

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
func (y Location) Diff(b Location) (d *Location, diff bool) {
	diff, d = false, new(Location)

	var (
		aV, bV = reflect.ValueOf(y), reflect.ValueOf(b)
		aT     = reflect.TypeOf(y)
		numA   = aV.NumField()
	)

	for i := 0; i < numA; i++ {
		var (
			nameA = aT.Field(i).Name
			valA  = aV.Field(i)
			valB  = bV.Field(i)
		)

		if valB.IsNil() {
			continue
		}

		if !reflect.DeepEqual(valA.Interface(), valB.Interface()) {
			if nameA == "CustomFields" {
				// deal with case where left is nil and right is empty
				if y.CustomFields == nil && b.CustomFields != nil {
					diff = true
				}
				d.CustomFields = make(map[string]interface{})
				for field, value := range b.CustomFields {
					if aValue, ok := y.CustomFields[field]; ok {
						if !reflect.DeepEqual(aValue, value) {
							d.CustomFields[field] = value
							diff = true
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

func (y Location) Copy() (n *Location) {
	n = new(Location)
	v := reflect.ValueOf(y)
	t := reflect.TypeOf(y)
	num := v.NumField()
	for i := 0; i < num; i++ {
		fieldName := t.Field(i).Name
		if fieldName == "CustomFields" {
			n.CustomFields = make(map[string]interface{})
			for field, value := range y.CustomFields {
				n.CustomFields[field] = value
			}
		} else {
			reflect.ValueOf(n).Elem().FieldByName(fieldName).Set(v.Field(i))
		}
	}
	return n
}
