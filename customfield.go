package yext

// CustomField is the representation of a Custom Field definition in Yext Location Manager.
// For details see https://www.yext.com/support/platform-api/#Administration_API/Custom_Fields.htm
type CustomField struct {
	Name    string            `json:"name"`
	Id      string            `json:"id"`
	Options map[string]string `json:"options"` // Only present for multi-option custom fields
	Type    string            `json:"type"`
}
