package yext

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestErrorDeserialzation(t *testing.T) {
	setup()
	defer teardown()

	errorResp := `{"meta": {"errors": [{
	  "message": "We had a problem with our software. Please contact support!",
	  "errorCode": 9
	}], "uuid": "123456789"}}`

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errorResp))
	})

	err := client.DoRequest("", "", nil)

	if _, ok := err.(*Response); !ok {
		t.Error("Expected to recieve *Response type, got", err, "instead")
	}

	if !strings.Contains(err.Error(), "message: We had a problem with") {
		t.Error("Issue parsing the error, was: ", err.Error())
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

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		request++
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errf(request)))
	})

	err := client.DoRequest("GET", "", nil)

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

	err := client.DoRequest("GET", "", nil)

	if err != nil {
		t.Error("Expected error to be nil when final attempt succeeded:", err)
	}
}

func TestRetryWithBody(t *testing.T) {
	setup().WithRetries(3)
	defer teardown()

	requests := 0
	body := map[string]interface{}{"foo": "bar"}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requests++

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

func TestMetaDeserialization(t *testing.T) {
	setup()
	defer teardown()

	resp := `{
		  "meta": {
		    "uuid": "0556abaf-e5fb-475f-8e2a-79688bf4bc18",
		    "errors": [
		    ]
		  },
		  "response": {
		    "count": 1,
		    "locations": [
		      {
		        "id": "3810730"
					}
		    ]
		  }
		}`

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(resp))
	})

	v := &locationListResponse{}
	err := client.DoRequest("", "", v)
	if err != nil {
		t.Error("Error came back in request")
	}

	if len(v.Locations) < 1 {
		t.Error("No Locations came back")
	}

	if *v.Locations[0].Id != "3810730" {
		t.Errorf("ID didn't come back correctly, it was %v instead", *v.Locations[0].Id)
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
