package v1

import (
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup) {
	r.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "user v1"})
	})

	r.GET("/info", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "user info"})
	})
}
