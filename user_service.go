package yext

import "fmt"

const (
	userPath = "users"
	rolePath = "roles"
)

var UserListMaxLimit = 50

type UserService struct {
	client *Client
}

type UserListResponse struct {
	Count int     `json:"count"`
	Users []*User `json:"users"`
}

type RolesListResponse struct {
	Count int     `json:"count"`
	Roles []*Role `json:"roles"`
}

func (u *UserService) ListAll() ([]*User, error) {
	var users []*User
	var lr listRetriever = func(opts *ListOptions) (int, int, error) {
		ulr, _, err := u.List(opts)
		if err != nil {
			return 0, 0, err
		}
		users = append(users, ulr.Users...)
		return len(ulr.Users), ulr.Count, err
	}

	if err := listHelper(lr, &ListOptions{Limit: UserListMaxLimit}); err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func (u *UserService) List(opts *ListOptions) (*UserListResponse, *Response, error) {
	requrl, err := addListOptions(userPath, opts)
	if err != nil {
		return nil, nil, err
	}

	v := &UserListResponse{}
	r, err := u.client.DoRequest("GET", requrl, v)
	if err != nil {
		return nil, r, err
	}
	return v, r, nil
}

func (u *User) pathToUser() string {
	return pathToUserId(u.GetId())
}

func pathToUserId(id string) string {
	return fmt.Sprintf("%s/%s", userPath, id)
}

func (u *UserService) Get(id string) (*User, *Response, error) {
	var v = &User{}
	r, err := u.client.DoRequest("GET", pathToUserId(id), v)
	if err != nil {
		return nil, r, err
	}
	return v, r, nil
}

func (u *UserService) Edit(y *User) (*Response, error) {
	return u.client.DoRequestJSON("PUT", y.pathToUser(), y, nil)
}

func (u *UserService) Create(y *User) (*Response, error) {
	return u.client.DoRequestJSON("POST", userPath, y, nil)
}

func (u *UserService) Delete(y *User) (*Response, error) {
	return u.client.DoRequest("DELETE", y.pathToUser(), nil)
}

func (u *UserService) ListRoles() (*RolesListResponse, *Response, error) {
	v := &RolesListResponse{}
	r, err := u.client.DoRequest("GET", rolePath, v)
	if err != nil {
		return nil, r, err
	}
	return v, r, nil
}

func (u *UserService) NewFolderACL(f *Folder, r Role) ACL {
	return ACL{
		Role:      r,
		On:        f.Id,
		AccountId: u.client.Config.AccountId,
		AccessOn:  ACCESS_FOLDER,
	}
}

func (u *UserService) NewAccountACL(r Role) ACL {
	return ACL{
		Role:      r,
		On:        u.client.Config.AccountId,
		AccountId: u.client.Config.AccountId,
		AccessOn:  ACCESS_ACCOUNT,
	}
}

func (u *UserService) NewLocationACL(l *Location, r Role) ACL {
	return ACL{
		Role:      r,
		On:        l.GetId(),
		AccountId: u.client.Config.AccountId,
		AccessOn:  ACCESS_LOCATION,
	}
}

func (u *UserService) NewEntityACL(e Entity, r Role) ACL {
    return ACL{
        Role:      r,
        On:        e.GetEntityId(),
        AccountId: u.client.Config.AccountId,
        AccessOn:  ACCESS_LOCATION,
    }
}
