package mubi

type Rating struct {
	Overall int `json:"overall"`
	Film    Film
}

type Ratings struct {
	Collection []Rating
}
