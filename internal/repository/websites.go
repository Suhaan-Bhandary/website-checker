package repository

// A interface which defines all the function of websites
type WebsitesStorer interface{}

type WebsitesStatus struct {
	Status      string
	LastFetched string
}
