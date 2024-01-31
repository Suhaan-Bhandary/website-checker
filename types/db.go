package types

type Status struct {
	Status      string
	LastFetched string
}

type DB struct {
	Websites map[string]Status
}
