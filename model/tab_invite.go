package model

import "github.com/jinzhu/gorm"

type Tab_Invite struct {
	gorm.Model
	Available  int    `json:"available" gorm:"comment:'可用次数'"`
	InviteUser string `json:"invite_user" gorm:"comment:'邀请人';type:varchar(50)"`
}
