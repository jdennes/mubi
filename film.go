package mubi

type Film struct {
	Title        string `json:title`
	Year         int    `json:year`
	CanonicalUrl string `json:"canonical_url"`
}
