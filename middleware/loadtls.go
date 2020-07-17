package middleware

import (
	"fmt"
	"perServer/global"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

// 用https把这个中间件在router里面use一下就好

func LoadTls() gin.HandlerFunc {
	address := fmt.Sprintf("localhost:%d", global.GVA_CONFIG.System.Addr)
	return func(c *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     address,
		})
		err := middleware.Process(c.Writer, c.Request)
		if err != nil {
			// 如果出现错误，请不要继续
			fmt.Println(err)
			return
		}
		// 继续往下处理
		c.Next()
	}
}
