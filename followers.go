package mubi

import (
	"encoding/json"
	"fmt"
	"log"
)

type FollowerItem struct {
	Id        int   `json:id`
	CreatedAt int64 `json:"created_at"`
	Follower  User  `json:follower`
}

func (api *MubiAPI) GetFollowers(userId int64, page int, perPage int) []FollowerItem {
	url := fmt.Sprintf(
		"https://mubi.com/services/api/followings/followers?user_id=%d&page=%d&per_page=%d",
		userId, page, perPage,
	)

	body := api.GetResponseBody(url)
	followers := make([]FollowerItem, 0)
	jsonErr := json.Unmarshal(body, &followers)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return followers
}
