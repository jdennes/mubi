package mubi

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestGetLiveFilmScreening(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		file, _ := os.Open("testdata/mubi-live-film-sample.json")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(file),
			Header:     make(http.Header),
		}
	})

	api := MubiAPI{client}
	screening := api.GetLiveFilmScreening()

	if screening.Film.Title != "Innocents with Dirty Hands" {
		t.Errorf("Screening film title was incorrect. Got: %s, expected: %s.", screening.Film.Title, "Innocents with Dirty Hands")
	}

	if screening.Film.Year != 1975 {
		t.Errorf("Screening film year was incorrect. Got: %d, expected: %d.", screening.Film.Year, 1975)
	}
}
