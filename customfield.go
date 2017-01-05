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
	Name    string `json:"name"`
	Id      string `json:"id"`
	Options []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"options"` // Only present for multi-option custom fields
	Type string `json:"type"`
}

type CustomFieldValue interface {
	CustomFieldTag() string
}

type CustomFieldValueComparable interface {
	Equal(CustomFieldValueComparable) bool
}

type YesNo bool

func (y YesNo) CustomFieldTag() string {
	return CUSTOMFIELDTYPE_YESNO
}

type SingleLineText string

func (s SingleLineText) CustomFieldTag() string {
	return CUSTOMFIELDTYPE_SINGLELINETEXT
}

type MultiLineText string

func (m MultiLineText) CustomFieldTag() string {
	return CUSTOMFIELDTYPE_MULTILINETEXT
}

type Url string

func (u Url) CustomFieldTag() string {
	return CUSTOMFIELDTYPE_URL
}

type Date string

func (d Date) CustomFieldTag() string {
	return CUSTOMFIELDTYPE_DATE
}

type Number string

func (n Number) CustomFieldTag() string {
	return CUSTOMFIELDTYPE_NUMBER
}

type OptionField interface {
	CustomFieldValue
	SetOptionId(id string)
	UnsetOptionId(id string)
	IsOptionIdSet(id string) bool
}

type SingleOption string

func (s SingleOption) CustomFieldTag() string {
	return CUSTOMFIELDTYPE_SINGLEOPTION
}

func (s *SingleOption) SetOptionId(id string) {
	*s = SingleOption(id)
}

func (s *SingleOption) UnsetOptionId(id string) {
	if string(*s) == id {
		*s = SingleOption("")
	}
}

func (s *SingleOption) IsOptionIdSet(id string) bool {
	return *s == SingleOption(id)
}

type MultiOption []string

func (m MultiOption) CustomFieldTag() string {
	return CUSTOMFIELDTYPE_MULTIOPTION
}

func (m MultiOption) Equal(c CustomFieldValueComparable) bool {
	n := c.(MultiOption)
	if len(m) != len(n) {
		return false
	}
	for i := 0; i < len(m); i++ {
		found := false
		for j := 0; j < len(n); j++ {
			if m[i] == n[j] {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (m *MultiOption) SetOptionId(id string) {
	if !m.IsOptionIdSet(id) {
		*m = append(*m, id)
	}
}

func (m *MultiOption) UnsetOptionId(id string) {
	if m.IsOptionIdSet(id) {
		t := []string(*m)
		indexOfTarget := -1
		for i := 0; i < len(*m); i++ {
			if t[i] == id {
				indexOfTarget = i
			}
		}
		if indexOfTarget >= 0 {
			*m = append(t[:indexOfTarget], t[indexOfTarget+1:]...)
		}
	}
}

func (m *MultiOption) IsOptionIdSet(id string) bool {
	for _, option := range *m {
		if option == id {
			return true
		}
	}
	return false
}

type TextList []string

func (t TextList) CustomFieldTag() string {
	return CUSTOMFIELDTYPE_TEXTLIST
}

type Photo struct {
	Url             string `json:"url,omitempty"`
	Description     string `json:"description,omitempty"`
	Details         string `json:"details,omitempty"`
	ClickThroughURL string `json:"clickthroughUrl,omitempty"`
}

func (p *Photo) CustomFieldTag() string {
	return CUSTOMFIELDTYPE_PHOTO
}

type Gallery []*Photo

func (g *Gallery) CustomFieldTag() string {
	return CUSTOMFIELDTYPE_GALLERY
}

type Video struct {
	Description string `json:"description"`
	Url         string `json:"url"`
}

type VideoGallery []Video

func (v *VideoGallery) CustomFieldTag() string {
	return CUSTOMFIELDTYPE_VIDEO
}

type Hours struct {
	AdditionalText string         `json:"additionalHoursText,omitempty"`
	Hours          string         `json:"hours,omitempty"`
	HolidayHours   []HolidayHours `json:"holidayHours,omitempty"`
}

func (h *Hours) CustomFieldTag() string {
	return CUSTOMFIELDTYPE_HOURS
}
