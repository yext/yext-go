package yext

import (
	"encoding/json"
)

type AccessOn string

const (
	ACCESS_CUSTOMER = AccessOn("CUSTOMER")
	ACCESS_FOLDER   = AccessOn("FOLDER")
	ACCESS_LOCATION = AccessOn("LOCATION")
)

type ACL struct {
	*Role
	On       *string  `json:"on,omitempty"`
	AccessOn AccessOn `json:"onType"`
}

func (a *ACL) GetOn() string {
	if a.On == nil {
		return ""
	}
	return *a.On
}

func (a *ACL) String() string {
	b, _ := json.Marshal(a)
	return string(b)
}
