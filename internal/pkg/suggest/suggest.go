package suggest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

func GetSuggestion(logger *logrus.Logger) string {
	var result = ""
	var words []string
	resp, err := http.Get("RandomWordApiURL")
	if err != nil {
		logger.Error(err)
		fmt.Println("Failed randomly to get some random words")
	}
	page, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err)
		fmt.Println("Failed randomly to get some random words")
	}
	err = json.Unmarshal([]byte(page), &words)
	if err != nil {
		logger.Error(err)
		fmt.Println("Failed to parse random word json")
	}
	result = fmt.Sprintf("How about you use '%s' ?", strings.Title(words[0])+strings.Title(words[1]))
	fmt.Println(result)
	return result
}
