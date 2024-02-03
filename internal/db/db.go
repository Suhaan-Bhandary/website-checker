package db

import "sync"

type DB struct {
	// [string]Status
	Websites sync.Map
}

type WebsitesStatus struct {
	Status      string
	LastFetched string
}

// Creating a in-memory database
var database = DB{}
