package yext

import (
	"encoding/json"
	"fmt"
)

const customFieldPath = "customfields"

var CustomFieldListMaxLimit = 1000

type CustomFieldService struct {
	CustomFieldManager *CustomFieldManager // TODO: do we need this?
	client             *Client
}

type CustomFieldResponse struct {
	Count        int            `json:"count"`
	CustomFields []*CustomField `json:"customFields"`
}

func (s *CustomFieldService) ListAll() ([]*CustomField, error) {
	var customFields []*CustomField
	var lr listRetriever = func(opts *ListOptions) (int, int, error) {
		cfr, _, err := s.List(opts)
		if err != nil {
			return 0, 0, err
		}
		customFields = append(customFields, cfr.CustomFields...)
		return len(cfr.CustomFields), cfr.Count, err
	}

	if err := listHelper(lr, &ListOptions{Limit: CustomFieldListMaxLimit}); err != nil {
		return nil, err
	} else {
		return customFields, nil
	}
}

func (s *CustomFieldService) List(opts *ListOptions) (*CustomFieldResponse, *Response, error) {
	requrl, err := addListOptions(customFieldPath, opts)
	if err != nil {
		return nil, nil, err
	}

	v := &CustomFieldResponse{}
	r, err := s.client.DoRequest("GET", requrl, v)
	if err != nil {
		return nil, r, err
	}
	return v, r, nil
}

func (s *CustomFieldService) Create(cf *CustomField) (*Response, error) {
	asJSON, err := json.Marshal(cf)
	if err != nil {
		return nil, err
	}
	var asMap map[string]interface{}
	err = json.Unmarshal(asJSON, &asMap)
	if err != nil {
		return nil, err
	}
	delete(asMap, "id")
	return s.client.DoRequestJSON("POST", customFieldPath, asMap, nil)
}

func (s *CustomFieldService) Edit(cf *CustomField) (*Response, error) {
	asJSON, err := json.Marshal(cf)
	if err != nil {
		return nil, err
	}
	var asMap map[string]interface{}
	err = json.Unmarshal(asJSON, &asMap)
	if err != nil {
		return nil, err
	}
	delete(asMap, "id")
	delete(asMap, "type")
	return s.client.DoRequestJSON("PUT", fmt.Sprintf("%s/%s", customFieldPath, cf.GetId()), asMap, nil)
}

func (c *CustomFieldManager) MustSetString(name string, value string, loc *Location) {
	Must(c.SetString(name, value, loc))
}

func (c *CustomFieldManager) SetPhoto(name string, v *CustomLocationPhoto, loc *Location) error {
	_, err := c.Set(name, v, loc)
	return err
}

func (c *CustomFieldManager) UnsetPhoto(name string, loc *Location) error {
	return c.SetPhoto(name, UnsetPhotoValue, loc)
}

func (c *CustomFieldManager) MustSetPhoto(name string, v *CustomLocationPhoto, loc *Location) {
	Must(c.SetPhoto(name, v, loc))
}

func (c *CustomFieldManager) MustUnsetPhoto(name string, loc *Location) {
	Must(c.SetPhoto(name, UnsetPhotoValue, loc))
}

func GetSingleOptionPointer(option SingleOption) *SingleOption {
	return &option
}
