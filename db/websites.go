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
		database.Websites[website] = types.Status{
			Status:      utils.NOT_FETCHED,
			LastFetched: time.Now().Format("01-02-2006 15:04:05"),
		}
	}
}

func GetWebsites() []string {
	websitesURL := make([]string, 0, len(database.Websites))
	for website := range database.Websites {
		websitesURL = append(websitesURL, website)
	}

	return websitesURL
}

func GetWebsiteStatus(website string) types.Status {
	return database.Websites[website]
}

func GetAllStatus() map[string]types.Status {
	return database.Websites
}

func DeleteWebsite(website string) error {
	if website == "all" {
		// clearing the map
		for key := range database.Websites {
			delete(database.Websites, key)
		}
		return nil
	}

	_, ok := database.Websites[website]
	if !ok {
		return errors.New("Website not found")
	}

	// Removing website from websites
	delete(database.Websites, website)

	return nil
}

func IsWebsitePresent(website string) bool {
	_, ok := database.Websites[website]
	return ok
}

func AssignStatus(status map[string]types.Status) {
	database.Websites = status
}
