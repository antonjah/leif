package lunches

import (
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"

	"github.com/antonjah/leif/internal/pkg/constants"
	"github.com/antonjah/leif/internal/pkg/utils"

	"github.com/PuerkitoBio/goquery"
)

var dayMap = map[int]string{
	1: "måndag",
	2: "tisdag",
	3: "onsdag",
	4: "torsdag",
	5: "fredag",
	6: "lördag",
	7: "söndag",
}

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

	weekday := int(time.Now().Weekday())

	var lunches []string

	doc.Find(".w-restaurant").Each(func(i int, s *goquery.Selection) {
		name := "\n" + s.Find(".i-restaurant.name").Find(".heading").Text()
		var lunchStr = name + "\n"
		s.Find(".i-restaurant.menu").Find(".dish").Each(func(i int, se *goquery.Selection) {
			if v, _ := se.Attr("data-day"); v == dayMap[weekday] {
				lunchStr += se.Find(".name.ckeditor.text-only-editor").Text() + "\n"
			}
		})
		lunches = append(lunches, lunchStr)
	})

	return lunches
}
