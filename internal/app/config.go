package app

import (
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
)

// Config app configuration
type Config struct {
	SlackToken    string `env:"SLACK_TOKEN,required"`
	GitLabToken   string `env:"GITLAB_TOKEN"`
	GitLabURL     string `env:"GITLAB_BASE_URL"`
	PostMordToken string `env:"POSTMORD_TOKEN"`
}

// NewConfig returns a new app configuration with ENVs loaded
func NewConfig(logger *logrus.Logger) Config {
	c := Config{}

	if err := env.Parse(&c); err != nil {
		logger.Fatal(err)
	}

	return c
}
