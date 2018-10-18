package yext

import "fmt"

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

type CFTAsset struct {
	Id          *string           `json:"id,omitempty"`
	Name        *string           `json:"name,omitempty"`
	Type        AssetType         `json:"type,omitempty"`
	Description *string           `json:"description,omitempty"`
	ForEntities *ForEntities      `json:"forEntities,omitempty"`
	Usage       *AssetUsageList     `json:"usage,omitempty"`
	Locale      *string           `json:"locale,omitempty"`
	Labels      *UnorderedStrings `json:"labels,omitempty"`
	Owner       *uint64           `json:"owner,omitempty"`
	Value       interface{}       `json:"value,omitempty"`
}

type ForEntities struct {
	MappingType   MappingType       `json:"mappingType,omitempty"`
	FolderId      *string           `json:"folderId,omitempty"`
	EntityIds     *UnorderedStrings `json:"entityIds,omitempty"`
	LabelIds      *UnorderedStrings `json:"labelIds,omitempty"`
	LabelOperator LabelOperator     `json:"labelOperator,omitempty"`
}

type AssetUsage struct {
	Type       UsageType         `json:"type,omitempty"`
	FieldNames *UnorderedStrings `json:"fieldNames,omitempty"`
}

func (a *AssetUsage) Equal(b Comparable) bool {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Value of A: %+v, Value of B:%+v, Type Of A: %T, Type Of B: %T\n", a, b, a, b)
            panic(r)
        }
    }()

    bAssetUsage := *b.(*AssetUsage)
    if a.Type != bAssetUsage.Type || !a.FieldNames.Equal(bAssetUsage.FieldNames) {
        return false
    }

    return true
}

type AssetUsageList []AssetUsage

func (a *AssetUsageList) Equal(b Comparable) bool {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Value of A: %+v, Value of B:%+v, Type Of A: %T, Type Of B: %T\n", a, b, a, b)
            panic(r)
        }
    }()

    if a == nil || b == nil {
        return false
    }

    var (
        u = []AssetUsage(*a)
        s = []AssetUsage(*b.(*AssetUsageList))
    )
    if len(u) != len(s) {
        return false
    }

    for i := 0; i < len(u); i++ {
        var found bool
        for j := 0; j < len(s); j++ {
            if u[i].Equal(&s[j]) {
                found = true
            }
        }
        if !found {
            return false
        }
    }
    return true
}

type TextValue string

type ImageValue struct {
	Url           string `json:"url,omitempty"`
	AlternateText string `json:"alternateText,omitempty"`
	Height        uint64 `json:"height,omitempty"`
	Width         uint64 `json:"width,omitempty"`
}

type ComplexImageValue struct {
	Image           *ImageValue `json:"image,omitempty"`
	Description     string      `json:"description,omitempty"`
	Details         string      `json:"details,omitempty"`
	ClickthroughURL string      `json:"clickthroughUrl,omitempty"`
}

type VideoValue struct {
	Url string `json:"url,omitempty"`
}

type ComplexVideoValue struct {
	Video       *VideoValue `json:"video,omitempty"`
	Description string      `json:"description,omitempty"`
}

func (a *CFTAsset) GetId() string {
	if a.Id == nil {
		return ""
	}
	return *a.Id
}

func (a *CFTAsset) GetAssetType() AssetType {
	return a.Type
}
