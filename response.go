package yext

import "encoding/json"

type Meta struct {
	UUID   string `json:"uuid"`
	Errors Errors `json:"errors,omitempty"`
}

type Response struct {
	// HttpResponse *http.Response Maybe want this in the future?
	Meta        Meta             `json:"meta"`
	Response    interface{}      `json:"-"`
	ResponseRaw *json.RawMessage `json:"response,omitempty"`
}
