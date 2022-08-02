package yext

import (
	"encoding/json"
	"reflect"
	"testing"
)

var CustomLocationEntityCFManager = &CustomFieldManager{
	CustomFields: []*CustomField{
		&CustomField{
			Name: "CF Url",
			Type: CUSTOMFIELDTYPE_URL,
			Id:   String("cf_Url"),
		},
		&CustomField{
			Name: "CF Text",
			Type: CUSTOMFIELDTYPE_SINGLELINETEXT,
			Id:   String("cf_Text"),
		},
	},
}

type CFT struct {
	Text  *string `json:"cft_text,omitempty"`
	Image **Image `json:"cft_image,omitempty"`
	Bool  **bool  `json:"cft_bool,omitempty"`
}

type CustomEntity struct {
	CFHours        **Hours           `json:"cf_Hours,omitempty"`
	CFUrl          *string           `json:"cf_Url,omitempty"`
	CFDailyTimes   **DailyTimes      `json:"cf_DailyTimes,omitempty"`
	CFTextList     *[]string         `json:"cf_TextList,omitempty"`
	CFGallery      *[]Photo          `json:"cf_Gallery,omitempty"`
	CFPhoto        **Photo           `json:"cf_Photo,omitempty"`
	CFVideos       *[]Video          `json:"cf_Videos,omitempty"`
	CFVideo        **Video           `json:"cf_Video,omitempty"`
	CFDate         *string           `json:"cf_Date,omitempty"`
	CFSingleOption **string          `json:"cf_SingleOption,omitempty"`
	CFMultiOption  *UnorderedStrings `json:"cf_MultiOption,omitempty"`
	CFYesNo        **bool            `json:"cf_YesNo,omitempty"`
	CFText         *string           `json:"cf_Text,omitempty"`
	CFMultiLine    *string           `json:"cf_MultiLineText,omitempty"`
	CFImage        **Image           `json:"cf_Image,omitempty"`
	CFType         *CFT              `json:"cf_Type,omitempty"`
}

type CustomLocationEntity struct {
	LocationEntity
	CustomEntity
}

func (y *CustomLocationEntity) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &y.LocationEntity); err != nil {
		return err
	}
	if err := json.Unmarshal(bytes, &y.CustomEntity); err != nil {
		return err
	}
	return nil
}

func (y CustomLocationEntity) String() string {
	b, _ := json.Marshal(y)
	return string(b)
}

func (c *CustomEntity) UnmarshalJSON(data []byte) error {
	type Alias CustomEntity
	a := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	return UnmarshalEntityJSON(c, data)
}

func entityToJSONString(entity Entity) (error, string) {
	buf, err := json.Marshal(entity)
	if err != nil {
		return err, ""
	}

	return nil, string(buf)
}

