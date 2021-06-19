package yext

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
	"testing"
)

func makeCustomFields(n int) []*CustomField {
	var cfs []*CustomField

	for i := 0; i < n; i++ {
		new := &CustomField{Id: String(strconv.Itoa(i))}
		cfs = append(cfs, new)
	}

	return cfs
}

func TestCustomFieldListAll(t *testing.T) {
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

func TestCustomFieldListMismatchCount(t *testing.T) {
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

func TestMustCacheCustomFields(t *testing.T) {
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

var cfm = &CustomFieldManager{
	CustomFields: []*CustomField{
		&CustomField{
			Name: "My Favorite Colors",
			Type: CUSTOMFIELDTYPE_MULTIOPTION,
			Id:   String("c_myFavoriteColors"),
			Options: []CustomFieldOption{
				CustomFieldOption{
					Key:   "c_blue",
					Value: "Blue",
				},
				CustomFieldOption{
					Key:   "c_red",
					Value: "Red",
				},
				CustomFieldOption{
					Key:   "c_green",
					Value: "Green",
				},
			},
		},
		&CustomField{
			Name: "My Favorite Food",
			Type: CUSTOMFIELDTYPE_MULTIOPTION,
			Id:   String("c_myFavoriteFood"),
			Options: []CustomFieldOption{
				CustomFieldOption{
					Key:   "c_cheese",
					Value: "Cheese",
				},
				CustomFieldOption{
					Key:   "c_olives",
					Value: "Olives",
				},
				CustomFieldOption{
					Key:   "c_chickenFingers",
					Value: "Chicken Fingers",
				},
			},
		},
	},
}

func TestMustIsMultiOptionSet(t *testing.T) {
	if !cfm.MustIsMultiOptionSet("My Favorite Colors", "Red", ToUnorderedStrings([]string{"c_red"})) {
		t.Error("TestMustIsMultiOptionSet: red is set but got false")
	}
	if cfm.MustIsMultiOptionSet("My Favorite Colors", "Red", ToUnorderedStrings([]string{"c_blue"})) {
		t.Error("TestMustIsMultiOptionSet: blue is not set but got true")
	}
	if !cfm.MustIsMultiOptionSet("My Favorite Colors", "Red", ToUnorderedStrings([]string{"c_blue", "c_red"})) {
		t.Error("TestMustIsMultiOptionSet: red is set but got false")
	}
	if cfm.MustIsMultiOptionSet("My Favorite Colors", "Red", ToUnorderedStrings([]string{})) {
		t.Error("TestMustIsMultiOptionSet: red is not set but got true")
	}
}

func TestMustIsSingleOptionSet(t *testing.T) {
	if !cfm.MustIsSingleOptionSet("My Favorite Food", "Cheese", NullableString("c_cheese")) {
		t.Error("TestMustIsSingleOptionSet: cheese is set but got false")
	}
	if cfm.MustIsSingleOptionSet("My Favorite Food", "Olives", NullableString("c_cheese")) {
		t.Error("TestMustIsSingleOptionSet: olives is not set but got true")
	}
	if cfm.MustIsSingleOptionSet("My Favorite Food", "Cheese", NullString()) {
		t.Error("TestMustIsSingleOptionSet: cheese is not set but got true")
	}
}

func TestMustGetMultiOptionNames(t *testing.T) {
	tests := []struct {
		got  []string
		want []string
	}{
		{
			got:  cfm.MustGetMultiOptionNames("My Favorite Food", ToUnorderedStrings([]string{"c_cheese", "c_chickenFingers"})),
			want: []string{"Cheese", "Chicken Fingers"},
		},
		{
			got:  cfm.MustGetMultiOptionNames("My Favorite Colors", ToUnorderedStrings([]string{})),
			want: []string{},
		},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(test.got, test.want) {
			t.Errorf("TestMustGetMultiOptionNames: wanted %v, got %v", test.want, test.got)
		}
	}
}

func TestMustGetSingleOptionName(t *testing.T) {
	tests := []struct {
		got  string
		want string
	}{
		{
			got:  cfm.MustGetSingleOptionName("My Favorite Food", NullableString("c_chickenFingers")),
			want: "Chicken Fingers",
		},
		{
			got:  cfm.MustGetSingleOptionName("My Favorite Colors", NullString()),
			want: "",
		},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("TestMustGetSingleOptionName: wanted %s, got %s", test.want, test.got)
		}
	}
}

func TestSetCustomField(t *testing.T) {
	tests := []struct {
		FieldName string
		Value     interface{}
		Entity    *CustomLocationEntity
		Want      *CustomLocationEntity
	}{
		{
			FieldName: "CF Text",
			Value:     String("New Text"),
			Entity: &CustomLocationEntity{
				CustomEntity: CustomEntity{
					CFUrl: String("www.yext.com"),
				},
			},
			Want: &CustomLocationEntity{
				CustomEntity: CustomEntity{
					CFUrl:  String("www.yext.com"),
					CFText: String("New Text"),
				},
			},
		},
	}

	for _, test := range tests {
		CustomLocationEntityCFManager.SetCustomFieldValue(&test.Entity.CustomEntity, test.FieldName, test.Value)
		if delta, diff, _ := Diff(test.Entity, test.Want); diff {
			t.Errorf("TestSetCustomField:\nWanted: %v\nGot: %v\nDelta: %v", test.Want, test.Entity, delta)
		}
	}
}

func TestGetCustomField(t *testing.T) {
	tests := []struct {
		FieldName string
		Entity    *CustomLocationEntity
		Want      interface{}
	}{
		{
			FieldName: "CF Text",
			Entity: &CustomLocationEntity{
				CustomEntity: CustomEntity{
					CFUrl:  String("www.yext.com"),
					CFText: String("New Text"),
				},
			},
			Want: String("New Text"),
		},
	}

	for _, test := range tests {
		got := CustomLocationEntityCFManager.GetCustomFieldValue(&test.Entity.CustomEntity, test.FieldName)
		if !reflect.DeepEqual(got, test.Want) {
			t.Errorf("TestGetCustomField:\nWanted: %v\nGot: %v", test.Want, got)
		}
	}
}
