package yext

import (
	"encoding/json"
	"fmt"
	"strconv"
)

const customFieldPath = "customfields"

var CustomFieldListMaxLimit = 1000

type CustomFieldManager struct {
	CustomFields []*CustomField
}

func (s *CustomFieldService) Edit(cf *CustomField) (*Response, error) {
	asJson, err := json.Marshal(cf)
	if err != nil {
		return nil, err
	}
	var asMap map[string]interface{}
	err = json.Unmarshal(asJson, &asMap)
	if err != nil {
		return nil, err
	}
	delete(asMap, "id")
	delete(asMap, "type")
	r, err := s.client.DoRequestJSON("PUT", fmt.Sprintf("%s/%s", customFieldPath, cf.Id), asMap, nil)
	if err != nil {
		return r, err
	}
	return r, nil
}

func (c *CustomFieldManager) Get(name string, loc *Location) (interface{}, error) {
	if loc == nil || loc.CustomFields == nil {
		return nil, nil
	}

	var (
		field *CustomField
		err   error
	)

	if field, err = c.CustomField(name); err != nil {
		return nil, err
	}

	return loc.CustomFields[field.Id], nil
}

func (c *CustomFieldManager) MustGet(name string, loc *Location) interface{} {
	if ret, err := c.Get(name, loc); err != nil {
		panic(err)
	} else {
		return ret
	}
}

func (c *CustomFieldManager) IsOptionSet(fieldName string, optionName string, loc *Location) (bool, error) {
	var (
		field interface{}
		err   error
		of    OptionField
		id    string
	)

	if field, err = c.Get(fieldName, loc); err != nil {
		return false, err
	}

	if field == nil {
		return false, nil
	}
	switch field.(type) {
	case nil:
		return false, nil
	case MultiOption:
		mo := field.(MultiOption)
		of = &mo
	case *MultiOption:
		of = field.(*MultiOption)
	case SingleOption:
		so := field.(SingleOption)
		of = &so
	case *SingleOption:
		of = field.(*SingleOption)
	default:
		return false, fmt.Errorf("'%s' is not an OptionField custom field, is %T", fieldName, field)
	}

	if id, err = c.CustomFieldOptionId(fieldName, optionName); err != nil {
		return false, err
	}

	return of.IsOptionIdSet(id), nil
}

func (c *CustomFieldManager) MustIsOptionSet(fieldName string, optionName string, loc *Location) bool {
	if set, err := c.IsOptionSet(fieldName, optionName, loc); err != nil {
		panic(err)
	} else {
		return set
	}
}

func (c *CustomFieldManager) SetOption(fieldName string, optionName string, loc *Location) (*Location, error) {
	var (
		field interface{}
		err   error
		of    OptionField
		ok    bool
		id    string
	)

	if field, err = c.Get(fieldName, loc); err != nil {
		return loc, err
	}

	if field == nil {
		var cf *CustomField
		if cf, err = c.CustomField(fieldName); err != nil {
			return loc, fmt.Errorf("problem getting '%s': %v", fieldName, err)
		}
		switch cf.Type {
		case CUSTOMFIELDTYPE_MULTIOPTION:
			of = new(MultiOption)
		case CUSTOMFIELDTYPE_SINGLEOPTION:
			of = new(SingleOption)
		default:
			return loc, fmt.Errorf("'%s' is not an OptionField is '%s'", cf.Name, cf.Type)
		}
	} else if of, ok = field.(OptionField); !ok {
		return loc, fmt.Errorf("'%s': %v is not an OptionField custom field is %T", fieldName, field, field)
	}

	if id, err = c.CustomFieldOptionId(fieldName, optionName); err != nil {
		return loc, err
	}

	of.SetOptionId(id)
	return c.Set(fieldName, of, loc)
}

func (c *CustomFieldManager) MustSetOption(fieldName string, optionName string, loc *Location) *Location {
	if loc, err := c.SetOption(fieldName, optionName, loc); err != nil {
		panic(err)
	} else {
		return loc
	}
}

