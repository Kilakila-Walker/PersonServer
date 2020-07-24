package api

//权限API
import (
	"fmt"
	"perServer/global/response"
	"perServer/global/token"
	"perServer/model/request"
	"perServer/service"
	"perServer/utils"

	"github.com/gin-gonic/gin"
)

// 更改角色api权限
func UpdateCasbin(c *gin.Context) {
	//绑定
	var params request.CasbinEdit
	_ = c.ShouldBindJSON(&params)
	//验证规则
	UserVerify := utils.Rules{
		"id":        {utils.NotEmpty()},
		"role_uid":  {utils.NotEmpty()},
		"api_token": {utils.NotEmpty()},
		"method":    {utils.NotEmpty()},
		"path":      {utils.NotEmpty()},
	}
	//验证
	UserVerifyErr := utils.Verify(params, UserVerify)
	if UserVerifyErr != nil {
		response.ToJson(response.ERROR, gin.H{}, UserVerifyErr.Error(), c)
		return
	}
	// 获取token信息
	waitUse, _ := token.GetClaims(c)
	roleUid := waitUse.RoleUid
	// 获取请求的URI
	uri := c.Request.URL.RequestURI()
	//验证此次提交是否有申请过token
	verCode := token.ApiTokenVeri(uri+roleUid, params.ApiToken)
	if verCode != 0 {
		response.ToJson(response.ERROR, gin.H{}, "不存在该请求记录", c)
		return
	}
	err := service.UpdateCasbin(params)
	if err != nil {
		response.ToJson(response.ERROR, gin.H{}, fmt.Sprintf("添加规则失败，%v", err), c)
	} else {
		response.ToJson(response.SUCCESS, gin.H{}, "添加规则成功", c)
	}
}

//添加角色API权限
func AddCasbin(c *gin.Context) {
	//绑定
	var params request.CasbinEdit
	_ = c.ShouldBindJSON(&params)
	//验证规则
	UserVerify := utils.Rules{
		"role_uid":  {utils.NotEmpty()},
		"api_token": {utils.NotEmpty()},
		"method":    {utils.NotEmpty()},
		"path":      {utils.NotEmpty()},
	}
	//验证
	UserVerifyErr := utils.Verify(params, UserVerify)
	if UserVerifyErr != nil {
		response.ToJson(response.ERROR, gin.H{}, UserVerifyErr.Error(), c)
		return
	}
	// 获取token信息
	waitUse, _ := token.GetClaims(c)
	roleUid := waitUse.RoleUid
	// 获取请求的URI
	uri := c.Request.URL.RequestURI()
	//验证此次提交是否有申请过token
	verCode := token.ApiTokenVeri(uri+roleUid, params.ApiToken)
	if verCode != 0 {
		response.ToJson(response.ERROR, gin.H{}, "不存在该请求记录", c)
		return
	}
	err := service.UpdateCasbin(params)
	if err != nil {
		response.ToJson(response.ERROR, gin.H{}, fmt.Sprintf("添加规则失败，%v", err), c)
	} else {
		response.ToJson(response.SUCCESS, gin.H{}, "添加规则成功", c)
	}
}

// 获取权限列表
func GetPolicyPathByAuthorityId(c *gin.Context) {
	//绑定
	var params request.CasbinGet
	_ = c.ShouldBindJSON(&params)
	//验证规则
	UserVerify := utils.Rules{
		"role_uid":  {utils.NotEmpty()},
		"api_token": {utils.NotEmpty()},
	}
	//验证
	UserVerifyErr := utils.Verify(params, UserVerify)
	if UserVerifyErr != nil {
		response.ToJson(response.ERROR, gin.H{}, UserVerifyErr.Error(), c)
		return
	}
	// 获取token信息
	waitUse, _ := token.GetClaims(c)
	roleUid := waitUse.RoleUid
	// 获取请求的URI
	uri := c.Request.URL.RequestURI()
	//验证此次提交是否有申请过token
	verCode := token.ApiTokenVeri(uri+roleUid, params.ApiToken)
	if verCode != 0 {
		response.ToJson(response.ERROR, gin.H{}, "不存在该请求记录", c)
		return
	}
	paths := service.GetCasbin(params.RoleUid)
	response.ToJson(response.SUCCESS, paths, "", c)
}
