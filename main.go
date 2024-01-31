package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Suhaan-Bhandary/website-checker/db"
	"github.com/Suhaan-Bhandary/website-checker/routes"
)

const SERVER_ADDRESS = "127.0.0.1:8080"

func StatusUpdateBackgroundJob() {
	for {
		db.UpdateAllWebsiteStatus()
		time.Sleep(10 * time.Second)
	}
}

func main() {
	fmt.Println("Starting Server...")

	// Fetch details
	go StatusUpdateBackgroundJob()

	// Listening to the server and assigning our custom router
	router := routes.Router()
	err := http.ListenAndServe(SERVER_ADDRESS, router)

	if err != nil {
		fmt.Println(err)
		return
	}
}
