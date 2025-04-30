package response

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"go-gin-layout/internal/errcode"
	"net/http"
)

// Response 响应结构体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	resData := fmtEmptyData(data)
	c.JSON(http.StatusOK, Response{
		Code:    errcode.Success,
		Message: errcode.GetErrorMessage(errcode.Success),
		Data:    resData,
	})
}

// Error 错误响应
func Error(c *gin.Context, code int, message string, data interface{}) {
	resData := fmtEmptyData(data)
	errMsg := message
	if errMsg == "" {
		errMsg = errcode.GetErrorMessage(code)
	}
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: errMsg,
		Data:    resData,
	})
}

// ErrorWithHttpCode 错误响应
func ErrorWithHttpCode(c *gin.Context, httpCode int, code int, message string, data interface{}) {
	resData := fmtEmptyData(data)
	errMsg := message
	if errMsg == "" {
		errMsg = errcode.GetErrorMessage(code)
	}
	c.JSON(httpCode, Response{
		Code:    code,
		Message: errMsg,
		Data:    resData,
	})
}

func fmtEmptyData(data interface{}) interface{} {
	if validate.IsEmpty(data) {
		return map[string]interface{}{}
	}
	return data
}
