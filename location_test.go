package yext

import (
	"encoding/json"
	"reflect"
	"testing"
)

func jsonString(l *Location) (error, string) {
	buf, err := json.Marshal(l)
	if err != nil {
		return err, ""
	}

	return nil, string(buf)
}

func TestJSONSerialization(t *testing.T) {
	type test struct {
		l    *Location
		want string
	}

	tests := []test{
		{&Location{}, `{}`},
		{&Location{City: nil}, `{}`},
		{&Location{City: String("")}, `{"city":""}`},
		{&Location{Languages: nil}, `{}`},
		{&Location{Languages: nil}, `{}`},
		{&Location{Languages: &[]string{}}, `{"languages":[]}`},
		{&Location{Languages: &[]string{"English"}}, `{"languages":["English"]}`},
	}

	for _, test := range tests {
		if err, got := jsonString(test.l); err != nil {
			t.Error("Unable to convert", test.l, "to JSON:", err)
		} else if got != test.want {
			t.Errorf("json.Marshal(%#v) = %s; expected %s", test.l, got, test.want)
		}
	}
}

func TestJSONDeserialization(t *testing.T) {
	type test struct {
		json string
		want *Location
	}

	tests := []test{
		{`{}`, &Location{}},
		{`{"emails": []}`, &Location{Emails: Strings([]string{})}},
		{`{"emails": ["mhupman@yext.com", "bmcginnis@yext.com"]}`, &Location{Emails: Strings([]string{"mhupman@yext.com", "bmcginnis@yext.com"})}},
	}

	for _, test := range tests {
		v := &Location{}

		if err := json.Unmarshal([]byte(test.json), v); err != nil {
			t.Error("Unable to deserialize", test.json, "from JSON:", err)
		} else if !reflect.DeepEqual(v, test.want) {
			t.Errorf("json.Unmarshal(%#v) = %s; expected %s", test.json, v, test.want)
		}
	}
}
