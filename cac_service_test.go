package yext

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

type DataType struct {
	path string
	data []byte
}

func TestConfigFieldListAll(t *testing.T) {
	tests := []struct {
		data []DataType
		want []*CustomField
	}{
		{
			data: []DataType{
				{
					path: "/accounts/me/config/resourcenames/km/field-eligibility-group",
					data: []byte(`{"meta": {"errors": [], "uuid": ""}, "response": ["entityTypeA.default"]}`),
				},
				{
					path: "/accounts/me/config/resources/km/field-eligibility-group/entityTypeA.default",
					data: []byte(`{"meta": {"errors": [], "uuid": ""}, "response": {"$id":"entityTypeA.default", "entityType": "entityTypeA", "fields": []}}`),
				},
				{
					path: "/accounts/me/config/resourcenames/km/field",
					data: []byte(`{"meta": {"errors": [], "uuid": ""}, "response": ["fieldA"]}`),
				},
				{
					path: "/accounts/me/config/resources/km/field/fieldA",
					data: []byte(`{"meta": {"errors": [], "uuid": ""}, "response": {"$id": "fieldA", "typeId": "boolean"}}`),
				},
			},
			want: []*CustomField{
				{
					Id:                 String("fieldA"),
					Type:               CUSTOMFIELDTYPE_YESNO,
					EntityAvailability: nil,
				},
			},
		},
		{
			data: []DataType{
				{
					path: "/accounts/me/config/resourcenames/km/field-eligibility-group",
					data: []byte(`{"meta": {"errors": [], "uuid": ""}, "response": ["entityTypeA.default"]}`),
				},
				{
					path: "/accounts/me/config/resources/km/field-eligibility-group/entityTypeA.default",
					data: []byte(`{"meta": {"errors": [], "uuid": ""}, "response": {"$id":"entityTypeA.default", "entityType": "entityTypeA", "fields": ["fieldA"]}}`),
				},
				{
					path: "/accounts/me/config/resourcenames/km/field",
					data: []byte(`{"meta": {"errors": [], "uuid": ""}, "response": ["fieldA"]}`),
				},
				{
					path: "/accounts/me/config/resources/km/field/fieldA",
					data: []byte(`{"meta": {"errors": [], "uuid": ""}, "response": {"$id": "fieldA", "typeId": "boolean"}}`),
				},
			},
			want: []*CustomField{
				{
					Id:                 String("fieldA"),
					Type:               CUSTOMFIELDTYPE_YESNO,
					EntityAvailability: []EntityType{"entityTypeA"},
				},
			},
		},
		{
			data: []DataType{
				{
					path: "/accounts/me/config/resourcenames/km/field-eligibility-group",
					data: []byte(`{"meta": {"errors": [], "uuid": ""}, "response": ["entityTypeA.default", "entityTypeB.default"]}`),
				},
				{
					path: "/accounts/me/config/resources/km/field-eligibility-group/entityTypeA.default",
					data: []byte(`{"meta": {"errors": [], "uuid": ""}, "response": {"$id":"entityTypeA.default", "entityType": "entityTypeA", "fields": ["fieldA"]}}`),
				},
				{
					path: "/accounts/me/config/resources/km/field-eligibility-group/entityTypeB.default",
					data: []byte(`{"meta": {"errors": [], "uuid": ""}, "response": {"$id":"entityTypeB.default", "entityType": "entityTypeB", "fields": ["fieldA", "fieldB"]}}`),
				},
				{
					path: "/accounts/me/config/resourcenames/km/field",
					data: []byte(`{"meta": {"errors": [], "uuid": ""}, "response": ["fieldA", "fieldB"]}`),
				},
				{
					path: "/accounts/me/config/resources/km/field/fieldA",
					data: []byte(`{"meta": {"errors": [], "uuid": ""}, "response": {"$id": "fieldA", "typeId": "boolean"}}`),
				},
				{
					path: "/accounts/me/config/resources/km/field/fieldB",
					data: []byte(`{"meta": {"errors": [], "uuid": ""}, "response": {"$id": "fieldB", "typeId": "decimal"}}`),
				},
			},
			want: []*CustomField{
				{
					Id:                 String("fieldA"),
					Type:               CUSTOMFIELDTYPE_YESNO,
					EntityAvailability: []EntityType{"entityTypeA", "entityTypeB"},
				},
				{
					Id:                 String("fieldB"),
					Type:               CUSTOMFIELDTYPE_NUMBER,
					EntityAvailability: []EntityType{"entityTypeB"},
				},
			},
		},
		{
			data: []DataType{
				{
					path: "/accounts/me/config/resourcenames/km/field-eligibility-group",
					data: []byte(`{"meta": {"errors": [], "uuid": ""}, "response": ["entityTypeA.default"]}`),
				},
				{
					path: "/accounts/me/config/resources/km/field-eligibility-group/entityTypeA.default",
					data: []byte(`{"meta": {"errors": [], "uuid": ""}, "response": {"$id":"entityTypeA.default", "entityType": "entityTypeA", "fields": ["fieldA"]}}`),
				},
				{
					path: "/accounts/me/config/resourcenames/km/field",
					data: []byte(`{"meta": {"errors": [], "uuid": ""}, "response": ["fieldA"]}`),
				},
				{
					path: "/accounts/me/config/resources/km/field/fieldA",
					data: []byte(
						`{
							"meta": {"errors": [], "uuid": ""}, 
							"response": {
								"$id": "fieldA", 
								"typeId": "list", 
								"type": {
									"listType": {
										"typeId": "option", 
										"type": {
											"optionType": {
												"option": [
													{"displayName": "Option A", "textValue": "optionA", "displayNameTranslation": [{"localeCode": "fr", "value": "Option A FR"}]},
													{"displayName": "Option B", "textValue": "optionB", "displayNameTranslation": [{"localeCode": "fr", "value": "Option B FR"}]}
												]
											}
										}
									}	
								}
							}
						}`),
				},
			},
			want: []*CustomField{
				{
					Id:                 String("fieldA"),
					Type:               CUSTOMFIELDTYPE_MULTIOPTION,
					EntityAvailability: []EntityType{"entityTypeA"},
					Options: []CustomFieldOption{
						{
							Key:   "optionA",
							Value: "Option A",
							Translations: []Translation{
								{
									LanguageCode: "fr",
									Value:        "Option A FR",
								},
							},
						},
						{
							Key:   "optionB",
							Value: "Option B",
							Translations: []Translation{
								{
									LanguageCode: "fr",
									Value:        "Option B FR",
								},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		setup()
		client.WithConfigAPIForCFs()
		reqs := 0
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			reqs++
			if reqs > len(test.data) {
				t.Errorf("Too many requests sent to config endpoint - got %d, expected %d", reqs, len(test.data))
			}
			for _, d := range test.data {
				if r.URL.Path == d.path {
					w.Write(d.data)
					return
				}
			}
			t.Errorf("Unexpected request to %s", r.URL.Path)
		})

		got, _ := client.CustomFieldService.ListAll()
		if !reflect.DeepEqual(got, test.want) {
			var (
				gotResults  []string
				wantResults []string
			)
			for _, g := range got {
				gotResults = append(gotResults, fmt.Sprintf("%+v", g))
			}
			for _, w := range test.want {
				wantResults = append(wantResults, fmt.Sprintf("%+v", w))
			}
			t.Errorf("CustomFieldService.ListAll() = %s, want %s", gotResults, wantResults)
		}
		teardown()
	}
}

func TestConfigFieldCreate(t *testing.T) {
	tests := []struct {
		count int
		cf    *CustomField
	}{
		{
			count: 1,
			cf: &CustomField{
				Id: String("fieldA"),
			},
		},
		{
			count: 3,
			cf: &CustomField{
				Id:                 String("fieldA"),
				EntityAvailability: []EntityType{"entityTypeA"},
			},
		},
		{
			count: 5,
			cf: &CustomField{
				Id:                 String("fieldA"),
				EntityAvailability: []EntityType{"entityTypeA", "entityTypeB"},
			},
		},
	}

	for _, test := range tests {
		setup()
		client.WithConfigAPIForCFs()
		reqs := 0
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			reqs++
			if reqs > test.count {
				t.Errorf("Too many requests sent to config endpoint - got %d, expected %d", reqs, test.count)
			}
			v := &mockResponse{}
			if strings.Contains(r.URL.Path, "config/resources/km/field-eligibility-group") {
				v = &mockResponse{Response: ConfigFieldEligibilityGroup{}}
			}
			data, _ := json.Marshal(v)
			w.Write(data)
		})

		_, err := client.CustomFieldService.Create(test.cf)
		if err != nil {
			t.Errorf("CustomFieldService.Create() returned error: %v", err)
		}
		if reqs != test.count {
			t.Errorf("Too few requests sent to config endpoint - got %d, expected %d", reqs, test.count)
		}
		teardown()
	}
}

func TestConfigFieldEdit(t *testing.T) {
	tests := []struct {
		count int
		cf    *CustomField
	}{
		{
			count: 1,
			cf: &CustomField{
				Id: String("fieldA"),
			},
		},
		{
			count: 2,
			cf: &CustomField{
				Id:   String("fieldA"),
				Type: CUSTOMFIELDTYPE_SINGLEOPTION,
			},
		},
		{
			count: 3,
			cf: &CustomField{
				Id:                 String("fieldA"),
				EntityAvailability: []EntityType{"entityTypeA"},
			},
		},
		{
			count: 5,
			cf: &CustomField{
				Id:                 String("fieldA"),
				EntityAvailability: []EntityType{"entityTypeA", "entityTypeB"},
			},
		},
		{
			count: 2,
			cf: &CustomField{
				Id:                 String("fieldC"),
				EntityAvailability: []EntityType{"entityTypeA"},
			},
		},
	}

	for _, test := range tests {
		setup()
		client.WithConfigAPIForCFs()
		reqs := 0
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			reqs++
			if reqs > test.count {
				t.Errorf("Too many requests sent to config endpoint - got %d, expected %d", reqs, test.count)
			}
			v := &mockResponse{}
			if strings.Contains(r.URL.Path, "config/resources/km/field-eligibility-group") {
				v = &mockResponse{Response: ConfigFieldEligibilityGroup{Fields: []string{"fieldC"}}}
			} else if strings.Contains(r.URL.Path, "config/resources/km/field") {
				v = &mockResponse{Response: ConfigField{Type: &ConfigType{OptionType: &ConfigOptionType{}}}}
			}
			data, _ := json.Marshal(v)
			w.Write(data)
		})

		_, err := client.CustomFieldService.Edit(test.cf)
		if err != nil {
			t.Errorf("CustomFieldService.Create() returned error: %v", err)
		}
		if reqs != test.count {
			t.Errorf("Too few requests sent to config endpoint - got %d, expected %d", reqs, test.count)
		}
		teardown()
	}
}

