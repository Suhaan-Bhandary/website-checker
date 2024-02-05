package websites

import (
	"github.com/Suhaan-Bhandary/website-checker/internal/pkg/helpers"
	"github.com/Suhaan-Bhandary/website-checker/internal/repository"
)

// This struct stores all the dependencies it needs from the repo layer
type service struct {
	websitesRepo repository.WebsitesStorer
}

// All the functions exposed to handler
type Service interface {
	InsertWebsites(websites []string)
}

func NewService(websitesRepo repository.WebsitesStorer) Service {
	return &service{
		websitesRepo: websitesRepo,
	}
}

func (os *service) InsertWebsites(websites []string) {
	for _, website := range websites {
		website := helpers.CleanString(website)
		os.websitesRepo.InsertWebsite(website)
	}
}
