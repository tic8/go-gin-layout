package handler

import (
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/mcuadros/go-gin-prometheus"
)

func RegisterMetrics(r *gin.Engine) {
	//r.GET("/metrics", gin.WrapH(promhttp.Handler()))\
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)
}
