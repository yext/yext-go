package yext

import (
	"encoding/json"
)

const ENTITYTYPE_PROVIDERFACILITY EntityType = "providerFacility"

// ProviderFacilityEntity is the representation of a Provider Facility entity, which links a Healthcare Professional
// (healthcareProvider) to a Healthcare Facility (healthcareFacility) where that provider practices.
type ProviderFacilityEntity struct {
	BaseEntity

	// Linked entities. Each holds the entity ID of the linked Healthcare Facility / Healthcare Professional.
	HealthcareFacility *string `json:"healthcareFacility,omitempty"`
	HealthcareProvider *string `json:"healthcareProvider,omitempty"`

	// Admin
	CategoryIds *[]string `json:"categoryIds,omitempty"`
	Closed      **bool    `json:"closed,omitempty"`
	ComingSoon  **bool    `json:"comingSoon,omitempty"`

	// Address Fields
	Name          *string  `json:"name,omitempty"`
	OfficeName    *string  `json:"officeName,omitempty"`
	Address       *Address `json:"address,omitempty"`
	AddressHidden **bool   `json:"addressHidden,omitempty"`
	ISORegionCode *string  `json:"isoRegionCode,omitempty"`
	Geomodifier   *string  `json:"geomodifier,omitempty"`
	Neighborhood  *string  `json:"neighborhood,omitempty"`
	Timezone      *string  `json:"timezone,omitempty"`

	What3WordsAddress             *string `json:"what3WordsAddress,omitempty"`
	DeliverListingsWithoutGeocode **bool  `json:"deliverListingsWithoutGeocode,omitempty"`

	// Contact Info
	AlternatePhone *string `json:"alternatePhone,omitempty"`
	Fax            *string `json:"fax,omitempty"`
	LocalPhone     *string `json:"localPhone,omitempty"`
	MobilePhone    *string `json:"mobilePhone,omitempty"`
	MainPhone      *string `json:"mainPhone,omitempty"`
	TollFreePhone  *string `json:"tollFreePhone,omitempty"`

	// Hours
	Hours               **Hours `json:"hours,omitempty"`
	AdditionalHoursText *string `json:"additionalHoursText,omitempty"`

	// Lats & Lngs
	CityCoordinate     **Coordinate `json:"cityCoordinate,omitempty"`
	DisplayCoordinate  **Coordinate `json:"displayCoordinate,omitempty"`
	DropoffCoordinate  **Coordinate `json:"dropoffCoordinate,omitempty"`
	GeocodedCoordinate **Coordinate `json:"geocodedCoordinate,omitempty"`
	PickupCoordinate   **Coordinate `json:"pickupCoordinate,omitempty"`
	RoutableCoordinate **Coordinate `json:"routableCoordinate,omitempty"`
	WalkableCoordinate **Coordinate `json:"walkableCoordinate,omitempty"`

	YextDisplayCoordinate  **Coordinate `json:"yextDisplayCoordinate,omitempty"`
	YextDropoffCoordinate  **Coordinate `json:"yextDropoffCoordinate,omitempty"`
	YextPickupCoordinate   **Coordinate `json:"yextPickupCoordinate,omitempty"`
	YextRoutableCoordinate **Coordinate `json:"yextRoutableCoordinate,omitempty"`
	YextWalkableCoordinate **Coordinate `json:"yextWalkableCoordinate,omitempty"`

	// Urls
	ReservationUrl **Website `json:"reservationUrl,omitempty"`
	WebsiteUrl     **Website `json:"websiteUrl,omitempty"`

	// Social Media / Google / Facebook
	FacebookDescriptor   *string `json:"facebookDescriptor,omitempty"`
	FacebookOverrideCity *string `json:"facebookOverrideCity,omitempty"`
	FacebookStoreId      *string `json:"facebookStoreId,omitempty"`
	FacebookVanityUrl    *string `json:"facebookVanityUrl,omitempty"`

	GoogleAccountId          *string                    `json:"googleAccountId,omitempty"`
	GoogleAttributes         *map[string][]string       `json:"googleAttributes,omitempty"`
	GoogleCid                *string                    `json:"googleCid,omitempty"`
	GoogleEntityRelationship **GoogleEntityRelationship `json:"googleEntityRelationship,omitempty"`
	GoogleMyBusinessLabels   *UnorderedStrings          `json:"googleMyBusinessLabels,omitempty"`
	GooglePlaceId            *string                    `json:"googlePlaceId,omitempty"`
	GoogleWebsiteOverride    **string                   `json:"googleWebsiteOverride,omitempty"`

	// Scout
	ScoutCustomKeywords *UnorderedStrings `json:"scoutCustomKeywords,omitempty"`
}

func (p *ProviderFacilityEntity) UnmarshalJSON(data []byte) error {
	type Alias ProviderFacilityEntity
	a := &struct {
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	return UnmarshalEntityJSON(p, data)
}

func (y ProviderFacilityEntity) GetId() string {
	if y.BaseEntity.Meta != nil && y.BaseEntity.Meta.Id != nil {
		return *y.BaseEntity.Meta.Id
	}
	return ""
}

func (y ProviderFacilityEntity) String() string {
	b, _ := json.Marshal(y)
	return string(b)
}

func (y ProviderFacilityEntity) GetName() string {
	if y.Name != nil {
		return *y.Name
	}
	return ""
}

func (y ProviderFacilityEntity) GetHealthcareFacility() string {
	if y.HealthcareFacility != nil {
		return *y.HealthcareFacility
	}
	return ""
}

func (y ProviderFacilityEntity) GetHealthcareProvider() string {
	if y.HealthcareProvider != nil {
		return *y.HealthcareProvider
	}
	return ""
}

func (y ProviderFacilityEntity) GetCategoryIds() (v []string) {
	if y.CategoryIds != nil {
		v = *y.CategoryIds
	}
	return v
}

func (y ProviderFacilityEntity) GetGooglePlaceId() string {
	if y.GooglePlaceId != nil {
		return *y.GooglePlaceId
	}
	return ""
}

func (y ProviderFacilityEntity) GetGoogleAttributes() map[string][]string {
	if y.GoogleAttributes != nil {
		return *y.GoogleAttributes
	}
	return nil
}

func (y ProviderFacilityEntity) GetHolidayHours() []HolidayHours {
	h := GetHours(y.Hours)
	if h != nil && h.HolidayHours != nil {
		return *h.HolidayHours
	}
	return nil
}

func (y ProviderFacilityEntity) IsClosed() bool {
	return GetNullableBool(y.Closed)
}

func (y ProviderFacilityEntity) IsComingSoon() bool {
	return GetNullableBool(y.ComingSoon)
}
