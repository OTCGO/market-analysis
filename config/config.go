package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type (
	// Config top level struct representing the config
	// for the node.
	Config struct {
		ApplicationConfiguration ApplicationConfiguration
		LogConfiguration         LogConfiguration
		MongoConfiguration       MongoConfiguration
	}

	// ApplicationConfiguration config specific to the node.
	ApplicationConfiguration struct {
		CoinMarketCapURL string
	}

	// LogConfiguration
	LogConfiguration struct {
		Level  string
		Output string
	}
	// LogConfiguration
	MongoConfiguration struct {
		Url       string
		PoolLimit int
		Db        string
	}
)

var conf *Config

// Loadattempts to load the config from the give
// path and netMode.
func Load(path string, name string) (config *Config, err error) {
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	viper.AddConfigPath(path)
	viper.SetConfigName(name)

	if err := viper.ReadInConfig(); err != nil {
		errors.Wrap(err, "viper ReadInConfig error")
	}

	// var configuration Config.ApplicationConfiguration
	err = viper.Unmarshal(&config)
	if err != nil {
		return &Config{}, errors.Wrap(err, "Problem unmarshaling config json data")
	}
	conf = config
	return config, nil
}

func GetConfig() (config *Config, err error) {
	if conf == nil {
		errors.Wrap(err, "viper ReadInConfig error")
	}

	return conf, err
}
