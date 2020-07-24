package response

import (
	"perServer/model"
)

type SysUserResponse struct {
	User model.Sys_User `json:"user"`
}

type LoginResponse struct {
	ID        uint   `json:"id"`
	Uuid      string `json:"uuid"`
	Username  string `json:"username"`
	NickName  string `json:"nick_name"`
	HeaderImg string `json:"header_img"`
	Mail      string `json:"mail"`
	Token     string `json:"token"`
	RoleUid   string `json:"role_uid"`
}
