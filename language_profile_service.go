package yext

import (
	"fmt"
)

const entityProfilesPath = "entityprofiles"

type LanguageProfileService struct {
	client   *Client
	registry Registry
}

type LanguageProfileListResponse struct {
	Profiles []interface{} `json:"profiles"`
}

func (l *LanguageProfileService) RegisterDefaultEntities() {
	l.registry = defaultEntityRegistry()
}

func (l *LanguageProfileService) RegisterEntity(t EntityType, entity interface{}) {
	l.registry.Register(string(t), entity)
}

func (l *LanguageProfileService) Get(id string, languageCode string) (*LanguageProfile, *Response, error) {
	var v map[string]interface{}
	r, err := l.client.DoRequest("GET", fmt.Sprintf("%s/%s/%s", entityProfilesPath, id, languageCode), &v)
	if err != nil {
		return nil, r, err
	}

	entity, err := toEntityType(v, l.registry)
	if err != nil {
		return nil, r, err
	}
	setNilIsEmpty(entity)

	return &LanguageProfile{Entity: entity}, r, nil
}

func (l *LanguageProfileService) List(id string) ([]*LanguageProfile, *Response, error) {
	var (
		v        LanguageProfileListResponse
		profiles = []*LanguageProfile{}
	)
	r, err := l.client.DoRequest("GET", fmt.Sprintf("%s/%s", entityProfilesPath, id), &v)
	if err != nil {
		return nil, r, err
	}

	typedProfiles, err := toEntityTypes(v.Profiles, l.registry)
	if err != nil {
		return nil, r, err
	}
	for _, profile := range typedProfiles {
		setNilIsEmpty(profile)
		profiles = append(profiles, &LanguageProfile{Entity: profile})
	}
	return profiles, r, nil
}

func (l *LanguageProfileService) Upsert(entity Entity, languageCode string) (*Response, error) {
	id := entity.GetEntityId()
	if id == "" {
		return nil, fmt.Errorf("entity profile service upsert: profile object had no id")
	}

	r, err := l.client.DoRequestJSON("PUT", fmt.Sprintf("%s/%s/%s", entityProfilesPath, id, languageCode), entity, nil)
	if err != nil {
		return r, err
	}
	return r, nil
}

func (l *LanguageProfileService) Delete(id string, languageCode string) (*Response, error) {
	r, err := l.client.DoRequest("DELETE", fmt.Sprintf("%s/%s/%s", entityProfilesPath, id, languageCode), nil)
	if err != nil {
		return r, err
	}
	return r, nil
}
