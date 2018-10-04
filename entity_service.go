package yext

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"net/url"
)

const entityPath = "entities"

type EntityService struct {
	client   *Client
	registry Registry
}

type EntityListOptions struct {
	ListOptions
	SearchID            string
	ResolvePlaceholders bool
}

type EntityListResponse struct {
	Count     int           `json:"count"`
	Entities  []interface{} `json:"entities"`
	PageToken string        `json:"pageToken"`
}

func (e *EntityService) RegisterDefaultEntities() {
	e.registry = make(Registry)
	e.RegisterEntity(ENTITYTYPE_LOCATION, &Location{})
	e.RegisterEntity(ENTITYTYPE_EVENT, &Event{})
}

func (e *EntityService) RegisterEntity(entityType EntityType, entity interface{}) {
	e.registry.Register(string(entityType), entity)
}

func (e *EntityService) LookupEntity(entityType EntityType) (interface{}, error) {
	return e.registry.Lookup(string(entityType))
}

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (e *EntityService) toEntityTypes(entityInterfaces []interface{}) ([]Entity, error) {
	var entities = []Entity{}
	for _, entityInterface := range entityInterfaces {
		entity, err := e.toEntityType(entityInterface)
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}
	return entities, nil
}

func (e *EntityService) toEntityType(entityInterface interface{}) (Entity, error) {
	// Determine Entity Type
	var entityValsByKey = entityInterface.(map[string]interface{})
	meta, ok := entityValsByKey["meta"]
	if !ok {
		return nil, fmt.Errorf("Unable to find meta attribute in %v", entityValsByKey)
	}

	var metaByKey = meta.(map[string]interface{})
	entityType, ok := metaByKey["entityType"]
	if !ok {
		return nil, fmt.Errorf("Unable to find entityType attribute in %v", metaByKey)
	}

	entityObj, err := e.LookupEntity(EntityType(entityType.(string)))
	if err != nil {
		return nil, err
	}

	// Convert into struct of Entity Type
	entityJSON, err := json.Marshal(entityValsByKey)
	if err != nil {
		return nil, fmt.Errorf("Error marshaling entity to JSON: %s", err)
	}

	err = json.Unmarshal(entityJSON, &entityObj)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling entity JSON: %s", err)
	}

	return entityObj.(Entity), nil
}

// TODO: Paging is not working here. Waiting on techops
// TODO: Add List for SearchID (similar to location-service). Follow up with Techops to see if SearchID is implemented
func (e *EntityService) ListAll(opts *EntityListOptions) ([]Entity, error) {
	var entities []Entity
	if opts == nil {
		opts = &EntityListOptions{}
	}
	opts.ListOptions = ListOptions{Limit: LocationListMaxLimit} // TODO: should this be EntityListMaxLimit
	var lg tokenListRetriever = func(listOptions *ListOptions) (string, error) {
		opts.ListOptions = *listOptions
		resp, _, err := e.List(opts)
		if err != nil {
			return "", err
		}

		typedEntities, err := e.toEntityTypes(resp.Entities)
		if err != nil {
			return "", err
		}
		for _, entity := range typedEntities {
			entities = append(entities, entity)
		}
		return resp.PageToken, err
	}

	if err := tokenListHelper(lg, &opts.ListOptions); err != nil {
		return nil, err
	} else {
		return entities, nil
	}
}

func (e *EntityService) List(opts *EntityListOptions) (*EntityListResponse, *Response, error) {
	var (
		requrl = entityPath
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

	v := &EntityListResponse{}
	r, err := e.client.DoRequest("GET", requrl, v)
	if err != nil {
		return nil, r, err
	}

	// TODO: nil is empty
	// for _, l := range v.Entities {
	// 	l.nilIsEmpty = true
	// }

	return v, nil, nil
}

func addEntityListOptions(requrl string, opts *EntityListOptions) (string, error) {
	if opts == nil {
		return requrl, nil
	}

	u, err := url.Parse(requrl)
	if err != nil {
		return "", err
	}

	q := u.Query()
	if opts.SearchID != "" {
		q.Add("searchId", opts.SearchID)
	}
	if opts.ResolvePlaceholders {
		q.Add("resolvePlaceholders", "true")
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}

func (e *EntityService) Get(id string) (Entity, *Response, error) {
	var v interface{}
	r, err := e.client.DoRequest("GET", fmt.Sprintf("%s/%s", entityPath, id), &v)
	if err != nil {
		return nil, r, err
	}

	entity, err := e.toEntityType(v)

	// TODO: nil is emtpy
	//v.nilIsEmpty = true

	return entity, r, nil
}

// TODO: Currently an error with API. Need to test this
func (e *EntityService) Create(y Entity) (*Response, error) {
	// TODO: custom field validation
	// if err := validateCustomFieldsKeys(y.CustomFields); err != nil {
	// 	return nil, err
	// }
	var requrl = entityPath
	u, err := url.Parse(requrl)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Add("entityType", string(y.GetEntityType()))
	u.RawQuery = q.Encode()
	r, err := e.client.DoRequestJSON("POST", u.String(), y, nil)
	if err != nil {
		return r, err
	}

	return r, nil
}

// TODO: There is an outstanding techops QA issue to allow the Id in the request but we may have to remove other things like account
func (e *EntityService) Edit(y Entity) (*Response, error) {
	// TODO: custom field validation
	// if err := validateCustomFieldsKeys(y.CustomFields); err != nil {
	// 	return nil, err
	// }
	r, err := e.client.DoRequestJSON("PUT", fmt.Sprintf("%s/%s", entityPath, y.GetEntityId()), y, nil)
	if err != nil {
		return r, err
	}

	return r, nil
}
