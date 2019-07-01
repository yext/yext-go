package yext

import (
	"encoding/json"
)

type EntityType string

type Entity interface {
	GetEntityId() string
	GetEntityType() EntityType
}

type EntityMeta struct {
	Id          *string           `json:"id,omitempty"`
	AccountId   *string           `json:"accountId,omitempty"`
	EntityType  EntityType        `json:"entityType,omitempty"`
	FolderId    *string           `json:"folderId,omitempty"`
	Labels      *UnorderedStrings `json:"labels,omitempty"`
	Language    *string           `json:"language,omitempty"`
	CountryCode *string           `json:"countryCode,omitempty"`
}

type BaseEntity struct {
	Meta       *EntityMeta `json:"meta,omitempty"`
	nilIsEmpty bool
}

func (b *BaseEntity) GetEntityId() string {
	if b.Meta != nil && b.Meta.Id != nil {
		return *b.Meta.Id
	}
	return ""
}

func (b *BaseEntity) GetEntityType() EntityType {
	if b.Meta != nil {
		return b.Meta.EntityType
	}
	return ""
}

func (b *BaseEntity) GetFolderId() string {
	if b == nil || b.Meta == nil {
		return ""
	}
	if b.Meta.FolderId != nil {
		return *b.Meta.FolderId
	}
	return ""
}

func (b *BaseEntity) GetCountryCode() string {
	if b == nil || b.Meta == nil {
		return ""
	}
	if b.Meta.CountryCode != nil {
		return *b.Meta.CountryCode
	}
	return ""
}

// GetLabels returns a list of labels.
// Labels are stored in the system by ID, not by name
// Given a label "Example Label" with ID "123"
// This function will return a list containing ["123"]
func (b *BaseEntity) GetLabels() (v UnorderedStrings) {
	if b == nil || b.Meta == nil {
		return nil
	}
	if b.Meta.Labels != nil {
		v = *b.Meta.Labels
	}
	return v
}

// SetLabels takes a list of strings.
// Labels are stored in the system by ID, not by name
// This means that passing the name of the label will not work, labels must be passed by ID
// Given a label "Example Label" with ID "123"
// You must pass a list containing "123" for the Label "Example Label" to get set in the platform
func (b *BaseEntity) SetLabels(v []string) {
	l := UnorderedStrings(v)
	b.SetLabelsWithUnorderedStrings(l)
}

func (b *BaseEntity) SetLabelsWithUnorderedStrings(v UnorderedStrings) {
	if v != nil {
		b.Meta.Labels = &v
	}
}

func (b *BaseEntity) GetNilIsEmpty() bool {
	return b.nilIsEmpty
}

func (b *BaseEntity) SetNilIsEmpty(val bool) {
	b.nilIsEmpty = val
}

type RawEntity map[string]interface{}

func (r *RawEntity) GetEntityId() string {
	if m, ok := (*r)["meta"]; ok {
		meta := m.(map[string]interface{})
		if id, ok := meta["id"]; ok {
			return id.(string)
		}
	}
	return ""
}

func (r *RawEntity) GetEntityType() EntityType {
	if m, ok := (*r)["meta"]; ok {
		meta := m.(map[string]interface{})
		if t, ok := meta["entityType"]; ok {
			return EntityType(t.(string))
		}
	}
	return EntityType("")
}

func (r *RawEntity) GetLanguage() string {
	if m, ok := (*r)["meta"]; ok {
		meta := m.(map[string]interface{})
		if l, ok := meta["language"]; ok {
			return l.(string)
		}
	}
	return ""
}

func (r *RawEntity) GetAccountId() string {
	if m, ok := (*r)["meta"]; ok {
		meta := m.(map[string]interface{})
		if a, ok := meta["accountId"]; ok {
			return a.(string)
		}
	}
	return ""
}

func ConvertToRawEntity(e Entity) (*RawEntity, error) {
	var raw RawEntity
	m, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(m, &raw)
	if err != nil {
		return nil, err
	}
	return &raw, nil
}
