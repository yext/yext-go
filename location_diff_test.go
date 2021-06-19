package yext

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

var examplePhoto = LocationPhoto{
	Url:         "http://www.google.com",
	Description: "An example image",
}

var complexOne = &Location{
	Id:   String("lock206"),
	Name: String("Farmers Insurance - Stephen Lockhart "),
	CustomFields: map[string]interface{}{
		"1857": "",
		"1858": "47291E",
		"1859": "Stephen",
		"1871": "Lockhart",
		"3004": "Agent",
		"7240": "Stephen Paul Lockhart",
		"7251": true,
		"7253": "4729",
		"7254": "slockhart",
		"7255": "2684",
		"7256": []string{"lock206"},
		"7261": false,
		"7263": true,
		"7265": "",
		"7266": "",
		"7269": "1E",
		"7270": true,
		"7271": "29",
		"7272": "User",
		"7273": "Tonawanda",
		"7274": "NY",
		"7275": "2690 Sheridan Dr W",
		"7276": "14150-9425",
		"7277": "47",
		"7278": false,
		"7279": true,
		"7283": "",
		"7284": "Paul",
		"7285": "1225598",
		"7286": "",
		"7287": false,
		"7288": true,
		"7296": "1225598",
		"7297": "",
		"7298": "",
		"7299": "1046846",
		"7300": true,
	},
	Address:         String("2690 Sheridan Dr W"),
	Address2:        String(""),
	City:            String("Tonawanda"),
	State:           String("NY"),
	Zip:             String("14150-9425"),
	Phone:           String("716-835-0306"),
	FaxPhone:        String("716-835-0415"),
	YearEstablished: String("2015"),
	Emails:          &[]string{"slockhart@farmersagent.com"},
	Services: &[]string{
		"Auto Insurance",
		"Home Insurance",
		"Homeowners Insurance",
		"Business Insurance",
		"Motorcyle Insurance",
		"Recreational Insurance",
		"Renters Insurance",
		"Umbrella Insurance",
	},
	Languages: &[]string{"English"},
	FolderId:  String("91760"),
}

var complexTwo = &Location{
	Id:   String("lock206"),
	Name: String("Farmers Insurance - Stephen Lockhart "),
	CustomFields: map[string]interface{}{
		"1857": "",
		"1858": "47291E",
		"1859": "Stephen",
		"1871": "Lockhart",
		"3004": "Agent",
		"7240": "Stephen Paul Lockhart",
		"7251": true,
		"7253": "4729",
		"7254": "slockhart",
		"7255": "2684",
		"7256": []string{"lock206"},
		"7261": false,
		"7263": true,
		"7265": "",
		"7266": "",
		"7269": "1E",
		"7270": true,
		"7271": "29",
		"7272": "User",
		"7273": "Tonawanda",
		"7274": "NY",
		"7275": "2690 Sheridan Dr W",
		"7276": "14150-9425",
		"7277": "47",
		"7278": false,
		"7279": true,
		"7283": "",
		"7284": "Paul",
		"7285": "1225598",
		"7286": "",
		"7287": false,
		"7288": true,
		"7296": "1225598",
		"7297": "",
		"7298": "",
		"7299": "1046846",
		"7300": true,
	},
	Address:         String("2690 Sheridan Dr W"),
	Address2:        String(""),
	City:            String("Tonawanda"),
	State:           String("NY"),
	Zip:             String("14150-9425"),
	Phone:           String("716-835-0306"),
	FaxPhone:        String("716-835-0415"),
	YearEstablished: String("2015"),
	Emails:          &[]string{"slockhart@farmersagent.com"},
	Services: &[]string{
		"Auto Insurance",
		"Home Insurance",
		"Homeowners Insurance",
		"Business Insurance",
		"Motorcyle Insurance",
		"Recreational Insurance",
		"Renters Insurance",
		"Umbrella Insurance",
	},
	Languages: &[]string{"English"},
	FolderId:  String("91760"),
}

var jsonData string = `{"id":"phai514","locationName":"Farmers Insurance - Aroun Phaisan ","customFields":{"1857":"","1858":"122191","1859":"Aroun","1871":"Phaisan","3004":"Agent","7240":"Aroun Phaisan","7251":true,"7253":"1221","7254":"aphaisan","7255":"2685","7256":["phai514"],"7261":false,"7263":true,"7265":"","7266":"","7269":"91","7270":true,"7271":"21","7272":"User_Dup","7273":"Lincoln","7274":"NE","7275":"5730 R St Ste B","7276":"68505-2309","7277":"12","7278":false,"7279":true,"7283":"","7284":"","7285":"16133384","7286":"","7287":true,"7288":true,"7296":"16133384","7297":"","7298":"","7299":"786200","7300":true},"address":"5730 R St","address2":"Ste B","city":"Lincoln","state":"NE","zip":"68505-2309","phone":"402-417-4266","faxPhone":"402-423-3141","yearEstablished":"2011","emails":["aphaisan@farmersagent.com"],"services":["Auto Insurance","Home Insurance","Homeowners Insurance","Business Insurance","Motorcyle Insurance","Recreational Insurance","Renters Insurance","Umbrella Insurance","Term Life Insurance","Whole Life Insurance"],"languages":["English"],"folderId":"91760"}`

var baseLocation Location = Location{
	Id:                     String("ding"),
	Name:                   String("ding"),
	AccountId:              String("ding"),
	Address:                String("ding"),
	Address2:               String("ding"),
	DisplayAddress:         String("ding"),
	City:                   String("ding"),
	State:                  String("ding"),
	Zip:                    String("ding"),
	CountryCode:            String("ding"),
	Phone:                  String("ding"),
	LocalPhone:             String("ding"),
	AlternatePhone:         String("ding"),
	FaxPhone:               String("ding"),
	MobilePhone:            String("ding"),
	TollFreePhone:          String("ding"),
	TtyPhone:               String("ding"),
	FeaturedMessage:        String("ding"),
	FeaturedMessageUrl:     String("ding"),
	WebsiteUrl:             String("ding"),
	DisplayWebsiteUrl:      String("ding"),
	ReservationUrl:         String("ding"),
	Hours:                  String("ding"),
	AdditionalHoursText:    String("ding"),
	Description:            String("ding"),
	TwitterHandle:          String("ding"),
	FacebookPageUrl:        String("ding"),
	YearEstablished:        String("ding"),
	FolderId:               String("ding"),
	SuppressAddress:        Bool(false),
	IsPhoneTracked:         Bool(true),
	DisplayLat:             Float(1234.0),
	DisplayLng:             Float(1234.0),
	RoutableLat:            Float(1234.0),
	RoutableLng:            Float(1234.0),
	Keywords:               &[]string{"ding", "ding"},
	PaymentOptions:         &[]string{"ding", "ding"},
	VideoUrls:              &[]string{"ding", "ding"},
	Emails:                 &[]string{"ding", "ding"},
	Specialties:            &[]string{"ding", "ding"},
	Services:               &[]string{"ding", "ding"},
	Brands:                 &[]string{"ding", "ding"},
	Languages:              &[]string{"ding", "ding"},
	Logo:                   &examplePhoto,
	FacebookCoverPhoto:     &examplePhoto,
	FacebookProfilePicture: &examplePhoto,
	Photos:                 &[]LocationPhoto{examplePhoto, examplePhoto, examplePhoto},
	ProductListIds:         &[]string{"1234", "5678"},
	Closed: &LocationClosed{
		IsClosed: Bool(false),
	},
	CustomFields: map[string]interface{}{
		"1234": "ding",
	},
}

