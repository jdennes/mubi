package mubi

type Film struct {
	Title        string `json:title`
	CanonicalUrl string `json:"canonical_url"`
}
