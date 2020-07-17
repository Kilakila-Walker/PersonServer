package service

import (
	"errors"
	"perServer/global"
	"perServer/model"
	"perServer/model/request"
	"strings"

	"github.com/casbin/casbin"
	"github.com/casbin/casbin/util"
	gormadapter "github.com/casbin/gorm-adapter"
)

// 更新casbin权限
func UpdateCasbin(authorityId string, casbinInfos []request.CasbinInfo) error {
	ClearCasbin(0, authorityId)
	for _, v := range casbinInfos {
		cm := model.Sys_Casbin{
			RoleId: 0,
			Path:   v.Path,
			Method: v.Method,
		}
		addflag := AddCasbin(cm)
		if addflag == false {
			return errors.New("存在相同api,添加失败,请联系管理员")
		}
	}
	return nil
}

// 添加权限

func AddCasbin(cm model.Sys_Casbin) bool {
	e := Casbin()
	return e.AddPolicy(cm.RoleId, cm.Path, cm.Method)
}

// API更新随动

func UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	var cs []model.Sys_Casbin
	err := global.GVA_DB.Table("casbin_rule").Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Find(&cs).Updates(map[string]string{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	return err
}

// 获取权限列表

func GetPolicyPathByAuthorityId(authorityId string) (pathMaps []request.CasbinInfo) {
	e := Casbin()
	list := e.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

// 清除匹配的权限
func ClearCasbin(v int, p ...string) bool {
	e := Casbin()
	return e.RemoveFilteredPolicy(v, p...)

}

// 持久化到数据库  引入自定义规则
func Casbin() *casbin.Enforcer {
	a := gormadapter.NewAdapterByDB(global.GVA_DB)
	e := casbin.NewEnforcer(global.GVA_CONFIG.Casbin.ModelPath, a)
	e.AddFunction("ParamsMatch", ParamsMatchFunc)
	_ = e.LoadPolicy()
	return e
}

// 自定义规则函数
func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}

// 自定义规则函数
func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return ParamsMatch(name1, name2), nil
}
