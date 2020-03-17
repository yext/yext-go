package yext

import (
	"encoding/json"
)

const ENTITYTYPE_FAQ EntityType = "faq"

// FAQ is the representation of a FAQ in Yext Location Manager.
type FAQEntity struct {
	BaseEntity

	Name           *string   `json:"name,omitempty"`
	Answer         *string   `json:"answer,omitempty"`
	LandingPageUrl *string   `json:"landingPageUrl,omitempty"`
	Keywords       *[]string `json:"keywords,omitempty"`
}

func (y FAQEntity) GetName() string {
	if y.Name != nil {
		return GetString(y.Name)
	}
	return ""
}

func (y FAQEntity) GetAnswer() string {
	if y.Answer != nil {
		return GetString(y.Answer)
	}
	return ""
}

func (y FAQEntity) GetLandingPageUrl() string {
	if y.LandingPageUrl != nil {
		return GetString(y.LandingPageUrl)
	}
	return ""
}

func (y FAQEntity) GetKeywords() (v []string) {
	if y.Keywords != nil {
		v = *y.Keywords
	}
	return v
}

func (l *FAQEntity) UnmarshalJSON(data []byte) error {
	type Alias FAQEntity
	a := &struct {
		*Alias
	}{
		Alias: (*Alias)(l),
	}
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	return UnmarshalEntityJSON(l, data)
}

func (j *FAQEntity) String() string {
	b, _ := json.Marshal(j)
	return string(b)
}
