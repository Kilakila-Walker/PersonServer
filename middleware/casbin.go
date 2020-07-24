package middleware

import (
	"perServer/global"
	"perServer/global/response"
	"perServer/global/token"
	"perServer/service"

	"github.com/gin-gonic/gin"
)

// 鉴权中间件 暂时未修改好 需要在jwt验证之后开启 开启方式在router文件夹
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		waitUse, code := token.GetClaims(c)
		if code != 0 {
			response.ToJson(
				response.ERROR,
				gin.H{"reload": true},
				"未登录或非法访问",
				c)
			c.Abort()
			return
		}
		// 获取请求的URI
		obj := c.Request.URL.RequestURI()
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的ID
		sub := waitUse.ID
		e := service.Casbin()
		// 判断策略中是否存在
		if global.GVA_CONFIG.System.Env == "develop" || e.Enforce(sub, obj, act) {
			c.Next()
		} else {
			response.ToJson(response.ERROR, gin.H{}, "权限不足", c)
			c.Abort()
			return
		}
	}
}
