package handler

import (
	"github.com/ribincao/ribin-dev-box/ribin-common/network"
	"github.com/ribincao/ribin-dev-box/ribin-protocol/base"
	"github.com/ribincao/ribin-dev-box/ribin-server/logic"
	"github.com/ribincao/ribin-dev-box/ribin-server/manager"
)

func OnClose(conn *network.WrapConnection) {
	roomId := manager.GetRoomIdByPlayerId(conn.PlayerId)
	if roomId == "" {
		return
	}
	room := manager.GetRoom[*logic.NormalRoom](roomId)
	if room == nil {
		return
	}
	room.RemovePlayer(conn.PlayerId)

	data := &base.BstBody{}
	room.Broadcast(base.Server2ClientBstType_E_PUSH_ROOM_MESSAGE, data, "") // TODO: Broadcast
}
