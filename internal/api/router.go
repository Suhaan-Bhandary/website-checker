package api

import (
	"net/http"

	"github.com/Suhaan-Bhandary/website-checker/internal/app"
	"github.com/gorilla/mux"
)

func NewRouter(deps app.Dependencies) *mux.Router {
	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/", func(w http.ResponseWriter,
		r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	return router
}
