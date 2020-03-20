package postmord

import (
	"fmt"
	"strings"

	"github.com/antonjah/postmord"
	log "github.com/sirupsen/logrus"
)

func Search(identifier string, logger *log.Logger, token string) string {
	c := postmord.NewClient(token, nil, nil)
	r, err := c.FindByIdentifier(strings.TrimSpace(identifier))
	if err != nil {
		logger.Error(err)
		return "Failed to get status, please check my logs"
	}

	answer := ""
	for _, shipment := range r.Tracking.Shipments {
		answer = answer + fmt.Sprintf(
			"%s - %s - %s\n\n",
			shipment.Consignor.Name,
			shipment.StatusText.Body,
			shipment.StatusText.EstimatedTimeOfArrival,
		)
	}

	if answer == "" {
		return "No statuses found for your parcel"
	}

	return answer
}
