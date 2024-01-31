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
