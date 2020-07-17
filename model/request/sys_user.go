package request

import uuid "github.com/satori/go.uuid"

// 注册请求结构
type RegisterStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nick_name"`
	Mail     string `json:"mail"`
	InviteId string `json:"InvitedId"`
}

//登录请求结构
type LoginStruct struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

// 修改密码请求结构
type ChangePasswordStruct struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// 修改角色
type SetUserAuth struct {
	UUID        uuid.UUID `json:"uuid"`
	AuthorityId string    `json:"authorityId"`
}
