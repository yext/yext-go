package yext

import (
	"encoding/json"
)

const ENTITYTYPE_EVENT EntityType = "event"

type EventEntity struct {
	BaseEntity

	// Admin
	CategoryIds *UnorderedStrings `json:"categoryIds,omitempty"`

	// Address Fields
	Name          *string  `json:"name,omitempty"`
	Address       *Address `json:"address,omitempty"`
	AddressHidden *bool    `json:"addressHidden,omitempty"`
	ISORegionCode *string  `json:"isoRegionCode,omitempty"`

	// Other Contact Info
	MainPhone      *string `json:"mainPhone,omitempty"`
	AlternatePhone *string `json:"alternatePhone,omitempty"`
	FaxPhone       *string `json:"fax,omitempty"`
	TollFreePhone  *string `json:"tollFreePhone,omitempty"`

	//Event Info
	Description        *string           `json:"description,omitempty"`
	TicketUrl          *string           `json:"ticketUrl,omitempty"`
	Hours              *Hours            `json:"hours,omitempty"`
	Brands             *[]string         `json:"brands,omitempty"`
	Logo               *LocationPhoto    `json:"logo,omitempty"`
	PaymentOptions     *[]string         `json:"paymentOptions,omitempty"`
	Timezone           *string           `json:"timezone,omitempty"`
	YearEstablished    *float64          `json:"yearEstablished,omitempty"`
	AgeRange           *AgeRange         `json:"ageRange,omitempty"`
	Time               *TimeRange        `json:"time,omitempty"`
	IsFreeEvent        *bool             `json:"isFreeEvent,omitempty"`
	IsTicketedEvent    *bool             `json:"isTicketedEvent,omitempty"`
	EventStatus        *string           `json:"eventStatus,omitempty"`
	VenueName          *string           `json:"venueName,omitempty"`
	TicketAvailability *string           `json:"ticketAvailability,omitempty"`
	TicketPriceRange   *TicketPriceRange `json:"ticketPriceRange,omitempty"`

	//Lats & Lngs
	DisplayCoordinate  *Coordinate `json:"yextDisplayCoordinate,omitempty"`
	RoutableCoordinate *Coordinate `json:"yextRoutableCoordinate,omitempty"`
	DropoffCoordinate  *Coordinate `json:"yextDropoffCoordinate,omitempty"`
	WalkableCoordinate *Coordinate `json:"yextWalkableCoordinate,omitempty"`
	PickupCoordinate   *Coordinate `json:"yextPickupCoordinate,omitempty"`

	//Event Organizer Info
	OrganizerEmail *string `json:"organizerEmail,omitempty"`
	OrganizerName  *string `json:"organizerName,omitempty"`
	OrganizerPhone *string `json:"organizerPhone,omitempty"`

	//Urls
	WebsiteUrl      *Website         `json:"websiteUrl,omitempty"`
	FeaturedMessage *FeaturedMessage `json:"featuredMessage,omitempty"`

	//Social Media
	FacebookCoverPhoto   *LocationPhoto `json:"facebookCoverPhoto,omitempty"`
	FacebookPageUrl      *string        `json:"facebookPageUrl,omitempty"`
	FacebookProfilePhoto *LocationPhoto `json:"facebookProfilePhoto,omitempty"`
	FacebookStoreId      *string        `json:"facebookStoreId,omitempty"`
	FacebookVanityUrl    *string        `json:"facebookVanityUrl,omitempty"`

	TwitterHandle *string `json:"twitterHandle,omitempty"`

	PhotoGallery *[]PhotoGalleryItem `json:"photoGallery,omitempty"`
	Videos       *[]Video            `json:"videos,omitempty"`

	GoogleAttributes *map[string][]string `json:"googleAttributes,omitempty"`
}

type AgeRange struct {
	MinValue *int64 `json:"minValue,omitempty"`
	MaxValue *int64 `json:"maxValue,omitempty"`
}

type TimeRange struct {
	Start *string `json:"start,omitempty"`
	End   *string `json:"end,omitempty"`
}

type TicketPriceRange struct {
	MinValue     *string `json:"minValue,omitempty"`
	MaxValue     *string `json:"maxValue,omitempty"`
	CurrencyCode *string `json:"currencyCode,omitempty"`
}

