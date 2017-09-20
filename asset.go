package yext

type AssetType string

const (
	AssetTypeText  AssetType = "AssetTypeText"
	AssetTypePhoto AssetType = "AssetTypePhoto"
	AssetTypeVideo AssetType = "AssetTypeVideo"
)

type LabelOperator string

const (
	LabelOperatorAnd LabelOperator = "AND"
	LabelOperatorOr  LabelOperator = "OR"
)

type Asset struct {
	Id                     string       `json:"id"`
	Name                   string       `json:"name"`
	Type                   AssetType    `json:"type"`
	ForLocations           ForLocations `json:"forLocations"`
	Description            string       `json:"description"`
	Labels                 []string     `json:"labels"`
	Contents               []Content    `json:"contents"`
	AssetTypePhotoUrl      string       `json:"AssetTypePhotoUrl"` // TODO: Is this right? Check documentation
	Details                string       `json:"details"`
	ClickthroughUrl        string       `json:"clickthroughUrl"`
	AlternateAssetTypeText string       `json:"alternateAssetTypeText"` // TODO: Is this right? Check documentation
	AssetTypeVideoUrl      string       `json:"AssetTypeVideoUrl"`      // TODO: Is this right? Check documentation
}

type ForLocations struct {
	MappingType   string        `json:"mappingType"`
	FolderId      string        `json:"folderId"`
	LocationIds   []string      `json:"locationIds"`
	LabelIds      []string      `json:"labelIds"`
	LabelOperator LabelOperator `json:"labelOperator"`
}

type Content struct {
	Content string `json:"content"`
	Locale  string `json:"locale"`
}