func TestEntityJSONSerialization(t *testing.T) {
	type test struct {
		entity Entity
		want   string
	}

	tests := []test{
		{&CustomLocationEntity{}, `{}`},
		{&CustomLocationEntity{LocationEntity: LocationEntity{Address: &Address{City: nil}}}, `{"address":{}}`},
		{&CustomLocationEntity{LocationEntity: LocationEntity{Address: &Address{City: String("")}}}, `{"address":{"city":""}}`},
		{&CustomLocationEntity{LocationEntity: LocationEntity{Languages: nil}}, `{}`},
		{&CustomLocationEntity{LocationEntity: LocationEntity{Languages: &[]string{}}}, `{"languages":[]}`},
		{&CustomLocationEntity{LocationEntity: LocationEntity{Languages: &[]string{"English"}}}, `{"languages":["English"]}`},
		{&CustomLocationEntity{LocationEntity: LocationEntity{Hours: nil}}, `{}`},
		{&CustomLocationEntity{LocationEntity: LocationEntity{Hours: NullHours()}}, `{"hours":null}`},
		{&CustomLocationEntity{LocationEntity: LocationEntity{Hours: NullableHours(&Hours{Monday: NullDayHours(), Tuesday: NullDayHours(), Wednesday: NullDayHours(), Thursday: NullDayHours(), Friday: NullDayHours(), Saturday: NullDayHours(), Sunday: NullDayHours()})}}, `{"hours":{"monday":null,"tuesday":null,"wednesday":null,"thursday":null,"friday":null,"saturday":null,"sunday":null}}`},
		{&CustomLocationEntity{CustomEntity: CustomEntity{CFUrl: String("")}}, `{"cf_Url":""}`},
		{&CustomLocationEntity{CustomEntity: CustomEntity{CFUrl: nil}}, `{}`},
		{&CustomLocationEntity{CustomEntity: CustomEntity{CFTextList: &[]string{}}}, `{"cf_TextList":[]}`},
		{&CustomLocationEntity{CustomEntity: CustomEntity{CFTextList: nil}}, `{}`},
		{&CustomLocationEntity{CustomEntity: CustomEntity{CFSingleOption: NullString()}}, `{"cf_SingleOption":null}`},
		{&CustomLocationEntity{CustomEntity: CustomEntity{CFMultiOption: ToUnorderedStrings([]string{})}}, `{"cf_MultiOption":[]}`},
		{&CustomLocationEntity{CustomEntity: CustomEntity{CFDate: String("")}}, `{"cf_Date":""}`},
		{&CustomLocationEntity{CustomEntity: CustomEntity{CFVideo: NullVideo()}}, `{"cf_Video":null}`},
		{&CustomLocationEntity{CustomEntity: CustomEntity{CFPhoto: NullPhoto()}}, `{"cf_Photo":null}`},
		{&CustomLocationEntity{CustomEntity: CustomEntity{CFGallery: &[]Photo{}}}, `{"cf_Gallery":[]}`},
		{&CustomLocationEntity{CustomEntity: CustomEntity{CFVideos: &[]Video{}}}, `{"cf_Videos":[]}`},
		{&CustomLocationEntity{CustomEntity: CustomEntity{CFDailyTimes: NullDailyTimes()}}, `{"cf_DailyTimes":null}`},
		{&CustomLocationEntity{CustomEntity: CustomEntity{CFHours: NullHours()}}, `{"cf_Hours":null}`},
		{&CustomLocationEntity{CustomEntity: CustomEntity{CFYesNo: NullBool()}}, `{"cf_YesNo":null}`},
		{&CustomLocationEntity{LocationEntity: LocationEntity{Name: String("Hello")}, CustomEntity: CustomEntity{CFYesNo: NullBool()}}, `{"name":"Hello","cf_YesNo":null}`},
		{&CustomLocationEntity{LocationEntity: LocationEntity{Name: String("")}, CustomEntity: CustomEntity{CFYesNo: NullBool()}}, `{"name":"","cf_YesNo":null}`},
		{&CustomLocationEntity{CustomEntity: CustomEntity{CFText: String("")}}, `{"cf_Text":""}`},
	}

	for _, test := range tests {
		if err, got := entityToJSONString(test.entity); err != nil {
			t.Error("Unable to convert", test.entity, "to JSON:", err)
		} else if got != test.want {
			t.Errorf("json.Marshal()\nGot:      %s\nExpected: %s", got, test.want)
		}
	}
}

