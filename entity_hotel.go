package yext

import (
	"encoding/json"
)

const ENTITYTYPE_HOTEL EntityType = "hotel"

type HotelEntity struct {
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
	AdditionalHoursText *string   `json:"additionalHoursText,omitempty"`
	Associations        *[]string `json:"associations,omitempty"`
	Brands              *[]string `json:"brands,omitempty"`
	Description         *string   `json:"description,omitempty"`
	Hours               **Hours   `json:"hours,omitempty"`
	CheckInTime         *string   `json:"checkInTime,omitempty"` // TODO: check type of time field
	CheckOutTime        *string   `json:"checkOutTime,omitempty"`
	Services            *[]string `json:"services,omitempty"`
	Languages           *[]string `json:"languages,omitempty"`

	// Social Media
	FacebookPageUrl *string `json:"facebookPageUrl,omitempty"`
	InstagramHandle *string `json:"instagramHandle,omitempty"`
	TwitterHandle   *string `json:"twitterHandle,omitempty"`

	FacebookCoverPhoto   **Image `json:"facebookCoverPhoto,omitempty"`
	FacebookProfilePhoto **Image `json:"facebookProfilePhoto,omitempty"`

	GoogleCoverPhoto      **Image              `json:"googleCoverPhoto,omitempty"`
	GooglePreferredPhoto  *string              `json:"googlePreferredPhoto,omitempty"`
	GoogleProfilePhoto    **Image              `json:"googleProfilePhoto,omitempty"`
	GoogleWebsiteOverride **string             `json:"googleWebsiteOverride,omitempty"`
	GoogleAttributes      *map[string][]string `json:"googleAttributes,omitempty"`

	// Media
	Logo         **Photo  `json:"logo,omitempty"`
	PhotoGallery *[]Photo `json:"photoGallery,omitempty"`
	Videos       *[]Video `json:"videos,omitempty"`

	// Lists
	Bios         **Lists `json:"bios,omitempty"`
	Menus        **Lists `json:"menus,omitempty"`
	ProductLists **Lists `json:"productLists,omitempty"`

	// URLs
	FeaturedMessage **FeaturedMessage `json:"featuredMessage,omitempty"`
	MenuUrl         **Website         `json:"menuUrl,omitempty"`
	OrderUrl        **Website         `json:"orderUrl,omitempty"`
	ReservationUrl  **Website         `json:"reservationUrl,omitempty"`
	WebsiteUrl      **Website         `json:"websiteUrl,omitempty"`

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

	// Property Info
	YearEstablished    **float64 `json:"yearEstablished,omitempty"`
	YearLastRenovated  **int     `json:"yearLastRenovated,omitempty"`
	RoomCount          **int     `json:"roomCount,omitempty"`
	FloorCount         **int     `json:"floorCount,omitempty"`
	BeachFrontProperty **Ternary `json:"beachFrontProperty,omitempty"`

	// Services
	ClassificationRating    *string   `json:"classificationRating,omitempty"`
	FrontDesk               *string   `json:"frontDesk,omitempty"`
	Laundry                 *string   `json:"laundry,omitempty"`
	Housekeeping            *string   `json:"housekeeping,omitempty"`
	Parking                 *string   `json:"parking,omitempty"`
	SelfParking             *string   `json:"selfParking,omitempty"`
	ValetParking            *string   `json:"valetParking,omitempty"`
	Concierge               **Ternary `json:"concierge,omitempty"`
	Elevator                **Ternary `json:"elevator,omitempty"`
	BaggageStorage          **Ternary `json:"baggageStorage,omitempty"`
	SocialHour              **Ternary `json:"socialHour,omitempty"`
	WakeUpCalls             **Ternary `json:"wakeUpCalls,omitempty"`
	ConvenienceStore        **Ternary `json:"convenienceStore,omitempty"`
	GiftShop                **Ternary `json:"giftShop,omitempty"`
	CurrencyExchange        **Ternary `json:"currencyExchange,omitempty"`
	TurndownService         **Ternary `json:"turndownService,omitempty"`
	ElectricChargingStation **Ternary `json:"electricChargingStation,omitempty"`

	// Policies
	PaymentOptions          *[]string `json:"paymentOptions,omitempty"`
	AllInclusive            *string   `json:"allInclusive,omitempty"`
	PetsAllowed             *string   `json:"petsAllowed,omitempty"`
	KidsStayFree            **Ternary `json:"kidsStayFree,omitempty"`
	MaxAgeOfKidsStayFree    **int     `json:"maxAgeOfKidsStayFree,omitempty"`
	MaxNumberOfKidsStayFree **int     `json:"maxNumberOfKidsStayFree,omitempty"`
	SmokeFreeProperty       **Ternary `json:"smokeFreeProperty,omitempty"`
	CatsAllowed             **Ternary `json:"catsAllowed,omitempty"`
	DogsAllowed             **Ternary `json:"dogsAllowed,omitempty"`

	// Food and Drink
	RoomService     *string           `json:"roomService,omitempty"`
	RestaurantCount **int             `json:"restaurantCount,omitempty"`
	Breakfast       *string           `json:"breakfast,omitempty"`
	TableService    **Ternary         `json:"tableService,omitempty"`
	Bar             **Ternary         `json:"bar,omitempty"`
	VendingMachine  **Ternary         `json:"vendingMachine,omitempty"`
	BuffetOptions   *UnorderedStrings `json:"buffetOptions,omitempty"`

	// Pools
	IndoorPoolCount  **int     `json:"indoorPoolCount,omitempty"`
	IndoorPools      **Ternary `json:"indoorPools,omitempty"`
	OutdoorPoolCount **int     `json:"outdoorPoolCount,omitempty"`
	OutdoorPools     **Ternary `json:"outdoorPools,omitempty"`
	Pools            **Ternary `json:"pools,omitempty"`
	HotTub           **Ternary `json:"hotTub,omitempty"`
	WaterSlide       **Ternary `json:"waterslide,omitempty"`
	LazyRiver        **Ternary `json:"lazyRiver,omitempty"`
	AdultPool        **Ternary `json:"adultPool,omitempty"`
	WadingPool       **Ternary `json:"wadingPool,omitempty"`
	WavePool         **Ternary `json:"wavePool,omitempty"`
	ThermalPool      **Ternary `json:"thermalPool,omitempty"`
	WaterPark        **Ternary `json:"waterPark,omitempty"`
	LifeGuard        **Ternary `json:"lifeguard,omitempty"`

	// Wellness
	FitnessCenter     *string   `json:"fitnessCenter,omitempty"`
	EllipticalMachine **Ternary `json:"ellipticalMachine,omitempty"`
	Treadmill         **Ternary `json:"treadmill,omitempty"`
	WeightMachine     **Ternary `json:"weightMachine,omitempty"`
	FreeWeights       **Ternary `json:"freeWeights,omitempty"`
	Spa               **Ternary `json:"spa,omitempty"`
	Salon             **Ternary `json:"salon,omitempty"`
	Sauna             **Ternary `json:"sauna,omitempty"`
	Massage           **Ternary `json:"massage,omitempty"`
	DoctorOnCall      **Ternary `json:"doctorOnCall,omitempty"`

	// Activities
	Bicycles        *string   `json:"bicycles,omitempty"`
	WaterCraft      *string   `json:"watercraft,omitempty"`
	GameRoom        **Ternary `json:"gameRoom,omitempty"`
	Nightclub       **Ternary `json:"nightclub,omitempty"`
	Casino          **Ternary `json:"casino,omitempty"`
	BoutiqueStores  **Ternary `json:"boutiqueStores,omitempty"`
	Tennis          **Ternary `json:"tennis,omitempty"`
	Golf            **Ternary `json:"golf,omitempty"`
	HorsebackRiding **Ternary `json:"horsebackRiding,omitempty"`
	Snorkeling      **Ternary `json:"snorkeling,omitempty"`
	Scuba           **Ternary `json:"scuba,omitempty"`
	WaterSkiing     **Ternary `json:"waterSkiing,omitempty"`
	BeachAccess     **Ternary `json:"beachAccess,omitempty"`
	PrivateBeach    **Ternary `json:"privateBeach,omitempty"`

	// Transportation
	AirportShuttle    *string   `json:"airportShuttle,omitempty"`
	PrivateCarService *string   `json:"privateCarService,omitempty"`
	AirportTransfer   **Ternary `json:"airportTransfer,omitempty"`
	LocalShuttle      **Ternary `json:"localShuttle,omitempty"`
	CarRental         **Ternary `json:"carRental,omitempty"`

	// Families
	BabySittingOffered **Ternary `json:"babysittingOffered,omitempty"`
	KidFriendly        **Ternary `json:"kidFriendly,omitempty"`
	KidsClub           **Ternary `json:"kidsClub,omitempty"`

	// Connectivity
	WiFiAvailable *string           `json:"wifiAvailable,omitempty"`
	WiFiDetails   *UnorderedStrings `json:"wifiDetails,omitempty"`

	// Business
	MeetingRoomCount **int     `json:"meetingRoomCount,omitempty"`
	BusinessCenter   **Ternary `json:"businessCenter,omitempty"`

	// Accessibility
	MobilityAccessible   **Ternary         `json:"mobilityAccessible,omitempty"`
	AccessibilityDetails *UnorderedStrings `json:"accessibilityDetails,omitempty"`

	// Covid & Cleanliness Fields
	DigitalGuestRoomKeys                  **Ternary `json:"digitalGuestRoomKeys,omitempty"`
	CommonAreasAdvancedCleaning           **Ternary `json:"commonAreasEnhancedCleaning,omitempty"`
	GuestRoomsEnhancedCleaning            **Ternary `json:"guestRoomsEnhancedCleaning,omitempty"`
	CommercialGradeDisinfectantUsed       **Ternary `json:"commercialGradeDisinfectantUsed,omitempty"`
	EmployeesTrainedInCleaningProcedures  **Ternary `json:"employeesTrainedInCleaningProcedures,omitempty"`
	EmployeesTrainedInHandWashing         **Ternary `json:"employeesTrainedInHandWashing,omitempty"`
	EmployeesWearProtectiveEquipment      **Ternary `json:"employeesWearProtectiveEquipment,omitempty"`
	HighTouchItemsRemovedFromGuestRooms   **Ternary `json:"highTouchItemsRemovedFromGuestRooms,omitempty"`
	HighTouchItemsRemovedFromCommonAreas  **Ternary `json:"highTouchItemsRemovedFromCommonAreas,omitempty"`
	PlasticKeycardsDisinfectedOrDiscarded **Ternary `json:"plasticKeycardsDisinfectedOrDiscarded,omitempty"`
	ContactlessCheckInCheckOut            **Ternary `json:"contactlessCheckinOrCheckout,omitempty"`
	PhysicalDistancingRequired            **Ternary `json:"physicalDistancingRequired,omitempty"`
	PlexiglassUsed                        **Ternary `json:"plexiglassUsed,omitempty"`
	LimitedOccupancyInSharedAreas         **Ternary `json:"limitedOccupancyInSharedAreas,omitempty"`
	PrivateSpacesInWellnessAreas          **Ternary `json:"privateSpacesInWellnessAreas,omitempty"`
	CommonAreasArrangedForDistancing      **Ternary `json:"commonAreasArrangedForDistancing,omitempty"`
	SanitizerAvailable                    **Ternary `json:"sanitizerAvailable,omitempty"`
	MasksRequired                         **Ternary `json:"masksRequired,omitempty"`
	IndividuallyPackagedMealsAvailable    **Ternary `json:"individuallyPackagedMealsAvailable,omitempty"`
	DisposableFlatware                    **Ternary `json:"disposableFlatware,omitempty"`
	SingleUseMenus                        **Ternary `json:"singleUseMenus,omitempty"`
	RoomBookingsBuffer                    **Ternary `json:"roomBookingsBuffer,omitempty"`
	RequestOnlyHousekeeping               **Ternary `json:"requestOnlyHousekeeping,omitempty"`
	InRoomHygieneKits                     **Ternary `json:"inRoomHygieneKits,omitempty"`
	ProtectiveEquipmentAvailable          **Ternary `json:"protectiveEquipmentAvailable,omitempty"`
	SafeHandlingForFoodServices           **Ternary `json:"safeHandlingForFoodServices,omitempty"`
	AdditionalSanitationInFoodAreas       **Ternary `json:"additionalSanitationInFoodAreas,omitempty"`
	EcoCertifications                     *UnorderedStrings `json:"ecoCertifications,omitempty"`
	EcoFriendlyToiletries                 **Ternary `json:"ecoFriendlyToiletries,omitempty"`
}

