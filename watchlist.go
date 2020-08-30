package mubi

import (
	"encoding/json"
	"fmt"
	"log"
)

type WatchlistItem struct {
	Timestamp int64 `json:"created_at"`
	Film      Film  `json:film`
}

func (api *MubiAPI) GetWatchlist(userId int64, page int, perPage int) []WatchlistItem {
	url := fmt.Sprintf(
		"https://mubi.com/services/api/wishes?user_id=%d&page=%d&per_page=%d",
		userId, page, perPage,
	)

	body := api.GetResponseBody(url)
	watchlist := make([]WatchlistItem, 0)
	jsonErr := json.Unmarshal(body, &watchlist)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return watchlist
}
