package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ribincao/ribin-dev-box/ribin-cgi/handler"
	"github.com/ribincao/ribin-dev-box/ribin-cgi/pool"
	"github.com/ribincao/ribin-dev-box/ribin-common/config"
	"github.com/ribincao/ribin-dev-box/ribin-common/logger"
	"go.uber.org/zap"
)

func init() {
	config.InitConfig("./conf.yaml")
	logger.InitLogger(config.GlobalConfig.LogConfig)
	pool.InitRpcPool()
}

func main() {
	engine := gin.Default()
	group := engine.Group("engine")
	{
		group.GET("testGet", handler.TestGet)
		group.POST("testPost", handler.TestPost)
	}

	userGroup := engine.Group("engine/user")
	{
		userGroup.GET("/testGet", handler.TestGet)
		userGroup.GET("/testPost", handler.TestPost)
	}

	port := fmt.Sprintf(":%d", config.GlobalConfig.ServiceConfig.Port)
	logger.Info("Cgi-Server Start Success", zap.Any("Port", port))
	if err := engine.Run(port); err != nil {
		panic(err)
	}
}
