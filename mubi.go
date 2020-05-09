package mubi

import (
	"net/http"
)

type MubiAPI struct {
	Client *http.Client
}
