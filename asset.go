package yext

const (
	AssetTypeText  AssetType = "AssetTypeText"
	AssetTypePhoto AssetType = "AssetTypePhoto"
	AssetTypeVideo AssetType = "AssetTypeVideo"
)

type Asset struct {
	Id              string       `json:"id"`
	Name            string       `json:"name"`
	Type            AssetType    `json:"type"`
	ForLocations    ForLocations `json:"forLocations"`
	Description     string       `json:"description"`
	Labels          []string     `json:"labels"`
	Contents        []Content    `json:"contents"`        // Type:Text
	PhotoUrl        string       `json:"photoUrl"`        // Type:Photo
	Details         string       `json:"details"`         // Type:Photo
	ClickthroughUrl string       `json:"clickthroughUrl"` // Type:Photo
	AlternateText   string       `json:"alternateText"`   // Type:Photo
	VideoUrl        string       `json:"videoUrl"`        // Type:Video
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
