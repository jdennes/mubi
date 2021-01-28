package mubi

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

// The only way I know how to get a list of films of the day at the moment is
// to scrape some JSON out of https://mubi.com/film-of-the-day

type FilmOfTheDay struct {
	Id          int    `json:"id"`
	AvailableAt string `json:"available_at"`
	ExpiresAt   string `json:"expires_at"`
	Film        Film   `json:"film"`
}

type FilmsOfTheDayData struct {
	Properties struct {
		InitialState struct {
			FilmOfTheDayContainer struct {
				FilmsOfTheDay []FilmOfTheDay `json:"filmProgrammings"`
			} `json:"filmProgramming"`
		} `json:"initialState"`
	} `json:"props"`
}

func (api *MubiAPI) GetFilmsOfTheDay() []FilmOfTheDay {
	res, requestErr := http.Get("https://mubi.com/film-of-the-day")
	if requestErr != nil {
		log.Fatal(requestErr)
	}

	defer res.Body.Close()
	doc, loadErr := goquery.NewDocumentFromReader(res.Body)
	if loadErr != nil {
		log.Fatal(loadErr)
	}

	// Find the <script id="__NEXT_DATA__" ...> element containing the JSON
	element := doc.Find("script[id='__NEXT_DATA__']")
	jsonContent := element.Text()

	var filmsData FilmsOfTheDayData
	jsonErr := json.Unmarshal([]byte(jsonContent), &filmsData)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return filmsData.Properties.InitialState.FilmOfTheDayContainer.FilmsOfTheDay
}
