package utils

import "strings"

func FindInSlice(searchArgument string, slice []string) []string {
	var found []string

	for _, b := range slice {
		if strings.Contains(strings.ToUpper(b), strings.ToUpper(searchArgument)) {
			found = append(found, b)
		}
	}

	return found
}