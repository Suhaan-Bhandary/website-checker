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
	return func(w http.ResponseWriter, _ *http.Request) {
		// get website
		websites := websitesSvc.GetWebsites()

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(dto.GetWebsitesResponse{
			Message:  "Websites in DB",
			Websites: websites,
		})
	}
}

func DeleteWebsiteHandler(websitesSvc websites.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		website := r.URL.Query().Get("website")
		if website == "" {
			fmt.Println("Please provide website / all in URL Query")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Please provide website / all in URL Query"))
			return
		}

		err := websitesSvc.DeleteWebsite(website)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Deleted website successfully"))
	}
}

func GetWebsiteStatusHandler(websitesSvc websites.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		website := r.URL.Query().Get("website")

		// If website is not present return all website status
		if website == "" {
			websiteStatusList := websitesSvc.GetAllStatus()

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(dto.AllWebsiteStatusResponse{
				Message: "List of all website status.",
				Status:  websiteStatusList,
			})
			return
		}

		websiteStatus, err := websitesSvc.GetWebsiteStatus(website)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(dto.WebsiteStatusResponse{
			Message: "Websites in DB",
			Status:  websiteStatus,
		})
	}
}
