package app

import (
	"strings"

	"github.com/nlopes/slack"
	"github.com/sirupsen/logrus"
)

// HandleResponse type checks the input message and sends the
// response to specified slack channel
func HandleResponse(message interface{}, event *slack.MessageEvent, client *slack.Client, logger *logrus.Logger) {
	var formattedMessage string

	switch v := message.(type) {
	case string:
		formattedMessage = message.(string)
	case []string:
		formattedMessage = strings.Join(message.([]string), "\n\n")
	case []byte:
		formattedMessage = string(message.([]byte))
	default:
		logger.Errorf("Unknown message type: %v", v)
	}

	if _, _, _, err := client.SendMessage(event.Channel, slack.MsgOptionText(formattedMessage, false)); err != nil {
		logger.Errorf("Failed to send message: %v", err)
	}
}
