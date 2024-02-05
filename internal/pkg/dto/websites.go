package dto

import (
	"errors"

	"github.com/Suhaan-Bhandary/website-checker/internal/repository"
)

type AddWebsiteRequest struct {
	Websites []string `json:"websites"`
}

func (req *AddWebsiteRequest) Validate() error {
	if len(req.Websites) == 0 {
		return errors.New("At least one website required")
	}
	return nil
}

type GetWebsitesResponse struct {
	Message  string
	Websites []string
}

type AllWebsiteStatusResponse struct {
	Message string
	Status  map[string]repository.WebsitesStatus
}

type WebsiteStatusResponse struct {
	Message string
	Status  repository.WebsitesStatus
}
