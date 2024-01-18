package yext

import (
	"encoding/json"
)

const (
	ENTITYTYPE_FINANCIALPROFESSIONAL EntityType = "financialProfessional"
)

// FinancialProfessional is the representation of a FinancialProfessional in Yext Location Manager.
type FinancialProfessional struct {
	BaseEntity
	// Address Fields
	Name              *string             `json:"name,omitempty"`
	Address           *Address            `json:"address,omitempty"`
	AddressHidden     **bool              `json:"addressHidden,omitempty"`
	ISORegionCode     *string             `json:"isoRegionCode,omitempty"`
	ServiceAreaPlaces *[]ServiceAreaPlace `json:"serviceAreaPlaces,omitempty"`
	Geomodifier       *string             `json:"geomodifier,omitempty"`
	Impressum         *string             `json:"impressum,omitempty"`
	Neighborhood      *string             `json:"neighborhood,omitempty"`

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
	NMLSNumber          *string   `json:"nmlsNumber,omitempty"`
	TeamNumber          *string   `json:"teamNumber,omitempty"`
	Hours               **Hours   `json:"hours,omitempty"`
	OnlineServiceHours  **Hours   `json:"onlineServiceHours,omitempty"`
	Closed              **bool    `json:"closed,omitempty"`
	Description         *string   `json:"description,omitempty"`
	AdditionalHoursText *string   `json:"additionalHoursText,omitempty"`
	Associations        *[]string `json:"associations,omitempty"`
	Certifications      *[]string `json:"certifications,omitempty"`
	Brands              *[]string `json:"brands,omitempty"`
	Products            *[]string `json:"products,omitempty"`
	Services            *[]string `json:"services,omitempty"`
	// Spelling of json tag 'specialities' is intentional to match mispelling in Yext API
	Specialties *[]string `json:"specialities,omitempty"`

	// Lists
	ProductLists **Lists `json:"productLists,omitempty"`

	Languages      *[]string `json:"languages,omitempty"`
	Logo           **Photo   `json:"logo,omitempty"`
	PaymentOptions *[]string `json:"paymentOptions,omitempty"`

	// Admin
	Keywords    *[]string `json:"keywords,omitempty"`
	CategoryIds *[]string `json:"categoryIds,omitempty"`

	// Urls
	MenuUrl         **Website         `json:"menuUrl,omitempty"`
	OrderUrl        **Website         `json:"orderUrl,omitempty"`
	ReservationUrl  **Website         `json:"reservationUrl,omitempty"`
	WebsiteUrl      **Website         `json:"websiteUrl,omitempty"`
	LandingPageUrl  *string           `json:"landingPageUrl,omitempty"`
	IOSAppURL       *string           `json:"iosAppUrl,omitempty"`
	AndroidAppURL   *string           `json:"androidAppUrl,omitempty"`
	DisclosureLink  *string           `json:"disclosureLink,omitempty"`
	FeaturedMessage **FeaturedMessage `json:"featuredMessage,omitempty"`

	// Social Media
	FacebookCallToAction **FacebookCTA `json:"facebookCallToAction,omitempty"`
	FacebookCoverPhoto   **Image       `json:"facebookCoverPhoto,omitempty"`
	FacebookPageUrl      *string       `json:"facebookPageUrl,omitempty"`
	FacebookProfilePhoto **Image       `json:"facebookProfilePhoto,omitempty"`
	FacebookOverrideCity *string       `json:"facebookOverrideCity,omitempty"`
	FacebookName         *string       `json:"facebookName,omitempty"`
	FacebookDescriptor   *string       `json:"facebookDescriptor,omitempty"`
	FacebookVanityURL    *string       `json:"facebookVanityUrl,omitempty"`
	FacebookStoreID      *string       `json:"facebookStoreId,omitempty"`
	FacebookAbout        *string       `json:"facebookAbout,omitempty"`
	FacebookParentPageId *string       `json:"facebookParentPageId,omitempty"`

	GoogleCoverPhoto       **Image   `json:"googleCoverPhoto,omitempty"`
	GooglePreferredPhoto   *string   `json:"googlePreferredPhoto,omitempty"`
	GoogleProfilePhoto     **Image   `json:"googleProfilePhoto,omitempty"`
	GoogleWebsiteOverride  **string  `json:"googleWebsiteOverride,omitempty"`
	GoogleAccountID        *string   `json:"googleAccountId,omitempty"`
	GooglePlaceID          *string   `json:"googlePlaceId,omitempty"`
	GoogleMyBusinessLabels *[]string `json:"googleMyBusinessLabels,omitempty"`
	GoogleShortName        *string   `json:"googleShortName,omitempty"`

	InstagramHandle *string `json:"instagramHandle,omitempty"`
	TwitterHandle   *string `json:"twitterHandle,omitempty"`

	PhotoGallery *[]Photo `json:"photoGallery,omitempty"`
	Videos       *[]Video `json:"videos,omitempty"`
	Headshot     **Image  `json:"headshot,omitempty"`

	GoogleAttributes *map[string][]string `json:"googleAttributes,omitempty"`

	TimeZoneUtcOffset                 *string   `json:"timeZoneUtcOffset,omitempty"`
	Timezone                          *string   `json:"timezone,omitempty"`
	QuestionsAndAnswers               **bool    `json:"questionsAndAnswers,omitempty"`
	DeliverListingsWithoutGeocode     **bool    `json:"deliverListingsWithoutGeocode,omitempty"`
	HolidayHoursConversationEnabled   **bool    `json:"holidayHoursConversationEnabled,omitempty"`
	ReviewResponseConversationEnabled **bool    `json:"reviewResponseConversationEnabled,omitempty"`
	NudgeEnabled                      **bool    `json:"nudgeEnabled,omitempty"`
	What3WordsAddress                 *string   `json:"what3WordsAddress,omitempty"`
	AppointmentOnly                   **bool    `json:"appointmentOnly,omitempty"`
	YearsOfExperience                 **int     `json:"yearsOfExperience,omitempty"`
	Interests                         *[]string `json:"interests,omitempty"`
	Hobbies                           *[]string `json:"hobbies,omitempty"`
	Awards                            *[]string `json:"awards,omitempty"`
}

