package yext

// TODO
// * Need better custom field accessors and helpers
// * The API will accept some things and return them in a different format - this makes diff'ing difficult:
// ** Phone: Send in 540-444-4444, get back 5404444444

import (
	"encoding/json"
)

const (
	ENTITYTYPE_LOCATION          EntityType = "location"
	FACEBOOKCTA_TYPE_BOOKNOW                = "BOOK_NOW"
	FACEBOOKCTA_TYPE_CALLNOW                = "CALL_NOW"
	FACEBOOKCTA_TYPE_CONTACTUS              = "CONTACT_US"
	FACEBOOKCTA_TYPE_LEARNMORE              = "LEARN_MORE"
	FACEBOOKCTA_TYPE_OFF                    = "NONE"
	FACEBOOKCTA_TYPE_PLAYGAME               = "PLAY_GAME"
	FACEBOOKCTA_TYPE_SENDEMAIL              = "SEND_EMAIL"
	FACEBOOKCTA_TYPE_SENDMESSAGE            = "SEND_MESSAGE"
	FACEBOOKCTA_TYPE_SHOPNOW                = "SHOP_NOW"
	FACEBOOKCTA_TYPE_SIGNUP                 = "SIGN_UP"
	FACEBOOKCTA_TYPE_USEAPP                 = "USE_APP"
	FACEBOOKCTA_TYPE_WATCHVIDEO             = "WATCH_VIDEO"
)

