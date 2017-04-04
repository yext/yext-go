package yext

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/jonboulle/clockwork"
)

func TestResponseDeserialzation(t *testing.T) {
	tests := []struct {
		data       string
		statuscode int
		want       *Response
		obj        interface{}
	}{
		{
			data: `{"meta": {"errors": [{
		  "message": "We had a problem with our software. Please contact support!",
		  "code": 9,
			"type": "FATAL_ERROR"
		}], "uuid": ""}}`,
			want: &Response{
				Meta: Meta{
					Errors: Errors{
						Error{
							Code:    9,
							Type:    ErrorTypeFatal,
							Message: "We had a problem with our software. Please contact support!",
						}}}},
		},
		{
			data: `{"meta": {"errors": [], "uuid": ""}}`,
			want: &Response{},
		},
		{
			data: `{"meta": {"errors": [], "uuid": ""}, "response": {"foo": true}}`,
			want: &Response{Response: map[string]interface{}{"foo": true}},
			obj:  map[string]interface{}{},
		},
		{
			data: `{"meta": {"errors": [], "uuid": "0556abaf-e5fb-475f-8e2a-79688bf4bc18"}}`,
			want: &Response{Meta: Meta{UUID: "0556abaf-e5fb-475f-8e2a-79688bf4bc18"}},
		},
	}

	for _, test := range tests {
		setup()

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(test.data))
		})

		resp, _ := client.DoRequest("", "", test.obj)

		// Avoid DeepEqual issues comparing yext.Errors(nil) and yext.Errors{} as
		// the API always returns `errors: []` vs `errors: null` to indicate 'no error'
		if len(resp.Meta.Errors) == 0 {
			resp.Meta.Errors = nil
		}

		if !reflect.DeepEqual(resp, test.want) {
			t.Errorf("ResponseData: %s\n\tWanted: %#v\n\tGot: %#v", test.data, test.want, resp)
		}

		teardown()
	}
}

func TestRetries(t *testing.T) {
	setup().WithRetries(3)
	defer teardown()

	requests := 0

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requests++
		w.WriteHeader(http.StatusInternalServerError)
	})

	client.DoRequest("GET", "", nil)

	if requests != 4 {
		t.Error("Expected 4 net attempts when error encountered, only got", requests)
	}
}

func TestLastRetryError(t *testing.T) {
	setup().WithRetries(3)
	defer teardown()

	request := 0

	errf := func(n int) string {
		return fmt.Sprintf("error from request #%d", n)
	}

	wraperr := func(m string) string {
		return fmt.Sprintf(`{"meta": {"errors": [{"message": "%s"}], "uuid": ""}}`, m)
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		request++
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(wraperr(errf(request))))
	})

	_, err := client.DoRequest("GET", "", nil)

	expectedErr := errf(*client.Config.RetryCount + 1)
	if !strings.Contains(err.Error(), expectedErr) {
		t.Errorf("Expected to get error `%s`, instead got `%s`", expectedErr, err)
	}
}

func TestBailout(t *testing.T) {
	setup().WithRetries(3)
	defer teardown()

	requests := 0

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requests++

		if requests == 4 {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

	})

	_, err := client.DoRequest("GET", "", nil)

	if err != nil {
		t.Error("Expected error to be nil when final attempt succeeded:", err)
	}
}

func TestRetryWithBody(t *testing.T) {
	setup().WithRetries(3)
	defer teardown()

	body := map[string]interface{}{"foo": "bar"}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		b, _ := ioutil.ReadAll(r.Body)
		var payload map[string]interface{}
		json.Unmarshal(b, &payload)
		if !reflect.DeepEqual(body, payload) {
			t.Error("Expected to get identical body in retry scenario, got", string(b), "instead")
		}

		// Force retries
		w.WriteHeader(http.StatusInternalServerError)
	})

	client.DoRequestJSON("POST", "", body, nil)
}

