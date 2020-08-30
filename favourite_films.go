package mubi

import (
	"encoding/json"
	"fmt"
	"log"
)

type FavouriteFilm struct {
	Timestamp int64 `json:"created_at"`
	Fannable  struct {
		Film Film `json:film`
	} `json:fannable`
}

func (api *MubiAPI) GetFavouriteFilms(userId int64, page int, perPage int) []FavouriteFilm {
	url := fmt.Sprintf(
		"https://mubi.com/services/api/favourites/films?user_id=%d&page=%d&per_page=%d",
		userId, page, perPage,
	)

	body := api.GetResponseBody(url)
	favourites := make([]FavouriteFilm, 0)
	jsonErr := json.Unmarshal(body, &favourites)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return favourites
}
