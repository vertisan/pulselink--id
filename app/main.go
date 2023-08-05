package main

import (
	"id-manager/cmd/api"
	"id-manager/config"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		panic("Cannot load configuration!")
	}

	api.Start(c)
}
