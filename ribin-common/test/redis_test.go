package test

import (
	"context"
	"testing"
	"time"

	"github.com/ribincao/ribin-dev-box/ribin-common/config"
	"github.com/ribincao/ribin-dev-box/ribin-common/db"
	"github.com/ribincao/ribin-dev-box/ribin-common/logger"
	"go.uber.org/zap"
)

func InitLocalRedisTestEnv() {
	config.InitConfig("./conf.yaml")
	logger.InitLogger(config.GlobalConfig.LogConfig)
	db.InitRedis()
}
func TestRedisString(t *testing.T) {
	InitLocalRedisTestEnv()
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*1)
	defer cancel()
	err := db.RedisGlobal.Set(ctx, "ping", "pong", time.Second*120)
	logger.Info("TestSet", zap.Any("err", err))
	val, err := db.RedisGlobal.Get(ctx, "ping")
	logger.Info("TestSet", zap.Any("value", val), zap.Any("err", err))
}
