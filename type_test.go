package yext

import (
	"testing"

	yext "github.com/yext/yext-go"
)

type TestStruct struct {
	Field1 *string
	Field2 *bool
}

func TestUnorderedSlicesEqual(t *testing.T) {
	tests := []struct {
		A    *UnorderedSlices
		B    *UnorderedSlices
		Want bool
	}{
		{
			A:    nil,
			B:    nil,
			Want: true,
		},
		{
			A:    nil,
			B:    ToUnorderedSlices([]*TestStruct{}),
			Want: false,
		},
		{
			A:    ToUnorderedSlices([]*TestStruct{}),
			B:    nil,
			Want: false,
		},
		{
			A: ToUnorderedSlices([]*TestStruct{
				&TestStruct{
					Field1: String("hi"),
					Field2: yext.Bool(true),
				},
			}),
			B: ToUnorderedSlices([]*TestStruct{
				&TestStruct{
					Field1: yext.String("hi"),
					Field2: yext.Bool(true),
				},
			}),
			Want: true,
		},
		{
			A: ToUnorderedSlices([]*TestStruct{
				&TestStruct{
					Field1: yext.String("hi"),
					Field2: yext.Bool(true),
				},
			}),
			B: ToUnorderedSlices([]*TestStruct{
				&TestStruct{
					Field1: yext.String("hi"),
					Field2: yext.Bool(false),
				},
			}),
			Want: false,
		},
		{
			A: ToUnorderedSlices([]*TestStruct{
				&TestStruct{
					Field1: yext.String("hi"),
					Field2: yext.Bool(true),
				},
				&TestStruct{
					Field1: yext.String("bye"),
					Field2: yext.Bool(false),
				},
			}),
			B: ToUnorderedSlices([]*TestStruct{
				&TestStruct{
					Field1: yext.String("bye"),
					Field2: yext.Bool(false),
				},
				&TestStruct{
					Field1: yext.String("hi"),
					Field2: yext.Bool(true),
				},
			}),
			Want: true,
		},
		{
			A: ToUnorderedSlices([]*TestStruct{
				&TestStruct{
					Field1: yext.String("hi"),
					Field2: yext.Bool(true),
				},
				&TestStruct{
					Field1: yext.String("bye"),
					Field2: yext.Bool(true),
				},
			}),
			B: ToUnorderedSlices([]*TestStruct{
				&TestStruct{
					Field1: yext.String("bye"),
					Field2: yext.Bool(false),
				},
				&TestStruct{
					Field1: yext.String("hi"),
					Field2: yext.Bool(true),
				},
			}),
			Want: false,
		},
	}
	for i, test := range tests {
		if got := test.A.Equal(test.B); got != test.Want {
			t.Errorf("TestUnorderedSlicesEqual %d: Expected: %t, Got: %t", i+1, test.Want, got)
		}
	}
}
