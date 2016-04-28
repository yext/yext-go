package yext

import (
	"fmt"
	"net/url"
	"strconv"
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

func (u *UserService) ListAll() ([]*User, error) {
	var (
		start     int
		allUsers  []*User
		increment = 1000
	)

	for {
		users, err := u.List(start, increment)
		if err != nil {
			return nil, err
		}

		allUsers = append(allUsers, users...)

		if len(users) < increment {
			break
		}

		start += increment
	}
	return allUsers, nil
}

func (u *UserService) List(start, limit int) ([]*User, error) {
	userUrl := url.URL{Path: userPath}

	if start > 0 {
		q := userUrl.Query()
		q.Set("start", strconv.Itoa(start))
		userUrl.RawQuery = q.Encode()
	}

	if limit > 0 {
		q := userUrl.Query()
		q.Set("limit", strconv.Itoa(limit))
		userUrl.RawQuery = q.Encode()
	}

	v := &userListResponse{}
	err := u.client.DoRequest("GET", userUrl.String(), v)
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

func (u *UserService) NewFolderACL(f *Folder, r Role) ACL {
	return ACL{
		Role:     r,
		On:       f.Id,
		AccessOn: ACCESS_FOLDER,
	}
}

func (u *UserService) NewCustomerACL(r Role) ACL {
	return ACL{
		Role:     r,
		On:       u.client.Config.CustomerId,
		AccessOn: ACCESS_CUSTOMER,
	}
}

func (u *UserService) NewLocationACL(l *Location, r Role) ACL {
	return ACL{
		Role:     r,
		On:       l.GetId(),
		AccessOn: ACCESS_LOCATION,
	}
}
