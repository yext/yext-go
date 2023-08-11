package yext

import "unicode"

const (
	CONFIGFIELDTYPE_YESNO           = "boolean"
	CONFIGFIELDTYPE_DECIMAL         = "decimal"
	CONFIGFIELDTYPE_DATE            = "date"
	CONFIGFIELDTYPE_DAILYTIMES      = "dailyTimes"
	CONFIGFIELDTYPE_HOURS           = "hours"
	CONFIGFIELDTYPE_IMAGE           = "image"
	CONFIGFIELDTYPE_OPTION          = "option"
	CONFIGFIELDTYPE_VIDEO           = "video"
	CONFIGFIELDTYPE_RICHTEXT        = "richText"
	CONFIGFIELDTYPE_RICHTEXTV2      = "richTextV2"
	CONFIGFIELDTYPE_LIST            = "list"
	CONFIGFIELDTYPE_STRING          = "string"
	CONFIGFIELDTYPE_MARKDOWN        = "markdown"
	CONFIGFIELDTYPE_CTA             = "cta"
	CONFIGFIELDTYPE_PRICE           = "price"
	CONFIGFIELDTYPE_ENTITYREFERENCE = "entityReference"
)

type ConfigFieldEligibilityGroup struct {
	Id         *string    `json:"$id,omitempty"`
	Schema     string     `json:"$schema,omitempty"`
	Name       string     `json:"name,omitempty"`
	EntityType EntityType `json:"entityType,omitempty"`
	Fields     []string   `json:"fields,omitempty"`
	Required   []string   `json:"required,omitempty"`
}

type ConfigTranslation struct {
	LocaleCode string `json:"localeCode,omitempty"`
	Value      string `json:"value,omitempty"`
}

type ConfigOption struct {
	DisplayName            string              `json:"displayName,omitempty"`
	DisplayNameTranslation []ConfigTranslation `json:"displayNameTranslation,omitempty"`
	TextValue              string              `json:"textValue,omitempty"`
}

type ConfigOptionType struct {
	Option    []ConfigOption `json:"option,omitempty"`
	NoOptions bool           `json:"noOptions,omitempty"`
}

type ConfigStringType struct {
	Stereotype string `json:"stereotype,omitempty"`
	MinLength  int    `json:"minLength,omitempty"`
	MaxLength  int    `json:"maxLength,omitempty"`
}

type ConfigDescriptionType struct {
	StereoType string `json:"stereoType,omitempty"`
	MinLength  int    `json:"minLength,omitempty"`
	MaxLength  int    `json:"maxLength,omitempty"`
}

type ConfigVideoType struct {
	IsSimpleVideo   bool                   `json:"isSimpleVideo,omitempty"`
	DescriptionType *ConfigDescriptionType `json:"descriptionType,omitempty"`
}