func TestEntityJSONDeserialization(t *testing.T) {
	type test struct {
		json string
		want Entity
	}

	tests := []test{
		{`{}`, &CustomLocationEntity{}},
		{`{"emails": []}`, &CustomLocationEntity{LocationEntity: LocationEntity{Emails: Strings([]string{})}}},
		{`{"emails": ["bob@email.com", "sue@email.com"]}`, &CustomLocationEntity{LocationEntity: LocationEntity{Emails: Strings([]string{"bob@email.com", "sue@email.com"})}}},
		{`{"emails": ["bob@email.com", "sue@email.com"], "cf_Url": "www.yext.com"}`, &CustomLocationEntity{LocationEntity: LocationEntity{Emails: Strings([]string{"bob@email.com", "sue@email.com"})}, CustomEntity: CustomEntity{CFUrl: String("www.yext.com")}}},
		{`{"cf_TextList": ["a", "b", "c"]}`, &CustomLocationEntity{CustomEntity: CustomEntity{CFTextList: Strings([]string{"a", "b", "c"})}}},
		{`{"address":{"city":""}}`, &CustomLocationEntity{LocationEntity: LocationEntity{Address: &Address{City: String("")}}}},
		{`{"languages":[]}`, &CustomLocationEntity{LocationEntity: LocationEntity{Languages: &[]string{}}}},
		{`{"languages":["English"]}`, &CustomLocationEntity{LocationEntity: LocationEntity{Languages: &[]string{"English"}}}},
		{`{"hours":null}`, &CustomLocationEntity{LocationEntity: LocationEntity{Hours: NullHours()}}},
		{`{"hours":{"monday":null,"tuesday":null,"wednesday":null,"thursday":null,"friday":null,"saturday":null,"sunday":null}}`, &CustomLocationEntity{LocationEntity: LocationEntity{Hours: NullableHours(&Hours{Monday: NullDayHours(), Tuesday: NullDayHours(), Wednesday: NullDayHours(), Thursday: NullDayHours(), Friday: NullDayHours(), Saturday: NullDayHours(), Sunday: NullDayHours()})}}},
		{`{"cf_Url":""}`, &CustomLocationEntity{CustomEntity: CustomEntity{CFUrl: String("")}}},
		{`{"cf_Url": "www.yext.com"}`, &CustomLocationEntity{CustomEntity: CustomEntity{CFUrl: String("www.yext.com")}}},
		{`{"cf_TextList":[]}`, &CustomLocationEntity{CustomEntity: CustomEntity{CFTextList: &[]string{}}}},
		{`{"cf_SingleOption":null}`, &CustomLocationEntity{CustomEntity: CustomEntity{CFSingleOption: NullString()}}},
		{`{"cf_MultiOption":[]}`, &CustomLocationEntity{CustomEntity: CustomEntity{CFMultiOption: ToUnorderedStrings([]string{})}}},
		{`{"cf_Date":""}`, &CustomLocationEntity{CustomEntity: CustomEntity{CFDate: String("")}}},
		{`{"cf_Video":null}`, &CustomLocationEntity{CustomEntity: CustomEntity{CFVideo: NullVideo()}}},
		{`{"cf_Photo":null}`, &CustomLocationEntity{CustomEntity: CustomEntity{CFPhoto: NullPhoto()}}},
		{`{"cf_Gallery":[]}`, &CustomLocationEntity{CustomEntity: CustomEntity{CFGallery: &[]Photo{}}}},
		{`{"cf_Videos":[]}`, &CustomLocationEntity{CustomEntity: CustomEntity{CFVideos: &[]Video{}}}},
		{`{"cf_DailyTimes":null}`, &CustomLocationEntity{CustomEntity: CustomEntity{CFDailyTimes: NullDailyTimes()}}},
		{`{"cf_Hours":null}`, &CustomLocationEntity{CustomEntity: CustomEntity{CFHours: NullHours()}}},
		{`{"cf_YesNo":null}`, &CustomLocationEntity{CustomEntity: CustomEntity{CFYesNo: NullBool()}}},
		{`{"name":"Hello","cf_YesNo":null}`, &CustomLocationEntity{LocationEntity: LocationEntity{Name: String("Hello")}, CustomEntity: CustomEntity{CFYesNo: NullBool()}}},
		{`{"name":"","cf_YesNo":null}`, &CustomLocationEntity{LocationEntity: LocationEntity{Name: String("")}, CustomEntity: CustomEntity{CFYesNo: NullBool()}}},
		{`{"cf_Text":""}`, &CustomLocationEntity{CustomEntity: CustomEntity{CFText: String("")}}},
	}

	for _, test := range tests {
		v := &CustomLocationEntity{}
		if err := json.Unmarshal([]byte(test.json), v); err != nil {
			t.Error("Unable to deserialize", test.json, "from JSON:", err)
		} else if !reflect.DeepEqual(v, test.want) {
			t.Errorf("json.Unmarshal()\nGot:      %s\nExpected: %s", v, test.want)
		}
	}
}

