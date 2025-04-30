package config

import (
	"github.com/axiaoxin-com/logging"
	"go.uber.org/zap"
)

func InitLogger() *zap.SugaredLogger {
	// 配置 Options 创建 logger
	options := logging.Options{
		Name:   "logging", // logger 名称
		Level:  "debug",   // zap 的 AtomicLevel ， logger 日志级别
		Format: "json",    // 日志输出格式为 json
		//OutputPaths:       []string{"stderr"}, // 日志输出位置为 stderr
		InitialFields:     nil,   // DefaultInitialFields 初始 logger 带有 pid 字段
		DisableCaller:     false, // 是否打印调用的代码行位置
		DisableStacktrace: false, // 错误日志是否打印调用栈信息
	}
	optionsLogger, _ := logging.NewLogger(options)
	return optionsLogger.Sugar()
}
