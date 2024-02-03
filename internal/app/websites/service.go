package websites

import "github.com/Suhaan-Bhandary/website-checker/internal/repository"

// This struct stores all the dependencies it needs from the repo layer
type service struct {
	websitesRepo repository.WebsitesStorer
}

// All the functions exposed to handler
type Service interface{}

func NewService(websitesRepo repository.WebsitesStorer) Service {
	return &service{
		websitesRepo: websitesRepo,
	}
}
