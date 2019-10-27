// Copyright (c) 2019 Braggart Inc. All Rights Reserved.
// This is licensed software from Braggart Inc, for limitations
// and restrictions contact your company contract manager.

package config

import (
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
)

// Config specified configuration option through env vars
type Config struct {
	Port     int    `env:"PORT"`
	BasePath string `env:"BASE_PATH"`

	DBHost     string `env:"DB_HOST"`
	DBPort     int    `env:"DB_PORT"`
	DBName     string `env:"DB_NAME"`
	DBUsername string `env:"DB_USERNAME"`
	DBPassword string `env:"DB_PASSWORD"`
}

// GetConfig get configuration from environment variables
func GetConfig() *Config {
	conf := &Config{}
	err := env.Parse(conf)
	if err != nil {
		logrus.WithFields(logrus.Fields{"module": "config"}).
			Fatal("unable to parse environment variables: ", err)
	}

	return conf
}