func TestDiffIdentical(t *testing.T) {
	secondLocation := &Location{
		Id:                     String("ding"),
		Name:                   String("ding"),
		AccountId:              String("ding"),
		Address:                String("ding"),
		Address2:               String("ding"),
		DisplayAddress:         String("ding"),
		City:                   String("ding"),
		State:                  String("ding"),
		Zip:                    String("ding"),
		CountryCode:            String("ding"),
		Phone:                  String("ding"),
		LocalPhone:             String("ding"),
		AlternatePhone:         String("ding"),
		FaxPhone:               String("ding"),
		MobilePhone:            String("ding"),
		TollFreePhone:          String("ding"),
		TtyPhone:               String("ding"),
		FeaturedMessage:        String("ding"),
		FeaturedMessageUrl:     String("ding"),
		WebsiteUrl:             String("ding"),
		DisplayWebsiteUrl:      String("ding"),
		ReservationUrl:         String("ding"),
		Hours:                  String("ding"),
		AdditionalHoursText:    String("ding"),
		Description:            String("ding"),
		TwitterHandle:          String("ding"),
		FacebookPageUrl:        String("ding"),
		YearEstablished:        String("ding"),
		FolderId:               String("ding"),
		SuppressAddress:        Bool(false),
		IsPhoneTracked:         Bool(true),
		DisplayLat:             Float(1234.0),
		DisplayLng:             Float(1234.0),
		RoutableLat:            Float(1234.0),
		RoutableLng:            Float(1234.0),
		Keywords:               &[]string{"ding", "ding"},
		PaymentOptions:         &[]string{"ding", "ding"},
		VideoUrls:              &[]string{"ding", "ding"},
		Emails:                 &[]string{"ding", "ding"},
		Specialties:            &[]string{"ding", "ding"},
		Services:               &[]string{"ding", "ding"},
		Brands:                 &[]string{"ding", "ding"},
		Languages:              &[]string{"ding", "ding"},
		Logo:                   &examplePhoto,
		FacebookCoverPhoto:     &examplePhoto,
		FacebookProfilePicture: &examplePhoto,
		Photos:                 &[]LocationPhoto{examplePhoto, examplePhoto, examplePhoto},
		ProductListIds:         &[]string{"1234", "5678"},
		Closed: &LocationClosed{
			IsClosed: Bool(false),
		},
		CustomFields: map[string]interface{}{
			"1234": "ding",
		},
	}
	d, isDiff := baseLocation.Diff(secondLocation)
	if isDiff == true {
		t.Errorf("Expected diff to be false was true, diff result %v", d)
	} else if d != nil {
		t.Errorf("Expected an empty diff location, but got %v", d)
	}
}

type stringTest struct {
	baseValue          *string
	newValue           *string
	isDiff             bool
	nilIsEmpty         bool
	expectedFieldValue *string
}

var stringTests = []stringTest{
	stringTest{String("ding"), String("ding"), false, false, nil},
	stringTest{String("ding"), String("dong"), true, false, String("dong")},
	stringTest{nil, String("dong"), true, false, String("dong")},
	stringTest{nil, String(""), true, false, String("")},
	stringTest{nil, String(""), false, true, nil},
	stringTest{String(""), String(""), false, false, nil},
	stringTest{String(""), nil, false, false, nil},
	stringTest{String(""), nil, false, true, nil},
	stringTest{String(""), String("dong"), true, false, String("dong")},
	{nil, nil, false, false, nil},
}

func formatStringPtr(s *string) string {
	if s == nil {
		return "nil"
	} else if *s == "" {
		return "empty string"
	} else {
		return *s
	}
}

func (t stringTest) formatErrorBase(index int) string {
	bv := formatStringPtr(t.baseValue)
	nv := formatStringPtr(t.newValue)
	return fmt.Sprintf("Failure with example %v:\n\tbase: '%v'\n\tnew: '%v'", index, bv, nv)
}

func TestStringDiffs(t *testing.T) {
	a, b := *new(Location), new(Location)

	for i, data := range stringTests {
		a.Name, b.Name = data.baseValue, data.newValue
		a.nilIsEmpty, b.nilIsEmpty = data.nilIsEmpty, data.nilIsEmpty

		d, isDiff := a.Diff(b)
		if isDiff != data.isDiff {
			t.Errorf("%vExpected diff to be %v\nbut was %v\ndiff struct was %v\n", data.formatErrorBase(i), data.isDiff, isDiff, d)
		}
		if d == nil && data.expectedFieldValue == nil {
			continue
		} else if d == nil && data.expectedFieldValue != nil {
			t.Errorf("%v\ndelta was nil but expected %v\n", data.formatErrorBase(i), formatStringPtr(data.expectedFieldValue))
		} else if d != nil && data.expectedFieldValue == nil {
			t.Errorf("%v\ndiff was%v\n", data.formatErrorBase(i), d)
		} else if *d.Name != *data.expectedFieldValue {
			t.Errorf("%v\ndiff was%v\n", data.formatErrorBase(i), d)
		}
	}
}

type boolTest struct {
	baseValue          *bool
	newValue           *bool
	isDiff             bool
	nilIsEmpty         bool
	expectedFieldValue *bool
}

var boolTests = []boolTest{
	{Bool(false), Bool(false), false, false, nil},
	{Bool(true), Bool(true), false, false, nil},
	{Bool(false), Bool(true), true, false, Bool(true)},
	{Bool(true), Bool(false), true, false, Bool(false)},
	{nil, Bool(false), true, false, Bool(false)},
	{nil, Bool(false), false, true, nil},
	{nil, Bool(true), true, false, Bool(true)},
	{Bool(false), nil, false, false, nil},
	{Bool(false), nil, false, true, nil},
	{Bool(true), nil, false, false, nil},
	{nil, nil, false, false, nil},
}

func formatBoolPtr(b *bool) string {
	if b == nil {
		return "nil"
	} else {
		return fmt.Sprintf("%v", *b)
	}
}

func (t boolTest) formatErrorBase(index int) string {
	bv := formatBoolPtr(t.baseValue)
	nv := formatBoolPtr(t.newValue)
	return fmt.Sprintf("Failure with example %v:\n\tbase: '%v'\n\tnew: '%v'\n", index, bv, nv)
}

func TestBoolDiffs(t *testing.T) {
	a, b := *new(Location), new(Location)
	for i, data := range boolTests {
		a.SuppressAddress, b.SuppressAddress = data.baseValue, data.newValue
		a.nilIsEmpty, b.nilIsEmpty = data.nilIsEmpty, data.nilIsEmpty
		d, isDiff := a.Diff(b)
		if isDiff != data.isDiff {
			t.Errorf("%v\nExpected diff to be %v\nbut was %v\ndiff struct was %v\n", data.formatErrorBase(i), data.isDiff, isDiff, d)
		}
		if d == nil && data.expectedFieldValue == nil {
			continue
		} else if d == nil && data.expectedFieldValue != nil {
			t.Errorf("%v\ndelta was nil but expected %v\n", data.formatErrorBase(i), formatBoolPtr(data.expectedFieldValue))
		} else if d != nil && data.expectedFieldValue == nil {
			t.Errorf("%v\ndiff was%v\n", data.formatErrorBase(i), d)
		} else if *d.SuppressAddress != *data.expectedFieldValue {
			t.Errorf("%v\ndiff was%v\n", data.formatErrorBase(i), d)
		}
	}
}

type googleAttributesTest struct {
	baseValue          *LocationGoogleAttributes
	newValue           *LocationGoogleAttributes
	isDiff             bool
	nilIsEmpty         bool
	expectedFieldValue []*LocationGoogleAttribute
}

