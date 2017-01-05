package yext

import (
	"encoding/json"
)

type Role struct {
	Id   int     `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

func (r *Role) GetName() string {
	if r.Name == nil {
		return ""
	}
	return *r.Name
}

func (r *Role) String() string {
	b, _ := json.Marshal(r)
	return string(b)
}
