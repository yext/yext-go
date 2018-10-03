package yext

const (
	AssetTypeText         = "text"
	AssetTypeImage        = "image"
	AssetTypeVideo        = "video"
	AssetTypeComplexImage = "complexImage"
	AssetTypeComplexVideo = "complexVideo"

	AssetMappingTypeNone     = "NO_ENTITIES"
	AssetMappingTypeAll      = "ALL_ENTITIES"
	AssetMappingTypeFolder   = "FOLDER"
	AssetMappingTypeEntities = "ENTITIES"

	AssetUsageTypeProfileFields  = "PROFILE_FIELDS"
	AssetUsageTypeReviewResponse = "REVIEW_RESPONSE"
	AssetUsageTypeSocialPosting  = "SOCIAL_POSTING"
	AssetUsageTypeAllFields      = "ALL_PROFILE_FIELDS"
)

type LabelOperator string

const (
	LabelOperatorAnd LabelOperator = "AND"
	LabelOperatorOr  LabelOperator = "OR"
)

type AssetUsageType string

type Asset struct {
	Id          *string       `json:"id"`
	Name        *string       `json:"name"`
	Type        *string       `json:"type"`
	Description *string       `json:"description,omitempty"`
	ForEntities *ForEntities  `json:"forEntities,omitempty"`
	Usage       *[]AssetUsage `json:"usage,omitempty"`
	Locale      *string       `json:"locale,omitempty"`
	Labels      *[]string     `json:"labels,omitempty"`
	Owner       *uint64       `json:"owner,omitempty"`
	Value       *AssetValue   `json:"value,omitempty"`
}

type ForEntities struct {
	MappingType   *string        `json:"mappingType"`
	FolderId      *string        `json:"folderId,omitempty"`
	EntityIds     *[]string      `json:"entityIds,omitempty"`
	LabelIds      *[]string      `json:"labelIds,omitempty"`
	LabelOperator *LabelOperator `json:"labelOperator,omitempty"`
}

type AssetValue struct {
	Text  *TextValue  `json:"text,omitempty"`
	Image *ImageValue `json:"image,omitempty"`
}

type AssetUsage struct {
	Type       *string   `json:"type"`
	FieldNames *[]string `json:"fieldNames"`
}

type TextValue string

type ImageValue struct {
	Url             *string `json:"url"`
	AlternateText   *string `json:"alternateText,omitempty"`
	ClickthroughURL *string `json:"clickthroughUrl,omitempty"`
	Details         *string `json:"details,omitempty"`
}
