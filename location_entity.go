package yext

// TODO
// * Need better custom field accessors and helpers
// * The API will accept some things and return them in a different format - this makes diff'ing difficult:
// ** Phone: Send in 540-444-4444, get back 5404444444

import (
	"encoding/json"
)

const ENTITYTYPE_LOCATION EntityType = "location"

// Location is the representation of a Location in Yext Location Manager.
// For details see https://www.yext.com/support/platform-api/#Administration_API/Locations.htm
type LocationEntity struct {
	BaseEntity

	// Admin
	FolderId    *string           `json:"folderId,omitempty"`
	LabelIds    *UnorderedStrings `json:"labelIds,omitempty"`
	CategoryIds *[]string         `json:"categoryIds,omitempty"`
	Closed      *LocationClosed   `json:"closed,omitempty"`
	Keywords    *[]string         `json:"keywords,omitempty"`
	Language    *string           `json:"language,omitempty"`

	// hydrated   bool
	// nilIsEmpty bool

	// Address Fields
	Name            *string  `json:"name,omitempty"`
	Address         *Address `json:"address,omitempty"`
	DisplayAddress  *string  `json:"displayAddress,omitempty"`
	CountryCode     *string  `json:"countryCode,omitempty"`
	SuppressAddress *bool    `json:"suppressAddress,omitempty"`

	// Other Contact Info
	AlternatePhone *string   `json:"alternatePhone,omitempty"`
	FaxPhone       *string   `json:"faxPhone,omitempty"`
	LocalPhone     *string   `json:"localPhone,omitempty"`
	MobilePhone    *string   `json:"mobilePhone,omitempty"`
	MainPhone      *string   `json:"mainPhone,omitempty"`
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
	Degrees              *[]string      `json:"degrees,omitempty"`

	// Location Info
	Description         *string        `json:"description,omitempty"`
	Hours               *Hours         `json:"hours,omitempty"`
	AdditionalHoursText *string        `json:"additionalHoursText,omitempty"`
	YearEstablished     *float64       `json:"yearEstablished,omitempty"`
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
	DisplayCoordinate *Coordinate `json:"yextDisplayCoordinate,omitempty"`
	// TODO: Update below
	DropoffLat         *float64    `json:"dropoffLat,omitempty"`
	DropoffLng         *float64    `json:"dropoffLng,omitempty"`
	WalkableLat        *float64    `json:"walkableLat,omitempty"`
	WalkableLng        *float64    `json:"walkableLng,omitempty"`
	RoutableCoordinate *Coordinate `json:"yextRoutableCoordinate,omitempty"`
	PickupLat          *float64    `json:"pickupLat,omitempty"`
	PickupLng          *float64    `json:"pickupLng,omitempty"`

	// ECLS
	BioListIds        *[]string `json:"bioListIds,omitempty"`
	BioListsLabel     *string   `json:"bioListsLabel,omitempty"`
	EventListIds      *[]string `json:"eventListIds,omitempty"`
	EventListsLabel   *string   `json:"eventListsLabel,omitempty"`
	MenuListsLabel    *string   `json:"menusLabel,omitempty"`
	MenuListIds       *[]string `json:"menuIds,omitempty"`
	ProductListIds    *[]string `json:"productListIds,omitempty"`
	ProductListsLabel *string   `json:"productListsLabel,omitempty"`

	// Urls
	MenuUrl         *Website         `json:"menuUrl,omitempty"`
	OrderUrl        *Website         `json:"orderUrl,omitempty"`
	ReservationUrl  *Website         `json:"reservationUrl,omitempty"`
	WebsiteUrl      *Website         `json:"websiteUrl,omitempty"`
	FeaturedMessage *FeaturedMessage `json:"featuredMessage,omitempty"`

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

	Photos    *[]LocationPhoto `json:"photos,omitempty"`
	VideoUrls *[]string        `json:"videoUrls,omitempty"`

	GoogleAttributes *GoogleAttributes `json:"googleAttributes,omitempty"`

	// Reviews
	ReviewBalancingURL   *string `json:"reviewBalancingURL,omitempty"`
	FirstPartyReviewPage *string `json:"firstPartyReviewPage,omitempty"`

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
	*/
}

