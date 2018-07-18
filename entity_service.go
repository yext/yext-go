package yext

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
)

const entityPath = "entities"

type EntityRegistry map[EntityType]Entity

var YextEntityRegistry = EntityRegistry{
	ENTITYTYPE_EVENT:    &EventEntity{},
	ENTITYTYPE_LOCATION: &Location{},
}

type EntityService struct {
	client   *Client
	registry map[EntityType]Entity
}

type EntityListOptions struct {
	ListOptions
	SearchID            string
	ResolvePlaceholders bool
	EntityTypes         []EntityType
}

type EntityListResponse struct {
	Count         int        `json:"count"`
	Entities      EntityList `json:"entities"`
	NextPageToken string     `json:"nextPageToken"`
}

func (e *EntityService) RegisterEntity(entityType EntityType, entity Entity) {
	e.registry[entityType] = entity
}

func (e *EntityService) LookupEntity(entityType EntityType) (Entity, error) {
	entity, ok := e.registry[entityType]
	if !ok {
		return nil, fmt.Errorf("Unable to find entity type %s in entity registry %v", entityType, e.registry)
	}
	return entity, nil
}

func (e *EntityService) PathName(entityType EntityType) (string, error) {
	entity, err := e.LookupEntity(entityType)
	if err != nil {
		return "", err
	}
	return entity.PathName(), nil
}

func (e *EntityService) ListTest() ([]Entity, error) {
	jsonFile, err := os.Open("entitiesList.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var resp EntityListResponse
	err = json.Unmarshal(byteValue, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Entities, nil
}

func (e *EntityService) ListAll(opts *EntityListOptions) ([]Entity, error) {
	var entities []Entity
	if opts == nil {
		opts = &EntityListOptions{}
	}
	opts.ListOptions = ListOptions{Limit: LocationListMaxLimit}
	var lg tokenListRetriever = func(listOptions *ListOptions) (string, error) {
		opts.ListOptions = *listOptions
		resp, _, err := e.List(opts)
		if err != nil {
			return "", err
		}
		// for _, entity := range resp.Entities {
		// 	entities = append(entities, entity)
		// }
		return resp.NextPageToken, err
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

	// if _, err := e.HydrateLocations(v.Locations); err != nil {
	// 	return nil, r, err
	// }

	// for _, l := range v.Entities {
	// 	l.nilIsEmpty = true
	// }

	return v, r, nil
}

func (e *EntityService) ListAllOfType(opts *EntityListOptions, entityType EntityType) ([]Entity, error) {
	var entities []Entity
	if opts == nil {
		opts = &EntityListOptions{}
	}
	opts.ListOptions = ListOptions{Limit: LocationListMaxLimit}
	var lg tokenListRetriever = func(listOptions *ListOptions) (string, error) {
		opts.ListOptions = *listOptions
		resp, _, err := e.ListOfType(opts, entityType)
		if err != nil {
			return "", err
		}
		// for _, entity := range resp.Entities {
		// 	entities = append(entities, entity)
		// }
		return resp.NextPageToken, err
	}

	if err := tokenListHelper(lg, &opts.ListOptions); err != nil {
		return nil, err
	} else {
		return entities, nil
	}
}

func (e *EntityService) ListOfType(opts *EntityListOptions, entityType EntityType) (*EntityListResponse, *Response, error) {
	var (
		requrl string
		err    error
	)

	pathName, err := e.PathName(entityType)
	if err != nil {
		return nil, nil, err
	}
	requrl = pathName

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

	// if _, err := e.HydrateLocations(v.Locations); err != nil {
	// 	return nil, r, err
	// }

	// for _, l := range v.Entities {
	// 	l.nilIsEmpty = true
	// }

	return v, r, nil
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
	if opts.EntityTypes != nil && len(opts.EntityTypes) > 0 {
		// TODO: does this want a list?
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}

func (e *EntityService) Get(id string, entityType EntityType) (Entity, *Response, error) {
	entity, err := e.LookupEntity(entityType)
	if err != nil {
		return nil, nil, err
	}
	r, err := e.client.DoRequest("GET", fmt.Sprintf("%s/%s", entity.PathName(), id), entity)
	if err != nil {
		return nil, r, err
	}

	// if _, err := HydrateLocation(&v, l.CustomFields); err != nil {
	// 	return nil, r, err
	// }

	//v.nilIsEmpty = true

	return entity, r, nil
}

func (e *EntityService) Create(y Entity) (*Response, error) {
	// if err := validateCustomFieldsKeys(y.CustomFields); err != nil {
	// 	return nil, err
	// }
	r, err := e.client.DoRequestJSON("POST", fmt.Sprintf("%s", y.PathName()), y, nil)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (e *EntityService) Edit(y Entity) (*Response, error) {
	// if err := validateCustomFieldsKeys(y.CustomFields); err != nil {
	// 	return nil, err
	// }
	r, err := e.client.DoRequestJSON("PUT", fmt.Sprintf("%s/%s", y.PathName(), y.EntityId()), y, nil)
	if err != nil {
		return r, err
	}

	return r, nil
}

// func (e *EntityService) ListEntitiesBySearchId(searchId string, entityType EntityType) ([]Entity, error) {
// 	//var locations []*Location
// 	var llo = &LocationListOptions{SearchID: searchId}
// 	llo.ListOptions = ListOptions{Limit: LocationListMaxLimit}
// 	var lg tokenListRetriever = func(opts *ListOptions) (string, error) {
// 		llo.ListOptions = *opts
// 		llr, _, err := l.List(llo)
// 		if err != nil {
// 			return "", err
// 		}
// 		entities = append(entities, llr.Locations...)
// 		return llr.NextPageToken, err
// 	}
//
// 	if err := tokenListHelper(lg, &llo.ListOptions); err != nil {
// 		return nil, err
// 	} else {
// 		return entities, nil
// 	}
// }
