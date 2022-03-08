package yext

// Custom field service to be used with legacy location service

import (
	"fmt"
)

type CustomFieldType string

const (
	CUSTOMFIELDTYPE_LOCATIONLIST = "LOCATION_LIST"
)

var (
	UnsetPhotoValue = (*CustomLocationPhoto)(nil)
)

type Translation struct {
	LanguageCode string `json:"languageCode"`
	Value        string `json:"value"`
}

type CustomFieldOption struct {
	Key          string        `json:"key,omitempty"`
	Value        string        `json:"value"`
	Translations []Translation `json:"translations,omitempty"`
}

type CustomFieldValue interface {
	CustomFieldTag() string
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

type MultiOption UnorderedStrings

func (m MultiOption) CustomFieldTag() string {
	return CUSTOMFIELDTYPE_MULTIOPTION
}

func (m MultiOption) Equal(c Comparable) bool {
	var n MultiOption
	switch v := c.(type) {
	case MultiOption:
		n = v
	case *MultiOption:
		n = *v
	default:
		panic(fmt.Errorf("%v is not a MultiOption is %T", c, c))
	}
	if len(m) != len(n) {
		return false
	}
	a := UnorderedStrings(m)
	b := UnorderedStrings(n)
	return (&a).Equal(&b)

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

type LocationList UnorderedStrings

func (l LocationList) CustomFieldTag() string {
	return CUSTOMFIELDTYPE_LOCATIONLIST
}

func (m LocationList) Equal(c Comparable) bool {
	var n LocationList
	switch v := c.(type) {
	case LocationList:
		n = v
	case *LocationList:
		n = *v
	default:
		panic(fmt.Errorf("%v is not a LocationList is %T", c, c))
	}
	if len(m) != len(n) {
		return false
	}
	a := UnorderedStrings(m)
	b := UnorderedStrings(n)
	return (&a).Equal(&b)
}

type CustomLocationPhoto struct {
	Url             string `json:"url,omitempty"`
	Description     string `json:"description,omitempty"`
	Details         string `json:"details,omitempty"`
	ClickThroughURL string `json:"clickthroughUrl,omitempty"`
}

func (p *CustomLocationPhoto) CustomFieldTag() string {
	return CUSTOMFIELDTYPE_PHOTO
}

type CustomLocationGallery []*CustomLocationPhoto

func (g *CustomLocationGallery) CustomFieldTag() string {
	return CUSTOMFIELDTYPE_GALLERY
}

type CustomLocationVideo struct {
	Description string `json:"description"`
	Url         string `json:"url"`
}

type VideoGallery []CustomLocationVideo

func (v *VideoGallery) CustomFieldTag() string {
	return CUSTOMFIELDTYPE_VIDEO
}

// CustomLocationHours is the Hours custom field format used by locations API
// Entities API uses the Hours struct in location_entities.go (profile and custom hours are defined the same way for entities)
type CustomLocationHours struct {
	AdditionalText string                 `json:"additionalHoursText,omitempty"`
	Hours          string                 `json:"hours,omitempty"`
	HolidayHours   []LocationHolidayHours `json:"holidayHours,omitempty"`
}

func (h CustomLocationHours) CustomFieldTag() string {
	return CUSTOMFIELDTYPE_HOURS
}

// CustomLocationDailyTimes is the DailyTimes custom field format used by locations API
// Entities API uses the new DailyTimes struct
type CustomLocationDailyTimes struct {
	DailyTimes string `json:"dailyTimes,omitempty"`
}

func (d CustomLocationDailyTimes) CustomFieldTag() string {
	return CUSTOMFIELDTYPE_DAILYTIMES
}

type DailyTimes struct {
	Sunday    string `json:"sunday,omitempty"`
	Monday    string `json:"monday,omitempty"`
	Tuesday   string `json:"tuesday,omitempty"`
	Wednesday string `json:"wednesday,omitempty"`
	Thursday  string `json:"thursday,omitempty"`
	Friday    string `json:"friday,omitempty"`
	Saturday  string `json:"saturday,omitempty"`
}
