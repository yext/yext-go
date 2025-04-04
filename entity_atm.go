package yext

import (
	"encoding/json"
)

const ENTITYTYPE_ATM EntityType = "atm"

// ATM is the representation of an ATM in Yext's Knowledge Graph.
type ATMEntity struct {
	BaseEntity

	// Admin
	CategoryIds *[]string `json:"categoryIds,omitempty"`
	Closed      **bool    `json:"closed,omitempty"`

	// Address Fields
	Name          *string  `json:"name,omitempty"`
	Address       *Address `json:"address,omitempty"`
	AddressHidden **bool   `json:"addressHidden,omitempty"`
	ISORegionCode *string  `json:"isoRegionCode,omitempty"`
	Geomodifier   *string  `json:"geomodifier,omitempty"`

	// Other Contact Info
	AlternatePhone *string `json:"alternatePhone,omitempty"`
	Fax            *string `json:"fax,omitempty"`
	LocalPhone     *string `json:"localPhone,omitempty"`
	MobilePhone    *string `json:"mobilePhone,omitempty"`
	MainPhone      *string `json:"mainPhone,omitempty"`
	TollFreePhone  *string `json:"tollFreePhone,omitempty"`
	TtyPhone       *string `json:"ttyPhone,omitempty"`

	// Location Info
	Hours               **Hours `json:"hours,omitempty"`
	AdditionalHoursText *string `json:"additionalHoursText,omitempty"`
	Logo                **Photo `json:"logo,omitempty"`

	// Lats & Lngs
	DisplayCoordinate  **Coordinate `json:"displayCoordinate,omitempty"`
	DropoffCoordinate  **Coordinate `json:"dropoffCoordinate,omitempty"`
	WalkableCoordinate **Coordinate `json:"walkableCoordinate,omitempty"`
	RoutableCoordinate **Coordinate `json:"routableCoordinate,omitempty"`
	PickupCoordinate   **Coordinate `json:"pickupCoordinate,omitempty"`

	YextDisplayCoordinate  **Coordinate `json:"yextDisplayCoordinate,omitempty"`
	YextDropoffCoordinate  **Coordinate `json:"yextDropoffCoordinate,omitempty"`
	YextWalkableCoordinate **Coordinate `json:"yextWalkableCoordinate,omitempty"`
	YextRoutableCoordinate **Coordinate `json:"yextRoutableCoordinate,omitempty"`
	YextPickupCoordinate   **Coordinate `json:"yextPickupCoordinate,omitempty"`

	// Urls
	WebsiteUrl      **Website         `json:"websiteUrl,omitempty"`
	FeaturedMessage **FeaturedMessage `json:"featuredMessage,omitempty"`

	// Social Media
	FacebookPageUrl *string `json:"facebookPageUrl,omitempty"`
}

func (l *ATMEntity) UnmarshalJSON(data []byte) error {
	type Alias ATMEntity
	a := &struct {
		*Alias
	}{
		Alias: (*Alias)(l),
	}
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	return UnmarshalEntityJSON(l, data)
}

func (y ATMEntity) GetId() string {
	if y.BaseEntity.Meta != nil && y.BaseEntity.Meta.Id != nil {
		return *y.BaseEntity.Meta.Id
	}
	return ""
}

func (y ATMEntity) GetCategoryIds() (v []string) {
	if y.CategoryIds != nil {
		v = *y.CategoryIds
	}
	return v
}

func (y ATMEntity) GetName() string {
	if y.Name != nil {
		return GetString(y.Name)
	}
	return ""
}

func (y ATMEntity) GetAccountId() string {
	if y.BaseEntity.Meta != nil && y.BaseEntity.Meta.AccountId != nil {
		return *y.BaseEntity.Meta.AccountId
	}
	return ""
}

func (y ATMEntity) GetLine1() string {
	if y.Address != nil && y.Address.Line1 != nil {
		return GetString(y.Address.Line1)
	}
	return ""
}

func (y ATMEntity) GetLine2() string {
	if y.Address != nil && y.Address.Line2 != nil {
		return GetString(y.Address.Line2)
	}
	return ""
}

func (y ATMEntity) GetAddressHidden() bool {
	return GetNullableBool(y.AddressHidden)
}

func (y ATMEntity) GetExtraDescription() string {
	if y.Address != nil && y.Address.ExtraDescription != nil {
		return GetString(y.Address.ExtraDescription)
	}
	return ""
}

func (y ATMEntity) GetCity() string {
	if y.Address != nil && y.Address.City != nil {
		return GetString(y.Address.City)
	}
	return ""
}

func (y ATMEntity) GetRegion() string {
	if y.Address != nil && y.Address.Region != nil {
		return GetString(y.Address.Region)
	}
	return ""
}

