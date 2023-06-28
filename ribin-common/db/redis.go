package db

import (
	"crypto/tls"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"github.com/ribincao/ribin-dev-box/ribin-common/config"
	"github.com/ribincao/ribin-dev-box/ribin-common/constant"
	"github.com/ribincao/ribin-dev-box/ribin-common/logger"
)

var (
	RedisGlobal IRedis
)

type ClientType int32
type ClusterOptions func(*redis.ClusterOptions)

const (
	T_Client  ClientType = 1
	T_Cluster ClientType = 2
	NilError             = redis.Nil
)

func InitRedis() {
	if config.GlobalConfig.ServiceConfig.Env == constant.ENV_LOCAL {
		RedisGlobal = NewRedisClient(
			config.GlobalConfig.ServiceConfig.RedisAddr,
			config.GlobalConfig.ServiceConfig.RedisUserName,
			config.GlobalConfig.ServiceConfig.RedisPasswd,
			T_Client)
	} else {
		RedisGlobal = NewRedisClient(
			config.GlobalConfig.ServiceConfig.RedisAddr,
			config.GlobalConfig.ServiceConfig.RedisUserName,
			config.GlobalConfig.ServiceConfig.RedisPasswd,
			T_Cluster)
	}
	logger.Info("[Engine-Tool] Redis Client Initialized!")
}

func NewRedisClient(addr, userName, userPwd string, clientType ClientType) IRedis {
	if clientType == T_Client {
		c := &RedisClient{}
		c.Client = redis.NewClient(&redis.Options{
			Addr:     addr,
			Username: userName,
			Password: userPwd,
		})
		return c
	}
	cc := createClusterClient(addr, userName, userPwd)
	rs := redsync.New(goredis.NewPool(cc))
	return &RedisCluster{Client: cc, address: addr, rs: rs}
}

func createClusterClient(addr, username, password string, cusClusterOptions ...ClusterOptions) *redis.ClusterClient {
	opts := &redis.ClusterOptions{
		Addrs:          []string{addr},
		Password:       password,
		Username:       username,
		MaxRedirects:   3,
		ReadOnly:       false,
		RouteByLatency: false,

		MaxRetries:      0,
		MinRetryBackoff: 8 * time.Millisecond,
		MaxRetryBackoff: 512 * time.Millisecond,

		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolTimeout:  4 * time.Second,

		PoolSize:     15,
		MinIdleConns: 1,

		IdleCheckFrequency: 60 * time.Second,
		IdleTimeout:        5 * time.Minute,
		MaxConnAge:         0 * time.Second,

		TLSConfig: &tls.Config{},
	}

	for _, cusOpt := range cusClusterOptions {
		cusOpt(opts)
	}

	return redis.NewClusterClient(opts)
}
