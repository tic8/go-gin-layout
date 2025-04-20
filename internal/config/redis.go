package config

import (
	"github.com/redis/go-redis/v9"
	"go-gin-layout/internal/global"
)

func InitRedis() {
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:     Config.GetString("redis.addr"),
		Password: Config.GetString("redis.password"),
		DB:       Config.GetInt("redis.db"),
	})
}
