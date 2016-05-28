package yext

import (
	"encoding/json"
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
    l *Location
    want string
  }

  tests := []test{
    {&Location{}, `{}`},
		{&Location{City: nil}, `{}`},
		{&Location{City: String("")}, `{"city":""}`},
		{&Location{Languages: nil}, `{}`},
		{&Location{Languages: []string(nil)}, `{}`},
		{&Location{Languages: []string{}}, `{"languages":[]}`},
		{&Location{Languages: []string{"English"}}, `{"languages":["English"]}`},
  }

  for _, test := range tests {
    if err, got := jsonString(test.l); err != nil {
      t.Error("Unable to convert", test.l, "to JSON:", err)
    } else if got != test.want {
      t.Errorf("json.Marshal(%#v) = %s; expected %s", test.l, got, test.want)
    }
  }
}
