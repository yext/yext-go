package yext

// TODO
// * Need better custom field accessors and helpers
// * The API will accept some things and return them in a different format - this makes diff'ing difficult:
// ** Phone: Send in 540-444-4444, get back 5404444444

import (
	"encoding/json"
	"fmt"
)

// Location is the representation of a Location in Yext Location Manager.
// For details see https://www.yext.com/support/platform-api/#Administration_API/Locations.htm
type Location struct {
	// Admin
	Id           *string                `json:"id,omitempty"`
	AccountId    *string                `json:"accountId,omitempty"`
	LocationType *string                `json:"locationType,omitempty"`
	FolderId     *string                `json:"folderId,omitempty"`
	LabelIds     *[]string              `json:"labelIds,omitempty"`
	CategoryIds  *[]string              `json:"categoryIds,omitempty"`
	Closed       *LocationClosed        `json:"closed,omitempty"`
	Keywords     *[]string              `json:"keywords,omitempty"`
	Language     *string                `json:"language,omitempty"`
	CustomFields map[string]interface{} `json:"customFields,omitempty"`

	hydrated bool

	// Address Fields
	Name            *string `json:"locationName,omitempty"`
	Address         *string `json:"address,omitempty"`
	Address2        *string `json:"address2,omitempty"`
	DisplayAddress  *string `json:"displayAddress,omitempty"`
	City            *string `json:"city,omitempty"`
	State           *string `json:"state,omitempty"`
	Sublocality     *string `json:"sublocality,omitempty"`
	Zip             *string `json:"zip,omitempty"`
	CountryCode     *string `json:"countryCode,omitempty"`
	SuppressAddress *bool   `json:"suppressAddress,omitempty"`

	// Other Contact Info
	AlternatePhone *string   `json:"alternatePhone,omitempty"`
	FaxPhone       *string   `json:"faxPhone,omitempty"`
	LocalPhone     *string   `json:"localPhone,omitempty"`
	MobilePhone    *string   `json:"mobilePhone,omitempty"`
	Phone          *string   `json:"phone,omitempty"`
	TollFreePhone  *string   `json:"tollFreePhone,omitempty"`
	TtyPhone       *string   `json:"ttyPhone,omitempty"`
	IsPhoneTracked *bool     `json:"isPhoneTracked,omitempty"`
	Emails         *[]string `json:"emails,omitempty"`

	// HealthCare fields
	FirstName            *string        `json:"firstName,omitempty"`
	MiddleName           *string        `json:"middleName,omitempty"`
	LastName             *string        `json:"lastName,omitempty"`
	Gender               *string        `json:"gender,omitempty"`
	Headshot             *LocationPhoto `json:"headshot,omitempty"`
	AcceptingNewPatients *bool          `json:"acceptingNewPatients,omitempty"`
	AdmittingHospitals   *[]string      `json:"admittingHospitals,omitempty"`
	ConditionsTreated    *[]string      `json:"conditionsTreated,omitempty"`
	InsuranceAccepted    *[]string      `json:"insuranceAccepted,omitempty"`
	NPI                  *string        `json:"npi,omitempty"`
	OfficeName           *string        `json:"officeName,omitempty"`

	// Location Info
	Description         *string        `json:"description,omitempty"`
	HolidayHours        []HolidayHours `json:"holidayHours,omitempty"`
	Hours               *string        `json:"hours,omitempty"`
	AdditionalHoursText *string        `json:"additionalHoursText,omitempty"`
	YearEstablished     *string        `json:"yearEstablished,omitempty"`
	Associations        *[]string      `json:"associations,omitempty"`
	Certifications      *[]string      `json:"certifications,omitempty"`
	Brands              *[]string      `json:"brands,omitempty"`
	Products            *[]string      `json:"products,omitempty"`
	Services            *[]string      `json:"services,omitempty"`
	Specialties         *[]string      `json:"specialties,omitempty"`
	Languages           *[]string      `json:"languages,omitempty"`
	Logo                *LocationPhoto `json:"logo,omitempty"`
	PaymentOptions      *[]string      `json:"paymentOptions,omitempty"`

	// Lats & Lngs
	DisplayLat  *float64 `json:"displayLat,omitempty"`
	DisplayLng  *float64 `json:"displayLng,omitempty"`
	DropoffLat  *float64 `json:"dropoffLat,omitempty"`
	DropoffLng  *float64 `json:"dropoffLng,omitempty"`
	WalkableLat *float64 `json:"walkableLat,omitempty"`
	WalkableLng *float64 `json:"walkableLng,omitempty"`
	RoutableLat *float64 `json:"routableLat,omitempty"`
	RoutableLng *float64 `json:"routableLng,omitempty"`
	PickupLat   *float64 `json:"pickupLat,omitempty"`
	PickupLng   *float64 `json:"pickupLng,omitempty"`

	// ECLS
	BioListIds        *[]string `json:"bioListIds,omitempty"`
	BioListsLabel     *string   `json:"bioListsLabel,omitempty"`
	EventListIds      *[]string `json:"eventListIds,omitempty"`
	EventListsLabel   *string   `json:"eventListsLabel,omitempty"`
	MenusLabel        *string   `json:"menusLabel,omitempty"`
	MenuIds           *[]string `json:"menuIds,omitempty"`
	ProductListIds    *[]string `json:"productListIds,omitempty"`
	ProductListsLabel *string   `json:"productListsLabel,omitempty"`

	// Urls
	MenuUrl               *string `json:"menuUrl,omitempty"`
	DisplayMenuUrl        *string `json:"displayMenuUrl,omitempty"`
	OrderUrl              *string `json:"orderUrl,omitempty"`
	DisplayOrderUrl       *string `json:"displayOrderUrl,omitempty"`
	ReservationUrl        *string `json:"reservationUrl,omitempty"`
	DisplayReservationUrl *string `json:"displayReservationUrl,omitempty"`
	DisplayWebsiteUrl     *string `json:"displayWebsiteUrl,omitempty"`
	WebsiteUrl            *string `json:"websiteUrl,omitempty"`
	FeaturedMessage       *string `json:"featuredMessage,omitempty"`
	FeaturedMessageUrl    *string `json:"featuredMessageUrl,omitempty"`

	// Uber
	UberClientId         *string `json:"uberClientId,omitempty"`
	UberLinkText         *string `json:"uberLinkText,omitempty"`
	UberLinkType         *string `json:"uberLinkType,omitempty"`
	UberTripBrandingText *string `json:"uberTripBrandingText,omitempty"`
	UberTripBrandingUrl  *string `json:"uberTripBrandingUrl,omitempty"`

	// Social Media
	FacebookCoverPhoto     *LocationPhoto `json:"facebookCoverPhoto,omitempty"`
	FacebookPageUrl        *string        `json:"facebookPageUrl,omitempty"`
	FacebookProfilePicture *LocationPhoto `json:"facebookProfilePicture,omitempty"`

	GoogleCoverPhoto      *LocationPhoto `json:"googleCoverPhoto,omitempty"`
	GooglePreferredPhoto  *string        `json:"googlePreferredPhoto,omitempty"`
	GoogleProfilePhoto    *LocationPhoto `json:"googleProfilePhoto,omitempty"`
	GoogleWebsiteOverride *string        `json:"googleWebsiteOverride,omitempty"`

	InstagramHandle *string `json:"instagramHandle,omitempty"`
	TwitterHandle   *string `json:"twitterHandle,omitempty"`

	Photos    []LocationPhoto `json:"photos,omitempty"`
	VideoUrls *[]string       `json:"videoUrls,omitempty"`

	/** TODO(bmcginnis) add the following fields:

	   ServiceArea       struct {
	 		Places *[]string `json:"places,omitempty"`
	 		Radius *int      `json:"radius,omitempty"`
	 		Unit   *string   `json:"unit,omitempty"`
	 	} `json:"serviceArea,omitempty"`

	  EducationList         []struct {
			InstitutionName *string `json:"institutionName,omitempty"`
			Type            *string `json:"type,omitempty"`
			YearCompleted   *string `json:"yearCompleted,omitempty"`
		} `json:"educationList,omitempty"`

	  GoogleAttributes   []struct {
	    Id        *string   `json:"id,omitempty"`
	    OptionIds *[]string `json:"optionIds,omitempty"`
	  } `json:"googleAttributes,omitempty"`
	*/
}

