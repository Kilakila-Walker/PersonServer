package model

import (
	"github.com/jinzhu/gorm"
)

// 字典表（下拉列表
type Sys_Dic struct {
	gorm.Model
	Field string `json:"field" gorm:"comment:'字段';type:varchar(50)"`
	Name  string `json:"name" gorm:"comment:'字典名';type:varchar(50)"`
	Desc  string `json:"desc" gorm:"comment:'字典描述';type:varchar(250)"`
}
