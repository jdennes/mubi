package mubi

import (
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
