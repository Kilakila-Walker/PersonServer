package model

import (
	"github.com/jinzhu/gorm"
)

// 客户端表
type Sys_Client struct {
	gorm.Model
	UserId    string `json:"userId" gorm:"comment:'所属用户ID';type:varchar(50)"`
	State     int    `json:"state" gorm:"comment:'连接状态'"`
	Authority int    `json:"Authority" gorm:"comment:'连接权限'"`
	IP        string `json:"ip" gorm:"comment:'当前IP地址';type:varchar(50)"`
	Port      string `json:"port" gorm:"comment:'开放端口';type:varchar(50)"`
	Desc      string `json:"desc" gorm:"comment:'服务器描述';type:varchar(400)"`
	Pic       string `json:"pic" gorm:"comment:'服务器图片';type:varchar(255)"`
}
