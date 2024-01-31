package utils

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/Suhaan-Bhandary/website-checker/types"
)

/*
	Status: UP, DOWN, ERROR
*/

const UP = "UP"
const DOWN = "DOWN"
const ERROR = "ERROR"

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
	statusMap := map[string]string{}

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

			statusMap[website] = status
		}(website)
	}

	wg.Wait()

	return statusMap
}
