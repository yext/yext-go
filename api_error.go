package yext

import (
	"encoding/json"
	"fmt"
)

type APIError struct {
	Message string `json:"message"`
	Code    int    `json:"errorCode"`
}

type APIErrorResponse struct {
	Errors   []*APIError `json:"errors"`
	HTTPCode int
}

func (e *APIError) Error() string {
	return fmt.Sprintf("Recieved error code: '%d', messsage was: '%s'", e.Code, e.Message)
}

func (e *APIErrorResponse) Error() string {
	var errStr string

	for _, err := range e.Errors {
		errStr += fmt.Sprintf("%s\n", err.Error())
	}
	return errStr
}

func UnmarhsalAPIError(httpCode int, respBody []byte) (e *APIErrorResponse) {
	if len(respBody) == 0 {
		return nil
	}

	var errResp = new(APIErrorResponse)
	err := json.Unmarshal(respBody, errResp)
	if err != nil {
		ae := &APIError{
			Code:    -1,
			Message: fmt.Sprintf("Unknown API Error %v", string(respBody)),
		}
		errResp.Errors = append(errResp.Errors, ae)
	}

	errResp.HTTPCode = httpCode

	return errResp
}