var googleAttributesTests = []googleAttributesTest{
	{nil, nil, false, false, nil},
	{
		baseValue:          nil,
		newValue:           &LocationGoogleAttributes{},
		isDiff:             true,
		nilIsEmpty:         false,
		expectedFieldValue: LocationGoogleAttributes{},
	},
	{
		baseValue:          nil,
		newValue:           &LocationGoogleAttributes{},
		isDiff:             false,
		nilIsEmpty:         true,
		expectedFieldValue: nil,
	},
	{
		baseValue:          &LocationGoogleAttributes{},
		newValue:           nil,
		isDiff:             false,
		nilIsEmpty:         true,
		expectedFieldValue: nil,
	},
	{
		baseValue:          &LocationGoogleAttributes{},
		newValue:           nil,
		isDiff:             false,
		nilIsEmpty:         false,
		expectedFieldValue: nil,
	},
	{
		baseValue:          &LocationGoogleAttributes{&LocationGoogleAttribute{Id: String("has_delivery"), OptionIds: Strings([]string{"true"})}},
		newValue:           &LocationGoogleAttributes{&LocationGoogleAttribute{Id: String("has_delivery"), OptionIds: Strings([]string{"true"})}},
		isDiff:             false,
		nilIsEmpty:         false,
		expectedFieldValue: nil,
	},
	{
		baseValue:          &LocationGoogleAttributes{&LocationGoogleAttribute{Id: String("has_delivery"), OptionIds: Strings([]string{"true"})}},
		newValue:           &LocationGoogleAttributes{&LocationGoogleAttribute{Id: String("has_delivery"), OptionIds: Strings([]string{"false"})}},
		isDiff:             true,
		nilIsEmpty:         false,
		expectedFieldValue: []*LocationGoogleAttribute{&LocationGoogleAttribute{Id: String("has_delivery"), OptionIds: Strings([]string{"false"})}},
	},
	{
		baseValue: &LocationGoogleAttributes{
			&LocationGoogleAttribute{Id: String("has_delivery"), OptionIds: Strings([]string{"true"})},
			&LocationGoogleAttribute{Id: String("has_catering"), OptionIds: Strings([]string{"true"})},
		},
		newValue: &LocationGoogleAttributes{
			&LocationGoogleAttribute{Id: String("has_delivery"), OptionIds: Strings([]string{"true"})},
			&LocationGoogleAttribute{Id: String("has_catering"), OptionIds: Strings([]string{"true"})},
		},
		isDiff:             false,
		nilIsEmpty:         false,
		expectedFieldValue: nil,
	},
	{
		baseValue: &LocationGoogleAttributes{
			&LocationGoogleAttribute{OptionIds: Strings([]string{"true"}), Id: String("has_delivery")},
			&LocationGoogleAttribute{Id: String("has_catering"), OptionIds: Strings([]string{"true"})},
		},
		newValue: &LocationGoogleAttributes{
			&LocationGoogleAttribute{Id: String("has_catering"), OptionIds: Strings([]string{"true"})},
			&LocationGoogleAttribute{Id: String("has_delivery"), OptionIds: Strings([]string{"true"})},
		},
		isDiff:             false,
		nilIsEmpty:         false,
		expectedFieldValue: nil,
	},
	{
		baseValue: &LocationGoogleAttributes{
			&LocationGoogleAttribute{Id: String("has_delivery"), OptionIds: Strings([]string{"false"})},
			&LocationGoogleAttribute{Id: String("has_catering"), OptionIds: Strings([]string{"false"})},
		},
		newValue: &LocationGoogleAttributes{
			&LocationGoogleAttribute{Id: String("has_catering"), OptionIds: Strings([]string{"true"})},
			&LocationGoogleAttribute{Id: String("has_delivery"), OptionIds: Strings([]string{"true"})},
		},
		isDiff:     true,
		nilIsEmpty: false,
		expectedFieldValue: LocationGoogleAttributes{
			&LocationGoogleAttribute{Id: String("has_catering"), OptionIds: Strings([]string{"true"})},
			&LocationGoogleAttribute{Id: String("has_delivery"), OptionIds: Strings([]string{"true"})},
		},
	},
	{
		baseValue: &LocationGoogleAttributes{
			&LocationGoogleAttribute{Id: String("has_catering"), OptionIds: Strings([]string{"false"})},
		},
		newValue: &LocationGoogleAttributes{
			&LocationGoogleAttribute{Id: String("has_catering"), OptionIds: Strings([]string{"false"})},
			&LocationGoogleAttribute{Id: String("has_delivery"), OptionIds: Strings([]string{"true"})},
		},
		isDiff:     true,
		nilIsEmpty: false,
		expectedFieldValue: LocationGoogleAttributes{
			&LocationGoogleAttribute{Id: String("has_catering"), OptionIds: Strings([]string{"false"})},
			&LocationGoogleAttribute{Id: String("has_delivery"), OptionIds: Strings([]string{"true"})},
		},
	},
	{
		baseValue: &LocationGoogleAttributes{
			&LocationGoogleAttribute{Id: String("has_catering"), OptionIds: Strings([]string{"false"})},
		},
		newValue: &LocationGoogleAttributes{
			&LocationGoogleAttribute{Id: String("has_catering"), OptionIds: nil},
		},
		isDiff:     true,
		nilIsEmpty: false,
		expectedFieldValue: LocationGoogleAttributes{
			&LocationGoogleAttribute{Id: String("has_catering"), OptionIds: Strings([]string{"false"})},
		},
	},
	{
		baseValue: &LocationGoogleAttributes{
			&LocationGoogleAttribute{Id: String("has_catering"), OptionIds: nil},
		},
		newValue: &LocationGoogleAttributes{
			&LocationGoogleAttribute{Id: String("has_catering"), OptionIds: Strings([]string{"false"})},
		},
		isDiff:     true,
		nilIsEmpty: false,
		expectedFieldValue: LocationGoogleAttributes{
			&LocationGoogleAttribute{Id: String("has_catering"), OptionIds: Strings([]string{"false"})},
		},
	},
	{
		baseValue: &LocationGoogleAttributes{
			&LocationGoogleAttribute{Id: String("has_catering"), OptionIds: nil},
		},
		newValue: &LocationGoogleAttributes{
			&LocationGoogleAttribute{Id: String("has_catering"), OptionIds: nil},
		},
		isDiff:             false,
		nilIsEmpty:         false,
		expectedFieldValue: nil,
	},
	{
		baseValue: &LocationGoogleAttributes{
			&LocationGoogleAttribute{Id: String("has_catering"), OptionIds: Strings([]string{"false"})},
		},
		newValue: &LocationGoogleAttributes{
			&LocationGoogleAttribute{Id: String("has_catering"), OptionIds: Strings([]string{})},
		},
		isDiff:     true,
		nilIsEmpty: false,
		expectedFieldValue: LocationGoogleAttributes{
			&LocationGoogleAttribute{Id: String("has_catering"), OptionIds: Strings([]string{"false"})},
		},
	},
}

func (t googleAttributesTest) formatErrorBase(index int) string {
	return fmt.Sprintf("Failure with example %v:\n\tbase: '%v'\n\tnew: '%v'\n", index, t.baseValue, t.newValue)
}

func TestGoogleAttributesDiffs(t *testing.T) {
	a, b := *new(Location), new(Location)
	for i, data := range googleAttributesTests {
		a.GoogleAttributes, b.GoogleAttributes = data.baseValue, data.newValue
		a.nilIsEmpty, b.nilIsEmpty = data.nilIsEmpty, data.nilIsEmpty
		d, isDiff := a.Diff(b)
		if isDiff != data.isDiff {
			t.Errorf("%vExpected diff to be %v\nbut was %v\ndiff struct was %v\n", data.formatErrorBase(i), data.isDiff, isDiff, d)
		}
		if d == nil && data.expectedFieldValue == nil {
			continue
		} else if d == nil && data.expectedFieldValue != nil {
			t.Errorf("%v\ndelta was nil but expected %v\n", data.formatErrorBase(i), data.expectedFieldValue)
		} else if len(d.GetGoogleAttributes()) != len(data.expectedFieldValue) {
			t.Errorf("%v\ndiff was%v\n", data.formatErrorBase(i), d)
		}
	}
}

type stringArrayTest struct {
	baseValue          *[]string
	newValue           *[]string
	isDiff             bool
	nilIsEmpty         bool
	expectedFieldValue []string
}

var stringArrayTests = []stringArrayTest{
	{&[]string{"ding", "dong"}, &[]string{"ding", "dong"}, false, false, nil},
	{&[]string{"ding", "dong"}, &[]string{"ding", "dong", "dang"}, true, false, []string{"ding", "dong", "dang"}},
	{&[]string{"ding", "dong", "dang"}, &[]string{"ding", "dong"}, true, false, []string{"ding", "dong"}},
	{&[]string{}, &[]string{}, false, false, nil},
	{&[]string{}, &[]string{"ding"}, true, false, []string{"ding"}},
	{&[]string{}, nil, false, false, nil},
	{&[]string{}, nil, false, true, nil},
	{nil, &[]string{}, true, false, []string{}},
	{nil, &[]string{}, false, true, nil},
	{nil, nil, false, false, nil},
	{&[]string{"ding"}, &[]string{}, true, false, []string{}},
	{&[]string{"ding"}, nil, false, false, nil},
}

func (t stringArrayTest) formatErrorBase(index int) string {
	return fmt.Sprintf("Failure with example %v:\n\tbase: '%v'\n\tnew: '%v'\n", index, t.baseValue, t.newValue)
}

