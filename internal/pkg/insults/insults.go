package insults

import (
	"html"
	"io/ioutil"
	"net/http"
)

// Get returns an insult from the evil insult API
func Get() (string, error) {
	r, err := http.Get("https://evilinsult.com/generate_insult.php?lang=en&type=text")
	if err != nil {
		return "The insult API isn't working right now, you'll have to do it yourself", err
	}
	defer r.Body.Close()

	b, _ := ioutil.ReadAll(r.Body)
	return html.UnescapeString(string(b)), nil
}
