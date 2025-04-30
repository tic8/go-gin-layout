package errcode

import (
	"fmt"
	"sync"
)

// 定义错误码常量
const (
	Success       = 0
	InvalidParam  = 4001
	Unauthorized  = 4002
	InternalError = 5001
)

// 定义错误信息
var errorMessages = map[int]string{
	Success:       "success",
	InvalidParam:  "parameter error",
	Unauthorized:  "unauthorized",
	InternalError: "internal server error",
}

var (
	mu         sync.RWMutex
	errorCodes = make(map[int]string)
)

// RegisterError 注册错误码
func RegisterError(code int, message string) error {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := errorCodes[code]; exists {
		return fmt.Errorf("error code %d already registered", code)
	}
	errorCodes[code] = message
	return nil
}

// GetErrorMessage 根据错误码获取错误信息
func GetErrorMessage(code int) string {
	mu.RLock()
	defer mu.RUnlock()
	if msg, exists := errorCodes[code]; exists {
		return msg
	}
	return "Unknown error"
}

// InitErrors 初始化所有错误码
func InitErrors() {
	for code, message := range errorMessages {
		MustRegisterError(code, message)
	}
}

// MustRegisterError 强制注册错误码（用于初始化时）
func MustRegisterError(code int, message string) {
	if err := RegisterError(code, message); err != nil {
		panic(err)
	}
}
