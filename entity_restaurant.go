package yext

import (
	"encoding/json"
)

const ENTITYTYPE_RESTAURANT EntityType = "restaurant"

type RestaurantEntity struct {
	BaseEntity

	// Admin
	CategoryIds *[]string `json:"categoryIds,omitempty"`
	Closed      **bool    `json:"closed,omitempty"`
	Keywords    *[]string `json:"keywords,omitempty"`

	// Address Fields
	Name          *string  `json:"name,omitempty"`
	Address       *Address `json:"address,omitempty"`
	AddressHidden **bool   `json:"addressHidden,omitempty"`
	ISORegionCode *string  `json:"isoRegionCode,omitempty"`

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
	Description               *string           `json:"description,omitempty"`
	Hours                     **Hours           `json:"hours,omitempty"`
	BrunchHours               **Hours           `json:"brunchHours,omitempty"`
	DriveThroughHours         **Hours           `json:"driveThroughHours,omitempty"`
	AdditionalHoursText       *string           `json:"additionalHoursText,omitempty"`
	DeliveryHours             **Hours           `json:"deliveryHours,omitempty"`
	YearEstablished           **float64         `json:"yearEstablished,omitempty"`
	Services                  *[]string         `json:"services,omitempty"`
	Languages                 *[]string         `json:"languages,omitempty"`
	Logo                      **Photo           `json:"logo,omitempty"`
	PaymentOptions            *[]string         `json:"paymentOptions,omitempty"`
	Geomodifier               *string           `json:"geomodifier,omitempty"`
	PickupAndDeliveryServices *[]string         `json:"pickupAndDeliveryServices,omitempty"`
	MealsServed               *UnorderedStrings `json:"mealsServed,omitempty"`
	AcceptsReservations       **Ternary         `json:"kitchenHours,omitempty"`
	KitchenHours              **Hours           `json:"kitchenHours,omitempty"`
	DineInHours               **Hours           `json:"dineInHours,omitempty"`

	// Lats & Lngs
	DisplayCoordinate  **Coordinate `json:"displayCoordinate,omitempty"`
	DropoffCoordinate  **Coordinate `json:"dropoffCoordinate,omitempty"`
	WalkableCoordinate **Coordinate `json:"walkableCoordinate,omitempty"`
	RoutableCoordinate **Coordinate `json:"routableCoordinate,omitempty"`
	PickupCoordinate   **Coordinate `json:"pickupCoordinate,omitempty"`

	// Lists
	Bios         **Lists `json:"bios,omitempty"`
	Calendars    **Lists `json:"calendars,omitempty"`
	Menus        **Lists `json:"menus,omitempty"`
	ProductLists **Lists `json:"productLists,omitempty"`

	// Urls
	MenuUrl         **Website         `json:"menuUrl,omitempty"`
	OrderUrl        **Website         `json:"orderUrl,omitempty"`
	ReservationUrl  **Website         `json:"reservationUrl,omitempty"`
	WebsiteUrl      **Website         `json:"websiteUrl,omitempty"`
	FeaturedMessage **FeaturedMessage `json:"featuredMessage,omitempty"`

	// Uber
	UberLink         **UberLink         `json:"uberLink,omitempty"`
	UberTripBranding **UberTripBranding `json:"uberTripBranding,omitempty"`

	// Social Media
	FacebookCoverPhoto         **Image `json:"facebookCoverPhoto,omitempty"`
	FacebookLocationDescriptor *string `json:"facebookDescriptor,omitempty"`
	FacebookPageUrl            *string `json:"facebookPageUrl,omitempty"`
	FacebookProfilePhoto       **Image `json:"facebookProfilePhoto,omitempty"`
	FacebookUsername           *string `json:"facebookVanityUrl,omitempty"`

	GoogleCoverPhoto      **Image  `json:"googleCoverPhoto,omitempty"`
	GooglePreferredPhoto  *string  `json:"googlePreferredPhoto,omitempty"`
	GoogleProfilePhoto    **Image  `json:"googleProfilePhoto,omitempty"`
	GoogleWebsiteOverride **string `json:"googleWebsiteOverride,omitempty"`

	InstagramHandle *string `json:"instagramHandle,omitempty"`
	TwitterHandle   *string `json:"twitterHandle,omitempty"`

	PhotoGallery *[]Photo `json:"photoGallery,omitempty"`
	Videos       *[]Video `json:"videos,omitempty"`

	GoogleAttributes *map[string][]string `json:"googleAttributes,omitempty"`

	// Reviews
	ReviewGenerationUrl  *string `json:"reviewGenerationUrl,omitempty"`
	FirstPartyReviewPage *string `json:"firstPartyReviewPage,omitempty"`

	TimeZoneUtcOffset *string `json:"timeZoneUtcOffset,omitempty"`
	Timezone          *string `json:"timezone,omitempty"`
}

