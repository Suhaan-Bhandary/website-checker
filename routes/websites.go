package routes

import (
	"net/http"

	"github.com/Suhaan-Bhandary/website-checker/handlers"
	"github.com/gorilla/mux"
)

func WebsiteRoutes(router *mux.Router) {
	// Register all the website route here
	router.HandleFunc("", handlers.WebsitePostHandler).Methods(http.MethodPost)
	router.HandleFunc("", handlers.WebsiteGetHandler).Methods(http.MethodGet)
	router.HandleFunc("", handlers.WebsiteRemoveHandler).Methods(http.MethodDelete)
	router.HandleFunc("/status", handlers.WebsiteStatusGetHandler).Methods(http.MethodGet)
}
