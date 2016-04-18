package yext

// TODO
// * Need better custom field accessors and helpers
// * The API will accept some things and return them in a different format - this makes diff'ing difficult:
// ** Phone: Send in 540-444-4444, get back 5404444444
// ** Custom Field Multi-Option: Send in options ["3", "2", "1"], get back ["1", "2", "3"]

import (
	"encoding/json"
	"fmt"
)

// Location is the representation of a Location in Yext Location Manager.
// For details see https://www.yext.com/support/platform-api/#Administration_API/Locations.htm
type Location struct {
	Id                     *string                `json:"id,omitempty"`
	Name                   *string                `json:"locationName,omitempty"`
	Lists                  []ECL                  `json:"lists,omitempty"`
	Keywords               []string               `json:"keywords,omitempty"`
	Associations           []string               `json:"associations,omitempty"`
	CustomFields           map[string]interface{} `json:"customFields,omitempty"`
	CustomerId             *string                `json:"customerId,omitempty"`
	Address                *string                `json:"address,omitempty"`
	Address2               *string                `json:"address2,omitempty"`
	SuppressAddress        *bool                  `json:"suppressAddress,omitempty"`
	DisplayAddress         *string                `json:"displayAddress,omitempty"`
	City                   *string                `json:"city,omitempty"`
	State                  *string                `json:"state,omitempty"`
	Zip                    *string                `json:"zip,omitempty"`
	CountryCode            *string                `json:"countryCode,omitempty"`
	Phone                  *string                `json:"phone,omitempty"`
	IsPhoneTracked         *bool                  `json:"isPhoneTracked,omitempty"`
	LocalPhone             *string                `json:"localPhone,omitempty"`
	AlternatePhone         *string                `json:"alternatePhone,omitempty"`
	FaxPhone               *string                `json:"faxPhone,omitempty"`
	MobilePhone            *string                `json:"mobilePhone,omitempty"`
	TollFreePhone          *string                `json:"tollFreePhone,omitempty"`
	TtyPhone               *string                `json:"ttyPhone,omitempty"`
	SpecialOffer           *string                `json:"specialOffer,omitempty"`
	SpecialOfferUrl        *string                `json:"specialOfferUrl,omitempty"`
	WebsiteUrl             *string                `json:"websiteUrl,omitempty"`
	DisplayWebsiteUrl      *string                `json:"displayWebsiteUrl,omitempty"`
	ReservationUrl         *string                `json:"reservationUrl,omitempty"`
	Hours                  *string                `json:"hours,omitempty"` // A concrete type would be nice for this
	AdditionalHoursText    *string                `json:"additionalHoursText,omitempty"`
	Description            *string                `json:"description,omitempty"`
	PaymentOptions         []string               `json:"paymentOptions,omitempty"`
	VideoUrls              []string               `json:"videoUrls,omitempty"`
	TwitterHandle          *string                `json:"twitterHandle,omitempty"`
	FacebookPageUrl        *string                `json:"facebookPageUrl,omitempty"`
	YearEstablished        *string                `json:"yearEstablished,omitempty"`
	DisplayLat             *float64               `json:"displayLat,omitempty"`
	DisplayLng             *float64               `json:"displayLng,omitempty"`
	RoutableLat            *float64               `json:"routableLat,omitempty"`
	RoutableLng            *float64               `json:"routableLng,omitempty"`
	Emails                 []string               `json:"emails,omitempty"`
	Specialties            []string               `json:"specialties,omitempty"`
	Services               []string               `json:"services,omitempty"`
	Brands                 []string               `json:"brands,omitempty"`
	Languages              []string               `json:"languages,omitempty"`
	FolderId               *string                `json:"folderId,omitempty"`
	FacebookCoverPhoto     *LocationPhoto         `json:"facebookCoverPhoto,omitempty"`
	FacebookProfilePicture *LocationPhoto         `json:"facebookProfilePicture,omitempty"`
	Logo                   *LocationPhoto         `json:"logo,omitempty"`
	Photos                 []LocationPhoto        `json:"photos,omitempty"`
	Closed                 *LocationClosed        `json:"closed,omitempty"`
	CategoryIds            []string               `json:"categoryIds,omitempty"`
	HolidayHours           []HolidayHours         `json:"holidayHours,omitempty"`
	hydrated               bool
}

func (y Location) GetId() string {
	if y.Id != nil {
		return *y.Id
	}
	return ""
}

func (y Location) GetName() string {
	if y.Name != nil {
		return *y.Name
	}
	return ""
}

func (y Location) GetCustomerId() string {
	if y.CustomerId != nil {
		return *y.CustomerId
	}
	return ""
}

func (y Location) GetAddress() string {
	if y.Address != nil {
		return *y.Address
	}
	return ""
}

func (y Location) GetAddress2() string {
	if y.Address2 != nil {
		return *y.Address2
	}
	return ""
}

func (y Location) GetSuppressAddress() bool {
	if y.SuppressAddress != nil {
		return *y.SuppressAddress
	}
	return false
}

func (y Location) GetDisplayAddress() string {
	if y.DisplayAddress != nil {
		return *y.DisplayAddress
	}
	return ""
}

