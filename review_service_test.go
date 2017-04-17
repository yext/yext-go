package yext

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testing"
)

func TestReviewListOptions(t *testing.T) {
	tests := []struct {
		opts *ReviewListOptions
		want map[string]string
	}{
		{
			opts: nil,
			want: map[string]string{
				"limit":                 "",
				"offset":                "",
				"LocationIds":           "",
				"FolderId":              "",
				"Countries":             "",
				"LocationLabels":        "",
				"PublisherIds":          "",
				"ReviewContent":         "",
				"MinRating":             "",
				"MaxRating":             "",
				"MinPublisherDate":      "",
				"MaxPublisherDate":      "",
				"MinLastYextUpdateDate": "",
				"MaxLastYextUpdateDate": "",
				"AwaitingResponse":      "",
				"MinNonOwnerComments":   "",
				"ReviewerName":          "",
				"ReviewerEmail":         "",
			},
		},
		{
			// The values are technically 0,0, but that doesn't make any sense in the context of a list request
			opts: &ReviewListOptions{},
			want: map[string]string{
				"limit":                 "",
				"offset":                "",
				"LocationIds":           "",
				"FolderId":              "",
				"Countries":             "",
				"LocationLabels":        "",
				"PublisherIds":          "",
				"ReviewContent":         "",
				"MinRating":             "",
				"MaxRating":             "",
				"MinPublisherDate":      "",
				"MaxPublisherDate":      "",
				"MinLastYextUpdateDate": "",
				"MaxLastYextUpdateDate": "",
				"AwaitingResponse":      "",
				"MinNonOwnerComments":   "",
				"ReviewerName":          "",
				"ReviewerEmail":         "",
			},
		},
		{
			opts: &ReviewListOptions{ListOptions: ListOptions{Limit: 10}},
			want: map[string]string{"limit": "10"},
		},
		{
			opts: &ReviewListOptions{ListOptions: ListOptions{Offset: 10}},
			want: map[string]string{"offset": "10"},
		},
		{
			opts: &ReviewListOptions{FolderId: "124"},
			want: map[string]string{"folderId": "124"},
		},
		{
			opts: &ReviewListOptions{Countries: []string{"usa", "china"}},
			want: map[string]string{"countries": "usa,china"},
		},
		{
			opts: &ReviewListOptions{LocationLabels: []string{"label1", "label2"}},
			want: map[string]string{"locationLabels": "label1,label2"},
		},
		{
			opts: &ReviewListOptions{PublisherIds: []string{"99ha", "xwJOE"}},
			want: map[string]string{"publisherIds": "99ha,xwJOE"},
		},
		{
			opts: &ReviewListOptions{ReviewContent: "great experience"},
			want: map[string]string{"reviewContent": "great experience"},
		},
		{
			opts: &ReviewListOptions{MinRating: 2.05},
			want: map[string]string{"minRating": "2.05"},
		},
		{
			opts: &ReviewListOptions{MaxRating: 4.11},
			want: map[string]string{"maxRating": "4.11"},
		},
		{
			opts: &ReviewListOptions{MinPublisherDate: "1999-03-15"},
			want: map[string]string{"minPublisherDate": "1999-03-15"},
		},
		{
			opts: &ReviewListOptions{MaxPublisherDate: "2013-04-01"},
			want: map[string]string{"maxPublisherDate": "2013-04-01"},
		},
		{
			opts: &ReviewListOptions{MinLastYextUpdateDate: "1851-02-24"},
			want: map[string]string{"minLastYextUpdateDate": "1851-02-24"},
		},
		{
			opts: &ReviewListOptions{MaxLastYextUpdateDate: "1900-01-01"},
			want: map[string]string{"maxLastYextUpdateDate": "1900-01-01"},
		},
		{
			opts: &ReviewListOptions{AwaitingResponse: "REVIEW"},
			want: map[string]string{"awaitingResponse": "REVIEW"},
		},
		{
			opts: &ReviewListOptions{MinNonOwnerComments: 2},
			want: map[string]string{"minNonOwnerComments": "2"},
		},
		{
			opts: &ReviewListOptions{ReviewerName: "jeff"},
			want: map[string]string{"reviewerName": "jeff"},
		},
		{
			opts: &ReviewListOptions{ReviewerEmail: "jump@too.high"},
			want: map[string]string{"reviewerEmail": "jump@too.high"},
		},
	}

	for _, test := range tests {
		setup()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			for opt, val := range test.want {
				if v := r.URL.Query().Get(opt); v != val {
					t.Errorf("Wanted limit %s, got %s", val, v)
				}
			}
		})

		client.ReviewService.List(test.opts)
		teardown()
	}
}

func makeRevs(n int) []*Review {
	var revs []*Review

	for i := 0; i < n; i++ {
		new := &Review{Id: &i}
		revs = append(revs, new)
	}

	return revs
}

func TestReviewList(t *testing.T) {
	maxLimit := strconv.Itoa(ReviewListMaxLimit)

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
				t.Errorf("Too many requests sent to review list - got %d, expected %d", reqs, len(test.reqs))
			}

			expectedreq := test.reqs[reqs-1]

			if v := r.URL.Query().Get("limit"); v != expectedreq.limit {
				t.Errorf("Wanted limit %s, got %s", expectedreq.limit, v)
			}

			if v := r.URL.Query().Get("offset"); v != expectedreq.offset {
				t.Errorf("Wanted offset %s, got %s", expectedreq.offset, v)
			}

			revs := []*Review{}
			remaining := test.count - ((reqs - 1) * ReviewListMaxLimit)
			if remaining > 0 {
				if remaining > ReviewListMaxLimit {
					remaining = ReviewListMaxLimit
				}
				revs = makeRevs(remaining)
			}

			v := &mockResponse{Response: &ReviewListResponse{Count: test.count, Reviews: revs}}
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