func TestSetLabels(t *testing.T) {
	var tests = []struct {
		Destination *LocationEntity
		Source      *LocationEntity
		Want        *LocationEntity
	}{
		{
			Destination: &LocationEntity{
				BaseEntity: BaseEntity{
					Meta: &EntityMeta{
						Labels: nil,
					},
				},
			},
			Source: &LocationEntity{
				BaseEntity: BaseEntity{
					Meta: &EntityMeta{
						Labels: nil,
					},
				},
			},
			Want: &LocationEntity{
				BaseEntity: BaseEntity{
					Meta: &EntityMeta{
						Labels: nil,
					},
				},
			},
		},
	}
	for _, test := range tests {
		test.Destination.SetLabelsWithUnorderedStrings(test.Source.GetLabels())
		if _, isDiff, _ := Diff(test.Want, test.Destination); isDiff {
			t.Errorf("SetLabelsWithUnorderedStrings: Wanted %v Got %v", test.Want, test.Destination)
		}
	}

	var labelTests = []struct {
		Entity *LocationEntity
		Labels []string
		Want   *LocationEntity
	}{
		{
			Entity: &LocationEntity{
				BaseEntity: BaseEntity{
					Meta: &EntityMeta{
						Labels: nil,
					},
				},
			},
			Labels: []string{},
			Want: &LocationEntity{
				BaseEntity: BaseEntity{
					Meta: &EntityMeta{
						Labels: ToUnorderedStrings([]string{}),
					},
				},
			},
		},
		{
			Entity: &LocationEntity{
				BaseEntity: BaseEntity{
					Meta: &EntityMeta{
						Labels: nil,
					},
				},
			},
			Labels: nil,
			Want: &LocationEntity{
				BaseEntity: BaseEntity{
					Meta: &EntityMeta{
						Labels: nil,
					},
				},
			},
		},
	}
	for _, test := range labelTests {
		test.Entity.SetLabels(test.Labels)
		if _, isDiff, _ := Diff(test.Want, test.Entity); isDiff {
			t.Errorf("SetLabels: Wanted %v Got %v", test.Want, test.Entity)
		}
	}
}

func TestEntitySampleJSONResponseDeserialization(t *testing.T) {
	entityService := EntityService{
		Registry: &EntityRegistry{},
	}
	entityService.RegisterEntity("location", &CustomLocationEntity{})
	mapOfStringToInterface := make(map[string]interface{})
	err := json.Unmarshal([]byte(sampleEntityJSON), &mapOfStringToInterface)
	if err != nil {
		t.Errorf("Unable to unmarshal sample entity JSON. Err: %s", err)
	}
	if _, err := entityService.ToEntityType(mapOfStringToInterface); err != nil {
		t.Errorf("Unable to convert JSON to entity type. Err: %s", err)
	}
}

