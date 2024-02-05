package db

import (
	"time"

	"github.com/Suhaan-Bhandary/website-checker/internal/pkg/helpers"
)

func GetWebsiteStatus(website string) WebsitesStatus {
	status, ok := database.Websites.Load(website)
	if !ok {
		return WebsitesStatus{
			Status:      helpers.ERROR,
			LastFetched: time.Now().Format("01-02-2006 15:04:05"),
		}
	}

	return status.(WebsitesStatus)
}

func GetAllStatus() map[string]WebsitesStatus {
	status := map[string]WebsitesStatus{}

	database.Websites.Range(func(key any, value any) bool {
		status[key.(string)] = value.(WebsitesStatus)
		return true
	})

	return status
}

func UpdateAllWebsiteStatus() {
	database.Websites.Range(func(key any, _ any) bool {
		go func() {
			status, err := helpers.GetWebsiteStatus(key.(string))
			if err != nil {
				status = helpers.ERROR
			}

			database.Websites.Store(key, WebsitesStatus{
				Status:      status,
				LastFetched: time.Now().Format("01-02-2006 15:04:05"),
			})
		}()

		return true
	})
}
