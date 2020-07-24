package model

import (
	"github.com/jinzhu/gorm"
)

//角色表
//初始化时请初始换一个uid为default的初始角色和一个uid为admin的管理员角色
type Sys_Role struct {
	gorm.Model
	Uid  string `json:"uid" gorm:"comment:'角色uid';type:varchar(50);unique_index"`
	Name string `json:"name" gorm:"comment:'角色名';type:varchar(50)"`
	Desc string `json:"desc" gorm:"comment:'角色描述';type:varchar(200)"`
}
