package router

import (
	"perServer/api"

	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("register", api.Register)
		BaseRouter.POST("login", api.Login)
		BaseRouter.POST("captcha", api.Captcha)
	}
}
