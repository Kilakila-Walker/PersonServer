package api

//权限API
import (
	"fmt"
	"perServer/global/response"
	"perServer/model/request"
	resp "perServer/model/response"
	"perServer/service"
	"perServer/utils"

	"github.com/gin-gonic/gin"
)

// 更改角色api权限
func UpdateCasbin(c *gin.Context) {
	var cmr request.CasbinInReceive
	_ = c.ShouldBindJSON(&cmr)
	AuthorityIdVerifyErr := utils.Verify(cmr, utils.CustomizeMap["AuthorityIdVerify"])
	if AuthorityIdVerifyErr != nil {
		response.ToJson(response.ERROR, gin.H{}, AuthorityIdVerifyErr.Error(), c)
		return
	}
	err := service.UpdateCasbin(cmr.AuthorityId, cmr.CasbinInfos)
	if err != nil {
		response.ToJson(response.ERROR, gin.H{}, fmt.Sprintf("添加规则失败，%v", err), c)
	} else {
		response.ToJson(response.SUCCESS, gin.H{}, "添加规则成功", c)
	}
}

// 获取权限列表
func GetPolicyPathByAuthorityId(c *gin.Context) {
	var cmr request.CasbinInReceive
	_ = c.ShouldBindJSON(&cmr)
	AuthorityIdVerifyErr := utils.Verify(cmr, utils.CustomizeMap["AuthorityIdVerify"])
	if AuthorityIdVerifyErr != nil {
		response.ToJson(response.ERROR, gin.H{}, AuthorityIdVerifyErr.Error(), c)
		return
	}
	paths := service.GetPolicyPathByAuthorityId(cmr.AuthorityId)
	response.ToJson(response.SUCCESS, resp.PolicyPathResponse{Paths: paths}, "成功", c)
}

// RESTFUL测试路由
func CasbinTest(c *gin.Context) {
	// 测试restful以及占位符代码  随意书写
	pathParam := c.Param("pathParam")
	query := c.Query("query")
	response.ToJson(response.SUCCESS, gin.H{"pathParam": pathParam, "query": query}, "获取规则成功", c)
}
