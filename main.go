package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title        string `json:title`
	CanonicalUrl string `json:"canonical_url"`
}

type Rating struct {
	Overall int `json:"overall"`
	Film    Film
}

type Ratings struct {
	Collection []Rating
}

func main() {
	userId := 7995037
	url := fmt.Sprintf("https://mubi.com/services/api/ratings?user_id=%d", userId)

	client := http.Client{
		Timeout: time.Second * 2,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "jdennes/mubi")
	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	ratings := make([]Rating, 0)
	jsonErr := json.Unmarshal(body, &ratings)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	for _, rating := range ratings {
		fmt.Printf("%s: %s\n", rating.Film.Title, rating.Film.CanonicalUrl)
		fmt.Println("Rating: ", rating.Overall)
		fmt.Println("-----")
	}
}
