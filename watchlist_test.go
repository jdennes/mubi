package mubi

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestGetWatchlist(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		file, _ := os.Open("testdata/mubi-watchlist-sample.json")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(file),
			Header:     make(http.Header),
		}
	})

	watchlistApi := WatchlistAPI{client}
	watchlist := watchlistApi.GetWatchlist(7995037)

	if len(watchlist) != 2 {
		t.Errorf("Number of watchlist items was incorrect. Got: %d, expected: %d.", len(watchlist), 2)
	}

	if watchlist[0].Film.Title != "Fireworks" {
		t.Errorf("First film title on watchlist was incorrect. Got: %s, expected: %s.", watchlist[0].Film.Title, "Fireworks")
	}

	if watchlist[1].Film.Title != "The Intruder" {
		t.Errorf("Second film title on watchlist was incorrect. Got: %s, expected: %s.", watchlist[1].Film.Title, "The Intruder")
	}
}