func TestStringArrayDiffs(t *testing.T) {
	a, b := *new(Location), new(Location)
	for i, data := range stringArrayTests {
		a.PaymentOptions, b.PaymentOptions = data.baseValue, data.newValue
		a.nilIsEmpty, b.nilIsEmpty = data.nilIsEmpty, data.nilIsEmpty
		d, isDiff := a.Diff(b)
		if isDiff != data.isDiff {
			t.Errorf("%vExpected diff to be %v\nbut was %v\ndiff struct was %v\n", data.formatErrorBase(i), data.isDiff, isDiff, d)
		}
		if d == nil && data.expectedFieldValue == nil {
			continue
		} else if d == nil && data.expectedFieldValue != nil {
			t.Errorf("%v\ndelta was nil but expected %v\n", data.formatErrorBase(i), data.expectedFieldValue)
		} else if len(d.GetPaymentOptions()) != len(data.expectedFieldValue) {
			t.Errorf("%v\ndiff was%v\n", data.formatErrorBase(i), d)
		} else {
			for i := 0; i < len(d.GetPaymentOptions()); i++ {
				if d.GetPaymentOptions()[i] != data.expectedFieldValue[i] {
					t.Errorf("%v\ndiff was%v\n", data.formatErrorBase(i), d)
				}
			}
		}
	}
}

type floatTest struct {
	baseValue          *float64
	newValue           *float64
	isDiff             bool
	nilIsEmpty         bool
	expectedFieldValue *float64
}

var floatTests = []floatTest{
	{Float(1234.0), Float(1234.0), false, false, nil},
	{Float(1234.0), nil, false, false, nil},
	{Float(0), nil, false, false, nil},
	{nil, nil, false, false, nil},
	{Float(0), Float(0), false, false, nil},
	{Float(0), Float(9876.0), true, false, Float(9876.0)},
	{Float(1234.0), Float(9876.0), true, false, Float(9876.0)},
	{nil, Float(9876.0), true, false, Float(9876.0)},
	{nil, Float(9876.0), true, true, Float(9876.0)},
	{nil, Float(0), true, false, Float(0)},
	{nil, Float(0), false, true, nil},
}

func formatFloatPtr(b *float64) string {
	if b == nil {
		return "nil"
	} else {
		return fmt.Sprintf("%v", *b)
	}
}

func (t floatTest) formatErrorBase(index int) string {
	return fmt.Sprintf("Failure with example %v:\n\tbase: '%v'\n\tnew: '%v'", index, formatFloatPtr(t.baseValue), formatFloatPtr(t.newValue))
}

func TestFloatDiffs(t *testing.T) {
	a, b := *new(Location), new(Location)
	for i, data := range floatTests {
		a.DisplayLat, b.DisplayLat = data.baseValue, data.newValue
		a.nilIsEmpty, b.nilIsEmpty = data.nilIsEmpty, data.nilIsEmpty
		d, isDiff := a.Diff(b)
		if isDiff != data.isDiff {
			t.Errorf("%vExpected diff to be %v\nbut was %v\ndiff struct was %v\n", data.formatErrorBase(i), data.isDiff, isDiff, d)
		}
		if d == nil && data.expectedFieldValue == nil {
			continue
		} else if d == nil && data.expectedFieldValue != nil {
			t.Errorf("%v\ndelta was nil but expected %v\n", data.formatErrorBase(i), formatFloatPtr(data.expectedFieldValue))
		} else if d != nil && data.expectedFieldValue == nil {
			t.Errorf("%v\ndiff was%v\n", data.formatErrorBase(i), d)
		} else if *d.DisplayLat != *data.expectedFieldValue {
			t.Errorf("%v\ndiff was%v\n", data.formatErrorBase(i), d)
		}
	}
}

type photoTest struct {
	baseValue          *LocationPhoto
	newValue           *LocationPhoto
	isDiff             bool
	nilIsEmpty         bool
	expectedFieldValue *LocationPhoto
}

func formatPhoto(b *LocationPhoto) string {
	if b == nil {
		return "nil"
	} else {
		return fmt.Sprintf("%v", *b)
	}
}

func (t photoTest) formatErrorBase(index int) string {
	return fmt.Sprintf("Failure with example %v:\n\tbase: '%v'\n\tnew: '%v'", index, formatPhoto(t.baseValue), formatPhoto(t.newValue))
}

var photoTests = []photoTest{
	{&LocationPhoto{Url: "ding", Description: "dong"}, &LocationPhoto{Url: "ding", Description: "dong"}, false, false, nil},
	{&LocationPhoto{Url: "ding", Description: "dong"}, nil, false, false, nil},
	{&LocationPhoto{}, nil, false, false, nil},
	{&LocationPhoto{}, nil, false, true, nil},
	{nil, &LocationPhoto{}, true, false, &LocationPhoto{}},
	{nil, &LocationPhoto{}, false, true, nil},
	{nil, &LocationPhoto{Url: "ding", Description: "dong"}, true, false, &LocationPhoto{Url: "ding", Description: "dong"}},
	{&LocationPhoto{Url: "ding"}, &LocationPhoto{Url: "ding", Description: "dong"}, true, false, &LocationPhoto{Url: "ding", Description: "dong"}},
	{&LocationPhoto{Description: "dong"}, &LocationPhoto{Url: "ding", Description: "dong"}, true, false, &LocationPhoto{Url: "ding", Description: "dong"}},
}

func TestPhotoDiffs(t *testing.T) {
	a, b := *new(Location), new(Location)
	for i, data := range photoTests {
		a.FacebookCoverPhoto, b.FacebookCoverPhoto = data.baseValue, data.newValue
		a.nilIsEmpty, b.nilIsEmpty = data.nilIsEmpty, data.nilIsEmpty
		d, isDiff := a.Diff(b)
		if isDiff != data.isDiff {
			t.Errorf("%vExpected diff to be %v\nbut was %v\ndiff struct was %v\n", data.formatErrorBase(i), data.isDiff, isDiff, d)
		}
		if d == nil && data.expectedFieldValue == nil {
			continue
		} else if d == nil && data.expectedFieldValue != nil {
			t.Errorf("%v\ndelta was nil but expected %v\n", data.formatErrorBase(i), formatPhoto(data.expectedFieldValue))
		} else if d != nil && data.expectedFieldValue == nil {
			t.Errorf("%v\ndiff was%v\n", data.formatErrorBase(i), d)
		} else if *d.FacebookCoverPhoto != *data.expectedFieldValue {
			t.Errorf("%v\ndiff was%v\n", data.formatErrorBase(i), d)
		}
	}
}

type photoArrayTest struct {
	baseValue          *[]LocationPhoto
	newValue           *[]LocationPhoto
	isDiff             bool
	nilIsEmpty         bool
	expectedFieldValue *[]LocationPhoto
}

var photoArrayTests = []photoArrayTest{
	{nil, &[]LocationPhoto{}, true, false, &[]LocationPhoto{}},
	{nil, &[]LocationPhoto{}, false, true, nil},
	{&[]LocationPhoto{}, nil, false, false, nil},
	{&[]LocationPhoto{}, nil, false, true, nil},
	{nil, nil, false, false, nil},
	{&[]LocationPhoto{LocationPhoto{Url: "ding", Description: "dong"}}, &[]LocationPhoto{}, true, false, &[]LocationPhoto{}},
	{&[]LocationPhoto{LocationPhoto{Url: "ding", Description: "dong"}}, nil, false, false, nil},
	{&[]LocationPhoto{LocationPhoto{Url: "ding", Description: "dong"}}, &[]LocationPhoto{LocationPhoto{Url: "dong", Description: "ding"}}, true, false, &[]LocationPhoto{LocationPhoto{Url: "dong", Description: "ding"}}},
	{&[]LocationPhoto{LocationPhoto{Url: "ding", Description: "dong"}}, &[]LocationPhoto{LocationPhoto{Url: "ding", Description: "dong"}, LocationPhoto{Url: "ding", Description: "dong"}}, true, false, &[]LocationPhoto{LocationPhoto{Url: "ding", Description: "dong"}, LocationPhoto{Url: "ding", Description: "dong"}}},
}

