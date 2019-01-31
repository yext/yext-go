package yext

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
	"testing"
)

func parseAs(cftype string, rawval interface{}) (interface{}, error) {
	cfs := []*CustomField{
		&CustomField{
			Id:   String("123"),
			Type: cftype,
		},
	}

	cfsraw := map[string]interface{}{
		"123": rawval,
	}

	newcfs, err := ParseCustomFields(cfsraw, cfs)
	return newcfs["123"], err
}

type customFieldParseTest struct {
	TypeName string
	Raw      interface{}
	Expected interface{}
}

func runParseTest(t *testing.T, index int, c customFieldParseTest) {
	cf, err := parseAs(c.TypeName, c.Raw)
	if err != nil {
		t.Error(err)
		return
	}
	var (
		cfType     = reflect.TypeOf(cf).String()
		expectType = reflect.TypeOf(c.Expected).String()
	)

	if cfType != expectType {
		t.Errorf("test #%d (%s) failed:\nexpected type %s\ngot type %s", index, c.TypeName, expectType, cfType)
	}

	if !reflect.DeepEqual(cf, c.Expected) {
		t.Errorf("test #%d (%s) failed\nexpected value %v\ngot value      %v", index, c.TypeName, c.Expected, cf)
	}
}

var (
	customPhotoRaw = map[string]interface{}{
		"details":         "A great picture",
		"description":     "This is a picture of an awesome event",
		"clickthroughUrl": "https://yext.com/event",
		"url":             "https://mktgcdn.com/awesome.jpg",
	}
	// hoursRaw is in the format used by HoursCustom for location-service
	hoursRaw = map[string]interface{}{
		"additionalHoursText": "This is an example of extra hours info",
		"hours":               "1:9:00:17:00,3:15:00:12:00,3:5:00:11:00,4:9:00:17:00,5:0:00:0:00,6:9:00:17:00,7:9:00:17:00",
		"holidayHours": []interface{}{
			map[string]interface{}{
				"date":  "2016-05-30",
				"hours": "",
			},
			map[string]interface{}{
				"date":  "2016-05-31",
				"hours": "9:00:17:00",
			},
		},
	}
	videoRaw = map[string]interface{}{
		"description": "An example caption for a video",
		"url":         "http://www.youtube.com/watch?v=M80FTIcEgZM",
	}
	// dailyTimesRaw is in the format used by DailyTimesCustom for location-service
	dailyTimesRaw = map[string]interface{}{
		"dailyTimes": "1:10:00;2:4:00;3:5:00;4:6:00;5:7:00;6:8:00;7:9:00",
	}
	parseTests = []customFieldParseTest{
		customFieldParseTest{"BOOLEAN", false, YesNo(false)},
		customFieldParseTest{"BOOLEAN", "false", YesNo(false)},
		customFieldParseTest{"NUMBER", "12345", Number("12345")},
		customFieldParseTest{"TEXT", "foo", SingleLineText("foo")},
		customFieldParseTest{"MULTILINE_TEXT", "foo", MultiLineText("foo")},
		customFieldParseTest{"SINGLE_OPTION", "foo", GetSingleOptionPointer(SingleOption("foo"))},
		customFieldParseTest{"URL", "foo", Url("foo")},
		customFieldParseTest{"DATE", "foo", Date("foo")},
		customFieldParseTest{"TEXT_LIST", []string{"a", "b", "c"}, TextList([]string{"a", "b", "c"})},
		customFieldParseTest{"TEXT_LIST", []interface{}{"a", "b", "c"}, TextList([]string{"a", "b", "c"})},
		customFieldParseTest{"MULTI_OPTION", []string{"a", "b", "c"}, MultiOption([]string{"a", "b", "c"})},
		customFieldParseTest{"MULTI_OPTION", []interface{}{"a", "b", "c"}, MultiOption([]string{"a", "b", "c"})},
		customFieldParseTest{"PHOTO", customPhotoRaw, &CustomLocationPhoto{
			Url:             "https://mktgcdn.com/awesome.jpg",
			Description:     "This is a picture of an awesome event",
			Details:         "A great picture",
			ClickThroughURL: "https://yext.com/event",
		}},
		customFieldParseTest{"PHOTO", nil, (*CustomLocationPhoto)(nil)},
		customFieldParseTest{"GALLERY", []interface{}{customPhotoRaw}, CustomLocationGallery{
			&CustomLocationPhoto{
				Url:             "https://mktgcdn.com/awesome.jpg",
				Description:     "This is a picture of an awesome event",
				Details:         "A great picture",
				ClickThroughURL: "https://yext.com/event",
			},
		}},
		customFieldParseTest{"VIDEO", videoRaw, CustomLocationVideo{
			Url:         "http://www.youtube.com/watch?v=M80FTIcEgZM",
			Description: "An example caption for a video",
		}},
		customFieldParseTest{"HOURS", hoursRaw, CustomLocationHours{
			AdditionalText: "This is an example of extra hours info",
			Hours:          "1:9:00:17:00,3:15:00:12:00,3:5:00:11:00,4:9:00:17:00,5:0:00:0:00,6:9:00:17:00,7:9:00:17:00",
			HolidayHours: []LocationHolidayHours{
				LocationHolidayHours{
					Date:  "2016-05-30",
					Hours: "",
				},
				LocationHolidayHours{
					Date:  "2016-05-31",
					Hours: "9:00:17:00",
				},
			},
		}},
		customFieldParseTest{"DAILY_TIMES", dailyTimesRaw, CustomLocationDailyTimes{
			DailyTimes: "1:10:00;2:4:00;3:5:00;4:6:00;5:7:00;6:8:00;7:9:00",
		}},
		customFieldParseTest{"LOCATION_LIST", []string{"a", "b", "c"}, LocationList([]string{"a", "b", "c"})},
	}
)

