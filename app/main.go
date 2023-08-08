package main

import (
	"id-manager/cmd/grpc"
	"id-manager/config"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		panic("Cannot load configuration!")
	}

	grpc.InitServer(c)
}
