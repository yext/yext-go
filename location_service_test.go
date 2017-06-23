package yext

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testing"
)

func TestListOptions(t *testing.T) {
	tests := []struct {
		opts      *LocationListOptions
		limit     string
		token     string
		searchIDs string
	}{
		{
			opts:      nil,
			limit:     "",
			token:     "",
			searchIDs: "",
		},
		{
			// The values are technically 0,0, but that doesn't make any sense in the context of a list request
			opts:      &LocationListOptions{ListOptions: ListOptions{}},
			limit:     "",
			token:     "",
			searchIDs: "",
		},
		{
			opts:      &LocationListOptions{ListOptions: ListOptions{Limit: 10}},
			limit:     "10",
			token:     "",
			searchIDs: "",
		},
		{
			opts:      &LocationListOptions{ListOptions: ListOptions{PageToken: "qwerty1234"}},
			limit:     "",
			token:     "qwerty1234",
			searchIDs: "",
		},
		{
			opts:      &LocationListOptions{ListOptions: ListOptions{Limit: 42, PageToken: "asdfgh4321"}},
			limit:     "42",
			token:     "asdfgh4321",
			searchIDs: "",
		},
		{
			opts:      &LocationListOptions{SearchIDs: []string{"1234"}, ListOptions: ListOptions{Limit: 42, PageToken: "asdfgh4321"}},
			limit:     "42",
			token:     "asdfgh4321",
			searchIDs: "1234",
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
		})

		client.LocationService.List(test.opts)
		teardown()
	}
}

func makeLocs(n int) []*Location {
	var locs []*Location

	for i := 0; i < n; i++ {
		new := &Location{Id: String(strconv.Itoa(i))}
		locs = append(locs, new)
	}

	return locs
}

func TestListBySearchIds(t *testing.T) {
	maxLimit := strconv.Itoa(LocationListMaxLimit)

	tests := []struct {
		limit                 string
		tokenResponses        []string
		expectedTokenRequests []string
		searchIds             []string
	}{
		{
			limit:                 maxLimit,
			tokenResponses:        []string{""},
			expectedTokenRequests: []string{""},
			searchIds:             []string{"1234"},
		},
		{
			limit:                 maxLimit,
			tokenResponses:        []string{"first_token"},
			expectedTokenRequests: []string{"", "first_token"},
			searchIds:             []string{"1234", "1234"},
		},
		{
			limit:                 maxLimit,
			tokenResponses:        []string{"first_token", "second_token", "third_token"},
			expectedTokenRequests: []string{"", "first_token", "second_token", "third_token"},
			searchIds:             []string{"1234", "1234", "1234", "1234"},
		},
	}

	for _, test := range tests {
		setup()
		reqs := 0
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			reqs++
			if reqs > len(test.expectedTokenRequests) {
				t.Errorf("Too many requests sent to location list - got %d, expected %d", reqs, len(test.expectedTokenRequests))
			}

			expectedreq := test.expectedTokenRequests[reqs-1]
			tokenresp := ""
			if reqs <= len(test.tokenResponses) {
				tokenresp = test.tokenResponses[reqs-1]
			}
			searchId := test.searchIds[reqs-1]

			if v := r.URL.Query().Get("searchIds"); v != searchId {
				t.Errorf("Wanted searchId %s, got %s", searchId, v)
			}

			if v := r.URL.Query().Get("limit"); v != test.limit {
				t.Errorf("Wanted limit %s, got %s", test.limit, v)
			}

			if v := r.URL.Query().Get("pageToken"); v != expectedreq {
				t.Errorf("Wanted offset %s, got %s", expectedreq, v)
			}

			if tokenresp != "" {
				v := &mockResponse{Response: &LocationListResponse{NextPageToken: tokenresp}}
				data, _ := json.Marshal(v)
				w.Write(data)
			}
		})

		client.LocationService.ListBySearchIds([]string{"1234"})
		if reqs < len(test.expectedTokenRequests) {
			t.Errorf("Too few requests sent to location list - got %d, expected %d", reqs, len(test.expectedTokenRequests))
		}

		teardown()
	}
}

func TestTokenListAll(t *testing.T) {
	maxLimit := strconv.Itoa(LocationListMaxLimit)

	tests := []struct {
		limit                 string
		tokenResponses        []string
		expectedTokenRequests []string
	}{
		{
			limit:                 maxLimit,
			tokenResponses:        []string{""},
			expectedTokenRequests: []string{""},
		},
		{
			limit:                 maxLimit,
			tokenResponses:        []string{"first_token"},
			expectedTokenRequests: []string{"", "first_token"},
		},
		{
			limit:                 maxLimit,
			tokenResponses:        []string{"first_token", "second_token", "third_token"},
			expectedTokenRequests: []string{"", "first_token", "second_token", "third_token"},
		},
	}

	for _, test := range tests {
		setup()
		reqs := 0
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			reqs++
			if reqs > len(test.expectedTokenRequests) {
				t.Errorf("Too many requests sent to location list - got %d, expected %d", reqs, len(test.expectedTokenRequests))
			}

			expectedreq := test.expectedTokenRequests[reqs-1]
			tokenresp := ""
			if reqs <= len(test.tokenResponses) {
				tokenresp = test.tokenResponses[reqs-1]
			}

			if v := r.URL.Query().Get("limit"); v != test.limit {
				t.Errorf("Wanted limit %s, got %s", test.limit, v)
			}

			if v := r.URL.Query().Get("pageToken"); v != expectedreq {
				t.Errorf("Wanted token %s, got %s", expectedreq, v)
			}

			if tokenresp != "" {
				v := &mockResponse{Response: &LocationListResponse{NextPageToken: tokenresp}}
				data, _ := json.Marshal(v)
				w.Write(data)
			}
		})

		client.LocationService.ListAll()
		if reqs != len(test.expectedTokenRequests) {
			t.Errorf("Wrong number of requests sent to location list - got %d, expected %d", reqs, len(test.expectedTokenRequests))
		}

		teardown()
	}
}
