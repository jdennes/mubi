package mubi

import (
	"encoding/json"
	"log"
)

type LiveFilmScreening struct {
	Film struct {
		Title     string `json:title`
		Year      int    `json:year`
		WebUrl    string `json:"web_url"`
		Directors string `json:directors`
	} `json:"film_programming"`
	Description string `json:description`
}

func (api *MubiAPI) GetLiveFilmScreening() LiveFilmScreening {
	url := "https://mubi.com/live.json"
	body := api.GetResponseBody(url)
	screening := LiveFilmScreening{}
	jsonErr := json.Unmarshal(body, &screening)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return screening
}