func (t photoArrayTest) formatErrorBase(index int) string {
	return fmt.Sprintf("Failure with example %v:\n\tbase: '%v'\n\tnew: '%v'", index, t.baseValue, t.newValue)
}

func TestPhotoArrayDiffs(t *testing.T) {
	a, b := *new(Location), new(Location)
	for i, data := range photoArrayTests {
		a.Photos, b.Photos = data.baseValue, data.newValue
		a.nilIsEmpty, b.nilIsEmpty = data.nilIsEmpty, data.nilIsEmpty
		d, isDiff := a.Diff(b)
		if isDiff != data.isDiff {
			t.Errorf("%vExpected diff to be %v\nbut was %v\ndiff struct was %v\n", data.formatErrorBase(i), data.isDiff, isDiff, d)
		}
		if d == nil && data.expectedFieldValue == nil {
			continue
		} else if d == nil && data.expectedFieldValue != nil {
			t.Errorf("%v\ndelta was nil but expected %v\n", data.formatErrorBase(i), data.expectedFieldValue)
		} else if len(*d.Photos) != len(*data.expectedFieldValue) {
			t.Errorf("%v\ndiff was%v\n", data.formatErrorBase(i), d)
		} else {
			for i := 0; i < len(*d.Photos); i++ {
				if (*d.Photos)[i] != (*data.expectedFieldValue)[i] {
					t.Errorf("%v\ndiff was%v\n", data.formatErrorBase(i), d)
				}
			}
		}
	}
}

type listTest struct {
	baseValue          []List
	newValue           []List
	isDiff             bool
	expectedFieldValue []List
}

func (t listTest) formatErrorBase(index int) string {
	return fmt.Sprintf("Failure with example %v:\n\tbase: '%v'\n\tnew: '%v'", index, t.baseValue, t.newValue)
}

type customFieldsTest struct {
	baseValue          map[string]interface{}
	newValue           map[string]interface{}
	isDiff             bool
	nilIsEmpty         bool
	expectedFieldValue map[string]interface{}
}

var baseCustomFields = map[string]interface{}{
	"62150": CustomLocationGallery{
		&CustomLocationPhoto{
			ClickThroughURL: "https://locations.yext.com",
			Description:     "This is the caption",
			Url:             "http://a.mktgcdn.com/p-sandbox/gRcmaehu-FoJtL3Ld6vNjYHpbZxmPSYZ1cTEF_UU7eY/1247x885.png",
		},
	},
	"62151": CustomLocationPhoto{
		ClickThroughURL: "https://locations.yext.com",
		Description:     "This is a caption on a single!",
		Url:             "http://a.mktgcdn.com/p-sandbox/bSZ_mKhfFYGih6-ry5mtbwB_JbKu930kFxHOaQRwZC4/1552x909.png",
	},
	"62152": []string{
		"this",
		"is",
		"a",
		"textlist",
	},
	"62153": "This is a\r\nmulti\r\nline\r\ntext",
	"62154": "This is a single line text.",
	// Hours CustomField Type not really working in the model right now
	// "62144": {
	//   "holidayTimes": [
	//     {
	//       "date": "2015-12-14",
	//       "time": "9:00"
	//     },
	//     {
	//       "date": "2015-12-15",
	//       "time": "0:-1"
	//     }
	//   ],
	//   "dailyTimes": "2:18:00,3:18:00,4:18:00,5:18:00,6:18:00"
	// },
	"62155": "https://locations.yext.com/this-is-a-url",
	"62145": "12/14/2015",
	"62156": "true",
	// Hours CustomField Type not really working in the model right now
	// "62146": {
	//   "additionalHoursText": "We have wacky hours!",
	//   "hours": "2:9:00:18:00,3:19:00:22:00,3:9:00:18:00,4:0:00:0:00,5:9:00:18:00,6:9:00:18:00",
	//   "holidayHours": [
	//     {
	//       "date": "2015-12-12",
	//       "hours": ""
	//     },
	//     {
	//       "date": "2015-12-13",
	//       "hours": "10:00:13:00"
	//     },
	//     {
	//       "date": "2015-12-14",
	//       "hours": "0:00:0:00"
	//     },
	//     {
	//       "date": "2015-12-15",
	//       "hours": "10:00:17:00"
	//     }
	//   ]
	// },
	"62157": map[string]interface{}{"url": "http://www.youtube.com/watch?v=sYMYktsKmSk"},
	"62147": "10",
	"62148": []string{
		"27348",
		"27349",
	},
	"62149": MultiOption{"1", "2"},
}

func copyCustomFields(cf map[string]interface{}) map[string]interface{} {
	n := map[string]interface{}{}
	for key, value := range cf {
		n[key] = value
	}
	return n
}

func appendJunkToCustomFields(cf map[string]interface{}) map[string]interface{} {
	n := copyCustomFields(cf)
	n["guy"] = "random junk"
	return n
}

func deleteKeyFromCustomField(cf map[string]interface{}) map[string]interface{} {
	n := copyCustomFields(cf)
	delete(n, "62148")
	return n
}

func modifyCF(cf map[string]interface{}) map[string]interface{} {
	n := copyCustomFields(cf)
	n["62153"] = "This is a\r\nMODIFIED multi\r\nline\r\ntext"
	return n
}

func reorderedMultiOptionCF(cf map[string]interface{}) map[string]interface{} {
	n := copyCustomFields(cf)
	n["62149"] = MultiOption{"2", "1"}
	return n
}

func zeroCFKEy(cf map[string]interface{}, key string) map[string]interface{} {
	n := copyCustomFields(cf)

	if value, ok := n[key]; ok {
		n[key] = reflect.Zero(reflect.TypeOf(value)).Interface()
	}

	return n
}

var (
	copyOfBase             = copyCustomFields(baseCustomFields)
	appendedCF             = appendJunkToCustomFields(baseCustomFields)
	trimmedCF              = deleteKeyFromCustomField(baseCustomFields)
	modifiedCF             = modifyCF(baseCustomFields)
	differentOptionOrderCF = reorderedMultiOptionCF(baseCustomFields)
)

var customFieldsTests = []customFieldsTest{
	{nil, nil, false, false, nil},
	{map[string]interface{}{}, nil, false, false, nil},
	{map[string]interface{}{}, nil, false, true, nil},
	{map[string]interface{}{}, map[string]interface{}{}, false, false, nil},
	{nil, map[string]interface{}{}, false, false, nil},
	{nil, map[string]interface{}{}, false, true, nil},
	{baseCustomFields, copyOfBase, false, false, nil},
	{baseCustomFields, appendedCF, true, false, map[string]interface{}{"guy": "random junk"}},
	{baseCustomFields, trimmedCF, false, false, nil},
	{baseCustomFields, modifiedCF, true, false, map[string]interface{}{"62153": "This is a\r\nMODIFIED multi\r\nline\r\ntext"}},
	{baseCustomFields, differentOptionOrderCF, false, false, nil},
}

func addZeroTests() {
	for key, val := range baseCustomFields {
		z := zeroCFKEy(baseCustomFields, key)
		zeroForKey := reflect.Zero(reflect.TypeOf(val)).Interface()
		test := customFieldsTest{baseCustomFields, z, true, false, map[string]interface{}{key: zeroForKey}}
		customFieldsTests = append(customFieldsTests, test)
	}
}

func (t customFieldsTest) formatErrorBase(index int) string {
	return fmt.Sprintf("Failure with example %v:\n\tbase: '%v'\n\tnew: '%v'", index, t.baseValue, t.newValue)
}

func TestCustomFieldsDiff(t *testing.T) {
	addZeroTests()
	a, b := *new(Location), new(Location)
	for i, data := range customFieldsTests {
		a.CustomFields, b.CustomFields = data.baseValue, data.newValue
		a.nilIsEmpty, b.nilIsEmpty = data.nilIsEmpty, data.nilIsEmpty
		d, isDiff := a.Diff(b)
		if isDiff != data.isDiff {
			t.Errorf("%vExpected diff to be %v\nbut was %v\ndiff struct was %v\n", data.formatErrorBase(i), data.isDiff, isDiff, d)
		}
		if d == nil && data.expectedFieldValue == nil {
			continue
		} else if d == nil && data.expectedFieldValue != nil {
			t.Errorf("%v\ndelta was nil but expected %v\n", data.formatErrorBase(i), data.expectedFieldValue)
		} else if !reflect.DeepEqual(data.expectedFieldValue, d.CustomFields) {
			t.Errorf("%v\ndiff was%v\n", data.formatErrorBase(i), d)
		}
	}
}

