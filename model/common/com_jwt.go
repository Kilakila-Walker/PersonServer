package common

import (
	"github.com/jinzhu/gorm"
)

type Com_Jwt struct {
	gorm.Model
	Jwt string `json:"jwt"`
}