func (y Location) GetId() string {
	if y.Id != nil {
		return *y.Id
	}
	return ""
}

func (y Location) GetLocationType() string {
	if y.LocationType != nil {
		return *y.LocationType
	}
	return ""
}

func (y Location) GetName() string {
	if y.Name != nil {
		return *y.Name
	}
	return ""
}

func (y Location) GetFirstName() string {
	if y.FirstName != nil {
		return *y.FirstName
	}
	return ""
}

func (y Location) GetMiddleName() string {
	if y.MiddleName != nil {
		return *y.MiddleName
	}
	return ""
}

func (y Location) GetLastName() string {
	if y.LastName != nil {
		return *y.LastName
	}
	return ""
}

func (y Location) GetNPI() string {
	if y.NPI != nil {
		return *y.NPI
	}
	return ""
}

func (y Location) GetAccountId() string {
	if y.AccountId != nil {
		return *y.AccountId
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

func (y Location) GetFeaturedMessage() string {
	if y.FeaturedMessage != nil {
		return *y.FeaturedMessage
	}
	return ""
}

func (y Location) GetFeaturedMessageUrl() string {
	if y.FeaturedMessageUrl != nil {
		return *y.FeaturedMessageUrl
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

func (y Location) GetKeywords() (v []string) {
	if y.Keywords != nil {
		v = *y.Keywords
	}
	return v
}

func (y Location) GetAssociations() (v []string) {
	if y.Associations != nil {
		v = *y.Associations
	}
	return v
}

func (y Location) GetEmails() (v []string) {
	if y.Emails != nil {
		v = *y.Emails
	}
	return v
}

func (y Location) GetSpecialties() (v []string) {
	if y.Specialties != nil {
		v = *y.Specialties
	}
	return v
}

func (y Location) GetServices() (v []string) {
	if y.Services != nil {
		v = *y.Services
	}
	return v
}

func (y Location) GetBrands() (v []string) {
	if y.Brands != nil {
		v = *y.Brands
	}
	return v
}

func (y Location) GetLanguages() (v []string) {
	if y.Languages != nil {
		v = *y.Languages
	}
	return v
}

func (y Location) GetLabelIds() (v []string) {
	if y.LabelIds != nil {
		v = *y.LabelIds
	}
	return v
}

func (y Location) GetCategoryIds() (v []string) {
	if y.CategoryIds != nil {
		v = *y.CategoryIds
	}
	return v
}

func (y Location) GetPaymentOptions() (v []string) {
	if y.PaymentOptions != nil {
		v = *y.PaymentOptions
	}
	return v
}

func (y Location) GetVideoUrls() (v []string) {
	if y.VideoUrls != nil {
		v = *y.VideoUrls
	}
	return v
}

func (y Location) GetAdmittingHospitals() (v []string) {
	if y.AdmittingHospitals != nil {
		v = *y.AdmittingHospitals
	}
	return v
}

// LocationPhoto represents a photo associated with a Location in Yext Location Manager.
// For details see https://www.yext.com/support/platform-api/#Administration_API/Locations.htm#Photo
type LocationPhoto struct {
	Url             string `json:"url,omitempty"`
	Description     string `json:"description,omitempty"`
	AlternateText   string `json:"alternateText"`
	ClickThroughUrl string `json:"clickthroughUrl,omitempty"`
}

func (l Photo) String() string {
	return fmt.Sprintf("Url: '%v', Description: '%v'", l.Url, l.Description)
}

// LocationClosed represents the 'closed' state of a Location in Yext Location Manager.
// For details see https://www.yext.com/support/platform-api/#Administration_API/Locations.htm#Closed
type LocationClosed struct {
	IsClosed   bool   `json:"isClosed"`
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
