package initialize

import (
	"fmt"
	_ "perServer/docs"
	"perServer/global"
	"perServer/middleware"
	"perServer/router"
	"time"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由
func RunServer() {
	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		Redis()
	}
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	var Router = gin.New()
	Router.Use(gin.Recovery())
	//使用默认log（只会打印）
	Router.Use(gin.Logger())
	//使用中间件 log记录（会打印并记录文件）
	//Router.Use(middleware.Logger())
	// 打开tls需要RunTLS关闭Run
	// Router.Use(middleware.LoadTls())
	// 跨域
	Router.Use(middleware.Cors())
	fmt.Printf("use middleware cors")
	//测试api用
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	fmt.Printf("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	ApiGroup := Router.Group("")
	router.InitUserRouter(ApiGroup)   // 注册用户路由
	router.InitBaseRouter(ApiGroup)   // 注册基础功能路由 不做鉴权
	router.InitSystemRouter(ApiGroup) // system相关路由

	//ssl配套使用 参数为ssl地址
	//Router.RunTLS(address,"../ssl.poe","../ssl.key")
	Router.Run(address)
	// 保证文本顺序输出
	time.Sleep(10 * time.Microsecond)
	fmt.Printf(`欢迎使用 perServer
	自动化文档地址:http://127.0.0.1%s/swagger/index.html`, address)
}
