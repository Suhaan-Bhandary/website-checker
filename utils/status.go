package utils

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/Suhaan-Bhandary/website-checker/types"
)

/*
	Status: UP, DOWN, ERROR, NOT_FETCHED
*/

const UP = "UP"
const DOWN = "DOWN"
const ERROR = "ERROR"
const NOT_FETCHED = "NOT_FETCHED"

func GetWebsiteStatus(website string) (string, error) {
	// check if website is present or not
	resp, err := http.Get(website)
	if err != nil {
		return "", err
	}

	statusCode := fmt.Sprint(resp.StatusCode)
	if statusCode[0] != '2' {
		return DOWN, nil
	}

	return UP, nil
}

func GetAllWebsiteStatus(websites []string) types.AllWebsiteStatus {
	statusMap := map[string]types.Status{}

	wg := sync.WaitGroup{}
	wg.Add(len(websites))

	for _, website := range websites {
		// creating a go routine
		go func(website string) {
			defer wg.Done()

			status, err := GetWebsiteStatus(website)
			if err != nil {
				status = ERROR
			}

			statusMap[website] = types.Status{
				Status:      status,
				LastFetched: time.Now().Format("01-02-2006 15:04:05"),
			}
		}(website)
	}

	wg.Wait()

	return statusMap
}
