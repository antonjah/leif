package app

import (
	"github.com/caarlos0/env"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"

	"github.com/antonjah/gleif/internal/utils"
)

type Config struct {
	Cache cache.Cache
	SlackToken string `env:"SLACK_TOKEN,required"`
}

func NewConfig() Config {
	appCache := utils.NewCache()
	config := Config{Cache: appCache}

	if err := env.Parse(&config); err != nil {
		log.Fatal(err)
	}

	return config
}
