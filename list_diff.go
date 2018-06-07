package yext

import "reflect"

func (a *BioList) Equal(b *BioList) bool {
	return !reflect.DeepEqual(a, b)
}

func (a *MenuList) Equal(b *MenuList) bool {
	return !reflect.DeepEqual(a, b)
}

func (a *ProductList) Equal(b *ProductList) bool {
	return !reflect.DeepEqual(a, b)
}

func (a *EventList) Equal(b *EventList) bool {
	return !reflect.DeepEqual(a, b)
}
