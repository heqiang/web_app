package routers

import (
	"github.com/gin-gonic/gin"
	"web_app/logger"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(c *gin.Context) {

	})
	return r
}
