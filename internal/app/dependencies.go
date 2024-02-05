package app

import (
	"github.com/Suhaan-Bhandary/website-checker/internal/app/websites"
	db "github.com/Suhaan-Bhandary/website-checker/internal/repository"
	repository "github.com/Suhaan-Bhandary/website-checker/internal/repository/memorydb"
)

type Dependencies struct {
	WebsiteService websites.Service
}

func NewServices(db *db.DB) Dependencies {
	websiteRepo := repository.NewWebsiteRepo(db)
	websitesService := websites.NewService(websiteRepo)

	return Dependencies{
		WebsiteService: websitesService,
	}
}
