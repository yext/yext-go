package yext

type EntityType string

// This is a little dangerous because if you embed an Entity (like Location) within another Entity
// You don't have to re-define these...but we'd need to re-define copy...
type Entity interface {
	GetEntityId() string
	GetEntityType() EntityType
	Copy() Entity
}

type EntityMeta struct {
	Id          *string           `json:"id,omitempty"`
	AccountId   *string           `json:"accountId,omitempty"`
	EntityType  EntityType        `json:"entityType,omitempty"`
	FolderId    *string           `json:"folderId,omitempty"`
	LabelIds    *UnorderedStrings `json:"labelIds,omitempty"`
	CategoryIds *[]string         `json:"categoryIds,omitempty"`
	Language    *string           `json:"language,omitempty"`
	CountryCode *string           `json:"countryCode,omitempty"`
	nilIsEmpty  bool
}