type PhotoGalleryItem struct {
	Image           *Image  `json:"image,omitempty"`
	ClickthroughUrl *string `json:"clickthroughUrl,omitempty"`
	Description     *string `json:"description,omitempty"`
	Details         *string `json:"details,omitempty"`
}

func (y EventEntity) GetId() string {
	if y.BaseEntity.Meta != nil && y.BaseEntity.Meta.Id != nil {
		return *y.BaseEntity.Meta.Id
	}
	return ""
}

func (y EventEntity) GetName() string {
	if y.Name != nil {
		return *y.Name
	}
	return ""
}

func (y EventEntity) GetCountryCode() string {
	if y.Address != nil && y.Address.CountryCode != nil {
		return *y.Address.CountryCode
	}
	return ""
}

func (y EventEntity) GetLine1() string {
	if y.Address != nil && y.Address.Line1 != nil {
		return *y.Address.Line1
	}
	return ""
}

func (y EventEntity) GetLine2() string {
	if y.Address != nil && y.Address.Line2 != nil {
		return *y.Address.Line2
	}
	return ""
}

func (y EventEntity) GetCity() string {
	if y.Address != nil && y.Address.City != nil {
		return *y.Address.City
	}
	return ""
}

func (y EventEntity) GetRegion() string {
	if y.Address != nil && y.Address.Region != nil {
		return *y.Address.Region
	}
	return ""
}

func (y EventEntity) GetPostalCode() string {
	if y.Address != nil && y.Address.PostalCode != nil {
		return *y.Address.PostalCode
	}
	return ""
}

func (y EventEntity) GetAddressHidden() bool {
	if y.AddressHidden != nil {
		return *y.AddressHidden
	}
	return false
}

func (y EventEntity) GetMainPhone() string {
	if y.MainPhone != nil {
		return *y.MainPhone
	}
	return ""
}

func (y EventEntity) GetAlternatePhone() string {
	if y.AlternatePhone != nil {
		return *y.AlternatePhone
	}
	return ""
}

func (y EventEntity) GetFaxPhone() string {
	if y.FaxPhone != nil {
		return *y.FaxPhone
	}
	return ""
}

func (y EventEntity) GetTollFreePhone() string {
	if y.TollFreePhone != nil {
		return *y.TollFreePhone
	}
	return ""
}

func (y EventEntity) GetOrganizerEmail() string {
	if y.OrganizerEmail != nil {
		return *y.OrganizerEmail
	}
	return ""
}

func (y EventEntity) GetOrganizerName() string {
	if y.OrganizerName != nil {
		return *y.OrganizerName
	}
	return ""
}

func (y EventEntity) GetOrganizerPhone() string {
	if y.OrganizerPhone != nil {
		return *y.OrganizerPhone
	}
	return ""
}

func (y EventEntity) GetDescription() string {
	if y.Description != nil {
		return *y.Description
	}
	return ""
}

func (y EventEntity) GetHolidayHours() []HolidayHours {
	if y.Hours != nil && y.Hours.HolidayHours != nil {
		return *y.Hours.HolidayHours
	}
	return nil
}

func (y EventEntity) GetBrands() (v []string) {
	if y.Brands != nil {
		v = *y.Brands
	}
	return v
}

func (y EventEntity) GetDisplayLat() float64 {
	if y.DisplayCoordinate != nil && y.DisplayCoordinate.Latitude != nil {
		return *y.DisplayCoordinate.Latitude
	}
	return 0
}

func (y EventEntity) GetDisplayLng() float64 {
	if y.DisplayCoordinate != nil && y.DisplayCoordinate.Longitude != nil {
		return *y.DisplayCoordinate.Longitude
	}
	return 0
}

func (y EventEntity) GetRoutableLat() float64 {
	if y.RoutableCoordinate != nil && y.RoutableCoordinate.Latitude != nil {
		return *y.RoutableCoordinate.Latitude
	}
	return 0
}

func (y EventEntity) GetRoutableLng() float64 {
	if y.RoutableCoordinate != nil && y.RoutableCoordinate.Longitude != nil {
		return *y.RoutableCoordinate.Longitude
	}
	return 0
}

