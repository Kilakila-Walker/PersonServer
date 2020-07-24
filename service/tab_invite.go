package service

//sql 尽可能参数化 不然会有注入风险
import (
	"perServer/global"
	"perServer/model"
	"perServer/model/request"

	uuid "github.com/satori/go.uuid"
)

//添加邀请码
func AddInvite(m model.Tab_Invite) (err error, invite model.Tab_Invite) {
	var model model.Tab_Invite
	m.Uid = uuid.NewV1().String()
	err = global.GVA_DB.Model(&model).Create(&m).Error
	return err, model
}

//查看邀请码
func GetInvite(info request.PageInfo) (err error, list interface{}, total int) {
	var invite model.Tab_Invite
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&invite)
	var invites []model.Tab_Invite
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&invites).Error
	return err, invites, total
}
