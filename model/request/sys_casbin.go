package request

// Casbin
type CasbinInfo struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

// Casbin structure for input parameters
type CasbinUpdate struct {
	RoleId      string       `json:"roleid"`
	CasbinInfos []CasbinInfo `json:"casbinInfos"`
}

//添加/修改casbin请求
type CasbinEdit struct {
	ID       uint   `json:"id"`
	ApiToken string `json:"api_token"`
	RoleUid  string `json:"role_uid"`
	Path     string `json:"path"`
	Method   string `json:"method"`
}

//仅包含角色ID的请求
type CasbinGet struct {
	ApiToken string `json:"api_token"`
	RoleUid  string `json:"role_uid"`
}
