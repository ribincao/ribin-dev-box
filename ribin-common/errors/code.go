package errors

const (
	TimerTickErrorCode       = 10001
	TimerBucketErrorCode     = 10002
	TimerTaskRepeatErrorCode = 10003
	TimerTaskAddErrorCode    = 10004
)
const (
	ConnectionCloseErrorCode = 20001
)
const (
	MessageDecodeErrorCode = 30001
)
const (
	RoomUnexistErrorCode     = -10001
	PlayerNotInRoomErrorCode = -10002
	ParamsErrorCode          = -10003
)

var (
	PlayerNotInRoomError = New(PlayerNotInRoomErrorCode, "player not in room")
	RoomUnexistError     = New(RoomUnexistErrorCode, "room unexist")
	RoomIdParamError     = New(ParamsErrorCode, "roomid empty")
	PlayerIdParamError   = New(ParamsErrorCode, "playerid empty")
)
