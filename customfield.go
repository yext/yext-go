package yext

import (
	"encoding/json"
)

const (
	CUSTOMFIELDTYPE_YESNO          = "BOOLEAN"
	CUSTOMFIELDTYPE_SINGLELINETEXT = "TEXT"
	CUSTOMFIELDTYPE_MULTILINETEXT  = "MULTILINE_TEXT"
	CUSTOMFIELDTYPE_SINGLEOPTION   = "SINGLE_OPTION"
	CUSTOMFIELDTYPE_URL            = "URL"
	CUSTOMFIELDTYPE_DATE           = "DATE"
	CUSTOMFIELDTYPE_NUMBER         = "NUMBER"
	CUSTOMFIELDTYPE_MULTIOPTION    = "MULTI_OPTION"
	CUSTOMFIELDTYPE_TEXTLIST       = "TEXT_LIST"
	CUSTOMFIELDTYPE_PHOTO          = "PHOTO"
	CUSTOMFIELDTYPE_GALLERY        = "GALLERY"
	CUSTOMFIELDTYPE_VIDEO          = "VIDEO"
	CUSTOMFIELDTYPE_HOURS          = "HOURS"
	CUSTOMFIELDTYPE_DAILYTIMES     = "DAILY_TIMES"
	CUSTOMFIELDTYPE_ENTITYLIST     = "ENTITY_LIST"
	CUSTOMFIELDTYPE_RICHTEXT       = "RICH_TEXT"
)

type EntityRelationship struct {
	SupportedDestinationEntityTypeIds []string `json:"supportedDestinationEntityTypeIds,omitempty"`
	Type                              string   `json:"type,omitempty"`
}

type CustomFieldValidation struct {
	MinCharLength      int                 `json:"minCharLength,omitempty"`
	MaxCharLength      int                 `json:"maxCharLength,omitempty"`
	MinItemCount       int                 `json:"minItemCount,omitempty"`
	MaxItemCount       int                 `json:"maxItemCount,omitempty"`
	MinValue           float64             `json:"minValue,omitempty"`
	MaxValue           float64             `json:"maxValue,omitempty"`
	MinDate            string              `json:"minDate,omitempty"`
	MaxDate            string              `json:"maxDate,omitempty"`
	AspectRatio        string              `json:"aspectRatio,omitempty"`
	MinWidth           int                 `json:"minWidth,omitempty"`
	MinHeight          int                 `json:"minHeight,omitempty"`
	EntityTypes        []string            `json:"entityTypes,omitempty"`
	RichTextFormats    []string            `json:"richTextFormats,omitempty"`
	EntityRelationship *EntityRelationship `json:"entityRelationship,omitempty"`
}

type CustomField struct {
	Id                         *string                `json:"id,omitempty"`
	Type                       string                 `json:"type"`
	Name                       string                 `json:"name"`
	Options                    []CustomFieldOption    `json:"options,omitempty"` // Only present for option custom fields
	Group                      string                 `json:"group"`
	Description                string                 `json:"description,omitempty"`
	AlternateLanguageBehaviour string                 `json:"alternateLanguageBehavior"`
	EntityAvailability         []EntityType           `json:"entityAvailability"`
	Validation                 *CustomFieldValidation `json:"validation,omitempty"` // Needed for rich text formatting
}

func (c *CustomField) UnmarshalJSON(data []byte) error {
	type Alias CustomField

	// name and description are strings on older API versions (v param), and objects
	// on the newer ones. For backwards compatibility, convert from object to string
	a := &struct {
		*Alias
		Name        json.RawMessage `json:"name"`
		Description json.RawMessage `json:"description"`
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	if a.Name != nil {
		// try to unmarshal as string first, fallback to object
		if err := json.Unmarshal(a.Name, &c.Name); err != nil {
			var nameValue struct {
				Value string `json:"value"`
			}

			if err := json.Unmarshal(a.Name, &nameValue); err != nil {
				return err
			}

			c.Name = nameValue.Value
		}
	}

	if a.Description != nil {
		// try to unmarshal as string first, fallback to object
		if err := json.Unmarshal(a.Description, &c.Description); err != nil {
			var val struct {
				Value string `json:"value"`
			}

			if err := json.Unmarshal(a.Description, &val); err != nil {
				return err
			}

			c.Description = val.Value
		}
	}
	return nil
}

func (c CustomField) GetId() string {
	if c.Id == nil {
		return ""
	}
	return *c.Id
}