func makeCustomFields(n int) []*CustomField {
	var cfs []*CustomField

	for i := 0; i < n; i++ {
		new := &CustomField{Id: String(strconv.Itoa(i))}
		cfs = append(cfs, new)
	}

	return cfs
}

func TestListAll(t *testing.T) {
	maxLimit := strconv.Itoa(CustomFieldListMaxLimit)

	type req struct {
		limit  string
		offset string
	}

	tests := []struct {
		count int
		reqs  []req
	}{
		{
			count: 0,
			reqs:  []req{{limit: maxLimit, offset: ""}},
		},
		{
			count: 1000,
			reqs:  []req{{limit: maxLimit, offset: ""}},
		},
		{
			count: 1001,
			reqs:  []req{{limit: maxLimit, offset: ""}, {limit: maxLimit, offset: "1000"}},
		},
		{
			count: 2000,
			reqs:  []req{{limit: maxLimit, offset: ""}, {limit: maxLimit, offset: "1000"}},
		},
	}

	for _, test := range tests {
		setup()
		reqs := 0
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			reqs++
			if reqs > len(test.reqs) {
				t.Errorf("Too many requests sent to custom field list - got %d, expected %d", reqs, len(test.reqs))
			}

			expectedreq := test.reqs[reqs-1]

			if v := r.URL.Query().Get("limit"); v != expectedreq.limit {
				t.Errorf("Wanted limit %s, got %s", expectedreq.limit, v)
			}

			if v := r.URL.Query().Get("offset"); v != expectedreq.offset {
				t.Errorf("Wanted offset %s, got %s", expectedreq.offset, v)
			}

			cfs := []*CustomField{}
			remaining := test.count - ((reqs - 1) * CustomFieldListMaxLimit)
			if remaining > 0 {
				if remaining > CustomFieldListMaxLimit {
					remaining = CustomFieldListMaxLimit
				}
				cfs = makeCustomFields(remaining)
			}

			v := &mockResponse{Response: &CustomFieldResponse{Count: test.count, CustomFields: cfs}}
			data, _ := json.Marshal(v)
			w.Write(data)
		})

		client.CustomFieldService.ListAll()
		if reqs < len(test.reqs) {
			t.Errorf("Too few requests sent to custom field list - got %d, expected %d", reqs, len(test.reqs))
		}

		teardown()
	}
}

func TestListMismatchCount(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		v := &mockResponse{Response: &CustomFieldResponse{Count: 25, CustomFields: makeCustomFields(24)}}
		data, _ := json.Marshal(v)
		w.Write(data)
	})

	rlr, err := client.CustomFieldService.ListAll()
	if rlr != nil {
		t.Error("Expected response to be nil when receiving mismatched count and number of custom fields")
	}
	if err == nil {
		t.Error("Expected error to be present when receiving mismatched count and number of custom fields")
	}
}

