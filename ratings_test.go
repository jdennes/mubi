package mubi

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

// Uses example JSON responses as fixtures from the testdata directory and
// replaces Transport on http.Client to avoid real HTTP requests.

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestGetRatings(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		file, _ := os.Open("testdata/mubi-ratings-sample.json")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(file),
			Header:     make(http.Header),
		}
	})

	ratingsApi := RatingsAPI{client}
	ratings := ratingsApi.GetRatings(7995037)

	if len(ratings) != 2 {
		t.Errorf("Ratings length was incorrect. Got: %d, expected: %d.", len(ratings), 2)
	}
}
