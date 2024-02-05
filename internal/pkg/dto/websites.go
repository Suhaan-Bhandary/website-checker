package dto

import "errors"

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
