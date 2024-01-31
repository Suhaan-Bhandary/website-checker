package db

type DB struct {
	websites map[string]struct{}
}

// Creating a in-memory database
var database = DB{
	websites: map[string]struct{}{},
}
