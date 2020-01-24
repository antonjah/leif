package utils

import (
	"math/rand"
	"strings"
	"time"
)

// FindInSlice searches for a string in a slice and returns all matches
func FindInSlice(searchArgument string, slice []string) []string {
	var found []string

	for _, b := range slice {
		if strings.Contains(strings.ToUpper(b), strings.ToUpper(searchArgument)) {
			found = append(found, b)
		}
	}

	return found
}

// GetRandom returns a random index from a slice
func GetRandom(stringSlice []string) string {
	rand.Seed(time.Now().Unix())
	return stringSlice[rand.Intn(len(stringSlice))]
}