package test

import (
	"testing"
	"time"

	"github.com/ribincao/ribin-dev-box/ribin-common/config"
	"github.com/ribincao/ribin-dev-box/ribin-common/logger"
	"github.com/ribincao/ribin-dev-box/ribin-common/utils"
	"go.uber.org/zap"
)

func f() {
	var infoMap map[int]string
	infoMap[1] = "a"
}

func TestRecover(t *testing.T) {
	config.InitConfig("./conf.yaml")
	logger.InitLogger(config.GlobalConfig.LogConfig)
	utils.GoWithRecover(f)

	time.Sleep(2 * time.Second)
}

func TestGenerater(t *testing.T) {
	config.InitConfig("./conf.yaml")
	logger.InitLogger(config.GlobalConfig.LogConfig)
	code := utils.GenRoomCode()
	logger.Info("TestGenerater", zap.Any("Code", code))
}
