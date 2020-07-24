package model

import "github.com/jinzhu/gorm"

//首字母大写为public
//数据库保存：大写字母会按照蛇形命名法进行修改 NickName->nick_name
//gorm字段tags见http://gorm.io/docs/models.html
type Sys_User struct {
	gorm.Model
	Uuid      string `json:"uuid" gorm:"comment:'用户UUID';type:varchar(50)"`
	Username  string `json:"username" gorm:"comment:'用户登录名';type:varchar(50)"`
	Password  string `json:"-"  gorm:"comment:'用户登录密码';type:varchar(50)"`
	NickName  string `json:"nick_name" gorm:"comment:'用户昵称';type:varchar(50)" `
	HeaderImg string `json:"header_img" gorm:"comment:'用户头像';type:varchar(250)"`
	Mail      string `json:"mail" gorm:"comment:'邮箱';type:varchar(100)"`
	RoleUid   string `json:"role_uid" gorm:"comment:'角色UID';type:varchar(50)"`
	InviteUid string `json:"invite_uid" gorm:"comment:'邀请人UID';type:varchar(50)"`
}
