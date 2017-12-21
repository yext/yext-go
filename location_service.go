package yext

import (
	"fmt"
	"net/url"
	"regexp"
)

const locationsPath = "locations"

var (
	LocationListMaxLimit = 50

	customFieldKeyRegex = regexp.MustCompile("^[0-9]+$")
)

type LocationService struct {
	client       *Client
	CustomFields []*CustomField
}

type LocationListOptions struct {
	ListOptions
	SearchID            string
	ResolvePlaceholders bool
}

type LocationListResponse struct {
	Count         int         `json:"count"`
	Locations     []*Location `json:"locations"`
	NextPageToken string      `json:"nextPageToken"`
}

func (l *LocationService) ListAll(llopts *LocationListOptions) ([]*Location, error) {
	var locations []*Location
	if llopts == nil {
		llopts = &LocationListOptions{}
	}
	llopts.ListOptions = ListOptions{Limit: LocationListMaxLimit}
	var lg tokenListRetriever = func(opts *ListOptions) (string, error) {
		llopts.ListOptions = *opts
		llr, _, err := l.List(llopts)
		if err != nil {
			return "", err
		}
		locations = append(locations, llr.Locations...)
		return llr.NextPageToken, err
	}

	if err := tokenListHelper(lg, &llopts.ListOptions); err != nil {
		return nil, err
	} else {
		return locations, nil
	}
}

func (l *LocationService) List(llopts *LocationListOptions) (*LocationListResponse, *Response, error) {
	var (
		requrl string
		err    error
	)

	requrl = locationsPath

	if llopts != nil {
		requrl, err = addLocationListOptions(requrl, llopts)
		if err != nil {
			return nil, nil, err
		}
	}

	if llopts != nil {
		requrl, err = addListOptions(requrl, &llopts.ListOptions)
		if err != nil {
			return nil, nil, err
		}
	}

	v := &LocationListResponse{}
	r, err := l.client.DoRequest("GET", requrl, v)
	if err != nil {
		return nil, r, err
	}

	if _, err := l.HydrateLocations(v.Locations); err != nil {
		return nil, r, err
	}

	return v, r, nil
}

func addLocationListOptions(requrl string, opts *LocationListOptions) (string, error) {
	if opts == nil {
		return requrl, nil
	}

	u, err := url.Parse(requrl)
	if err != nil {
		return "", err
	}

	q := u.Query()
	if opts.SearchID != "" {
		q.Add("searchId", opts.SearchID)
	}
	if opts.ResolvePlaceholders {
		q.Add("resolvePlaceholders", "true")
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}

func addGetOptions(requrl string, opts *LocationListOptions) (string, error) {
	if opts == nil {
		return requrl, nil
	}

	u, err := url.Parse(requrl)
	if err != nil {
		return "", err
	}

	q := u.Query()
	if opts.ResolvePlaceholders {
		q.Add("resolvePlaceholders", "true")
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}

func (l *LocationService) Edit(y *Location) (*Response, error) {
	if err := validateCustomFields(y.CustomFields); err != nil {
		return nil, err
	}
	r, err := l.client.DoRequestJSON("PUT", fmt.Sprintf("%s/%s", locationsPath, y.GetId()), y, nil)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (l *LocationService) Create(y *Location) (*Response, error) {
	if err := validateCustomFields(y.CustomFields); err != nil {
		return nil, err
	}
	r, err := l.client.DoRequestJSON("POST", fmt.Sprintf("%s", locationsPath), y, nil)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (l *LocationService) Get(id string) (*Location, *Response, error) {
	var v Location
	r, err := l.client.DoRequest("GET", fmt.Sprintf("%s/%s", locationsPath, id), &v)
	if err != nil {
		return nil, r, err
	}

	if _, err := HydrateLocation(&v, l.CustomFields); err != nil {
		return nil, r, err
	}

	return &v, r, nil
}

func (l *LocationService) GetWithOptions(id string, llopts *LocationListOptions) (*Location, *Response, error) {
	var (
		requrl string
		err    error
		v      Location
	)
	requrl = fmt.Sprintf("%s/%s", locationsPath, id)
	if llopts != nil {
		requrl, err = addGetOptions(requrl, llopts)
		if err != nil {
			return nil, nil, err
		}
	}
	r, err := l.client.DoRequest("GET", requrl, &v)
	if err != nil {
		return nil, r, err
	}

	if _, err := HydrateLocation(&v, l.CustomFields); err != nil {
		return nil, r, err
	}

	return &v, r, nil
}

func (l *LocationService) ListBySearchId(searchId string) ([]*Location, error) {
	var locations []*Location
	var llo = &LocationListOptions{SearchID: searchId}
	llo.ListOptions = ListOptions{Limit: LocationListMaxLimit}
	var lg tokenListRetriever = func(opts *ListOptions) (string, error) {
		llo.ListOptions = *opts
		llr, _, err := l.List(llo)
		if err != nil {
			return "", err
		}
		locations = append(locations, llr.Locations...)
		return llr.NextPageToken, err
	}

	if err := tokenListHelper(lg, &llo.ListOptions); err != nil {
		return nil, err
	} else {
		return locations, nil
	}
}

// TODO: This mutates the locations, no need to return them
func (l *LocationService) HydrateLocations(locs []*Location) ([]*Location, error) {
	if l.CustomFields == nil {
		return locs, nil
	}

	for _, loc := range locs {
		_, err := HydrateLocation(loc, l.CustomFields)
		if err != nil {
			return locs, err
		}
	}

	return locs, nil
}
