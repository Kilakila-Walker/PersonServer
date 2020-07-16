package router

import (
	"perServer/api"
	"perServer/middleware"

	"github.com/gin-gonic/gin"
)

func InitCasbinRouter(Router *gin.RouterGroup) {
	CasbinRouter := Router.Group("casbin").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		CasbinRouter.POST("updateCasbin", api.UpdateCasbin)
		CasbinRouter.POST("getPolicyPathByAuthorityId", api.GetPolicyPathByAuthorityId)
		CasbinRouter.GET("casbinTest/:pathParam", api.CasbinTest)
	}
}
