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

type CoronaResponse struct {
	Data []CountryInfo `json:"data"`
}

type CountryInfo struct {
	Country             string `json:"country"`
	Cases               int    `json:"cases"`
	TodayCases          int    `json:"todayCases"`
	Deaths              int    `json:"deaths"`
	TodayDeaths         int    `json:"todayDeaths"`
	Recovered           int    `json:"recovered"`
	Active              int    `json:"active"`
	Critical            int    `json:"critical"`
	CasesPerOneMillion  int    `json:"casesPerOneMillion"`
	DeathsPerOneMillion int    `json:"deathsPerOneMillion"`
	Confirmed           int    `json:"confirmed"`
}

func (c *CoronaResponse) toString() string {
	return fmt.Sprintf(
		"Confirmed: %d\n"+
			"Cases: %d\n"+
			"Cases today: %d\n"+
			"Deaths: %d\n"+
			"Deaths today: %d\n"+
			"Recovered: %d\n"+
			"Active: %d\n"+
			"Critical: %d\n"+
			"Cases per million: %d\n"+
			"Deaths per million: %d",
		c.Data[0].Confirmed,
		c.Data[0].Cases,
		c.Data[0].TodayCases,
		c.Data[0].Deaths,
		c.Data[0].TodayDeaths,
		c.Data[0].Recovered,
		c.Data[0].Active,
		c.Data[0].Critical,
		c.Data[0].CasesPerOneMillion,
		c.Data[0].DeathsPerOneMillion,
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
	} else if r.StatusCode / 100 != 2 {
		logger.Error(r.Status)
		return "Failed to get corona status, please check my logs"
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error(err)
		return "Failed to get corona status, please check my logs"
	}
	defer r.Body.Close()

	var res CoronaResponse
	if err := json.Unmarshal(b, &res); err != nil {
		logger.Fatal(err)
		return "Failed to read corona status, please check my logs"
	}

	cache.Set(country, res.toString(), 24*time.Hour)

	return res.toString()
}
