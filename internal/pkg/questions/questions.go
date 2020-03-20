package questions

import (
	"fmt"

	"gitlab.com/psheets/ddgquery"
)

// GetAnswer tries to get an answer or definition for user input from DDG
func GetAnswer(arg string) string {
	var answer string = "Failed to get an answer for that, please check my logs"

	results, _ := ddgquery.Query(arg, 1)
	if len(results) > 0 {
		answer = results[0].Info
		if results[0].Ref != "" {
			answer = fmt.Sprintf("%s [%s]", answer, results[0].Ref)
		}
	}

	return answer
}
