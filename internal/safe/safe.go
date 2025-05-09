package safe

import (
	"fmt"
	"go-gin-layout/internal/alert"
	"log"
	"runtime/debug"
)

// Go 启动一个安全的 goroutine，捕获 panic 并发送飞书告警
func Go(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				// 捕获 panic 并记录日志
				log.Printf("goroutine panic: %v\n%s", r, debug.Stack())

				// 发送飞书告警
				title := "服务异常告警"
				content := fmt.Sprintf("捕获到 panic: %v\n\n堆栈信息:\n%s", r, debug.Stack())
				alert.SendAlert(title, content)
			}
		}()
		fn()
	}()
}
