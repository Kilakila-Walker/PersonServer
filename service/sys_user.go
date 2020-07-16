package service

import (
	"errors"
	"perServer/global"
	"perServer/model"
	"perServer/model/request"
	"perServer/utils"

	uuid "github.com/satori/go.uuid"
)

// @title    Register
// @description   register, 用户注册
// @auth                     （2020/04/05  20:22）
// @param     u               model.SysUser
// @return    err             error
// @return    userInter       *SysUser

func Register(u model.Sys_User) (err error, userInter model.Sys_User) {
	var user model.Sys_User
	// 判断用户名是否注册
	notRegister := global.GVA_DB.Where("username = ?", u.Username).First(&user).RecordNotFound()
	// notRegister为false表明读取到了 不能注册
	if !notRegister {
		return errors.New("用户名已注册"), userInter
	} else {
		// 否则 附加uuid 密码des加密 注册
		u.Password = utils.EncryDES_Str(u.Password)
		u.Uuid = uuid.NewV4().String()
		err = global.GVA_DB.Create(&u).Error
	}
	return err, u
}

// @title    Login
// @description   login, 用户登录
// @auth                     （2020/04/05  20:22）
// @param     u               *model.SysUser
// @return    err             error
// @return    userInter       *SysUser

func Login(u *model.Sys_User) (err error, userInter *model.Sys_User) {
	var user model.Sys_User
	u.Password = utils.EncryDES_Str(u.Password)
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	return err, &user
}

// @title    ChangePassword
// @description   change the password of a certain user, 修改用户密码
// @auth                     （2020/04/05  20:22）
// @param     u               *model.SysUser
// @param     newPassword     string
// @return    err             error
// @return    userInter       *SysUser

func ChangePassword(u *model.Sys_User, newPassword string) (err error, userInter *model.Sys_User) {
	var user model.Sys_User
	u.Password = utils.EncryDES_Str(u.Password)
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.EncryDES_Str(newPassword)).Error
	return err, u
}

// @title    GetInfoList
// @description   get user list by pagination, 分页获取数据
// @auth                      （2020/04/05  20:22）
// @param     info             request.PageInfo
// @return    err              error
// @return    list             interface{}
// @return    total            int

func GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.Sys_User{})
	var userList []model.Sys_User
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&userList).Error
	return err, userList, total
}

// @title    UploadHeaderImg
// @description   upload avatar, 用户头像上传更新地址
// @auth                     （2020/04/05  20:22）
// @param     uuid            UUID
// @param     filePath        string
// @return    err             error
// @return    userInter       *SysUser

func UploadHeaderImg(uuid uuid.UUID, filePath string) (err error, userInter *model.Sys_User) {
	var user model.Sys_User
	err = global.GVA_DB.Where("uuid = ?", uuid).First(&user).Update("header_img", filePath).First(&user).Error
	return err, &user
}
