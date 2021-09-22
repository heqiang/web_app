package main

import (
	"fmt"
	"go.uber.org/zap"
	"web_app/logger"
	"web_app/settings"
)

func main() {
	//加载配置
	if err := settings.Init(); err != nil {
		fmt.Println("init setting failed,err:%v\n", err)
		return
	}
	//初始化日志
	if err := logger.InitLogger(); err != nil {
		fmt.Println("init logger failed,err:%v\n", err)
		return
	}
	zap.L().Debug("Logger init successed")
	//初始化mysql连接
	if err := mysql.Init(); err != nil {
		fmt.Println("init mysql failed,err:%v\n", err)
		return
	}
	//初始化reids连接
	if err := redis.Init(); err != nil {
		fmt.Println("init redis failed,err:%v\n", err)
		return
	}
	//注册路由

	//启动服务

}
