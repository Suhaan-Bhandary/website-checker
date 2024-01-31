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
	w.Write([]byte("Website Get"))
}

func WebsiteStatusGetHandler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Website Status"))
}

func WebsiteRemoveHandler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Website Remove"))
}
