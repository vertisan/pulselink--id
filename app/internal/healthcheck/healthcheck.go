package healthcheck

import (
	"github.com/gin-gonic/gin"
	"id-manager/config"
	healthcheck "id-manager/internal/healthcheck/service"
)

func RegisterRoutes(router *gin.Engine, config config.Config) {
	router.GET("/healthcheck", func(ctx *gin.Context) {
		handleGetHealthcheck(ctx, config)
	})
}

func handleGetHealthcheck(ctx *gin.Context, c config.Config) {
	ctx.JSON(healthcheck.GetHealthcheckStatus(c))
}
