package app

import (
	"github.com/antonjah/leif/internal/pkg/utils"
	"github.com/nlopes/slack"
)

// Run loads configs, initializes handlers and starts Leif
func Run() {
	logger := InitLogger()

	logger.Info("Loading config...")
	config := NewConfig(logger)

	client := slack.New(config.SlackToken)

	cache := utils.NewCache()
	eventHandler := EventHandler{client, cache, logger, config}

	logger.Info("Connecting to slack...")
	rtm := client.NewRTM()
	go rtm.ManageConnection()

	logger.Info("Leif is now running!")

	logEnabledPlugins(config, logger)
	for msg := range rtm.IncomingEvents {
		eventHandler.Handle(msg)
	}
}
