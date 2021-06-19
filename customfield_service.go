package yext

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

const customFieldPath = "customfields"

var CustomFieldListMaxLimit = 1000

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
	}
	return customFields, nil

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

type CustomFieldManager struct {
	CustomFields []*CustomField
}

func (c *CustomFieldManager) CustomField(name string) (*CustomField, error) {
	names := []string{}
	for _, cf := range c.CustomFields {
		if name == cf.Name {
			return cf, nil
		}
		names = append(names, cf.Name)
	}

	return nil, fmt.Errorf("Unable to find custom field with name %s, available fields: %v", name, names)
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
		return cf.GetId(), nil
	}
}

func (c *CustomFieldManager) MustCustomFieldId(name string) string {
	if id, err := c.CustomFieldId(name); err != nil {
		panic(err)
	} else {
		return id
	}
}

func (c *CustomFieldManager) CustomFieldName(id string) (string, error) {
	ids := []string{}
	for _, cf := range c.CustomFields {
		if id == cf.GetId() {
			return cf.Name, nil
		}
		ids = append(ids, cf.GetId())
	}

	return "", fmt.Errorf("Unable to find custom field with Id %s, available Ids: %v", id, ids)
}

func (c *CustomFieldManager) MustCustomFieldName(id string) string {
	if name, err := c.CustomFieldName(id); err != nil {
		panic(err)
	} else {
		return name
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

func (c *CustomFieldManager) MustSingleOptionId(fieldName, optionName string) **string {
	if optionName == "" {
		return c.NullSingleOption()
	}
	id := c.MustCustomFieldOptionId(fieldName, optionName)
	return NullableString(id)
}

func (c *CustomFieldManager) MustIsSingleOptionSet(fieldName, optionName string, setOptionId **string) bool {
	id := c.MustCustomFieldOptionId(fieldName, optionName)
	return GetNullableString(setOptionId) == id
}

func (c *CustomFieldManager) NullSingleOption() **string {
	return NullString()
}

func (c *CustomFieldManager) MustMultiOptionIds(fieldName string, optionNames ...string) *UnorderedStrings {
	if len(optionNames) == 0 {
		return c.NullMultiOption()
	}
	var optionIds = []string{}
	for _, optionName := range optionNames {
		id := c.MustCustomFieldOptionId(fieldName, optionName)

		shouldAddOptionId := true
		for _, optionId := range optionIds {
			if id == optionId { // Don't add duplicate option ids
				shouldAddOptionId = false
				break
			}
		}

		if shouldAddOptionId {
			optionIds = append(optionIds, id)
		}
	}
	return ToUnorderedStrings(optionIds)
}

func (c *CustomFieldManager) MustIsMultiOptionSet(fieldName string, optionName string, setOptionIds *UnorderedStrings) bool {
	if setOptionIds == nil {
		return false
	}
	optionId := c.MustCustomFieldOptionId(fieldName, optionName)
	for _, v := range *setOptionIds {
		if v == optionId {
			return true
		}
	}
	return false
}

func (c *CustomFieldManager) NullMultiOption() *UnorderedStrings {
	return ToUnorderedStrings([]string{})
}

func (c *CustomFieldManager) GetMultiOptionNames(fieldName string, optionIds *UnorderedStrings) ([]string, error) {
	optionNames := []string{}
	for _, optionId := range *optionIds {
		optionName, err := c.CustomFieldOptionName(fieldName, optionId)
		if err != nil {
			return nil, err
		}
		optionNames = append(optionNames, optionName)
	}
	return optionNames, nil
}

// MustGetMultiOptionNames returns a slice containing the names of the options that are set
// Example Use:
// Custom field name: "Available Breeds"
// Custom field option ID to name map: {
//		"2421": "Labrador Retriever",
//		"2422": "Chihuahua",
//		"2423": "Boston Terrier"
// }
// Custom field options that are set: "2421" and "2422"
// cfmanager.MustGetMultiOptionNames("Available Breeds", loc.AvailableBreeds) returns ["Labrador Retriever", "Chihuahua"]
func (c *CustomFieldManager) MustGetMultiOptionNames(fieldName string, options *UnorderedStrings) []string {
	if ids, err := c.GetMultiOptionNames(fieldName, options); err != nil {
		panic(err)
	} else {
		return ids
	}
}

func (c *CustomFieldManager) GetSingleOptionName(fieldName string, optionId **string) (string, error) {
	if GetNullableString(optionId) == "" {
		return "", nil
	}

	optionName, err := c.CustomFieldOptionName(fieldName, GetNullableString(optionId))
	if err != nil {
		return "", err
	}
	return optionName, nil
}

// MustGetSingleOptionName returns the name of a set option. It returns an empty string if the single
// option is unset (i.e. if its value is a NullString)
// Example Use:
// Custom field name: "Dog Height"
// Custom field option ID to name map: {
//		"1423": "Tall",
//		"1424": "Short",
// }
// Custom field option that is set: "1424"
// cfmanager.MustGetSingleOptionName("Dog Height", loc.DogHeight) returns "Short"
//
func (c *CustomFieldManager) MustGetSingleOptionName(fieldName string, optionId **string) string {
	if id, err := c.GetSingleOptionName(fieldName, optionId); err != nil {
		panic(err)
	} else {
		return id
	}
}

func (c *CustomFieldManager) CustomFieldOptionName(cfName string, optionId string) (string, error) {
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

func (c *CustomFieldManager) MustCustomFieldOptionName(fieldName, optionId string) string {
	if id, err := c.CustomFieldOptionName(fieldName, optionId); err != nil {
		panic(err)
	} else {
		return id
	}
}

func (c *CustomFieldService) CacheCustomFields() ([]*CustomField, error) {
	cfs, err := c.ListAll()
	if err != nil {
		return nil, err
	}

	c.CustomFieldManager = &CustomFieldManager{CustomFields: cfs}
	return c.CustomFieldManager.CustomFields, nil
}

func (c *CustomFieldService) MustCacheCustomFields() []*CustomField {
	slice, err := c.CacheCustomFields()
	if err != nil {
		panic(err)
	}
	return slice
}

// SetCustomFieldValue sets the value of a given custom field (fieldName)
// A pointer to the custom entity (&CustomEntity) should be passed in as the customEntity interface{}
func (c *CustomFieldManager) SetCustomFieldValue(customEntity interface{}, fieldName string, valToSet interface{}) {
	cfId := c.MustCustomFieldId(fieldName)
	SetFieldByJSONTag(customEntity, cfId, valToSet)
}

func SetFieldByJSONTag(i interface{}, fieldTag string, valToSet interface{}) {
	var (
		valueOfValToSet = reflect.ValueOf(valToSet)
		v               = Indirect(reflect.ValueOf(i))
		t               = v.Type()
		num             = v.NumField()
	)

	for n := 0; n < num; n++ {
		var (
			field = t.Field(n)
			name  = field.Name
			tag   = strings.Replace(field.Tag.Get("json"), ",omitempty", "", -1)
			val   = v.Field(n)
		)
		if tag == fieldTag {
			Indirect(reflect.ValueOf(i)).FieldByName(name).Set(valueOfValToSet)
		} else {
			Indirect(reflect.ValueOf(i)).FieldByName(name).Set(val)
		}
	}
}

// GetCustomFieldValue gets the value of a given custom field (fieldName)
// A pointer to the custom entity (&CustomEntity) should be passed in as the customEntity interface{}
func (c *CustomFieldManager) GetCustomFieldValue(customEntity interface{}, fieldName string) interface{} {
	cfId := c.MustCustomFieldId(fieldName)
	return GetFieldByJSONTag(customEntity, cfId)
}

func GetFieldByJSONTag(i interface{}, fieldTag string) interface{} {
	var (
		v   = Indirect(reflect.ValueOf(i))
		t   = v.Type()
		num = v.NumField()
	)

	for n := 0; n < num; n++ {
		var (
			field = t.Field(n)
			name  = field.Name
			tag   = strings.Replace(field.Tag.Get("json"), ",omitempty", "", -1)
		)
		if tag == fieldTag {
			return Indirect(reflect.ValueOf(i)).FieldByName(name).Interface()
		}
	}
	return nil
}
