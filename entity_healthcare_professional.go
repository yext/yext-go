package yext

// TODO
// * Need better custom field accessors and helpers
// * The API will accept some things and return them in a different format - this makes diff'ing difficult:
// ** Phone: Send in 540-444-4444, get back 5404444444

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	EDUCATION_FELLOWSHIP                         = "FELLOWSHIP"
	EDUCATION_INTERNSHIP                         = "INTERNSHIP"
	EDUCATION_MEDICAL_SCHOOL                     = "MEDICAL_SCHOOL"
	EDUCATION_RESIDENCY                          = "RESIDENCY"
	ENTITYTYPE_HEALTHCAREPROFESSIONAL EntityType = "healthcareProfessional"
	GENDER_MALE                                  = "MALE"
	GENDER_FEMALE                                = "FEMALE"
	GENDER_UNSPECIFIED                           = "UNSPECIFIED"
	GENDER_NONBINARY                             = "NONBINARY"
)

type HealthcareProfessionalEntity struct {
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
	Services            *[]string `json:"services,omitempty"`
	// Spelling of json tag 'specialities' is intentional to match mispelling in Yext API
	Specialties    *[]string `json:"specialities,omitempty"`
	Languages      *[]string `json:"languages,omitempty"`
	Logo           **Photo   `json:"logo,omitempty"`
	PaymentOptions *[]string `json:"paymentOptions,omitempty"`

	// Healthcare
	FirstName            *string           `json:"firstName,omitempty"`
	MiddleName           *string           `json:"middleName,omitempty"`
	LastName             *string           `json:"lastName,omitempty"`
	Gender               **string          `json:"gender,omitempty"`
	Headshot             **Image           `json:"headshot,omitempty"`
	AcceptingNewPatients *bool             `json:"acceptingNewPatients,omitempty"`
	AdmittingHospitals   *[]string         `json:"admittingHospitals,omitempty"`
	ConditionsTreated    *[]string         `json:"conditionsTreated,omitempty"`
	InsuranceAccepted    *[]string         `json:"insuranceAccepted,omitempty"`
	NPI                  *string           `json:"npi,omitempty"`
	OfficeName           *string           `json:"officeName,omitempty"`
	Degrees              *UnorderedStrings `json:"degrees,omitempty"`
	EducationList        *[]Education      `json:"educationList,omitempty"`

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
	WebsiteUrl      **Website         `json:"websiteUrl,omitempty"`
	FeaturedMessage **FeaturedMessage `json:"featuredMessage,omitempty"`
	ReservationUrl  **Website         `json:"reservationUrl,omitempty"`
	TelehealthUrl   *string           `json:"telehealthUrl,omitempty"`

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

type Education struct {
	InstitutionName string `json:"institutionName,omitempty"`
	Type            string `json:"type,omitempty"`
	YearCompleted   int    `json:"yearCompleted,omitempty"`
}

func (e Education) String() string {
	return fmt.Sprintf("Institution Name: '%v', Type: '%v', Year Completed: '%v'", e.InstitutionName, e.Type, e.YearCompleted)
}

// Equal compares Education
func (a *Education) Equal(b Comparable) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Value of A: %+v, Value of B:%+v, Type Of A: %T, Type Of B: %T\n", a, b, a, b)
			panic(r)
		}
	}()

	if a == nil || b == nil {
		return false
	}

	var (
		u = Education(*a)
		s = Education(*b.(*Education))
	)
	if u.InstitutionName != s.InstitutionName {
		return false
	}

	if u.Type != s.Type {
		return false
	}

	if u.YearCompleted != s.YearCompleted {
		return false
	}

	return true
}

func (h *HealthcareProfessionalEntity) UnmarshalJSON(data []byte) error {
	type Alias HealthcareProfessionalEntity
	a := &struct {
		*Alias
	}{
		Alias: (*Alias)(h),
	}
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	return UnmarshalEntityJSON(h, data)
}

