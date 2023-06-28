package test

import (
	"testing"
	"time"

	"github.com/ribincao/ribin-dev-box/ribin-common/config"
	"github.com/ribincao/ribin-dev-box/ribin-common/logger"
	"github.com/ribincao/ribin-dev-box/ribin-common/timer"
	"go.uber.org/zap"
)

func InitLocalTimerTestEnv() {
	config.InitConfig("./conf.yaml")
	logger.InitLogger(config.GlobalConfig.LogConfig)
}

func TestTimerWheel(t *testing.T) {
	InitLocalTimerTestEnv()
	tw, _ := timer.NewTimeWheel(time.Millisecond*100, 4, timer.TickSafeMode())
	tw.Start()
	defer tw.Stop()

	i := 0
	tw.AddTask("Circle", time.Second*1, func() {
		logger.Info("-", zap.Any("i", i), zap.Any("now", time.Now().Second()))
		i++
	}, timer.CIRCLE_MODE)

	tw.AddTask("NotCircle", time.Second*10, func() {
		logger.Info("+ end", zap.Any("now", time.Now().Second()))
	}, timer.NOT_CIRCLE_MODE)

	time.Sleep(time.Second * 15)
}
