package yext

import (
	"encoding/json"
)

const ENTITYTYPE_HEALTHCAREFACILITY EntityType = "healthcareFacility"

// Location is the representation of a Location in Yext Location Manager.
// For details see https://www.yext.com/support/platform-api/#Administration_API/Locations.htm

type HealthcareFacilityEntity struct {
	BaseEntity

	// Admin
	CategoryIds *[]string `json:"categoryIds,omitempty"`
	Closed      **bool    `json:"closed,omitempty"`
	Keywords    *[]string `json:"keywords,omitempty"`
	Slug        *string   `json:"slug,omitempty"`

	// Address Fields
	Name          *string  `json:"name,omitempty"`
	Address       *Address `json:"address,omitempty"`
	AddressHidden **bool   `json:"addressHidden,omitempty"`
	ISORegionCode *string  `json:"isoRegionCode,omitempty"`
	Geomodifier   *string  `json:"geomodifier,omitempty"`

	// Other Contact Info
	AlternatePhone *string   `json:"alternatePhone,omitempty"`
	Fax            *string   `json:"fax,omitempty"`
	LocalPhone     *string   `json:"localPhone,omitempty"`
	MobilePhone    *string   `json:"mobilePhone,omitempty"`
	MainPhone      *string   `json:"mainPhone,omitempty"`
	TollFreePhone  *string   `json:"tollFreePhone,omitempty"`
	TtyPhone       *string   `json:"ttyPhone,omitempty"`
	Emails         *[]string `json:"emails,omitempty"`

	// Location Info
	Description         *string   `json:"description,omitempty"`
	Hours               **Hours   `json:"hours,omitempty"`
	AdditionalHoursText *string   `json:"additionalHoursText,omitempty"`
	YearEstablished     **float64 `json:"yearEstablished,omitempty"`
	Associations        *[]string `json:"associations,omitempty"`
	Certifications      *[]string `json:"certifications,omitempty"`
	Brands              *[]string `json:"brands,omitempty"`
	Products            *[]string `json:"products,omitempty"`
	Services            *[]string `json:"services,omitempty"`
	// Spelling of json tag 'specialities' is intentional to match mispelling in Yext API
	Specialties          *[]string `json:"specialities,omitempty"`
	Languages            *[]string `json:"languages,omitempty"`
	Logo                 **Photo   `json:"logo,omitempty"`
	PaymentOptions       *[]string `json:"paymentOptions,omitempty"`
	InsuranceAccepted    *[]string `json:"insuranceAccepted,omitempty"`
	AcceptingNewPatients **bool    `json:"acceptingNewPatients,omitempty"`
	NPI                  *string   `json:"npi,omitempty"`

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

	// Lists
	Bios         **Lists `json:"bios,omitempty"`
	Calendars    **Lists `json:"calendars,omitempty"`
	Menus        **Lists `json:"menus,omitempty"`
	ProductLists **Lists `json:"productLists,omitempty"`

	// Urls
	MenuUrl               **Website         `json:"menuUrl,omitempty"`
	OrderUrl              **Website         `json:"orderUrl,omitempty"`
	ReservationUrl        **Website         `json:"reservationUrl,omitempty"`
	TelehealthUrl         *string           `json:"telehealthUrl,omitempty"`
	WebsiteUrl            **Website         `json:"websiteUrl,omitempty"`
	FeaturedMessage       **FeaturedMessage `json:"featuredMessage,omitempty"`
	COVID19InformationUrl *string           `json:"covid19InformationUrl,omitempty"`

	// Uber
	UberLink         **UberLink         `json:"uberLink,omitempty"`
	UberTripBranding **UberTripBranding `json:"uberTripBranding,omitempty"`

	// Social Media
	FacebookCoverPhoto   **Image `json:"facebookCoverPhoto,omitempty"`
	FacebookPageUrl      *string `json:"facebookPageUrl,omitempty"`
	FacebookProfilePhoto **Image `json:"facebookProfilePhoto,omitempty"`

	GoogleCoverPhoto         **Image                    `json:"googleCoverPhoto,omitempty"`
	GooglePreferredPhoto     *string                    `json:"googlePreferredPhoto,omitempty"`
	GoogleProfilePhoto       **Image                    `json:"googleProfilePhoto,omitempty"`
	GoogleWebsiteOverride    **string                   `json:"googleWebsiteOverride,omitempty"`
	GooglePlaceId            *string                    `json:"googlePlaceId,omitempty"`
	GoogleMyBusinessLabels   *UnorderedStrings          `json:"googleMyBusinessLabels,omitempty"`
	GoogleEntityRelationship **GoogleEntityRelationship `json:"googleEntityRelationship,omitempty"`

	InstagramHandle *string `json:"instagramHandle,omitempty"`
	TwitterHandle   *string `json:"twitterHandle,omitempty"`

	PhotoGallery *[]Photo `json:"photoGallery,omitempty"`
	Videos       *[]Video `json:"videos,omitempty"`

	GoogleAttributes *map[string][]string `json:"googleAttributes,omitempty"`

	// Reviews
	ReviewGenerationUrl  *string `json:"reviewGenerationUrl,omitempty"`
	FirstPartyReviewPage *string `json:"firstPartyReviewPage,omitempty"`
}

