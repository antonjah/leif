package lunches

import (
	"net/http"
	"regexp"
	"time"

	"github.com/antonjah/gleif/internal/constants"

	"github.com/PuerkitoBio/goquery"
	"github.com/patrickmn/go-cache"

	"github.com/antonjah/gleif/internal/utils"
)

// LunchHandler provides methods to lunch requests
type LunchHandler struct {
	Cache cache.Cache
}

// NewLunchHandler returns a new LunchHandler
func NewLunchHandler(cache cache.Cache) LunchHandler {
	return LunchHandler{Cache: cache}
}

// GetAll returns all daily lunches
func (h LunchHandler) GetAll() ([]string, error) {
	cachedLunches, found := h.Cache.Get("lunches")
	if found {
		lunches := cachedLunches.([]string)
		return lunches, nil
	}

	response, err := http.Get(constants.DAGENSLUNCHURL)
	if err != nil {
		return []string{}, err
	}
	defer response.Body.Close()

	lunches, err := lunchListFromResponse(response)
	if err != nil {
		return []string{}, err
	}

	h.Cache.Set("lunches", lunches, 1*time.Hour)

	return lunches, nil
}

// Search tries to find a specific keyword in the name of all restaurants
// or in all of their menus
func (h LunchHandler) Search(searchArgument string) ([]string, error) {
	lunches, err := h.GetAll()
	if err != nil {
		return []string{}, err
	}

	matches := utils.FindInSlice(searchArgument, lunches)

	return matches, nil
}

func lunchListFromResponse(response *http.Response) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return []string{}, err
	}

	var lunches []string

	doc.Find(".card-body.pt-2").Each(func(i int, s *goquery.Selection) {
		s.Attr("class")
		re := regexp.MustCompile(`\s+`)
		lunchRow := re.ReplaceAllString(s.Text(), " ")
		lunches = append(lunches, lunchRow)
	})

	return lunches, nil
}
