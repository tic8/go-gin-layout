package config

import (
	"github.com/robfig/cron/v3"
	"go-gin-layout/internal/global"
	"time"
)

func InitCron() {
	lock := &global.RedisLock{
		Key:        "init:cron_lock", // 分布式锁的键
		Expiration: 3 * time.Second,
	}
	if !lock.Acquire() {
		global.Logger.Warn("InitCron is already running, skipping initialization")
		return
	}
	
	location, _ := time.LoadLocation("Asia/Shanghai")
	global.Cron = cron.New(cron.WithSeconds(), cron.WithLocation(location))

	// Add your cron jobs here
	// Example: global.Cron.AddFunc("0 0 0 * * *", func() { fmt.Println("Every day at midnight") })

	global.Cron.Start()
}
