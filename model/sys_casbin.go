package model

import (
	"github.com/jinzhu/gorm"
)

//权限表 根据角色划分权限
type Sys_Casbin struct {
	gorm.Model
	RoleId uint   `json:"role_id" gorm:"comment:'角色id'"`
	Path   string `json:"path" gorm:"comment:'API路径';type:varchar(50)"`
	Method string `json:"method" gorm:"comment:'访问形式';type:varchar(50)"`
}
