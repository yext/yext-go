package yext

import (
	"errors"
	"fmt"
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

type LocationListResponse struct {
	Count     int         `json:"count"`
	Locations []*Location `json:"locations"`
}

func (l *LocationService) ListAll() ([]*Location, error) {
	var locations []*Location
	var lg listRetriever = func(opts *ListOptions) (int, int, error) {
		llr, _, err := l.List(opts)
		if err != nil {
			return 0, 0, err
		}
		locations = append(locations, llr.Locations...)
		return len(llr.Locations), llr.Count, err
	}

	if err := listHelper(lg, &ListOptions{Limit: LocationListMaxLimit}); err != nil {
		return nil, err
	} else {
		return locations, nil
	}
}

func (l *LocationService) List(opts *ListOptions) (*LocationListResponse, *Response, error) {
	requrl, err := addListOptions(locationsPath, opts)
	if err != nil {
		return nil, nil, err
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

// TODO: Do searchids work in APIv2?
func (l *LocationService) ListBySearchIds(searchIds []string) ([]*Location, error) {
	v := &LocationListResponse{}
	_, err := l.client.DoRequest("GET", fmt.Sprintf("%s?searchIds=%s", locationsPath, strings.Join(searchIds, ",")), v)
	if err != nil {
		return nil, err
	}
	return l.HydrateLocations(v.Locations)
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
