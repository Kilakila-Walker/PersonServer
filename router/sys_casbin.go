package router

import (
	"perServer/api"
	"perServer/middleware"

	"github.com/gin-gonic/gin"
)

func InitCasbinRouter(Router *gin.RouterGroup) {
	//使用中间件 JWT鉴权
	CasbinRouter := Router.Group("casbin").Use(middleware.JWTAuth())
	{
		CasbinRouter.POST("addCasbin", api.AddCasbin)
		CasbinRouter.POST("getPolicyPathByAuthorityId", api.GetPolicyPathByAuthorityId)

	}
}
