package repository

import "github.com/Suhaan-Bhandary/website-checker/internal/db"

// Creating a in-memory database
func InitializeDatabase() *db.DB {
	return &db.DB{}
}
