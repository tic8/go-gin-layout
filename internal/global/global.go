package global

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	RedisClient *redis.Client
	Cron        *cron.Cron
	Logger      *zap.Logger
	Ctx         = context.Background()
)
