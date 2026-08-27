package crypto_hasher

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"golang.org/x/crypto/bcrypt"
)

type Config struct {
	Cost int `envconfig:"COST" default:"12"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("BCRYPT", &config); err != nil {
		return Config{}, fmt.Errorf("process bcrypt envconfig: %w", err)
	}

	if config.Cost < bcrypt.MinCost || config.Cost > bcrypt.MaxCost {
		return Config{}, fmt.Errorf(
			"invalid bcrypt cost %d: must be between %d and %d",
			config.Cost,
			bcrypt.MinCost,
			bcrypt.MaxCost,
		)
	}

	return config, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		panic(fmt.Errorf("get bcrypt config: %w", err))
	}

	return config
}
