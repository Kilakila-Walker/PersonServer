package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

//统一的返回格式
func ToJson(code int, data interface{}, message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		message,
	})
}

//提供错误码的返回形式
func Json(status int, code int, data interface{}, message string, c *gin.Context) {
	c.JSON(status, Response{
		code,
		data,
		message,
	})
}
