package test

import (
	"testing"

	"github.com/ribincao/ribin-dev-box/ribin-common/config"
	"github.com/ribincao/ribin-dev-box/ribin-common/logger"

	"go.uber.org/zap"
)

func TestLog(t *testing.T) {
	config.InitConfig("./conf.yaml")
	logger.InitLogger(config.GlobalConfig.LogConfig)
	logger.Debug("DebugTest :", zap.String("Env", config.GlobalConfig.ServiceConfig.Env))
	logger.Info("InfoTest :", zap.String("Env", config.GlobalConfig.ServiceConfig.Env))
	logger.Error("ErrorTest :", zap.String("Env", config.GlobalConfig.ServiceConfig.Env))
	logger.Warn("WarnTest :", zap.String("Env", config.GlobalConfig.ServiceConfig.Env))
	logger.Fatal("FatalTest :", zap.String("Env", config.GlobalConfig.ServiceConfig.Env))
}
