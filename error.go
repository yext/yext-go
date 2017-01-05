package yext

import (
	"fmt"
	"strings"
)

type ErrorType string

const (
	ErrorTypeFatal    = "FATAL_ERROR"
	ErrorTypeNonFatal = "NON_FATAL_ERROR"
	ErrorTypeWarning  = "WARNING"
)

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Type    string `json:"type"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("type: %s code: %d message: %s", e.Type, e.Code, e.Message)
}

func (e *Error) IsError() bool {
	return e.Type == ErrorTypeFatal || e.Type == ErrorTypeNonFatal
}

func (e *Error) IsWarning() bool {
	return e.Type == ErrorTypeWarning
}

type Errors []Error

func (e Errors) Error() string {
	errs := make([]string, len(e))

	for i, err := range e {
		errs[i] = err.Error()
	}

	return strings.Join(errs, "; ")
}

func (e Errors) Errors() []Error {
	var errors []Error
	for _, err := range e {
		if err.IsError() {
			errors = append(errors, err)
		}
	}
	return errors
}

func (e Errors) Warnings() []Error {
	var warnings []Error
	for _, err := range e {
		if err.IsWarning() {
			warnings = append(warnings, err)
		}
	}
	return warnings
}
