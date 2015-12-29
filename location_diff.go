package yext

import (
	"reflect"
)

func (a Location) Diff(b Location) (d *Location, diff bool) {
	diff = false

	var (
		aV, bV = reflect.ValueOf(a), reflect.ValueOf(b)
		aT     = reflect.TypeOf(a)
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
			d = new(Location)
			if nameA == "CustomFields" {
				// deal with case where left is nil and right is empty
				if a.CustomFields == nil && b.CustomFields != nil {
					diff = true
				}
				d.CustomFields = make(map[string]interface{})
				for field, value := range b.CustomFields {
					if aValue, ok := a.CustomFields[field]; ok {
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
