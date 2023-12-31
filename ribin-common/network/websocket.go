package network

import (
	"net"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/ribincao/ribin-dev-box/ribin-common/logger"
	"go.uber.org/zap"
)

type Functor struct {
	Func func(code int, text string, args ...interface{}) error
	Args []interface{}
}

type WebSocketConn struct {
	Mu                sync.Mutex
	WebsocketConn     *websocket.Conn
	onCloseFunctors   map[string]*Functor
	onConnectFunctors map[string]*Functor
}

func (c *WebSocketConn) ReadMessage() (messageType int, p []byte, err error) {
	c.WebsocketConn.SetReadDeadline(time.Now().Add(10 * time.Second))
	return c.WebsocketConn.ReadMessage()
}

func (c *WebSocketConn) WriteMessage(messageType int, data []byte) error {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.WebsocketConn.SetWriteDeadline(time.Now().Add(2 * time.Second))
	return c.WebsocketConn.WriteMessage(messageType, data)
}

func (c *WebSocketConn) Close() {
	c.WebsocketConn.Close()
}

func (c *WebSocketConn) RemoteAddr() net.Addr {
	return c.WebsocketConn.RemoteAddr()
}

func (c *WebSocketConn) InitCloseHandler() {
	c.WebsocketConn.SetCloseHandler(c.onClose)
}

func (c *WebSocketConn) onClose(code int, text string) error {
	defaultCloseFunctor := func(code int, text string) error {
		msg := websocket.FormatCloseMessage(code, "")
		c.WebsocketConn.WriteControl(websocket.CloseMessage, msg, time.Now().Add(time.Second))
		return nil
	}

	if c.onCloseFunctors == nil {
		return defaultCloseFunctor(code, text)
	}

	for funcName, functor := range c.onCloseFunctors {
		if functor == nil {
			continue
		}

		err := functor.Func(code, text, functor.Args...)
		if err != nil {
			logger.Error("ConnectionOnCloseError",
				zap.String("FunName", funcName),
				zap.Int("Code", code),
				zap.String("Text", text),
				zap.Error(err))
		}
	}
	return defaultCloseFunctor(code, text)
}

func (c *WebSocketConn) RegisterCloseFunctor(name string, function *Functor) {
	if c.onCloseFunctors == nil {
		c.onCloseFunctors = make(map[string]*Functor)
	}
	if _, ok := c.onCloseFunctors[name]; ok {
		logger.Warn("OnCloseFuntorName Already Exist", zap.String("Name", name))
		return
	}
	c.onCloseFunctors[name] = function
}

func (c *WebSocketConn) onConnect() {
	if c.onConnectFunctors == nil {
		return
	}
	for funcName, functor := range c.onConnectFunctors {
		if functor == nil {
			continue
		}

		err := functor.Func(0, "", functor.Args...)
		if err != nil {
			logger.Error("ConnectionOnConnectError",
				zap.String("FunName", funcName),
				zap.Error(err))
		}
	}
}

func (c *WebSocketConn) RegisterConnectFunctor(name string, function *Functor) {
	if c.onConnectFunctors == nil {
		c.onConnectFunctors = make(map[string]*Functor)
	}
	if _, ok := c.onConnectFunctors[name]; ok {
		logger.Warn("OnConnectFuntorName Already Exist", zap.String("Name", name))
		return
	}
	c.onConnectFunctors[name] = function
}
