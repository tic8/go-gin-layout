package redislog

import (
	"context"
	"github.com/spf13/cast"
	"go-gin-layout/internal/global"
	"time"

	"github.com/axiaoxin-com/logging"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var TraceIdKey = "trace_id"
var LoggerKey = "redis"

type RedisLoggerHook struct{}

var _ redis.Hook = RedisLoggerHook{}

func (h RedisLoggerHook) DialHook(existing redis.DialHook) redis.DialHook {
	return existing // 不修改 Dial 行为
}

func (h RedisLoggerHook) ProcessHook(existing redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		start := time.Now()
		traceId := ctx.Value(TraceIdKey)
		ctx = context.WithValue(ctx, TraceIdKey, ctx.Value(TraceIdKey))
		ctx = context.WithValue(ctx, "redis_start_time", start)

		ctxLogger := logging.CtxLogger(ctx)
		ctx, ctxLogger = logging.NewCtxLogger(ctx, logging.CloneLogger(LoggerKey), cast.ToString(traceId))
		logging.ReplaceLogger(global.ZapLogger)

		err := existing(ctx, cmd) // 执行原来的 Hook（或核心命令）

		ctxLogger.Info("Redis Command Executed",
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
		traceId := ctx.Value(TraceIdKey)
		ctx = context.WithValue(ctx, TraceIdKey, ctx.Value(TraceIdKey))
		ctx = context.WithValue(ctx, "redis_start_time", start)

		ctxLogger := logging.CtxLogger(ctx)
		ctx, ctxLogger = logging.NewCtxLogger(ctx, logging.CloneLogger(LoggerKey), cast.ToString(traceId))
		logging.ReplaceLogger(global.ZapLogger)
		err := existing(ctx, cmds)

		var allCmds string
		for _, cmd := range cmds {
			allCmds += cmd.String() + "; "
		}

		ctxLogger.Info("Redis Pipeline Executed",
			zap.String("cmds", allCmds),
			zap.Duration("cost", time.Since(start)),
			zap.Error(err),
		)

		return err
	}
}
