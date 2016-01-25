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

func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// client configured to use test server
	client = NewClient("", "", "customerId", Config{})
	client.baseUrl = server.URL

	// Disable retries in test
	client.retryAttempts = 0

	// 0 delay between retries for test
	DefaultBackoffPolicy = BackoffPolicy{[]int{0}}
}

func teardown() {
	server.Close()
}
