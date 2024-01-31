package db

import (
	"errors"

	"github.com/Suhaan-Bhandary/website-checker/utils"
)

func InsertWebsites(websites []string) {
	for _, website := range websites {
		website := utils.CleanString(website)
		database.websites[website] = struct{}{}
	}
}

func GetWebsites() []string {
	websitesURL := make([]string, 0, len(database.websites))
	for website := range database.websites {
		websitesURL = append(websitesURL, website)
	}

	return websitesURL
}

func DeleteWebsite(website string) error {
	if website == "all" {
		// clearing the map
		for key := range database.websites {
			delete(database.websites, key)
		}
		return nil
	}

	_, ok := database.websites[website]
	if !ok {
		return errors.New("Website not found")
	}

	// Removing website from websites
	delete(database.websites, website)

	return nil
}

func IsWebsitePresent(website string) bool {
	_, ok := database.websites[website]
	return ok
}
