package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Suhaan-Bhandary/website-checker/internal/api"
	"github.com/Suhaan-Bhandary/website-checker/internal/app"
	"github.com/Suhaan-Bhandary/website-checker/internal/repository"
)

const SERVER_ADDRESS = "127.0.0.1:3000"

func main() {
	context := context.Background()

	fmt.Println("Starting Website Checker Server...")
	defer fmt.Println("Shutting Down Website Checker Server...")

	// Getting the db reference
	memoryDB := repository.InitializeDatabase()

	// Passing the dependency to services
	services := app.NewServices(memoryDB)

	// passing the services to handler
	router := api.NewRouter(services)

	// Fetch details
	go services.WebsiteService.StatusUpdateBackgroundJob(context)

	// Listening to the server and assigning our custom router
	err := http.ListenAndServe(SERVER_ADDRESS, router)
	if err != nil {
		fmt.Println(err)
		return
	}
}