func TestMustCache(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		v := &mockResponse{Response: &CustomFieldResponse{Count: 24, CustomFields: makeCustomFields(24)}}
		data, _ := json.Marshal(v)
		w.Write(data)
	})
	m := client.CustomFieldService.MustCacheCustomFields()
	n := makeCustomFields(24)
	for i := 0; i < 24; i++ {
		if m[i].GetId() != n[i].GetId() {
			t.Error("Must Cache Custom fields should return the same custom field slice as makeCustomFields")
		}
	}
}

func TestParsing(t *testing.T) {
	for i, testData := range parseTests {
		runParseTest(t, i, testData)
	}
}

func TestParseLeaveUnknownTypes(t *testing.T) {
	type peanut []string
	cf, err := parseAs("BLAH", peanut([]string{"a", "b", "c"}))
	if err != nil {
		t.Error(err)
	}

	if _, ok := cf.(peanut); !ok {
		t.Errorf("Expected type peanut, got type %T", cf)
	}
}

func TestGetString(t *testing.T) {
	var cfManager = &LocationCustomFieldManager{
		CustomFields: []*CustomField{
			&CustomField{
				Name: "Single Line Text",
				Id:   String("SingleLineText"),
				Type: CUSTOMFIELDTYPE_SINGLELINETEXT,
			},
			&CustomField{
				Name: "Multi Line Text",
				Id:   String("MultiLineText"),
				Type: CUSTOMFIELDTYPE_MULTILINETEXT,
			},
			&CustomField{
				Name: "Date",
				Id:   String("Date"),
				Type: CUSTOMFIELDTYPE_DATE,
			},
			&CustomField{
				Name: "Number",
				Id:   String("Number"),
				Type: CUSTOMFIELDTYPE_NUMBER,
			},
			&CustomField{
				Name: "Single Option",
				Id:   String("SingleOption"),
				Type: CUSTOMFIELDTYPE_SINGLEOPTION,
				Options: []CustomFieldOption{
					CustomFieldOption{
						Key:   "SingleOptionOneKey",
						Value: "Single Option One Value",
					},
				},
			},
			&CustomField{
				Name: "Url",
				Id:   String("Url"),
				Type: CUSTOMFIELDTYPE_URL,
			},
		},
	}

	loc := &Location{
		CustomFields: map[string]interface{}{
			cfManager.MustCustomFieldId("Single Line Text"): SingleLineText("Single Line Text Value"),
			cfManager.MustCustomFieldId("Multi Line Text"):  MultiLineText("Multi Line Text Value"),
			cfManager.MustCustomFieldId("Date"):             Date("04/16/2018"),
			cfManager.MustCustomFieldId("Number"):           Number("2"),
			cfManager.MustCustomFieldId("Single Option"):    GetSingleOptionPointer(SingleOption(cfManager.MustCustomFieldOptionId("Single Option", "Single Option One Value"))),
			cfManager.MustCustomFieldId("Url"):              Url("www.url.com"),
		},
	}

	blankLoc := &Location{
		CustomFields: map[string]interface{}{
			cfManager.MustCustomFieldId("Single Line Text"): SingleLineText(""),
			cfManager.MustCustomFieldId("Multi Line Text"):  MultiLineText(""),
			cfManager.MustCustomFieldId("Date"):             Date(""),
			cfManager.MustCustomFieldId("Number"):           Number(""),
			cfManager.MustCustomFieldId("Single Option"):    GetSingleOptionPointer(SingleOption("")),
			cfManager.MustCustomFieldId("Url"):              Url(""),
		},
	}

	tests := []struct {
		CFName string
		Loc    *Location
		Want   string
	}{
		{
			CFName: "Single Line Text",
			Loc:    loc,
			Want:   "Single Line Text Value",
		},
		{
			CFName: "Multi Line Text",
			Loc:    loc,
			Want:   "Multi Line Text Value",
		},
		{
			CFName: "Single Option",
			Loc:    loc,
			Want:   "Single Option One Value",
		},
		{
			CFName: "Date",
			Loc:    loc,
			Want:   "04/16/2018",
		},
		{
			CFName: "Number",
			Loc:    loc,
			Want:   "2",
		},
		{
			CFName: "Url",
			Loc:    loc,
			Want:   "www.url.com",
		},
		{
			CFName: "Single Line Text",
			Loc:    blankLoc,
			Want:   "",
		},
		{
			CFName: "Multi Line Text",
			Loc:    blankLoc,
			Want:   "",
		},
		{
			CFName: "Single Option",
			Loc:    blankLoc,
			Want:   "",
		},
		{
			CFName: "Date",
			Loc:    blankLoc,
			Want:   "",
		},
		{
			CFName: "Number",
			Loc:    blankLoc,
			Want:   "",
		},
		{
			CFName: "Url",
			Loc:    blankLoc,
			Want:   "",
		},
	}

	for _, test := range tests {
		if got, err := cfManager.GetString(test.CFName, test.Loc); err != nil {
			t.Errorf("Get String: got err for custom field %s: %s", test.CFName, err)
		} else if got != test.Want {
			t.Errorf("Get String: got '%s', wanted '%s' for custom field %s", got, test.Want, test.CFName)
		}
	}
}

