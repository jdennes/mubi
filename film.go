package mubi

type Director struct {
	Id   int    `json:id`
	Name string `json:name`
}

type Film struct {
	Title        string     `json:title`
	Year         int        `json:year`
	CanonicalUrl string     `json:"canonical_url"`
	WebUrl       string     `json:"web_url"`
	Directors    []Director `json:directors`
}
