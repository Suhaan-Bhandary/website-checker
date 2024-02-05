package repository

// A interface which defines all the function of websites
type WebsitesStorer interface {
	InsertWebsite(websites string)
	GetWebsites() []string

	DeleteWebsite(website string)
	DeleteAllWebsites()

	IsWebsitePresent(website string) bool

	GetWebsiteStatus(website string) WebsitesStatus
	GetAllWebsiteStatus() map[string]WebsitesStatus

	UpdateAllWebsiteStatus()
}

const UP = "UP"
const DOWN = "DOWN"
const ERROR = "ERROR"
const NOT_FETCHED = "NOT_FETCHED"

type WebsitesStatus struct {
	Status      string
	LastFetched string
}
