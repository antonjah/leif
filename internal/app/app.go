package app

import (
	"github.com/nlopes/slack"

	"github.com/antonjah/leif/internal/pkg/utils"
)

// Run loads configs, initializes handlers and starts Leif
func Run() {
	logger := InitLogger()

	logger.Info("Loading config...")
	config := NewConfig()

	client := slack.New(config.SlackToken)

	cache := utils.NewCache()
	eventHandler := NewEventHandler(client, cache, logger, config)

	logger.Info("Connecting to slack...")
	rtm := client.NewRTM()
	go rtm.ManageConnection()

	logger.Info("Leif is now running!")
	for msg := range rtm.IncomingEvents {
		eventHandler.Handle(msg)
	}
}
