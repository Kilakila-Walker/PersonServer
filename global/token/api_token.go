package token

import (
	"perServer/service"

	"github.com/go-redis/redis"
)

//根据key获取并验证apitoken 这里Key不需要添加"api"字符串
//value=MD5(uid+api+act+time)
//key="api"+api+uid 模块 uri uid
func ApiTokenVeri(key string, apitoken string) int {
	err, value := service.GetRedis("api" + key)
	if err == redis.Nil {
		return -1 //不存在
	} else if err != nil {
		return -3 //执行错误
	}
	if value != apitoken {
		return -2 //token错误
	}
	service.DelRedis(key)
	return 0 //存在且一致
}
