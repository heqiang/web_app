package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_app/controller"
	"web_app/logger"
	"web_app/middlerware"
	"web_app/settings"
)

func Setup(conf *settings.AppConfig) *gin.Engine {
	gin.SetMode(conf.Mode)
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	v1 := r.Group("/api/v1")
	v1.GET("/", func(c *gin.Context) {})
	// 注册 业务路由
	v1.POST("/register", controller.RegisterHandle)
	v1.POST("/Login", controller.LoginHadle)
	v1.Use(middlerware.JWTAuthMiddleware())
	{
		v1.GET("/community", controller.CommunityHandle)
	}
	r.GET("/ping", middlerware.JWTAuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "已登录",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
