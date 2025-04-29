package config

import (
	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     Config.GetString("redis.addr"),
		Password: Config.GetString("redis.password"),
		DB:       Config.GetInt("redis.db"),
	})
}
