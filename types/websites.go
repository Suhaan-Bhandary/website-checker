package types

type WebsitePostBody struct {
	Websites []string `json:"websites"`
}

type WebsitesGetResponseBody struct {
	Message  string   `json:"message"`
	Websites []string `json:"websites"`
}

type WebsiteStatus struct {
	Website string `json:"website"`
	Status  string `json:"status"`
}

type WebsiteStatusResponse struct {
	Message string `json:"message"`
	Website string `json:"website"`
	Status  Status `json:"status"`
}

type AllWebsiteStatus map[string]Status

type AllWebsiteStatusResponse struct {
	Message string           `json:"message"`
	Status  AllWebsiteStatus `json:"statusList"`
}
