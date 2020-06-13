package initredis

import (
	"apt/conf"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

// 初始化连接
func InitClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     conf.Redisconfig(),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
