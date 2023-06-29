package handler

import (
	"context"

	"github.com/ribincao/ribin-dev-box/ribin-common/logger"
	serverData "github.com/ribincao/ribin-dev-box/ribin-protocol/server-data"
	"go.uber.org/zap"
)

func (s *EngineServer) HeartBeat(ctx context.Context, ping *serverData.Ping) (*serverData.Pong, error) {
	logger.Info("HeartBeat", zap.Any("Req", ping))
	resp := &serverData.Pong{
		Uid: ping.Uid,
		Seq: ping.Seq,
		Ts:  ping.Ts,
	}
	logger.Info("HeartBeat", zap.Any("Rsp", ping))
	return resp, nil
}
