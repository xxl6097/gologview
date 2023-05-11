package util

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type WebSocketUtil struct {
	upgrader websocket.Upgrader
	//conn     *websocket.Conn
	conns map[string]*websocket.Conn
}

func NewWebSocket() WebSocketUtil {
	ww := WebSocketUtil{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		conns: make(map[string]*websocket.Conn),
	}

	return ww
}

func (this *WebSocketUtil) Echo(w http.ResponseWriter, r *http.Request) {
	conn, err := this.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	//this.conn = conn

	this.conns[conn.RemoteAddr().String()] = conn
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("read", err)
			return
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println("write", err)
			return
		}
	}
}

//func (this *WebSocketUtil) Send(data []byte) {
//	//this.conn.WriteMessage(websocket.BinaryMessage, data)
//}

func (this *WebSocketUtil) SendAll(data []byte) {
	for _, v := range this.conns {
		v.WriteMessage(websocket.TextMessage, data)
	}
}