func (c *CustomFieldManager) UnsetOption(fieldName string, optionName string, loc *Location) (*Location, error) {
	var (
		field interface{}
		err   error
		id    string
	)

	if field, err = c.Get(fieldName, loc); err != nil {
		return loc, err
	}

	if field == nil {
		return loc, fmt.Errorf("'%s' is not currently set", fieldName)
	}

	option, ok := field.(OptionField)
	if !ok {
		return loc, fmt.Errorf("'%s' is not an OptionField custom field", fieldName)
	}

	if id, err = c.CustomFieldOptionId(fieldName, optionName); err != nil {
		return loc, err
	}

	option.UnsetOptionId(id)
	return c.Set(fieldName, option, loc)
}

func (c *CustomFieldManager) MustUnsetOption(fieldName string, optionName string, loc *Location) *Location {
	if loc, err := c.UnsetOption(fieldName, optionName, loc); err != nil {
		panic(err)
	} else {
		return loc
	}
}

func (c *CustomFieldManager) Set(name string, value CustomFieldValue, loc *Location) (*Location, error) {
	field, err := c.CustomField(name)
	if err != nil {
		return loc, err
	}
	loc.CustomFields[field.Id] = value
	return loc, nil
}

func (c *CustomFieldManager) MustSet(name string, value CustomFieldValue, loc *Location) *Location {
	if loc, err := c.Set(name, value, loc); err != nil {
		panic(err)
	} else {
		return loc
	}
}

func (c *CustomFieldManager) CustomField(name string) (*CustomField, error) {
	for _, cf := range c.CustomFields {
		if name == cf.Name {
			return cf, nil
		}
	}

	return nil, fmt.Errorf("Unable to find custom field with name %s, available field: %v", name, c.CustomFields)
}

func (c *CustomFieldManager) MustCustomField(name string) *CustomField {
	if cf, err := c.CustomField(name); err != nil {
		panic(err)
	} else {
		return cf
	}
}

func (c *CustomFieldManager) CustomFieldId(name string) (string, error) {
	if cf, err := c.CustomField(name); err != nil {
		return "", err
	} else {
		return cf.Id, nil
	}
}

func (c *CustomFieldManager) MustCustomFieldId(name string) string {
	if id, err := c.CustomFieldId(name); err != nil {
		panic(err)
	} else {
		return id
	}
}

func (c *CustomFieldManager) CustomFieldOptionId(fieldName, optionName string) (string, error) {
	cf, err := c.CustomField(fieldName)
	if err != nil {
		return "", err
	}

	if cf.Options == nil {
		return "", fmt.Errorf("Custom field %s doesn't have any options", fieldName)
	}

	for _, option := range cf.Options {
		if option.Value == optionName {
			return option.Key, nil
		}
	}

	return "", fmt.Errorf("Unable to find custom field option with name %s", optionName)
}

func (c *CustomFieldManager) MustCustomFieldOptionId(fieldName, optionName string) string {
	if id, err := c.CustomFieldOptionId(fieldName, optionName); err != nil {
		panic(err)
	} else {
		return id
	}
}

type CustomFieldService struct {
	CustomFieldManager *CustomFieldManager
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

func (s *CustomFieldService) CacheCustomFields() error {
	cfs, err := s.ListAll()
	if err != nil {
		return err
	}

	s.CustomFieldManager = &CustomFieldManager{CustomFields: cfs}
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
		if _, ok := v.(CustomFieldValue); ok {
			parsed[k] = v
			continue
		}

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
			var cfp Photo
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
		case CUSTOMFIELDTYPE_DAILYTIMES:
			asJson, err := json.Marshal(v)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not re-marshal '%v' as json for DailyT imes Field %v", v, err)
			}
			var cf DailyTimes
			err = json.Unmarshal(asJson, &cf)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not unmarshal '%v' into Daily Times Field %v", v, err)
			}
			newval = cf
		default:
			newval = v
		}

		parsed[k] = newval
	}

	return parsed, nil
}
