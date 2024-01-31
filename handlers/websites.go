package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Suhaan-Bhandary/website-checker/db"
	"github.com/Suhaan-Bhandary/website-checker/types"
)

func WebsitePostHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody types.WebsitePostBody

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body."))
		return
	}

	if len(requestBody.Websites) == 0 {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Atleast one website required"))
		return
	}

	db.InsertWebsites(requestBody.Websites)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Registered the given websites"))
}

func WebsiteGetHandler(w http.ResponseWriter, _ *http.Request) {
	websites := db.GetWebsites()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(types.WebsitesGetResponseBody{
		Message:  "Websites in DB",
		Websites: websites,
	})
}

func WebsiteStatusGetHandler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Website Status"))
}

func WebsiteRemoveHandler(w http.ResponseWriter, r *http.Request) {
	website := r.URL.Query().Get("website")

	if website == "" {
		fmt.Println("Please provide website / all in URL Query")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please provide website / all in URL Query"))
		return
	}

	err := db.DeleteWebsite(website)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted website successfully"))
}
