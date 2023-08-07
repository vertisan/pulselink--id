package id

import (
	"github.com/gin-gonic/gin"
	"id-manager/config"
	idGenerator "id-manager/internal/id/service"
)

func RegisterRoutes(router *gin.RouterGroup, config config.Config) {
	idRoutes := router.Group("/id")
	{
		idRoutes.POST("", handleGetIdGenerate)
		idRoutes.POST("/", handleGetIdGenerate) // TODO: Any alternative without a redirect?
	}
}

func handleGetIdGenerate(ctx *gin.Context) {
	ctx.JSON(idGenerator.GenerateId())
}
