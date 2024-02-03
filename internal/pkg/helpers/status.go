package helpers

import (
	"fmt"
	"net/http"
)

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
