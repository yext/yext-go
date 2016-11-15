package yext

import (
	"net/http"
	"net/http/httptest"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

func setup() *Config {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// client configured to use test server
	config := NewConfig().
		WithHTTPClient(http.DefaultClient).
		WithBaseUrl(server.URL). // Use test server
		WithApiKey("apikey").    // Customer ID needs to be set to something to avoid '//' in the URL path
		WithRetries(0)           // No retries

	client = NewClient(config)
	// No delay between attempts
	DefaultBackoffPolicy = BackoffPolicy{[]int{0}}

	return config
}

func teardown() {
	server.Close()
}
