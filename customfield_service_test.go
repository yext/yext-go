package yext

import (
	"encoding/json"
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
		"sizes": []interface{}{
			map[string]interface{}{
				"Height": 100,
				"Width":  100,
				"Url":    "https://mktgcdn.com/awesome.jpg",
			},
		},
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
	parseTests = []customFieldParseTest{
		customFieldParseTest{CUSTOMFIELDTYPE_YESNO, false, reflect.TypeOf(YesNo(false))},
		customFieldParseTest{CUSTOMFIELDTYPE_YESNO, "false", reflect.TypeOf(YesNo(false))},
		customFieldParseTest{CUSTOMFIELDTYPE_NUMBER, "12345", reflect.TypeOf(Number(""))},
		customFieldParseTest{CUSTOMFIELDTYPE_SINGLELINETEXT, "foo", reflect.TypeOf(SingleLineText(""))},
		customFieldParseTest{CUSTOMFIELDTYPE_MULTILINETEXT, "foo", reflect.TypeOf(MultiLineText(""))},
		customFieldParseTest{CUSTOMFIELDTYPE_SINGLEOPTION, "foo", reflect.TypeOf(SingleOption(""))},
		customFieldParseTest{CUSTOMFIELDTYPE_URL, "foo", reflect.TypeOf(Url(""))},
		customFieldParseTest{CUSTOMFIELDTYPE_DATE, "foo", reflect.TypeOf(Date(""))},
		customFieldParseTest{CUSTOMFIELDTYPE_TEXTLIST, []string{"a", "b", "c"}, reflect.TypeOf(TextList([]string{}))},
		customFieldParseTest{CUSTOMFIELDTYPE_TEXTLIST, []interface{}{"a", "b", "c"}, reflect.TypeOf(TextList([]string{}))},
		customFieldParseTest{CUSTOMFIELDTYPE_MULTIOPTION, []string{"a", "b", "c"}, reflect.TypeOf(MultiOption([]string{}))},
		customFieldParseTest{CUSTOMFIELDTYPE_MULTIOPTION, []interface{}{"a", "b", "c"}, reflect.TypeOf(MultiOption([]string{}))},
		customFieldParseTest{CUSTOMFIELDTYPE_PHOTO, customPhoto, reflect.TypeOf(CustomPhoto{})},
		customFieldParseTest{CUSTOMFIELDTYPE_GALLERY, []interface{}{customPhoto}, reflect.TypeOf(Gallery{})},
		customFieldParseTest{CUSTOMFIELDTYPE_VIDEO, video, reflect.TypeOf(Video{})},
		customFieldParseTest{CUSTOMFIELDTYPE_HOURS, hours, reflect.TypeOf(Hours{})},
	}
)

func TestParsing(t *testing.T) {
	for i, testData := range parseTests {
		runParseTest(t, i, testData)
	}
}

// other tests
func TestMapMarshalling(t *testing.T) {
	target := map[string]interface{}{}
	var cp = &CustomPhoto{}
	target["123"] = cp

	raw := map[string]interface{}{"123": customPhoto, "234": customPhoto}

	b, _ := json.Marshal(raw)
	err := json.Unmarshal(b, &target)

	t.Logf("Err: %v", err)
	t.Logf("Got: %#v", target)
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
