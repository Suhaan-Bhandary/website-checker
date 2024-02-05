package websites

import (
	"context"
	"errors"
	"time"

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

	GetAllStatus() map[string]repository.WebsitesStatus
	GetWebsiteStatus(website string) (repository.WebsitesStatus, error)

	StatusUpdateBackgroundJob(ctx context.Context)
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

func (os *service) GetAllStatus() map[string]repository.WebsitesStatus {
	return os.websitesRepo.GetAllWebsiteStatus()
}

func (os *service) GetWebsiteStatus(website string) (repository.WebsitesStatus, error) {
	if ok := os.websitesRepo.IsWebsitePresent(website); !ok {
		return repository.WebsitesStatus{}, errors.New("Website not found")
	}

	websiteStatus := os.websitesRepo.GetWebsiteStatus(website)
	return websiteStatus, nil
}

func (os *service) StatusUpdateBackgroundJob(ctx context.Context) {
	for {
		os.websitesRepo.UpdateAllWebsiteStatus()
		time.Sleep(time.Minute)
	}
}
