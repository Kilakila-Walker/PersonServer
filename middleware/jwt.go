package middleware

import (
	"perServer/global/response"
	"perServer/global/token"

	"github.com/gin-gonic/gin"
)

//jwt验证 除开登录 其他都需要检验jwt
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息
		x_token := c.Request.Header.Get("x-token")
		if x_token == "" {
			response.ToJson(
				response.ERROR,
				gin.H{"reload": true},
				"未登录或非法访问",
				c)
			c.Abort()
			return
		}

		j := token.NewJWT()
		// parseToken 解析token包含的信息
		_, err := j.ParseJwt(x_token)
		if err != nil {
			if err == token.TokenExpired {
				response.ToJson(
					response.ERROR,
					gin.H{"reload": true},
					"授权已过期",
					c)
				c.Abort()
				return
			}
			response.ToJson(
				response.ERROR,
				gin.H{"reload": true},
				err.Error(),
				c)
			c.Abort()
			return
		}
		c.Next()
	}
}