func (y Location) GetCity() string {
	if y.City != nil {
		return *y.City
	}
	return ""
}

func (y Location) GetState() string {
	if y.State != nil {
		return *y.State
	}
	return ""
}

func (y Location) GetZip() string {
	if y.Zip != nil {
		return *y.Zip
	}
	return ""
}

func (y Location) GetCountryCode() string {
	if y.CountryCode != nil {
		return *y.CountryCode
	}
	return ""
}

func (y Location) GetPhone() string {
	if y.Phone != nil {
		return *y.Phone
	}
	return ""
}

func (y Location) GetIsPhoneTracked() bool {
	if y.IsPhoneTracked != nil {
		return *y.IsPhoneTracked
	}
	return false
}

func (y Location) GetLocalPhone() string {
	if y.LocalPhone != nil {
		return *y.LocalPhone
	}
	return ""
}

func (y Location) GetAlternatePhone() string {
	if y.AlternatePhone != nil {
		return *y.AlternatePhone
	}
	return ""
}

func (y Location) GetFaxPhone() string {
	if y.FaxPhone != nil {
		return *y.FaxPhone
	}
	return ""
}

func (y Location) GetMobilePhone() string {
	if y.MobilePhone != nil {
		return *y.MobilePhone
	}
	return ""
}

func (y Location) GetTollFreePhone() string {
	if y.TollFreePhone != nil {
		return *y.TollFreePhone
	}
	return ""
}

func (y Location) GetTtyPhone() string {
	if y.TtyPhone != nil {
		return *y.TtyPhone
	}
	return ""
}

func (y Location) GetSpecialOffer() string {
	if y.SpecialOffer != nil {
		return *y.SpecialOffer
	}
	return ""
}

func (y Location) GetSpecialOfferUrl() string {
	if y.SpecialOfferUrl != nil {
		return *y.SpecialOfferUrl
	}
	return ""
}

func (y Location) GetWebsiteUrl() string {
	if y.WebsiteUrl != nil {
		return *y.WebsiteUrl
	}
	return ""
}

func (y Location) GetDisplayWebsiteUrl() string {
	if y.DisplayWebsiteUrl != nil {
		return *y.DisplayWebsiteUrl
	}
	return ""
}

func (y Location) GetReservationUrl() string {
	if y.ReservationUrl != nil {
		return *y.ReservationUrl
	}
	return ""
}

func (y Location) GetHours() string {
	if y.Hours != nil {
		return *y.Hours
	}
	return ""
}

func (y Location) GetAdditionalHoursText() string {
	if y.AdditionalHoursText != nil {
		return *y.AdditionalHoursText
	}
	return ""
}

func (y Location) GetDescription() string {
	if y.Description != nil {
		return *y.Description
	}
	return ""
}

func (y Location) GetTwitterHandle() string {
	if y.TwitterHandle != nil {
		return *y.TwitterHandle
	}
	return ""
}

func (y Location) GetFacebookPageUrl() string {
	if y.FacebookPageUrl != nil {
		return *y.FacebookPageUrl
	}
	return ""
}

func (y Location) GetYearEstablished() string {
	if y.YearEstablished != nil {
		return *y.YearEstablished
	}
	return ""
}

func (y Location) GetDisplayLat() float64 {
	if y.DisplayLat != nil {
		return *y.DisplayLat
	}
	return 0
}

func (y Location) GetDisplayLng() float64 {
	if y.DisplayLng != nil {
		return *y.DisplayLng
	}
	return 0
}

func (y Location) GetRoutableLat() float64 {
	if y.RoutableLat != nil {
		return *y.RoutableLat
	}
	return 0
}

func (y Location) GetRoutableLng() float64 {
	if y.RoutableLng != nil {
		return *y.RoutableLng
	}
	return 0
}

func (y Location) GetFolderId() string {
	if y.FolderId != nil {
		return *y.FolderId
	}
	return ""
}

func (y Location) String() string {
	b, _ := json.Marshal(y)
	return string(b)
}

// Photo represents a photo associated with a Location in Yext Location Manager.
// For details see https://www.yext.com/support/platform-api/#Administration_API/Locations.htm#Photo
type LocationPhoto struct {
	Url             string `json:"url,omitempty"`
	Description     string `json:"description,omitempty"`
	ClickThroughURL string `json:"clickthroughUrl,omitempty"`
}

func (l Photo) String() string {
	return fmt.Sprintf("Url: '%v', Description: '%v'", l.Url, l.Description)
}

// LocationClosed represents the 'closed' state of a Location in Yext Location Manager.
// For details see https://www.yext.com/support/platform-api/#Administration_API/Locations.htm#Closed
type LocationClosed struct {
	IsClosed   string `json:"isClosed"`
	ClosedDate string `json:"closedDate,omitempty"`
}

func (l LocationClosed) String() string {
	return fmt.Sprintf("isClosed: %v, closedDate: '%v'", l.IsClosed, l.ClosedDate)
}

// HolidayHours represents individual exceptions to a Location's regular hours in Yext Location Manager.
// For details see
type HolidayHours struct {
	Date  string `json:"date"`
	Hours string `json:"hours"`
}