type Address struct {
	Line1       *string `json:"line1,omitempty"`
	Line2       *string `json:"line2,omitempty"`
	City        *string `json:"city,omitempty"`
	Region      *string `json:"region,omitempty"`
	Sublocality *string `json:"sublocality,omitempty"`
	PostalCode  *string `json:"postalCode,omitempty"`
}

type FeaturedMessage struct {
	Description *string `json:"description,omitempty"`
	Url         *string `json:"url,omitempty"`
}

type Website struct {
	DisplayUrl       *string `json:"displayUrl,omitempty"`
	Url              *string `json:"url,omitempty"`
	PreferDisplayUrl *bool   `json:"preferDisplayUrl,omitempty"`
}

type Coordinate struct {
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
}

type Hours struct {
	Monday       *DayHours       `json:"monday,omitempty"`
	Tuesday      *DayHours       `json:"tuesday,omitempty"`
	Wednesday    *DayHours       `json:"wednesday,omitempty"`
	Thursday     *DayHours       `json:"thursday,omitempty"`
	Friday       *DayHours       `json:"friday,omitempty"`
	Saturday     *DayHours       `json:"saturday,omitempty"`
	Sunday       *DayHours       `json:"sunday,omitempty"`
	HolidayHours *[]HolidayHours `json:"holidayHours,omitempty"`
}

type DayHours struct {
	OpenIntervals []*Interval `json:"openIntervals,omitempty"`
	IsClosed      *bool       `json:"isClosed,omitempty"`
}

type Interval struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

func (y LocationEntity) GetId() string {
	if y.BaseEntity.Meta != nil && y.BaseEntity.Meta.Id != nil {
		return *y.BaseEntity.Meta.Id
	}
	return ""
}

func (y LocationEntity) GetName() string {
	if y.Name != nil {
		return *y.Name
	}
	return ""
}

func (y LocationEntity) GetFirstName() string {
	if y.FirstName != nil {
		return *y.FirstName
	}
	return ""
}

func (y LocationEntity) GetMiddleName() string {
	if y.MiddleName != nil {
		return *y.MiddleName
	}
	return ""
}

func (y LocationEntity) GetLastName() string {
	if y.LastName != nil {
		return *y.LastName
	}
	return ""
}

func (y LocationEntity) GetGender() string {
	if y.Gender != nil {
		return *y.Gender
	}
	return ""
}

func (y LocationEntity) GetAcceptingNewPatients() bool {
	if y.AcceptingNewPatients != nil {
		return *y.AcceptingNewPatients
	}
	return false
}

func (y LocationEntity) GetCertifications() []string {
	if y.Certifications != nil {
		return *y.Certifications
	}
	return nil
}

func (y LocationEntity) GetNPI() string {
	if y.NPI != nil {
		return *y.NPI
	}
	return ""
}

func (y LocationEntity) GetOfficeName() string {
	if y.OfficeName != nil {
		return *y.OfficeName
	}
	return ""
}

func (y LocationEntity) GetDegrees() []string {
	if y.Degrees != nil {
		return *y.Degrees
	}
	return nil
}

func (y LocationEntity) GetAccountId() string {
	if y.BaseEntity.Meta != nil && y.BaseEntity.Meta.AccountId != nil {
		return *y.BaseEntity.Meta.AccountId
	}
	return ""
}

func (y LocationEntity) GetLine1() string {
	if y.Address != nil && y.Address.Line1 != nil {
		return *y.Address.Line1
	}
	return ""
}

func (y LocationEntity) GetLine2() string {
	if y.Address != nil && y.Address.Line2 != nil {
		return *y.Address.Line2
	}
	return ""
}

func (y LocationEntity) GetSuppressAddress() bool {
	if y.SuppressAddress != nil {
		return *y.SuppressAddress
	}
	return false
}

func (y LocationEntity) GetDisplayAddress() string {
	if y.DisplayAddress != nil {
		return *y.DisplayAddress
	}
	return ""
}

func (y LocationEntity) GetCity() string {
	if y.Address != nil && y.Address.City != nil {
		return *y.Address.City
	}
	return ""
}

func (y LocationEntity) GetRegion() string {
	if y.Address != nil && y.Address.Region != nil {
		return *y.Address.Region
	}
	return ""
}

