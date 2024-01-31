package types

type WebsitePostBody struct {
	Websites []string `json:"websites"`
}

type WebsitesGetResponseBody struct {
	Message  string   `json:"message"`
	Websites []string `json:"websites"`
}
