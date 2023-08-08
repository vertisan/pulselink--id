package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	GrpcPort string `envconfig:"GRPC_PORT" default:"9090"`
}

func LoadConfig() (config Config, err error) {
	err = envconfig.Process("", &config)
	if err != nil {
		panic(err)
	}

	return
}
