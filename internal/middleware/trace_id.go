package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// TraceIDMiddleware 检查 header 中是否有 trace_id，如果没有则生成并设置
func TraceIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.GetHeader("trace-id")
		if traceID == "" {
			traceID = uuid.New().String()
			c.Request.Header.Set("trace-id", traceID)
		}
		c.Next()
	}
}
