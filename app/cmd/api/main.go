package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"id-manager/config"
	"id-manager/internal/healthcheck"
	"id-manager/internal/id_generator"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		panic("Cannot load configuration!")
	}

	r := gin.Default()

	healthcheck.RegisterRoutes(r, c)
	id_generator.RegisterRoutes(r, c)

	err = r.Run(fmt.Sprintf(":%s", c.Port))
	if err != nil {
		panic(fmt.Sprintf("Cannot start the server on port :%s", c.Port))
		return
	}
}
