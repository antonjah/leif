package app

import (
	"github.com/nlopes/slack"

	"github.com/antonjah/leif/internal/lunches"
)

// Run loads configs and initializes handlers.
// Then it starts all goroutines and starts the app
func Run() {
	logger := InitLogger()

	logger.Info("Loading config...")
	config := NewConfig()

	client := slack.New(config.SlackToken)

	lunchHandler := lunches.NewLunchHandler(config.Cache)
	eventHandler := NewEventHandler(client, lunchHandler, logger)

	logger.Info("Connecting to slack...")
	rtm := client.NewRTM()
	go rtm.ManageConnection()

	logger.Info("Leif is now running!")
	for msg := range rtm.IncomingEvents {
		eventHandler.Handle(msg)
	}
}
