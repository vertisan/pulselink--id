package id_generator

import (
	"github.com/gin-gonic/gin"
	"id-manager/config"
	idGenerator "id-manager/internal/id_generator/service"
)

func RegisterRoutes(router *gin.Engine, config config.Config) {
	idRoutes := router.Group("/id")
	{
		idRoutes.GET("/generate", func(ctx *gin.Context) {
			handleGetIdGenerate(ctx)
		})
	}
}

func handleGetIdGenerate(ctx *gin.Context) {
	ctx.JSON(idGenerator.GenerateId())
}
