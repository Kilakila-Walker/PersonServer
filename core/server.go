package core

import (
	"fmt"
	"net/http"
	"perServer/global"
	"perServer/initialize"
	"time"
)

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	//InstallPlugs(Router)
	// end 插件描述

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := &http.Server{
		Addr:           address,
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 保证文本顺序输出
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Debug("server run success on ", address)

	fmt.Printf(`欢迎使用 perServer
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080`, s.Addr)
	global.GVA_LOG.Error(s.ListenAndServe())
}
