package model

import (
	"github.com/jinzhu/gorm"
)

type Tab_FilesInfo struct {
	gorm.Model
	ClientId  string `json:"clientId" gorm:"comment:'客户端ID';type:varchar(50)"`
	RelatedId string `json:"relatedId" gorm:"comment:'文件与信息关联ID';type:varchar(50)"`
	Title     string `json:"title" gorm:"comment:'标题';type:varchar(50)"`
	Desc      string `json:"desc" gorm:"comment:'desc';type:varchar(500)"`
	UpUser    string `json:"upUser" gorm:"comment:'上传用户';type:varchar(50)"`
	Torrent   string `json:"torrent" gorm:"comment:'种子信息';type:varchar(200)"`
}
