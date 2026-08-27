package core_http_server

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Addr            string        `envconfig:"ADDR"             required:"true"`
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" required:"true"`
	SecureCookies   bool          `envconfig:"SECURE_COOKIES"   default:"false"`
	// AllowedOrigins  []string      `envconfig:"ALLOWED_ORIGINS" required:"true"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("HTTP", &config); err != nil {
		return Config{}, fmt.Errorf("process envconfig: %w", err)
	}
	return config, nil
}

func NewConfigMust() Config {
	var config Config

	if err := envconfig.Process("HTTP", &config); err != nil {
		err = fmt.Errorf("get HTTP serbver config %w", err)
		panic(err)
	}
	return config
}