func TestCustomFieldOptionId(t *testing.T) {
	tests := []struct {
		data        []DataType
		fieldName   string
		optionValue string
		wantId      string
	}{
		{
			data: []DataType{
				{
					path: "/accounts/me/config/resourcenames/km/field",
					data: []byte(`{"meta": {"errors": [], "uuid": ""}, "response": ["fieldA"]}`),
				},
				{
					path: "/accounts/me/config/resources/km/field/fieldA",
					data: []byte(
						`{
							"meta": {"errors": [], "uuid": ""}, 
							"response": {
								"$id": "fieldA", 
								"displayName": "Field A",
								"typeId": "option", 
								"type": {
									"optionType": {
										"option": [
											{"displayName": "Option A", "textValue": "optionA", "displayNameTranslation": [{"localeCode": "fr", "value": "Option A FR"}]},
											{"displayName": "Option B", "textValue": "optionB", "displayNameTranslation": [{"localeCode": "fr", "value": "Option B FR"}]}
										]
									}	
								}
							}
						}`),
				},
			},
			fieldName:   "Field A",
			optionValue: "Option A",
			wantId:      "optionA",
		},
		{
			data: []DataType{
				{
					path: "/accounts/me/config/resourcenames/km/field",
					data: []byte(`{"meta": {"errors": [], "uuid": ""}, "response": ["fieldA", "fieldB"]}`),
				},
				{
					path: "/accounts/me/config/resources/km/field/fieldA",
					data: []byte(
						`{
							"meta": {"errors": [], "uuid": ""}, 
							"response": {
								"$id": "fieldA", 
								"displayName": "Field A",
								"typeId": "option", 
								"type": {
									"optionType": {
										"option": [
											{"displayName": "Option A", "textValue": "optionA", "displayNameTranslation": [{"localeCode": "fr", "value": "Option A FR"}]},
											{"displayName": "Option B", "textValue": "optionB", "displayNameTranslation": [{"localeCode": "fr", "value": "Option B FR"}]}
										]
									}	
								}
							}
						}`),
				},
				{
					path: "/accounts/me/config/resources/km/field/fieldB",
					data: []byte(
						`{
							"meta": {"errors": [], "uuid": ""}, 
							"response": {
								"$id": "fieldB", 
								"displayName": "Field B",
								"typeId": "list", 
								"type": {
									"listType": {
										"typeId": "option", 
										"type": {
											"optionType": {
												"option": [
													{"displayName": "Option C", "textValue": "optionC", "displayNameTranslation": [{"localeCode": "fr", "value": "Option C FR"}]},
													{"displayName": "Option D", "textValue": "optionD", "displayNameTranslation": [{"localeCode": "fr", "value": "Option D FR"}]}
												]
											}
										}
									}	
								}
							}
						}`),
				},
			},
			fieldName:   "Field B",
			optionValue: "Option D",
			wantId:      "optionD",
		},
	}
	for _, test := range tests {
		setup()
		client.WithConfigAPIForCFs()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			for _, d := range test.data {
				if r.URL.Path == d.path {
					w.Write(d.data)
					return
				}
			}
			v := &mockResponse{}
			if strings.Contains(r.URL.Path, "config/resources/km/field-eligibility-group") {
				v = &mockResponse{Response: ConfigFieldEligibilityGroup{}}
			} else if strings.Contains(r.URL.Path, "config/resourcenames/km/field-eligibility-group") {
				v = &mockResponse{Response: []string{}}
			}
			data, _ := json.Marshal(v)
			w.Write(data)
		})

		client.CustomFieldService.MustCacheCustomFields()

		gotId, err := client.CustomFieldService.CustomFieldManager.CustomFieldOptionId(test.fieldName, test.optionValue)
		if err != nil {
			t.Errorf("CustomFieldService.CustomFieldOptionId() returned error: %v", err)
		}
		if gotId != test.wantId {
			t.Errorf("CustomFieldService.CustomFieldOptionId() = %v, want %v", gotId, test.wantId)
		}
		teardown()
	}
}
