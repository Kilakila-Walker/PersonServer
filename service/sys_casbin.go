package service

import (
	"perServer/global"
	"perServer/model"
	"perServer/model/request"
	"perServer/model/response"

	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
)

//更新权限
func UpdateCasbin(ce request.CasbinEdit) error {
	var cs model.Sys_Casbin
	upCasbin := model.Sys_Casbin{
		ID:      ce.ID,
		RoleUid: ce.RoleUid,
		Path:    ce.Path,
		Method:  ce.Method,
	}
	err := global.GVA_DB.Model(&cs).Where("id = ?", ce.ID).Updates(upCasbin).Error
	return err
}

// 添加权限
func AddCasbin(ce request.CasbinEdit) bool {
	e := Casbin()
	return e.AddPolicy(ce.RoleUid, ce.Path, ce.Method)
}

// API更新随动
func UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	var cs model.Sys_Casbin
	err := global.GVA_DB.Model(&cs).Where("path = ? AND method = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"path":   newPath,
		"method": newMethod,
	}).Error
	return err
}

// 获取权限列表
func GetCasbin(roleUid string) (CasbinList []response.Casbin_Res) {
	e := Casbin()
	list := e.GetFilteredPolicy(0, roleUid)
	for _, v := range list {
		CasbinList = append(CasbinList, response.Casbin_Res{
			RoleId: roleUid,
			Path:   v[1],
			Method: v[2],
		})
	}
	return CasbinList
}

//删除某项权限
func DelCasbin(ce request.CasbinEdit) bool {
	e := Casbin()
	return e.RemovePolicy(ce.RoleUid, ce.Path, ce.Method)
}

// 清除匹配的权限
func ClearCasbin(v int, p ...string) bool {
	e := Casbin()
	return e.RemoveFilteredPolicy(v, p...)
}

// 持久化到数据库
func Casbin() *casbin.Enforcer {
	a := gormadapter.NewAdapterByDB(global.GVA_DB)
	e := casbin.NewEnforcer(global.GVA_CONFIG.Casbin.ModelPath, a)
	_ = e.LoadPolicy()
	return e
}
