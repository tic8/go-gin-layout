package config

import (
	"github.com/redis/go-redis/v9"
	"go-gin-layout/internal/pkg/redislog"
)

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     Cfg.Redis.Addr,
		Password: Cfg.Redis.Password,
		DB:       Cfg.Redis.DB,
	})

	if Cfg.Redis.Debug {
		rdb.AddHook(&redislog.RedisLoggerHook{})
	}
	return rdb
}