func (y EventEntity) GetDropoffLat() float64 {
	if y.DropoffCoordinate != nil && y.DropoffCoordinate.Latitude != nil {
		return *y.DropoffCoordinate.Latitude
	}
	return 0
}

func (y EventEntity) GetDropoffLng() float64 {
	if y.DropoffCoordinate != nil && y.DropoffCoordinate.Longitude != nil {
		return *y.DropoffCoordinate.Longitude
	}
	return 0
}

func (y EventEntity) GetPickupLat() float64 {
	if y.PickupCoordinate != nil && y.PickupCoordinate.Latitude != nil {
		return *y.PickupCoordinate.Latitude
	}
	return 0
}

func (y EventEntity) GetPickupLng() float64 {
	if y.PickupCoordinate != nil && y.PickupCoordinate.Longitude != nil {
		return *y.PickupCoordinate.Longitude
	}
	return 0
}

func (y EventEntity) GetWalkableLat() float64 {
	if y.WalkableCoordinate != nil && y.WalkableCoordinate.Latitude != nil {
		return *y.WalkableCoordinate.Latitude
	}
	return 0
}

func (y EventEntity) GetWalkableLng() float64 {
	if y.WalkableCoordinate != nil && y.WalkableCoordinate.Longitude != nil {
		return *y.WalkableCoordinate.Longitude
	}
	return 0
}

func (y EventEntity) GetWebsiteUrl() string {
	if y.WebsiteUrl != nil && y.WebsiteUrl.Url != nil {
		return *y.WebsiteUrl.Url
	}
	return ""
}

func (y EventEntity) GetDisplayWebsiteUrl() string {
	if y.WebsiteUrl != nil && y.WebsiteUrl.DisplayUrl != nil {
		return *y.WebsiteUrl.DisplayUrl
	}
	return ""
}

func (y EventEntity) GetTicketUrl() string {
	if y.TicketUrl != nil {
		return *y.TicketUrl
	}
	return ""
}

func (y EventEntity) GetFeaturedMessageDescription() string {
	if y.FeaturedMessage != nil && y.FeaturedMessage.Description != nil {
		return *y.FeaturedMessage.Description
	}
	return ""
}

func (y EventEntity) GetFeaturedMessageUrl() string {
	if y.FeaturedMessage != nil && y.FeaturedMessage.Url != nil {
		return *y.FeaturedMessage.Url
	}
	return ""
}

func (y EventEntity) GetFacebookPageUrl() string {
	if y.FacebookPageUrl != nil {
		return *y.FacebookPageUrl
	}
	return ""
}

func (y EventEntity) GetGoogleAttributes() map[string][]string {
	if y.GoogleAttributes != nil {
		return *y.GoogleAttributes
	}
	return nil
}

func (y EventEntity) GetPaymentOptions() (v []string) {
	if y.PaymentOptions != nil {
		v = *y.PaymentOptions
	}
	return v
}

func (y EventEntity) GetTimezone() string {
	if y.Timezone != nil {
		return *y.Timezone
	}
	return ""
}

func (y EventEntity) GetYearEstablished() float64 {
	if y.YearEstablished != nil {
		return *y.YearEstablished
	}
	return 0
}

func (y EventEntity) GetIsTicketedEvent() bool {
	if y.IsTicketedEvent != nil {
		return *y.IsTicketedEvent
	}
	return false
}

func (y EventEntity) GetIsFreeEvent() bool {
	if y.IsFreeEvent != nil {
		return *y.IsFreeEvent
	}
	return false
}

func (y EventEntity) GetEventStatus() string {
	if y.EventStatus != nil {
		return *y.EventStatus
	}
	return ""
}

func (y EventEntity) GetVenueName() string {
	if y.VenueName != nil {
		return *y.VenueName
	}
	return ""
}

func (y EventEntity) GetTicketAvailability() string {
	if y.TicketAvailability != nil {
		return *y.TicketAvailability
	}
	return ""
}

func (b *EventEntity) GetCategoryIds() (v UnorderedStrings) {
	if b.CategoryIds != nil {
		v = *b.CategoryIds
	}
	return v
}

func (e *EventEntity)SingleString() string {
	b, _ := json.Marshal(e)
	return string(b)
}
