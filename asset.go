package yext

type MappingType string

const (
	MappingTypeNone     MappingType = "NO_ENTITIES"
	MappingTypeAll      MappingType = "ALL_ENTITIES"
	MappingTypeFolder   MappingType = "FOLDER"
	MappingTypeEntities MappingType = "ENTITIES"
)

type UsageType string

const (
	UsageTypeProfileFields  UsageType = "PROFILE_FIELDS"
	UsageTypeReviewResponse UsageType = "REVIEW_RESPONSE"
	UsageTypeSocialPosting  UsageType = "SOCIAL_POSTING"
	UsageTypeAllFields      UsageType = "ALL_PROFILE_FIELDS"
)

type LabelOperator string

const (
	LabelOperatorAnd LabelOperator = "AND"
	LabelOperatorOr  LabelOperator = "OR"
)

type AssetType string

const (
	ASSETTYPE_TEXT         AssetType = "text"
	ASSETTYPE_IMAGE        AssetType = "image"
	ASSETTYPE_VIDEO        AssetType = "video"
	ASSETTYPE_COMPLEXIMAGE AssetType = "complexImage"
	ASSETTYPE_COMPLEXVIDEO AssetType = "complexVideo"
)

type AssetUsageType string

type Asset struct {
	Id          *string           `json:"id"`
	Name        *string           `json:"name"`
	Type        AssetType         `json:"type"`
	Description *string           `json:"description,omitempty"`
	ForEntities *ForEntities      `json:"forEntities,omitempty"`
	Usage       *[]AssetUsage     `json:"usage,omitempty"`
	Locale      *string           `json:"locale,omitempty"`
	Labels      *UnorderedStrings `json:"labels,omitempty"`
	Owner       *uint64           `json:"owner,omitempty"`
	Value       interface{}       `json:"value,omitempty"`
}

type ForEntities struct {
	MappingType   MappingType       `json:"mappingType"`
	FolderId      *string           `json:"folderId,omitempty"`
	EntityIds     *[]string         `json:"entityIds,omitempty"`
	LabelIds      *UnorderedStrings `json:"labelIds,omitempty"`
	LabelOperator *LabelOperator    `json:"labelOperator,omitempty"`
}

type AssetUsage struct {
	Type       UsageType `json:"type"`
	FieldNames *[]string `json:"fieldNames,omitempty"`
}

type AssetValue interface {
	GetAssetType() AssetType
}

type TextValue string

func (t TextValue) GetAssetType() AssetType {
	return ASSETTYPE_TEXT
}

type Image struct {
	Url           string  `json:"url"`
	AlternateText *string `json:"alternateText,omitempty"`
}

type ImageValue struct {
	Image *Image `json:"image,omitempty"`
}

func (i ImageValue) GetAssetType() AssetType {
	return ASSETTYPE_IMAGE
}

type ComplexImageValue struct {
	Image           *Image `json:"image,omitempty"`
	ClickthroughURL string `json:"clickthroughUrl,omitempty"`
	Details         string `json:"details,omitempty"`
}

func (ci ComplexImageValue) GetAssetType() AssetType {
	return ASSETTYPE_COMPLEXIMAGE
}

type VideoValue struct {
	Video *Video `json:"video,omitempty"`
}

func (i VideoValue) GetAssetType() AssetType {
	return ASSETTYPE_VIDEO
}

type ComplexVideoValue struct {
	Video           *Video `json:"image,omitempty"`
	ClickthroughURL string `json:"clickthroughUrl,omitempty"`
	Details         string `json:"details,omitempty"`
}

func (ci ComplexVideoValue) GetAssetType() AssetType {
	return ASSETTYPE_COMPLEXVIDEO
}