type customFieldsNilIsEmptyTest struct {
	baseValue          map[string]interface{}
	baseNilIsEmpty     bool
	newValue           map[string]interface{}
	newNilIsEmpty      bool
	isDiff             bool
	expectedFieldValue map[string]interface{}
}

var customFieldsNilIsEmptyTests = []customFieldsNilIsEmptyTest{
	{
		baseValue:          map[string]interface{}{},
		baseNilIsEmpty:     true,
		newValue:           map[string]interface{}{"65": ""},
		newNilIsEmpty:      false,
		isDiff:             false,
		expectedFieldValue: nil,
	},
	{
		baseValue:          map[string]interface{}{},
		baseNilIsEmpty:     false,
		newValue:           map[string]interface{}{"65": ""},
		newNilIsEmpty:      false,
		isDiff:             true,
		expectedFieldValue: map[string]interface{}{"65": ""},
	},
	{
		baseValue:          map[string]interface{}{},
		baseNilIsEmpty:     true,
		newValue:           map[string]interface{}{"65": "yext"},
		newNilIsEmpty:      false,
		isDiff:             true,
		expectedFieldValue: map[string]interface{}{"65": "yext"},
	},
	{
		baseValue:          map[string]interface{}{},
		baseNilIsEmpty:     true,
		newValue:           map[string]interface{}{"65": []string{}},
		newNilIsEmpty:      false,
		isDiff:             false,
		expectedFieldValue: nil,
	},
	{
		baseValue:          map[string]interface{}{},
		baseNilIsEmpty:     false,
		newValue:           map[string]interface{}{"65": []string{}},
		newNilIsEmpty:      false,
		isDiff:             true,
		expectedFieldValue: map[string]interface{}{"65": []string{}},
	},
	{
		baseValue:          map[string]interface{}{},
		baseNilIsEmpty:     false,
		newValue:           map[string]interface{}{"65": []string{"ding"}},
		newNilIsEmpty:      false,
		isDiff:             true,
		expectedFieldValue: map[string]interface{}{"65": []string{"ding"}},
	},
	{
		baseValue:          map[string]interface{}{},
		baseNilIsEmpty:     true,
		newValue:           map[string]interface{}{"65": nil},
		newNilIsEmpty:      false,
		isDiff:             false,
		expectedFieldValue: nil,
	},
	{
		baseValue:          map[string]interface{}{},
		baseNilIsEmpty:     false,
		newValue:           map[string]interface{}{"65": nil},
		newNilIsEmpty:      false,
		isDiff:             true,
		expectedFieldValue: map[string]interface{}{"65": nil},
	},
}

func (t customFieldsNilIsEmptyTest) formatErrorBase(index int) string {
	return fmt.Sprintf("Failure with example %v:\n\tbase: '%v'\n\tnew: '%v'", index, t.baseValue, t.newValue)
}

func TestCustomFieldsNilIsEmptyDiff(t *testing.T) {
	a, b := *new(Location), new(Location)
	for i, test := range customFieldsNilIsEmptyTests {
		a.CustomFields, b.CustomFields = test.baseValue, test.newValue
		a.nilIsEmpty, b.nilIsEmpty = test.baseNilIsEmpty, test.newNilIsEmpty
		d, isDiff := a.Diff(b)
		if isDiff != test.isDiff {
			t.Errorf("%vExpected diff to be %v\nbut was %v\ndiff struct was %v\n", test.formatErrorBase(i), test.isDiff, isDiff, d)
		}
		if d == nil && test.expectedFieldValue == nil {
			continue
		} else if d == nil && test.expectedFieldValue != nil {
			t.Errorf("%v\ndelta was nil but expected %v\n", test.formatErrorBase(i), test.expectedFieldValue)
		} else if !reflect.DeepEqual(test.expectedFieldValue, d.CustomFields) {
			t.Errorf("%v\ndiff was%v\n", test.formatErrorBase(i), d)
		}
	}
}

type closedTest struct {
	baseValue          *LocationClosed
	newValue           *LocationClosed
	isDiff             bool
	nilIsEmpty         bool
	expectedFieldValue *LocationClosed
}

var closedTests = []closedTest{
	{nil, nil, false, false, nil},
	{&LocationClosed{}, nil, false, false, nil},
	{&LocationClosed{}, nil, false, true, nil},
	{&LocationClosed{}, &LocationClosed{}, false, false, nil},
	{nil, &LocationClosed{}, true, false, &LocationClosed{}},
	{nil, &LocationClosed{}, false, true, nil},
	{
		nil,
		&LocationClosed{
			IsClosed:   Bool(true),
			ClosedDate: "1/1/2001",
		},
		true,
		false,
		&LocationClosed{
			IsClosed:   Bool(true),
			ClosedDate: "1/1/2001",
		},
	},
	{
		&LocationClosed{
			IsClosed:   Bool(true),
			ClosedDate: "1/1/2001",
		},
		&LocationClosed{
			IsClosed:   Bool(true),
			ClosedDate: "1/1/2001",
		},
		false,
		false,
		nil,
	},
	{
		&LocationClosed{
			IsClosed:   Bool(true),
			ClosedDate: "1/1/2001",
		},
		nil,
		false,
		false,
		nil,
	},
	{
		&LocationClosed{
			IsClosed:   Bool(true),
			ClosedDate: "1/1/2001",
		},
		&LocationClosed{
			IsClosed:   Bool(false),
			ClosedDate: "1/1/2001",
		},
		true,
		false,
		&LocationClosed{
			IsClosed:   Bool(false),
			ClosedDate: "1/1/2001",
		},
	},
	{
		&LocationClosed{
			IsClosed:   Bool(true),
			ClosedDate: "1/1/2001",
		},
		&LocationClosed{
			IsClosed:   Bool(true),
			ClosedDate: "1/1/2002",
		},
		true,
		false,
		&LocationClosed{
			IsClosed:   Bool(true),
			ClosedDate: "1/1/2002",
		},
	},
}

func formatClosed(b *LocationClosed) string {
	if b == nil {
		return "nil"
	} else {
		return fmt.Sprintf("%v", *b)
	}
}

func (t closedTest) formatErrorBase(index int) string {
	return fmt.Sprintf("Failure with example %v:\n\tbase: '%v'\n\tnew: '%v'", index, formatClosed(t.baseValue), formatClosed(t.newValue))
}

func TestClosedDiffs(t *testing.T) {
	a, b := *new(Location), new(Location)
	for i, data := range closedTests {
		a.Closed, b.Closed = data.baseValue, data.newValue
		a.nilIsEmpty, b.nilIsEmpty = data.nilIsEmpty, data.nilIsEmpty
		d, isDiff := a.Diff(b)
		if isDiff != data.isDiff {
			t.Errorf("%vExpected diff to be %v\nbut was %v\ndiff struct was %v\n", data.formatErrorBase(i), data.isDiff, isDiff, d)
		}
		if d == nil && data.expectedFieldValue == nil {
			continue
		} else if d == nil && data.expectedFieldValue != nil {
			t.Errorf("%v\ndelta was nil but expected %v\n", data.formatErrorBase(i), formatClosed(data.expectedFieldValue))
		} else if d != nil && data.expectedFieldValue == nil {
			t.Errorf("%v\ndelta was not nil but expected nil\n diff:%v\n", data.formatErrorBase(i), d)
		} else if GetBool(d.Closed.IsClosed) != GetBool(data.expectedFieldValue.IsClosed) {
			t.Errorf("%v\ndiff was %v\n", data.formatErrorBase(i), d)
		} else if d.Closed.ClosedDate != data.expectedFieldValue.ClosedDate {
			t.Errorf("%v\ndiff was %v\n", data.formatErrorBase(i), d)
		}
	}
}

