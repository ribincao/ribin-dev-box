package pool

import (
	"context"
	"fmt"
	"time"

	"github.com/ribincao/ribin-dev-box/ribin-common/config"
	"github.com/ribincao/ribin-dev-box/ribin-common/logger"
	grpcPool "github.com/ribincao/ribin-dev-box/ribin-common/pool"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var ServerDataPool *grpcPool.Pool

func InitRpcPool() {
	var err error
	xrayInterceptors := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	ServerDataPool, err = grpcPool.Init(func() (*grpc.ClientConn, error) {
		return grpc.Dial(config.GlobalConfig.ServiceConfig.BackendAddr, xrayInterceptors...)
	}, 10, 100, 10*time.Second, 5*time.Minute, time.Second, grpcPool.PoolGetModeLoose)
	if err != nil {
		fmt.Println("grpc engine data pool init error", err)
	}
}

func GetRpcClient(ctx context.Context, pool *grpcPool.Pool) (*grpcPool.Client, error) {
	conn, err := pool.Get(ctx)
	if err != nil {
		logger.Error("grpc conn error", zap.Error(err))
		return nil, err
	}
	return conn, err
}
