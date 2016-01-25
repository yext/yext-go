package yext

import (
	"net/http"
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
