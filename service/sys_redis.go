package service

import (
	"perServer/global"
	"time"
)

// 获取redis
func GetRedis(key string) (err error, value string) {
	value, err = global.GVA_REDIS.Get(key).Result()
	return err, value
}

//删除redis
func DelRedis(key string) (err error) {
	err = global.GVA_REDIS.Del(key).Err()
	return err
}

// 设置redis
func SetRedis(value string, key string, countMin int) (err error) {
	duration := 1000 * 1000 * 1000 * countMin
	err = global.GVA_REDIS.Set(key, value, time.Duration(duration)).Err()
	return err
}
