package yext

// TODO
// * Need better custom field accessors and helpers
// * The API will accept some things and return them in a different format - this makes diff'ing difficult:
// ** Phone: Send in 540-444-4444, get back 5404444444

import (
	"encoding/json"

	yext "github.com/yext/yext-go"
)

const ENTITYTYPE_LOCATION EntityType = "location"

// Location is the representation of a Location in Yext Location Manager.
// For details see https://www.yext.com/support/platform-api/#Administration_API/Locations.htm
type LocationEntity struct {
	BaseEntity

	// Admin
	CategoryIds *[]string `json:"categoryIds,omitempty"`
	Closed      *bool     `json:"closed,omitempty"`
	Keywords    *[]string `json:"keywords,omitempty"`

	// Address Fields
	Name          *string  `json:"name,omitempty"`
	Address       *Address `json:"address,omitempty"`
	AddressHidden *bool    `json:"addressHidden,omitempty"`

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
	Hours               *Hours    `json:"hours,omitempty"`
	AdditionalHoursText *string   `json:"additionalHoursText,omitempty"`
	YearEstablished     *float64  `json:"yearEstablished,omitempty"`
	Associations        *[]string `json:"associations,omitempty"`
	Certifications      *[]string `json:"certifications,omitempty"`
	Brands              *[]string `json:"brands,omitempty"`
	Products            *[]string `json:"products,omitempty"`
	Services            *[]string `json:"services,omitempty"`
	Specialties         *[]string `json:"specialties,omitempty"`
	Languages           *[]string `json:"languages,omitempty"`
	Logo                *Image    `json:"logo,omitempty"`
	PaymentOptions      *[]string `json:"paymentOptions,omitempty"`

	// Lats & Lngs
	DisplayCoordinate  *Coordinate `json:"displayCoordinate,omitempty"`
	DropoffCoordinate  *Coordinate `json:"dropoffCoordinate,omitempty"`
	WalkableCoordinate *Coordinate `json:"walkableCoordinate,omitempty"`
	RoutableCoordinate *Coordinate `json:"routableCoordinate,omitempty"`
	PickupCoordinate   *Coordinate `json:"pickupCoordinate,omitempty"`

	// Lists
	Bios         *Lists `json:"bios,omitempty"`
	Calendars    *Lists `json:"calendars,omitempty"`
	Menus        *Lists `json:"menus,omitempty"`
	ProductLists *Lists `json:"productLists,omitempty"`

	// Urls
	MenuUrl         *Website         `json:"menuUrl,omitempty"`
	OrderUrl        *Website         `json:"orderUrl,omitempty"`
	ReservationUrl  *Website         `json:"reservationUrl,omitempty"`
	WebsiteUrl      *Website         `json:"websiteUrl,omitempty"`
	FeaturedMessage *FeaturedMessage `json:"featuredMessage,omitempty"`

	// Uber
	//UberClientId         *string `json:"uberClientId,omitempty"`
	UberLink         *UberLink         `json:"uberLink,omitempty"`
	UberTripBranding *UberTripBranding `json:"uberTripBranding,omitempty"`

	// Social Media
	FacebookCoverPhoto   *LocationPhoto `json:"facebookCoverPhoto,omitempty"`
	FacebookPageUrl      *string        `json:"facebookPageUrl,omitempty"`
	FacebookProfilePhoto *LocationPhoto `json:"facebookProfilePhoto,omitempty"`

	GoogleCoverPhoto      *LocationPhoto `json:"googleCoverPhoto,omitempty"`
	GooglePreferredPhoto  *string        `json:"googlePreferredPhoto,omitempty"`
	GoogleProfilePhoto    *LocationPhoto `json:"googleProfilePhoto,omitempty"`
	GoogleWebsiteOverride *string        `json:"googleWebsiteOverride,omitempty"`

	InstagramHandle *string `json:"instagramHandle,omitempty"`
	TwitterHandle   *string `json:"twitterHandle,omitempty"`

	Photos    *[]LocationPhoto `json:"photos,omitempty"`
	VideoUrls *[]string        `json:"videoUrls,omitempty"`

	GoogleAttributes *map[string][]string `json:"googleAttributes,omitempty"`

	// Reviews
	//ReviewBalancingURL   *string `json:"reviewBalancingURL,omitempty"`
	FirstPartyReviewPage *string `json:"firstPartyReviewPage,omitempty"`
}

