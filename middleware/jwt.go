package middleware

import (
	"perServer/global/response"
	"perServer/global/token"

	"github.com/gin-gonic/gin"
)

//jwt验证 除开登录 其他都需要检验jwt
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localSstorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
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
		claims, err := j.ParseToken(x_token)
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
		c.Set("claims", claims)
		c.Next()
	}
}