func (l *FinancialProfessional) UnmarshalJSON(data []byte) error {
	type Alias FinancialProfessional
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

func (y FinancialProfessional) GetId() string {
	if y.BaseEntity.Meta != nil && y.BaseEntity.Meta.Id != nil {
		return *y.BaseEntity.Meta.Id
	}
	return ""
}

func (y FinancialProfessional) GetCategoryIds() (v []string) {
	if y.CategoryIds != nil {
		v = *y.CategoryIds
	}
	return v
}

func (y FinancialProfessional) GetName() string {
	if y.Name != nil {
		return GetString(y.Name)
	}
	return ""
}

func (y FinancialProfessional) GetAccountId() string {
	if y.BaseEntity.Meta != nil && y.BaseEntity.Meta.AccountId != nil {
		return *y.BaseEntity.Meta.AccountId
	}
	return ""
}

func (y FinancialProfessional) GetLine1() string {
	if y.Address != nil && y.Address.Line1 != nil {
		return GetString(y.Address.Line1)
	}
	return ""
}

func (y FinancialProfessional) GetLine2() string {
	if y.Address != nil && y.Address.Line2 != nil {
		return GetString(y.Address.Line2)
	}
	return ""
}

func (y FinancialProfessional) GetAddressHidden() bool {
	return GetNullableBool(y.AddressHidden)
}

func (y FinancialProfessional) GetExtraDescription() string {
	if y.Address != nil && y.Address.ExtraDescription != nil {
		return GetString(y.Address.ExtraDescription)
	}
	return ""
}

func (y FinancialProfessional) GetCity() string {
	if y.Address != nil && y.Address.City != nil {
		return GetString(y.Address.City)
	}
	return ""
}

func (y FinancialProfessional) GetRegion() string {
	if y.Address != nil && y.Address.Region != nil {
		return GetString(y.Address.Region)
	}
	return ""
}

func (y FinancialProfessional) GetCountryCode() string {
	if y.BaseEntity.Meta != nil && y.BaseEntity.Meta.CountryCode != nil {
		return GetString(y.BaseEntity.Meta.CountryCode)
	}
	return ""
}

func (y FinancialProfessional) GetPostalCode() string {
	if y.Address != nil && y.Address.PostalCode != nil {
		return GetString(y.Address.PostalCode)
	}
	return ""
}

func (y FinancialProfessional) GetMainPhone() string {
	if y.MainPhone != nil {
		return *y.MainPhone
	}
	return ""
}

func (y FinancialProfessional) GetLocalPhone() string {
	if y.LocalPhone != nil {
		return *y.LocalPhone
	}
	return ""
}

func (y FinancialProfessional) GetAlternatePhone() string {
	if y.AlternatePhone != nil {
		return *y.AlternatePhone
	}
	return ""
}

func (y FinancialProfessional) GetFax() string {
	if y.Fax != nil {
		return *y.Fax
	}
	return ""
}

func (y FinancialProfessional) GetMobilePhone() string {
	if y.MobilePhone != nil {
		return *y.MobilePhone
	}
	return ""
}

func (y FinancialProfessional) GetTollFreePhone() string {
	if y.TollFreePhone != nil {
		return *y.TollFreePhone
	}
	return ""
}

func (y FinancialProfessional) GetTtyPhone() string {
	if y.TtyPhone != nil {
		return *y.TtyPhone
	}
	return ""
}

func (y FinancialProfessional) GetFeaturedMessage() string {
	f := GetFeaturedMessage(y.FeaturedMessage)
	if f != nil {
		return GetString(f.Description)
	}
	return ""
}

func (y FinancialProfessional) GetFeaturedMessageUrl() string {
	f := GetFeaturedMessage(y.FeaturedMessage)
	if f != nil {
		return GetString(f.Url)
	}
	return ""
}

func (y FinancialProfessional) GetWebsiteUrl() string {
	w := GetWebsite(y.WebsiteUrl)
	if w != nil {
		return GetString(w.Url)
	}
	return ""
}

func (y FinancialProfessional) GetDisplayWebsiteUrl() string {
	w := GetWebsite(y.WebsiteUrl)
	if w != nil {
		return GetString(w.DisplayUrl)
	}
	return ""
}

func (y FinancialProfessional) GetReservationUrl() string {
	w := GetWebsite(y.ReservationUrl)
	if w != nil {
		return GetString(w.Url)
	}
	return ""
}

func (y FinancialProfessional) GetHours() *Hours {
	return GetHours(y.Hours)
}

func (y FinancialProfessional) GetAdditionalHoursText() string {
	if y.AdditionalHoursText != nil {
		return *y.AdditionalHoursText
	}
	return ""
}

func (y FinancialProfessional) GetDescription() string {
	if y.Description != nil {
		return *y.Description
	}
	return ""
}

func (y FinancialProfessional) GetTwitterHandle() string {
	if y.TwitterHandle != nil {
		return *y.TwitterHandle
	}
	return ""
}

func (y FinancialProfessional) GetFacebookPageUrl() string {
	if y.FacebookPageUrl != nil {
		return *y.FacebookPageUrl
	}
	return ""
}

func (y FinancialProfessional) String() string {
	b, _ := json.Marshal(y)
	return string(b)
}

func (y FinancialProfessional) GetKeywords() (v []string) {
	if y.Keywords != nil {
		v = *y.Keywords
	}
	return v
}

func (y FinancialProfessional) GetLanguage() (v string) {
	if y.BaseEntity.Meta.Language != nil {
		v = *y.BaseEntity.Meta.Language
	}
	return v
}

func (y FinancialProfessional) GetAssociations() (v []string) {
	if y.Associations != nil {
		v = *y.Associations
	}
	return v
}

func (y FinancialProfessional) GetEmails() (v []string) {
	if y.Emails != nil {
		v = *y.Emails
	}
	return v
}

func (y FinancialProfessional) GetSpecialties() (v []string) {
	if y.Specialties != nil {
		v = *y.Specialties
	}
	return v
}

func (y FinancialProfessional) GetProductLists() (v *Lists) {
	return GetLists(y.ProductLists)
}

func (y FinancialProfessional) GetServices() (v []string) {
	if y.Services != nil {
		v = *y.Services
	}
	return v
}

func (y FinancialProfessional) GetBrands() (v []string) {
	if y.Brands != nil {
		v = *y.Brands
	}
	return v
}

func (y FinancialProfessional) GetLanguages() (v []string) {
	if y.Languages != nil {
		v = *y.Languages
	}
	return v
}

func (y FinancialProfessional) GetPaymentOptions() (v []string) {
	if y.PaymentOptions != nil {
		v = *y.PaymentOptions
	}
	return v
}

func (y FinancialProfessional) GetVideos() (v []Video) {
	if y.Videos != nil {
		v = *y.Videos
	}
	return v
}

func (y FinancialProfessional) GetGoogleAttributes() map[string][]string {
	if y.GoogleAttributes != nil {
		return *y.GoogleAttributes
	}
	return nil
}

func (y FinancialProfessional) GetHolidayHours() []HolidayHours {
	h := GetHours(y.Hours)
	if h != nil && h.HolidayHours != nil {
		return *h.HolidayHours
	}
	return nil
}

func (y FinancialProfessional) IsClosed() bool {
	return GetNullableBool(y.Closed)
}