var sampleEntityJSON = `{
  "additionalHoursText": "Some additional hours text",
  "address": {
     "line1": "7900 Westpark",
     "city": "McLean",
     "region": "VA",
     "postalCode": "22102",
     "extraDescription": "Galleria Shopping Center"
  },
  "addressHidden": false,
  "description": "This is my description",
  "hours": {
     "monday": {
			 "openIntervals": [
	         {
	             "start": "09:00",
	             "end": "17:00"
	         }
				]
     },
     "tuesday": {
			 "openIntervals": [
	         {
	             "start": "09:00",
	             "end": "17:00"
	         }
				]
     },
     "wednesday": {
			 	"openIntervals": [
	         {
	             "start": "09:00",
	             "end": "17:00"
	         }
				]
     },
     "thursday": {
		 		"openIntervals": [
					 {
	             "start": "09:00",
	             "end": "17:00"
	         }
				]
     },
     "friday": {
			 "openIntervals": [
	         {
	             "start": "09:00",
	             "end": "17:00"
	         }
				]
     },
		 "saturday": {
			 "isClosed": true
     },
     "sunday": {
			  "openIntervals": [
	         {
	             "start": "00:00",
	             "end": "23:59"
	         }
				]
     },
     "holidayHours": [
         {
             "date": "2018-12-25",
						 "isClosed": true,
             "isRegularHours": false
         }
     ]
  },
  "name": "Yext Consulting",
  "cf_Hours": {
     "monday": {
			 "openIntervals": [
	         {
	             "start": "09:00",
	             "end": "17:00"
	         }
				]
     },
     "tuesday": {
			 "openIntervals": [
         {
             "start": "09:00",
             "end": "17:00"
         }
				]
     },
     "wednesday": {
			 "openIntervals": [
         {
             "start": "09:00",
             "end": "17:00"
         }
		 	 ]
     },
     "thursday": {
			 "openIntervals": [
	         {
	             "start": "09:00",
	             "end": "17:00"
	         }
				]
     },
     "friday": {
			 	"openIntervals": [
	         {
	             "start": "09:00",
	             "end": "14:00"
	         },
	         {
	             "start": "15:00",
	             "end": "17:00"
	         }
				]
     },
     "saturday": {
			 	"openIntervals": [
	         {
	             "start": "00:00",
	             "end": "23:59"
	         }
				]
     },
     "holidayHours": [
         {
             "date": "2018-10-13",
             "hours": [
                 {
                     "start": "10:00",
                     "end": "16:00"
                 }
             ],
             "isRegularHours": false
         }
     ]
  },
  "cf_DailyTimes": {
     "monday": "09:00",
     "tuesday": "09:00",
     "wednesday": "09:00",
     "thursday": "09:00",
     "friday": "09:00",
     "saturday": "09:00"
  },
  "cf_Date": "2018-09-28",
  "cf_Gallery": [
     {
         "image": {
             "url": "http://a.mktgcdn.com/p/CCCUglaMWv5i6Ede4KJuEFttpN416TTKGppUz-vcBqI/920x640.jpg",
             "width": 920,
             "height": 640,
             "derivatives": [
                 {
                     "url": "http://a.mktgcdn.com/p/CCCUglaMWv5i6Ede4KJuEFttpN416TTKGppUz-vcBqI/619x430.jpg",
                     "width": 619,
                     "height": 430
                 },
                 {
                     "url": "http://a.mktgcdn.com/p/CCCUglaMWv5i6Ede4KJuEFttpN416TTKGppUz-vcBqI/600x417.jpg",
                     "width": 600,
                     "height": 417
                 },
                 {
                     "url": "http://a.mktgcdn.com/p/CCCUglaMWv5i6Ede4KJuEFttpN416TTKGppUz-vcBqI/196x136.jpg",
                     "width": 196,
                     "height": 136
                 }
             ]
         },
         "description": "Corgi Puppy"
     },
     {
         "image": {
             "url": "http://a.mktgcdn.com/p/T88YcpE1osKn6higLxP5W0yYr3PU7iQJueQ29nsCFkA/103x103.jpg",
             "width": 103,
             "height": 103
         },
         "description": "[[name]]"
     }
  ],
  "cf_MultiOpion": [
     "4118",
     "8938"
  ],
  "cf_Photo": {
     "image": {
         "url": "http://a.mktgcdn.com/p/aYN_mOHTqqLfizsnR7b17ldLAPe5P3vX--wBkpTHx14/590x350.jpg",
         "width": 590,
         "height": 350,
         "derivatives": [
             {
                 "url": "http://a.mktgcdn.com/p/aYN_mOHTqqLfizsnR7b17ldLAPe5P3vX--wBkpTHx14/196x116.jpg",
                 "width": 196,
                 "height": 116
             }
         ]
     }
  },
  "cf_SingleOption": "3035",
  "cf_TextList": [
     "List Item 1",
     "List Item 2",
     "List Item 3"
  ],
  "cf_Videos": [
     {
         "video": {
             "url": "http://www.youtube.com/watch?v=TYRDgd3Tb44"
         },
         "description": "video description"
     }
  ],
  "cf_Video": {
      "video": {
          "url": "http://www.youtube.com/watch?v=fC3Cthm0HFU"
      },
      "description": "video description"
  },
  "cf_Url": "http://yext.com/careers",
  "emails": [
     "cdworak@yext.com"
  ],
  "featuredMessage": {
     "description": "This is my featured message",
     "url":"http://www.bestbuy.com/site/electronics/black-friday/pcmcat225600050002.c?id\u003dpcmcat225600050002\u0026ref\u003dNS\u0026loc\u003dns100"
  },
  "isoRegionCode": "VA",
  "mainPhone": "+18888888888",
  "timezone": "America/New_York",
  "websiteUrl": {
     "url": "http://yext.com",
     "displayUrl": "http://yext.com",
     "preferDisplayUrl": false
  },
  "yextDisplayCoordinate": {
     "latitude": 38.92475,
     "longitude": -77.21718
  },
  "yextRoutableCoordinate": {
     "latitude": 38.9243983751914,
     "longitude": -77.2178385786886
  },
	"googleAttributes": {
		"wi_fi": ["free_wi_fi"],
		"welcomes_dogs": ["true"]
	},
  "meta": {
     "accountId": "3549951188342570541",
     "uid": "b3JxON",
     "id": "CTG",
     "categoryIds": [
         "668"
     ],
     "folderId": "0",
     "language": "en",
     "countryCode": "US",
     "entityType": "location"
  }
}`

