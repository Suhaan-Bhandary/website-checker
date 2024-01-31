package db

import (
	"errors"
	"time"

	"github.com/Suhaan-Bhandary/website-checker/types"
	"github.com/Suhaan-Bhandary/website-checker/utils"
)

func InsertWebsites(websites []string) {
	for _, website := range websites {
		website := utils.CleanString(website)
		database.Websites.Store(website, types.Status{
			Status:      utils.NOT_FETCHED,
			LastFetched: time.Now().Format("01-02-2006 15:04:05"),
		})
	}
}

func GetWebsites() []string {
	websitesURL := make([]string, 0)

	database.Websites.Range(func(key any, _ any) bool {
		websitesURL = append(websitesURL, key.(string))
		return true
	})

	return websitesURL
}

func GetWebsiteStatus(website string) types.Status {
	status, ok := database.Websites.Load(website)
	if !ok {
		return types.Status{
			Status:      utils.ERROR,
			LastFetched: time.Now().Format("01-02-2006 15:04:05"),
		}
	}

	return status.(types.Status)
}

func GetAllStatus() types.AllWebsiteStatus {
	status := types.AllWebsiteStatus{}

	database.Websites.Range(func(key any, value any) bool {
		status[key.(string)] = value.(types.Status)
		return true
	})

	return status
}

func DeleteWebsite(website string) error {
	if website == "all" {
		// clearing the map
		database.Websites.Range(func(key any, _ any) bool {
			database.Websites.Delete(key)
			return true
		})

		return nil
	}

	_, ok := database.Websites.Load(website)
	if !ok {
		return errors.New("Website not found")
	}

	// Removing website from websites
	database.Websites.Delete(website)

	return nil
}

func IsWebsitePresent(website string) bool {
	_, ok := database.Websites.Load(website)
	return ok
}

func UpdateAllWebsiteStatus() {
	database.Websites.Range(func(key any, _ any) bool {
		go func() {
			status, err := utils.GetWebsiteStatus(key.(string))
			if err != nil {
				status = utils.ERROR
			}

			database.Websites.Store(key, types.Status{
				Status:      status,
				LastFetched: time.Now().Format("01-02-2006 15:04:05"),
			})
		}()

		return true
	})
}
