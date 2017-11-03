package yext

import (
	"encoding/json"
	"fmt"
)

const profilesPath = "profiles"

type LanguageProfileService struct {
	client       *Client
	CustomFields []*CustomField
}

type LanguageProfileListResponse struct {
	LanguageProfiles []*LanguageProfile `json:"languageProfiles"`
}

func (l *LanguageProfileListResponse) ResponseAsLocations() []*Location {
	languageProfilesAsLocs := []*Location{}
	for _, lp := range l.LanguageProfiles {
		languageProfilesAsLocs = append(languageProfilesAsLocs, &lp.Location)
	}
	return languageProfilesAsLocs
}

func (l *LanguageProfileService) GetAll(id string) (*LanguageProfileListResponse, *Response, error) {
	var v LanguageProfileListResponse
	r, err := l.client.DoRequest("GET", fmt.Sprintf("%s/%s/%s", locationsPath, id, profilesPath), &v)
	if err != nil {
		return nil, r, err
	}

	if _, err := l.HydrateLocations(v.LanguageProfiles); err != nil {
		return nil, r, err
	}

	return &v, r, nil
}

func (l *LanguageProfileService) Get(id string, languageCode string) (*LanguageProfile, *Response, error) {
	var v LanguageProfile
	r, err := l.client.DoRequest("GET", fmt.Sprintf("%s/%s/%s/%s", locationsPath, id, profilesPath, languageCode), &v)
	if err != nil {
		return nil, r, err
	}

	if _, err := HydrateLocation(&v.Location, l.CustomFields); err != nil {
		return nil, r, err
	}

	return &v, r, nil
}

func (l *LanguageProfileService) Upsert(y *LanguageProfile, languageCode string) (*Response, error) {
	id := y.GetId()
	if id == "" {
		return nil, fmt.Errorf("language profile service: upsert: profile object had no id")
	}
	asJSON, err := json.Marshal(y)
	if err != nil {
		return nil, err
	}
	var asMap map[string]interface{}
	err = json.Unmarshal(asJSON, &asMap)
	if err != nil {
		return nil, err
	}
	delete(asMap, "id")

	if err := validateCustomFields(y.CustomFields); err != nil {
		return nil, err
	}
	r, err := l.client.DoRequestJSON("PUT", fmt.Sprintf("%s/%s/%s/%s", locationsPath, id, profilesPath, languageCode), asMap, nil)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (l *LanguageProfileService) Delete(id string, languageCode string) (*Response, error) {
	r, err := l.client.DoRequest("DELETE", fmt.Sprintf("%s/%s/%s/%s", locationsPath, id, profilesPath, languageCode), nil)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (l *LanguageProfileService) HydrateLocations(languageProfiles []*LanguageProfile) ([]*LanguageProfile, error) {
	if l.CustomFields == nil {
		return languageProfiles, nil
	}

	for _, profile := range languageProfiles {
		_, err := HydrateLocation(&profile.Location, l.CustomFields)
		if err != nil {
			return languageProfiles, err
		}
	}

	return languageProfiles, nil
}
