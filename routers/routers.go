package routers

import (
	"github.com/gin-gonic/gin"
	"web_app/controller"
	"web_app/logger"
	"web_app/settings"
)

func Setup(conf *settings.AppConfig) *gin.Engine {
	gin.SetMode(conf.Mode)
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(c *gin.Context) {})
	// 注册 业务路由
	r.POST("/register", controller.RegisterHandle)
	r.POST("/Login", controller.LoginHadle)
	return r
}
