package corona

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"

	"github.com/antonjah/leif/internal/pkg/constants"
)

type countryStatus struct {
	Country      string `json:"country"`
	Province     string `json:"province"`
	Confirmed    int    `json:"confirmed"`
	Recovered    int    `json:"recovered"`
	Deaths       int    `json:"deaths"`
	MortalityPer string `json:"mortalityPer"`
	RecoveredPer string `json:"recoveredPer"`
	LastUpdated  string `json:"lastUpdated"`
}

func (s *countryStatus) toString() string {
	if s.Province != "" {
		return fmt.Sprintf(
			"Province: %s\nConfirmed: %d\nRecovered: %d\nDeaths: %d\nMortality Rate: %s%%\nRecover Rate: %s%%\nLast Updated: %s\n\n",
			s.Province, s.Confirmed, s.Recovered, s.Deaths, s.MortalityPer, s.RecoveredPer, s.LastUpdated,
		)
	} else {
		return fmt.Sprintf(
			"Confirmed: %d\nRecovered: %d\nDeaths: %d\nMortality Rate: %s%%\nRecover Rate: %s%%\nLast Updated: %s\n\n",
			s.Confirmed, s.Recovered, s.Deaths, s.MortalityPer, s.RecoveredPer, s.LastUpdated,
		)
	}
}

func GetStatuses(country string, cache cache.Cache, logger *logrus.Logger) string {
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
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error(err)
		return "Failed to get corona status, please check my logs"
	}
	defer r.Body.Close()

	var countryStatuses []countryStatus
	if err := json.Unmarshal(b, &countryStatuses); err != nil {
		return "Country not found"
	}

	var cs string
	for _, s := range countryStatuses {
		cs = cs + s.toString()
	}

	cache.Set(country, cs, 24*time.Hour)

	return cs
}
