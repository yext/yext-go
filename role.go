package yext

import (
	"encoding/json"
)

type Role struct {
	Id   *string `json:"roleId,omitempty"`
	Name *string `json:"roleName,omitempty"`
}

func (r *Role) GetId() string {
	if r.Id == nil {
		return ""
	}
	return *r.Id
}

func (r *Role) GetName() string {
	if r.Name == nil {
		return ""
	}
	return *r.Name
}

func (r *Role)SingleString() string {
	b, _ := json.Marshal(r)
	return string(b)
}
