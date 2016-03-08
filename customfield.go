package yext

type CustomFieldType string

const (
	CUSTOMFIELDTYPE_YESNO          = "BOOLEAN"
	CUSTOMFIELDTYPE_SINGLELINETEXT = "TEXT"
	CUSTOMFIELDTYPE_MULTILINETEXT  = "MULTILINETEXT"
	CUSTOMFIELDTYPE_SINGLEOPTION   = "SINGLEOPTION"
	CUSTOMFIELDTYPE_URL            = "URL"
	CUSTOMFIELDTYPE_DATE           = "DATE"
	CUSTOMFIELDTYPE_NUMBER         = "NUMBER"
	CUSTOMFIELDTYPE_MULTIOPTION    = "MULTIOPTION"
	CUSTOMFIELDTYPE_TEXTLIST       = "TEXTLIST"
	CUSTOMFIELDTYPE_PHOTO          = "PHOTO"
	CUSTOMFIELDTYPE_GALLERY        = "GALLERY"
	CUSTOMFIELDTYPE_VIDEO          = "VIDEO"
	CUSTOMFIELDTYPE_HOURS          = "HOURS"
	// not sure what to do with "DAILYTIMES", omitting
)

// CustomField is the representation of a Custom Field definition in Yext Location Manager.
// For details see https://www.yext.com/support/platform-api/#Administration_API/Custom_Fields.htm
type CustomField struct {
	Name    string            `json:"name"`
	Id      string            `json:"id"`
	Options map[string]string `json:"options"` // Only present for multi-option custom fields
	Type    string            `json:"type"`
}

type YesNo bool

type SingleLineText string
type MultiLineText string
type SingleOption string
type Url string
type Date string
type Number string

type MultiOption []string
type TextList []string

type Image struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Url    string `json:"url"`
}

type CustomPhoto struct {
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	ClickthroughUrl string  `json:"clickthroughUrl"`
	Sizes           []Image `json:"sizes"`
}

type Gallery []CustomPhoto

type Video struct {
	Description string `json:"description"`
	Url         string `json:"url"`
}

type VideoGallery []Video

type Hours struct {
	AdditionalText string         `json:"additionalHoursText"`
	Hours          string         `json:"hours"`
	HolidayHours   []HolidayHours `json:"holidayHours"`
}