const (
	// Single-option IDs
	OptionNotApplicable                  = "NOT_APPLICABLE"
	OptionPetsWelcomeForFree             = "PETS_WELCOME_FOR_FREE"
	OptionPetsWelcome                    = "PETS_WELCOME"
	OptionFrontDeskAvailable24Hours      = "FRONT_DESK_AVAILABLE_24_HOURS"
	OptionFrontDeskAvailable             = "FRONT_DESK_AVAILABLE"
	OptionHousekeepingAvailableDaily     = "HOUSEKEEPING_AVAILABLE_DAILY"
	OptionHousekeepingAvailable          = "HOUSEKEEPING_AVAILABLE"
	OptionFullServiceLaundry             = "FULL_SERVICE"
	OptionSelfServiceLaundry             = "SELF_SERVICE"
	Option24HourRoomService              = "ROOM_SERVICE_AVAILABLE_24_HOURS"
	OptionRoomServiceAvailable           = "ROOM_SERVICE_AVAILABLE"
	OptionBicycleRentals                 = "BICYCLE_RENTALS"
	OptionBicycleRentalsForFree          = "BICYCLE_RENTALS_FOR_FREE"
	OptionFitnessCenterAvailable         = "FITNESS_CENTER_AVAILABLE"
	OptionFitnessCenterAvailableForFree  = "FITNESS_CENTER_AVAILABLE_FOR_FREE"
	OptionParkingAvailableForFree        = "PARKING_AVAILABLE_FOR_FREE"
	OptionParkingAvailable               = "PARKING_AVAILABLE"
	OptionSelfParkingAvailableForFree    = "SELF_PARKING_AVAILABLE_FOR_FREE"
	OptionSelfParkingAvailable           = "SELF_PARKING_AVAILABLE"
	OptionValetParkingAvailableForFree   = "VALET_PARKING_AVAILABLE_FOR_FREE"
	OptionValetParkingAvailable          = "VALET_PARKING_AVAILABLE"
	OptionAirportShuttleAvailableForFree = "AIRPORT_SHUTTLE_AVAILABLE_FOR_FREE"
	OptionAirportShuttleAvailable        = "AIRPORT_SHUTTLE_AVAILABLE"
	OptionBreakfastAvailable             = "BREAKFAST_AVAILABLE"
	OptionBreakfastAvailableForFree      = "BREAKFAST_AVAILABLE_FOR_FREE"
	OptionWiFiAvailable                  = "WIFI_AVAILABLE"
	OptionWiFiAvailableForFree           = "WIFI_AVAILABLE_FOR_FREE"
	OptionBuffetBreakfast                = "BUFFET_BREAKFAST"
	OptionBuffetDinner                   = "BUFFET_DINNER"
	OptionBuffet                         = "BUFFET"
	OptionPrivateCarService              = "PRIVATE_CAR_SERVICE"
	OptionPrivateCarServiceForFree       = "PRIVATE_CAR_SERVICE_FOR_FREE"
	OptionWatercraftRentals              = "WATERCRAFT_RENTALS"
	OptionWatercraftRentalsForFree       = "WATERCRAFT_RENTALS_FOR_FREE"
	OptionAllInclusiveRatesAvailable     = "ALL_INCLUSIVE_RATES_AVAILABLE"
	OptionAllInclusiveRatesOnly          = "ALL_INCLUSIVE_RATES_ONLY"

	// Multi-option IDs
	OptionWiFiInPublicAreas      = "WIFI_IN_PUBLIC_AREAS"
	OptionPublicInternetTerminal = "PUBLIC_INTERNET_TERMINAL"
	OptionAccessibleParking      = "ACCESSIBLE_PARKING"
	OptionAccessibleElevator     = "ACCESSIBLE_ELEVATOR"
	OptionAccessiblePool         = "ACCESSIBLE_POOL"
)