func TestComplexDiffs(t *testing.T) {
	matt, ben := Location{Name: String("matt"), Emails: &[]string{"matt@yext.com"}}, Location{Name: String("ben"), Emails: &[]string{"ben@yext.com"}}
	delta, isDiff := matt.Diff(&ben)
	if !isDiff {
		t.Errorf("Expected true diff was false\ndelta was:\n%v\n", delta)
	}
	if delta.GetName() != "ben" || len(delta.GetEmails()) != 1 || delta.GetEmails()[0] != "ben@yext.com" {
		t.Errorf("Delta was not as expected\ndelta was:\n%v\nexpected\v%v\n", delta, ben)
	}
}

func TestComplexIdentical(t *testing.T) {
	delta, isDiff := complexOne.Diff(complexTwo)
	if isDiff {
		t.Errorf("Expected false but was true, delta was:\n%v\n", delta)
	}
	if delta != nil {
		t.Errorf("Expected nil delta but was non-nil, delta was:\n%v\n", delta)
	}
}

func TestDiffIds(t *testing.T) {
	tests := []struct {
		BaseLoc *Location
		NewLoc  *Location
		IsDiff  bool
		Delta   *Location
	}{
		{
			BaseLoc: &Location{
				Id: String("1"),
			},
			NewLoc: &Location{
				Id: String("2"),
			},
			IsDiff: true,
			Delta: &Location{
				Id: String("2"),
			},
		},
		{
			BaseLoc: &Location{
				Id: String("1"),
			},
			NewLoc: &Location{
				Id: String("1"),
			},
			IsDiff: false,
		},
		{
			BaseLoc: &Location{
				Id: String("1"),
			},
			NewLoc: &Location{
				Id:   String("2"),
				Name: String("New Name"),
			},
			IsDiff: true,
			Delta: &Location{
				Id:   String("2"),
				Name: String("New Name"),
			},
		},
		{
			BaseLoc: &Location{
				Id: String("1"),
			},
			NewLoc: &Location{
				Id:   String("1"),
				Name: String("New Name"),
			},
			IsDiff: true,
			Delta: &Location{
				Id:   String("1"), // Historically, even though there is no diff in the Id, we always want the Id set
				Name: String("New Name"),
			},
		},
	}
	for _, test := range tests {
		delta, isDiff := test.BaseLoc.Diff(test.NewLoc)
		if isDiff != test.IsDiff {
			t.Errorf("Expected %t but was %t, delta was:\n%v\n", test.IsDiff, isDiff, delta)
		} else if test.IsDiff && !reflect.DeepEqual(test.Delta, delta) {
			t.Errorf("Expected %v for delta, delta was \n%v\n", test.Delta, delta)
		}
	}
}

func TestUnmarshal(t *testing.T) {
	var one, two = new(Location), new(Location)
	err := json.Unmarshal([]byte(jsonData), one)
	if err != nil {
		t.Errorf(err.Error())
	}

	err = json.Unmarshal([]byte(jsonData), two)
	if err != nil {
		t.Errorf(err.Error())
	}

	delta, isDiff := one.Diff(two)
	if isDiff {
		t.Errorf("Expected false but was true, delta was:\n%v\n", delta)
	}
	if delta != nil {
		t.Errorf("Expected nil delta but was non-nil, delta was:\n%v\n", delta)
	}
}

type Scenario struct {
	A         *Location
	B         *Location
	WantDelta *Location
	WantDiff  bool
}

func TestLabels(t *testing.T) {
	var (
		one   = UnorderedStrings([]string{"One", "Two", "Three"})
		two   = UnorderedStrings([]string{"One", "Two", "Three"})
		three = UnorderedStrings([]string{"Three", "One", "Two"})
		four  = UnorderedStrings([]string{"One", "Two"})
		tests = []Scenario{
			Scenario{
				A: &Location{
					Id:       String("1"),
					LabelIds: &one,
				},
				B: &Location{
					Id:       String("1"),
					LabelIds: &two,
				},
				WantDelta: nil,
				WantDiff:  false,
			},
			Scenario{
				A: &Location{
					Id:       String("1"),
					LabelIds: &one,
				},
				B: &Location{
					Id:       String("1"),
					LabelIds: &three,
				},
				WantDelta: nil,
				WantDiff:  false,
			},
			Scenario{
				A: &Location{
					Id:       String("1"),
					LabelIds: &one,
				},
				B: &Location{
					Id:       String("1"),
					LabelIds: &four,
				},
				WantDelta: &Location{
					Id:       String("1"),
					LabelIds: &four,
				},
				WantDiff: true,
			},
			Scenario{
				A: &Location{
					LabelIds: nil,
				},
				B: &Location{
					LabelIds: nil,
				},
				WantDiff: false,
			},
		}
	)

	for i, test := range tests {
		test.A.hydrated = true
		test.B.hydrated = true
		if delta, diff := test.A.Diff(test.B); diff != test.WantDiff {
			t.Errorf("Test %d:\n\tA:\t%+v\n\tB:\t%+v\n\tDelta:\t%+v\n\tWanted:%+v\n\tDiff was: %t wanted %t", i, test.A, test.B, delta, test.WantDelta, diff, test.WantDiff)
		}
	}
}

// Designed to test nilIsTrue of Location
// Location without any fields should be equal to Location with fields that are there but are zero values
// Float 0 is possible so nilIsTrue does not affect that
func TestLocationNils(t *testing.T) {
	a, b := *new(Location), new(Location)
	b.Name = String("")
	b.Emails = &[]string{}
	b.Headshot = &LocationPhoto{}
	b.GoogleAttributes = &LocationGoogleAttributes{}

	a.nilIsEmpty, b.nilIsEmpty = true, true
	d, isDiff := a.Diff(b)

	if isDiff != false {
		t.Errorf("Expected diff to be false but was %v\ndiff struct was %v\n", isDiff, d)
	}
}

func TestLocationCustomFieldEmptyComparision(t *testing.T) {
	a, b := *new(Location), new(Location)
	a.Id = String("blah")
	b.CustomFields = map[string]interface{}{}

	a.hydrated, b.hydrated = true, true
	d, isDiff := a.Diff(b)

	if isDiff != false {
		t.Errorf("Expected diff to be false but was %v\ndiff struct was %v\n", isDiff, d)
	}
}

func TestCustomFieldPointerComparison(t *testing.T) {
	a, b := *new(Location), new(Location)
	a.Id = String("blah")
	a.CustomFields = map[string]interface{}{
		"1": CustomLocationHours{Hours: "1:09:00:18:00"},
	}
	b.CustomFields = map[string]interface{}{
		"1": &CustomLocationHours{Hours: "1:09:00:18:00"},
	}

	a.hydrated, b.hydrated = true, true
	d, isDiff := a.Diff(b)

	if isDiff != false {
		t.Errorf("Expected diff to be false but was %v\ndiff struct was %v\n", isDiff, d)
	}
}

func TestZeroValuesAndNilDiffing(t *testing.T) {
	tests := []Scenario{
		Scenario{
			A:         &Location{SuppressAddress: nil},
			B:         &Location{SuppressAddress: nil},
			WantDelta: nil,
			WantDiff:  false,
		},
		Scenario{
			A:         &Location{SuppressAddress: Bool(false)},
			B:         &Location{SuppressAddress: nil},
			WantDelta: nil,
			WantDiff:  false,
		},
		Scenario{
			A:         &Location{SuppressAddress: nil},
			B:         &Location{SuppressAddress: Bool(false)},
			WantDelta: &Location{SuppressAddress: Bool(false)},
			WantDiff:  true,
		},
		Scenario{
			A:         &Location{nilIsEmpty: true, SuppressAddress: nil},
			B:         &Location{SuppressAddress: Bool(false)},
			WantDelta: nil,
			WantDiff:  false,
		},
		Scenario{
			A:         &Location{},
			B:         &Location{CustomFields: map[string]interface{}{"1": UnsetPhotoValue}},
			WantDelta: &Location{CustomFields: map[string]interface{}{"1": UnsetPhotoValue}},
			WantDiff:  true,
		},
		Scenario{
			A:         &Location{nilIsEmpty: true},
			B:         &Location{CustomFields: map[string]interface{}{"1": UnsetPhotoValue}},
			WantDelta: nil,
			WantDiff:  false,
		},
		Scenario{
			A:         &Location{CustomFields: map[string]interface{}{"1": UnsetPhotoValue}},
			B:         &Location{CustomFields: map[string]interface{}{"1": UnsetPhotoValue}},
			WantDelta: nil,
			WantDiff:  false,
		},
	}

	for i, test := range tests {
		if delta, diff := test.A.Diff(test.B); diff != test.WantDiff {
			t.Errorf("Test %d:\n\tA:\t%+v\n\tB:\t%+v\n\tDelta:\t%+v\n\tWanted:%+v\n\tDiff was: %t wanted %t", i, test.A, test.B, delta, test.WantDelta, diff, test.WantDiff)
		}
	}
}

