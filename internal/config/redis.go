package config

import (
	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     Cfg.Redis.Addr,
		Password: Cfg.Redis.Password,
		DB:       Cfg.Redis.DB,
	})
}
