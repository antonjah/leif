package app

import (
	"github.com/nlopes/slack"

	"github.com/antonjah/gleif/internal/lunches"
)

func Run() {
	config := NewConfig()

	lunchHandler := lunches.NewHandler(config.Cache)

	client := slack.New(config.SlackToken)
	rtm := client.NewRTM()
	go rtm.ManageConnection()

	eventHandler := NewEventHandler(client, lunchHandler)

	for msg := range rtm.IncomingEvents {
		eventHandler.Handle(msg)
	}
}
