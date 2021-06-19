package yext

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type LocationCustomFieldService struct {
	CustomFieldManager *LocationCustomFieldManager
	client             *Client
}

func (s *LocationCustomFieldService) ListAll() ([]*CustomField, error) {
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

func (s *LocationCustomFieldService) List(opts *ListOptions) (*CustomFieldResponse, *Response, error) {
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

func (s *LocationCustomFieldService) Create(cf *CustomField) (*Response, error) {
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

func (s *LocationCustomFieldService) Edit(cf *CustomField) (*Response, error) {
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

func (s *LocationCustomFieldService) Delete(customFieldId string) (*Response, error) {
	return s.client.DoRequest("DELETE", fmt.Sprintf("%s/%s", customFieldPath, customFieldId), nil)
}

type LocationCustomFieldManager struct {
	CustomFields []*CustomField
}

func (c *LocationCustomFieldManager) Get(name string, loc *Location) (interface{}, error) {
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

	return loc.CustomFields[field.GetId()], nil
}

func (c *LocationCustomFieldManager) MustGet(name string, loc *Location) interface{} {
	if ret, err := c.Get(name, loc); err != nil {
		panic(err)
	} else {
		return ret
	}
}

func (c *LocationCustomFieldManager) IsOptionSet(fieldName string, optionName string, loc *Location) (bool, error) {
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

func (c *LocationCustomFieldManager) MustIsOptionSet(fieldName string, optionName string, loc *Location) bool {
	if set, err := c.IsOptionSet(fieldName, optionName, loc); err != nil {
		panic(err)
	} else {
		return set
	}
}

func (c *LocationCustomFieldManager) SetOption(fieldName string, optionName string, loc *Location) (*Location, error) {
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

func (c *LocationCustomFieldManager) MustSetOption(fieldName string, optionName string, loc *Location) *Location {
	if loc, err := c.SetOption(fieldName, optionName, loc); err != nil {
		panic(err)
	} else {
		return loc
	}
}

func (c *LocationCustomFieldManager) UnsetOption(fieldName string, optionName string, loc *Location) (*Location, error) {
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

func (c *LocationCustomFieldManager) MustUnsetOption(fieldName string, optionName string, loc *Location) *Location {
	if loc, err := c.UnsetOption(fieldName, optionName, loc); err != nil {
		panic(err)
	} else {
		return loc
	}
}

// TODO: Why does this return a location?
// TODO: Should we validate the the type we received matches the type of the field?  Probably.
func (c *LocationCustomFieldManager) Set(name string, value CustomFieldValue, loc *Location) (*Location, error) {
	field, err := c.CustomField(name)
	if err != nil {
		return loc, err
	}
	loc.CustomFields[field.GetId()] = value
	return loc, nil
}

func (c *LocationCustomFieldManager) MustSet(name string, value CustomFieldValue, loc *Location) *Location {
	if loc, err := c.Set(name, value, loc); err != nil {
		panic(err)
	} else {
		return loc
	}
}

func (c *LocationCustomFieldManager) CustomField(name string) (*CustomField, error) {
	names := []string{}
	for _, cf := range c.CustomFields {
		if name == cf.Name {
			return cf, nil
		}
		names = append(names, cf.Name)
	}

	return nil, fmt.Errorf("Unable to find custom field with name %s, available fields: %v", name, names)
}

func (c *LocationCustomFieldManager) MustCustomField(name string) *CustomField {
	if cf, err := c.CustomField(name); err != nil {
		panic(err)
	} else {
		return cf
	}
}

func (c *LocationCustomFieldManager) CustomFieldId(name string) (string, error) {
	if cf, err := c.CustomField(name); err != nil {
		return "", err
	} else {
		return cf.GetId(), nil
	}
}

func (c *LocationCustomFieldManager) MustCustomFieldId(name string) string {
	if id, err := c.CustomFieldId(name); err != nil {
		panic(err)
	} else {
		return id
	}
}

func (c *LocationCustomFieldManager) CustomFieldName(id string) (string, error) {
	ids := []string{}
	for _, cf := range c.CustomFields {
		if id == cf.GetId() {
			return cf.Name, nil
		}
		ids = append(ids, cf.GetId())
	}

	return "", fmt.Errorf("Unable to find custom field with Id %s, available Ids: %v", id, ids)
}

func (c *LocationCustomFieldManager) MustCustomFieldName(id string) string {
	if name, err := c.CustomFieldName(id); err != nil {
		panic(err)
	} else {
		return name
	}
}

func (c *LocationCustomFieldManager) CustomFieldOptionId(fieldName, optionName string) (string, error) {
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

func (c *LocationCustomFieldManager) MustCustomFieldOptionId(fieldName, optionName string) string {
	if id, err := c.CustomFieldOptionId(fieldName, optionName); err != nil {
		panic(err)
	} else {
		return id
	}
}

func (c *LocationCustomFieldService) CacheCustomFields() ([]*CustomField, error) {
	cfs, err := c.ListAll()
	if err != nil {
		return nil, err
	}

	c.CustomFieldManager = &LocationCustomFieldManager{CustomFields: cfs}
	return c.CustomFieldManager.CustomFields, nil
}

func (c *LocationCustomFieldService) MustCacheCustomFields() []*CustomField {
	slice, err := c.CacheCustomFields()
	if err != nil {
		panic(err)
	}
	return slice
}

func ParseCustomFields(cfraw map[string]interface{}, cfs []*CustomField) (map[string]interface{}, error) {
	typefor := func(id string) string {
		for _, cf := range cfs {
			if cf.GetId() == id {
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
				newval = GetSingleOptionPointer(SingleOption(typedVal))
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
			asJSON, err := json.Marshal(v)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not re-marshal '%v' as json for Text List Field %v", v, err)
			}
			var cf TextList
			err = json.Unmarshal(asJSON, &cf)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not unmarshal '%v' into Text List Field %v", v, err)
			}
			newval = cf
		case CUSTOMFIELDTYPE_MULTIOPTION:
			asJSON, err := json.Marshal(v)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not re-marshal '%v' as json for Multi-Option Field %v", v, err)
			}
			var cf MultiOption
			err = json.Unmarshal(asJSON, &cf)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not unmarshal '%v' into Multi-Option Field %v", v, err)
			}
			newval = cf
		case CUSTOMFIELDTYPE_PHOTO:
			asJSON, err := json.Marshal(v)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not re-marshal '%v' as json for Photo Field %v", v, err)
			}
			var cfp *CustomLocationPhoto
			err = json.Unmarshal(asJSON, &cfp)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not unmarshal '%v' into Photo Field %v", v, err)
			}
			newval = cfp
		case CUSTOMFIELDTYPE_GALLERY:
			asJSON, err := json.Marshal(v)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not re-marshal '%v' as json for Photo Gallery Field %v", v, err)
			}
			var g CustomLocationGallery
			err = json.Unmarshal(asJSON, &g)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not unmarshal '%v' into Photo Gallery Field %v", v, err)
			}
			newval = g
		case CUSTOMFIELDTYPE_VIDEO:
			asJSON, err := json.Marshal(v)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not re-marshal '%v' as json for Video Field %v", v, err)
			}
			var cf CustomLocationVideo
			err = json.Unmarshal(asJSON, &cf)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not unmarshal '%v' into Custom Location Video Field %v", v, err)
			}
			newval = cf
		case CUSTOMFIELDTYPE_HOURS:
			asJSON, err := json.Marshal(v)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not re-marshal '%v' as json for Hours Field %v", v, err)
			}
			var cf CustomLocationHours
			err = json.Unmarshal(asJSON, &cf)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not unmarshal '%v' into Hours Field %v", v, err)
			}
			newval = cf
		case CUSTOMFIELDTYPE_DAILYTIMES:
			asJSON, err := json.Marshal(v)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not re-marshal '%v' as json for DailyT imes Field %v", v, err)
			}
			var cf CustomLocationDailyTimes
			err = json.Unmarshal(asJSON, &cf)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not unmarshal '%v' into Daily Times Field %v", v, err)
			}
			newval = cf
		case CUSTOMFIELDTYPE_LOCATIONLIST:
			asJSON, err := json.Marshal(v)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not re-marshal '%v' as json for Location List Field %v", v, err)
			}
			var cf LocationList
			err = json.Unmarshal(asJSON, &cf)
			if err != nil {
				return nil, fmt.Errorf("parse custom fields failure: could not unmarshal '%v' into Location List Field %v", v, err)
			}
			newval = cf
		default:
			newval = v
		}

		parsed[k] = newval
	}

	return parsed, nil
}

