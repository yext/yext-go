package yext

type CustomField struct {
	Name    string            `json:"name"`
	Id      string            `json:"id"`
	Options map[string]string `json:"options"` // Only for multi-option
	Type    string            `json:"type"`
}
