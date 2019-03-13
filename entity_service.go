package yext

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

const (
	entityPath         = "entities"
	EntityListMaxLimit = 50
)

type EntityService struct {
	client   *Client
	Registry *EntityRegistry
}

type EntityListOptions struct {
	ListOptions
	SearchIDs           []string
	ResolvePlaceholders bool
	EntityTypes         []string
}

// Used for Create and Edit
type EntityServiceOptions struct {
	TemplateId     string   `json:"templateId,omitempty"`
	TemplateFields []string `json:"templateFields,omitempty"`
}

type EntityListResponse struct {
	Count        int           `json:"count"`
	Entities     []interface{} `json:"entities"`
	typedEntites []Entity
	PageToken    string `json:"pageToken"`
}

func (e *EntityService) RegisterDefaultEntities() {
	e.Registry = defaultEntityRegistry()
}

func (e *EntityService) RegisterEntity(t EntityType, entity interface{}) {
	e.Registry.RegisterEntity(t, entity)
}

func (e *EntityService) InitializeEntity(t EntityType) (Entity, error) {
	return e.Registry.InitializeEntity(t)
}

func (e *EntityService) ToEntityTypes(entities []interface{}) ([]Entity, error) {
	return e.Registry.ToEntityTypes(entities)
}

func (e *EntityService) ToEntityType(entity interface{}) (Entity, error) {
	return e.Registry.ToEntityType(entity)
}

// TODO: Add List for SearchID (similar to location-service). Follow up with Techops to see if SearchID is implemented
func (e *EntityService) ListAll(opts *EntityListOptions) ([]Entity, error) {
	var entities []Entity
	if opts == nil {
		opts = &EntityListOptions{}
	}
	opts.ListOptions = ListOptions{Limit: EntityListMaxLimit}
	var lg tokenListRetriever = func(listOptions *ListOptions) (string, error) {
		opts.ListOptions = *listOptions
		resp, _, err := e.List(opts)
		if err != nil {
			return "", err
		}

		entities = append(entities, resp.typedEntites...)
		return resp.PageToken, nil
	}

	if err := tokenListHelper(lg, &opts.ListOptions); err != nil {
		return nil, err
	}
	return entities, nil
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

	typedEntities, err := e.ToEntityTypes(v.Entities)
	if err != nil {
		return nil, r, err
	}
	entities := []Entity{}
	for _, entity := range typedEntities {
		setNilIsEmpty(entity)
		entities = append(entities, entity)
	}
	v.typedEntites = entities
	return v, r, nil
}

func addEntityServiceOptions(requrl string, opts *EntityServiceOptions) (string, error) {
	u, err := url.Parse(requrl)
	if err != nil {
		return "", err
	}

	if opts == nil {
		return requrl, nil
	}

	q := u.Query()
	if opts.TemplateId != "" {
		q.Add("templateId", opts.TemplateId)
	}
	if len(opts.TemplateFields) > 0 {
		q.Add("templateFields", strings.Join(opts.TemplateFields, ","))
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
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
	if len(opts.SearchIDs) > 0 {
		q.Add("searchIds", strings.Join(opts.SearchIDs, ","))
	}
	if opts.ResolvePlaceholders {
		q.Add("resolvePlaceholders", "true")
	}
	if len(opts.EntityTypes) > 0 {
		q.Add("entityTypes", strings.Join(opts.EntityTypes, ","))
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}

func (e *EntityService) Get(id string) (Entity, *Response, error) {
	var v map[string]interface{}
	r, err := e.client.DoRequest("GET", fmt.Sprintf("%s/%s", entityPath, id), &v)
	if err != nil {
		return nil, r, err
	}

	entity, err := e.ToEntityType(v)
	if err != nil {
		return nil, r, err
	}

	setNilIsEmpty(entity)

	return entity, r, nil
}

func setNilIsEmpty(i interface{}) {
	m := reflect.ValueOf(i).MethodByName("SetNilIsEmpty")
	if m.IsValid() {
		m.Call([]reflect.Value{reflect.ValueOf(true)})
	}
}

func getNilIsEmpty(i interface{}) bool {
	m := reflect.ValueOf(i).MethodByName("GetNilIsEmpty")
	if m.IsValid() {
		values := m.Call([]reflect.Value{})
		if len(values) == 1 {
			return values[0].Interface().(bool)
		}
	}
	return false
}

func (e *EntityService) Create(y Entity) (*Response, error) {
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

func (e *EntityService) CreateWithOptions(y Entity, opts *EntityServiceOptions) (*Response, error) {
	var (
		requrl = entityPath
		err    error
	)

	requrl, err = addEntityServiceOptions(requrl, opts)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(requrl)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Add("entityType", string(y.GetEntityType()))
	u.RawQuery = q.Encode()
	return e.client.DoRequestJSON("POST", u.String(), y, nil)
}

func (e *EntityService) EditWithOptions(y Entity, id string, opts *EntityServiceOptions) (*Response, error) {
	var (
		requrl = fmt.Sprintf("%s/%s", entityPath, id)
		err    error
	)

	requrl, err = addEntityServiceOptions(requrl, opts)
	if err != nil {
		return nil, err
	}

	r, err := e.client.DoRequestJSON("PUT", requrl, y, nil)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (e *EntityService) EditWithId(y Entity, id string) (*Response, error) {
	return e.client.DoRequestJSON("PUT", fmt.Sprintf("%s/%s", entityPath, id), y, nil)
}

func (e *EntityService) Edit(y Entity) (*Response, error) {
	return e.EditWithId(y, y.GetEntityId())
}

// DeleteWithId sends a request to the Knowledge Entities API to delete an entity with a given id
func (e *EntityService) DeleteWithId(id string) (*Response, error) {
	return e.client.DoRequestJSON("DELETE", fmt.Sprintf("%s/%s", entityPath, id), nil, nil)
}
