package v1

import (
	"github.com/gin-gonic/gin"
	"go-gin-layout/internal/api/response"
	"go-gin-layout/internal/service"
)

type UserRouter struct {
	UserService *service.UserService
}

func RegisterUserRoutes(r *gin.RouterGroup) {
	userRouter := &UserRouter{
		UserService: service.NewUserService(), // 依赖注入
	}

	r.GET("", userRouter.GetAllUsersHandler)
	r.GET("/info", userRouter.GetUserInfoHandler)
}

func (u *UserRouter) GetUserInfoHandler(c *gin.Context) {
	userInfo := u.UserService.GetUserInfo(c)
	response.Success(c, userInfo)
}

func (u *UserRouter) GetAllUsersHandler(c *gin.Context) {
	users := u.UserService.GetAllUsers()
	response.Success(c, users)
}
