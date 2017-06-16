package yext

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"
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
	SearchIDs []string
}

type LocationListResponse struct {
	Count     int         `json:"count"`
	Locations []*Location `json:"locations"`
}

func (l *LocationService) ListAll() ([]*Location, error) {
	var locations []*Location
	var llo = &LocationListOptions{}
	llo.ListOptions = ListOptions{Limit: LocationListMaxLimit}
	var lg listRetriever = func(opts *ListOptions) (int, int, error) {
		llo.ListOptions = *opts
		llr, _, err := l.List(llo)
		if err != nil {
			return 0, 0, err
		}
		locations = append(locations, llr.Locations...)
		return len(llr.Locations), llr.Count, err
	}

	if err := listHelper(lg, &llo.ListOptions); err != nil {
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
	if opts.SearchIDs != nil {
		q.Add("searchIds", strings.Join(opts.SearchIDs, ","))
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

	if _, err := l.HydrateLocation(&v); err != nil {
		return nil, r, err
	}

	return &v, r, nil
}

func (l *LocationService) ListBySearchIds(searchIds []string) ([]*Location, error) {
	var locations []*Location
	var llo = &LocationListOptions{SearchIDs: searchIds}
	llo.ListOptions = ListOptions{Limit: LocationListMaxLimit}
	var lg listRetriever = func(opts *ListOptions) (int, int, error) {
		llo.ListOptions = *opts
		llr, _, err := l.List(llo)
		if err != nil {
			return 0, 0, err
		}
		locations = append(locations, llr.Locations...)
		return len(llr.Locations), llr.Count, err
	}

	if err := listHelper(lg, &llo.ListOptions); err != nil {
		return nil, err
	} else {
		return locations, nil
	}
}

func (l *LocationService) HydrateLocation(loc *Location) (*Location, error) {
	if loc == nil || loc.CustomFields == nil || l.CustomFields == nil {
		return loc, nil
	}

	hydrated, err := ParseCustomFields(loc.CustomFields, l.CustomFields)
	if err != nil {
		return loc, fmt.Errorf("hydration failure: location: '%v' %v", loc.String(), err)
	}

	loc.CustomFields = hydrated
	loc.hydrated = true

	return loc, nil
}

func (l *LocationService) HydrateLocations(locs []*Location) ([]*Location, error) {
	if l.CustomFields == nil {
		return locs, nil
	}

	for _, loc := range locs {
		_, err := l.HydrateLocation(loc)
		if err != nil {
			return locs, err
		}
	}

	return locs, nil
}

func validateCustomFields(cfs map[string]interface{}) error {
	for k, _ := range cfs {
		if !customFieldKeyRegex.MatchString(k) {
			return errors.New(fmt.Sprintf("custom fields must be specified by their id, not name: %s", k))
		}
	}
	return nil
}
