package suggest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/antonjah/leif/internal/pkg/constants"
	"github.com/sirupsen/logrus"
)

func GetSuggestion(logger *logrus.Logger) string {
	var words []string
	resp, err := http.Get(constants.RandomWordApiURL)
	if err != nil {
		logger.Error(err)
		return "Failed to get random word, please check my logs"
	}
	page, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err)
		return "Failed to get random word, please check my logs"
	}
	err = json.Unmarshal([]byte(page), &words)
	if err != nil {
		logger.Error(err)
		return "Failed to get random word, please check my logs"
	}
	result := fmt.Sprintf("How about you use '%s'?", strings.Title(words[0])+strings.Title(words[1]))
	return result
}