type Video struct {
	VideoUrl    VideoUrl `json:"video,omitempty"`
	Description string   `json:"description,omitempty"`
}

type VideoUrl struct {
	Url string `json:"url,omitempty"`
}

type UberLink struct {
	Text         *string `json:"text,omitempty"`
	Presentation *string `json:"presentation,omitempty"`
}

type UberTripBranding struct {
	Text        *string `json:"text,omitempty"`
	Url         *string `json:"url,omitempty"`
	Description *string `json:"description,omitempty"`
}

type Lists struct {
	Label *string   `json:"label,omitempty"`
	Ids   *[]string `json:"ids,omitempty"`
}

type Photo struct {
	Image           *Image  `json:"image,omitempty"`
	ClickthroughUrl *string `json:"clickthroughUrl,omitempty"`
	Description     *string `json:"description,omitempty"`
	Details         *string `json:"details,omitempty"`
}

type Image struct {
	Url           *string `json:"url,omitempty"`
	AlternateText *string `json:"alternateText,omitempty"`
}

type Address struct {
	Line1            *string `json:"line1,omitempty"`
	Line2            *string `json:"line2,omitempty"`
	City             *string `json:"city,omitempty"`
	Region           *string `json:"region,omitempty"`
	Sublocality      *string `json:"sublocality,omitempty"`
	PostalCode       *string `json:"postalCode,omitempty"`
	CountryCode      *string `json:"countryCode,omitempty"`
	ExtraDescription *string `json:"extraDescription,omitempty"`
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
	OpenIntervals *[]Interval `json:"openIntervals,omitempty"`
	IsClosed      *bool       `json:"isClosed,omitempty"`
}

func (d *DayHours) SetClosed() {
	d.IsClosed = yext.Bool(true)
	d.OpenIntervals = nil
}

func (d *DayHours) AddHours(start string, end string) {
	intervals := []Interval{}
	d.IsClosed = nil
	if d.OpenIntervals != nil {
		intervals = *d.OpenIntervals
	}
	intervals = append(intervals, Interval{
		Start: start,
		End:   end,
	})
	d.OpenIntervals = &intervals
}

func (d *DayHours) SetHours(start string, end string) {
	d.IsClosed = nil
	d.OpenIntervals = &[]Interval{
		Interval{
			Start: start,
			End:   end,
		},
	}
}

type Interval struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

func NewInterval(start string, end string) *Interval {
	return &Interval{Start: start, End: end}
}

func (h *Hours) GetDayHours(w Weekday) *DayHours {
	switch w {
	case Sunday:
		return h.Sunday
	case Monday:
		return h.Monday
	case Tuesday:
		return h.Tuesday
	case Wednesday:
		return h.Wednesday
	case Thursday:
		return h.Thursday
	case Friday:
		return h.Friday
	case Saturday:
		return h.Saturday
	}
	return nil
}

func (h *Hours) SetClosedAllWeek() {
	h = &Hours{
		Sunday:    &DayHours{},
		Monday:    &DayHours{},
		Tuesday:   &DayHours{},
		Wednesday: &DayHours{},
		Thursday:  &DayHours{},
		Friday:    &DayHours{},
		Saturday:  &DayHours{},
	}
	h.Sunday.SetClosed()
	h.Monday.SetClosed()
	h.Tuesday.SetClosed()
	h.Wednesday.SetClosed()
	h.Thursday.SetClosed()
	h.Friday.SetClosed()
	h.Saturday.SetClosed()
}

