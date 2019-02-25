package yext

import (
	"reflect"
)

func (a *ReviewCreate) Diff(b *ReviewCreate) (d *ReviewCreate, diff bool) {
	d = &ReviewCreate{LocationId: String(a.GetLocationId())}

	var (
		aValue, bValue = reflect.ValueOf(a).Elem(), reflect.ValueOf(b).Elem()
		aType          = reflect.TypeOf(*a)
		numA           = aValue.NumField()
	)

	for i := 0; i < numA; i++ {
		var (
			nameA = aType.Field(i).Name
			valA  = aValue.Field(i)
			valB  = bValue.Field(i)
		)

		if valB.IsNil() || nameA == "Id" {
			continue
		}

		if !reflect.DeepEqual(valA.Interface(), valB.Interface()) {
			diff = true
			reflect.ValueOf(d).Elem().FieldByName(nameA).Set(valB)
		}
	}
	if !diff {
		// ensure that d remains nil if nothing has changed
		d = nil
	}
	return d, diff
}
