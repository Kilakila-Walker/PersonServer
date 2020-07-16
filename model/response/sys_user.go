package response

import (
	"perServer/model"
)

type SysUserResponse struct {
	User model.Sys_User `json:"user"`
}

type LoginResponse struct {
	User  model.Sys_User `json:"user"`
	Token string         `json:"token"`
}
