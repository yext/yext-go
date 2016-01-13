package yext

import (
	"net/http"
	"testing"
)

var (
	licenseId  = "12345"
	locationId = "0A8804"
)

func TestGetIsGET(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}
	})

	client.LicenseService.Get(licenseId)
}

func TestGetResponse(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(licenseResponse))
	})

	resp, err := client.LicenseService.Get(licenseId)
	if err != nil {
		t.Error("Error making request:", err)
	}

	if resp.Status != "ACTIVE" {
		t.Error("Did not get successful status from response.")
	}
}

func TestGetResponseError(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(licenseDoesNotExistResponse))
	})

	resp, err := client.LicenseService.Get(licenseId)
	if err != nil {
		t.Error("Error making request:", err)
	}

	if len(resp.LocationIds) > 0 {
		t.Error("Did not get successful status from response.")
	}
}

func TestAddLocationToLicenseIsPUT(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := "PUT"; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}
	})

	client.LicenseService.AddLocationToLicense(licenseId, locationId)
}

func TestAddLocationToLicense(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(licenseResponse))
	})

	resp, err := client.LicenseService.AddLocationToLicense(licenseId, locationId)
	if err != nil {
		t.Error("Error making request:", err)
	}

	for _, v := range resp.LocationIds {
		if v == locationId {
			return
		}
		t.Error("Location not correctly added to the license.")
	}
}

func TestAddLocationToLicenseError(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(licenseError))
	})

	resp, err := client.LicenseService.AddLocationToLicense(licenseId, locationId)
	if err != nil {
		t.Errorf("Error making request:", err)
	}

	if len(resp.LocationIds) > 0 {
		t.Error("Did not get successful status from response.")
	}
}

func TestARemoveLocationFromLicenseIsDELETE(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := "DELETE"; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}
	})

	client.LicenseService.RemoveLocationFromLicense(licenseId, locationId)
}

func TestRemoveLocationFromLicense(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(licenseRemoveResponse))
	})

	resp, err := client.LicenseService.RemoveLocationFromLicense(licenseId, locationId)
	if err != nil {
		t.Error("Error making request:", err)
	}

	for _, v := range resp.LocationIds {
		if v == locationId {
			t.Error("Location not correctly removed to the license.")
		}
	}
}

func TestRemoveLocationFromLicenseError(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(licenseError))
	})

	resp, err := client.LicenseService.RemoveLocationFromLicense(licenseId, locationId)
	if err != nil {
		t.Error("Error making request:", err)
	}

	if len(resp.LocationIds) > 0 {
		t.Error("Did not get successful status from response.")
	}
}

var licenseResponse = `
{
  "locationIds": [
    "0A8804",
    "0A8823",
    "0A8845",
    "0A8404",
    "0A8789",
    "0A8908",
    "0B0291",
    "0A8906",
    "0A8787"
  ],
  "features": [
    "Yellow Cauldron [TEST]",
    "2findlocal",
    "Custom Fields"
  ],
  "id": 16091,
  "quantity": 1000,
  "status": "ACTIVE"
}`

var licenseDoesNotExistResponse = `
{
  "errors": [
    {
      "message": "License pack does not exist.",
      "errorCode": 3034
    }
  ]
}`

var licenseError = `
{
  "errors": [
    {
      "message": "Location not found: 0A880",
      "errorCode": 2000
    }
  ]
}`

var licenseRemoveResponse = `
{
  "locationIds": [
    "0A8823",
    "0A8845",
    "0A8404",
    "0A8789",
    "0A8908",
    "0B0291",
    "0A8906",
    "0A8787"
  ],
  "features": [
    "Yellow Cauldron [TEST]",
    "2findlocal",
    "Custom Fields"
  ],
  "id": 16091,
  "quantity": 1000,
  "status": "ACTIVE"
}`
