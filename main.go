package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web_app/dao/mysqlc"
	"web_app/dao/redis"
	"web_app/logger"
	"web_app/routers"
	"web_app/settings"
)

func main() {
	//加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init setting failed,err:%v\n\n", err)
		return
	}
	//初始化日志
	if err := logger.InitLogger(settings.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed,err:%v\n\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("Logger init successed")
	//初始化mysql连接
	if err := mysqlc.InitMySQL(settings.Conf.MysqlConfig); err != nil {
		fmt.Printf("init mysql failed,err:%v\n\n", err)
		return
	}
	//初始化reids连接
	if err := redis.InitClient(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed,err:%v\n\n", err)
		return
	}
	defer redis.Close()
	//注册路由
	r := routers.Setup()
	//启动服务 优雅关机
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", viper.GetString("app.port")),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	log.Println("Server exiting")

}
