package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ribincao/ribin-dev-box/ribin-cgi/handler"
	"github.com/ribincao/ribin-dev-box/ribin-common/config"
	"github.com/ribincao/ribin-dev-box/ribin-common/logger"
)

func init() {
	config.InitConfig("./conf.yaml")
	logger.InitLogger(config.GlobalConfig.LogConfig)
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
}
