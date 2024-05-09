package yext

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"

	"github.com/schollz/progressbar/v3"
)

const (
	entityPath         = "entities"
	EntityListMaxLimit = 50
)

type RichTextFormat int

const (
	RichTextFormatDefault RichTextFormat = iota
	RichTextFormatHTML
	RichTextFormatMarkdown
	RichTextFormatNone
)

func (r RichTextFormat) ToString() string {
	switch r {
	case RichTextFormatHTML:
		return "html"
	case RichTextFormatMarkdown:
		return "markdown"
	case RichTextFormatNone:
		return "none"
	default:
		return ""
	}
}

type EntityService struct {
	client         *Client
	Registry       *EntityRegistry
	nilBoolIsEmpty bool
}

type EntityListOptions struct {
	ListOptions
	SearchIDs             []string
	ResolvePlaceholders   bool
	Rendered              bool // will return resolved placeholders for language profiles
	EntityTypes           []string
	Fields                []string
	Filter                string
	Format                RichTextFormat
	HideProgressBar       bool
	ConvertRichTextToHTML bool
}

// Used for Create and Edit
type EntityServiceOptions struct {
	TemplateId              string   `json:"templateId,omitempty"`
	TemplateFields          []string `json:"templateFields,omitempty"`
	Format                  string   `json:"format,omitempty"`
	StripUnsupportedFormats bool     `json:"stripUnsupportedFormats,omitempty"`
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

func (e *EntityService) WithClient(c *Client) *EntityService {
	e.client = c
	return e
}

// TODO: Add List for SearchID (similar to location-service). Follow up with Techops to see if SearchID is implemented
func (e *EntityService) ListAll(opts *EntityListOptions) ([]Entity, error) {
	var (
		entities            []Entity
		totalCountRetrieved = false
		bar                 = progressbar.Default(-1)
	)

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

		//show progress bar if number of entities > 250
		if resp.Count > 250 && !opts.HideProgressBar {
			if !totalCountRetrieved {
				bar = progressbar.Default(int64(resp.Count))
				bar.Add(len(resp.typedEntites))
				totalCountRetrieved = true
			} else {
				bar.Add(len(resp.typedEntites))
			}
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
		if e.nilBoolIsEmpty {
			setNilBoolIsEmpty(entity)
		}
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
	if opts.Format != "" {
		q.Add("format", opts.Format)
	}
	if opts.StripUnsupportedFormats {
		q.Add("stripUnsupportedFormats", "true")
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
	if len(opts.Fields) > 0 {
		q.Add("fields", strings.Join(opts.Fields, ","))
	}
	if opts.ResolvePlaceholders {
		q.Add("resolvePlaceholders", "true")
	}
	if opts.Rendered {
		q.Add("rendered", "true")
	}
	if len(opts.EntityTypes) > 0 {
		q.Add("entityTypes", strings.Join(opts.EntityTypes, ","))
	}
	if len(opts.Filter) > 0 {
		q.Add("filter", opts.Filter)
	}
	if opts.Format != RichTextFormatDefault {
		q.Add("format", opts.Format.ToString())
	}
	if opts.ConvertRichTextToHTML {
		q.Add("convertRichTextToHTML", "true")
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
	if e.nilBoolIsEmpty {
		setNilBoolIsEmpty(entity)
	}

	return entity, r, nil
}

func setNilIsEmpty(i interface{}) {
	m := reflect.ValueOf(i).MethodByName("SetNilIsEmpty")
	if m.IsValid() {
		m.Call([]reflect.Value{reflect.ValueOf(true)})
	}
}

func setNilBoolIsEmpty(i interface{}) {
	m := reflect.ValueOf(i).MethodByName("SetNilBoolIsEmpty")
	if m.IsValid() {
		m.Call([]reflect.Value{reflect.ValueOf(true)})
	}
}

func GetNilIsEmpty(i interface{}) bool {
	m := reflect.ValueOf(i).MethodByName("GetNilIsEmpty")
	if m.IsValid() {
		values := m.Call([]reflect.Value{})
		if len(values) == 1 {
			return values[0].Interface().(bool)
		}
	}
	return false
}

func GetNilBoolIsEmpty(i interface{}) bool {
	m := reflect.ValueOf(i).MethodByName("GetNilBoolIsEmpty")
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

// Delete sends a request to the Knowledge Entities API to delete an entity with a given id
func (e *EntityService) Delete(id string) (*Response, error) {
	return e.client.DoRequestJSON("DELETE", fmt.Sprintf("%s/%s", entityPath, id), nil, nil)
}

func (e *EntityService) SetNilBoolIsEmpty(nilBoolIsEmpty bool) {
	e.nilBoolIsEmpty = nilBoolIsEmpty
}
