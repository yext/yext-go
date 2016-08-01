package yext

const categoryPath = "categories"

type CategoryService struct {
	client *Client
}

type categoryListResponse struct {
	Categories []*Category `json:"categories"`
}

func (s *CategoryService) List() ([]*Category, error) {
	v := &categoryListResponse{}
	err := s.client.DoRawRequest("GET", categoryPath, v)
	return v.Categories, err
}
