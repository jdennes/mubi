package mubi

import (
	"encoding/json"
	"fmt"
	"log"
)

type FollowingItem struct {
	Id        int   `json:id`
	CreatedAt int64 `json:"created_at"`
	Followee  User  `json:followee`
}

func (api *MubiAPI) GetFollowing(userId int64, page int, perPage int) []FollowingItem {
	url := fmt.Sprintf(
		"https://mubi.com/services/api/followings?user_id=%d&page=%d&per_page=%d",
		userId, page, perPage,
	)

	body := api.GetResponseBody(url)
	following := make([]FollowingItem, 0)
	jsonErr := json.Unmarshal(body, &following)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return following
}
