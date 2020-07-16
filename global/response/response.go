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

func ToJson(code int, data interface{}, message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		message,
	})
}