func (y LocationEntity) GetPostalCode() string {
	if y.Address != nil && y.Address.PostalCode != nil {
		return *y.Address.PostalCode
	}
	return ""
}

func (y LocationEntity) GetCountryCode() string {
	if y.CountryCode != nil {
		return *y.CountryCode
	}
	return ""
}

func (y LocationEntity) GetMainPhone() string {
	if y.MainPhone != nil {
		return *y.MainPhone
	}
	return ""
}

func (y LocationEntity) GetIsPhoneTracked() bool {
	if y.IsPhoneTracked != nil {
		return *y.IsPhoneTracked
	}
	return false
}

func (y LocationEntity) GetLocalPhone() string {
	if y.LocalPhone != nil {
		return *y.LocalPhone
	}
	return ""
}

func (y LocationEntity) GetAlternatePhone() string {
	if y.AlternatePhone != nil {
		return *y.AlternatePhone
	}
	return ""
}

func (y LocationEntity) GetFaxPhone() string {
	if y.FaxPhone != nil {
		return *y.FaxPhone
	}
	return ""
}

func (y LocationEntity) GetMobilePhone() string {
	if y.MobilePhone != nil {
		return *y.MobilePhone
	}
	return ""
}

func (y LocationEntity) GetTollFreePhone() string {
	if y.TollFreePhone != nil {
		return *y.TollFreePhone
	}
	return ""
}

func (y LocationEntity) GetTtyPhone() string {
	if y.TtyPhone != nil {
		return *y.TtyPhone
	}
	return ""
}

func (y LocationEntity) GetFeaturedMessageDescription() string {
	if y.FeaturedMessage != nil && y.FeaturedMessage.Description != nil {
		return *y.FeaturedMessage.Description
	}
	return ""
}

func (y LocationEntity) GetFeaturedMessageUrl() string {
	if y.FeaturedMessage != nil && y.FeaturedMessage.Url != nil {
		return *y.FeaturedMessage.Url
	}
	return ""
}

func (y LocationEntity) GetWebsiteUrl() string {
	if y.WebsiteUrl != nil && y.WebsiteUrl.Url != nil {
		return *y.WebsiteUrl.Url
	}
	return ""
}

func (y LocationEntity) GetDisplayWebsiteUrl() string {
	if y.WebsiteUrl != nil && y.WebsiteUrl.DisplayUrl != nil {
		return *y.WebsiteUrl.DisplayUrl
	}
	return ""
}

func (y LocationEntity) GetReservationUrl() string {
	if y.ReservationUrl != nil && y.ReservationUrl.Url != nil {
		return *y.ReservationUrl.Url
	}
	return ""
}

func (y LocationEntity) GetAdditionalHoursText() string {
	if y.AdditionalHoursText != nil {
		return *y.AdditionalHoursText
	}
	return ""
}

func (y LocationEntity) GetDescription() string {
	if y.Description != nil {
		return *y.Description
	}
	return ""
}

func (y LocationEntity) GetTwitterHandle() string {
	if y.TwitterHandle != nil {
		return *y.TwitterHandle
	}
	return ""
}

func (y LocationEntity) GetFacebookPageUrl() string {
	if y.FacebookPageUrl != nil {
		return *y.FacebookPageUrl
	}
	return ""
}

func (y LocationEntity) GetYearEstablished() float64 {
	if y.YearEstablished != nil {
		return *y.YearEstablished
	}
	return 0
}

func (y LocationEntity) GetDisplayLat() float64 {
	if y.DisplayCoordinate != nil && y.DisplayCoordinate.Latitude != nil {
		return *y.DisplayCoordinate.Latitude
	}
	return 0
}

func (y LocationEntity) GetDisplayLng() float64 {
	if y.DisplayCoordinate != nil && y.DisplayCoordinate.Longitude != nil {
		return *y.DisplayCoordinate.Longitude
	}
	return 0
}

func (y LocationEntity) GetRoutableLat() float64 {
	if y.RoutableCoordinate != nil && y.RoutableCoordinate.Latitude != nil {
		return *y.RoutableCoordinate.Latitude
	}
	return 0
}

func (y LocationEntity) GetRoutableLng() float64 {
	if y.RoutableCoordinate != nil && y.RoutableCoordinate.Longitude != nil {
		return *y.RoutableCoordinate.Longitude
	}
	return 0
}

