package yext

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestErrorDeserialzation(t *testing.T) {
	setup()
	defer teardown()

	errorResp := `{"errors": [{
  "message": "We had a problem with our software. Please contact support!",
  "errorCode": 9
}]}`

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errorResp))
	})

	err := client.DoRequest("", "", nil)

	if _, ok := err.(*ErrorResponse); !ok {
		t.Error("Expected to recieve *ErrorResponse type, got", err, "instead")
	}
}

func TestRetries(t *testing.T) {
	setup()
	defer teardown()
	client.retryAttempts = 3

	requests := 0

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requests++
		w.WriteHeader(http.StatusInternalServerError)
	})

	client.DoRequest("", "", nil)

	if requests != 4 {
		t.Error("Expected 4 net attempts when error encountered, only got", requests)
	}
}

func TestLastRetryError(t *testing.T) {
	setup()
	defer teardown()
	client.retryAttempts = 3

	request := 0

	errf := func(n int) string {
		return fmt.Sprintf("error from request #%d", n)
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		request++
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errf(request)))
	})

	err := client.DoRequest("", "", nil)

	expectedErr := errf(client.retryAttempts + 1)
	if !strings.Contains(err.Error(), expectedErr) {
		t.Errorf("Expected to get error `%s`, instead got `%s`", expectedErr, err)
	}
}

func TestBailout(t *testing.T) {
	setup()
	client.retryAttempts = 3
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

	err := client.DoRequest("", "", nil)

	if err != nil {
		t.Error("Expected error to be nil when final attempt succeeded:", err)
	}
}
