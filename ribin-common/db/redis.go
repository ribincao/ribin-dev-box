package db

import (
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/ribincao/ribin-dev-box/ribin-common/config"
	"github.com/ribincao/ribin-dev-box/ribin-common/logger"
)

var (
	openRedisDB sync.Once
	RedisGlobal *RedisClient
)

type RedisClient struct {
	Client *redis.Client
}

type RedisCluster struct {
	Client *redis.ClusterClient
}

func InitRedis() {
	RedisGlobal = &RedisClient{}
	RedisGlobal.Client = redis.NewClient(&redis.Options{
		Addr:     config.GlobalConfig.ServiceConfig.RedisAddr,
		Password: config.GlobalConfig.ServiceConfig.RedisPasswd,
		DB:       0,
	})
	logger.Info("[Engine-Tool] Redis Client Initialized!")
}

func (client *RedisClient) TestGet() (string, error) {
	val, err := client.Client.Get("ping").Result()
	return val, err
}
func (client *RedisClient) TestSet() (string, error) {
	val, err := client.Client.Set("ping", "pong", time.Second*120).Result()
	return val, err
}
