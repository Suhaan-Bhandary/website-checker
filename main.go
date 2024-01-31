package main

import (
	"fmt"
	"net/http"

	"github.com/Suhaan-Bhandary/website-checker/routes"
)

const SERVER_ADDRESS = "127.0.0.1:8080"

func main() {
	fmt.Println("Starting Server...")

	// Listening to the server and assigning our custom router
	router := routes.Router()
	http.ListenAndServe(SERVER_ADDRESS, router)
}
