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

	// Show requests for debugging
	// client.ShowRequest = true
}

func teardown() {
	server.Close()
}