func TestRawEntityIsZeroValue(t *testing.T) {
	var tests = []struct {
		Name string
		Raw  *RawEntity
		Want bool
	}{
		{
			Name: "Zero - Empty entity",
			Raw:  &RawEntity{},
			Want: true,
		},
		{
			Name: "Zero - Empty single field",
			Raw: &RawEntity{
				"name": nil,
			},
			Want: true,
		},
		{
			Name: "Zero - Empty nested map field",
			Raw: &RawEntity{
				"address": map[string]interface{}{
					"line1": nil,
				},
			},
			Want: true,
		},
		{
			Name: "Zero - Empty nested slice subfield",
			Raw: &RawEntity{
				"hours": map[string]interface{}{
					"monday": map[string]interface{}{
						"openIntervals": []interface{}{
							map[string]interface{}{
								"start": nil,
								"end":   nil,
							},
						},
					},
				},
			},
			Want: true,
		},
		{
			Name: "Zero - double empty nested slice subfield",
			Raw: &RawEntity{
				"providerData": []interface{}{
					[]interface{}{
						map[string]interface{}{
							"name": nil,
						},
					},
					map[string]interface{}{
						"name": nil,
						"id":   nil,
					},
				},
			},
			Want: true,
		},
		{
			Name: "Not Zero - valid single field",
			Raw: &RawEntity{
				"name": "Name",
			},
			Want: false,
		},
		{
			Name: "Not Zero - valid single field (but it has zero value - empty string)",
			Raw: &RawEntity{
				"name": "",
			},
			Want: false,
		},
		{
			Name: "Not Zero - valid single field (but it has zero value - false)",
			Raw: &RawEntity{
				"closed": false,
			},
			Want: false,
		},
		{
			Name: "Not Zero - valid nested slice subfield",
			Raw: &RawEntity{
				"hours": map[string]interface{}{
					"monday": map[string]interface{}{
						"openIntervals": []interface{}{
							map[string]interface{}{
								"start": "10:00",
								"end":   "18:00",
							},
						},
					},
				},
			},
			Want: false,
		},
		{
			Name: "Not Zero - valid slice subfield, even with one invalid element",
			Raw: &RawEntity{
				"providerData": []interface{}{
					map[string]interface{}{
						"name": nil,
						"id":   nil,
					},
					map[string]interface{}{
						"name": "Name",
					},
				},
			},
			Want: false,
		},
		{
			Name: "Not Zero - double nested slice subfield, even with one invalid element in each slice",
			Raw: &RawEntity{
				"providerData": []interface{}{
					[]interface{}{
						map[string]interface{}{
							"name": false,
						},
						map[string]interface{}{
							"name": nil,
						},
					},
					map[string]interface{}{
						"name": nil,
						"id":   nil,
					},
				},
			},
			Want: false,
		},
		{
			Name: "Not Zero - double nested slice subfield, even with one invalid element",
			Raw: &RawEntity{
				"providerData": []interface{}{
					[]interface{}{
						map[string]interface{}{
							"name": nil,
						},
					},
					map[string]interface{}{
						"name": "",
					},
				},
			},
			Want: false,
		},
		{
			Name: "Disney ETL Test that was failing",
			Raw: &RawEntity{
				"c_oSID":         "80007922",
				"c_propertyType": "land",
				"c_storeName":    "frontierland",
				"description":    "",
				"hours":          map[string]interface{}{},
				"meta": map[string]interface{}{
					"entityType": "location",
					"id":         "80007922;entityType=land",
					"labels": []interface{}{
						"83533",
					},
				},
				"name":         "Frontierland",
				"photoGallery": []interface{}{},
			},
			Want: false,
		},
	}

	for _, test := range tests {
		got := test.Raw.IsZeroValue()
		if got != test.Want {
			t.Errorf("Got: %v, Wanted: %v", got, test.Want)
		}
	}
}

