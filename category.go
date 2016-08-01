package yext

// Category is a representation of a Category in Yext Location Manager.
// For details see https://www.yext.com/support/platform-api/#Info_API/Categories.htm
type Category struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	FullName   string `json:"fullName"`
	Selectable bool   `json:"selectable"`
	ParentId   string `json:"parentId"`
}
