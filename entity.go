package yext

import (
	"encoding/json"
	"fmt"
)

type EntityType string

type Entity interface {
	EntityId() string // Would this be necessary if it's always in the Core?
	Type() EntityType
	PathName() string
}

type EntityMeta struct {
	Id          *string           `json:"id,omitempty"`
	AccountId   *string           `json:"accountId,omitempty"`
	EntityType  *EntityType       `json:"locationType,omitempty"`
	FolderId    *string           `json:"folderId,omitempty"`
	LabelIds    *UnorderedStrings `json:"labelIds,omitempty"`
	CategoryIds *[]string         `json:"categoryIds,omitempty"`
	Language    *string           `json:"language,omitempty"`
	CountryCode *string           `json:"countryCode,omitempty"`
}

type EntityList []Entity

func (e *EntityList) UnmarshalJSON(b []byte) error {
	var raw []json.RawMessage
	err := json.Unmarshal(b, &raw)
	if err != nil {
		return err
	}

	for _, r := range raw {
		var obj map[string]interface{}
		err := json.Unmarshal(r, &obj)
		if err != nil {
			return err
		}

		var entityType string
		if t, ok := obj["entityType"].(string); ok {
			entityType = t
		}

		typedEntity, ok := YextEntityRegistry[EntityType(entityType)]
		if !ok {
			return fmt.Errorf("Entity type %s is not in the registry", entityType)
		}

		err = json.Unmarshal(r, typedEntity)
		if err != nil {
			return err
		}
		*e = append(*e, typedEntity)

	}
	return nil

}
