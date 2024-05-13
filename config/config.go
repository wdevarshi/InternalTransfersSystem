package config

import (
	"context"

	cbConfig "github.com/go-coldbrew/core/config"
	"github.com/go-coldbrew/log"
	"github.com/kelseyhightower/envconfig"
)

// defaultConfig is the default configuration for the application
// It is loaded from environment variables
var defaultConfig Config

type Config struct {
	cbConfig.Config
	PanicOnConfigError bool `envconfig:"PANIC_ON_CONFIG_ERROR" default:"true"`
	// App configuration
	Prefix string `envconfig:"PREFIX" default:"got"`
	DBUser string `envconfig:"DBUser" default:"devarshiwaghela"`

	//TODO: Move this config to vault
	DBPassword string `envconfig:"DBPassword" default:"waghela"`
	DBName     string `envconfig:"DBName" default:"internaltransferssystem"`
}

func init() {
	err := envconfig.Process("", &defaultConfig)
	// fail on error
	if err != nil {
		if defaultConfig.PanicOnConfigError {
			panic(err)
		} else {
			log.Error(context.Background(), "msg", "error while loading config", "err", err)
		}
	}
}

// Get returns the default configuration
// This is used by the application to load the configuration
func Get() Config {
	return defaultConfig
}

// GetColdBrewConfig returns the default configuration
// This is used by the coldbrew framework to load the configuration for the application
func GetColdBrewConfig() cbConfig.Config {
	return defaultConfig.Config
}
