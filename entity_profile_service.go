package yext

import "fmt"

const entityProfilesPath = "entityprofiles"

type EntityProfileService struct {
	client   *Client
	registry Registry
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
