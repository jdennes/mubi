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

	api := MubiAPI{client}
	userId, page, perPage := int64(7995037), 1, 20
	ratings := api.GetRatings(userId, page, perPage)

	if len(ratings) != 2 {
		t.Errorf("Number of ratings was incorrect. Got: %d, expected: %d.", len(ratings), 2)
	}

	if ratings[0].Film.Title != "Fallen Angels" {
		t.Errorf("First film title in ratings was incorrect. Got: %s, expected: %s.", ratings[0].Film.Title, "Fallen Angels")
	}
	if ratings[0].Overall != 4 {
		t.Errorf("First rating was incorrect. Got: %d, expected: %d.", ratings[0].Overall, 4)
	}

	if ratings[1].Film.Title != "Little Odessa" {
		t.Errorf("Second film title in ratings was incorrect. Got: %s, expected: %s.", ratings[1].Film.Title, "Little Odessa")
	}
	if ratings[1].Overall != 4 {
		t.Errorf("Second rating was incorrect. Got: %d, expected: %d.", ratings[1].Overall, 4)
	}
}