func (r *RestaurantEntity) UnmarshalJSON(data []byte) error {
	type Alias RestaurantEntity
	a := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	return UnmarshalEntityJSON(r, data)
}

func (r RestaurantEntity) GetId() string {
	if r.BaseEntity.Meta != nil && r.BaseEntity.Meta.Id != nil {
		return *r.BaseEntity.Meta.Id
	}
	return ""
}

func (r RestaurantEntity) GetName() string {
	if r.Name != nil {
		return GetString(r.Name)
	}
	return ""
}

func (r RestaurantEntity) GetAccountId() string {
	if r.BaseEntity.Meta != nil && r.BaseEntity.Meta.AccountId != nil {
		return *r.BaseEntity.Meta.AccountId
	}
	return ""
}

func (r RestaurantEntity) GetLine1() string {
	if r.Address != nil && r.Address.Line1 != nil {
		return GetString(r.Address.Line1)
	}
	return ""
}

func (r RestaurantEntity) GetLine2() string {
	if r.Address != nil && r.Address.Line2 != nil {
		return GetString(r.Address.Line2)
	}
	return ""
}

func (r RestaurantEntity) GetAddressHidden() bool {
	return GetNullableBool(r.AddressHidden)
}

func (r RestaurantEntity) GetExtraDescription() string {
	if r.Address != nil && r.Address.ExtraDescription != nil {
		return GetString(r.Address.ExtraDescription)
	}
	return ""
}

func (r RestaurantEntity) GetCity() string {
	if r.Address != nil && r.Address.City != nil {
		return GetString(r.Address.City)
	}
	return ""
}

func (r RestaurantEntity) GetRegion() string {
	if r.Address != nil && r.Address.Region != nil {
		return GetString(r.Address.Region)
	}
	return ""
}

func (r RestaurantEntity) GetPostalCode() string {
	if r.Address != nil && r.Address.PostalCode != nil {
		return GetString(r.Address.PostalCode)
	}
	return ""
}

func (r RestaurantEntity) GetMainPhone() string {
	if r.MainPhone != nil {
		return *r.MainPhone
	}
	return ""
}

func (r RestaurantEntity) GetLocalPhone() string {
	if r.LocalPhone != nil {
		return *r.LocalPhone
	}
	return ""
}

func (r RestaurantEntity) GetAlternatePhone() string {
	if r.AlternatePhone != nil {
		return *r.AlternatePhone
	}
	return ""
}

func (r RestaurantEntity) GetFax() string {
	if r.Fax != nil {
		return *r.Fax
	}
	return ""
}

func (r RestaurantEntity) GetMobilePhone() string {
	if r.MobilePhone != nil {
		return *r.MobilePhone
	}
	return ""
}

func (r RestaurantEntity) GetTollFreePhone() string {
	if r.TollFreePhone != nil {
		return *r.TollFreePhone
	}
	return ""
}

func (r RestaurantEntity) GetTtyPhone() string {
	if r.TtyPhone != nil {
		return *r.TtyPhone
	}
	return ""
}

func (r RestaurantEntity) GetFeaturedMessage() string {
	f := GetFeaturedMessage(r.FeaturedMessage)
	if f != nil {
		return GetString(f.Description)
	}
	return ""
}

func (r RestaurantEntity) GetFeaturedMessageUrl() string {
	f := GetFeaturedMessage(r.FeaturedMessage)
	if f != nil {
		return GetString(f.Url)
	}
	return ""
}

func (r RestaurantEntity) GetWebsiteUrl() string {
	w := GetWebsite(r.WebsiteUrl)
	if w != nil {
		return GetString(w.Url)
	}
	return ""
}

func (r RestaurantEntity) GetDisplayWebsiteUrl() string {
	w := GetWebsite(r.WebsiteUrl)
	if w != nil {
		return GetString(w.DisplayUrl)
	}
	return ""
}

func (r RestaurantEntity) GetReservationUrl() string {
	w := GetWebsite(r.ReservationUrl)
	if w != nil {
		return GetString(w.Url)
	}
	return ""
}

