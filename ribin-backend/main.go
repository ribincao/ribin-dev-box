package main

import (
	"fmt"
	"net"

	"github.com/ribincao/ribin-dev-box/ribin-backend/handler"
	"github.com/ribincao/ribin-dev-box/ribin-common/config"
	"github.com/ribincao/ribin-dev-box/ribin-common/logger"
	serverData "github.com/ribincao/ribin-dev-box/ribin-protocol/server-data"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func init() {
	config.InitConfig("./conf.yaml")
	logger.InitLogger(config.GlobalConfig.LogConfig)
}

func main() {
	port := fmt.Sprintf(":%s", config.GlobalConfig.ServiceConfig.Port)
	listen, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	serverData.RegisterServerDataServer(server, &handler.EngineServer{})

	logger.Info("Backend-Server Start Success", zap.Any("Port", port))
	if err = server.Serve(listen); err != nil {
		panic(err)
	}
}