func (h *Hours) SetClosed(w Weekday) {
	d := &DayHours{}
	d.SetClosed()
	switch w {
	case Sunday:
		h.Sunday = d
	case Monday:
		h.Monday = d
	case Tuesday:
		h.Tuesday = d
	case Wednesday:
		h.Wednesday = d
	case Thursday:
		h.Thursday = d
	case Friday:
		h.Friday = d
	case Saturday:
		h.Saturday = d
	}
}

func (h *Hours) SetUnspecified(w Weekday) {
	switch w {
	case Sunday:
		h.Sunday = nil
	case Monday:
		h.Monday = nil
	case Tuesday:
		h.Tuesday = nil
	case Wednesday:
		h.Wednesday = nil
	case Thursday:
		h.Thursday = nil
	case Friday:
		h.Friday = nil
	case Saturday:
		h.Saturday = nil
	}
}

func (h *Hours) AddHours(w Weekday, start string, end string) {
	d := h.GetDayHours(w)
	if d == nil {
		d = &DayHours{}
	}
	d.AddHours(start, end)
	switch w {
	case Sunday:
		h.Sunday = d
	case Monday:
		h.Monday = d
	case Tuesday:
		h.Tuesday = d
	case Wednesday:
		h.Wednesday = d
	case Thursday:
		h.Thursday = d
	case Friday:
		h.Friday = d
	case Saturday:
		h.Saturday = d
	}
}

func (h *Hours) SetHours(w Weekday, start string, end string) {
	d := &DayHours{}
	d.AddHours(start, end)
	switch w {
	case Sunday:
		h.Sunday = d
	case Monday:
		h.Monday = d
	case Tuesday:
		h.Tuesday = d
	case Wednesday:
		h.Wednesday = d
	case Thursday:
		h.Thursday = d
	case Friday:
		h.Friday = d
	case Saturday:
		h.Saturday = d
	}
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

func (y LocationEntity) GetAddressHidden() bool {
	if y.AddressHidden != nil {
		return *y.AddressHidden
	}
	return false
}

func (y LocationEntity) GetExtraDescription() string {
	if y.Address != nil && y.Address.ExtraDescription != nil {
		return *y.Address.ExtraDescription
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

func (y LocationEntity) GetMainPhone() string {
	if y.MainPhone != nil {
		return *y.MainPhone
	}
	return ""
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

func (y LocationEntity) GetFax() string {
	if y.Fax != nil {
		return *y.Fax
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

func (y LocationEntity) GetBios() (v Lists) {
	if y.Bios != nil {
		v = *y.Bios
	}
	return v
}

func (y LocationEntity) GetCalendars() (v Lists) {
	if y.Calendars != nil {
		v = *y.Calendars
	}
	return v
}

func (y LocationEntity) GetProductLists() (v Lists) {
	if y.ProductLists != nil {
		v = *y.ProductLists
	}
	return v
}

func (y LocationEntity) GetMenus() (v Lists) {
	if y.Menus != nil {
		v = *y.Menus
	}
	return v
}

// func (y LocationEntity) GetReviewBalancingURL() string {
// 	if y.ReviewBalancingURL != nil {
// 		return *y.ReviewBalancingURL
// 	}
// 	return ""
// }

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
	if y.BaseEntity.Meta.Language != nil {
		v = *y.BaseEntity.Meta.Language
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

func (y LocationEntity) GetPaymentOptions() (v []string) {
	if y.PaymentOptions != nil {
		v = *y.PaymentOptions
	}
	return v
}

func (y LocationEntity) GetVideos() (v []Video) {
	if y.Videos != nil {
		v = *y.Videos
	}
	return v
}

func (y LocationEntity) GetGoogleAttributes() map[string][]string {
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
	if y.Closed != nil {
		return *y.Closed
	}
	return false
}

// HolidayHours represents individual exceptions to a Location's regular hours in Yext Location Manager.
// For details see
type HolidayHours struct {
	Date           string      `json:"date"`
	OpenIntervals  *[]Interval `json:"openIntervals,omitempty"`
	IsClosed       *bool       `json:"isClosed,omitempty"`
	IsRegularHours *bool       `json:"isRegularHours,omitempty"`
}
