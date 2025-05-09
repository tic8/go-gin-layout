package global

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var (
	DB          *gorm.DB
	RedisClient *redis.Client
	Cron        *cron.Cron
	ZapLogger   *zap.Logger
	Logger      *zap.SugaredLogger
	Ctx         context.Context
	Cancel      context.CancelFunc
)

type RedisLock struct {
	Key        string
	Expiration time.Duration
}

// Acquire 创建一个新的 Redis 锁 false:已经执行过 true:未执行
func (lock *RedisLock) Acquire() bool {
	if lock.Expiration == 0 {
		lock.Expiration = 5 * time.Second // 设置默认值为 5 秒
	}
	result, err := RedisClient.SetNX(Ctx, lock.Key, "locked", lock.Expiration).Result()
	if err != nil {
		Logger.Errorf("Failed to acquire lock for key %s: %s", lock.Key, err)
		return false
	}
	return result
}
