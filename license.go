package yext

type LicensePack struct {
	LocationIds []string `json:"locationIds"`
	Features    []string `json:"features"`
	Id          int      `json:"id"`
	Quantity    int      `json:"quantity"`
	Status      string   `json:"status"`
}
