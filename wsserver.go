package main

import (
	// "io"
	"code.google.com/p/go.net/websocket"
	"github.com/bitly/go-simplejson"
	"log"
)

// Echo the data received on the WebSocket.
func gameServer(ws *websocket.Conn) {

	roomManeger.GetRoomsInfo()

	var err error
	player := NewPlayer(ws) // link ws with Player

	// player.ws = ws
	// need loop to keep socket connect
	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			log.Printf("connect closed!")
			player.Leave()
			break
		}

		js, err := simplejson.NewJson([]byte(reply))
		if err != nil {
			// TODO: Send error json back to client
			log.Printf("parse json error:", err)
			continue
		}

		commandDispatcher(player, js)
	}

}

func Send(ws *websocket.Conn, rtnJson string) {
	if err := websocket.Message.Send(ws, rtnJson); err != nil {
		log.Printf("Send fail for cmLoginHander")
	}
}
