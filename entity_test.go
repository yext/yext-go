package yext

import (
	"encoding/json"
	"reflect"
	"testing"
)

type CustomLocationEntity struct {
	LocationEntity
	CFHours        *Hours            `json:"cf_Hours,omitempty"`
	CFUrl          *string           `json:"cf_Url,omitempty"`
	CFDailyTimes   *DailyTimes       `json:"cf_DailyTimes,omitempty"`
	CFTextList     *[]string         `json:"cf_TextList,omitempty"`
	CFGallery      []*Photo          `json:"cf_Gallery,omitempty"`
	CFPhoto        *Photo            `json:"cf_Photo,omitempty"`
	CFVideos       []*Video          `json:"cf_Videos,omitempty"`
	CFVideo        *Video            `json:"cf_Video,omitempty"`
	CFDate         *Date             `json:"cf_Date,omitempty"`
	CFSingleOption *string           `json:"cf_SingleOtpion,omitempty"`
	CFMultiOption  *UnorderedStrings `json:"cf_MultiOption,omitempty"`
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
		{&CustomLocationEntity{LocationEntity: LocationEntity{Address: &Address{City: nil}}}, `{"address":{}}`}, // TODO: verify this is correct
		{&CustomLocationEntity{LocationEntity: LocationEntity{Address: &Address{City: String("")}}}, `{"address":{"city":""}}`},
		{&CustomLocationEntity{LocationEntity: LocationEntity{Languages: nil}}, `{}`},
		{&CustomLocationEntity{LocationEntity: LocationEntity{Languages: nil}}, `{}`},
		{&CustomLocationEntity{LocationEntity: LocationEntity{Languages: &[]string{}}}, `{"languages":[]}`},
		{&CustomLocationEntity{LocationEntity: LocationEntity{Languages: &[]string{"English"}}}, `{"languages":["English"]}`},
		{&CustomLocationEntity{LocationEntity: LocationEntity{Hours: nil}}, `{}`},
		{&CustomLocationEntity{LocationEntity: LocationEntity{Hours: &Hours{}}}, `{"hours":{}}`},
		{&CustomLocationEntity{CFUrl: String("")}, `{"cf_Url":""}`},
		{&CustomLocationEntity{CFUrl: nil}, `{}`},
		{&CustomLocationEntity{CFTextList: &[]string{}}, `{"cf_TextList":[]}`},
		{&CustomLocationEntity{CFTextList: nil}, `{}`},
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
		{`{"emails": ["mhupman@yext.com", "bmcginnis@yext.com"]}`, &CustomLocationEntity{LocationEntity: LocationEntity{Emails: Strings([]string{"mhupman@yext.com", "bmcginnis@yext.com"})}}},
		{`{"cf_Url": "www.yext.com"}`, &CustomLocationEntity{CFUrl: String("www.yext.com")}},
		{`{"cf_TextList": ["a", "b", "c"]}`, &CustomLocationEntity{CFTextList: Strings([]string{"a", "b", "c"})}},
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

func TestEntitySampleJSONResponseDeserialization(t *testing.T) {
	entityService := EntityService{
		registry: make(Registry),
	}
	entityService.RegisterEntity("LOCATION", &CustomLocationEntity{})
	mapOfStringToInterface := make(map[string]interface{})
	err := json.Unmarshal([]byte(sampleEntityJSON), &mapOfStringToInterface)
	if err != nil {
		t.Errorf("Unable to unmarshal sample entity JSON. Err: %s", err)
	}
	if _, err := entityService.toEntityType(mapOfStringToInterface); err != nil {
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
     "entityType": "LOCATION"
  }
}`
