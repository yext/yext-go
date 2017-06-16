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
		offset    string
		searchIDs string
	}{
		{
			opts:      nil,
			limit:     "",
			offset:    "",
			searchIDs: "",
		},
		{
			// The values are technically 0,0, but that doesn't make any sense in the context of a list request
			opts:      &LocationListOptions{ListOptions: ListOptions{}},
			limit:     "",
			offset:    "",
			searchIDs: "",
		},
		{
			opts:      &LocationListOptions{ListOptions: ListOptions{Limit: 10}},
			limit:     "10",
			offset:    "",
			searchIDs: "",
		},
		{
			opts:      &LocationListOptions{ListOptions: ListOptions{Offset: 10}},
			limit:     "",
			offset:    "10",
			searchIDs: "",
		},
		{
			opts:      &LocationListOptions{ListOptions: ListOptions{Limit: 42, Offset: 33}},
			limit:     "42",
			offset:    "33",
			searchIDs: "",
		},

		{
			opts:      &LocationListOptions{SearchIDs: []string{"1234"}, ListOptions: ListOptions{Limit: 42, Offset: 33}},
			limit:     "42",
			offset:    "33",
			searchIDs: "1234",
		},
	}

	for _, test := range tests {
		setup()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if v := r.URL.Query().Get("limit"); v != test.limit {
				t.Errorf("Wanted limit %s, got %s", test.limit, v)
			}

			if v := r.URL.Query().Get("offset"); v != test.offset {
				t.Errorf("Wanted offset %s, got %s", test.offset, v)
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

func TestListAll(t *testing.T) {
	maxLimit := strconv.Itoa(LocationListMaxLimit)

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
			count: 50,
			reqs:  []req{{limit: maxLimit, offset: ""}},
		},
		{
			count: 51,
			reqs:  []req{{limit: maxLimit, offset: ""}, {limit: maxLimit, offset: "50"}},
		},
		{
			count: 100,
			reqs:  []req{{limit: maxLimit, offset: ""}, {limit: maxLimit, offset: "50"}},
		},
	}

	for _, test := range tests {
		setup()
		reqs := 0
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			reqs++
			if reqs > len(test.reqs) {
				t.Errorf("Too many requests sent to location list - got %d, expected %d", reqs, len(test.reqs))
			}

			expectedreq := test.reqs[reqs-1]

			if v := r.URL.Query().Get("limit"); v != expectedreq.limit {
				t.Errorf("Wanted limit %s, got %s", expectedreq.limit, v)
			}

			if v := r.URL.Query().Get("offset"); v != expectedreq.offset {
				t.Errorf("Wanted offset %s, got %s", expectedreq.offset, v)
			}

			locs := []*Location{}
			remaining := test.count - ((reqs - 1) * LocationListMaxLimit)
			if remaining > 0 {
				if remaining > LocationListMaxLimit {
					remaining = LocationListMaxLimit
				}
				locs = makeLocs(remaining)
			}

			v := &mockResponse{Response: &LocationListResponse{Count: test.count, Locations: locs}}
			data, _ := json.Marshal(v)
			w.Write(data)
		})

		client.LocationService.ListAll()
		if reqs < len(test.reqs) {
			t.Errorf("Too few requests sent to location list - got %d, expected %d", reqs, len(test.reqs))
		}

		teardown()
	}
}

func TestListMismatchCount(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		v := &mockResponse{Response: &LocationListResponse{Count: 25, Locations: makeLocs(24)}}
		data, _ := json.Marshal(v)
		w.Write(data)
	})

	llr, err := client.LocationService.ListAll()
	if llr != nil {
		t.Error("Expected response to be nil when recieving mismatched count and number of locations")
	}
	if err == nil {
		t.Error("Expected error to be present when recieving mismatched count and number of locations")
	}
}

func TestListBySearchIds(t *testing.T) {
	maxLimit := strconv.Itoa(LocationListMaxLimit)

	type req struct {
		limit     string
		offset    string
		searchIds string
	}

	tests := []struct {
		count int
		reqs  []req
	}{
		{
			count: 0,
			reqs:  []req{{limit: maxLimit, offset: "", searchIds: "1234"}},
		},
		{
			count: 50,
			reqs:  []req{{limit: maxLimit, offset: "", searchIds: "1234"}},
		},
		{
			count: 51,
			reqs:  []req{{limit: maxLimit, offset: "", searchIds: "1234"}, {limit: maxLimit, offset: "50", searchIds: "1234"}},
		},
		{
			count: 100,
			reqs:  []req{{limit: maxLimit, offset: "", searchIds: "1234"}, {limit: maxLimit, offset: "50", searchIds: "1234"}},
		},
	}

	for _, test := range tests {
		setup()
		reqs := 0
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			reqs++
			if reqs > len(test.reqs) {
				t.Errorf("Too many requests sent to location list - got %d, expected %d", reqs, len(test.reqs))
			}

			expectedreq := test.reqs[reqs-1]

			if v := r.URL.Query().Get("searchIds"); v != expectedreq.searchIds {
				t.Errorf("Wanted searchId %s, got %s", expectedreq.searchIds, v)
			}

			if v := r.URL.Query().Get("limit"); v != expectedreq.limit {
				t.Errorf("Wanted limit %s, got %s", expectedreq.limit, v)
			}

			if v := r.URL.Query().Get("offset"); v != expectedreq.offset {
				t.Errorf("Wanted offset %s, got %s", expectedreq.offset, v)
			}

			locs := []*Location{}
			remaining := test.count - ((reqs - 1) * LocationListMaxLimit)
			if remaining > 0 {
				if remaining > LocationListMaxLimit {
					remaining = LocationListMaxLimit
				}
				locs = makeLocs(remaining)
			}

			v := &mockResponse{Response: &LocationListResponse{Count: test.count, Locations: locs}}
			data, _ := json.Marshal(v)
			w.Write(data)
		})

		client.LocationService.ListBySearchIds([]string{"1234"})
		if reqs < len(test.reqs) {
			t.Errorf("Too few requests sent to location list - got %d, expected %d", reqs, len(test.reqs))
		}

		teardown()
	}
}