func TestGetStringSlice(t *testing.T) {
	var cfManager = &LocationCustomFieldManager{
		CustomFields: []*CustomField{
			&CustomField{
				Name: "Text List",
				Id:   String("TextList"),
				Type: CUSTOMFIELDTYPE_TEXTLIST,
			},
			&CustomField{
				Name: "Location List",
				Id:   String("LocationList"),
				Type: CUSTOMFIELDTYPE_LOCATIONLIST,
			},
			&CustomField{
				Name: "Multi Option",
				Id:   String("MultiOption"),
				Type: CUSTOMFIELDTYPE_MULTIOPTION,
				Options: []CustomFieldOption{
					CustomFieldOption{
						Key:   "MultiOptionOneKey",
						Value: "Multi Option One Value",
					},
					CustomFieldOption{
						Key:   "MultiOptionTwoKey",
						Value: "Multi Option Two Value",
					},
					CustomFieldOption{
						Key:   "MultiOptionThreeKey",
						Value: "Multi Option Three Value",
					},
				},
			},
		},
	}

	loc := &Location{
		CustomFields: map[string]interface{}{
			cfManager.MustCustomFieldId("Text List"):     TextList([]string{"A", "B", "C"}),
			cfManager.MustCustomFieldId("Location List"): LocationList(UnorderedStrings([]string{"1", "2", "3"})),
			cfManager.MustCustomFieldId("Multi Option"):  MultiOption(UnorderedStrings([]string{"MultiOptionOneKey", "MultiOptionTwoKey", "MultiOptionThreeKey"})),
		},
	}

	blankLoc := &Location{
		CustomFields: map[string]interface{}{
			cfManager.MustCustomFieldId("Text List"):     TextList([]string{}),
			cfManager.MustCustomFieldId("Location List"): LocationList(UnorderedStrings([]string{})),
			cfManager.MustCustomFieldId("Multi Option"):  MultiOption(UnorderedStrings([]string{})),
		},
	}

	tests := []struct {
		CFName string
		Loc    *Location
		Want   []string
	}{
		{
			CFName: "Text List",
			Loc:    loc,
			Want:   []string{"A", "B", "C"},
		},
		{
			CFName: "Location List",
			Loc:    loc,
			Want:   []string{"1", "2", "3"},
		},
		{
			CFName: "Multi Option",
			Loc:    loc,
			Want:   []string{"Multi Option One Value", "Multi Option Two Value", "Multi Option Three Value"},
		},
		{
			CFName: "Text List",
			Loc:    blankLoc,
			Want:   []string{},
		},
		{
			CFName: "Location List",
			Loc:    blankLoc,
			Want:   []string{},
		},
		{
			CFName: "Multi Option",
			Loc:    blankLoc,
			Want:   []string{},
		},
	}

	for _, test := range tests {
		if got, err := cfManager.GetStringSlice(test.CFName, test.Loc); err != nil {
			t.Errorf("Get String Slice: got err for custom field %s: %s", test.CFName, err)
		} else if got == nil && test.Want != nil || got != nil && test.Want == nil {
			t.Errorf("Get String Slice: got %v, wanted %v for custom field %s", got, test.Want, test.CFName)
		} else {
			for i, _ := range got {
				if got[i] != test.Want[i] {
					t.Errorf("Get String Slice: got %v, wanted %v for custom field %s", got, test.Want, test.CFName)
				}
			}
		}
	}
}
