package yext

import "encoding/json"

// TODO: delete this file when the new endpoints become available

var listResponse = `{
    "count": 4,
    "entities": [
      {
        "id": "US5523",
        "uid": "0Za9W",
        "timestamp": 1530543889575,
        "accountId": "648517391220866580",
        "locationName": "Safelite AutoGlass",
        "address": "2124 W 2nd St",
        "city": "Grand Island",
        "state": "NE",
        "zip": "68803",
        "countryCode": "US",
        "language": "en",
        "suppressAddress": false,
        "phone": "8888432798",
        "isPhoneTracked": false,
        "localPhone": "3083828689",
        "alternatePhone": "8008002727",
        "categoryIds": [
            "28",
            "2040",
            "26",
            "27",
            "588",
            "524",
            "1325579"
        ],
        "entityType": "LOCATION"
      },
      {
        "id": "Entity2",
        "uid": "0Za9W",
        "timestamp": 1530543889575,
        "accountId": "648517391220866580",
        "locationName": "Safelite AutoGlass",
        "address": "2124 W 2nd St",
        "city": "Grand Island",
        "state": "NE",
        "zip": "68803",
        "countryCode": "US",
        "language": "en",
        "suppressAddress": false,
        "phone": "8888432798",
        "isPhoneTracked": false,
        "localPhone": "3083828689",
        "alternatePhone": "8008002727",
        "categoryIds": [
            "28",
            "2040",
            "26",
            "27",
            "588",
            "524",
            "1325579"
        ],
        "entityType": "LOCATION"
      },
      {
        "id": "eventId",
        "name": "event name",
        "description": "event description",
        "entityType": "EVENT"
      },
      {
        "id": "12345",
        "firstName": "Catherine",
        "nickname": "CTG",
        "entityType": "CONSULTING_ENGINEER"
      }
    ]
  }
`

// Used to fake a JSON response for Enitites:List
func fakeEntityListResponse() (*EntityListResponse, error) {
	byteValue := []byte(listResponse)
	var resp *EntityListResponse
	err := json.Unmarshal(byteValue, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
