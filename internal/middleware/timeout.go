package middleware

import (
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"go-gin-layout/internal/api/response"
	"net/http"
	"time"
)

func Timeout() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(8*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(func(c *gin.Context) {
			response.ErrorWithHttpCode(c, http.StatusRequestTimeout, http.StatusRequestTimeout, "timeout", nil)
		}),
	)
}
