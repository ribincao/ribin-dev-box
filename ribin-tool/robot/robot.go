package robot

import (
	"fmt"
	"sync"
	"time"

	"github.com/ribincao/ribin-dev-box/ribin-common/logger"
	"github.com/ribincao/ribin-dev-box/ribin-common/utils"
	"github.com/ribincao/ribin-dev-box/ribin-protocol/base"
	"go.uber.org/atomic"
	"go.uber.org/zap"
)

var robotMng sync.Map

func AddRobot(robotId string) string {
	_, ok := robotMng.Load(robotId)
	if ok {
		logger.Error("RobotId Repeat", zap.String("RobotId", robotId))
		return "[ERROR]RobotId Repeat"
	}
	robot := NewRobot(robotId)
	robotMng.Store(robotId, robot)
	return "[INFO]Robot Add Success."
}

type Robot struct {
	Id         string
	RoomId     string
	wrapconn   *WrapConnection
	Position   *base.Position
	SeqCounter *atomic.Int32
}

func NewRobot(robotId string) *Robot {
	return &Robot{
		Id:         robotId,
		RoomId:     "test",
		SeqCounter: atomic.NewInt32(0),
	}
}

func (r *Robot) EnterRoom(roomId string) {
	r.wrapconn = DialWrapConn(r.Id, roomId)
	r.Position = &base.Position{
		X: 5,
		Y: 2,
		Z: 1,
	}

	enterRoomReq := &base.Client2ServerReq{
		Cmd: base.Client2ServerReqCmd_E_CMD_ROOM_ENTER,
		Seq: fmt.Sprintf("%s-ENTER_ROOM", r.Id),
		Body: &base.ReqBody{
			PlayerId:     r.Id,
			RoomId:       r.RoomId,
			EnterRoomReq: &base.EnterRoomReq{},
		},
	}
	r.wrapconn.SendMessage(enterRoomReq)
	go r.wrapconn.RoomHeartBeat()
	go r.wrapconn.ReadMessage()

	utils.GoWithRecover(func() {
		r.SendFrame()
	})
}

func (r *Robot) GetSeq() string {
	return fmt.Sprintf("%s-%d", r.Id, r.SeqCounter.Add(1))
}

func (r *Robot) GetFrameReq() *base.Client2ServerReq {
	frame := &base.Frame{
		Position: r.Position,
	}
	body := &base.ReqBody{
		SendframeReq: frame,
	}
	return &base.Client2ServerReq{
		Cmd:  base.Client2ServerReqCmd_E_CMD_ROOM_FRAME,
		Seq:  r.GetSeq(),
		Body: body,
	}
}

func (r *Robot) SendFrame() {
	ticker := time.NewTicker(60 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			frameReq := r.GetFrameReq()
			r.wrapconn.SendMessage(frameReq)
		default:
			if r.wrapconn.isClose.Load() {
				r.wrapconn.roomConn.Close()
				return
			}
			time.Sleep(time.Millisecond * 10)
		}
	}
}
