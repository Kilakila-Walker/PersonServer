package initialize

import (
	"perServer/global"
	"perServer/model"
)

// 数据库表迁移（结构转数据表)
func DBTables() {
	db := global.GVA_DB
	db.AutoMigrate(model.Sys_User{},
		model.ExaFile{},
		model.ExaFileChunk{},
	)
	global.GVA_LOG.Debug("register table success")
}
