package yext

import (
	"encoding/json"
	"fmt"
)

const (
	entityProfilesPath          = "entityprofiles"
	LanguageProfileListMaxLimit = 50
)

type LanguageProfileService struct {
	client   *Client
	registry *EntityRegistry
}

type LanguageProfileListResponse struct {
	Profiles []interface{} `json:"profiles"`
}

type LanguageProfileListAllResponse struct {
	Count        int    `json:"count"`
	PageToken    string `json:"pageToken"`
	ProfileLists []struct {
		Profiles []interface{} `json:"profiles"`
	} `json:"profileLists"`
	typedProfiles []Entity
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

func (l *LanguageProfileService) List(id string) ([]Entity, *Response, error) {
	var (
		v        LanguageProfileListResponse
		profiles = []Entity{}
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
		profiles = append(profiles, profile)
	}
	return profiles, r, nil
}

func (l *LanguageProfileService) ListAll(opts *EntityListOptions) ([]Entity, error) {
	var entities []Entity
	if opts == nil {
		opts = &EntityListOptions{}
	}
	opts.ListOptions = ListOptions{Limit: EntityListMaxLimit}
	var lg tokenListRetriever = func(listOptions *ListOptions) (string, error) {
		opts.ListOptions = *listOptions
		resp, _, err := l.listAllHelper(opts)
		if err != nil {
			return "", err
		}

		for _, entity := range resp.typedProfiles {
			entities = append(entities, entity)
		}
		return resp.PageToken, nil
	}

	if err := tokenListHelper(lg, &opts.ListOptions); err != nil {
		return nil, err
	}
	return entities, nil
}

func (l *LanguageProfileService) listAllHelper(opts *EntityListOptions) (*LanguageProfileListAllResponse, *Response, error) {
	var (
		requrl = entityProfilesPath
		err    error
	)

	if opts != nil {
		requrl, err = addEntityListOptions(requrl, opts)
		if err != nil {
			return nil, nil, err
		}
	}

	if opts != nil {
		requrl, err = addListOptions(requrl, &opts.ListOptions)
		if err != nil {
			return nil, nil, err
		}
	}

	v := &LanguageProfileListAllResponse{}
	r, err := l.client.DoRequest("GET", requrl, v)
	if err != nil {
		return nil, r, err
	}

	entityInterfaces := []interface{}{}
	for _, list := range v.ProfileLists {
		for _, entity := range list.Profiles {
			entityInterfaces = append(entityInterfaces, entity)
		}
	}

	typedEntities, err := l.registry.ToEntityTypes(entityInterfaces)
	if err != nil {
		return nil, r, err
	}

	for _, entity := range typedEntities {
		setNilIsEmpty(entity)
	}

	v.typedProfiles = typedEntities
	return v, r, nil
}

func (l *LanguageProfileService) Upsert(entity Entity, id string, languageCode string) (*Response, error) {
	asJSON, err := json.Marshal(entity)
	if err != nil {
		return nil, err
	}
	var asMap map[string]interface{}
	err = json.Unmarshal(asJSON, &asMap)
	if err != nil {
		return nil, err
	}
	delete(asMap["meta"].(map[string]interface{}), "id")

	r, err := l.client.DoRequestJSON("PUT", fmt.Sprintf("%s/%s/%s", entityProfilesPath, id, languageCode), asMap, nil)
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
