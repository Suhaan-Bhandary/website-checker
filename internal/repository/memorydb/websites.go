package repository

import (
	"github.com/Suhaan-Bhandary/website-checker/internal/db"
	"github.com/Suhaan-Bhandary/website-checker/internal/pkg/helpers"
	"github.com/Suhaan-Bhandary/website-checker/internal/repository"
)

// A struct to define methods on it
// Since in memory db is a dependency for our repository methods we store it in our struct
type websitesStore struct {
	DB *db.DB
}

func NewWebsiteRepo(DB *db.DB) repository.WebsitesStorer {
	return &websitesStore{
		DB: DB,
	}
}

func (repo *websitesStore) InsertWebsite(website string) {
	repo.DB.Websites.Store(website, repository.WebsitesStatus{
		Status:      repository.NOT_FETCHED,
		LastFetched: helpers.GetCurrentTimeFormated(),
	})
}

func (repo *websitesStore) GetWebsites() []string {
	websitesURL := make([]string, 0)

	repo.DB.Websites.Range(func(key any, _ any) bool {
		websitesURL = append(websitesURL, key.(string))
		return true
	})

	return websitesURL
}
