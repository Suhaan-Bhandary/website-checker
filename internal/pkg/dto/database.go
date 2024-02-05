package dto

import "sync"

type DB struct {
	// [string]Status
	Websites sync.Map
}
