package app

import (
	"github.com/nlopes/slack"

	"github.com/antonjah/leif/internal/lunches"
)

// Run loads configs and initializes handlers.
// Then it starts all goroutines and starts the app
func Run() {
	config := NewConfig()

	client := slack.New(config.SlackToken)

	lunchHandler := lunches.NewLunchHandler(config.Cache)
	eventHandler := NewEventHandler(client, lunchHandler)

	rtm := client.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		eventHandler.Handle(msg)
	}
}
