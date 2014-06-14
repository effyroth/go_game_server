package main

import (
	// "io"
	"code.google.com/p/go.net/websocket"
	"github.com/bitly/go-simplejson"
	"log"
)

// Echo the data received on the WebSocket.
func gameServer(ws *websocket.Conn) {

	var err error
	var this Player // link ws with Player

	this.ws = ws
	// need loop to keep socket connect
	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			log.Printf("connect closed!")
			break
		}

		js, err := simplejson.NewJson([]byte(reply))
		if err != nil {
			// TODO: Send error json back to client
			log.Printf("parse json error:", err)
			continue
		}

		commandDispatcher(&this, js)
	}

}
