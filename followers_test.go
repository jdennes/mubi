package mubi

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestGetFollowers(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		file, _ := os.Open("testdata/mubi-followers-sample.json")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(file),
			Header:     make(http.Header),
		}
	})

	api := MubiAPI{client}
	userId, page, perPage := int64(6164674), 1, 20
	followers := api.GetFollowers(userId, page, perPage)

	if len(followers) != 20 {
		t.Errorf("Number of follower items was incorrect. Got: %d, expected: %d.", len(followers), 20)
	}

	if followers[0].Follower.Name != "Ryuu Baron" {
		t.Errorf("Name of first followed user was incorrect. Got: %s, expected: %s.", followers[0].Follower.Name, "Ryuu Baron")
	}

	if followers[1].Follower.Name != "Filip Klokan" {
		t.Errorf("Name of second followed user was incorrect. Got: %s, expected: %s.", followers[1].Follower.Name, "Filip Klokan")
	}
}
