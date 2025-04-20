package config

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"go-gin-layout/internal/global"
)

func InitCron() {
	global.Cron = cron.New()
	_, err := global.Cron.AddFunc("@every 1m", func() {
		fmt.Println("⏰ Running scheduled task...")
	})
	if err != nil {
		panic(fmt.Errorf("failed to schedule task: %w", err))
	}
	global.Cron.Start()
}
