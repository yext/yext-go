package yext

// Folder is a representation of a Folder in Yext Location Manager.
// For details see https://www.yext.com/support/platform-api/#Administration_API/Folders.htm
type Folder struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ParentId string `json:"parentId"`
}
