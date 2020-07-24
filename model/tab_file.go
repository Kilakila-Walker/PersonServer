package model

import (
	"github.com/jinzhu/gorm"
)

type Tab_File struct {
	gorm.Model
	RelatedId string `json:"relatedId" gorm:"comment:'文件与信息关联ID';type:varchar(50)"`
	Name      string `json:"name" gorm:"comment:'文件名';type:varchar(50)"`
	FileType  string `json:"fileType" gorm:"comment:'文件类型';type:varchar(50)"`
	UserId    uint   `json:"userId" gorm:"comment:'所属用户ID';type:varchar(50)"`
}
