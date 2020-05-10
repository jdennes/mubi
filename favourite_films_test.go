package mubi

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestGetFavouriteFilms(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		file, _ := os.Open("testdata/mubi-favourite-films-sample.json")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(file),
			Header:     make(http.Header),
		}
	})

	api := MubiAPI{client}
	userId, page, perPage := int64(7995037), 1, 20
	favourites := api.GetFavouriteFilms(userId, page, perPage)

	if len(favourites) != 2 {
		t.Errorf("Number of favourite films was incorrect. Got: %d, expected: %d.", len(favourites), 2)
	}

	if favourites[0].Fannable.Film.Title != "4 Months, 3 Weeks and 2 Days" {
		t.Errorf(
			"First film title on favourites list was incorrect. Got: %s, expected: %s.",
			favourites[0].Fannable.Film.Title, "4 Months, 3 Weeks and 2 Days",
		)
	}

	if favourites[1].Fannable.Film.Title != "Herr Lehmann" {
		t.Errorf(
			"Second film title on favourites list was incorrect. Got: %s, expected: %s.",
			favourites[1].Fannable.Film.Title, "Herr Lehmann",
		)
	}
}