// validateLocationCustomFieldsKeys can be used with Location API to validate custom fields
func validateLocationCustomFieldsKeys(cfs map[string]interface{}) error {
	for k, _ := range cfs {
		if !customFieldKeyRegex.MatchString(k) {
			return errors.New(fmt.Sprintf("custom fields must be specified by their id, not name: %s", k))
		}
	}
	return nil
}

func (c *LocationCustomFieldManager) GetBool(name string, loc *Location) (bool, error) {
	value, err := c.Get(name, loc)
	if err != nil {
		return false, err
	}

	if value == nil {
		return false, nil
	}

	switch t := value.(type) {
	case YesNo:
		return bool(value.(YesNo)), nil
	case *YesNo:
		return bool(*value.(*YesNo)), nil
	default:
		return false, fmt.Errorf("GetBool failure: Field '%v' is not of a YesNo type", t)
	}
}

func (c *LocationCustomFieldManager) MustGetBool(name string, loc *Location) bool {
	if ret, err := c.GetBool(name, loc); err != nil {
		panic(err)
	} else {
		return ret
	}
}

// GetStringAliasCustomField returns the string value from a string type alias
// custom field. It will return an error if the field is not a string type.
func (c *LocationCustomFieldManager) GetString(name string, loc *Location) (string, error) {
	fv, err := c.Get(name, loc)
	if err != nil {
		return "", err
	}
	if fv == nil {
		return "", nil
	}
	switch fv.(type) {
	case SingleLineText:
		return string(fv.(SingleLineText)), nil
	case *SingleLineText:
		return string(*fv.(*SingleLineText)), nil
	case MultiLineText:
		return string(fv.(MultiLineText)), nil
	case *MultiLineText:
		return string(*fv.(*MultiLineText)), nil
	case Url:
		return string(fv.(Url)), nil
	case *Url:
		return string(*fv.(*Url)), nil
	case Date:
		return string(fv.(Date)), nil
	case *Date:
		return string(*fv.(*Date)), nil
	case Number:
		return string(fv.(Number)), nil
	case *Number:
		return string(*fv.(*Number)), nil
	case SingleOption:
		if string(fv.(SingleOption)) == "" {
			return "", nil
		}
		return c.CustomFieldOptionName(name, string(fv.(SingleOption)))
	case *SingleOption:
		if string(*fv.(*SingleOption)) == "" {
			return "", nil
		}
		return c.CustomFieldOptionName(name, string(*fv.(*SingleOption)))
	default:
		return "", fmt.Errorf("%s is not a string custom field type, is %T", name, fv)
	}
}

