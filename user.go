package yext

import (
	"encoding/json"
)

type User struct {
	Id           *string `json:"id,omitempty"`        // req in post
	FirstName    *string `json:"firstName,omitempty"` // req in post
	LastName     *string `json:"lastName,omitempty"`  // req in post
	UserName     *string `json:"username,omitempty"`
	EmailAddress *string `json:"emailAddress,omitempty"` // req in post
	PhoneNumber  *string `json:"phoneNumber,omitempty"`
	Password     *string `json:"password,omitempty"`
	ACLs         []ACL   `json:"acl,omitempty"`
}

func (u *User) GetId() string {
	if u.Id == nil {
		return ""
	}
	return *u.Id
}

func (u *User) GetFirstName() string {
	if u.FirstName == nil {
		return ""
	}
	return *u.FirstName
}

func (u *User) GetLastName() string {
	if u.LastName == nil {
		return ""
	}
	return *u.LastName
}

func (u *User) GetUserName() string {
	if u.UserName == nil {
		return ""
	}
	return *u.UserName
}
func (u *User) GetEmailAddress() string {
	if u.EmailAddress == nil {
		return ""
	}
	return *u.EmailAddress
}

func (u *User) GetPhoneNumber() string {
	if u.PhoneNumber == nil {
		return ""
	}
	return *u.PhoneNumber
}

func (u *User) GetPassword() string {
	if u.Password == nil {
		return ""
	}
	return *u.Password
}

func (u *User) String() string {
	b, _ := json.Marshal(u)
	return string(b)
}
