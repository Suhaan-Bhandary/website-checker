package websites

import (
	"errors"

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
	GetWebsites() []string
	DeleteWebsite(website string) error
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

func (os *service) GetWebsites() []string {
	return os.websitesRepo.GetWebsites()
}

func (os *service) DeleteWebsite(website string) error {
	if website == "all" {
		os.websitesRepo.DeleteAllWebsites()
		return nil
	}

	if ok := os.websitesRepo.IsWebsitePresent(website); !ok {
		return errors.New("Website not found")
	}

	os.websitesRepo.DeleteWebsite(website)
	return nil
}
