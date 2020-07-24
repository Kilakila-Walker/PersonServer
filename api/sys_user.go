package api

//api处理数据组装、调用、返回的关系处理
//用户表API
import (
	"fmt"
	"perServer/global"
	"perServer/global/response"
	"perServer/global/token"
	"perServer/model"
	"perServer/model/common"
	"perServer/model/request"
	resp "perServer/model/response"
	"perServer/service"
	"perServer/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
)

// 用户注册账号
func Register(c *gin.Context) {
	var R request.RegisterStruct
	_ = c.ShouldBindJSON(&R)
	UserVerify := utils.Rules{
		"Username": {utils.NotEmpty()},
		"NickName": {utils.NotEmpty()},
		"Password": {utils.NotEmpty()},
		"Mail":     {utils.NotEmpty()},
		"InviteId": {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(R, UserVerify)
	if UserVerifyErr != nil {
		response.ToJson(response.ERROR, gin.H{}, UserVerifyErr.Error(), c)
		return
	}
	user := &model.Sys_User{Username: R.Username, NickName: R.NickName, Password: R.Password, Mail: R.Mail}
	err, userReturn := service.Register(*user)
	if err != nil {
		response.ToJson(response.ERROR, userReturn, fmt.Sprintf("%v", err), c)
	} else {
		response.ToJson(response.SUCCESS, userReturn, "注册成功", c)
	}
}

// 用户登录
func Login(c *gin.Context) {
	var L request.LoginStruct
	_ = c.ShouldBindJSON(&L)
	UserVerify := utils.Rules{
		"CaptchaId": {utils.NotEmpty()},
		"Captcha":   {utils.NotEmpty()},
		"Username":  {utils.NotEmpty()},
		"Password":  {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(L, UserVerify)
	if UserVerifyErr != nil {
		response.ToJson(response.ERROR, gin.H{}, UserVerifyErr.Error(), c)
		return
	}
	if store.Verify(L.CaptchaId, L.Captcha, true) {
		U := &model.Sys_User{Username: L.Username, Password: L.Password}
		if err, user := service.Login(U); err != nil {
			response.ToJson(response.ERROR, gin.H{}, fmt.Sprintf("用户名密码错误或%v", err), c)
			return
		} else {
			tokenNext(c, *user)
		}
	} else {
		response.ToJson(response.ERROR, gin.H{}, "验证码错误", c)
	}

}

// 登录以后签发jwt
func tokenNext(c *gin.Context, user model.Sys_User) {
	countMin := 60 * 60 * 24 * 7 //一周
	err, oldtoken := service.GetRedis(user.Username)
	j := token.NewJWT()
	clams := common.JWToken{
		Uuid:     user.Uuid,
		ID:       user.ID,
		NickName: user.NickName,
		RoleUid:  user.RoleUid,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),                   //签发时间
			NotBefore: time.Now().Unix() - 1000,            // 签名生效时间
			ExpiresAt: time.Now().Unix() + int64(countMin), // 过期时间
			Issuer:    "admin",                             // 签名的发行者
		},
	}
	newtoken, err := j.CreateJwt(clams) //创建新token
	if err == redis.Nil {               //不存在这个token
		service.SetRedis(newtoken, user.Username, countMin) //设置新的token
		response.ToJson(
			response.SUCCESS,
			resp.LoginResponse{
				ID:        user.ID,
				Uuid:      user.Uuid,
				Username:  user.Username,
				NickName:  user.NickName,
				HeaderImg: user.HeaderImg,
				Mail:      user.Mail,
				RoleUid:   user.RoleUid,
				Token:     newtoken,
			},
			"成功",
			c)
		return
	} else if err != nil { //执行错误
		response.ToJson(response.ERROR, gin.H{}, "获取token失败", c)
		return
	} else { //存在这个jwt
		if !global.GVA_CONFIG.System.UseMultipoint { //采用多点登录

			response.ToJson(
				response.SUCCESS,
				resp.LoginResponse{
					ID:        user.ID,
					Uuid:      user.Uuid,
					Username:  user.Username,
					NickName:  user.NickName,
					HeaderImg: user.HeaderImg,
					Mail:      user.Mail,
					Token:     oldtoken,
				},
				"成功",
				c)
			return
		} else { //不采用多点登录时使用新的token旧的token过期
			service.SetRedis(newtoken, user.Username, countMin)
			response.ToJson(
				response.SUCCESS,
				resp.LoginResponse{
					ID:        user.ID,
					Uuid:      user.Uuid,
					Username:  user.Username,
					NickName:  user.NickName,
					HeaderImg: user.HeaderImg,
					Mail:      user.Mail,
					Token:     newtoken,
				},
				"成功",
				c)
			return
		}
	}
}

// 用户修改密码
func ChangePassword(c *gin.Context) {
	//绑定
	var params request.ChangePasswordStruct
	_ = c.ShouldBindJSON(&params)
	//验证规则
	UserVerify := utils.Rules{
		"api_token":    {utils.NotEmpty()},
		"username":     {utils.NotEmpty()},
		"password":     {utils.NotEmpty()},
		"new_password": {utils.NotEmpty()},
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
	//组装数据并调用修改service
	U := &model.Sys_User{Username: params.Username, Password: params.Password}
	err, _ := service.ChangePassword(U, params.NewPassword)
	if err != nil {
		response.ToJson(response.ERROR, gin.H{}, "修改失败，请检查用户名密码", c)
		return
	} else {
		response.ToJson(response.SUCCESS, gin.H{}, "修改成功", c)
		return
	}
}

// 用户上传头像
func UploadHeaderImg(c *gin.Context) {
	//绑定
	var params request.ApitokenOnly
	_ = c.ShouldBindJSON(&params)
	//验证规则
	UserVerify := utils.Rules{
		"api_token": {utils.NotEmpty()},
	}
	//验证
	UserVerifyErr := utils.Verify(params, UserVerify)
	if UserVerifyErr != nil {
		response.ToJson(response.ERROR, gin.H{}, UserVerifyErr.Error(), c)
		return
	}
	req := c.Request
	uid := uuid.NewV4().String()
	// 获取token信息
	waitUse, _ := token.GetClaims(c)
	roleUid := waitUse.RoleUid
	uuid := waitUse.Uuid
	// 获取请求的URI
	uri := c.Request.URL.RequestURI()
	//验证此次提交是否有申请过token
	verCode := token.ApiTokenVeri(uri+roleUid, params.ApiToken)
	if verCode != 0 {
		response.ToJson(response.ERROR, gin.H{}, "不存在该请求记录", c)
		return
	}
	err, filePath, code := utils.Upload(uid, req)
	if code == -1 {
		response.ToJson(response.ERROR, gin.H{}, "上传方式错误", c)
		return
	} else if code == -2 {
		response.ToJson(response.ERROR, gin.H{}, "缺少postFile内容", c)
		return
	} else if code == -3 {
		response.ToJson(response.ERROR, gin.H{}, "请上传jpg/png/gif格式的图片", c)
		return
	} else if code == -4 {
		response.ToJson(response.ERROR, gin.H{}, "后台打开文件失败", c)
		return
	}
	// 获取头像文件
	if err != nil {
		response.ToJson(response.ERROR, gin.H{}, fmt.Sprintf("后台错误，%v", err), c)
		return
	} else {
		// 修改数据库后得到修改后的user并且返回供前端使用
		err, user := service.UploadHeaderImg(uuid, filePath)
		if err != nil {
			response.ToJson(response.ERROR, gin.H{}, fmt.Sprintf("修改数据库链接失败，%v", err), c)
			return
		} else {
			response.ToJson(response.ERROR, resp.SysUserResponse{User: *user}, "成功", c)
			return
		}
	}

}

// 分页获取用户列表
func GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	//验证规则
	UserVerify := utils.Rules{
		"api_token": {utils.NotEmpty()},
		"page":      {utils.NotEmpty()},
		"pageSize":  {utils.NotEmpty()},
	}
	//验证
	UserVerifyErr := utils.Verify(pageInfo, UserVerify)
	if UserVerifyErr != nil {
		response.ToJson(response.ERROR, gin.H{}, UserVerifyErr.Error(), c)
		return
	}
	// 获取token信息
	waitUse, _ := token.GetClaims(c)
	uuid := waitUse.Uuid
	// 获取请求的URI
	uri := c.Request.URL.RequestURI()
	//验证此次提交是否有申请过token
	verCode := token.ApiTokenVeri(uri+uuid, pageInfo.ApiToken)
	if verCode != 0 {
		response.ToJson(response.ERROR, gin.H{}, "不存在该请求记录", c)
		return
	}
	//获取用户信息
	err, list, total := service.GetUserInfoList(pageInfo)
	if err != nil {
		response.ToJson(response.ERROR, gin.H{}, fmt.Sprintf("获取数据失败，%v", err), c)
		return
	} else {
		response.ToJson(response.SUCCESS, resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "成功", c)
		return
	}
}
