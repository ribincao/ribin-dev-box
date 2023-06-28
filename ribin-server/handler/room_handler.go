package handler

import (
	"context"
	"time"

	errs "github.com/ribincao/ribin-dev-box/ribin-common/errors"
	"github.com/ribincao/ribin-dev-box/ribin-common/logger"
	"github.com/ribincao/ribin-dev-box/ribin-common/network"
	"github.com/ribincao/ribin-dev-box/ribin-protocol/base"
	"github.com/ribincao/ribin-dev-box/ribin-server/logic"
	"github.com/ribincao/ribin-dev-box/ribin-server/manager"
	"go.uber.org/zap"
)

func HandleServerMessage(ctx context.Context, conn *network.WrapConnection, req *base.Client2ServerReq) (*base.Server2ClientRsp, error) {
	var (
		err *errs.Error
		rsp = &base.Server2ClientRsp{
			Seq: req.Seq,
		}
		rspBody = &base.RspBody{}
	)

	logger.Debug("HandleServerMessage-IN", zap.Any("Req", req))
	switch req.Cmd {
	case base.Client2ServerReqCmd_E_CMD_HEART_BEAT:
		rspBody, err = handleHeartBeat(ctx, conn, req.Body, req.Seq)
	}

	if err != nil {
		rsp.Code = err.Code
		rsp.Msg = err.Message
	}
	rsp.Body = rspBody
	logger.Debug("HandleServerMessage-OUT", zap.Any("Rsp", rsp))
	return rsp, err
}

func CheckReqParam(req *base.ReqBody) (*logic.NormalRoom, *logic.NormalPlayer, *errs.Error) {
	playerId := req.GetPlayerId()
	if playerId == "" {
		return nil, nil, errs.New(-1, "invalid uid")
	}
	roomId := manager.GetRoomIdByPlayerId(playerId)
	if roomId == "" {
		return nil, nil, errs.New(-1, "invalid roomid")
	}

	room := manager.GetRoom[*logic.NormalRoom](roomId)
	if room == nil {
		return nil, nil, errs.New(-1, "server not exist")
	}

	player := room.GetPlayer(playerId)
	if player == nil {
		return nil, nil, errs.New(-1, "player not in server")
	}
	return room, player, nil
}

// HeartBeat
func handleHeartBeat(ctx context.Context, conn *network.WrapConnection, heartBeatReq *base.ReqBody, seq string) (*base.RspBody, *errs.Error) {
	var (
		err          *errs.Error
		heartBeatRsp = &base.RspBody{}
	)
	_, player, err := CheckReqParam(heartBeatReq)
	if err != nil {
		return heartBeatRsp, err
	}
	conn.UpdateLastActiveTime(time.Now().UnixMilli())
	player.LastActiveTime = time.Now()
	playerId := player.GetId()
	if playerId == "" {
		conn.PlayerId = playerId
	}
	player.SetRoomConn(conn)
	return heartBeatRsp, err
}
