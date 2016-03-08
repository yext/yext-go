package yext

import (
	"encoding/json"
	"fmt"
	"strconv"
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

	return nil, fmt.Errorf("Unable to find custom field with name %s", name)
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
		return "", fmt.Errorf("Custom field %s doesn't have any options", fieldName)
	}

	for id, name := range cf.Options {
		if name == optionName {
			return id, nil
		}
	}

	return "", fmt.Errorf("Unable to find custom field option with name %s", optionName)
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

func ParseCustomFields(cfraw map[string]interface{}, cfs []*CustomField) (map[string]interface{}, error) {
	typefor := func(id string) string {
		for _, cf := range cfs {
			if cf.Id == id {
				return cf.Type
			}
		}
		return ""
	}

	parsed := map[string]interface{}{}

	for k, v := range cfraw {
		var newval interface{}

		switch typefor(k) {
		case CUSTOMFIELDTYPE_YESNO:
			if typedVal, ok := v.(bool); ok {
				newval = YesNo(typedVal)
			} else if typedVal, ok := v.(string); ok {
				b, err := strconv.ParseBool(typedVal)
				if err != nil {
					return nil, fmt.Errorf("parse custom fields failure: could not parse '%v' as yes/no %v", v, err)
				}
				newval = YesNo(b)
			} else {
				return nil, fmt.Errorf("parse custom fields failure: could not parse '%v' as yes/no, expected bool got %T", v, v)
			}
		case CUSTOMFIELDTYPE_NUMBER:
			if typedVal, ok := v.(string); ok {
				newval = Number(typedVal)
			} else {
				return nil, fmt.Errorf("parse custom fields failure: could not parse '%v' as number, expected string got %T", v, v)
			}
		case CUSTOMFIELDTYPE_SINGLELINETEXT:
			if typedVal, ok := v.(string); ok {
				newval = SingleLineText(typedVal)
			} else {
				return nil, fmt.Errorf("parse custom fields failure: could not parse '%v' as single-line text, expected string got %T", v, v)
			}
		case CUSTOMFIELDTYPE_MULTILINETEXT:
			if typedVal, ok := v.(string); ok {
				newval = MultiLineText(typedVal)
			} else {
				return nil, fmt.Errorf("parse custom fields failure: could not parse '%v' as multi-line text, expected string got %T", v, v)
			}
		case CUSTOMFIELDTYPE_SINGLEOPTION:
			if typedVal, ok := v.(string); ok {
				newval = SingleOption(typedVal)
			} else {
				return nil, fmt.Errorf("parse custom fields failure: could not parse '%v' as single-option field, expected string got %T", v, v)
			}
		case CUSTOMFIELDTYPE_URL:
			if typedVal, ok := v.(string); ok {
				newval = Url(typedVal)
			} else {
				return nil, fmt.Errorf("parse custom fields failure: could not parse '%v' as url field, expected string got %T", v, v)
			}
		case CUSTOMFIELDTYPE_DATE:
			if typedVal, ok := v.(string); ok {
				newval = Date(typedVal)
			} else {
				return nil, fmt.Errorf("parse custom fields failure: could not parse '%v' as date field, expected string got %T", v, v)
			}
		case CUSTOMFIELDTYPE_TEXTLIST:
			asJson, err := json.Marshal(v)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not re-marshal '%v' as json for Text List Field %v", v, err)
			}
			var cf TextList
			err = json.Unmarshal(asJson, &cf)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not unmarshal '%v' into Text List Field %v", v, err)
			}
			newval = cf
		case CUSTOMFIELDTYPE_MULTIOPTION:
			asJson, err := json.Marshal(v)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not re-marshal '%v' as json for Multi-Option Field %v", v, err)
			}
			var cf MultiOption
			err = json.Unmarshal(asJson, &cf)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not unmarshal '%v' into Multi-Option Field %v", v, err)
			}
			newval = cf
		case CUSTOMFIELDTYPE_PHOTO:
			asJson, err := json.Marshal(v)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not re-marshal '%v' as json for Photo Field %v", v, err)
			}
			var cfp CustomPhoto
			err = json.Unmarshal(asJson, &cfp)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not unmarshal '%v' into Photo Field %v", v, err)
			}
			newval = cfp
		case CUSTOMFIELDTYPE_GALLERY:
			asJson, err := json.Marshal(v)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not re-marshal '%v' as json for Photo Gallery Field %v", v, err)
			}
			var g Gallery
			err = json.Unmarshal(asJson, &g)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not unmarshal '%v' into Photo Gallery Field %v", v, err)
			}
			newval = g
		case CUSTOMFIELDTYPE_VIDEO:
			asJson, err := json.Marshal(v)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not re-marshal '%v' as json for Video Field %v", v, err)
			}
			var cf Video
			err = json.Unmarshal(asJson, &cf)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not unmarshal '%v' into Video Field %v", v, err)
			}
			newval = cf
		case CUSTOMFIELDTYPE_HOURS:
			asJson, err := json.Marshal(v)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not re-marshal '%v' as json for Hours Field %v", v, err)
			}
			var cf Hours
			err = json.Unmarshal(asJson, &cf)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not unmarshal '%v' into Hours Field %v", v, err)
			}
			newval = cf
		default:
			newval = v
		}

		parsed[k] = newval
	}

	return parsed, nil
}
