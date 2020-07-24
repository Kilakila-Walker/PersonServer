package common

import (
	"github.com/dgrijalva/jwt-go"
)

// jwt structure
type JWToken struct {
	Uuid     string
	ID       uint
	NickName string
	RoleUid  string
	jwt.StandardClaims
}
