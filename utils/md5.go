package utils

import (
	"crypto/md5"
	"encoding/hex"
)
//获取md5
func MD5V(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}
