package mubi

import (
	"encoding/json"
	"fmt"
	"log"
)

type List struct {
	Id             int    `json:id`
	Title          string `json:title`
	ListFilmsCount int    `json:"list_films_count"`
	CreatedAt      int64  `json:"created_at"`
	UpdatedAt      int64  `json:"updated_at"`
	CanonicalUrl   string `json:"canonical_url"`
}

func (api *MubiAPI) GetLists(userId int64, page int, perPage int) []List {
	url := fmt.Sprintf(
		"https://mubi.com/services/api/lists?user_id=%d&sort=updated_at&page=%d&per_page=%d",
		userId, page, perPage,
	)

	body := api.GetResponseBody(url)
	lists := make([]List, 0)
	jsonErr := json.Unmarshal(body, &lists)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return lists
}
