package model

import (
	"github.com/jinzhu/gorm"
)

//评论表
type Tab_Comment struct {
	gorm.Model
	FileInfoId string `json:"fileInfoId" gorm:"comment:'文件信息ID';type:varchar(50)"`
	UserId     string `json:"userId" gorm:"comment:'用户ID';type:varchar(50)"`
	UserNick   string `json:"userNick" gorm:"comment:'用户昵称';type:varchar(50)"`
	UserPic    string `json:"userPic" gorm:"comment:'用户头像';type:varchar(255)"`
	Comment    string `json:"comment" gorm:"comment:'评论内容';type:type:varchar(400)"`
}
