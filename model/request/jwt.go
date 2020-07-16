package request

import (
	"github.com/dgrijalva/jwt-go"
)

// Custom claims structure
type CustomClaims struct {
	Uuid     string
	ID       uint
	NickName string
	jwt.StandardClaims
}