func (c *LocationCustomFieldManager) CustomFieldOptionName(cfName string, optionId string) (string, error) {
	cf, err := c.CustomField(cfName)
	if err != nil {
		return "", err
	}
	for _, option := range cf.Options {
		if option.Key == optionId {
			return option.Value, nil
		}
	}
	return "", fmt.Errorf("Unable to find option for key %s for custom field %s", optionId, cfName)
}

func (c *LocationCustomFieldManager) MustCustomFieldOptionName(fieldName, optionId string) string {
	if id, err := c.CustomFieldOptionName(fieldName, optionId); err != nil {
		panic(err)
	} else {
		return id
	}
}

func (c *LocationCustomFieldManager) MustGetString(name string, loc *Location) string {
	if ret, err := c.GetString(name, loc); err != nil {
		panic(err)
	} else {
		return ret
	}
}

// GetStringArrayAliasCustomField returns the string array value from a string array
// type alias custom field. It will return an error if the field is not a string
// array type.
func (c *LocationCustomFieldManager) GetStringSlice(name string, loc *Location) ([]string, error) {
	fv, err := c.Get(name, loc)
	if err != nil {
		return nil, err
	}
	if fv == nil {
		return nil, nil
	}
	switch fv.(type) {
	case UnorderedStrings:
		return []string(fv.(UnorderedStrings)), nil
	case *UnorderedStrings:
		return []string(*fv.(*UnorderedStrings)), nil
	case LocationList:
		return []string(fv.(LocationList)), nil
	case *LocationList:
		return []string(*fv.(*LocationList)), nil
	case TextList:
		return []string(fv.(TextList)), nil
	case *TextList:
		return []string(*fv.(*TextList)), nil
	case MultiOption:
		return c.CustomFieldOptionNames(name, []string(fv.(MultiOption)))
	case *MultiOption:
		return c.CustomFieldOptionNames(name, []string(*fv.(*MultiOption)))
	default:
		return nil, fmt.Errorf("%s is not a string array custom field type, is %T", name, fv)
	}
}

