package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	v1 "go-gin-layout/internal/api/v1"
	"go-gin-layout/internal/middleware"
)

func SetupRouter() *gin.Engine {
	// 创建 Gin 引擎
	r := gin.New()

	RegisterMetrics(r)
	RegisterSwagger(r)

	// 注册全局中间件
	r.Use(middleware.Logger())            // 日志中间件
	r.Use(gin.Recovery())                 // 恢复中间件，防止程序崩溃
	r.Use(cors.Default())                 // 跨域中间件
	r.Use(middleware.TraceIDMiddleware()) // 自定义 trace_id 中间件
	r.Use(middleware.Timeout())

	// 注册路由
	RegisterRoutes(r)

	return r
}

func RegisterRoutes(r *gin.Engine) {

	// 注册路由分组
	apiV1 := r.Group("/api/v1")
	v1.RegisterRoutes(apiV1)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
}
