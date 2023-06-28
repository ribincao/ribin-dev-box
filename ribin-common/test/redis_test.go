package test

import (
	"testing"

	"github.com/ribincao/ribin-dev-box/ribin-common/config"
	"github.com/ribincao/ribin-dev-box/ribin-common/db"
	"github.com/ribincao/ribin-dev-box/ribin-common/logger"
	"go.uber.org/zap"
)

func initLocalEnv() {
	config.InitConfig("./conf.yaml")
	logger.InitLogger(config.GlobalConfig.LogConfig)
	db.InitRedis()
}

func TestRedisSetAndGet(t *testing.T) {
	initLocalEnv()
	val, err := db.RedisGlobal.TestSet()
	logger.Info("TestSet", zap.Any("value", val), zap.Any("err", err))
	val, err = db.RedisGlobal.TestGet()
	logger.Info("TestSet", zap.Any("value", val), zap.Any("err", err))
}