func TestGetValue(t *testing.T) {
	var tests = []struct {
		Raw  *RawEntity
		Keys []string
		Want interface{}
	}{
		{
			Raw: &RawEntity{
				"meta": map[string]interface{}{
					"entityType": "location",
				},
			},
			Keys: []string{"meta", "entityType"},
			Want: "location",
		},
		{
			Raw: &RawEntity{
				"meta": map[string]interface{}{
					"entityType": "location",
				},
			},
			Keys: []string{"meta", "id"},
			Want: nil,
		},
	}

	for _, test := range tests {
		got := test.Raw.GetValue(test.Keys)
		if got != test.Want {
			t.Errorf("Got: %v, Wanted: %v", got, test.Want)
		}
	}
}

func TestSetValue(t *testing.T) {
	var tests = []struct {
		Raw   *RawEntity
		Keys  []string
		Value interface{}
		Want  *RawEntity
	}{
		{
			Raw: &RawEntity{
				"meta": map[string]interface{}{
					"entityType": "location",
				},
			},
			Keys:  []string{"meta", "entityType"},
			Value: "hotel",
			Want: &RawEntity{
				"meta": map[string]interface{}{
					"entityType": "hotel",
				},
			},
		},
		{
			Raw: &RawEntity{
				"meta": map[string]interface{}{
					"entityType": "location",
				},
			},
			Keys:  []string{"meta", "id"},
			Value: "1234",
			Want: &RawEntity{
				"meta": map[string]interface{}{
					"entityType": "location",
					"id":         "1234",
				},
			},
		},
		{
			Raw:   &RawEntity{},
			Keys:  []string{"meta", "id"},
			Value: "1234",
			Want: &RawEntity{
				"meta": map[string]interface{}{
					"id": "1234",
				},
			},
		},
	}

	for _, test := range tests {
		err := test.Raw.SetValue(test.Keys, test.Value)
		if err != nil {
			t.Errorf("Got err: %s", err)
		}
		if delta, isDiff := RawEntityDiff(*test.Raw, *test.Want, false, false, false, false); isDiff {
			t.Errorf("Got: %v, Wanted: %v, Delta: %v", test.Raw, test.Want, delta)
		}
	}
}
