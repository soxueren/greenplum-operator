package ws

import (
	"errors"
	"sync"

	"github.com/gorilla/websocket"
)

type Connection struct {
	WSConn    *websocket.Conn
	InChan    chan []byte
	OutChan   chan []byte
	Mutex     sync.Mutex
	CloseChan chan byte
	Isclose   bool
}

func InitConnection(c *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		WSConn:    c,
		InChan:    make(chan []byte, 1024),
		OutChan:   make(chan []byte, 1024),
		CloseChan: make(chan byte, 1),
		Isclose:   false,
	}

	go conn.ReadLoop()

	go conn.WriteLoop()

	return
}

func (conn *Connection) Close() {
	conn.Mutex.Lock()
	conn.WSConn.Close()
	if !conn.Isclose {
		close(conn.CloseChan)
		conn.Isclose = true
	}
	conn.Mutex.Unlock()
	return
}

func (conn *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.InChan:
	case <-conn.CloseChan:
		err = errors.New("connection is closed")
	}
	return
}

func (conn *Connection) WriteMessage(data []byte) (err error) {
	select {
	case conn.OutChan <- data:
	case <-conn.CloseChan:
		err = errors.New("connection is closed")
	}
	return
}

func (conn *Connection) ReadLoop() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = conn.WSConn.ReadMessage(); err != nil {
			goto ERR
		}
		select {
		case conn.InChan <- data:
		case <-conn.CloseChan:
			goto ERR
		}
	}

ERR:
	conn.Close()
}

func (conn *Connection) WriteLoop() {
	var (
		data []byte
		err  error
	)
	for {
		select {
		case data = <-conn.OutChan:
		case <-conn.CloseChan:
			goto ERR
		}
		if err = conn.WSConn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}
