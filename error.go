package yext

import (
	"fmt"
	"strconv"
	"strings"
)

type ErrorType string

const (
	ErrorTypeFatal    = "FATAL_ERROR"
	ErrorTypeNonFatal = "NON_FATAL_ERROR"
	ErrorTypeWarning  = "WARNING"
)

type Error struct {
	Message     string `json:"message"`
	Code        int    `json:"code"`
	Type        string `json:"type"`
	RequestUUID string `json:"request_uuid"`
}

func (e Error) Error() string {
	return fmt.Sprintf("type: %s code: %d message: %s, request uuid: %s", e.Type, e.Code, e.Message, e.RequestUUID)
}

func (e Error) ErrorWithoutUUID() string {
	return fmt.Sprintf("type: %s code: %d message: %s", e.Type, e.Code, e.Message)
}

func (e Error) IsError() bool {
	return e.Type == ErrorTypeFatal || e.Type == ErrorTypeNonFatal
}

func (e Error) IsWarning() bool {
	return e.Type == ErrorTypeWarning
}

type Errors []*Error

func (e Errors) Error() string {
	var (
		errs = make([]string, len(e))
		uuid = ""
	)

	for i, err := range e {
		errs[i] = err.ErrorWithoutUUID()
		uuid = err.RequestUUID
	}

	return fmt.Sprintf("%s; request uuid: %s", strings.Join(errs, "; "), uuid)
}

func (e Errors) Errors() []*Error {
	var errors []*Error
	for _, err := range e {
		if err.IsError() {
			errors = append(errors, err)
		}
	}
	return errors
}

func (e Errors) Warnings() []*Error {
	var warnings []*Error
	for _, err := range e {
		if err.IsWarning() {
			warnings = append(warnings, err)
		}
	}
	return warnings
}

func IsNotFoundError(err error) bool {
	if e, ok := err.(Errors); ok {
		for _, innerError := range e {
			if IsNotFoundError(innerError) {
				return true
			}
		}
	} else if e, ok := err.(*Error); ok {
		if e.Code == 2000 || e.Code == 6004 {
			return true
		}
	}
	return false
}

func IsErrorCode(err error, code int) bool {
	if e, ok := err.(Errors); ok {
		for _, innerError := range e {
			if IsErrorCode(innerError, code) {
				return true
			}
		}
	} else if e, ok := err.(*Error); ok {
		if e.Code == code {
			return true
		}
	}
	return false
}

func SplitStrAtWord(str string, word string) (string, string) {
	words := strings.Split(str, " ")
	found := false
	before := ""
	after := ""
	for _, w := range words {
		if w == word {
			found = true
		} else if found {
			if after != "" {
				after += " "
			}
			after += w
		} else {
			if before != "" {
				before += " "
			}
			before += w
		}
	}
	return before, after
}

func StrToError(str string) *Error {
	strRemaining := strings.TrimLeft(str, "type: ")
	typ, strRemaining := SplitStrAtWord(strRemaining, "code:")
	code, message := SplitStrAtWord(strRemaining, "message:")
	codeInt, _ := strconv.Atoi(code)
	return &Error{Type: typ, Code: codeInt, Message: message}
}

func ParseErrorStr(errorStr string) []*Error {
	errStrList := strings.Split(errorStr, "; ")
	var errors []*Error
	uuid := strings.TrimLeft(errStrList[len(errStrList)-1], "request uuid: ")
	for i := 0; i < len(errStrList)-1; i++ {
		err := StrToError(errStrList[i])
		err.RequestUUID = uuid
		errors = append(errors, err)
	}
	return errors
}
