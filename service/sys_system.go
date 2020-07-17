package service

import (
	"perServer/global"
	"perServer/model"
	"perServer/model/config"
	"perServer/utils"
)

// 读取配置文件
func GetSystemConfig() (err error, conf config.Server) {
	return nil, global.GVA_CONFIG
}

// 设置配置文件
func SetSystemConfig(system model.System) (err error) {
	cs := utils.StructToMap(system.Config)
	for k, v := range cs {
		global.GVA_VP.Set(k, v)
	}
	err = global.GVA_VP.WriteConfig()
	return err
}