func (y HealthcareProfessionalEntity) GetId() string {
	if y.BaseEntity.Meta != nil && y.BaseEntity.Meta.Id != nil {
		return *y.BaseEntity.Meta.Id
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetName() string {
	if y.Name != nil {
		return GetString(y.Name)
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetAccountId() string {
	if y.BaseEntity.Meta != nil && y.BaseEntity.Meta.AccountId != nil {
		return *y.BaseEntity.Meta.AccountId
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetLine1() string {
	if y.Address != nil && y.Address.Line1 != nil {
		return GetString(y.Address.Line1)
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetLine2() string {
	if y.Address != nil && y.Address.Line2 != nil {
		return GetString(y.Address.Line2)
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetAddressHidden() bool {
	return GetNullableBool(y.AddressHidden)
}

func (y HealthcareProfessionalEntity) GetExtraDescription() string {
	if y.Address != nil && y.Address.ExtraDescription != nil {
		return GetString(y.Address.ExtraDescription)
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetCity() string {
	if y.Address != nil && y.Address.City != nil {
		return GetString(y.Address.City)
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetRegion() string {
	if y.Address != nil && y.Address.Region != nil {
		return GetString(y.Address.Region)
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetCountryCode() string {
	if y.BaseEntity.Meta != nil && y.BaseEntity.Meta.CountryCode != nil {
		return GetString(y.BaseEntity.Meta.CountryCode)
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetPostalCode() string {
	if y.Address != nil && y.Address.PostalCode != nil {
		return GetString(y.Address.PostalCode)
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetMainPhone() string {
	if y.MainPhone != nil {
		return *y.MainPhone
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetLocalPhone() string {
	if y.LocalPhone != nil {
		return *y.LocalPhone
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetAlternatePhone() string {
	if y.AlternatePhone != nil {
		return *y.AlternatePhone
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetFax() string {
	if y.Fax != nil {
		return *y.Fax
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetMobilePhone() string {
	if y.MobilePhone != nil {
		return *y.MobilePhone
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetTollFreePhone() string {
	if y.TollFreePhone != nil {
		return *y.TollFreePhone
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetTtyPhone() string {
	if y.TtyPhone != nil {
		return *y.TtyPhone
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetFeaturedMessage() string {
	f := GetFeaturedMessage(y.FeaturedMessage)
	if f != nil {
		return GetString(f.Description)
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetFeaturedMessageUrl() string {
	f := GetFeaturedMessage(y.FeaturedMessage)
	if f != nil {
		return GetString(f.Url)
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetWebsiteUrl() string {
	w := GetWebsite(y.WebsiteUrl)
	if w != nil {
		return GetString(w.Url)
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetDisplayWebsiteUrl() string {
	w := GetWebsite(y.WebsiteUrl)
	if w != nil {
		return GetString(w.DisplayUrl)
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetReservationUrl() string {
	w := GetWebsite(y.ReservationUrl)
	if w != nil {
		return GetString(w.Url)
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetDisplayReservationUrl() string {
	w := GetWebsite(y.ReservationUrl)
	if w != nil {
		return GetString(w.DisplayUrl)
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetHours() *Hours {
	return GetHours(y.Hours)
}

func (y HealthcareProfessionalEntity) GetAdditionalHoursText() string {
	if y.AdditionalHoursText != nil {
		return *y.AdditionalHoursText
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetDescription() string {
	if y.Description != nil {
		return *y.Description
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetTwitterHandle() string {
	if y.TwitterHandle != nil {
		return *y.TwitterHandle
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetFacebookPageUrl() string {
	if y.FacebookPageUrl != nil {
		return *y.FacebookPageUrl
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetYearEstablished() float64 {
	if y.YearEstablished != nil {
		return GetNullableFloat(y.YearEstablished)
	}
	return 0
}

func (y HealthcareProfessionalEntity) GetDisplayLat() float64 {
	c := GetCoordinate(y.DisplayCoordinate)
	if c != nil {
		return GetNullableFloat(c.Latitude)
	}
	return 0
}

func (y HealthcareProfessionalEntity) GetDisplayLng() float64 {
	c := GetCoordinate(y.DisplayCoordinate)
	if c != nil {
		return GetNullableFloat(c.Longitude)
	}
	return 0
}

func (y HealthcareProfessionalEntity) GetRoutableLat() float64 {
	c := GetCoordinate(y.RoutableCoordinate)
	if c != nil {
		return GetNullableFloat(c.Latitude)
	}
	return 0
}

func (y HealthcareProfessionalEntity) GetRoutableLng() float64 {
	c := GetCoordinate(y.RoutableCoordinate)
	if c != nil {
		return GetNullableFloat(c.Longitude)
	}
	return 0
}

func (y HealthcareProfessionalEntity) GetBios() (v *Lists) {
	return GetLists(y.Bios)
}

func (y HealthcareProfessionalEntity) GetCalendars() (v *Lists) {
	return GetLists(y.Calendars)
}

func (y HealthcareProfessionalEntity) GetProductLists() (v *Lists) {
	return GetLists(y.ProductLists)
}

func (y HealthcareProfessionalEntity) GetReviewGenerationUrl() string {
	if y.ReviewGenerationUrl != nil {
		return *y.ReviewGenerationUrl
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetFirstPartyReviewPage() string {
	if y.FirstPartyReviewPage != nil {
		return *y.FirstPartyReviewPage
	}
	return ""
}

func (y HealthcareProfessionalEntity) String() string {
	b, _ := json.Marshal(y)
	return string(b)
}

func (y HealthcareProfessionalEntity) GetKeywords() (v []string) {
	if y.Keywords != nil {
		v = *y.Keywords
	}
	return v
}

func (y HealthcareProfessionalEntity) GetLanguage() (v string) {
	if y.BaseEntity.Meta.Language != nil {
		v = *y.BaseEntity.Meta.Language
	}
	return v
}

func (y HealthcareProfessionalEntity) GetAssociations() (v []string) {
	if y.Associations != nil {
		v = *y.Associations
	}
	return v
}

func (y HealthcareProfessionalEntity) GetEmails() (v []string) {
	if y.Emails != nil {
		v = *y.Emails
	}
	return v
}

func (y HealthcareProfessionalEntity) GetSpecialties() (v []string) {
	if y.Specialties != nil {
		v = *y.Specialties
	}
	return v
}

func (y HealthcareProfessionalEntity) GetServices() (v []string) {
	if y.Services != nil {
		v = *y.Services
	}
	return v
}

func (y HealthcareProfessionalEntity) GetBrands() (v []string) {
	if y.Brands != nil {
		v = *y.Brands
	}
	return v
}

func (y HealthcareProfessionalEntity) GetLanguages() (v []string) {
	if y.Languages != nil {
		v = *y.Languages
	}
	return v
}

func (y HealthcareProfessionalEntity) GetPaymentOptions() (v []string) {
	if y.PaymentOptions != nil {
		v = *y.PaymentOptions
	}
	return v
}

func (y HealthcareProfessionalEntity) GetVideos() (v []Video) {
	if y.Videos != nil {
		v = *y.Videos
	}
	return v
}

func (y HealthcareProfessionalEntity) GetGoogleAttributes() map[string][]string {
	if y.GoogleAttributes != nil {
		return *y.GoogleAttributes
	}
	return nil
}

func (y HealthcareProfessionalEntity) GetHolidayHours() []HolidayHours {
	h := GetHours(y.Hours)
	if h != nil && h.HolidayHours != nil {
		return *h.HolidayHours
	}
	return nil
}

func (y HealthcareProfessionalEntity) GetNPI() string {
	if y.NPI != nil {
		return *y.NPI
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetAdmittingHospitals() (v []string) {
	if y.AdmittingHospitals != nil {
		v = *y.AdmittingHospitals
	}
	return v
}

func (y HealthcareProfessionalEntity) IsClosed() bool {
	return GetNullableBool(y.Closed)
}

func (y HealthcareProfessionalEntity) GetCategoryIds() (v []string) {
	if y.CategoryIds != nil {
		v = *y.CategoryIds
	}
	return v
}

// sets all degrees that are valid
// returns invalid degrees
// handle parseError for invalid degrees outside of function call
func (y *HealthcareProfessionalEntity) SetValidDegrees(degrees []string) []string {
	var (
		invalidDegrees   = []string{}
		validDegrees     = []string{}
		validYextDegrees = []string{"ANP", "APN", "APRN", "ARNP", "CNM", "CNP", "CNS", "CPNP", "CRNA", "CRNP", "DC",
			"DDS", "DMD", "DNP", "DO", "DPM", "DVM", "FNP", "GNP", "LAC", "LPN", "MBA", "MBBS", "MD",
			"MPH", "ND", "NNP", "NP", "OD", "PA", "PAC", "PHARMD", "PHD", "PNP", "PSYD", "VMD", "WHNP"}
	)

	contains := func(list []string, item string) bool {
		for _, l := range list {
			if l == item {
				return true
			}
		}
		return false
	}

	for _, degree := range degrees {
		if contains(validYextDegrees, strings.ToUpper(degree)) {
			validDegrees = append(validDegrees, degree)
		} else {
			invalidDegrees = append(invalidDegrees, degree)
		}
	}

	y.Degrees = ToUnorderedStrings(validDegrees)

	return invalidDegrees
}

func (y HealthcareProfessionalEntity) GetGooglePlaceId() string {
	if y.GooglePlaceId != nil {
		return *y.GooglePlaceId
	}
	return ""
}

func (y HealthcareProfessionalEntity) GetGeomodifier() string {
	if y.Geomodifier != nil {
		return GetString(y.Geomodifier)
	}
	return ""
}
