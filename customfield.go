package yext

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

type CustomFieldValidation struct {
	MinCharLength   int      `json:"minCharLength,omitempty"`
	MaxCharLength   int      `json:"maxCharLength,omitempty"`
	MinItemCount    int      `json:"minItemCount,omitempty"`
	MaxItemCount    int      `json:"maxItemCount,omitempty"`
	MinValue        int      `json:"minValue,omitempty"`
	MaxValue        int      `json:"maxValue,omitempty"`
	MinDate         string   `json:"minDate,omitempty"`
	MaxDate         string   `json:"maxDate,omitempty"`
	AspectRatio     string   `json:"aspectRatio,omitempty"`
	MinWidth        int      `json:"minWidth,omitempty"`
	MinHeight       int      `json:"minHeight,omitempty"`
	EntityTypes     []string `json:"entityTypes,omitempty"`
	RichTextFormats []string `json:"richTextFormats,omitempty"`
}

type CustomField struct {
	Id                         *string                `json:"id,omitempty"`
	Type                       string                 `json:"type"`
	Name                       string                 `json:"name"`
	Options                    []CustomFieldOption    `json:"options,omitempty"` // Only present for option custom fields
	Group                      string                 `json:"group"`
	Description                string                 `json:"description"`
	AlternateLanguageBehaviour string                 `json:"alternateLanguageBehavior"`
	EntityAvailability         []EntityType           `json:"entityAvailability"`
	Validation                 *CustomFieldValidation `json:"validation,omitempty"` // Needed for rich text formatting
}

func (c CustomField) GetId() string {
	if c.Id == nil {
		return ""
	}
	return *c.Id
}
