package repository

import "sync"

type DB struct {
	// [string]Status
	Websites sync.Map
}

// Creating a in-memory database
func InitializeDatabase() *DB {
	return &DB{}
}
