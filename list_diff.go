package yext

import "reflect"

func (a *BioList) Diff(b *BioList) bool {
	return !reflect.DeepEqual(a, b)
}

func (a *MenuList) Diff(b *MenuList) bool {
	return !reflect.DeepEqual(a, b)
}

func (a *ProductList) Diff(b *ProductList) bool {
	return !reflect.DeepEqual(a, b)
}

func (a *EventList) Diff(b *EventList) bool {
	return !reflect.DeepEqual(a, b)
}
