package app

import (
	"strings"

	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
)

func HandleResponse(message interface{}, event *slack.MessageEvent, client *slack.Client) {
	var formattedMessage string

	switch v := message.(type) {
	case string:
		formattedMessage = message.(string)
	case []string:
		formattedMessage = strings.Join(message.([]string), "\n\n")
	case []byte:
		formattedMessage = string(message.([]byte))
	default:
		log.Error("Unknown message type: %v", v)
	}

	if _, _, _, err := client.SendMessage(event.Channel, slack.MsgOptionText(formattedMessage, false)); err != nil {
		log.Errorf("Failed to send message: %v", err)
	}
}