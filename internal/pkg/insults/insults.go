package insults

import (
	"html"
	"io/ioutil"
	"net/http"

	"github.com/antonjah/leif/internal/pkg/constants"
)

// Get returns an insult from the evil insult API
func Get() (string, error) {
	r, err := http.Get(constants.InsultURL)
	if err != nil {
		return "Failed to get an insult, please check my logs", err
	}
	defer r.Body.Close()

	b, _ := ioutil.ReadAll(r.Body)
	return html.UnescapeString(string(b)), nil
}
