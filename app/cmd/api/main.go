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

	router := gin.Default()

	api := router.Group("/api/v1")
	{
		id_generator.RegisterRoutes(api, c)
	}

	healthcheck.RegisterRoutes(router, c)

	err = router.Run(fmt.Sprintf(":%s", c.Port))
	if err != nil {
		panic(fmt.Sprintf("Cannot start the server on port :%s", c.Port))
	}
}
