package tldr

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/antonjah/leif/internal/pkg/constants"
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
		return "Failed to get tldr, please check my logs"
	}

	switch resp.StatusCode {
	case http.StatusOK:
		page, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Error(err)
			return "Failed to get tldr, please check my logs"
		}
		return string(page)

	case http.StatusNotFound:
		return fmt.Sprintf("Found nothing on %s", arg)

	default:
		logger.Errorf("Failed to get tldr: %s", resp.Status)
		return "Failed to get tldr, please check my logs"
	}
}
