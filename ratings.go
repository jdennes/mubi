package mubi

import (
	"encoding/json"
	"fmt"
	"log"
)

type Rating struct {
	Overall   int   `json:overall`
	Timestamp int64 `json:"updated_at"`
	Film      Film  `json:film`
}

func (api *MubiAPI) GetRatings(userId int64, page int, perPage int) []Rating {
	url := fmt.Sprintf(
		"https://mubi.com/services/api/ratings?user_id=%d&page=%d&per_page=%d",
		userId, page, perPage,
	)

	body := api.GetResponseBody(url)
	ratings := make([]Rating, 0)
	jsonErr := json.Unmarshal(body, &ratings)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return ratings
}