func (l *HealthcareFacilityEntity) UnmarshalJSON(data []byte) error {
	type Alias HealthcareFacilityEntity
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

func (y HealthcareFacilityEntity) GetId() string {
	if y.BaseEntity.Meta != nil && y.BaseEntity.Meta.Id != nil {
		return *y.BaseEntity.Meta.Id
	}
	return ""
}

func (y HealthcareFacilityEntity) GetCategoryIds() (v []string) {
	if y.CategoryIds != nil {
		v = *y.CategoryIds
	}
	return v
}

func (y HealthcareFacilityEntity) GetName() string {
	if y.Name != nil {
		return GetString(y.Name)
	}
	return ""
}

func (y HealthcareFacilityEntity) GetAccountId() string {
	if y.BaseEntity.Meta != nil && y.BaseEntity.Meta.AccountId != nil {
		return *y.BaseEntity.Meta.AccountId
	}
	return ""
}

func (y HealthcareFacilityEntity) GetLine1() string {
	if y.Address != nil && y.Address.Line1 != nil {
		return GetString(y.Address.Line1)
	}
	return ""
}

func (y HealthcareFacilityEntity) GetLine2() string {
	if y.Address != nil && y.Address.Line2 != nil {
		return GetString(y.Address.Line2)
	}
	return ""
}

func (y HealthcareFacilityEntity) GetAddressHidden() bool {
	return GetNullableBool(y.AddressHidden)
}

func (y HealthcareFacilityEntity) GetExtraDescription() string {
	if y.Address != nil && y.Address.ExtraDescription != nil {
		return GetString(y.Address.ExtraDescription)
	}
	return ""
}

func (y HealthcareFacilityEntity) GetCity() string {
	if y.Address != nil && y.Address.City != nil {
		return GetString(y.Address.City)
	}
	return ""
}

func (y HealthcareFacilityEntity) GetRegion() string {
	if y.Address != nil && y.Address.Region != nil {
		return GetString(y.Address.Region)
	}
	return ""
}

func (y HealthcareFacilityEntity) GetCountryCode() string {
	if y.BaseEntity.Meta != nil && y.BaseEntity.Meta.CountryCode != nil {
		return GetString(y.BaseEntity.Meta.CountryCode)
	}
	return ""
}

func (y HealthcareFacilityEntity) GetPostalCode() string {
	if y.Address != nil && y.Address.PostalCode != nil {
		return GetString(y.Address.PostalCode)
	}
	return ""
}

func (y HealthcareFacilityEntity) GetMainPhone() string {
	if y.MainPhone != nil {
		return *y.MainPhone
	}
	return ""
}

func (y HealthcareFacilityEntity) GetLocalPhone() string {
	if y.LocalPhone != nil {
		return *y.LocalPhone
	}
	return ""
}

func (y HealthcareFacilityEntity) GetAlternatePhone() string {
	if y.AlternatePhone != nil {
		return *y.AlternatePhone
	}
	return ""
}

func (y HealthcareFacilityEntity) GetFax() string {
	if y.Fax != nil {
		return *y.Fax
	}
	return ""
}

func (y HealthcareFacilityEntity) GetMobilePhone() string {
	if y.MobilePhone != nil {
		return *y.MobilePhone
	}
	return ""
}

func (y HealthcareFacilityEntity) GetTollFreePhone() string {
	if y.TollFreePhone != nil {
		return *y.TollFreePhone
	}
	return ""
}

func (y HealthcareFacilityEntity) GetTtyPhone() string {
	if y.TtyPhone != nil {
		return *y.TtyPhone
	}
	return ""
}

func (y HealthcareFacilityEntity) GetFeaturedMessage() string {
	f := GetFeaturedMessage(y.FeaturedMessage)
	if f != nil {
		return GetString(f.Description)
	}
	return ""
}

func (y HealthcareFacilityEntity) GetFeaturedMessageUrl() string {
	f := GetFeaturedMessage(y.FeaturedMessage)
	if f != nil {
		return GetString(f.Url)
	}
	return ""
}

func (y HealthcareFacilityEntity) GetWebsiteUrl() string {
	w := GetWebsite(y.WebsiteUrl)
	if w != nil {
		return GetString(w.Url)
	}
	return ""
}

func (y HealthcareFacilityEntity) GetDisplayWebsiteUrl() string {
	w := GetWebsite(y.WebsiteUrl)
	if w != nil {
		return GetString(w.DisplayUrl)
	}
	return ""
}

func (y HealthcareFacilityEntity) GetReservationUrl() string {
	w := GetWebsite(y.ReservationUrl)
	if w != nil {
		return GetString(w.Url)
	}
	return ""
}

func (y HealthcareFacilityEntity) GetHours() *Hours {
	return GetHours(y.Hours)
}

func (y HealthcareFacilityEntity) GetAdditionalHoursText() string {
	if y.AdditionalHoursText != nil {
		return *y.AdditionalHoursText
	}
	return ""
}

func (y HealthcareFacilityEntity) GetDescription() string {
	if y.Description != nil {
		return *y.Description
	}
	return ""
}

func (y HealthcareFacilityEntity) GetTwitterHandle() string {
	if y.TwitterHandle != nil {
		return *y.TwitterHandle
	}
	return ""
}

func (y HealthcareFacilityEntity) GetFacebookPageUrl() string {
	if y.FacebookPageUrl != nil {
		return *y.FacebookPageUrl
	}
	return ""
}

func (y HealthcareFacilityEntity) GetYearEstablished() float64 {
	if y.YearEstablished != nil {
		return GetNullableFloat(y.YearEstablished)
	}
	return 0
}

func (y HealthcareFacilityEntity) GetDisplayLat() float64 {
	c := GetCoordinate(y.DisplayCoordinate)
	if c != nil {
		return GetNullableFloat(c.Latitude)
	}
	return 0
}

func (y HealthcareFacilityEntity) GetDisplayLng() float64 {
	c := GetCoordinate(y.DisplayCoordinate)
	if c != nil {
		return GetNullableFloat(c.Longitude)
	}
	return 0
}

func (y HealthcareFacilityEntity) GetRoutableLat() float64 {
	c := GetCoordinate(y.RoutableCoordinate)
	if c != nil {
		return GetNullableFloat(c.Latitude)
	}
	return 0
}

func (y HealthcareFacilityEntity) GetRoutableLng() float64 {
	c := GetCoordinate(y.RoutableCoordinate)
	if c != nil {
		return GetNullableFloat(c.Longitude)
	}
	return 0
}

func (y HealthcareFacilityEntity) GetYextDisplayLat() float64 {
	c := GetCoordinate(y.YextDisplayCoordinate)
	if c != nil {
		return GetNullableFloat(c.Latitude)
	}
	return 0
}

func (y HealthcareFacilityEntity) GetYextDisplayLng() float64 {
	c := GetCoordinate(y.YextDisplayCoordinate)
	if c != nil {
		return GetNullableFloat(c.Longitude)
	}
	return 0
}

func (y HealthcareFacilityEntity) GetYextRoutableLat() float64 {
	c := GetCoordinate(y.YextRoutableCoordinate)
	if c != nil {
		return GetNullableFloat(c.Latitude)
	}
	return 0
}

func (y HealthcareFacilityEntity) GetYextRoutableLng() float64 {
	c := GetCoordinate(y.YextRoutableCoordinate)
	if c != nil {
		return GetNullableFloat(c.Longitude)
	}
	return 0
}

func (y HealthcareFacilityEntity) GetBios() (v *Lists) {
	return GetLists(y.Bios)
}

func (y HealthcareFacilityEntity) GetCalendars() (v *Lists) {
	return GetLists(y.Calendars)
}

func (y HealthcareFacilityEntity) GetProductLists() (v *Lists) {
	return GetLists(y.ProductLists)
}

func (y HealthcareFacilityEntity) GetMenus() (v *Lists) {
	return GetLists(y.Menus)
}

func (y HealthcareFacilityEntity) GetReviewGenerationUrl() string {
	if y.ReviewGenerationUrl != nil {
		return *y.ReviewGenerationUrl
	}
	return ""
}

func (y HealthcareFacilityEntity) GetFirstPartyReviewPage() string {
	if y.FirstPartyReviewPage != nil {
		return *y.FirstPartyReviewPage
	}
	return ""
}

func (y HealthcareFacilityEntity) GetGooglePlaceId() string {
	if y.GooglePlaceId != nil {
		return *y.GooglePlaceId
	}
	return ""
}

func (y HealthcareFacilityEntity) String() string {
	b, _ := json.Marshal(y)
	return string(b)
}

func (y HealthcareFacilityEntity) GetKeywords() (v []string) {
	if y.Keywords != nil {
		v = *y.Keywords
	}
	return v
}

func (y HealthcareFacilityEntity) GetLanguage() (v string) {
	if y.BaseEntity.Meta.Language != nil {
		v = *y.BaseEntity.Meta.Language
	}
	return v
}

func (y HealthcareFacilityEntity) GetAssociations() (v []string) {
	if y.Associations != nil {
		v = *y.Associations
	}
	return v
}

func (y HealthcareFacilityEntity) GetEmails() (v []string) {
	if y.Emails != nil {
		v = *y.Emails
	}
	return v
}

func (y HealthcareFacilityEntity) GetSpecialties() (v []string) {
	if y.Specialties != nil {
		v = *y.Specialties
	}
	return v
}

func (y HealthcareFacilityEntity) GetServices() (v []string) {
	if y.Services != nil {
		v = *y.Services
	}
	return v
}

func (y HealthcareFacilityEntity) GetBrands() (v []string) {
	if y.Brands != nil {
		v = *y.Brands
	}
	return v
}

func (y HealthcareFacilityEntity) GetLanguages() (v []string) {
	if y.Languages != nil {
		v = *y.Languages
	}
	return v
}

func (y HealthcareFacilityEntity) GetPaymentOptions() (v []string) {
	if y.PaymentOptions != nil {
		v = *y.PaymentOptions
	}
	return v
}

func (y HealthcareFacilityEntity) GetAcceptingNewPatients() bool {
	return GetNullableBool(y.AcceptingNewPatients)
}

func (y HealthcareFacilityEntity) GetVideos() (v []Video) {
	if y.Videos != nil {
		v = *y.Videos
	}
	return v
}

func (y HealthcareFacilityEntity) GetGoogleAttributes() map[string][]string {
	if y.GoogleAttributes != nil {
		return *y.GoogleAttributes
	}
	return nil
}

func (y HealthcareFacilityEntity) GetHolidayHours() []HolidayHours {
	h := GetHours(y.Hours)
	if h != nil && h.HolidayHours != nil {
		return *h.HolidayHours
	}
	return nil
}

func (y HealthcareFacilityEntity) IsClosed() bool {
	return GetNullableBool(y.Closed)
}

func (y HealthcareFacilityEntity) GetNPI() string {
	if y.NPI != nil {
		return GetString(y.NPI)
	}
	return ""
}
