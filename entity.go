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
	v := r.GetValue([]string{"meta", "id"})
	if v == nil {
		return ""
	}
	return v.(string)
}

func (r *RawEntity) GetEntityType() EntityType {
	v := r.GetValue([]string{"meta", "entityType"})
	if v == nil {
		return ""
	}
	if _, ok := v.(string); ok {
		return EntityType(v.(string))
	}
	if _, ok := v.(EntityType); ok {
		return v.(EntityType)
	}
	return ""
}

func (r *RawEntity) GetLanguage() string {
	v := r.GetValue([]string{"meta", "language"})
	if v == nil {
		return ""
	}
	return v.(string)
}

func (r *RawEntity) GetAccountId() string {
	v := r.GetValue([]string{"meta", "accountId"})
	if v == nil {
		return ""
	}
	return v.(string)
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

func (r *RawEntity) IsZeroValue() bool {
    m := (map[string]interface{})(*r)
    return rawEntityIsZeroValue(m)
}

// Function that recursively checks each field on a RawEntity to determine if the entity overall is empty or not
// keys with nil values should still cause this func to return true
func rawEntityIsZeroValue(rawEntity map[string]interface{}) bool {
    for _, v := range rawEntity {
        if v != nil {
            // if v is a map, it means we've got a nested field, so we need to recurse to check the subfields of that map too
            if m, ok := v.(map[string]interface{}); ok {
                return rawEntityIsZeroValue(m)
            } else {
                // if the value we're looking at isn't nil AND isn't a map (nested field), we can assume this field on
                // the raw entity has some value, therefore the whole entity is not zero value
                return false
            }
        }
    }

    // if we never encounter a field with a value in the above, it's an empty rawEntity so we can return true
    return true
}

func (r *RawEntity) GetValue(keys []string) interface{} {
	if len(keys) == 0 {
		return nil
	}
	m := (map[string]interface{})(*r)
	return getValue(m, keys, 0)
}

func getValue(m map[string]interface{}, keys []string, index int) interface{} {
	for k, v := range m {
		if k == keys[index] {
			if index == len(keys)-1 {
				return v
			} else {
				return getValue((v).(map[string]interface{}), keys, index+1)
			}
		}
	}
	return nil
}

func (r *RawEntity) SetValue(keys []string, val interface{}) error {
	if len(keys) == 0 {
		return nil
	}
	m := (map[string]interface{})(*r)
	v, err := setValue(m, keys, 0, val)
	if err != nil {
		return err
	}
	raw := RawEntity(v)
	r = &raw
	return nil
}

func setValue(m map[string]interface{}, keys []string, index int, val interface{}) (map[string]interface{}, error) {
	if m == nil {
		return nil, nil
	}
	if index == len(keys)-1 {
		m[keys[index]] = val
	} else {
		if _, ok := m[keys[index]]; !ok {
			m[keys[index]] = map[string]interface{}{}
		}
		v, err := setValue(m[keys[index]].(map[string]interface{}), keys, index+1, val)
		if err != nil {
			return v, err
		}
		m[keys[index]] = v
	}

	return m, nil
}

func (r *RawEntity) RemoveEntry(keys []string) error {
	if len(keys) == 0 {
		return nil
	}
	m := (map[string]interface{})(*r)
	v, err := removeEntry(m, keys, 0)
	if err != nil {
		return err
	}
	raw := RawEntity(v)
	r = &raw
	return nil
}

func removeEntry(m map[string]interface{}, keys []string, index int) (map[string]interface{}, error) {
	if m == nil {
		return nil, nil
	}
	if index == len(keys)-1 {
		delete(m, keys[index])
	} else {
		if _, ok := m[keys[index]]; !ok {
			m[keys[index]] = map[string]interface{}{}
		}
		v, err := removeEntry(m[keys[index]].(map[string]interface{}), keys, index+1)
		if err != nil {
			return v, err
		}
		m[keys[index]] = v
	}

	return m, nil
}

func ConvertStringsToEntityTypes(types []string) []EntityType {
	entityTypes := []EntityType{}
	for _, stringType := range types {
		entityTypes = append(entityTypes, EntityType(stringType))
	}

	return entityTypes
}