func (r RestaurantEntity) GetHours() *Hours {
	return GetHours(r.Hours)
}

func (r RestaurantEntity) GetAdditionalHoursText() string {
	if r.AdditionalHoursText != nil {
		return *r.AdditionalHoursText
	}
	return ""
}

func (r RestaurantEntity) GetDescription() string {
	if r.Description != nil {
		return *r.Description
	}
	return ""
}

func (r RestaurantEntity) GetTwitterHandle() string {
	if r.TwitterHandle != nil {
		return *r.TwitterHandle
	}
	return ""
}

func (r RestaurantEntity) GetFacebookPageUrl() string {
	if r.FacebookPageUrl != nil {
		return *r.FacebookPageUrl
	}
	return ""
}

func (r RestaurantEntity) GetYearEstablished() float64 {
	if r.YearEstablished != nil {
		return GetNullableFloat(r.YearEstablished)
	}
	return 0
}

func (r RestaurantEntity) GetDisplayLat() float64 {
	c := GetCoordinate(r.DisplayCoordinate)
	if c != nil {
		return GetNullableFloat(c.Latitude)
	}
	return 0
}

func (r RestaurantEntity) GetDisplayLng() float64 {
	c := GetCoordinate(r.DisplayCoordinate)
	if c != nil {
		return GetNullableFloat(c.Longitude)
	}
	return 0
}

func (r RestaurantEntity) GetRoutableLat() float64 {
	c := GetCoordinate(r.RoutableCoordinate)
	if c != nil {
		return GetNullableFloat(c.Latitude)
	}
	return 0
}

func (r RestaurantEntity) GetRoutableLng() float64 {
	c := GetCoordinate(r.RoutableCoordinate)
	if c != nil {
		return GetNullableFloat(c.Longitude)
	}
	return 0
}

func (r RestaurantEntity) GetBios() (v *Lists) {
	return GetLists(r.Bios)
}

func (r RestaurantEntity) GetCalendars() (v *Lists) {
	return GetLists(r.Calendars)
}

func (r RestaurantEntity) GetProductLists() (v *Lists) {
	return GetLists(r.ProductLists)
}

func (r RestaurantEntity) GetMenus() (v *Lists) {
	return GetLists(r.Menus)
}

func (r RestaurantEntity) GetReviewGenerationUrl() string {
	if r.ReviewGenerationUrl != nil {
		return *r.ReviewGenerationUrl
	}
	return ""
}

func (r RestaurantEntity) GetFirstPartyReviewPage() string {
	if r.FirstPartyReviewPage != nil {
		return *r.FirstPartyReviewPage
	}
	return ""
}

func (r RestaurantEntity) String() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r RestaurantEntity) GetKeywords() (v []string) {
	if r.Keywords != nil {
		v = *r.Keywords
	}
	return v
}

func (r RestaurantEntity) GetLanguage() (v string) {
	if r.BaseEntity.Meta.Language != nil {
		v = *r.BaseEntity.Meta.Language
	}
	return v
}

func (r RestaurantEntity) GetEmails() (v []string) {
	if r.Emails != nil {
		v = *r.Emails
	}
	return v
}

func (r RestaurantEntity) GetServices() (v []string) {
	if r.Services != nil {
		v = *r.Services
	}
	return v
}

func (r RestaurantEntity) GetLanguages() (v []string) {
	if r.Languages != nil {
		v = *r.Languages
	}
	return v
}

func (r RestaurantEntity) GetPaymentOptions() (v []string) {
	if r.PaymentOptions != nil {
		v = *r.PaymentOptions
	}
	return v
}

func (r RestaurantEntity) GetGeomodifier() string {
	if r.Geomodifier != nil {
		return GetString(r.Geomodifier)
	}
	return ""
}

func (r RestaurantEntity) GetVideos() (v []Video) {
	if r.Videos != nil {
		v = *r.Videos
	}
	return v
}

func (r RestaurantEntity) GetGoogleAttributes() map[string][]string {
	if r.GoogleAttributes != nil {
		return *r.GoogleAttributes
	}
	return nil
}

func (r RestaurantEntity) GetHolidayHours() []HolidayHours {
	h := GetHours(r.Hours)
	if h != nil && h.HolidayHours != nil {
		return *h.HolidayHours
	}
	return nil
}

func (r RestaurantEntity) IsClosed() bool {
	return GetNullableBool(r.Closed)
}
