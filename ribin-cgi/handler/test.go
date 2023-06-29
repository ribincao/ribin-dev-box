package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ribincao/ribin-dev-box/ribin-cgi/pool"
	serverData "github.com/ribincao/ribin-dev-box/ribin-protocol/server-data"
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
		c.JSON(http.StatusInternalServerError, &TestResp{
			RetCode: -2,
			Rmsg:    "rpc error",
			Data:    &TestData{},
		})
		return
	}

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
