package model

//权限表 根据角色划分权限
type Sys_Casbin struct {
	ID      uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	RoleUid string `json:"role_uid" gorm:"comment:'角色uid';type:varchar(50);"`
	Path    string `json:"path" gorm:"comment:'路径';type:varchar(50);"`
	Method  string `json:"method" gorm:"comment:'方式';type:varchar(50);"`
}
