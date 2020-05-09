package mubi

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

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
		t.Errorf("Number of ratings was incorrect. Got: %d, expected: %d.", len(ratings), 2)
	}
}