var hoursTests = []struct {
	A, B                     *string
	WantEquivalent, WantDiff bool
}{
	{
		A:              String(""),
		B:              String(""),
		WantEquivalent: true,
		WantDiff:       false,
	},
	{
		A:              String("1:closed,2:closed,3:closed,4:closed,5:closed,6:closed,7:closed"),
		B:              String(""),
		WantEquivalent: true,
		WantDiff:       false,
	},
	// This might seem odd, but we're still working out hours semantics with Product, so I'd rather err on the side
	// of a limited set of 'closed' equivalencies for now:
	{
		A:              String("1:closed"),
		B:              String(""),
		WantEquivalent: false,
		WantDiff:       true,
	},
	{
		A:              String("1:closed,2:closed,3:closed,4:closed,5:closed,6:closed,7:closed"),
		B:              String("1:closed,2:closed,3:closed,4:closed,5:closed,6:closed,7:closed"),
		WantEquivalent: true,
		WantDiff:       false,
	},
	{
		A:              String("1:11:00"),
		B:              String("1:closed,2:closed,3:closed,4:closed,5:closed,6:closed,7:closed"),
		WantEquivalent: false,
		WantDiff:       true,
	},
	{
		A:              String("1:11:00:20:00,2:10:00:21:00,3:10:00:21:00,4:10:00:21:00,5:10:00:21:00,6:10:00:21:00,7:10:00:21:00"),
		B:              String("1:11:00:20:00,2:10:00:21:00,3:10:00:21:00,4:10:00:21:00,5:10:00:21:00,6:10:00:21:00,7:10:00:21:00"),
		WantEquivalent: true,
		WantDiff:       false,
	},
	{
		A:              String("1:11:00:20:00,2:10:00:21:00,3:10:00:21:00,4:10:00:21:00,5:10:00:21:00,6:10:00:21:00,7:10:00:21:00"),
		B:              String("1:11:01:20:00,2:10:00:21:00,3:10:00:21:00,4:10:00:21:00,5:10:00:21:00,6:10:00:21:00,7:10:00:21:00"),
		WantEquivalent: false,
		WantDiff:       true,
	},
	{
		A:              String("1:11:00:20:00"),
		B:              String("1:11:00:20:00"),
		WantEquivalent: true,
		WantDiff:       false,
	},
	{
		A:              String("1:11:00:20:00"),
		B:              String("1:11:01:20:00"),
		WantEquivalent: false,
		WantDiff:       true,
	},
	{
		A:              nil,
		B:              String("1:11:01:20:00"),
		WantEquivalent: false,
		WantDiff:       true,
	},
	{
		A:              String("1:11:01:20:00"),
		B:              nil,
		WantEquivalent: false,
		WantDiff:       false,
	},
	{
		A:              nil,
		B:              nil,
		WantEquivalent: true,
		WantDiff:       false,
	},
}

func TestHoursAreEquivalent(t *testing.T) {
	for _, test := range hoursTests {
		if test.A != nil && test.B != nil {
			if got := HoursAreEquivalent(*test.A, *test.B); got != test.WantEquivalent {
				t.Errorf(`HoursAreEquivalent("%s", "%s")=%t, wanted %t`, stringify(test.A), stringify(test.B), got, test.WantEquivalent)
			}
			if got := HoursAreEquivalent(*test.B, *test.A); got != test.WantEquivalent {
				t.Errorf(`HoursAreEquivalent("%s", "%s")=%t, wanted %t`, stringify(test.B), stringify(test.A), got, test.WantEquivalent)
			}
		}
	}
}

func stringify(v *string) string {
	if v != nil {
		return *v
	}
	return "nil"
}

func TestHoursAreEquivalentDiff(t *testing.T) {
	for _, test := range hoursTests {
		a := &Location{Hours: test.A}
		b := &Location{Hours: test.B}

		if _, isDiff := a.Diff(b); isDiff != test.WantDiff {
			t.Errorf(`Diff("%s", "%s")=%t, wanted %t`, stringify(test.A), stringify(test.B), isDiff, test.WantDiff)
		}
	}
}

func TestIsZeroValue(t *testing.T) {
	tests := []struct {
		name       string
		i          interface{}
		nilIsEmpty bool
		want       bool
	}{
		{
			name: "Non-Empty String",
			i:    "CTG",
			want: false,
		},
		{
			name:       "Non-Empty String (nil is empty)",
			i:          "CTG",
			nilIsEmpty: true,
			want:       false,
		},
		{
			name: "Empty String",
			i:    "",
			want: true,
		},
		{
			name:       "Empty String (nil is empty)",
			i:          "",
			nilIsEmpty: true,
			want:       true,
		},
		{
			name: "Empty String",
			i:    "",
			want: true,
		},
		{
			name: "Nil *string",
			i:    (*string)(nil),
			want: false,
		},
		{
			name:       "Nil *string (nil is empty)",
			i:          (*string)(nil),
			nilIsEmpty: true,
			want:       true,
		},
		{
			name: "Empty *string",
			i:    String(""),
			want: true,
		},
		{
			name:       "Empty *string (nil is empty)",
			i:          String(""),
			nilIsEmpty: true,
			want:       true,
		},
		{
			name: "nil ptr to *Address struct ",
			i:    (*Address)(nil),
			want: false,
		},
		{
			name:       "nil ptr to *Address struct (nil is empty)",
			i:          (*Address)(nil),
			nilIsEmpty: true,
			want:       true,
		},
		{
			name: "empty *Address struct",
			i:    &Address{},
			want: true,
		},
		{
			name:       "empty *Address struct (nil is empty)",
			i:          &Address{},
			nilIsEmpty: true,
			want:       true,
		},
		{
			name: "non-empty *Address struct",
			i: &Address{
				Line1: String("7900 Westpark"),
			},
			want: false,
		},
		{
			name: "non-empty *Address struct with empty values",
			i: &Address{
				Line1: String(""),
			},
			want: true,
		},
		{
			name: "*Address struct with empty values (nil is empty)",
			i: &Address{
				Line1: String(""),
			},
			nilIsEmpty: true,
			want:       true,
		},
		{
			name: "*Address struct with zero values",
			i: &Address{
				Line1:       String(""),
				Line2:       String(""),
				City:        String(""),
				Region:      String(""),
				Sublocality: String(""),
				PostalCode:  String(""),
			},
			want: true,
		},
		{
			name: "struct with zero values (*Address)",
			i: &Address{
				Line1:       String(""),
				Line2:       String(""),
				City:        String(""),
				Region:      String(""),
				Sublocality: String(""),
				PostalCode:  String(""),
			},
			nilIsEmpty: true,
			want:       true,
		},
		{
			name: "list",
			i: &[]HolidayHours{
				HolidayHours{
					Date:     String("01-21-2019"),
					IsClosed: NullableBool(true),
				},
			},
			nilIsEmpty: true,
			want:       false,
		},
		{
			name: "**struct",
			i: NullableDayHours(&DayHours{
				IsClosed: NullableBool(true),
			}),
			nilIsEmpty: true,
			want:       false,
		},
		{
			name: "**struct",
			i: NullableDayHours(&DayHours{
				IsClosed: NullableBool(false),
			}),
			nilIsEmpty: true,
			want:       true,
		},
		{
			name:       "**struct",
			i:          NullDayHours(),
			nilIsEmpty: false,
			want:       true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if isZeroValue := IsZeroValue(reflect.ValueOf(test.i), test.nilIsEmpty); test.want != isZeroValue {
				t.Errorf(`Expected IsZeroValue: %t\nGot:%t`, test.want, isZeroValue)
			}
		})
	}
}
