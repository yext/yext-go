package yext

import (
	"reflect"
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
		customFieldParseTest{"SINGLE_OPTION", "foo", reflect.TypeOf(SingleOption(""))},
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
	}
)

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
