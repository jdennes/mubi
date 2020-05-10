package mubi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	favourites := make([]FavouriteFilm, 0)
	jsonErr := json.Unmarshal(body, &favourites)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return favourites
}
