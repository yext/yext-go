package yext

import "fmt"

func HydrateLocation(loc *Location, customFields []*CustomField) (*Location, error) {
	if loc == nil || loc.CustomFields == nil || customFields == nil {
		return loc, nil
	}

	hydrated, err := ParseCustomFields(loc.CustomFields, customFields)
	if err != nil {
		return loc, fmt.Errorf("hydration failure: location: '%v' %v", loc.String(), err)
	}

	loc.CustomFields = hydrated
	loc.hydrated = true

	return loc, nil
}
