package api

//系统配置API
import (
	"fmt"
	"perServer/global/response"
	"perServer/model"
	resp "perServer/model/response"
	"perServer/service"

	"github.com/gin-gonic/gin"
)

// 获取配置文件内容
func GetSystemConfig(c *gin.Context) {
	err, config := service.GetSystemConfig()
	if err != nil {
		response.ToJson(response.ERROR, gin.H{}, fmt.Sprintf("获取失败，%v", err), c)
	} else {
		response.ToJson(response.SUCCESS, resp.SysConfigResponse{Config: config}, "成功", c)
	}
}

// 设置配置文件内容
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
