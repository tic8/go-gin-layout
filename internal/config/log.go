package config

import (
	"go-gin-layout/internal/global"
	"go.uber.org/zap"
)

func InitLogger() {
	logger, _ := zap.NewProduction()
	global.Logger = logger
}
