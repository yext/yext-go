package yext

import (
	"encoding/json"
	"fmt"
)

type AccessOn string

const (
	ACCESS_CUSTOMER = AccessOn("CUSTOMER")
	ACCESS_FOLDER   = AccessOn("FOLDER")
	ACCESS_LOCATION = AccessOn("LOCATION")
)

type ACL struct {
	Role
	On       string   `json:"on,omitempty"`
	AccessOn AccessOn `json:"onType"`
}

func (a ACL) String() string {
	b, _ := json.Marshal(a)
	return string(b)
}

func (a ACL) Hash() string {
	return fmt.Sprintf("%d-%s-%s", a.Id, a.On, a.AccessOn)
}
