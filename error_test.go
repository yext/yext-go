package yext

import (
	"errors"
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
			ErrorMessage: "",
			Want:         nil,
		},
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

func TestToUserFriendlyMessage(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Non-Yext errors should use default processing",
			args: args{err: errors.New("test")},
			want: "test",
		},
		{name: "Single Yext error should work",
			args: args{err: Error{
				Message:     "test",
				Code:        0,
				Type:        ErrorTypeNonFatal,
				RequestUUID: "testing-UUID",
			}},
			want: "test",
		},
		{name: "Error list with single error should work",
			args: args{err: Errors{
				&Error{
					Message:     "test",
					Code:        0,
					Type:        ErrorTypeNonFatal,
					RequestUUID: "testing-UUID",
				},
			}},
			want: "test",
		},
		{name: "Error list with multiple error should list all errors comma separated",
			args: args{err: Errors{
				&Error{
					Message:     "test",
					Code:        0,
					Type:        ErrorTypeNonFatal,
					RequestUUID: "testing-UUID",
				},
				&Error{
					Message:     "message 2",
					Code:        0,
					Type:        ErrorTypeNonFatal,
					RequestUUID: "testing-UUID",
				},
			}},
			want: "test, message 2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUserFriendlyMessage(tt.args.err); got != tt.want {
				t.Errorf("ToUserFriendlyMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
