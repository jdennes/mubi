package mubi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type WatchlistItem struct {
	Timestamp int64 `json:"created_at"`
	Film      Film
}

type WatchlistAPI struct {
	Client *http.Client
}

func (api *WatchlistAPI) GetWatchlist(userId int64) []WatchlistItem {
	url := fmt.Sprintf("https://mubi.com/services/api/wishes?user_id=%d", userId)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "jdennes/mubi")
	res, getErr := api.Client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	watchlist := make([]WatchlistItem, 0)
	jsonErr := json.Unmarshal(body, &watchlist)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return watchlist
}
