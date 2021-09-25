package routers

import (
	"github.com/gin-gonic/gin"
	"web_app/controller"
	"web_app/logger"
)

func Setup() *gin.Engine {

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(c *gin.Context) {})
	// 注册 业务路由
	r.POST("/register", controller.RegisterHandle)
	return r
}
