package mubi

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestGetLists(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		file, _ := os.Open("testdata/mubi-lists-sample.json")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(file),
			Header:     make(http.Header),
		}
	})

	api := MubiAPI{client}
	userId, page, perPage := int64(56195), 1, 20
	lists := api.GetLists(userId, page, perPage)

	if len(lists) != 20 {
		t.Errorf("Number of lists was incorrect. Got: %d, expected: %d.", len(lists), 20)
	}

	if lists[0].Title != "100 DIRECTORS' ESSENTIAL FILMS" {
		t.Errorf("Title of first list was incorrect. Got: %s, expected: %s.", lists[0].Title, "100 DIRECTORS' ESSENTIAL FILMS")
	}

	if lists[1].Title != "MY FAVOURITE FILMS BY WOMEN" {
		t.Errorf("Title of second list user was incorrect. Got: %s, expected: %s.", lists[1].Title, "MY FAVOURITE FILMS BY WOMEN")
	}
}
