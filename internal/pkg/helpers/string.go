package helpers

import "strings"

func CleanString(str string) string {
	return strings.Trim(str, " \n")
}
