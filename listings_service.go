package yext

import (
	"net/url"
	"strings"
)

const (
	listingsPath         = "listings"
	ListingsListMaxLimit = 100
)

type ListingsJSONResponse struct {
	Meta     Meta         `json:"meta"`
	Response ListingsData `json:"response"`
}

type AlternateBrands struct {
	BrandName  string `json:"brandName"`
	ListingURL string `json:"listingUrl"`
}

type Listing struct {
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
	Count     int       `json:"count"`
	Listings  []Listing `json:"listings"`
	PageToken string    `json:"pageToken"`
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
	EntityIds    []string `json:"entityIds"`
	Language     string   `json:"language"`
	PublisherIds []string `json:"publisherIds"`
	Statuses     []string `json:"statuses"`
}

// ListAll performs the API call outlined here
// https://hitchhikers.yext.com/docs/knowledgeapis/listings/listingsmanagement/listings/
func (l *ListingsService) ListAll(opts *ListingsListOptions) ([]Listing, error) {
	var (
		listings []Listing
	)

	if opts == nil {
		opts = &ListingsListOptions{}
	}

	opts.ListOptions = ListOptions{Limit: ListingsListMaxLimit}
	var lg tokenListRetriever = func(listOptions *ListOptions) (string, error) {
		opts.ListOptions = *listOptions
		resp, _, err := l.List(opts)
		if err != nil {
			return "", err
		}

		listings = append(listings, resp.Response.Listings...)

		return resp.Response.PageToken, nil
	}

	if err := tokenListHelper(lg, &opts.ListOptions); err != nil {
		return nil, err
	}
	return listings, nil
}

// List performs the API call outlined here
// https://hitchhikers.yext.com/docs/knowledgeapis/listings/listingsmanagement/listings/
func (l *ListingsService) List(opts *ListingsListOptions) (*ListingsJSONResponse, *Response, error) {
	var (
		requrl = listingsPath + "/listings"
		err    error
	)

	if opts != nil {
		requrl, err = addListingListOptions(requrl, opts)
		if err != nil {
			return nil, nil, err
		}
	}

	if opts != nil {
		requrl, err = addListOptions(requrl, &opts.ListOptions)
		if err != nil {
			return nil, nil, err
		}
	}

	v := &ListingsJSONResponse{}
	r, err := l.client.DoRequest("GET", requrl, v)
	if err != nil {
		return nil, r, err
	}

	return v, r, nil
}

// addListingListOptions adds options to query that are specific to the listings API
func addListingListOptions(requrl string, opts *ListingsListOptions) (string, error) {
	if opts == nil {
		return requrl, nil
	}

	u, err := url.Parse(requrl)
	if err != nil {
		return "", err
	}

	q := u.Query()
	if len(opts.EntityIds) > 0 {
		q.Add("entityIds", strings.Join(opts.EntityIds, ","))
	}
	if len(opts.Statuses) > 0 {
		q.Add("statuses", strings.Join(opts.Statuses, ","))
	}
	if len(opts.PublisherIds) > 0 {
		q.Add("publisherIds", strings.Join(opts.PublisherIds, ","))
	}
	if opts.Language != "" {
		q.Add("language", opts.Language)
	}

	u.RawQuery = q.Encode()

	return u.String(), nil

}
