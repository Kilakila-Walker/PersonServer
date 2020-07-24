package api

//用户表API
import (
	"perServer/global"
	"perServer/global/response"
	"perServer/global/token"
	"perServer/model/request"
	"perServer/service"
	"perServer/utils"
	"time"

	"github.com/gin-gonic/gin"
)

//给每个表单发一个token 避免重复提交以及跨站攻击
func GetApiToken(c *gin.Context) {
	//绑定
	var params request.GetApiToken
	_ = c.ShouldBindJSON(&params)
	//验证规则
	UserVerify := utils.Rules{
		"uri":    {utils.NotEmpty()},
		"path":   {utils.NotEmpty()},
		"method": {utils.NotEmpty()},
	}
	//验证
	UserVerifyErr := utils.Verify(params, UserVerify)
	if UserVerifyErr != nil {
		response.ToJson(response.ERROR, gin.H{}, UserVerifyErr.Error(), c)
		return
	}
	waitUse, code := token.GetClaims(c)
	if code != 0 {
		response.ToJson(
			response.ERROR,
			gin.H{"reload": true},
			"未登录或非法访问",
			c)
		c.Abort()
		return
	}
	// 获取请求的API
	uri := params.Uri
	// 获取请求方法
	act := params.Method
	// 获取角色的UID
	uid := waitUse.RoleUid
	e := service.Casbin()
	// 判断策略中是否存在
	if global.GVA_CONFIG.System.Env != "develop" || !e.Enforce(uid, uri, act) {
		response.ToJson(response.ERROR, gin.H{}, "权限不足", c)
		c.Abort()
		return
	}
	valStr := uid + uri + act + time.Now().String() //这个公式告诉别人也很安全 反正服务器没有保存的东西不会通过验证
	token := utils.MD5V([]byte(valStr))             //value=MD5(uid+rui+act+time)
	key := "api" + uri + uid                        //key="api"+uri+uid 模块 uri uid
	err := service.SetRedis(token, key, 10)         //保存十分钟 十分钟内不提交表单，那是什么神奇的表单
	if err != nil {
		response.ToJson(response.ERROR, gin.H{}, "服务器遇到一些错误", c)
		return
	}
	response.ToJson(response.SUCCESS, gin.H{"api_token": token}, "需提供参数：api", c)
}
