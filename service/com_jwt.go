package service

import (
	"perServer/global"
	"perServer/model/common"
)

// 获取redis jwt
func GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.GVA_REDIS.Get(userName).Result()
	return err, redisJWT
}

// 设置redis jwt
func SetRedisJWT(jwtList common.Com_Jwt, userName string) (err error) {
	err = global.GVA_REDIS.Set(userName, jwtList.Jwt, 1000*1000*1000*60*60*24*7).Err()
	return err
}
