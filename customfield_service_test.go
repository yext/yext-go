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
			Id:   "123",
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
	TypeName     string
	Raw          interface{}
	ExpectedType reflect.Type
}

func runParseTest(t *testing.T, index int, c customFieldParseTest) {
	cf, err := parseAs(c.TypeName, c.Raw)
	if err != nil {
		t.Error(err)
		return
	}
	ty := reflect.TypeOf(cf).String()
	expect := c.ExpectedType.String()
	if ty != expect {
		t.Errorf("test #%d failed: expected type %s, got type %s", index, expect, ty)
	}
}

var (
	customPhoto = map[string]interface{}{
		"title":           "A great picture",
		"description":     "This is a picture of an awesome event",
		"clickthroughUrl": "https://yext.com/event",
		"url":             "https://mktgcdn.com/awesome.jpg",
	}
	hours = map[string]interface{}{
		"additionalHoursText": "This is an example of extra hours info",
		"hours":               "1:9:00:17:00,3:15:00:12:00,3:5:00:11:00,4:9:00:17:00,5:0:00:0:00,6:9:00:17:00,7:9:00:17:00",
		"holidayHours": []interface{}{
			map[string]interface{}{
				"date":  "2016-05-30",
				"hours": "",
			},
			map[string]interface{}{
				"date":  "2016-07-04",
				"hours": "0:00:0:00",
			},
			map[string]interface{}{
				"date":  "2016-09-05",
				"hours": "6:00:10:00",
			},
			map[string]interface{}{
				"date":  "2016-09-10",
				"hours": "9:00:17:00",
			},
		},
	}
	video = map[string]interface{}{
		"description": "An example caption for a video",
		"url":         "http://www.youtube.com/watch?v=M80FTIcEgZM",
	}
	dailyTimes = map[string]interface{}{
		"dailyTimes": "2:7:00,3:7:00,4:7:00,5:7:00,6:7:00,7:7:00,1:7:00",
	}
	parseTests = []customFieldParseTest{
		customFieldParseTest{"BOOLEAN", false, reflect.TypeOf(YesNo(false))},
		customFieldParseTest{"BOOLEAN", "false", reflect.TypeOf(YesNo(false))},
		customFieldParseTest{"NUMBER", "12345", reflect.TypeOf(Number(""))},
		customFieldParseTest{"TEXT", "foo", reflect.TypeOf(SingleLineText(""))},
		customFieldParseTest{"MULTILINE_TEXT", "foo", reflect.TypeOf(MultiLineText(""))},
		customFieldParseTest{"SINGLE_OPTION", "foo", reflect.TypeOf(GetSingleOptionPointer(SingleOption("")))},
		customFieldParseTest{"URL", "foo", reflect.TypeOf(Url(""))},
		customFieldParseTest{"DATE", "foo", reflect.TypeOf(Date(""))},
		customFieldParseTest{"TEXT_LIST", []string{"a", "b", "c"}, reflect.TypeOf(TextList([]string{}))},
		customFieldParseTest{"TEXT_LIST", []interface{}{"a", "b", "c"}, reflect.TypeOf(TextList([]string{}))},
		customFieldParseTest{"MULTI_OPTION", []string{"a", "b", "c"}, reflect.TypeOf(MultiOption([]string{}))},
		customFieldParseTest{"MULTI_OPTION", []interface{}{"a", "b", "c"}, reflect.TypeOf(MultiOption([]string{}))},
		customFieldParseTest{"PHOTO", customPhoto, reflect.TypeOf(Photo{})},
		customFieldParseTest{"GALLERY", []interface{}{customPhoto}, reflect.TypeOf(Gallery{})},
		customFieldParseTest{"VIDEO", video, reflect.TypeOf(Video{})},
		customFieldParseTest{"HOURS", hours, reflect.TypeOf(Hours{})},
		customFieldParseTest{"DAILY_TIMES", dailyTimes, reflect.TypeOf(DailyTimes{})},
		customFieldParseTest{"LOCATION_LIST", []string{"a", "b", "c"}, reflect.TypeOf(LocationList([]string{}))},
	}
)

func makeCustomFields(n int) []*CustomField {
	var cfs []*CustomField

	for i := 0; i < n; i++ {
		new := &CustomField{Id: strconv.Itoa(i)}
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
