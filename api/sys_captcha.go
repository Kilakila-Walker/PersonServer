package api

//验证码API
import (
	"fmt"
	"perServer/global"
	"perServer/global/response"
	resp "perServer/model/response"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

// 生成验证码
func Captcha(c *gin.Context) {
	//字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.GVA_CONFIG.Captcha.ImgHeight, global.GVA_CONFIG.Captcha.ImgWidth, global.GVA_CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		response.ToJson(response.ERROR, gin.H{}, fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.ToJson(
			response.SUCCESS,
			resp.SysCaptchaResponse{
				CaptchaId: id,
				PicPath:   b64s,
			}, "验证码获取成功", c)
	}
}
