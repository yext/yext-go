package yext

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
	LabelIds    *UnorderedStrings `json:"labelIds,omitempty"`
	CategoryIds *[]string         `json:"categoryIds,omitempty"`
	Language    *string           `json:"language,omitempty"`
	CountryCode *string           `json:"countryCode,omitempty"`
	// TODO: See if we still need and implement
	nilIsEmpty bool
}

type BaseEntity struct {
	Meta *EntityMeta `json:"meta,omitempty"`
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
