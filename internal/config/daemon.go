package config

import (
	"go-gin-layout/internal/global"
)

func InitDaemon() {
	go func() {
		lock := &global.RedisLock{
			Key: "daemon:single_execution", // 分布式锁的键
		}
		if !lock.Acquire() {
			global.Logger.Warn("Another instance is already running the daemon, skipping execution")
			return
		}

		// 启动多个任务
		go func() {}()
		go func() {}()
		go func() {}()

		// 主 Goroutine 等待上下文取消信号
		<-global.Ctx.Done()
		global.Logger.Info("Context canceled, exiting daemon")
	}()
}
