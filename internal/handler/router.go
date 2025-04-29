package handler

import (
	"context"
	"fmt"
	"github.com/axiaoxin-com/logging"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	v1 "go-gin-layout/internal/api/v1"
	"go-gin-layout/internal/middleware"
)

func SetupRouter() *gin.Engine {
	// 创建 Gin 引擎
	r := gin.New()

	RegisterMetrics(r)
	RegisterSwagger(r)

	// 注册全局中间件
	// you can custom the config or use logging.GinLogger() by default config
	conf := logging.GinLoggerConfig{
		Formatter: func(c context.Context, m logging.GinLogDetails) string {
			return fmt.Sprintf("%s use %s request %s at %v, handler %s use %f seconds to respond it with %d",
				m.ClientIP, m.Method, m.RequestURI, m.ReqTime, m.HandlerName, m.Latency, m.StatusCode)
		},
		SkipPaths:     []string{},
		EnableDetails: false,
		TraceIDFunc: func(ctx context.Context) string {
			// 从上下文中获取 trace_id
			if traceID, ok := ctx.Value("trace_id").(string); ok {
				return traceID
			}
			// 如果上下文中没有 trace_id，则生成一个新的
			return uuid.New().String()
		},
	}
	r.Use(logging.GinLoggerWithConfig(conf)) // 日志中间件
	r.Use(gin.Recovery())                    // 恢复中间件，防止程序崩溃
	r.Use(cors.Default())                    // 跨域中间件
	r.Use(middleware.TraceIDMiddleware())    // 自定义 trace_id 中间件

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
