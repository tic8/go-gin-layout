package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-layout/internal/config"
	"go-gin-layout/internal/handler"
)

func main() {
	config.InitConfig()
	config.InitLogger()
	config.InitMySQL()
	config.InitRedis()
	config.InitCron()

	r := gin.Default()
	handler.RegisterRoutes(r)
	handler.RegisterSwagger(r)
	handler.RegisterMetrics(r)

	fmt.Println("Starting server on :8080")
	r.Run(":8080")
}
