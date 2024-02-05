package repository

import (
	"time"

	"github.com/Suhaan-Bhandary/website-checker/internal/pkg/helpers"
	"github.com/Suhaan-Bhandary/website-checker/internal/repository"
)

// A struct to define methods on it
// Since in memory db is a dependency for our repository methods we store it in our struct
type websitesStore struct {
	DB *repository.DB
}

func NewWebsiteRepo(DB *repository.DB) repository.WebsitesStorer {
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

func (repo *websitesStore) DeleteWebsite(website string) {
	repo.DB.Websites.Delete(website)
}

func (repo *websitesStore) DeleteAllWebsites() {
	// clearing the map
	repo.DB.Websites.Range(func(key any, _ any) bool {
		repo.DB.Websites.Delete(key)
		return true
	})
}

func (repo *websitesStore) IsWebsitePresent(website string) bool {
	_, ok := repo.DB.Websites.Load(website)
	return ok
}

func (repo *websitesStore) GetWebsiteStatus(website string) repository.WebsitesStatus {
	status, ok := repo.DB.Websites.Load(website)
	if !ok {
		return repository.WebsitesStatus{
			Status:      helpers.ERROR,
			LastFetched: time.Now().Format("01-02-2006 15:04:05"),
		}
	}

	return status.(repository.WebsitesStatus)
}

func (repo *websitesStore) GetAllWebsiteStatus() map[string]repository.WebsitesStatus {
	status := map[string]repository.WebsitesStatus{}

	repo.DB.Websites.Range(func(key any, value any) bool {
		status[key.(string)] = value.(repository.WebsitesStatus)
		return true
	})

	return status
}

func (repo *websitesStore) UpdateAllWebsiteStatus() {
	repo.DB.Websites.Range(func(key any, _ any) bool {
		go func() {
			status, err := helpers.GetWebsiteStatus(key.(string))
			if err != nil {
				status = helpers.ERROR
			}

			repo.DB.Websites.Store(key, repository.WebsitesStatus{
				Status:      status,
				LastFetched: time.Now().Format("01-02-2006 15:04:05"),
			})
		}()

		return true
	})
}
