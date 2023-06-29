package main

import (
	"net"

	"github.com/ribincao/ribin-dev-box/ribin-backend/handler"
	"github.com/ribincao/ribin-dev-box/ribin-common/config"
	"github.com/ribincao/ribin-dev-box/ribin-common/logger"
	serverData "github.com/ribincao/ribin-dev-box/ribin-protocol/server-data"
	"google.golang.org/grpc"
)

func init() {
	config.InitConfig("./conf.yaml")
	logger.InitLogger(config.GlobalConfig.LogConfig)
}

func main() {
	listen, err := net.Listen("tcp", config.GlobalConfig.ServiceConfig.Port)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	serverData.RegisterServerDataServer(server, &handler.EngineServer{})

	if err = server.Serve(listen); err != nil {
		panic(err)
	}
}