func (c *LocationCustomFieldManager) CustomFieldOptionNames(cfName string, optionIds []string) ([]string, error) {
	var optionNames = []string{}
	for _, optionId := range optionIds {
		optionName, err := c.CustomFieldOptionName(cfName, optionId)
		if err != nil {
			return nil, err
		}
		optionNames = append(optionNames, optionName)
	}
	return optionNames, nil
}

func (c *LocationCustomFieldManager) MustGetStringSlice(name string, loc *Location) []string {
	if ret, err := c.GetStringSlice(name, loc); err != nil {
		panic(err)
	} else {
		return ret
	}
}

func (c *LocationCustomFieldManager) SetBool(name string, value bool, loc *Location) error {
	field, err := c.CustomField(name)
	if err != nil {
		return err
	}

	if field.Type != CUSTOMFIELDTYPE_YESNO {
		return fmt.Errorf("SetBool failure: custom field '%v' is of type '%v' and not boolean", name, field.Type)
	}

	loc.CustomFields[field.GetId()] = YesNo(value)
	return nil
}

func (c *LocationCustomFieldManager) MustSetBool(name string, value bool, loc *Location) {
	if err := c.SetBool(name, value, loc); err != nil {
		panic(err)
	} else {
		return
	}
}

func (c *LocationCustomFieldManager) SetStringSlice(name string, value []string, loc *Location) error {
	field, err := c.CustomField(name)
	if err != nil {
		return err
	}

	switch field.Type {
	case CUSTOMFIELDTYPE_MULTIOPTION:
		for _, element := range value {
			c.MustSetOption(name, element, loc)
		}
		return nil
	case CUSTOMFIELDTYPE_TEXTLIST:
		loc.CustomFields[field.GetId()] = TextList(value)
		return nil
	case CUSTOMFIELDTYPE_LOCATIONLIST:
		loc.CustomFields[field.GetId()] = UnorderedStrings(value)
		return nil
	default:
		return fmt.Errorf("SetStringSlice failure: custom field '%v' is of type '%v' and can not take a string slice", name, field.Type)
	}
}

func (c *LocationCustomFieldManager) MustSetStringSlice(name string, value []string, loc *Location) {
	if err := c.SetStringSlice(name, value, loc); err != nil {
		panic(err)
	} else {
		return
	}
}

func (c *LocationCustomFieldManager) SetString(name string, value string, loc *Location) error {
	field, err := c.CustomField(name)
	if err != nil {
		return err
	}

	switch field.Type {
	case CUSTOMFIELDTYPE_SINGLEOPTION:
		c.MustSetOption(name, value, loc)
		return nil
	case CUSTOMFIELDTYPE_SINGLELINETEXT:
		loc.CustomFields[field.GetId()] = SingleLineText(value)
		return nil
	case CUSTOMFIELDTYPE_MULTILINETEXT:
		loc.CustomFields[field.GetId()] = MultiLineText(value)
		return nil
	case CUSTOMFIELDTYPE_URL:
		loc.CustomFields[field.GetId()] = Url(value)
		return nil
	case CUSTOMFIELDTYPE_DATE:
		loc.CustomFields[field.GetId()] = Date(value)
		return nil
	case CUSTOMFIELDTYPE_NUMBER:
		loc.CustomFields[field.GetId()] = Number(value)
		return nil
	default:
		return fmt.Errorf("SetString failure: custom field '%v' is of type '%v' and can not take a string", name, field.Type)
	}
}

func (c *LocationCustomFieldManager) MustSetString(name string, value string, loc *Location) {
	Must(c.SetString(name, value, loc))
}

func (c *LocationCustomFieldManager) SetPhoto(name string, v *CustomLocationPhoto, loc *Location) error {
	_, err := c.Set(name, v, loc)
	return err
}

func (c *LocationCustomFieldManager) UnsetPhoto(name string, loc *Location) error {
	return c.SetPhoto(name, UnsetPhotoValue, loc)
}

func (c *LocationCustomFieldManager) MustSetPhoto(name string, v *CustomLocationPhoto, loc *Location) {
	Must(c.SetPhoto(name, v, loc))
}

func (c *LocationCustomFieldManager) MustUnsetPhoto(name string, loc *Location) {
	Must(c.SetPhoto(name, UnsetPhotoValue, loc))
}

func GetSingleOptionPointer(option SingleOption) *SingleOption {
	return &option
}
