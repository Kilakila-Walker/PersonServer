package model

import (
	"github.com/jinzhu/gorm"
)

type Sys_Role struct {
	gorm.Model
	Name     string `json:"name" gorm:"comment:'角色名';type:varchar(50)"`
	ParentId string `json:"parentId" gorm:"comment:'父角色ID';type:varchar(50)"`
	Desc     string `json:"desc" gorm:"comment:'角色描述';type:varchar(200)"`
}
