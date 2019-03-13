package yext

import (
	"fmt"
)

const entityProfilesPath = "entityprofiles"

type LanguageProfileService struct {
	client   *Client
	registry *EntityRegistry
}

type LanguageProfileListResponse struct {
	Profiles []interface{} `json:"profiles"`
}

func (l *LanguageProfileService) RegisterDefaultEntities() {
	l.registry = defaultEntityRegistry()
}

func (l *LanguageProfileService) RegisterEntity(t EntityType, entity interface{}) {
	l.registry.RegisterEntity(t, entity)
}

func (l *LanguageProfileService) Get(id string, languageCode string) (*Entity, *Response, error) {
	var v map[string]interface{}
	r, err := l.client.DoRequest("GET", fmt.Sprintf("%s/%s/%s", entityProfilesPath, id, languageCode), &v)
	if err != nil {
		return nil, r, err
	}

	entity, err := l.registry.ToEntityType(v)
	if err != nil {
		return nil, r, err
	}
	setNilIsEmpty(entity)

	return &entity, r, nil
}

func (l *LanguageProfileService) List(id string) ([]*Entity, *Response, error) {
	var (
		v        LanguageProfileListResponse
		profiles = []*Entity{}
	)
	r, err := l.client.DoRequest("GET", fmt.Sprintf("%s/%s", entityProfilesPath, id), &v)
	if err != nil {
		return nil, r, err
	}

	typedProfiles, err := l.registry.ToEntityTypes(v.Profiles)
	if err != nil {
		return nil, r, err
	}
	for _, profile := range typedProfiles {
		setNilIsEmpty(profile)
		profiles = append(profiles, &profile)
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
