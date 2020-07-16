package model

import (
	"github.com/jinzhu/gorm"
)

// 标签表
type Sys_Label struct {
	gorm.Model
	Name     string `json:"name" gorm:"comment:'标签名';type:varchar(50)"`
	ParentId string `json:"ParentId" gorm:"comment:'父标签ID';type:varchar(50)"`
	Desc     string `json:"desc" gorm:"comment:'标签描述';type:varchar(250)"`
}
