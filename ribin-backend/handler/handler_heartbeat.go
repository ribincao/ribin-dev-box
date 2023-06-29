package handler

import (
	"context"

	serverData "github.com/ribincao/ribin-dev-box/ribin-protocol/server-data"
)

func (s *EngineServer) HeartBeat(ctx context.Context, ping *serverData.Ping) (*serverData.Pong, error) {
	resp := &serverData.Pong{
		Uid: ping.Uid,
		Seq: ping.Seq,
		Ts:  ping.Ts,
	}
	return resp, nil
}
