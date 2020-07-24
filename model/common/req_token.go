package common

// post之前需要获取一个服务器提供的token，避免重复提交表单
type ApiToken struct {
	Api     string `json:"api"`
	UserUid string `json:"useruid"`
	Time    int64  `json:"time"`
}
