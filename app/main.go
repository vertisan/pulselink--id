package main

import (
	"vrsf-playground/id-manager/cmd/grpc"
	"vrsf-playground/id-manager/config"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		panic("Cannot load configuration!")
	}

	grpc.InitServer(c)
}
