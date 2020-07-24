package request

// 注册请求结构
type RegisterStruct struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	NickName  string `json:"nick_name"`
	Mail      string `json:"mail"`
	InviteUid string `json:"invited_id"`
}

//登录请求结构
type LoginStruct struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captcha_id"`
}

// 修改密码请求结构
type ChangePasswordStruct struct {
	ApiToken    string `json:"api_token"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

// 修改角色
type SetUserAuth struct {
	Username string `json:"username"`
	RoleUid  string `json:"role_id"`
}
