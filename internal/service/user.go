package service

import (
	"context"
	"go-gin-layout/internal/global"
	"go-gin-layout/internal/model"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUserInfo(ctx context.Context) map[string]interface{} {
	global.DB.WithContext(ctx).Model(&model.User{}).First(&model.User{})
	// 模拟获取用户信息的逻辑
	return map[string]interface{}{
		"id":   1,
		"name": "John Doe",
		"age":  30,
	}
}

func (s *UserService) GetAllUsers() []map[string]interface{} {
	// 模拟获取所有用户的逻辑
	return []map[string]interface{}{
		{"id": 1, "name": "John Doe", "age": 30},
		{"id": 2, "name": "Jane Doe", "age": 25},
	}
}
