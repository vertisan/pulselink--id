package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port        string `envconfig:"PORT" default:"80"`
	ServiceName string `envconfig:"SERVICE_NAME" required:"true"`
}

func LoadConfig() (config Config, err error) {
	err = envconfig.Process("", &config)
	if err != nil {
		panic(err)
	}

	return
}
