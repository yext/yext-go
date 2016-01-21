package yext

import (
	"fmt"
	"net/http"
	"strings"
)

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"errorCode"`
}

type ErrorResponse struct {
	Errors   []Error `json:"errors"`
	Response *http.Response
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %d message: %s", e.Code, e.Message)
}

func (e *ErrorResponse) Error() string {
	errs := make([]string, len(e.Errors))

	for i, err := range e.Errors {
		errs[i] = err.Error()
	}

	return strings.Join(errs, "; ")
}
