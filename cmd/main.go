package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-layout/internal/config"
	"go-gin-layout/internal/global"
	"go-gin-layout/internal/handler"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 初始化配置和资源
	config.InitConfig()
	config.InitLogger()
	config.InitMySQL()
	config.InitRedis()
	config.InitCron()

	// 创建 Gin 引擎
	r := gin.Default()
	handler.RegisterRoutes(r)
	handler.RegisterSwagger(r)
	handler.RegisterMetrics(r)

	// 启动服务器
	server := &http.Server{
		Addr:    ":8011",
		Handler: r,
	}

	// 在单独的 Goroutine 中启动 HTTP 服务
	go func() {
		fmt.Println("Starting server on :8011")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Server error: %s\n", err)
		}
	}()

	// 捕获系统信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // 阻塞，直到接收到信号

	fmt.Println("Shutting down server...")

	// 创建上下文，设置超时时间为 5 秒
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 优雅关闭 HTTP 服务器
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown: %s\n", err)
	}

	// 关闭其他资源
	if global.Cron != nil {
		global.Cron.Stop()
	}
	if global.RedisClient != nil {
		if err := global.RedisClient.Close(); err != nil {
			fmt.Printf("Failed to close Redis: %s\n", err)
		}
	}
	if config.DB != nil {
		sqlDB, _ := config.DB.DB()
		if err := sqlDB.Close(); err != nil {
			fmt.Printf("Failed to close MySQL: %s\n", err)
		}
	}

	fmt.Println("Server exiting")
}
