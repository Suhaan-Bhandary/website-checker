package types

import "sync"

type Status struct {
	Status      string
	LastFetched string
}

type DB struct {
	// [string]Status
	Websites sync.Map
}