type ConfigMinSizeType struct {
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

type ConfigImageType struct {
	IsSimpleImage      bool               `json:"isSimpleImage,omitempty"`
	MinSize            *ConfigMinSizeType `json:"minSize,omitempty"`
	AllowedAspectRatio *[]struct {
		HorizontalFactor int `json:"horizontalFactor,omitempty"`
		VerticalFactor   int `json:"verticalFactor,omitempty"`
	} `json:"allowedAspectRatio,omitempty"`
	UnconstrainedAspectRatioAllowed bool `json:"unconstrainedAspectRatioAllowed,omitempty"`
	NoAllowedAspectRatio            bool `json:"noAllowedAspectRatio,omitempty"`
}

type ConfigEntityReferenceType struct {
	SupportedEntityTypeIds []string `json:"supportedEntityTypeIds,omitempty"`
	RelatedFieldId         string   `json:"relatedFieldId,omitempty"`
	Type                   string   `json:"type,omitempty"`
}

type ConfigListSubType struct {
	*ConfigType
	EntityReferenceType *ConfigEntityReferenceType `json:"entityReferenceType,omitempty"`
}

type ConfigListType struct {
	TypeId    string            `json:"typeId,omitempty"`
	Type      ConfigListSubType `json:"type,omitempty"`
	MinLength int               `json:"minLength,omitempty"`
	MaxLength int               `json:"maxLength,omitempty"`
}

type ConfigRichTextType struct {
	MinLength          int      `json:"minLength,omitempty"`
	MaxLength          int      `json:"maxLength,omitempty"`
	SupportedFormat    []string `json:"supportedFormat,omitempty"`
	NoSupportedFormats bool     `json:"noSupportedFormats,omitempty"`
}

type ConfigRichTextV2Type struct {
	MinLength             int `json:"minLength,omitempty"`
	MaxLength             int `json:"maxLength,omitempty"`
	SupportedFormatFilter *[]struct {
		SupportedFormats []string `json:"supportedFormats,omitempty"`
	} `json:"supportedFormatFilter,omitempty"`
	NoSupportedFormat bool `json:"noSupportedFormat,omitempty"`
}

type ConfigMarkdownType struct {
	MinLength int `json:"minLength,omitempty"`
	MaxLength int `json:"maxLength,omitempty"`
}

type ConfigBooleanType map[string]interface{}

type ConfigDayType struct {
	Year  string `json:"year,omitempty"`
	Month string `json:"month,omitempty"`
	Day   string `json:"day,omitempty"`
}

type ConfigDateType struct {
	MinValue *ConfigDayType `json:"minValue,omitempty"`
	MaxValue *ConfigDayType `json:"maxValue,omitempty"`
}

type ConfigDecimalType struct {
	MinValue string `json:"minValue,omitempty"`
	MaxValue string `json:"maxValue,omitempty"`
}

type ConfigType struct {
	BooleanType    *ConfigBooleanType    `json:"booleanType,omitempty"`
	DateType       *ConfigDateType       `json:"dateType,omitempty"`
	DecimalType    *ConfigDecimalType    `json:"decimalType,omitempty"`
	ImageType      *ConfigImageType      `json:"imageType,omitempty"`
	RichTextV2Type *ConfigRichTextV2Type `json:"richTextV2Type,omitempty"`
	ListType       *ConfigListType       `json:"listType,omitempty"`
	MarkdownType   *ConfigMarkdownType   `json:"markdownType,omitempty"`
	OptionType     *ConfigOptionType     `json:"optionType,omitempty"`
	RichTextType   *ConfigRichTextType   `json:"richTextType,omitempty"`
	StringType     *ConfigStringType     `json:"stringType,omitempty"`
	VideoType      *ConfigVideoType      `json:"videoType,omitempty"`
}

type ConfigField struct {
	Id           *string     `json:"$id,omitempty"`
	Description  string      `json:"description,omitempty"`
	Schema       string      `json:"$schema,omitempty"`
	DisplayName  string      `json:"displayName,omitempty"`
	Group        string      `json:"group,omitempty"`
	Localization string      `json:"localization,omitempty"`
	Type         *ConfigType `json:"type,omitempty"`
	TypeId       string      `json:"typeId,omitempty"`
}

func (c *ConfigField) TransformToCustomField(entityAvailability []EntityType) *CustomField {
	r := &CustomField{
		Id:                         c.Id,
		Type:                       c.GetCustomFieldType(),
		Name:                       c.DisplayName,
		AlternateLanguageBehaviour: c.Localization,
		Options:                    c.GetCustomFieldOptions(),
		Group:                      "NONE",
		Description:                c.Description,
		EntityAvailability:         entityAvailability,
	}

	if c.Localization == "LOCALE_SPECIFIC" {
		r.AlternateLanguageBehaviour = "LANGUAGE_SPECIFIC"
	}

	for i, char := range c.Group {
		if unicode.IsDigit(char) {
			r.Group = c.Group[:i] + "_" + c.Group[i:]
		}
	}

	return r
}

func (c *ConfigField) GetCustomFieldType() string {
	switch c.TypeId {
	case CONFIGFIELDTYPE_YESNO:
		return CUSTOMFIELDTYPE_YESNO
	case CONFIGFIELDTYPE_DECIMAL:
		return CUSTOMFIELDTYPE_NUMBER
	case CONFIGFIELDTYPE_DATE:
		return CUSTOMFIELDTYPE_DATE
	case CONFIGFIELDTYPE_DAILYTIMES:
		return CUSTOMFIELDTYPE_DAILYTIMES
	case CONFIGFIELDTYPE_HOURS:
		return CUSTOMFIELDTYPE_HOURS
	case CONFIGFIELDTYPE_OPTION:
		return CUSTOMFIELDTYPE_SINGLEOPTION
	case CONFIGFIELDTYPE_RICHTEXT:
		return CUSTOMFIELDTYPE_RICHTEXT
	case CONFIGFIELDTYPE_RICHTEXTV2:
		return CUSTOMFIELDTYPE_RICHTEXTV2
	case CONFIGFIELDTYPE_MARKDOWN:
		return CUSTOMFIELDTYPE_MARKDOWN
	case CONFIGFIELDTYPE_CTA:
		return CUSTOMFIELDTYPE_CALLTOACTION
	case CONFIGFIELDTYPE_PRICE:
		return CUSTOMFIELDTYPE_PRICE
	case CONFIGFIELDTYPE_IMAGE:
		if c.Type.ImageType != nil && c.Type.ImageType.IsSimpleImage {
			return CUSTOMFIELDTYPE_SIMPLEPHOTO
		}
		return CUSTOMFIELDTYPE_PHOTO
	case CONFIGFIELDTYPE_VIDEO:
		if c.Type.VideoType != nil && c.Type.VideoType.IsSimpleVideo {
			return CUSTOMFIELDTYPE_SIMPLEVIDEO
		}
		return CUSTOMFIELDTYPE_VIDEO
	case CONFIGFIELDTYPE_LIST:
		if c.Type.ListType != nil {
			switch c.Type.ListType.TypeId {
			case CONFIGFIELDTYPE_OPTION:
				return CUSTOMFIELDTYPE_MULTIOPTION
			case CONFIGFIELDTYPE_STRING:
				if c.Type.ListType.Type.StringType != nil && c.Type.ListType.Type.StringType.Stereotype == "SIMPLE" {
					return CUSTOMFIELDTYPE_TEXTLIST
				}
			case CONFIGFIELDTYPE_IMAGE:
				if c.Type.ListType.Type.ImageType != nil && c.Type.ListType.Type.ImageType.IsSimpleImage {
					return CUSTOMFIELDTYPE_SIMPLEPHOTOGALLERY
				} else {
					return CUSTOMFIELDTYPE_GALLERY
				}
			case CONFIGFIELDTYPE_VIDEO:
				if c.Type.ListType.Type.VideoType != nil && c.Type.ListType.Type.VideoType.IsSimpleVideo {
					return CUSTOMFIELDTYPE_SIMPLYVIDEOGALLERY
				} else {
					return CUSTOMFIELDTYPE_VIDEOGALLERY
				}
			case CONFIGFIELDTYPE_ENTITYREFERENCE:
				return CUSTOMFIELDTYPE_ENTITYLIST
			}
		}
		return CUSTOMFIELDTYPE_LIST
	case CONFIGFIELDTYPE_STRING:
		if c.Type.StringType != nil {
			switch c.Type.StringType.Stereotype {
			case "SIMPLE":
				return CUSTOMFIELDTYPE_SINGLELINETEXT
			case "MULTILINE":
				return CUSTOMFIELDTYPE_MULTILINETEXT
			case "URL":
				return CUSTOMFIELDTYPE_URL
			case "SLUG":
				return CUSTOMFIELDTYPE_SLUG
			}
		}
		return CUSTOMFIELDTYPE_SINGLELINETEXT
	}

	return CUSTOMFIELDTYPE_STRUCT
}

func (c *ConfigField) GetCustomFieldOptions() []CustomFieldOption {
	var (
		r []CustomFieldOption
		o []ConfigOption
	)

	customFieldType := c.GetCustomFieldType()
	if customFieldType == CUSTOMFIELDTYPE_SINGLEOPTION {
		o = c.Type.OptionType.Option
	} else if customFieldType == CUSTOMFIELDTYPE_MULTIOPTION {
		o = c.Type.ListType.Type.OptionType.Option
	} else {
		return nil
	}

	for _, entry := range o {
		var translations = []Translation{}
		for _, translation := range entry.DisplayNameTranslation {
			translations = append(translations, Translation{
				LanguageCode: translation.LocaleCode,
				Value:        translation.Value,
			})
		}

		r = append(r, CustomFieldOption{
			Key:          entry.TextValue,
			Value:        entry.DisplayName,
			Translations: translations,
		})
	}

	return r
}