// Location is the representation of a Location in Yext Location Manager.
// For details see https://www.yext.com/support/platform-api/#Administration_API/Locations.htm
type LocationEntity struct {
	BaseEntity
	// Address Fields
	Name               *string             `json:"name,omitempty"`
	Address            *Address            `json:"address,omitempty"`
	AddressHidden      **bool              `json:"addressHidden,omitempty"`
	ISORegionCode      *string             `json:"isoRegionCode,omitempty"`
	ServiceAreaPlaces  *[]ServiceAreaPlace `json:"serviceAreaPlaces,omitempty"`
	Geomodifier        *string             `json:"geomodifier,omitempty"`
	BlackOwnedBusiness **bool              `json:"blackOwnedBusiness,omitempty"`
	Impressum          *string             `json:"impressum,omitempty"`

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
	Hours                     **Hours           `json:"hours,omitempty"`
	AccessHours               **Hours           `json:"accessHours,omitempty"`
	DriveThroughHours         **Hours           `json:"driveThroughHours,omitempty"`
	PickupHours               **Hours           `json:"pickupHours,omitempty"`
	DeliveryHours             **Hours           `json:"deliveryHours,omitempty"`
	Closed                    **bool            `json:"closed,omitempty"`
	Description               *string           `json:"description,omitempty"`
	AdditionalHoursText       *string           `json:"additionalHoursText,omitempty"`
	YearEstablished           **float64         `json:"yearEstablished,omitempty"`
	Associations              *[]string         `json:"associations,omitempty"`
	Certifications            *[]string         `json:"certifications,omitempty"`
	Brands                    *[]string         `json:"brands,omitempty"`
	Products                  *[]string         `json:"products,omitempty"`
	Services                  *[]string         `json:"services,omitempty"`
	PickupAndDeliveryServices *UnorderedStrings `json:"pickupAndDeliveryServices,omitempty"`
	// Spelling of json tag 'specialities' is intentional to match mispelling in Yext API
	Specialties *[]string `json:"specialities,omitempty"`

	FrequentlyAskedQuestions *[]FAQField `json:"frequentlyAskedQuestions,omitempty"`

	Languages      *[]string `json:"languages,omitempty"`
	Logo           **Photo   `json:"logo,omitempty"`
	PaymentOptions *[]string `json:"paymentOptions,omitempty"`

	// Admin
	Keywords    *[]string `json:"keywords,omitempty"`
	CategoryIds *[]string `json:"categoryIds,omitempty"`

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
	MenuUrl         **Website         `json:"menuUrl,omitempty"`
	OrderUrl        **Website         `json:"orderUrl,omitempty"`
	ReservationUrl  **Website         `json:"reservationUrl,omitempty"`
	WebsiteUrl      **Website         `json:"websiteUrl,omitempty"`
	FeaturedMessage **FeaturedMessage `json:"featuredMessage,omitempty"`

	// Uber
	UberLink         **UberLink         `json:"uberLink,omitempty"`
	UberTripBranding **UberTripBranding `json:"uberTripBranding,omitempty"`

	// Social Media
	FacebookCallToAction **FacebookCTA `json:"facebookCallToAction,omitempty"`
	FacebookCoverPhoto   **Image       `json:"facebookCoverPhoto,omitempty"`
	FacebookPageUrl      *string       `json:"facebookPageUrl,omitempty"`
	FacebookProfilePhoto **Image       `json:"facebookProfilePhoto,omitempty"`

	GoogleCoverPhoto       **Image   `json:"googleCoverPhoto,omitempty"`
	GoogleMyBusinessLabels *[]string `json:"googleMyBusinessLabels,omitempty"`
	GooglePreferredPhoto   *string   `json:"googlePreferredPhoto,omitempty"`
	GoogleProfilePhoto     **Image   `json:"googleProfilePhoto,omitempty"`
	GoogleWebsiteOverride  **string  `json:"googleWebsiteOverride,omitempty"`

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

func (l *LocationEntity) UnmarshalJSON(data []byte) error {
	type Alias LocationEntity
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

type FacebookCTA struct {
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

func NullableFacebookCTA(f *FacebookCTA) **FacebookCTA {
	return &f
}

func NullFacebookCTA() **FacebookCTA {
	var f *FacebookCTA
	return &f
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

func NullableUberLink(u *UberLink) **UberLink {
	return &u
}

func GetUberLink(u **UberLink) *UberLink {
	if u == nil {
		return nil
	}
	return *u
}

func NullUberLink() **UberLink {
	var u *UberLink
	return &u
}

type UberTripBranding struct {
	Text        *string `json:"text,omitempty"`
	Url         *string `json:"url,omitempty"`
	Description *string `json:"description,omitempty"`
}

func NullableUberTripBranding(u *UberTripBranding) **UberTripBranding {
	return &u
}

func GetUberTripBranding(u **UberTripBranding) *UberTripBranding {
	if u == nil {
		return nil
	}
	return *u
}

func NullUberTripBranding() **UberTripBranding {
	var u *UberTripBranding
	return &u
}

type Lists struct {
	Label *string   `json:"label,omitempty"`
	Ids   *[]string `json:"ids,omitempty"`
}

func NullableLists(l *Lists) **Lists {
	return &l
}

func GetLists(l **Lists) *Lists {
	if l == nil {
		return nil
	}
	return *l
}

func NullLists() **Lists {
	var l *Lists
	return &l
}

type Photo struct {
	Image           *Image  `json:"image,omitempty"`
	ClickthroughUrl *string `json:"clickthroughUrl,omitempty"`
	Description     *string `json:"description,omitempty"`
	Details         *string `json:"details,omitempty"`
}

type Image struct {
	Url           *string      `json:"url,omitempty"`
	SourceUrl     *string      `json:"sourceUrl,omitempty"`
	AlternateText *string      `json:"alternateText,omitempty"`
	Width         *int         `json:"width,omitempty"`
	Height        *int         `json:"height,omitempty"`
	Thumbnails    *[]Thumbnail `json:"thumnails,omitempty"`
}

type Thumbnail struct {
	Url    *string `json:"url,omitempty"`
	Width  *int    `json:"width,omitempty"`
	Height *int    `json:"height,omitempty"`
}

func NullableImage(i *Image) **Image {
	return &i
}

func GetImage(i **Image) *Image {
	if i == nil {
		return nil
	}
	return *i
}

func NullImage() **Image {
	var i *Image
	return &i
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

func NullableFeaturedMessage(f *FeaturedMessage) **FeaturedMessage {
	return &f
}

func GetFeaturedMessage(f **FeaturedMessage) *FeaturedMessage {
	if f == nil {
		return nil
	}
	return *f
}

func NullFeaturedMessage() **FeaturedMessage {
	var f *FeaturedMessage
	return &f
}

type Website struct {
	DisplayUrl       *string `json:"displayUrl,omitempty"`
	Url              *string `json:"url,omitempty"`
	PreferDisplayUrl **bool  `json:"preferDisplayUrl,omitempty"`
}

func (y Photo) String() string {
	b, _ := json.Marshal(y)
	return string(b)
}

func NullableWebsite(w *Website) **Website {
	return &w
}

func GetWebsite(w **Website) *Website {
	if w == nil {
		return nil
	}
	return *w
}

func NullWebsite() **Website {
	var w *Website
	return &w
}

type Coordinate struct {
	Latitude  **float64 `json:"latitude,omitempty"`
	Longitude **float64 `json:"longitude,omitempty"`
}

func NullableCoordinate(c *Coordinate) **Coordinate {
	return &c
}

func GetCoordinate(c **Coordinate) *Coordinate {
	if c == nil {
		return nil
	}
	return *c
}

func NullCoordinate() **Coordinate {
	var c *Coordinate
	return &c
}

type Hours struct {
	Monday       **DayHours      `json:"monday,omitempty"`
	Tuesday      **DayHours      `json:"tuesday,omitempty"`
	Wednesday    **DayHours      `json:"wednesday,omitempty"`
	Thursday     **DayHours      `json:"thursday,omitempty"`
	Friday       **DayHours      `json:"friday,omitempty"`
	Saturday     **DayHours      `json:"saturday,omitempty"`
	Sunday       **DayHours      `json:"sunday,omitempty"`
	HolidayHours *[]HolidayHours `json:"holidayHours,omitempty"`
	ReopenDate   *string         `json:"reopenDate,omitempty"`
}

func (h Hours) String() string {
	b, _ := json.Marshal(h)
	return string(b)
}

func NewHoursClosedAllWeek() *Hours {
	h := &Hours{}
	h.SetClosedAllWeek()
	return h
}

func NewHoursUnspecifiedAllWeek() *Hours {
	h := &Hours{}
	h.SetUnspecifiedAllWeek()
	return h
}

type DayHours struct {
	OpenIntervals *[]Interval `json:"openIntervals,omitempty"`
	IsClosed      **bool      `json:"isClosed,omitempty"`
}

func NullableDayHours(d *DayHours) **DayHours {
	return &d
}

func NullDayHours() **DayHours {
	var v *DayHours
	return &v
}

func GetDayHours(d **DayHours) *DayHours {
	if d == nil {
		return nil
	}
	return *d
}

func (d *DayHours) SetClosed() {
	d.IsClosed = NullableBool(true)
	d.OpenIntervals = nil
}

func (d *DayHours) GetIntervals() []Interval {
	if d != nil && d.OpenIntervals != nil {
		return *d.OpenIntervals
	}
	return nil
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
	if h == nil {
		return nil
	}
	switch w {
	case Sunday:
		return GetDayHours(h.Sunday)
	case Monday:
		return GetDayHours(h.Monday)
	case Tuesday:
		return GetDayHours(h.Tuesday)
	case Wednesday:
		return GetDayHours(h.Wednesday)
	case Thursday:
		return GetDayHours(h.Thursday)
	case Friday:
		return GetDayHours(h.Friday)
	case Saturday:
		return GetDayHours(h.Saturday)
	}
	return nil
}

func (h *Hours) SetClosedAllWeek() {
	h.SetClosed(Sunday)
	h.SetClosed(Monday)
	h.SetClosed(Tuesday)
	h.SetClosed(Wednesday)
	h.SetClosed(Thursday)
	h.SetClosed(Friday)
	h.SetClosed(Saturday)
}

func (h *Hours) SetClosed(w Weekday) {
	d := &DayHours{}
	d.SetClosed()
	switch w {
	case Sunday:
		h.Sunday = NullableDayHours(d)
	case Monday:
		h.Monday = NullableDayHours(d)
	case Tuesday:
		h.Tuesday = NullableDayHours(d)
	case Wednesday:
		h.Wednesday = NullableDayHours(d)
	case Thursday:
		h.Thursday = NullableDayHours(d)
	case Friday:
		h.Friday = NullableDayHours(d)
	case Saturday:
		h.Saturday = NullableDayHours(d)
	}
}

func (h *Hours) SetUnspecifiedAllWeek() {
	h.SetUnspecified(Sunday)
	h.SetUnspecified(Monday)
	h.SetUnspecified(Tuesday)
	h.SetUnspecified(Wednesday)
	h.SetUnspecified(Thursday)
	h.SetUnspecified(Friday)
	h.SetUnspecified(Saturday)
}

func (h *Hours) SetUnspecified(w Weekday) {
	switch w {
	case Sunday:
		h.Sunday = NullableDayHours(&DayHours{
			OpenIntervals: nil,
			IsClosed:      nil,
		})
	case Monday:
		h.Monday = NullableDayHours(&DayHours{
			OpenIntervals: nil,
			IsClosed:      nil,
		})
	case Tuesday:
		h.Tuesday = NullableDayHours(&DayHours{
			OpenIntervals: nil,
			IsClosed:      nil,
		})
	case Wednesday:
		h.Wednesday = NullableDayHours(&DayHours{
			OpenIntervals: nil,
			IsClosed:      nil,
		})
	case Thursday:
		h.Thursday = NullableDayHours(&DayHours{
			OpenIntervals: nil,
			IsClosed:      nil,
		})
	case Friday:
		h.Friday = NullableDayHours(&DayHours{
			OpenIntervals: nil,
			IsClosed:      nil,
		})
	case Saturday:
		h.Saturday = NullableDayHours(&DayHours{
			OpenIntervals: nil,
			IsClosed:      nil,
		})
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
		h.Sunday = NullableDayHours(d)
	case Monday:
		h.Monday = NullableDayHours(d)
	case Tuesday:
		h.Tuesday = NullableDayHours(d)
	case Wednesday:
		h.Wednesday = NullableDayHours(d)
	case Thursday:
		h.Thursday = NullableDayHours(d)
	case Friday:
		h.Friday = NullableDayHours(d)
	case Saturday:
		h.Saturday = NullableDayHours(d)
	}
}

func (h *Hours) SetHours(w Weekday, start string, end string) {
	d := &DayHours{}
	d.AddHours(start, end)
	switch w {
	case Sunday:
		h.Sunday = NullableDayHours(d)
	case Monday:
		h.Monday = NullableDayHours(d)
	case Tuesday:
		h.Tuesday = NullableDayHours(d)
	case Wednesday:
		h.Wednesday = NullableDayHours(d)
	case Thursday:
		h.Thursday = NullableDayHours(d)
	case Friday:
		h.Friday = NullableDayHours(d)
	case Saturday:
		h.Saturday = NullableDayHours(d)
	}
}

func (y LocationEntity) GetId() string {
	if y.BaseEntity.Meta != nil && y.BaseEntity.Meta.Id != nil {
		return *y.BaseEntity.Meta.Id
	}
	return ""
}

func (y LocationEntity) GetCategoryIds() (v []string) {
	if y.CategoryIds != nil {
		v = *y.CategoryIds
	}
	return v
}

func (y LocationEntity) GetName() string {
	if y.Name != nil {
		return GetString(y.Name)
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
		return GetString(y.Address.Line1)
	}
	return ""
}

func (y LocationEntity) GetLine2() string {
	if y.Address != nil && y.Address.Line2 != nil {
		return GetString(y.Address.Line2)
	}
	return ""
}

func (y LocationEntity) GetAddressHidden() bool {
	return GetNullableBool(y.AddressHidden)
}

func (y LocationEntity) GetExtraDescription() string {
	if y.Address != nil && y.Address.ExtraDescription != nil {
		return GetString(y.Address.ExtraDescription)
	}
	return ""
}

func (y LocationEntity) GetCity() string {
	if y.Address != nil && y.Address.City != nil {
		return GetString(y.Address.City)
	}
	return ""
}

func (y LocationEntity) GetRegion() string {
	if y.Address != nil && y.Address.Region != nil {
		return GetString(y.Address.Region)
	}
	return ""
}

func (y LocationEntity) GetCountryCode() string {
	if y.BaseEntity.Meta != nil && y.BaseEntity.Meta.CountryCode != nil {
		return GetString(y.BaseEntity.Meta.CountryCode)
	}
	return ""
}

func (y LocationEntity) GetPostalCode() string {
	if y.Address != nil && y.Address.PostalCode != nil {
		return GetString(y.Address.PostalCode)
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

func (y LocationEntity) GetFeaturedMessage() string {
	f := GetFeaturedMessage(y.FeaturedMessage)
	if f != nil {
		return GetString(f.Description)
	}
	return ""
}

func (y LocationEntity) GetFeaturedMessageUrl() string {
	f := GetFeaturedMessage(y.FeaturedMessage)
	if f != nil {
		return GetString(f.Url)
	}
	return ""
}

func (y LocationEntity) GetWebsiteUrl() string {
	w := GetWebsite(y.WebsiteUrl)
	if w != nil {
		return GetString(w.Url)
	}
	return ""
}

func (y LocationEntity) GetDisplayWebsiteUrl() string {
	w := GetWebsite(y.WebsiteUrl)
	if w != nil {
		return GetString(w.DisplayUrl)
	}
	return ""
}

func (y LocationEntity) GetReservationUrl() string {
	w := GetWebsite(y.ReservationUrl)
	if w != nil {
		return GetString(w.Url)
	}
	return ""
}

func (y LocationEntity) GetHours() *Hours {
	return GetHours(y.Hours)
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
		return GetNullableFloat(y.YearEstablished)
	}
	return 0
}

func (y LocationEntity) GetDisplayLat() float64 {
	c := GetCoordinate(y.DisplayCoordinate)
	if c != nil {
		return GetNullableFloat(c.Latitude)
	}
	return 0
}

func (y LocationEntity) GetDisplayLng() float64 {
	c := GetCoordinate(y.DisplayCoordinate)
	if c != nil {
		return GetNullableFloat(c.Longitude)
	}
	return 0
}

func (y LocationEntity) GetRoutableLat() float64 {
	c := GetCoordinate(y.RoutableCoordinate)
	if c != nil {
		return GetNullableFloat(c.Latitude)
	}
	return 0
}

func (y LocationEntity) GetRoutableLng() float64 {
	c := GetCoordinate(y.RoutableCoordinate)
	if c != nil {
		return GetNullableFloat(c.Longitude)
	}
	return 0
}

func (y LocationEntity) GetYextDisplayLat() float64 {
	c := GetCoordinate(y.YextDisplayCoordinate)
	if c != nil {
		return GetNullableFloat(c.Latitude)
	}
	return 0
}

func (y LocationEntity) GetYextDisplayLng() float64 {
	c := GetCoordinate(y.YextDisplayCoordinate)
	if c != nil {
		return GetNullableFloat(c.Longitude)
	}
	return 0
}

func (y LocationEntity) GetYextRoutableLat() float64 {
	c := GetCoordinate(y.YextRoutableCoordinate)
	if c != nil {
		return GetNullableFloat(c.Latitude)
	}
	return 0
}

func (y LocationEntity) GetYextRoutableLng() float64 {
	c := GetCoordinate(y.YextRoutableCoordinate)
	if c != nil {
		return GetNullableFloat(c.Longitude)
	}
	return 0
}

func (y LocationEntity) GetBios() (v *Lists) {
	return GetLists(y.Bios)
}

func (y LocationEntity) GetCalendars() (v *Lists) {
	return GetLists(y.Calendars)
}

func (y LocationEntity) GetProductLists() (v *Lists) {
	return GetLists(y.ProductLists)
}

func (y LocationEntity) GetMenus() (v *Lists) {
	return GetLists(y.Menus)
}

func (y LocationEntity) GetReviewGenerationUrl() string {
	if y.ReviewGenerationUrl != nil {
		return *y.ReviewGenerationUrl
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
	h := GetHours(y.Hours)
	if h != nil && h.HolidayHours != nil {
		return *h.HolidayHours
	}
	return nil
}

func (y LocationEntity) IsClosed() bool {
	return GetNullableBool(y.Closed)
}

// HolidayHours represents individual exceptions to a Location's regular hours in Yext Location Manager.
// For details see
type HolidayHours struct {
	Date           *string     `json:"date"`
	OpenIntervals  *[]Interval `json:"openIntervals,omitempty"`
	IsClosed       **bool      `json:"isClosed,omitempty"`
	IsRegularHours **bool      `json:"isRegularHours,omitempty"`
}

func (y HolidayHours) String() string {
	b, _ := json.Marshal(y)
	return string(b)
}

func ToHolidayHours(y []HolidayHours) *[]HolidayHours {
	return &y
}

type ServiceAreaPlace struct {
	Name *string `json:"name"`
	Type *string `json:"type"`
}

type CTA struct {
	Text     *string `json:"label,omitempty"`
	URL      *string `json:"link,omitempty"`
	LinkType *string `json:"linkType,omitempty"`
}

func (y CTA) String() string {
	b, _ := json.Marshal(y)
	return string(b)
}

func ToCTA(c CTA) *CTA {
	return &c
}

func NullableCTA(c CTA) **CTA {
	p := &c
	return &p
}

type FAQField struct {
	Answer   *string `json:"answer,omitempty"`
	Question *string `json:"question,omitempty"`
}

func (y FAQField) String() string {
	b, _ := json.Marshal(y)
	return string(b)
}

func ToFAQ(c FAQField) *FAQField {
	return &c
}

func NullableFAQ(c FAQField) **FAQField {
	p := &c
	return &p
}
