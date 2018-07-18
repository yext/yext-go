package yext

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
