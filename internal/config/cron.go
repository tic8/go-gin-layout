package config

import (
	"github.com/robfig/cron/v3"
	"go-gin-layout/internal/global"
)

func InitCron() {
	global.Cron = cron.New()
	global.Cron.Start()
}
