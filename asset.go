package yext

type AssetType string

const (
	TEXT  AssetType = "TEXT"
	PHOTO AssetType = "PHOTO"
	VIDEO AssetType = "VIDEO"
)

type LabelOperatorType string

const (
	AND AssetType = "AND"
	OR  AssetType = "OR"
)

type Asset struct {
	Id              string       `json:"id"`
	Name            string       `json:"name"`
	Type            AssetType    `json:"type"`
	ForLocations    ForLocations `json:"forLocations"`
	Description     string       `json:"description"`
	Labels          []string     `json:"labels"`
	Contents        []Content    `json:"contents"`
	PhotoUrl        string       `json:"photoUrl"`
	Details         string       `json:"details"`
	ClickthroughUrl string       `json:"clickthroughUrl"`
	AlternateText   string       `json:"alternateText"`
	VideoUrl        string       `json:"videoUrl"`
}

type ForLocations struct {
	MappingType   string            `json:"mappingType"`
	FolderId      string            `json:"folderId"`
	LocationIds   []string          `json:"locationIds"`
	LabelIds      []string          `json:"labelIds"`
	LabelOperator LabelOperatorType `json:"labelOperator"`
}

type Content struct {
	Content string `json:"content"`
	Locale  string `json:"locale"`
}
