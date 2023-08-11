package yext

import (
	"fmt"
)

const (
	ConfigResourceNamesPath            = "config/resourcenames/km"
	ConfigResourcesPath                = "config/resources/km"
	ConfigFieldSubType                 = "field"
	ConfigFieldEligibilityGroupSubType = "field-eligibility-group"
)

// ConfigFieldService utilises the config API to retrieve & edit information about custom fields using the same
// structs as the custom field service. This is so that it can be used as a drop-in replacement for the custom field
// service once it is deprecated. See CR-3126 for more information.
type ConfigFieldService struct {
	client *Client
}

func (s *ConfigFieldService) ListAll() ([]*CustomField, error) {
	var (
		customFields                  []*CustomField
		fieldEligibilityGroupsByField = make(map[string][]EntityType)
		fieldEligibilityGroups        []string
		fields                        []string
	)

	_, err := s.client.DoRequest("GET", fmt.Sprintf("%s/%s", ConfigResourceNamesPath, ConfigFieldEligibilityGroupSubType), &fieldEligibilityGroups)
	if err != nil {
		return nil, err
	}

	for _, fieldEligibilityGroup := range fieldEligibilityGroups {
		v := &ConfigFieldEligibilityGroup{}
		_, err := s.client.DoRequest("GET", fmt.Sprintf("%s/%s/%s", ConfigResourcesPath, ConfigFieldEligibilityGroupSubType, fieldEligibilityGroup), v)
		if err != nil {
			return nil, err
		}

		for _, field := range v.Fields {
			fieldEligibilityGroupsByField[field] = append(fieldEligibilityGroupsByField[field], v.EntityType)
		}
	}

	_, err = s.client.DoRequest("GET", fmt.Sprintf("%s/%s", ConfigResourceNamesPath, ConfigFieldSubType), &fields)
	if err != nil {
		return nil, err
	}

	for _, field := range fields {
		v := &ConfigField{}
		_, err := s.client.DoRequest("GET", fmt.Sprintf("%s/%s/%s", ConfigResourcesPath, ConfigFieldSubType, field), v)
		if err != nil {
			return nil, err
		}
		customFields = append(customFields, v.TransformToCustomField(fieldEligibilityGroupsByField[field]))
	}
	return customFields, nil
}

func (s *ConfigFieldService) Create(cf *CustomField) (*Response, error) {
	r, err := s.client.DoRequestJSON("POST", fmt.Sprintf("%s/%s", ConfigResourcesPath, ConfigFieldSubType), cf.TransformToConfigField(), nil)
	if err != nil {
		return r, err
	}

	for _, entityType := range cf.EntityAvailability {
		v := &ConfigFieldEligibilityGroup{}
		r, err := s.client.DoRequest("GET", fmt.Sprintf("%s/%s/%s.default", ConfigResourcesPath, ConfigFieldEligibilityGroupSubType, entityType), v)
		if err != nil {
			return r, err
		}

		v.Fields = append(v.Fields, GetString(cf.Id))
		r, err = s.client.DoRequestJSON("PATCH", fmt.Sprintf("%s/%s/%s.default", ConfigResourcesPath, ConfigFieldEligibilityGroupSubType, entityType), v, nil)
		if err != nil {
			return r, err
		}
	}

	return r, nil
}

func (s *ConfigFieldService) Edit(cf *CustomField) (*Response, error) {
	updatedDefinition := cf.TransformToConfigField()
	if cf.Type == CUSTOMFIELDTYPE_SINGLEOPTION || cf.Type == CUSTOMFIELDTYPE_MULTIOPTION {
		// get current field definition since we don't want to overwrite any validation settings
		existingDefinition := &ConfigField{}
		r, err := s.client.DoRequest("GET", fmt.Sprintf("%s/%s/%s", ConfigResourcesPath, ConfigFieldSubType, GetString(cf.Id)), existingDefinition)
		if err != nil {
			return r, err
		}

		updatedDefinition.Type = existingDefinition.Type
		if cf.Type == CUSTOMFIELDTYPE_SINGLEOPTION {
			updatedDefinition.Type.OptionType.Option = cf.GetConfigOptions()
		} else {
			updatedDefinition.Type.ListType.Type.ConfigType.OptionType.Option = cf.GetConfigOptions()
		}
	} else {
		updatedDefinition.Type = nil
	}

	r, err := s.client.DoRequestJSON("PATCH", fmt.Sprintf("%s/%s/%s", ConfigResourcesPath, ConfigFieldSubType, GetString(cf.Id)), cf.TransformToConfigField(), nil)
	if err != nil {
		return r, err
	}

	for _, entityType := range cf.EntityAvailability {
		v := &ConfigFieldEligibilityGroup{}
		r, err := s.client.DoRequest("GET", fmt.Sprintf("%s/%s/%s.default", ConfigResourcesPath, ConfigFieldEligibilityGroupSubType, entityType), v)
		if err != nil {
			return r, err
		}

		alreadyInFieldEligibilityGroup := false
		for _, field := range v.Fields {
			if field == GetString(cf.Id) {
				alreadyInFieldEligibilityGroup = true
				break
			}
		}
		if !alreadyInFieldEligibilityGroup {
			v.Fields = append(v.Fields, GetString(cf.Id))
			r, err = s.client.DoRequestJSON("PUT", fmt.Sprintf("%s/%s/%s.default", ConfigResourcesPath, ConfigFieldEligibilityGroupSubType, entityType), v, nil)
			if err != nil {
				return r, err
			}
		}
	}

	return r, nil
}
