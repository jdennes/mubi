package mubi

type User struct {
	Id           int    `json:id`
	Name         string `json:name`
	CanonicalUrl string `json:"canonical_url"`
	AvartarUrl   string `json:"avatar_url"`
	Bio          string `json:bio`
	Subscriber   bool   `json:subscriber`
}
