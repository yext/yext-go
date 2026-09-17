package yext

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
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
				"pageToken":             "",
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
				"Status":                "",
			},
		},
		{
			// The values are technically 0,0, but that doesn't make any sense in the context of a list request
			opts: &ReviewListOptions{},
			want: map[string]string{
				"limit":                 "",
				"offset":                "",
				"pageToken":             "",
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
				"Status":                "",
			},
		},
		{
			opts: &ReviewListOptions{ListOptions: ListOptions{PageToken: "qwerty1234"}},
			want: map[string]string{"pageToken": "qwerty1234"},
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
			opts: &ReviewListOptions{ListOptions: ListOptions{PageToken: "qwerty1234", Offset: 10}},
			want: map[string]string{"pageToken": "qwerty1234", "offset": ""},
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
		{
			opts: &ReviewListOptions{Status: "LIVE"},
			want: map[string]string{"status": "LIVE"},
		},
	}

	for _, test := range tests {
		setup()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			for opt, val := range test.want {
				if v := r.URL.Query().Get(opt); v != val {
					t.Errorf("Wanted %s %s, got %s", opt, val, v)
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
				t.Errorf("Too many requests sent to review list - got %d, expected %d", reqs, len(test.expectedTokenRequests))
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

		client.ReviewService.ListAll()
		if reqs < len(test.expectedTokenRequests) {
			t.Errorf("Too few requests sent to review list - got %d, expected %d", reqs, len(test.expectedTokenRequests))
		}

		teardown()
	}
}

func TestReviewSyndicate(t *testing.T) {
	tests := []struct {
		id           int
		opts         *ReviewSyndicateOptions
		wantBody     string
		mockResponse *ReviewSyndicateResponse
		want         *ReviewSyndicateResponse
	}{
		{
			id:       123,
			opts:     &ReviewSyndicateOptions{PublisherId: 456},
			wantBody: `{"publisherId":456}`,
			mockResponse: &ReviewSyndicateResponse{
				Id:               "0",
				Outcome:          "SUCCESS",
				ExternalReviewId: "ext-123",
			},
			want: &ReviewSyndicateResponse{
				Id:               "0",
				Outcome:          "SUCCESS",
				ExternalReviewId: "ext-123",
			},
		},
		{
			id:       789,
			opts:     nil,
			wantBody: `null`,
			mockResponse: &ReviewSyndicateResponse{
				Id:      "789",
				Outcome: "FAILURE",
			},
			want: &ReviewSyndicateResponse{
				Id:      "789",
				Outcome: "FAILURE",
			},
		},
		{
			id:       789,
			opts:     &ReviewSyndicateOptions{},
			wantBody: `{}`,
			mockResponse: &ReviewSyndicateResponse{
				Id:      "789",
				Outcome: "SUCCESS",
			},
			want: &ReviewSyndicateResponse{
				Id:      "789",
				Outcome: "SUCCESS",
			},
		},
	}

	for _, test := range tests {
		setup()

		var gotPath, gotMethod, gotBody string
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			gotPath = r.URL.Path
			gotMethod = r.Method

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("Failed to read request body: %v", err)
			}
			gotBody = strings.TrimSpace(string(body))

			v := &mockResponse{Response: test.mockResponse}
			data, _ := json.Marshal(v)
			w.Write(data)
		})

		got, _, err := client.ReviewService.Syndicate(test.id, test.opts)
		if err != nil {
			t.Errorf("Unexpected error calling Syndicate: %v", err)
		}

		wantPath := "/accounts/" + client.Config.AccountId + "/reviews/" + strconv.Itoa(test.id) + "/syndicate"
		if gotPath != wantPath {
			t.Errorf("Wanted path %s, got %s", wantPath, gotPath)
		}

		if gotMethod != http.MethodPost {
			t.Errorf("Wanted method %s, got %s", http.MethodPost, gotMethod)
		}

		if gotBody != test.wantBody {
			t.Errorf("Wanted body %s, got %s", test.wantBody, gotBody)
		}

		if got.Id != test.want.Id || got.Outcome != test.want.Outcome || got.ExternalReviewId != test.want.ExternalReviewId {
			t.Errorf("Wanted response %+v, got %+v", test.want, got)
		}

		teardown()
	}
}
