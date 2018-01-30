package yext

import "testing"

func TestIsNotFoundError(t *testing.T) {
	tests := []struct {
		Err  error
		Want bool
	}{
		{
			Err: Errors{
				&Error{
					Type:    "NON_FATAL_ERROR",
					Code:    202,
					Message: "Some message",
				},
				&Error{
					Type:    "FATAL_ERROR",
					Code:    6004,
					Message: "Some message",
				},
			},
			Want: true,
		},
		{
			Err: &Error{
				Type:    "FATAL_ERROR",
				Code:    6004,
				Message: "Some message",
			},
			Want: true,
		},
		{
			Err: &Error{
				Type:    "FATAL_ERROR",
				Code:    2000,
				Message: "Some message",
			},
			Want: true,
		},
		{
			Err: &Error{
				Type:    "NON_FATAL_ERROR",
				Code:    202,
				Message: "Some message",
			},
			Want: false,
		},
	}

	for i, test := range tests {
		if got := IsNotFoundError(test.Err); got != test.Want {
			t.Errorf("TestIsNotFoundError %d failed: Wanted %t, got %t\n", i+1, test.Want, got)
		}
	}
}
