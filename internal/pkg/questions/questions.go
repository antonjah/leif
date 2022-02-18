package questions

import (
	"github.com/sirupsen/logrus"
	"strings"

	"github.com/ajanicij/goduckgo/goduckgo"
)

var (
	errorMessage   = "Failed to get an answer for that, please check my logs"
	defaultMessage = "Sorry I don't know what that means"
)

// GetAnswer tries to get an answer or definition for user input from DDG
func GetAnswer(arg string, logger *logrus.Logger) string {
	arg = strings.TrimLeft(arg, " ")
	answer, err := goduckgo.Query(arg)
	if err != nil {
		logger.Error(err)
		return errorMessage
	}

	if answer.AbstractText == "" {
		return defaultMessage
	}

	return answer.AbstractText
}
