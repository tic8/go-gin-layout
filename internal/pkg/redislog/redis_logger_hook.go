package redislog

import (
	"context"
	"time"

	"github.com/axiaoxin-com/logging"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type RedisLoggerHook struct{}

var _ redis.Hook = RedisLoggerHook{}

func (h RedisLoggerHook) DialHook(existing redis.DialHook) redis.DialHook {
	return existing // 不修改 Dial 行为
}

func (h RedisLoggerHook) ProcessHook(existing redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		start := time.Now()
		ctx = context.WithValue(ctx, "redis_start_time", start)

		err := existing(ctx, cmd) // 执行原来的 Hook（或核心命令）

		logging.Info(ctx, "Redis Command Executed",
			zap.String("cmd", cmd.String()),
			zap.Duration("cost", time.Since(start)),
			zap.Error(err),
		)

		return err
	}
}

func (h RedisLoggerHook) ProcessPipelineHook(existing redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		start := time.Now()
		ctx = context.WithValue(ctx, "redis_start_time", start)

		err := existing(ctx, cmds)

		var allCmds string
		for _, cmd := range cmds {
			allCmds += cmd.String() + "; "
		}

		logging.Info(ctx, "Redis Pipeline Executed",
			zap.String("cmds", allCmds),
			zap.Duration("cost", time.Since(start)),
			zap.Error(err),
		)

		return err
	}
}
