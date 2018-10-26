package yext

import (
	"fmt"
)

const entityProfilesPath = "entityprofiles"

type EntityProfileService struct {
	client   *Client
	registry Registry
}

type EntityProfileListResponse struct {
	Profiles []interface{} `json:"profiles"`
}

func (e *EntityProfileService) RegisterDefaultEntities() {
	e.registry = defaultEntityRegistry()
}

func (e *EntityProfileService) RegisterEntity(t EntityType, entity interface{}) {
	e.registry.Register(string(t), entity)
}

func (e *EntityProfileService) Get(id string, languageCode string) (*EntityProfile, *Response, error) {
	var v map[string]interface{}
	r, err := e.client.DoRequest("GET", fmt.Sprintf("%s/%s/%s", entityProfilesPath, id, languageCode), &v)
	if err != nil {
		return nil, r, err
	}

	entity, err := toEntityType(v, e.registry)
	if err != nil {
		return nil, r, err
	}
	setNilIsEmpty(entity)

	return &EntityProfile{Entity: entity}, r, nil
}

func (e *EntityProfileService) List(id string) ([]*EntityProfile, *Response, error) {
	var (
		v        EntityProfileListResponse
		profiles = []*EntityProfile{}
	)
	r, err := e.client.DoRequest("GET", fmt.Sprintf("%s/%s", entityProfilesPath, id), &v)
	if err != nil {
		return nil, r, err
	}

	typedProfiles, err := toEntityTypes(v.Profiles, e.registry)
	if err != nil {
		return nil, r, err
	}
	for _, profile := range typedProfiles {
		setNilIsEmpty(profile)
		profiles = append(profiles, &EntityProfile{Entity: profile})
	}
	return profiles, r, nil
}

func (e *EntityProfileService) Upsert(entity Entity, languageCode string) (*Response, error) {
	id := entity.GetEntityId()
	if id == "" {
		return nil, fmt.Errorf("entity profile service upsert: profile object had no id")
	}

	r, err := e.client.DoRequestJSON("PUT", fmt.Sprintf("%s/%s/%s", entityProfilesPath, id, languageCode), entity, nil)
	if err != nil {
		return r, err
	}
	return r, nil
}

func (e *EntityProfileService) Delete(id string, languageCode string) (*Response, error) {
	r, err := e.client.DoRequest("DELETE", fmt.Sprintf("%s/%s/%s", entityProfilesPath, id, languageCode), nil)
	if err != nil {
		return r, err
	}
	return r, nil
}
