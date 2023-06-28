package handler

import (
	"context"
	"time"

	"github.com/ribincao/ribin-dev-box/ribin-common/errors"
	"github.com/ribincao/ribin-dev-box/ribin-common/logger"
	"github.com/ribincao/ribin-dev-box/ribin-common/network"
	"github.com/ribincao/ribin-dev-box/ribin-protocol/base"
	"github.com/ribincao/ribin-dev-box/ribin-server/logic"
	"github.com/ribincao/ribin-dev-box/ribin-server/manager"
	"go.uber.org/zap"
)

func HandleServerMessage(ctx context.Context, conn *network.WrapConnection, req *base.Client2ServerReq) (*base.Server2ClientRsp, error) {
	var (
		err *errors.Error
		rsp = &base.Server2ClientRsp{
			Seq: req.Seq,
		}
		rspBody = &base.RspBody{}
	)

	logger.Debug("HandleServerMessage-IN", zap.Any("Req", req))
	switch req.Cmd {
	case base.Client2ServerReqCmd_E_CMD_ROOM_ENTER:
		rspBody, err = handleEnterRoom(ctx, conn, req.Body, req.Seq)
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

func CheckReqParam(req *base.ReqBody) (*logic.NormalRoom, *logic.NormalPlayer, *errors.Error) {
	playerId := req.GetPlayerId()
	if playerId == "" {
		return nil, nil, errors.PlayerIdParamError
	}
	roomId := manager.GetRoomIdByPlayerId(playerId)
	if roomId == "" {
		return nil, nil, errors.RoomIdParamError
	}

	room := manager.GetRoom[*logic.NormalRoom](roomId)
	if room == nil {
		return nil, nil, errors.RoomUnexistError
	}

	player := room.GetPlayer(playerId)
	if player == nil {
		return nil, nil, errors.PlayerNotInRoomError
	}
	return room, player, nil
}

// HeartBeat
func handleHeartBeat(ctx context.Context, conn *network.WrapConnection, heartBeatReq *base.ReqBody, seq string) (*base.RspBody, *errors.Error) {
	var (
		err          *errors.Error
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

func handleEnterRoom(ctx context.Context, conn *network.WrapConnection, enterRoomReq *base.ReqBody, seq string) (*base.RspBody, *errors.Error) {
	logger.Info("HandleEnterRoom-IN", zap.Any("EnterRoomReq", enterRoomReq), zap.String("Seq", seq))
	var (
		err          *errors.Error
		enterRoomRsp = &base.RspBody{
			EnterRoomRsp: &base.EnterRoomRsp{},
		}
	)
	defer func() {
		logger.Info("HandleEnterRoom-OUT", zap.Any("EnterRoomRsp", enterRoomRsp), zap.String("Seq", seq), zap.Error(err))
	}()

	manager.AddRoomToPlayerMap(enterRoomReq.RoomId, enterRoomReq.PlayerId)

	room, _, err := CheckReqParam(enterRoomReq)
	if err == errors.RoomUnexistError {
		roomInfo, err := CreateRoom(enterRoomReq, conn)
		enterRoomRsp.EnterRoomRsp.RoomInfo = roomInfo
		return enterRoomRsp, err
	}
	if err == errors.PlayerNotInRoomError {
		roomInfo, err := JoinRoom(room, enterRoomReq.PlayerId, conn)
		enterRoomRsp.EnterRoomRsp.RoomInfo = roomInfo
		return enterRoomRsp, err
	}
	return enterRoomRsp, err
}

func CreateRoom(enterRoomReq *base.ReqBody, conn *network.WrapConnection) (*base.RoomInfo, *errors.Error) {
	roomId, playerId := enterRoomReq.RoomId, enterRoomReq.PlayerId
	room := logic.NewNormalRoom(roomId)
	manager.RoomMng.AddRoom(room)
	return JoinRoom(room, playerId, conn)
}

func JoinRoom(room *logic.NormalRoom, playerId string, conn *network.WrapConnection) (*base.RoomInfo, *errors.Error) {
	player := logic.NewNormalPlayer(playerId, "TEST")
	player.SetRoomConn(conn)
	room.AddPlayer(player)
	return room.RoomInfo, nil
}
