package model

import (
	"github.com/jinzhu/gorm"
)

//用户-角色关联表 一个用户可拥有多个角色 角色权限暂未设置
type Sys_RoleToUser struct {
	gorm.Model
	RoleId uint `json:"roleId" gorm:"comment:'角色ID';type:varchar(50)"`
	UserId uint `json:"userId" gorm:"comment:'用户ID';type:varchar(50)"`
}
