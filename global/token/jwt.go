package token

//JWT
import (
	"errors"
	"perServer/global"
	"perServer/model/common"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")            //过期
	TokenNotValidYet = errors.New("Token not active yet")        //失活
	TokenMalformed   = errors.New("That's not even a token")     //非token
	TokenInvalid     = errors.New("Couldn't handle this token:") //修改
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.GVA_CONFIG.JWT.SigningKey),
	}
}

//能获取jwt封装的一些数据  请在jwt验证之后再使用
func GetClaims(c *gin.Context) (*common.JWToken, int) {
	var result *common.JWToken
	var code = 0
	x_token := c.Request.Header.Get("x-token")
	if x_token == "" {
		code = -1
		return result, code
	}
	j := NewJWT()
	token, _ := jwt.ParseWithClaims(x_token, &common.JWToken{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if token != nil {
		result, ok := token.Claims.(*common.JWToken)
		if ok && token.Valid {
			return result, 0
		}
	}
	code = -1
	return result, code
}

// 创建一个token
func (j *JWT) CreateJwt(claims common.JWToken) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析 token
func (j *JWT) ParseJwt(tokenString string) (*common.JWToken, error) {
	token, err := jwt.ParseWithClaims(tokenString, &common.JWToken{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*common.JWToken); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}

}
