package v1

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	RegisterUserRoutes(r.Group("/user"))
}
