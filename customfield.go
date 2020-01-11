package yext

import "fmt"

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
)

type CustomField struct {
	Id                         *string             `json:"id,omitempty"`
	Type                       string              `json:"type"`
	Name                       string              `json:"name"`
	Options                    []CustomFieldOption `json:"options,omitempty"` // Only present for option custom fields
	Group                      string              `json:"group"`
	Description                string              `json:"description"`
	AlternateLanguageBehaviour string              `json:"alternateLanguageBehavior"`
	EntityAvailability         []EntityType        `json:"entityAvailability"`
}

func (c CustomField) GetId() string {
	if c.Id == nil {
		return ""
	}
	return *c.Id
}

type EntityList UnorderedStrings

func (l EntityList) CustomFieldTag() string {
	return CUSTOMFIELDTYPE_ENTITYLIST
}

func (m EntityList) Equal(c Comparable) bool {
	var n EntityList
	switch v := c.(type) {
	case EntityList:
		n = v
	case *EntityList:
		n = *v
	default:
		panic(fmt.Errorf("%v is not a EntityList is %T", c, c))
	}
	if len(m) != len(n) {
		return false
	}
	a := UnorderedStrings(m)
	b := UnorderedStrings(n)
	return (&a).Equal(&b)
}
