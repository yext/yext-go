package yext

type LicensePack struct {
	LocationIds []string `json:"locationIds,omitempty"`
	Features    []string `json:"features,omitempty"`
	Id          int      `json:"id,omitempty"`
	Quantity    int      `json:"quantity,omitempty"`
	Status      string   `json:"status,omitempty"`
}
