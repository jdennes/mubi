package mubi

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type MubiAPI struct {
	Client *http.Client
}

func NewMubiAPI() *MubiAPI {
	client := http.Client{
		Timeout: time.Second * 5,
	}
	return &MubiAPI{Client: &client}
}

func (api *MubiAPI) GetResponseBody(url string) []byte {
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

	return body
}
