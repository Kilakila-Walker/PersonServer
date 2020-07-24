package request

// Paging common input parameter structure
type PageInfo struct {
	ApiToken string `json:"api_token"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}

// Find by id structure
type GetById struct {
	Id float64 `json:"id" form:"id"`
}

//请求apitoken的参数
type GetApiToken struct {
	Uri     string `json:"uri"`
	Path    string `json:"path"`
	Method  string `json:"method"`
	RoleUid string `json:"role_uid"`
}

//请求只包含apitoken
type ApitokenOnly struct {
	ApiToken string `json:"api_token"`
}
