package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ribincao/ribin-dev-box/ribin-cgi/pool"
	"github.com/ribincao/ribin-dev-box/ribin-common/logger"
	serverData "github.com/ribincao/ribin-dev-box/ribin-protocol/server-data"
	"go.uber.org/zap"
)

type TestData struct {
	Pong *serverData.Pong
}

type TestResp struct {
	RetCode int32     `json:"result"`
	Rmsg    string    `json:"rmsg"`
	Data    *TestData `json:"data"`
}

func TestGet(c *gin.Context) {
	uid := c.Query("uid")
	logger.Info("TestGet Req", zap.Any("Uid", uid))

	segCtx := c.Request.Context()
	ctx, cancel := context.WithTimeout(segCtx, time.Second*2)
	defer cancel()

	req := &serverData.Ping{
		Uid: uid,
		Seq: "test",
		Ts:  time.Now().Unix(),
	}
	client, err := pool.TryAtLeastOnce(ctx, pool.ServerDataPool)
	if err != nil {
		logger.Error("TestGetError -1", zap.Any("Uid", uid))
		c.JSON(http.StatusInternalServerError, &TestResp{
			RetCode: -1,
			Rmsg:    "rpc error",
			Data:    &TestData{},
		})
		return
	}
	defer client.Close()
	conn := serverData.NewServerDataClient(client)
	pong, err := conn.HeartBeat(ctx, req)
	if err != nil {
		logger.Error("TestGetError -2", zap.Any("Uid", uid))
		c.JSON(http.StatusInternalServerError, &TestResp{
			RetCode: -2,
			Rmsg:    "rpc error",
			Data:    &TestData{},
		})
		return
	}

	logger.Info("TestGet Rsp", zap.Any("Uid", uid))
	c.JSON(http.StatusOK, &TestResp{
		RetCode: 0,
		Rmsg:    "",
		Data: &TestData{
			Pong: pong,
		},
	})
}

func TestPost(c *gin.Context) {
}