func (h *HotelEntity) UnmarshalJSON(data []byte) error {
	type Alias HotelEntity
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

func (y HotelEntity) GetId() string {
	if y.BaseEntity.Meta != nil && y.BaseEntity.Meta.Id != nil {
		return *y.BaseEntity.Meta.Id
	}
	return ""
}

func (y HotelEntity) GetCategoryIds() (v []string) {
	if y.CategoryIds != nil {
		v = *y.CategoryIds
	}
	return v
}

func (y HotelEntity) GetName() string {
	if y.Name != nil {
		return GetString(y.Name)
	}
	return ""
}

func (y HotelEntity) GetAccountId() string {
	if y.BaseEntity.Meta != nil && y.BaseEntity.Meta.AccountId != nil {
		return *y.BaseEntity.Meta.AccountId
	}
	return ""
}

func (y HotelEntity) GetLine1() string {
	if y.Address != nil && y.Address.Line1 != nil {
		return GetString(y.Address.Line1)
	}
	return ""
}

func (y HotelEntity) GetLine2() string {
	if y.Address != nil && y.Address.Line2 != nil {
		return GetString(y.Address.Line2)
	}
	return ""
}

func (y HotelEntity) GetAddressHidden() bool {
	return GetNullableBool(y.AddressHidden)
}

func (y HotelEntity) GetExtraDescription() string {
	if y.Address != nil && y.Address.ExtraDescription != nil {
		return GetString(y.Address.ExtraDescription)
	}
	return ""
}

func (y HotelEntity) GetCity() string {
	if y.Address != nil && y.Address.City != nil {
		return GetString(y.Address.City)
	}
	return ""
}

func (y HotelEntity) GetRegion() string {
	if y.Address != nil && y.Address.Region != nil {
		return GetString(y.Address.Region)
	}
	return ""
}

func (y HotelEntity) GetCountryCode() string {
	if y.BaseEntity.Meta != nil && y.BaseEntity.Meta.CountryCode != nil {
		return GetString(y.BaseEntity.Meta.CountryCode)
	}
	if y.Address != nil && y.Address.CountryCode != nil {
		return GetString(y.Address.CountryCode)
	}
	return ""
}

func (y HotelEntity) GetPostalCode() string {
	if y.Address != nil && y.Address.PostalCode != nil {
		return GetString(y.Address.PostalCode)
	}
	return ""
}

func (y HotelEntity) GetMainPhone() string {
	return GetString(y.MainPhone)
}

func (y HotelEntity) GetFeaturedMessage() string {
	f := GetFeaturedMessage(y.FeaturedMessage)
	if f != nil {
		return GetString(f.Description)
	}
	return ""
}

func (y HotelEntity) GetFeaturedMessageUrl() string {
	f := GetFeaturedMessage(y.FeaturedMessage)
	if f != nil {
		return GetString(f.Url)
	}
	return ""
}

func (y HotelEntity) GetWebsiteUrl() string {
	w := GetWebsite(y.WebsiteUrl)
	if w != nil {
		return GetString(w.Url)
	}
	return ""
}

func (y HotelEntity) GetDisplayWebsiteUrl() string {
	w := GetWebsite(y.WebsiteUrl)
	if w != nil {
		return GetString(w.DisplayUrl)
	}
	return ""
}

func (y HotelEntity) GetReservationUrl() string {
	w := GetWebsite(y.ReservationUrl)
	if w != nil {
		return GetString(w.Url)
	}
	return ""
}

func (y HotelEntity) GetHours() *Hours {
	return GetHours(y.Hours)
}

func (y HotelEntity) GetYearEstablished() float64 {
	if y.YearEstablished != nil {
		return GetNullableFloat(y.YearEstablished)
	}
	return 0
}

func (y HotelEntity) GetDisplayLat() float64 {
	c := GetCoordinate(y.DisplayCoordinate)
	if c != nil {
		return GetNullableFloat(c.Latitude)
	}
	return 0
}

func (y HotelEntity) GetDisplayLng() float64 {
	c := GetCoordinate(y.DisplayCoordinate)
	if c != nil {
		return GetNullableFloat(c.Longitude)
	}
	return 0
}

func (y HotelEntity) GetRoutableLat() float64 {
	c := GetCoordinate(y.RoutableCoordinate)
	if c != nil {
		return GetNullableFloat(c.Latitude)
	}
	return 0
}

func (y HotelEntity) GetRoutableLng() float64 {
	c := GetCoordinate(y.RoutableCoordinate)
	if c != nil {
		return GetNullableFloat(c.Longitude)
	}
	return 0
}

func (y HotelEntity) GetYextDisplayLat() float64 {
	c := GetCoordinate(y.YextDisplayCoordinate)
	if c != nil {
		return GetNullableFloat(c.Latitude)
	}
	return 0
}

func (y HotelEntity) GetYextDisplayLng() float64 {
	c := GetCoordinate(y.YextDisplayCoordinate)
	if c != nil {
		return GetNullableFloat(c.Longitude)
	}
	return 0
}

func (y HotelEntity) GetYextRoutableLat() float64 {
	c := GetCoordinate(y.YextRoutableCoordinate)
	if c != nil {
		return GetNullableFloat(c.Latitude)
	}
	return 0
}

func (y HotelEntity) GetYextRoutableLng() float64 {
	c := GetCoordinate(y.YextRoutableCoordinate)
	if c != nil {
		return GetNullableFloat(c.Longitude)
	}
	return 0
}

func (y HotelEntity) GetBios() (v *Lists) {
	return GetLists(y.Bios)
}

func (y HotelEntity) GetProductLists() (v *Lists) {
	return GetLists(y.ProductLists)
}

func (y HotelEntity) GetMenus() (v *Lists) {
	return GetLists(y.Menus)
}

func (y HotelEntity) GetKeywords() (v []string) {
	if y.Keywords != nil {
		v = *y.Keywords
	}
	return v
}

func (y HotelEntity) GetLanguage() (v string) {
	if y.BaseEntity.Meta.Language != nil {
		v = *y.BaseEntity.Meta.Language
	}
	return v
}

func (y HotelEntity) GetAssociations() (v []string) {
	if y.Associations != nil {
		v = *y.Associations
	}
	return v
}

func (y HotelEntity) GetEmails() (v []string) {
	if y.Emails != nil {
		v = *y.Emails
	}
	return v
}

func (y HotelEntity) GetServices() (v []string) {
	if y.Services != nil {
		v = *y.Services
	}
	return v
}

func (y HotelEntity) GetBrands() (v []string) {
	if y.Brands != nil {
		v = *y.Brands
	}
	return v
}

func (y HotelEntity) GetLanguages() (v []string) {
	if y.Languages != nil {
		v = *y.Languages
	}
	return v
}

func (y HotelEntity) GetPaymentOptions() (v []string) {
	if y.PaymentOptions != nil {
		v = *y.PaymentOptions
	}
	return v
}

func (y HotelEntity) GetVideos() (v []Video) {
	if y.Videos != nil {
		v = *y.Videos
	}
	return v
}

func (y HotelEntity) GetHolidayHours() []HolidayHours {
	h := GetHours(y.Hours)
	if h != nil && h.HolidayHours != nil {
		return *h.HolidayHours
	}
	return nil
}

func (y HotelEntity) GetBuffetOptions() []string {
	if y.BuffetOptions != nil {
		v := *y.BuffetOptions
		return []string(v)
	}
	return []string{}
}

func (y HotelEntity) GetWiFiDetails() []string {
	if y.WiFiDetails != nil {
		v := *y.WiFiDetails
		return []string(v)
	}
	return []string{}
}

func (y HotelEntity) GetAccessibilityDetails() []string {
	if y.AccessibilityDetails != nil {
		v := *y.AccessibilityDetails
		return []string(v)
	}
	return []string{}
}

func (y HotelEntity) IsClosed() bool {
	return GetNullableBool(y.Closed)
}

func (y HotelEntity) String() string {
	b, _ := json.Marshal(y)
	return string(b)
}
