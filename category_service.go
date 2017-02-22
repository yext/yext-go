package yext

import "net/url"

const categoryPath = "categories"

// Category is a representation of a Category in Yext Location Manager.
// For details see http://developer.yext.com/docs/api-reference/#operation/getBusinessCategories
type Category struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	FullName   string `json:"fullName"`
	Selectable bool   `json:"selectable"`
	ParentId   string `json:"parentId"`
}

type CategoryService struct {
	client *Client
}

type CategoryListOptions struct {
	Language *string
	Country  *string
}

func (s *CategoryService) List(opts *CategoryListOptions) ([]*Category, error) {
	u, err := url.Parse(categoryPath)
	if err != nil {
		return nil, err
	}

	if opts != nil {
		q := u.Query()
		if opts.Language != nil {
			q.Add("language", *opts.Language)
		}
		if opts.Country != nil {
			q.Add("country", *opts.Country)
		}
		u.RawQuery = q.Encode()
	}

	v := []*Category{}
	_, err = s.client.DoRootRequest("GET", u.String(), v)
	return v, err
}
