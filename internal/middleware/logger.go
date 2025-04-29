package middleware

import (
	"context"
	"fmt"
	"github.com/axiaoxin-com/logging"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var TraceIDKey = "trace-id"

func Logger() gin.HandlerFunc {
	conf := logging.GinLoggerConfig{
		Formatter: func(c context.Context, m logging.GinLogDetails) string {
			return fmt.Sprintf("%s use %s request %s at %v, handler %s use %f seconds to respond it with %d",
				m.ClientIP, m.Method, m.RequestURI, m.ReqTime, m.HandlerName, m.Latency, m.StatusCode)
		},
		SkipPaths:          []string{},
		EnableDetails:      false,
		EnableRequestBody:  true,
		EnableResponseBody: true,
		TraceIDFunc: func(ctx context.Context) string {
			// 从上下文中获取 trace_id
			if traceID, ok := ctx.Value(TraceIDKey).(string); ok {
				return traceID
			}
			// 如果上下文中没有 trace_id，则生成一个新的
			return uuid.New().String()
		},
	}
	return logging.GinLoggerWithConfig(conf)
}
