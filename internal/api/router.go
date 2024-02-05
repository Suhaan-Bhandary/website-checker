package api

import (
	"net/http"

	"github.com/Suhaan-Bhandary/website-checker/internal/app"
	"github.com/gorilla/mux"
)

func NewRouter(deps app.Dependencies) *mux.Router {
	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/websites", AddWebsitesHandler(deps.WebsiteService)).Methods(http.MethodPost)
	router.HandleFunc("/websites", GetWebsitesHandler(deps.WebsiteService)).Methods(http.MethodGet)
	router.HandleFunc("/websites", DeleteWebsiteHandler(deps.WebsiteService)).Methods(http.MethodDelete)
	router.HandleFunc("/websites/status", GetWebsiteStatusHandler(deps.WebsiteService)).Methods(http.MethodGet)

	return router
}
