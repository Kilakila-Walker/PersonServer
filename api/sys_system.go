package api

import (
	"fmt"
	"perServer/global/response"
	"perServer/model"
	resp "perServer/model/response"
	"perServer/service"

	"github.com/gin-gonic/gin"
)

// @Tags system
// @Summary 获取配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /system/getSystemConfig [post]
func GetSystemConfig(c *gin.Context) {
	err, config := service.GetSystemConfig()
	if err != nil {
		response.ToJson(response.ERROR, gin.H{}, fmt.Sprintf("获取失败，%v", err), c)
	} else {
		response.ToJson(response.SUCCESS, resp.SysConfigResponse{Config: config}, "成功", c)
	}
}

// @Tags system
// @Summary 设置配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body model.System true "设置配置文件内容"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /system/setSystemConfig [post]
func SetSystemConfig(c *gin.Context) {
	var sys model.System
	_ = c.ShouldBindJSON(&sys)
	err := service.SetSystemConfig(sys)
	if err != nil {
		response.ToJson(response.ERROR, gin.H{}, fmt.Sprintf("设置失败，%v", err), c)
	} else {
		response.ToJson(response.ERROR, gin.H{}, "设置成功", c)
	}
}

// 本方法开发中 开发者windows系统 缺少linux系统所需的包 因此搁置
// @Tags system
// @Summary 设置配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body model.System true "设置配置文件内容"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /system/ReloadSystem [post]
func ReloadSystem(c *gin.Context) {
	var sys model.System
	_ = c.ShouldBindJSON(&sys)
	err := service.SetSystemConfig(sys)
	if err != nil {
		response.ToJson(response.ERROR, gin.H{}, fmt.Sprintf("设置失败，%v", err), c)
	} else {
		response.ToJson(response.SUCCESS, gin.H{}, "设置成功", c)
	}
}