func TestRetryWith400Error(t *testing.T) {
	setup().WithRetries(3)
	defer teardown()

	requests := 0

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requests++
		w.WriteHeader(http.StatusBadRequest)
	})

	client.DoRequest("GET", "", nil)

	if requests != 1 {
		t.Errorf("Expected 1 net attempts when %d encountered, got %d", http.StatusBadRequest, requests)
	}
}

func TestRetryWith404Error(t *testing.T) {
	setup().WithRetries(3)
	defer teardown()

	requests := 0

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requests++
		w.WriteHeader(http.StatusNotFound)
	})

	client.DoRequest("GET", "", nil)

	if requests != 1 {
		t.Errorf("Expected 1 net attempts when %d encountered, got %d", http.StatusNotFound, requests)
	}
}

func TestAddListOptions(t *testing.T) {
	tests := []struct {
		requrl string
		opts   *ListOptions
		want   string
	}{
		{
			requrl: "locations",
			want:   "locations",
		},
		{
			requrl: "locations",
			opts:   &ListOptions{Limit: 99},
			want:   "locations?limit=99",
		},
		{
			requrl: "locations",
			opts:   &ListOptions{Offset: 99},
			want:   "locations?offset=99",
		},
	}

	for _, test := range tests {
		if got, err := addListOptions(test.requrl, test.opts); err != nil {
			t.Errorf("addListOptions(%s, %#v) error %s", test.requrl, test.opts, err.Error())
		} else if got != test.want {
			t.Errorf("addListOptions(%s, %#v) = %s, wanted %s", test.requrl, test.opts, got, test.want)
		}
	}
}

func TestNoRetryOnDeserError(t *testing.T) {
	setup().WithRetries(3)
	defer teardown()

	requests := 0

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requests++
		w.Write([]byte(`{"meta": {"errors": [], "uuid": ""}, "response": false}`))
	})

	var v map[string]interface{}
	_, err := client.DoRequest("GET", "", &v)

	if err == nil {
		t.Error("Expected deserialization error")
	}

	if requests != 1 {
		t.Errorf("Expected 1 net attempts when %s encountered, got %d", err.Error(), requests)
	}
}

func TestRateLimitParsing(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Rate-Limit-Limit", "5000")
		w.Header().Set("Rate-Limit-Remaining", "4000")
		w.Header().Set("Rate-Limit-Reset", "1490799600")
		w.WriteHeader(http.StatusOK)
	})

	resp, err := client.DoRequest("GET", "", nil)

	if err != nil {
		t.Error(err)
	}

	if v := resp.RateLimitLimit; v != 5000 {
		t.Errorf("Expected RateLimitLimit of 5000, got %d", v)
	}

	if v := resp.RateLimitRemaining; v != 4000 {
		t.Errorf("Expected RateLimitRemaining of 4000, got %d", v)
	}

	if v := resp.RateLimitReset; v != 1490799600 {
		t.Errorf("Expected RateLimitReset of 1490799600, got %d", v)
	}
}

func TestRateLimitWaiting(t *testing.T) {
	conf := setup().WithRateLimitRetry().WithMockClock()
	defer teardown()
	fClock := conf.Clock.(clockwork.FakeClock)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resetTime := strconv.FormatInt(fClock.Now().Unix()+1, 10)
		w.Header().Set("Rate-Limit-Limit", "5000")
		w.Header().Set("Rate-Limit-Remaining", "4000")
		w.Header().Set("Rate-Limit-Reset", resetTime)
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte(`{"meta": {"errors": [{"message": "", "code": 1}], "uuid": ""}}`))
	})

	timeBefore := fClock.Now()
	go func() {
		client.DoRequest("GET", "", nil)
	}()
	go func() {
		time.Sleep(500 * time.Millisecond)
		fClock.Advance(time.Hour)
		fClock.Sleep(time.Hour)
	}()
	fClock.BlockUntil(1)
	if fClock.Since(timeBefore) < time.Minute {
		return
	} else {
		t.Errorf("No sleep initiated after waiting .5 seconds")
	}
}
