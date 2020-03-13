package f1

import (
	"fmt"

	"github.com/foobarmeow/go-ergast"
	"github.com/sirupsen/logrus"
)

func GetLatestResult(logger *logrus.Logger) string {
	r, err := ergast.Latest()
	if err != nil {
		logger.Error(err)
		return "Failed to get results, please check my logs"
	}

	result := fmt.Sprintf("%d %s - %d/%d\n\n", r.Season, r.RaceName, r.Date.Day(), r.Date.Month())

	for _, res := range r.Results {
		result = result + fmt.Sprintf("%d. %s %s %s (%s)\n", res.Position, res.Driver.GivenName, res.Driver.FamilyName, res.Status, res.FastestLap.Time)
	}

	return result
}
