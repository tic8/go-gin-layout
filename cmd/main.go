package main

import (
	"context"
	"fmt"
	"go-gin-layout/internal/config"
	"go-gin-layout/internal/errcode"
	"go-gin-layout/internal/global"
	"go-gin-layout/internal/handler"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 初始化资源
	initResources()

	// 创建 Gin 引擎
	r := handler.SetupRouter()

	// 从配置中读取端口
	listen := fmt.Sprintf("%s:%d", config.Cfg.Server.Host, config.Cfg.Server.Port)

	// 启动服务器
	server := &http.Server{
		Addr:    listen,
		Handler: r,
	}

	// 在单独的 Goroutine 中启动 HTTP 服务
	go func() {
		global.Logger.Infof("Starting server on %s", listen)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Fatalf("Server error: %s", err)
		}
	}()

	// 捕获系统信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // 阻塞，直到接收到信号

	global.Logger.Info("Shutting down server...")

	// 优雅关闭服务
	shutdownServer(server)

	global.Logger.Info("Server exiting")
}

// initResources 初始化配置和资源
func initResources() {
	// 初始化全局上下文和取消函数
	ctx, cancel := context.WithCancel(context.Background())
	global.Ctx = ctx
	global.Cancel = cancel

	// 初始化配置、日志、数据库、Redis、定时任务和守护进程
	config.InitConfig()
	global.Logger = config.InitLogger()
	global.DB = config.InitMySQL()
	global.RedisClient = config.InitRedis()
	config.InitCron()
	config.InitDaemon()
	// 初始化错误码
	errcode.InitErrors()
}

// shutdownServer 优雅关闭服务器和资源
func shutdownServer(server *http.Server) {
	// 创建上下文，设置超时时间为 5 秒
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 取消InitDaemon的执行
	if global.Cancel != nil {
		global.Cancel()
	}

	// 优雅关闭 HTTP 服务器
	if err := server.Shutdown(ctx); err != nil {
		global.Logger.Errorf("Server forced to shutdown: %s", err)
	}

	// 关闭其他资源
	if global.Cron != nil {
		global.Cron.Stop()
	}
	if global.RedisClient != nil {
		if err := global.RedisClient.Close(); err != nil {
			global.Logger.Errorf("Failed to close Redis: %s", err)
		}
	}
	if global.DB != nil {
		sqlDB, _ := global.DB.DB()
		if err := sqlDB.Close(); err != nil {
			global.Logger.Errorf("Failed to close MySQL: %s", err)
		}
	}
}
