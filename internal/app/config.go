package app

import (
	"github.com/caarlos0/env"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"

	"github.com/antonjah/leif/internal/pkg/utils"
)

// Config app configuration
type Config struct {
	Cache       cache.Cache
	SlackToken  string `env:"SLACK_TOKEN,required"`
	GitLabToken string `env:"GITLAB_TOKEN"`
	GitLabURL   string `env:"GITLAB_BASE_URL"`
}

// NewConfig returns a new app configuration with ENVs loaded
func NewConfig() Config {
	appCache := utils.NewCache()
	config := Config{Cache: appCache}

	if err := env.Parse(&config); err != nil {
		log.Fatal(err)
	}

	return config
}
