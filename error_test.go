package yext

import (
	"testing"
)

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

func TestErrorsFromString(t *testing.T) {
	tests := []struct {
		ErrorMessage string
		Want         []*Error
	}{
		{
			ErrorMessage: "type: FATAL_ERROR code: 2015 message: Unknown folder; request uuid: 7199948d-9f0d-4649-9625-495b33ad4940",
			Want: []*Error{
				&Error{
					Type:        "FATAL_ERROR",
					Code:        2015,
					Message:     "Unknown folder",
					RequestUUID: "7199948d-9f0d-4649-9625-495b33ad4940",
				},
			},
		},
		{
			ErrorMessage: "type: FATAL_ERROR code: 2106 message: featuredMessageUrl: The url provided is invalid.; type: FATAL_ERROR code: 2103 message: websiteUrl: The url provided is invalid.; request uuid: 3b03b517-51c5-4a64-8285-a3466ce875f6",
			Want: []*Error{
				&Error{
					Type:        "FATAL_ERROR",
					Code:        2106,
					Message:     "featuredMessageUrl: The url provided is invalid.",
					RequestUUID: "3b03b517-51c5-4a64-8285-a3466ce875f6",
				},
				&Error{
					Type:        "FATAL_ERROR",
					Code:        2103,
					Message:     "websiteUrl: The url provided is invalid.",
					RequestUUID: "3b03b517-51c5-4a64-8285-a3466ce875f6",
				},
			},
		},
	}

	for i, test := range tests {
		got, err := ErrorsFromString(test.ErrorMessage)
		if err != nil {
			t.Errorf("TestErrorsFromString %d failed: %s", i+1, err.Error())
		} else if len(got) != len(test.Want) {
			t.Errorf("TestErrorsFromString %d failed: \ngot\n\t%v\nexpected\n\t%v", i+1, got, test.Want)
		} else {
			for i, errObj := range got {
				if *errObj != *test.Want[i] {
					t.Errorf("TestErrorsFromString %d failed: \ngot\n\t%v\nexpected\n\t%v", i+1, got, test.Want)
				}
			}
		}
	}
}
