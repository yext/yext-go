package yext

type ListingsJSONResponse struct {
	Meta     Meta         `json:"meta"`
	Response ListingsData `json:"response"`
}

type AlternateBrands struct {
	BrandName  string `json:"brandName"`
	ListingURL string `json:"listingUrl"`
}

type Listings struct {
	ID               string             `json:"id"`
	LocationID       string             `json:"locationId"`
	AccountID        string             `json:"accountId"`
	PublisherID      string             `json:"publisherId"`
	Status           string             `json:"status"`
	AdditionalStatus string             `json:"additionalStatus,omitempty"`
	ListingURL       string             `json:"listingUrl"`
	ScreenshotURL    string             `json:"screenshotUrl"`
	AlternateBrands  *[]AlternateBrands `json:"alternateBrands"`
	LoginURL         string             `json:"loginUrl,omitempty"`
}

type ListingsData struct {
	Count    int        `json:"count"`
	Listings []Listings `json:"listings"`
}

type TokenResponseObject struct {
	AccessToken string `json:"access_token"`
	InstanceURL string `json:"instance_url"`
	Id          string `json:"id"`
	TokenType   string `json:"token_type"`
	IssuedAt    string `json:"issued_at"`
	Signature   string `json:"signature"`
}

type ListingsService struct {
	client *Client
}

type ListingsListOptions struct {
	ListOptions
}
