package handler

import (
	"archive/zip"
	"context"
	"encoding/json"
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
	client, err := pool.GetRpcClient(ctx, pool.ServerDataPool)
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

type Person struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func handleDownloadResource(c *gin.Context) {
	// 设置响应头
	c.Header("Content-Type", "application/ictet-stream")
	c.Header("Content-Disposition", "attachment;filename=resource.zip")

	// 创建 ZIP 编写器
	zipWriter := zip.NewWriter(c.Writer)
	defer zipWriter.Close()

	// 将每个人的数据写入单独的 JSON 文件，并添加到 ZIP 文件
	var people []Person
	people = append(people, Person{
		ID:   1,
		Name: "ribin",
		Age:  28,
	})
	for _, person := range people {
		// 创建 JSON 文件
		jsonFileName := person.Name + ".json"
		jsonFile, err := zipWriter.Create(jsonFileName)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to create JSON file")
			return
		}

		// 将数据写入 JSON 文件
		encoder := json.NewEncoder(jsonFile)
		if err := encoder.Encode(person); err != nil {
			c.String(http.StatusInternalServerError, "Failed to write data to JSON file")
			return
		}
	}
}
