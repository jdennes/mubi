package mubi

import (
	"net/http"
)

// Tests use example JSON responses as fixtures from the testdata directory and
// replace Transport on http.Client to avoid real HTTP requests.

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}
