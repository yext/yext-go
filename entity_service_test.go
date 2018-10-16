package yext

import "testing"

func TestSetNilIsEmpty(t *testing.T) {
	type randomStruct struct{}
	tests := []struct {
		i      interface{}
		before bool
		after  bool
	}{
		{
			i:      &BaseEntity{},
			before: false,
			after:  true,
		},
		{
			i: &BaseEntity{
				nilIsEmpty: true,
			},
			before: true,
			after:  true,
		},
		{
			i:      &LocationEntity{},
			before: false,
			after:  true,
		},
		{
			i:      &randomStruct{},
			before: false,
			after:  false,
		},
	}

	for _, test := range tests {
		before := getNilIsEmpty(test.i)
		if before != test.before {
			t.Errorf("Before set nil is empty: Expected %t, got %t", test.before, before)
		}
		setNilIsEmpty(test.i)
		after := getNilIsEmpty(test.i)
		if after != test.after {
			t.Errorf("After set nil is empty: Expected %t, got %t", test.after, after)
		}
	}
}
