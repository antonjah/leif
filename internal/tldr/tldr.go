package tldr

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/antonjah/leif/internal/constants"
)

// GetTLDR tries to get information about a command
// from the awesome TLDR.io
func GetTLDR(arg string, logger *logrus.Logger) string {
	req, _ := http.NewRequest("GET", fmt.Sprintf(constants.TLDRBaseURL, arg), nil)
	req.Header.Add("accept", "application/vnd.github.VERSION.raw")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error(err)
		return "I encountered a problem trying to get your tldr information, please check my logs"
	}

	switch resp.StatusCode {
	case http.StatusOK:
		page, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Error(err)
			return "I encountered a problem trying to read the tldr information, please check my logs"
		}
		return string(page)

	case http.StatusNotFound:
		return fmt.Sprintf("Found nothing on %s", arg)

	default:
		return "The GitHub API shit the bed, please check my logs"
	}
}
