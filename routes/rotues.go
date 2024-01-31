package routes

import "github.com/gorilla/mux"

func UseRoute(prefix string, subRouterPathFunc func(router *mux.Router), router *mux.Router) {
	subRouter := router.PathPrefix(prefix).Subrouter()
	subRouterPathFunc(subRouter)
}

func Router() *mux.Router {
	router := mux.NewRouter()

	// Register routes
	UseRoute("/websites", WebsiteRoutes, router)

	return router
}
