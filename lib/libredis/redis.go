package libredis

import (
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"gopkg.in/redis.v5"
)

var redisCache *redis.Client

func init() {
	redisHost, _ := web.AppConfig.String("redis.conn")
	dataBase, _ := web.AppConfig.Int("redis.dbNum")
	password, _ := web.AppConfig.String("redis.password")
	redisCache, _ = createClient(redisHost, password, dataBase)
}

// 创建 redis 客户端
func createClient(redisHost string, password string, dataBase int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: password,
		DB:       dataBase,
	})

	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	_, err := client.Ping().Result()
	if err != nil{
		logs.Error("redis connect failed, err: %v", err)
		return nil, err
	}

	return client, nil
}

func SetStr(key, value string, time time.Duration) error {
	if err := redisCache.Set(key, value, time).Err(); err != nil {
		logs.Error("[SetStr] redisCache.Set, err: %v, key: %v", err, key)
		return err
	}
	return nil
}

func GetStr(key string) (string, error) {
	v, err := redisCache.Get(key).Result()
	if err != nil {
		logs.Error("[GetStr] redisCache.Get, err: %v, key: %v", err, key)
		return "", err
	}
	return v, nil
}

func GetInt64(key string) (int64, error) {
	v, err := redisCache.Get(key).Int64()
	if err != nil {
		logs.Error("[GetInt64] redisCache.Get, err: %v, key: %v", err, key)
		return 0, err
	}
	return v, nil
}

func DelKey(key string) error {
	if err := redisCache.Del(key).Err(); err != nil {
		logs.Error("[DelKey] redisCache.Del, err: %v, key: %v", err, key)
		return err
	}
	return nil
}

func TTL(key string) (time.Duration, error) {
	ttl, err := redisCache.TTL(key).Result()
	if err != nil {
		logs.Error("[TTL] redisCache.TTL, err: %v, key: %v", err, key)
		return 0, err
	}
	return ttl, nil
}