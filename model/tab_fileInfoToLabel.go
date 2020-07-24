package model

import (
	"github.com/jinzhu/gorm"
)

//文件信息与标签关联表
type Tab_FileInfoToLabel struct {
	gorm.Model
	FileinfoId uint   `json:"fileinfo_id" gorm:"comment:'文件信息ID';type:varchar(50)"`
	LabelId    uint   `json:"label_id" gorm:"comment:'标签ID';type:varchar(50)"`
	LabelName  string `json:"label_name" gorm:"comment:'标签名';type:varchar(50)"`
}
