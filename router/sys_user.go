package router

import (
	"perServer/api"
	"perServer/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		UserRouter.POST("changePassword", api.ChangePassword)   // 修改密码
		UserRouter.POST("uploadHeaderImg", api.UploadHeaderImg) // 上传头像
		UserRouter.POST("getUserList", api.GetUserList)         // 分页获取用户列表
	}
}
