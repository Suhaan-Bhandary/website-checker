package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Suhaan-Bhandary/website-checker/db"
	"github.com/Suhaan-Bhandary/website-checker/routes"
	"github.com/Suhaan-Bhandary/website-checker/utils"
)

const SERVER_ADDRESS = "127.0.0.1:8080"

func main() {
	fmt.Println("Starting Server...")

	// Fetch details
	go func() {
		for {
			// Update the map list
			websites := db.GetWebsites()
			status := utils.GetAllWebsiteStatus(websites)
			db.AssignStatus(status)

			time.Sleep(time.Minute)
		}
	}()

	// Listening to the server and assigning our custom router
	router := routes.Router()
	err := http.ListenAndServe(SERVER_ADDRESS, router)

	if err != nil {
		fmt.Println(err)
		return
	}
}
