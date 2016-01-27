package yext

import (
	"fmt"
)

const (
	userPath = "users"
	rolePath = "roles"
)

type UserService struct {
	client *Client
}

type userListResponse struct {
	Users []*User `json:"users"`
}

type rolesRepsonse struct {
	Roles []*Role `json:"roles"`
}

func (u *UserService) List() ([]*User, error) {
	v := &userListResponse{}
	err := u.client.DoRequest("GET", userPath, v)
	return v.Users, err
}

func (u *User) pathToUser() string {
	return pathToUserId(u.GetId())
}

func pathToUserId(id string) string {
	return fmt.Sprintf("%s/%s", userPath, id)
}

func (u *UserService) Get(id string) (*User, error) {
	var v = &User{}
	err := u.client.DoRequest("GET", pathToUserId(id), v)
	return v, err
}

func (u *UserService) Edit(y *User) (*User, error) {
	var v = &User{}
	err := u.client.DoRequestJSON("PUT", y.pathToUser(), y, v)
	return v, err
}

func (u *UserService) Create(y *User) (*User, error) {
	var v = &User{}
	err := u.client.DoRequestJSON("POST", userPath, y, v)
	return v, err
}

func (u *UserService) Delete(y *User) error {
	return u.client.DoRequest("DELETE", y.pathToUser(), nil)
}

func (u *UserService) AvailableRoles() ([]*Role, error) {
	v := &rolesRepsonse{}
	err := u.client.DoRawRequest("GET", rolePath, v)
	return v.Roles, err
}

func (u *UserService) NewFolderACL(f *Folder, r *Role) *ACL {
	return &ACL{
		Role:     r,
		On:       String(f.Id),
		AccessOn: ACCESS_FOLDER,
	}
}

func (u *UserService) NewCustomerACL(r *Role) *ACL {
	return &ACL{
		Role:     r,
		On:       String(u.client.customerId),
		AccessOn: ACCESS_CUSTOMER,
	}
}

func (u *UserService) NewLocationACL(l *Location, r *Role) *ACL {
	return &ACL{
		Role:     r,
		On:       l.Id,
		AccessOn: ACCESS_LOCATION,
	}
}
