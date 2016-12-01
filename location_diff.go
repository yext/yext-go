package yext

import "reflect"

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
	// TODO (bjm):
	// remove this once the ETL is upgraded to use hydration
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

		if !reflect.DeepEqual(valA.Interface(), valB.Interface()) {
			if nameA == "CustomFields" {
				// deal with case where left is nil and right is empty
				if y.CustomFields == nil && b.CustomFields != nil {
					diff = true
				}
				d.CustomFields = make(map[string]interface{})
				for field, value := range b.CustomFields {
					if aValue, ok := y.CustomFields[field]; ok {
						var valueDiff bool
						if v, ok := aValue.(CustomFieldValueComparable); ok {
							valueDiff = !v.Equal(value.(CustomFieldValueComparable))
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
