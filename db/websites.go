package db

import (
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
