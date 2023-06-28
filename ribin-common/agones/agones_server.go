package agones

import (
	"context"
	"time"

	"github.com/ribincao/ribin-dev-box/ribin-common/logger"

	coresdk "agones.dev/agones/pkg/sdk"
	sdk "agones.dev/agones/sdks/go"
)

var agonesSDK *sdk.SDK

func InitServer(ctx context.Context) {
	var err error
	agonesSDK, err = sdk.NewSDK()
	if err != nil {
		panic(err)
	}

	//启动监听回调
	WatchEvents()

	// 健康上报
	DoHealth(ctx)
}

func GetAgonesSDK() *sdk.SDK {
	return agonesSDK
}

func GetGameServer() *coresdk.GameServer {
	gs, err := agonesSDK.GameServer()
	if err != nil {
		checkError(err)
	}
	return gs
}

func Exit() {
	err := agonesSDK.Shutdown()
	if err != nil {
		checkError(err)
	}
}

func Ready() {
	err := agonesSDK.Ready()
	if err != nil {
		checkError(err)
	}
}

func DoHealth(ctx context.Context) {
	go func() {
		tick := time.Tick(2 * time.Second)
		for {
			err := agonesSDK.Health()
			if err != nil {
				checkError(err)
			}
			select {
			case <-ctx.Done():
				logger.Debug("stop agones health check")
				return
			case <-tick:
			}
		}
	}()
}

func checkError(err error) {
}
