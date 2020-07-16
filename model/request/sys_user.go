package request

import uuid "github.com/satori/go.uuid"

// User register structure
type RegisterStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nick_name"`
	Mail     string `json:"mail"`
	InviteId string `json:"InvitedId"`
}

// User login structure
type LoginStruct struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

// Modify password structure
type ChangePasswordStruct struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// Modify  user's auth structure
type SetUserAuth struct {
	UUID        uuid.UUID `json:"uuid"`
	AuthorityId string    `json:"authorityId"`
}
