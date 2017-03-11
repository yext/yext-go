package yext

import (
	"encoding/json"
	"fmt"
)

type AccessOn string

const (
	ACCESS_ACCOUNT  = AccessOn("ACCOUNT")
	ACCESS_FOLDER   = AccessOn("FOLDER")
	ACCESS_LOCATION = AccessOn("LOCATION")
)

type ACL struct {
	Role
	On        string   `json:"on,omitempty"`
	AccessOn  AccessOn `json:"onType"`
	AccountId string   `json:"accountId"`
}

func (a ACL) String() string {
	b, _ := json.Marshal(a)
	return string(b)
}

// Hash returns a string representation of the elements that make an ACL
// functionally unique in Yext.  Useful for comparing ACLs for "real-world"
// equality - "Do two ACLs have the same effect?"
func (a ACL) Hash() string {
	// Not including AccountId in the hash is intentional - APIv2 has partial support
	// for the alias `me` to mean AccountId associated to the current API key.
	// This complicates our diffing since the APIv2 returns the numeric AccountId
	// instead of the `me` alias when return users, so diffing against them might
	// lead to spurious diffs.
	return fmt.Sprintf("%s-%s-%s", a.GetId(), a.On, a.AccessOn)
}
