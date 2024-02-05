package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Suhaan-Bhandary/website-checker/internal/app/websites"
	"github.com/Suhaan-Bhandary/website-checker/internal/pkg/dto"
)

func AddWebsitesHandler(websitesSvc websites.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implement this
		var req dto.AddWebsiteRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid request body."))
			return
		}

		err = req.Validate()
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		websitesSvc.InsertWebsites(req.Websites)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Registered the given websites"))
	}
}

func GetWebsitesHandler(websitesSvc websites.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// get website
		websites := websitesSvc.GetWebsites()

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(dto.GetWebsitesResponse{
			Message:  "Websites in DB",
			Websites: websites,
		})
	}
}
