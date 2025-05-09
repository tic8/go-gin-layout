package config

import (
	"github.com/axiaoxin-com/logging"
	"go.uber.org/zap"
)

var ZapLogger *zap.Logger

func InitLogger(logPath string) *zap.SugaredLogger {
	// scheme 为 lumberjack ，日志文件为 /tmp/x.log , 保存 7 天，保留 10 份文件，文件大小超过 100M ，使用压缩备份，压缩文件名使用 localtime
	sink := logging.NewLumberjackSink("lumberjack", logPath+"app.log", 7, 10, 1024, true, true)
	// 配置 Options 创建 logger
	options := logging.Options{
		LumberjackSink: sink,
		//Name:              "logging",               // logger 名称
		Level:             "debug",                 // zap 的 AtomicLevel ， logger 日志级别
		Format:            "json",                  // 日志输出格式为 json
		OutputPaths:       []string{"lumberjack:"}, // 日志输出位置为 stderr
		InitialFields:     nil,                     // DefaultInitialFields 初始 logger 带有 pid 字段
		DisableCaller:     false,                   // 是否打印调用的代码行位置
		DisableStacktrace: false,                   // 错误日志是否打印调用栈信息
		//EncoderConfig:     &zapcore.EncoderConfig{},
	}
	optionsLogger, _ := logging.NewLogger(options)
	//logging.ReplaceLogger(optionsLogger)

	//logging.RegisterLumberjackSink(sink)

	ZapLogger = optionsLogger
	return optionsLogger.Sugar()
}
