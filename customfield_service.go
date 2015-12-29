package yext

import (
	"errors"
	"fmt"
)

const customFieldPath = "customFields"

type CustomFieldCache struct {
	CustomFields []*CustomField
}

func (c *CustomFieldCache) CustomField(name string) (*CustomField, error) {
	for _, cf := range c.CustomFields {
		if name == cf.Name {
			return cf, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Unable to find custom field with name %s", name))
}

func (c *CustomFieldCache) MustCustomField(name string) *CustomField {
	if cf, err := c.CustomField(name); err != nil {
		panic(err)
	} else {
		return cf
	}
}

func (c *CustomFieldCache) CustomFieldId(name string) (string, error) {
	if cf, err := c.CustomField(name); err != nil {
		return "", err
	} else {
		return cf.Id, nil
	}
}

func (c *CustomFieldCache) MustCustomFieldId(name string) string {
	if id, err := c.CustomFieldId(name); err != nil {
		panic(err)
	} else {
		return id
	}
}

func (c *CustomFieldCache) CustomFieldOptionId(fieldName, optionName string) (string, error) {
	cf, err := c.CustomField(fieldName)
	if err != nil {
		return "", err
	}

	if cf.Options == nil {
		return "", errors.New(fmt.Sprintf("Custom field %s doesn't have any options", fieldName))
	}

	for id, name := range cf.Options {
		if name == optionName {
			return id, nil
		}
	}

	return "", errors.New(fmt.Sprintf("Unable to find custom field option with name %s", optionName))
}

func (c *CustomFieldCache) MustCustomFieldOptionId(fieldName, optionName string) string {
	if id, err := c.CustomFieldOptionId(fieldName, optionName); err != nil {
		panic(err)
	} else {
		return id
	}
}

type CustomFieldService struct {
	CustomFieldCache *CustomFieldCache
	client           *Client
}

type customFieldResponse struct {
	CustomFields []*CustomField `json:"customFields"`
}

func (s *CustomFieldService) List() ([]*CustomField, error) {
	v := &customFieldResponse{}
	err := s.client.DoRequest("GET", customFieldPath, v)
	return v.CustomFields, err
}

func (s *CustomFieldService) CacheCustomFields() error {
	cfs, err := s.List()
	if err != nil {
		return err
	}

	s.CustomFieldCache = &CustomFieldCache{CustomFields: cfs}
	return nil
}

func (s *CustomFieldService) MustCacheCustomFields() {
	if err := s.CacheCustomFields(); err != nil {
		panic(err)
	}
}
