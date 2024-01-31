package handlers

import "net/http"

func WebsitePostHandler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Website Post"))
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
