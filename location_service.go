package yext

import (
	"errors"
	"fmt"
	"regexp"
)

const locationsPath = "locations"

var customFieldKeyRegex = regexp.MustCompile("^[0-9]+$")

type LocationService struct {
	client *Client
}

type locationListResponse struct {
	Locations []*Location `json:"locations"`
}

func (l *LocationService) List() ([]*Location, error) {
	v := &locationListResponse{}
	err := l.client.DoRequest("GET", locationsPath, v)
	return v.Locations, err
}

func (l *LocationService) Edit(y *Location) (*Location, error) {
	if err := validateCustomFields(y.CustomFields); err != nil {
		return nil, err
	}
	var v Location
	err := l.client.DoRequestJSON("PUT", fmt.Sprintf("%s/%s", locationsPath, y.GetId()), y, &v)
	return &v, err
}

func (l *LocationService) Create(y *Location) (*Location, error) {
	if err := validateCustomFields(y.CustomFields); err != nil {
		return nil, err
	}
	var v Location
	err := l.client.DoRequestJSON("POST", fmt.Sprintf("%s", locationsPath), y, &v)
	return &v, err
}

func (l *LocationService) Get(id string) (*Location, error) {
	var v Location
	err := l.client.DoRequest("GET", fmt.Sprintf("%s/%s", locationsPath, id), &v)
	return &v, err
}

func (l *LocationService) ListBySearchIds(searchIds []string) ([]*Location, error) {
	v := &locationListResponse{}
	err := l.client.DoRequest("GET", fmt.Sprintf("%s?searchIds=%s", locationsPath, strings.Join(searchIds, ",")), v)
	return v.Locations, err
}

func validateCustomFields(cfs map[string]interface{}) error {
	for k, _ := range cfs {
		if !customFieldKeyRegex.Match([]byte(k)) {
			return errors.New(fmt.Sprintf("custom fields must be specified by their id, not name: %s", k))
		}
	}
	return nil
}
