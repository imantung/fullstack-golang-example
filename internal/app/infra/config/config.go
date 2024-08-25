package config

import (
	"github.com/imantung/fullstack-golang-example/internal/app/infra/di"
	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		Address string `envconfig:"ADDRESS" required:"true" default:":8080"`
	}
)

const Prefix = "APP"

var _ = di.Provide(NewConfig)

func NewConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process(Prefix, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
