package lunches

import (
	"net/http"
	"regexp"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"

	"github.com/antonjah/leif/internal/pkg/constants"
	"github.com/antonjah/leif/internal/pkg/utils"

	"github.com/PuerkitoBio/goquery"
)

// GetAll returns all daily lunches
func GetAll(cache cache.Cache, logger *logrus.Logger) []string {
	cachedLunches, found := cache.Get("lunches")
	if found {
		lunches := cachedLunches.([]string)
		return lunches
	}

	response, err := http.Get(constants.DagensLunchURL)
	if err != nil {
		logger.Error(err)
		return []string{"Failed to get lunches, please check my logs"}
	}
	defer response.Body.Close()

	lunches := lunchListFromResponse(response, logger)

	cache.Set("lunches", lunches, 1*time.Hour)

	return lunches
}

// Search tries to find a specific keyword in the name of all restaurants
// or in all of their menus
func Search(searchArgument string, cache cache.Cache, logger *logrus.Logger) ([]string, error) {
	lunches := GetAll(cache, logger)
	matches := utils.FindInSlice(searchArgument, lunches)

	return matches, nil
}

func lunchListFromResponse(response *http.Response, logger *logrus.Logger) []string {
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		logger.Error(err)
		return []string{"Failed to get lunches, please check my logs"}
	}

	var lunches []string

	doc.Find(".card-body.pt-2").Each(func(i int, s *goquery.Selection) {
		s.Attr("class")
		re := regexp.MustCompile(`\s+`)
		lunchRow := "-" + re.ReplaceAllString(s.Text(), " ")
		lunches = append(lunches, lunchRow)
	})

	return lunches
}
