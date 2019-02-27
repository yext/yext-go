package yext

import (
	"net/http"
	"testing"
)

func TestSetNilIsEmpty(t *testing.T) {
	type randomStruct struct{}
	tests := []struct {
		i      interface{}
		before bool
		after  bool
	}{
		{
			i:      &BaseEntity{},
			before: false,
			after:  true,
		},
		{
			i: &BaseEntity{
				nilIsEmpty: true,
			},
			before: true,
			after:  true,
		},
		{
			i:      &LocationEntity{},
			before: false,
			after:  true,
		},
		{
			i:      &randomStruct{},
			before: false,
			after:  false,
		},
	}

	for _, test := range tests {
		before := getNilIsEmpty(test.i)
		if before != test.before {
			t.Errorf("Before set nil is empty: Expected %t, got %t", test.before, before)
		}
		setNilIsEmpty(test.i)
		after := getNilIsEmpty(test.i)
		if after != test.after {
			t.Errorf("After set nil is empty: Expected %t, got %t", test.after, after)
		}
	}
}

func TestEntityListOptions(t *testing.T) {
	tests := []struct {
		opts                *EntityListOptions
		limit               string
		token               string
		searchIDs           string
		entityTypes         string
		resolvePlaceholders bool
	}{
		{
			opts:      nil,
			limit:     "",
			token:     "",
			searchIDs: "",
		},
		{
			// The values are technically 0,0, but that doesn't make any sense in the context of a list request
			opts:      &EntityListOptions{ListOptions: ListOptions{}},
			limit:     "",
			token:     "",
			searchIDs: "",
		},
		{
			opts:      &EntityListOptions{ListOptions: ListOptions{Limit: 10}},
			limit:     "10",
			token:     "",
			searchIDs: "",
		},
		{
			opts:        &EntityListOptions{EntityTypes: []string{"location"}},
			limit:       "",
			token:       "",
			searchIDs:   "",
			entityTypes: "location",
		},
		{
			opts:        &EntityListOptions{EntityTypes: []string{"location,event"}},
			limit:       "",
			token:       "",
			searchIDs:   "",
			entityTypes: "location,event",
		},
		{
			opts:      &EntityListOptions{ListOptions: ListOptions{PageToken: "qwerty1234"}},
			limit:     "",
			token:     "qwerty1234",
			searchIDs: "",
		},
		{
			opts:      &EntityListOptions{ListOptions: ListOptions{Limit: 42, PageToken: "asdfgh4321"}},
			limit:     "42",
			token:     "asdfgh4321",
			searchIDs: "",
		},
		{
			opts:      &EntityListOptions{SearchIDs: []string{"1234"}},
			limit:     "",
			token:     "",
			searchIDs: "1234",
		},
		{
			opts:      &EntityListOptions{SearchIDs: []string{"1234", "5678"}, ListOptions: ListOptions{Limit: 42, PageToken: "asdfgh4321"}},
			limit:     "42",
			token:     "asdfgh4321",
			searchIDs: "1234,5678",
		},
		{
			opts:                &EntityListOptions{ResolvePlaceholders: true, ListOptions: ListOptions{Limit: 42, PageToken: "asdfgh4321"}},
			limit:               "42",
			token:               "asdfgh4321",
			resolvePlaceholders: true,
		},
	}

	for _, test := range tests {
		setup()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if v := r.URL.Query().Get("limit"); v != test.limit {
				t.Errorf("Wanted limit %s, got %s", test.limit, v)
			}
			if v := r.URL.Query().Get("pageToken"); v != test.token {
				t.Errorf("Wanted token %s, got %s", test.token, v)
			}
			if v := r.URL.Query().Get("searchIds"); v != test.searchIDs {
				t.Errorf("Wanted searchId %s, got %s", test.searchIDs, v)
			}
			if v := r.URL.Query().Get("entityTypes"); v != test.entityTypes {
				t.Errorf("Wanted entityTypes %s, got %s", test.entityTypes, v)
			}
			v := r.URL.Query().Get("resolvePlaceholders")
			if v == "true" && !test.resolvePlaceholders || v == "" && test.resolvePlaceholders || v == "false" && test.resolvePlaceholders {
				t.Errorf("Wanted resolvePlaceholders %t, got %s", test.resolvePlaceholders, v)
			}
		})

		client.EntityService.List(test.opts)
		teardown()
	}
}
