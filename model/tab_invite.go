package model

import "github.com/jinzhu/gorm"

type Tab_Invite struct {
	gorm.Model
	Uid        string `json:"uid" gorm:"comment:'邀请UID'"`
	Available  int    `json:"available" gorm:"comment:'可用次数'"`
	InviteUser string `json:"invite_user" gorm:"comment:'邀请人UID';type:varchar(50)"`
}
