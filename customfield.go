package yext

import (
	"encoding/json"
	"strings"
)

const (
	CUSTOMFIELDTYPE_YESNO              = "BOOLEAN"
	CUSTOMFIELDTYPE_SINGLELINETEXT     = "TEXT"
	CUSTOMFIELDTYPE_MULTILINETEXT      = "MULTILINE_TEXT"
	CUSTOMFIELDTYPE_SINGLEOPTION       = "SINGLE_OPTION"
	CUSTOMFIELDTYPE_URL                = "URL"
	CUSTOMFIELDTYPE_DATE               = "DATE"
	CUSTOMFIELDTYPE_NUMBER             = "NUMBER"
	CUSTOMFIELDTYPE_MULTIOPTION        = "MULTI_OPTION"
	CUSTOMFIELDTYPE_TEXTLIST           = "TEXT_LIST"
	CUSTOMFIELDTYPE_SIMPLEPHOTO        = "SIMPLE_PHOTO"
	CUSTOMFIELDTYPE_PHOTO              = "PHOTO"
	CUSTOMFIELDTYPE_GALLERY            = "GALLERY"
	CUSTOMFIELDTYPE_SIMPLEVIDEO        = "SIMPLE_VIDEO"
	CUSTOMFIELDTYPE_VIDEO              = "VIDEO"
	CUSTOMFIELDTYPE_VIDEOGALLERY       = "VIDEO_GALLERY"
	CUSTOMFIELDTYPE_HOURS              = "HOURS"
	CUSTOMFIELDTYPE_DAILYTIMES         = "DAILY_TIMES"
	CUSTOMFIELDTYPE_ENTITYLIST         = "ENTITY_LIST"
	CUSTOMFIELDTYPE_RICHTEXT           = "RICH_TEXT"
	CUSTOMFIELDTYPE_RICHTEXTV2         = "RICH_TEXT_V2"
	CUSTOMFIELDTYPE_SIMPLYVIDEOGALLERY = "SIMPLE_VIDEO_GALLERY"
	CUSTOMFIELDTYPE_SIMPLEPHOTOGALLERY = "SIMPLE_PHOTO_GALLERY"
	CUSTOMFIELDTYPE_MARKDOWN           = "MARKDOWN"
	CUSTOMFIELDTYPE_SLUG               = "SLUG"
	CUSTOMFIELDTYPE_CALLTOACTION       = "CALL_TO_ACTION"
	CUSTOMFIELDTYPE_PRICE              = "PRICE"
	CUSTOMFIELDTYPE_LIST               = "LIST"
	CUSTOMFIELDTYPE_STRUCT             = "STRUCT"
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

func (c *CustomField) TransformToConfigField() *ConfigField {
	r := &ConfigField{
		Id:           c.Id,
		Schema:       "https://schema.yext.com/config/km/field/v1",
		Description:  c.Description,
		DisplayName:  c.Name,
		Group:        strings.Replace(c.Group, "_", "", -1),
		Localization: c.AlternateLanguageBehaviour,
		Type:         c.GetConfigType(),
		TypeId:       c.GetConfigFieldTypeId(),
	}

	if c.AlternateLanguageBehaviour == "LANGUAGE_SPECIFIC" {
		r.Localization = "LOCALE_SPECIFIC"
	}

	return r
}

func (c *CustomField) GetConfigFieldTypeId() string {
	switch c.Type {
	case CUSTOMFIELDTYPE_YESNO:
		return CONFIGFIELDTYPE_YESNO
	case CUSTOMFIELDTYPE_NUMBER:
		return CONFIGFIELDTYPE_DECIMAL
	case CUSTOMFIELDTYPE_DATE:
		return CONFIGFIELDTYPE_DATE
	case CUSTOMFIELDTYPE_DAILYTIMES:
		return CONFIGFIELDTYPE_DAILYTIMES
	case CUSTOMFIELDTYPE_HOURS:
		return CONFIGFIELDTYPE_HOURS
	case CUSTOMFIELDTYPE_SINGLEOPTION:
		return CONFIGFIELDTYPE_OPTION
	case CUSTOMFIELDTYPE_RICHTEXT:
		return CONFIGFIELDTYPE_RICHTEXT
	case CUSTOMFIELDTYPE_RICHTEXTV2:
		return CONFIGFIELDTYPE_RICHTEXTV2
	case CUSTOMFIELDTYPE_MARKDOWN:
		return CONFIGFIELDTYPE_MARKDOWN
	case CUSTOMFIELDTYPE_CALLTOACTION:
		return CONFIGFIELDTYPE_CTA
	case CUSTOMFIELDTYPE_PRICE:
		return CONFIGFIELDTYPE_PRICE
	case CUSTOMFIELDTYPE_SIMPLEPHOTO, CUSTOMFIELDTYPE_PHOTO:
		return CONFIGFIELDTYPE_IMAGE
	case CUSTOMFIELDTYPE_SIMPLEVIDEO, CUSTOMFIELDTYPE_VIDEO:
		return CONFIGFIELDTYPE_VIDEO
	case CUSTOMFIELDTYPE_MULTIOPTION,
		CUSTOMFIELDTYPE_TEXTLIST,
		CUSTOMFIELDTYPE_SIMPLEPHOTOGALLERY,
		CUSTOMFIELDTYPE_GALLERY,
		CUSTOMFIELDTYPE_SIMPLYVIDEOGALLERY,
		CUSTOMFIELDTYPE_VIDEOGALLERY,
		CUSTOMFIELDTYPE_ENTITYLIST:
		return CONFIGFIELDTYPE_LIST
	case CUSTOMFIELDTYPE_SINGLELINETEXT,
		CUSTOMFIELDTYPE_MULTILINETEXT,
		CUSTOMFIELDTYPE_URL,
		CUSTOMFIELDTYPE_SLUG:
		return CONFIGFIELDTYPE_STRING
	}
	return c.Type
}

func (c *CustomField) GetConfigType() *ConfigType {
	configFieldTypeId := c.GetConfigFieldTypeId()
	switch configFieldTypeId {
	case CONFIGFIELDTYPE_YESNO:
		return &ConfigType{
			BooleanType: &ConfigBooleanType{},
		}
	case CONFIGFIELDTYPE_DECIMAL:
		return &ConfigType{
			DecimalType: &ConfigDecimalType{},
		}
	case CONFIGFIELDTYPE_DATE:
		return &ConfigType{
			DateType: &ConfigDateType{},
		}
	case CONFIGFIELDTYPE_IMAGE:
		r := &ConfigType{
			ImageType: &ConfigImageType{
				UnconstrainedAspectRatioAllowed: true,
			},
		}

		if c.Type == CUSTOMFIELDTYPE_SIMPLEPHOTO {
			r.ImageType.IsSimpleImage = true
		}
		return r
	case CONFIGFIELDTYPE_VIDEO:
		r := &ConfigType{
			VideoType: &ConfigVideoType{},
		}

		if c.Type == CUSTOMFIELDTYPE_SIMPLEVIDEO {
			r.VideoType.IsSimpleVideo = true
		}
		return r
	case CONFIGFIELDTYPE_LIST:
		switch c.Type {
		case CUSTOMFIELDTYPE_MULTIOPTION:
			return &ConfigType{
				ListType: &ConfigListType{
					TypeId: "option",
					Type: ConfigListSubType{
						ConfigType: &ConfigType{
							OptionType: &ConfigOptionType{
								Option: c.GetConfigOptions(),
							},
						},
					},
				},
			}
		case CUSTOMFIELDTYPE_TEXTLIST:
			return &ConfigType{
				ListType: &ConfigListType{
					TypeId: "string",
					Type: ConfigListSubType{
						ConfigType: &ConfigType{
							StringType: &ConfigStringType{
								Stereotype: "SIMPLE",
							},
						},
					},
				},
			}
		case CUSTOMFIELDTYPE_SIMPLEPHOTOGALLERY, CUSTOMFIELDTYPE_GALLERY:
			r := &ConfigType{
				ListType: &ConfigListType{
					TypeId: "image",
					Type: ConfigListSubType{
						ConfigType: &ConfigType{
							ImageType: &ConfigImageType{
								UnconstrainedAspectRatioAllowed: true,
							},
						},
					},
				},
			}

			if c.Type == CUSTOMFIELDTYPE_SIMPLEPHOTOGALLERY {
				r.ListType.Type.ImageType.IsSimpleImage = true
			}

			return r
		case CUSTOMFIELDTYPE_SIMPLYVIDEOGALLERY, CUSTOMFIELDTYPE_VIDEOGALLERY:
			r := &ConfigType{
				ListType: &ConfigListType{
					TypeId: "video",
					Type: ConfigListSubType{
						ConfigType: &ConfigType{
							VideoType: &ConfigVideoType{},
						},
					},
				},
			}

			if c.Type == CUSTOMFIELDTYPE_SIMPLYVIDEOGALLERY {
				r.ListType.Type.VideoType.IsSimpleVideo = true
			}

			return r
		case CUSTOMFIELDTYPE_ENTITYLIST:
			return &ConfigType{
				ListType: &ConfigListType{
					TypeId: "entityReference",
					Type: ConfigListSubType{
						EntityReferenceType: &ConfigEntityReferenceType{},
					},
				},
			}
		}
	case CONFIGFIELDTYPE_OPTION:
		return &ConfigType{
			OptionType: &ConfigOptionType{
				Option: c.GetConfigOptions(),
			},
		}
	case CONFIGFIELDTYPE_RICHTEXT:
		return &ConfigType{
			RichTextType: &ConfigRichTextType{},
		}
	case CONFIGFIELDTYPE_RICHTEXTV2:
		return &ConfigType{
			RichTextV2Type: &ConfigRichTextV2Type{},
		}
	case CONFIGFIELDTYPE_MARKDOWN:
		return &ConfigType{
			MarkdownType: &ConfigMarkdownType{},
		}
	case CONFIGFIELDTYPE_STRING:
		r := &ConfigType{
			StringType: &ConfigStringType{},
		}

		switch c.Type {
		case CUSTOMFIELDTYPE_SINGLELINETEXT:
			r.StringType.Stereotype = "SIMPLE"
		case CUSTOMFIELDTYPE_MULTILINETEXT:
			r.StringType.Stereotype = "MULTILINE"
		case CUSTOMFIELDTYPE_URL:
			r.StringType.Stereotype = "URL"
		case CUSTOMFIELDTYPE_SLUG:
			r.StringType.Stereotype = "SLUG"
		}

		return r
	}

	return nil
}

func (c *CustomField) GetConfigOptions() []ConfigOption {
	var (
		r []ConfigOption
	)

	configFieldTypeId := c.GetConfigFieldTypeId()
	if !(configFieldTypeId == CONFIGFIELDTYPE_OPTION || configFieldTypeId == CONFIGFIELDTYPE_LIST) {
		return nil
	}

	for _, entry := range c.Options {
		var translations []ConfigTranslation
		for _, translation := range entry.Translations {
			translations = append(translations, ConfigTranslation{
				LocaleCode: translation.LanguageCode,
				Value:      translation.Value,
			})
		}
		r = append(r, ConfigOption{
			TextValue:              entry.Key,
			DisplayName:            entry.Value,
			DisplayNameTranslation: translations,
		})
	}

	return r
}

func (c CustomField) GetId() string {
	if c.Id == nil {
		return ""
	}
	return *c.Id
}
