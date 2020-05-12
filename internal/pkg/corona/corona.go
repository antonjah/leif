package corona

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"

	"github.com/antonjah/leif/internal/pkg/constants"
)

type CountryInfo struct {
	Confirmed   int    `json:"Confirmed"`
	Deaths      int    `json:"Deaths"`
	Recovered   int    `json:"Recovered"`
	Active      int    `json:"Active"`
	LastUpdated string `json:"Date"`
}

func (c *CountryInfo) toString() string {
	return fmt.Sprintf(
		"Confirmed: %d\n"+
			"Deaths: %d\n"+
			"Recovered: %d\n"+
			"Active: %d\n"+
			"Last updated: %s",
		c.Confirmed,
		c.Deaths,
		c.Recovered,
		c.Active,
		c.LastUpdated,
	)
}

func GetStatuses(country string, cache cache.Cache, logger *logrus.Logger) string {
	country = strings.ToLower(country)
	cachedCountry, found := cache.Get(country)
	if found {
		cached := cachedCountry.(string)
		return cached
	}

	c := http.DefaultClient
	r, err := c.Get(fmt.Sprintf(constants.CoronaBaseURL, country))
	if err != nil {
		logger.Error(err)
		return "Failed to get corona status, please check my logs"
	} else if r.StatusCode == http.StatusNotFound {
		return "Country not found"
	} else if r.StatusCode/100 != 2 {
		logger.Error(r.Status)
		return "Failed to get corona status, please check my logs"
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error(err)
		return "Failed to get corona status, please check my logs"
	}
	defer r.Body.Close()

	var res []CountryInfo
	if err := json.Unmarshal(b, &res); err != nil {
		logger.Fatal(err)
		return "Failed to read corona status, please check my logs"
	}

	cache.Set(country, res[len(res)-1].toString(), 24*time.Hour)

	return res[len(res)-1].toString()
}
