package middleware

import (
	"context"
	"fmt"
	"github.com/axiaoxin-com/logging"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var TraceIDKey = "trace_id"

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
			traceID := uuid.New().String()
			// 将 trace_id 设置到上下文中
			ctx = context.WithValue(ctx, TraceIDKey, traceID)
			return traceID
		},
	}
	return logging.GinLoggerWithConfig(conf)
}