func (y LocationEntity) GetBioListIds() (v []string) {
	if y.BioListIds != nil {
		v = *y.BioListIds
	}
	return v
}

func (y LocationEntity) GetEventListIds() (v []string) {
	if y.EventListIds != nil {
		v = *y.EventListIds
	}
	return v
}

func (y LocationEntity) GetProductListIds() (v []string) {
	if y.ProductListIds != nil {
		v = *y.ProductListIds
	}
	return v
}

func (y LocationEntity) GetMenuListIds() (v []string) {
	if y.MenuListIds != nil {
		v = *y.MenuListIds
	}
	return v
}

func (y LocationEntity) GetReviewBalancingURL() string {
	if y.ReviewBalancingURL != nil {
		return *y.ReviewBalancingURL
	}
	return ""
}

func (y LocationEntity) GetFirstPartyReviewPage() string {
	if y.FirstPartyReviewPage != nil {
		return *y.FirstPartyReviewPage
	}
	return ""
}

func (y LocationEntity) String() string {
	b, _ := json.Marshal(y)
	return string(b)
}

func (y LocationEntity) GetKeywords() (v []string) {
	if y.Keywords != nil {
		v = *y.Keywords
	}
	return v
}

func (y LocationEntity) GetLanguage() (v string) {
	if y.Language != nil {
		v = *y.Language
	}
	return v
}

func (y LocationEntity) GetAssociations() (v []string) {
	if y.Associations != nil {
		v = *y.Associations
	}
	return v
}

func (y LocationEntity) GetEmails() (v []string) {
	if y.Emails != nil {
		v = *y.Emails
	}
	return v
}

func (y LocationEntity) GetSpecialties() (v []string) {
	if y.Specialties != nil {
		v = *y.Specialties
	}
	return v
}

func (y LocationEntity) GetServices() (v []string) {
	if y.Services != nil {
		v = *y.Services
	}
	return v
}

func (y LocationEntity) GetBrands() (v []string) {
	if y.Brands != nil {
		v = *y.Brands
	}
	return v
}

func (y LocationEntity) GetLanguages() (v []string) {
	if y.Languages != nil {
		v = *y.Languages
	}
	return v
}

func (y LocationEntity) GetFolderId() string {
	if y.FolderId != nil {
		return *y.FolderId
	}
	return ""
}

func (y LocationEntity) GetCategoryIds() (v []string) {
	if y.CategoryIds != nil {
		v = *y.CategoryIds
	}
	return v
}

func (y LocationEntity) GetLabelIds() (v UnorderedStrings) {
	if y.LabelIds != nil {
		v = *y.LabelIds
	}
	return v
}

func (y LocationEntity) SetLabelIds(v []string) {
	l := UnorderedStrings(v)
	y.SetLabelIdsWithUnorderedStrings(l)
}

func (y LocationEntity) SetLabelIdsWithUnorderedStrings(v UnorderedStrings) {
	y.LabelIds = &v
}

func (y LocationEntity) GetPaymentOptions() (v []string) {
	if y.PaymentOptions != nil {
		v = *y.PaymentOptions
	}
	return v
}

func (y LocationEntity) GetVideoUrls() (v []string) {
	if y.VideoUrls != nil {
		v = *y.VideoUrls
	}
	return v
}

func (y LocationEntity) GetAdmittingHospitals() (v []string) {
	if y.AdmittingHospitals != nil {
		v = *y.AdmittingHospitals
	}
	return v
}

func (y LocationEntity) GetGoogleAttributes() GoogleAttributes {
	if y.GoogleAttributes != nil {
		return *y.GoogleAttributes
	}
	return nil
}

func (y LocationEntity) GetHolidayHours() []HolidayHours {
	if y.Hours != nil && y.Hours.HolidayHours != nil {
		return *y.Hours.HolidayHours
	}
	return nil
}

func (y LocationEntity) IsClosed() bool {
	if y.Closed != nil && y.Closed.IsClosed {
		return true
	}
	return false
}

// HolidayHours represents individual exceptions to a Location's regular hours in Yext Location Manager.
// For details see
type HolidayHours struct {
	Date  string      `json:"date"`
	Hours []*Interval `json:"hours"`
}
