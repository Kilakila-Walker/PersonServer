package service

import (
	"perServer/global"
	"time"
)

// 获取redis jwt
func GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.GVA_REDIS.Get(userName).Result()
	return err, redisJWT
}

// 设置redis token
func SetRedisJWT(token string, userName string, countMin int) (err error) {
	duration := 1000 * 1000 * 1000 * countMin
	err = global.GVA_REDIS.Set(userName, token, time.Duration(duration)).Err()
	return err
}
