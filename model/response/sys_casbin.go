package response

//casbin
type Casbin_Res struct {
	RoleId string `json:"role_id"`
	Path   string `json:"path"`
	Method string `json:"method"`
}
