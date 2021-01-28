package mubi

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestGetFilmsOfTheDay(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		file, _ := os.Open("testdata/mubi.com-film-of-the-day.html")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(file),
			Header:     make(http.Header),
		}
	})

	api := MubiAPI{client}
	fotd := api.GetFilmsOfTheDay()

	if len(fotd) != 30 {
		t.Errorf("Number of films of the day was incorrect. Got: %d, expected: %d.", len(fotd), 30)
	}
}
