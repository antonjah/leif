package lunches

import (
	"fmt"
	"net/http"
	"strings"
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

	doc.Find(".lunchMenuListItem").Each(func(i int, restaurantSel *goquery.Selection) {
		var name = strings.TrimSpace(restaurantSel.Find(".lunchMenuListItem__restaurantName").Text())

		var dishList []string
		restaurantSel.Find(".dish").Each(func(i int, dishSel *goquery.Selection) {
			dishSel.Find(".dish__topRow").Each(func(i int, dishInfoSel *goquery.Selection) {
				dish := strings.ReplaceAll(strings.TrimSpace(dishInfoSel.Find(".dish__name").Text()), "\t", "")
				additional := dishSel.Find(".dish__bottomRow").Text()
				dishList = append(dishList, fmt.Sprintf("* %s %s", dish, additional))
			})
		})

		formatted := fmt.Sprintf("\n*%s*\n%s", name, strings.Join(dishList, "\n"))
		lunches = append(lunches, formatted)
	})

	return lunches
}