func (y ATMEntity) GetCountryCode() string {
	if y.BaseEntity.Meta != nil && y.BaseEntity.Meta.CountryCode != nil {
		return GetString(y.BaseEntity.Meta.CountryCode)
	}
	return ""
}

func (y ATMEntity) GetPostalCode() string {
	if y.Address != nil && y.Address.PostalCode != nil {
		return GetString(y.Address.PostalCode)
	}
	return ""
}

func (y ATMEntity) GetMainPhone() string {
	if y.MainPhone != nil {
		return *y.MainPhone
	}
	return ""
}

func (y ATMEntity) GetLocalPhone() string {
	if y.LocalPhone != nil {
		return *y.LocalPhone
	}
	return ""
}

func (y ATMEntity) GetAlternatePhone() string {
	if y.AlternatePhone != nil {
		return *y.AlternatePhone
	}
	return ""
}

func (y ATMEntity) GetFax() string {
	if y.Fax != nil {
		return *y.Fax
	}
	return ""
}

func (y ATMEntity) GetMobilePhone() string {
	if y.MobilePhone != nil {
		return *y.MobilePhone
	}
	return ""
}

func (y ATMEntity) GetTollFreePhone() string {
	if y.TollFreePhone != nil {
		return *y.TollFreePhone
	}
	return ""
}

func (y ATMEntity) GetTtyPhone() string {
	if y.TtyPhone != nil {
		return *y.TtyPhone
	}
	return ""
}

func (y ATMEntity) GetFeaturedMessage() string {
	f := GetFeaturedMessage(y.FeaturedMessage)
	if f != nil {
		return GetString(f.Description)
	}
	return ""
}

func (y ATMEntity) GetFeaturedMessageUrl() string {
	f := GetFeaturedMessage(y.FeaturedMessage)
	if f != nil {
		return GetString(f.Url)
	}
	return ""
}

func (y ATMEntity) GetWebsiteUrl() string {
	w := GetWebsite(y.WebsiteUrl)
	if w != nil {
		return GetString(w.Url)
	}
	return ""
}

func (y ATMEntity) GetDisplayWebsiteUrl() string {
	w := GetWebsite(y.WebsiteUrl)
	if w != nil {
		return GetString(w.DisplayUrl)
	}
	return ""
}

func (y ATMEntity) GetHours() *Hours {
	return GetHours(y.Hours)
}

func (y ATMEntity) GetAdditionalHoursText() string {
	if y.AdditionalHoursText != nil {
		return *y.AdditionalHoursText
	}
	return ""
}

func (y ATMEntity) GetFacebookPageUrl() string {
	if y.FacebookPageUrl != nil {
		return *y.FacebookPageUrl
	}
	return ""
}

func (y ATMEntity) GetDisplayLat() float64 {
	c := GetCoordinate(y.DisplayCoordinate)
	if c != nil {
		return GetNullableFloat(c.Latitude)
	}
	return 0
}

func (y ATMEntity) GetDisplayLng() float64 {
	c := GetCoordinate(y.DisplayCoordinate)
	if c != nil {
		return GetNullableFloat(c.Longitude)
	}
	return 0
}

func (y ATMEntity) GetRoutableLat() float64 {
	c := GetCoordinate(y.RoutableCoordinate)
	if c != nil {
		return GetNullableFloat(c.Latitude)
	}
	return 0
}

func (y ATMEntity) GetRoutableLng() float64 {
	c := GetCoordinate(y.RoutableCoordinate)
	if c != nil {
		return GetNullableFloat(c.Longitude)
	}
	return 0
}

func (y ATMEntity) GetYextDisplayLat() float64 {
	c := GetCoordinate(y.YextDisplayCoordinate)
	if c != nil {
		return GetNullableFloat(c.Latitude)
	}
	return 0
}

func (y ATMEntity) GetYextDisplayLng() float64 {
	c := GetCoordinate(y.YextDisplayCoordinate)
	if c != nil {
		return GetNullableFloat(c.Longitude)
	}
	return 0
}

func (y ATMEntity) GetYextRoutableLat() float64 {
	c := GetCoordinate(y.YextRoutableCoordinate)
	if c != nil {
		return GetNullableFloat(c.Latitude)
	}
	return 0
}

func (y ATMEntity) GetYextRoutableLng() float64 {
	c := GetCoordinate(y.YextRoutableCoordinate)
	if c != nil {
		return GetNullableFloat(c.Longitude)
	}
	return 0
}

func (y ATMEntity) String() string {
	b, _ := json.Marshal(y)
	return string(b)
}

func (y ATMEntity) GetHolidayHours() []HolidayHours {
	h := GetHours(y.Hours)
	if h != nil && h.HolidayHours != nil {
		return *h.HolidayHours
	}
	return nil
}

func (y ATMEntity) IsClosed() bool {
	return GetNullableBool(y.Closed)
}
