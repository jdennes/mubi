package mubi

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestGetFollowing(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		file, _ := os.Open("testdata/mubi-following-sample.json")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(file),
			Header:     make(http.Header),
		}
	})

	api := MubiAPI{client}
	userId, page, perPage := int64(6164674), 1, 20
	following := api.GetFollowing(userId, page, perPage)

	if len(following) != 20 {
		t.Errorf("Number of following items was incorrect. Got: %d, expected: %d.", len(following), 20)
	}

	if following[0].Followee.Name != "mubianer" {
		t.Errorf("Name of first followed user was incorrect. Got: %s, expected: %s.", following[0].Followee.Name, "mubianer")
	}

	if following[1].Followee.Name != "RogerTheMovieManiac88" {
		t.Errorf("Name of second followed user was incorrect. Got: %s, expected: %s.", following[1].Followee.Name, "RogerTheMovieManiac88")
	}
}
