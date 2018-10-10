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

func (b *BaseEntity) GetNilIsEmpty() bool {
	return b.nilIsEmpty
}

func (b *BaseEntity) SetNilIsEmpty(val bool) {
	b.nilIsEmpty = val
}

func (b *BaseEntity) GetFolderId() string {
	if b.Meta.FolderId != nil {
		return *b.Meta.FolderId
	}
	return ""
}

func (b *BaseEntity) GetCategoryIds() (v []string) {
	if b.Meta.CategoryIds != nil {
		v = *b.Meta.CategoryIds
	}
	return v
}

func (b *BaseEntity) GetLabelIds() (v UnorderedStrings) {
	if b.Meta.LabelIds != nil {
		v = *b.Meta.LabelIds
	}
	return v
}

func (b *BaseEntity) SetLabelIds(v []string) {
	l := UnorderedStrings(v)
	b.SetLabelIdsWithUnorderedStrings(l)
}

func (b *BaseEntity) SetLabelIdsWithUnorderedStrings(v UnorderedStrings) {
	b.Meta.LabelIds = &v
}
