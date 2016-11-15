package yext

import (
	"fmt"
	"net/http"
	"strings"
)

type Response struct {
	Meta         Meta        `json:"meta"`
	Body         interface{} `json:"response"`
	HTTPResponse *http.Response
}

type Meta struct {
	Errors []Error `json:"errors"`
	UUID   string  `json:"uuid"`
}

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Type    string `json:"type"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s code: %d message: %s", e.Type, e.Code, e.Message)
}

func (m *Meta) Error() string {
	errs := make([]string, len(m.Errors))

	for i, err := range m.Errors {
		errs[i] = err.Error()
	}

	return strings.Join(errs, "; ")
}

func (r *Response) Error() string {
	return r.Meta.Error()
}
